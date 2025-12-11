// SPDX-License-Identifier: UNLICENSED
//
// Copyright (c) 2025, Unikraft GmbH.  All rights reserved.
//
// This software and related documentation ("Unikraft Software") are protected
// under relevant copyright laws.  The information contained herein is
// confidential and proprietary to Unikraft GmbH ("Unikraft") and/or its
// licensors.  Without the prior written permission of Unikraft GmbH and/or its
// licensors, any reproduction, modification, use or disclosure of Unikraft
// Software, and information contained herein, in whole or in part, shall be
// strictly prohibited.
//
// BY OPENING THIS FILE, RECEIVER HEREBY UNEQUIVOCALLY ACKNOWLEDGES AND AGREES
// THAT THE SOFTWARE/FIRMWARE AND ITS DOCUMENTATIONS ("UNIKRAFT SOFTWARE")
// RECEIVED FROM UNIKRAFT AND/OR ITS REPRESENTATIVES ARE PROVIDED TO RECEIVER ON
// AN "AS-IS" BASIS ONLY.  UNIKRAFT EXPRESSLY DISCLAIMS ANY AND ALL WARRANTIES,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE IMPLIED WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE OR NONINFRINGEMENT. NEITHER
// DOES UNIKRAFT PROVIDE ANY WARRANTY WHATSOEVER WITH RESPECT TO THE SOFTWARE OF
// ANY THIRD PARTY WHICH MAY BE USED BY, INCORPORATED IN, OR SUPPLIED WITH THE
// UNIKRAFT SOFTWARE, AND RECEIVER AGREES TO LOOK ONLY TO SUCH THIRD PARTY FOR
// ANY WARRANTY CLAIM RELATING THERETO.  RECEIVER EXPRESSLY ACKNOWLEDGES THAT IT
// IS RECEIVER'S SOLE RESPONSIBILITY TO OBTAIN FROM ANY THIRD PARTY ALL PROPER
// LICENSES CONTAINED IN UNIKRAFT SOFTWARE.  UNIKRAFT SHALL ALSO NOT BE
// RESPONSIBLE FOR ANY UNIKRAFT SOFTWARE RELEASES MADE TO RECEIVER'S
// SPECIFICATION OR TO CONFORM TO A PARTICULAR STANDARD OR OPEN FORUM.
// RECEIVER'S SOLE AND EXCLUSIVE REMEDY AND UNIKRAFT'S ENTIRE AND CUMULATIVE
// LIABILITY WITH RESPECT TO THE UNIKRAFT SOFTWARE RELEASED HEREUNDER WILL BE,
// AT UNIKRAFT'S OPTION, TO REVISE OR REPLACE THE UNIKRAFT SOFTWARE AT ISSUE, OR
// REFUND ANY SOFTWARE LICENSE FEES OR SERVICE CHARGE PAID BY RECEIVER TO
// UNIKRAFT FOR SUCH UNIKRAFT SOFTWARE AT ISSUE.

package controller

import (
	"context"
	"fmt"
	"time"

	"github.com/go-logr/logr"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	ukcclient "github.com/unikraft-cloud/k8s-operator/internal/client"
)

const resourceFinalizer = "cloud.unikraft.v1/resource"

type Reconciler[CRD client.Object, SpecT any, StatusT any] struct {
	client.Client
	Scheme    *runtime.Scheme
	GetSpec   func(obj CRD) *SpecT
	SetStatus func(obj CRD, status *StatusT)
	NewObject func() CRD

	ResourceClient ukcclient.ResourceClient[CRD, SpecT, StatusT]

	Logger logr.Logger
}

func (r *Reconciler[CRD, SpecT, StatusT]) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	if r.NewObject == nil {
		return reconcile.Result{}, fmt.Errorf("Reconciler.NewObject must be provided")
	}

	obj := r.NewObject()
	if err := r.Get(ctx, req.NamespacedName, obj); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if obj.GetDeletionTimestamp() != nil {
		r.Logger.Info("deleting", "namespacedName", req.NamespacedName)

		if controllerutil.ContainsFinalizer(obj, resourceFinalizer) {
			status, err := r.ResourceClient.DeleteResource(ctx, obj)
			if err != nil {
				r.SetStatus(obj, status)
				_ = r.Status().Update(ctx, obj)
				return reconcile.Result{}, err
			}

			controllerutil.RemoveFinalizer(obj, resourceFinalizer)
			err = r.Update(ctx, obj)
			if err != nil {
				return ctrl.Result{}, err
			}
		}
		return ctrl.Result{}, nil
	}

	if !controllerutil.ContainsFinalizer(obj, resourceFinalizer) {
		controllerutil.AddFinalizer(obj, resourceFinalizer)
		err := r.Update(ctx, obj)
		if err != nil {
			return ctrl.Result{}, err
		}
	}

	exists, _, err := r.ResourceClient.ResourceExists(ctx, obj)
	if err != nil {
		return ctrl.Result{}, err
	}

	if exists {
		// TODO(petar-cvit): implement resource updates
		return ctrl.Result{}, nil
	}

	r.Logger.Info("creating", "namespacedName", req.NamespacedName)

	status, err := r.ResourceClient.CreateResource(ctx, obj)

	r.SetStatus(obj, status)
	if err := r.Status().Update(ctx, obj); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, err
}

func (r *Reconciler[CRD, SpecT, StatusT]) SetupWithManager(mgr ctrl.Manager) error {
	if r.NewObject == nil {
		return fmt.Errorf("Reconciler.NewObject must be provided")
	}

	rateLimiter := workqueue.NewMaxOfRateLimiter(
		workqueue.NewItemExponentialFailureRateLimiter(1*time.Second, 64*time.Second),
	)

	return ctrl.NewControllerManagedBy(mgr).
		For(r.NewObject()).
		WithOptions(controller.Options{
			RateLimiter: rateLimiter,
		}).
		Complete(r)
}
