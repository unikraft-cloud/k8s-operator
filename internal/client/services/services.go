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

package services

import (
	"context"

	"unikraft.com/cloud/sdk/pkg/httpclient"
	ukcplatform "unikraft.com/cloud/sdk/platform"

	unikraftv1alpha1 "github.com/unikraft-cloud/k8s-operator/api/v1alpha1"
	"github.com/unikraft-cloud/k8s-operator/api/v1alpha1/platform"
	"github.com/unikraft-cloud/k8s-operator/internal/client"
)

type Client struct {
	client ukcplatform.Client
}

func NewClient(metro, token string) *Client {
	return &Client{
		client: ukcplatform.NewClient(
			ukcplatform.WithToken(token),
			ukcplatform.WithHTTPClient(httpclient.NewHTTPClient()),
			ukcplatform.WithDefaultMetro(metro),
		),
	}
}

func (c *Client) ResourceExists(ctx context.Context, obj *unikraftv1alpha1.Service) (bool, error) {
	resp, err := c.client.GetServiceGroups(ctx, []ukcplatform.NameOrUUID{{Name: obj.Spec.Name}}, false)
	if err != nil {
		if ukcplatform.ErrorContainsOnly(err, ukcplatform.APIHTTPErrorNotFound) {
			return false, nil
		}
		return false, err
	}

	if resp == nil || resp.Data == nil || len(resp.Data.ServiceGroups) != 1 {
		return false, nil
	}

	return true, nil
}

func (c *Client) CreateResource(ctx context.Context, obj *unikraftv1alpha1.Service) (*platform.CreateServiceGroupResponse, error) {
	req, err := client.Convert[platform.CreateServiceGroupRequest, ukcplatform.CreateServiceGroupRequest](obj.Spec)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.CreateServiceGroup(ctx, req)
	if err != nil && !ukcplatform.ErrorContainsOnly(err, ukcplatform.APIHTTPErrorAlreadyExists) {
		return &platform.CreateServiceGroupResponse{
			Status:  ukcplatform.Ptr(platform.ResponseStatus(ukcplatform.ResponseStatusERROR)),
			Message: ukcplatform.Ptr(err.Error()),
			Data:    obj.Status.Data,
		}, err
	}

	return client.Convert[*ukcplatform.Response[ukcplatform.CreateServiceGroupResponseData], *platform.CreateServiceGroupResponse](resp)
}

func (c *Client) UpdateResource(ctx context.Context, obj *unikraftv1alpha1.Service) (*platform.CreateServiceGroupResponse, error) {
	// TODO(petar-cvit): implement updates
	return nil, nil
}

func (c *Client) DeleteResource(ctx context.Context, obj *unikraftv1alpha1.Service) (*platform.CreateServiceGroupResponse, error) {
	if obj.Status.Data == nil || len(obj.Status.Data.ServiceGroups) == 0 {
		return nil, nil
	}

	serviceNames := make([]ukcplatform.NameOrUUID, 0)
	for _, service := range obj.Status.Data.ServiceGroups {
		if service.Name != nil {
			serviceNames = append(serviceNames, ukcplatform.NameOrUUID{
				Name: service.Name,
			})
		}
	}

	resp, err := c.client.DeleteServiceGroups(ctx, serviceNames)
	if err != nil {
		return &platform.CreateServiceGroupResponse{
			Status:  ukcplatform.Ptr(platform.ResponseStatus(ukcplatform.ResponseStatusERROR)),
			Message: ukcplatform.Ptr(err.Error()),
			Data:    obj.Status.Data,
		}, err
	}

	return client.Convert[*ukcplatform.Response[ukcplatform.DeleteServiceGroupsResponseData], *platform.CreateServiceGroupResponse](resp)
}
