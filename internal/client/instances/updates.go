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

package instances

import (
	"reflect"

	ukcplatform "unikraft.com/cloud/sdk/platform"

	unikraftv1alpha1 "github.com/unikraft-cloud/k8s-operator/api/v1alpha1"
	"github.com/unikraft-cloud/k8s-operator/api/v1alpha1/platform"
)

const (
	defaultMemoryMB = 128
	defaultVCPU     = 1
)

func instanceUpdates(curr ukcplatform.Instance, obj *unikraftv1alpha1.Instance) []ukcplatform.UpdateInstancesRequestItem {
	if obj == nil || obj.Status.Data == nil {
		return nil
	}

	instances := obj.Status.Data.Instances
	var updates []ukcplatform.UpdateInstancesRequestItem

	addUpdate := func(prop ukcplatform.UpdateInstancesRequestItemProp, op ukcplatform.UpdateInstancesRequestItemOp, value interface{}) {
		for _, instance := range instances {
			updates = append(updates, ukcplatform.UpdateInstancesRequestItem{
				Name:  instance.Name,
				Prop:  prop,
				Op:    op,
				Value: ukcplatform.Ptr(value),
			})
		}
	}

	if !reflect.DeepEqual(curr.Args, obj.Spec.Args) {
		addUpdate(ukcplatform.UpdateInstancesRequestItemPropArgs, ukcplatform.UpdateInstancesRequestItemOpSet, obj.Spec.Args)
	}

	if !reflect.DeepEqual(curr.Env, obj.Spec.Env) {
		addUpdate(ukcplatform.UpdateInstancesRequestItemPropEnv, ukcplatform.UpdateInstancesRequestItemOpSet, obj.Spec.Env)
	}

	if !ptrValuesEqualWithDefault(curr.MemoryMb, obj.Spec.MemoryMb, defaultMemoryMB) {
		addUpdate(ukcplatform.UpdateInstancesRequestItemPropMemory_mb, ukcplatform.UpdateInstancesRequestItemOpSet, obj.Spec.MemoryMb)
	}

	if !ptrValuesEqualWithDefault(curr.Vcpus, obj.Spec.Vcpus, defaultVCPU) {
		addUpdate(ukcplatform.UpdateInstancesRequestItemPropVcpus, ukcplatform.UpdateInstancesRequestItemOpSet, obj.Spec.Vcpus)
	}

	if !scaleToZeroEqual(curr.ScaleToZero, obj.Spec.ScaleToZero) {
		addUpdate(ukcplatform.UpdateInstancesRequestItemPropScale_to_zero, ukcplatform.UpdateInstancesRequestItemOpSet, obj.Spec.ScaleToZero)
	}

	return updates
}

func scaleToZeroEqual(curr *ukcplatform.InstanceScaleToZero, spec *platform.CreateInstanceRequestScaleToZero) bool {
	if curr == nil && spec == nil {
		return true
	}
	if curr == nil || spec == nil {
		return false
	}

	if (curr.Policy == nil) != (spec.Policy == nil) {
		return false
	}
	if curr.Policy != nil && spec.Policy != nil {
		if string(*curr.Policy) != string(*spec.Policy) {
			return false
		}
	}

	if !reflect.DeepEqual(curr.Stateful, spec.Stateful) {
		return false
	}

	if !reflect.DeepEqual(curr.CooldownTimeMs, spec.CooldownTimeMs) {
		return false
	}

	return true
}

func ptrValuesEqualWithDefault[T1, T2 int32 | int64 | uint32 | uint64](a *T1, b *T2, defaultVal int64) bool {
	aVal := defaultVal
	bVal := defaultVal
	if a != nil {
		aVal = int64(*a)
	}
	if b != nil {
		bVal = int64(*b)
	}
	return aVal == bVal
}
