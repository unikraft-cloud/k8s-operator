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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/unikraft-cloud/k8s-operator/api/v1alpha1/platform"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="Image",type="string",JSONPath=".spec.image"
// +kubebuilder:printcolumn:name="Sync status",type="string",JSONPath=".status.status"

// Instance is the Schema for the instances API
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +kubebuilder:validation:XValidation:rule="(!has(oldSelf.name) && !has(self.name)) || (has(oldSelf.name) && has(self.name) && self.name == oldSelf.name)",message="spec.name is immutable"
	// +kubebuilder:validation:XValidation:rule="(!has(oldSelf.service_group) && !has(self.service_group)) || (has(oldSelf.service_group) && has(self.service_group) && self.service_group == oldSelf.service_group)",message="spec.service_group is immutable"
	// +kubebuilder:validation:XValidation:rule="(!has(oldSelf.volumes) && !has(self.volumes)) || (has(oldSelf.volumes) && has(self.volumes) && self.volumes == oldSelf.volumes)",message="spec.volumes is immutable"
	// +kubebuilder:validation:XValidation:rule="(!has(oldSelf.autostart) && !has(self.autostart)) || (has(oldSelf.autostart) && has(self.autostart) && self.autostart == oldSelf.autostart)",message="spec.autostart is immutable"
	// +kubebuilder:validation:XValidation:rule="(!has(oldSelf.replicas) && !has(self.replicas)) || (has(oldSelf.replicas) && has(self.replicas) && self.replicas == oldSelf.replicas)",message="spec.replicas is immutable"
	// +kubebuilder:validation:XValidation:rule="(!has(oldSelf.restart_policy) && !has(self.restart_policy)) || (has(oldSelf.restart_policy) && has(self.restart_policy) && self.restart_policy == oldSelf.restart_policy)",message="spec.restart_policy is immutable"
	// +kubebuilder:validation:XValidation:rule="(!has(oldSelf.wait_timeout_ms) && !has(self.wait_timeout_ms)) || (has(oldSelf.wait_timeout_ms) && has(self.wait_timeout_ms) && self.wait_timeout_ms == oldSelf.wait_timeout_ms)",message="spec.wait_timeout_ms is immutable"
	// +kubebuilder:validation:XValidation:rule="(!has(oldSelf.features) && !has(self.features)) || (has(oldSelf.features) && has(self.features) && self.features == oldSelf.features)",message="spec.features is immutable"
	Spec   platform.CreateInstanceRequest `json:"spec,omitempty"`
	Status platform.GetInstancesResponse  `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceList contains a list of Instance
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Instance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Instance{}, &InstanceList{})
}
