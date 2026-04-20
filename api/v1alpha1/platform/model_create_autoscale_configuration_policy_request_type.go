// This file is auto-generated. DO NOT EDIT.
// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2025, Unikraft GmbH.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

// +k8s:deepcopy-gen=package
package platform

// The policy type to add to the autoscale configuration.
// Metric to use for the step policy.
// +kubebuilder:validation:Enum=cpu
type CreateAutoscaleConfigurationPolicyRequestTypeMetric string

const (
	CreateAutoscaleConfigurationPolicyRequestTypeMetricCpu CreateAutoscaleConfigurationPolicyRequestTypeMetric = "cpu"
)

// The type of adjustment to be made in the step policy.
// +kubebuilder:validation:Enum=change;exact;percentage
type CreateAutoscaleConfigurationPolicyRequestTypeAdjustmentType string

const (
	CreateAutoscaleConfigurationPolicyRequestTypeAdjustmentTypeChange     CreateAutoscaleConfigurationPolicyRequestTypeAdjustmentType = "change"
	CreateAutoscaleConfigurationPolicyRequestTypeAdjustmentTypeExact      CreateAutoscaleConfigurationPolicyRequestTypeAdjustmentType = "exact"
	CreateAutoscaleConfigurationPolicyRequestTypeAdjustmentTypePercentage CreateAutoscaleConfigurationPolicyRequestTypeAdjustmentType = "percentage"
)

type CreateAutoscaleConfigurationPolicyRequestType struct {
	// The name of the policy.
	Name *string `json:"name,omitempty"`
	// (Only applies when using global control plane).
	// The metro of the service group the policy applies to.
	Metro *string `json:"metro,omitempty"`
	// If the policy is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Metric to use for the step policy.
	Metric *CreateAutoscaleConfigurationPolicyRequestTypeMetric `json:"metric,omitempty"`
	// The type of adjustment to be made in the step policy.
	AdjustmentType *CreateAutoscaleConfigurationPolicyRequestTypeAdjustmentType `json:"adjustment_type,omitempty"`
	// The steps for the step policy.
	// Each step defines an adjustment value and optional bounds.
	Steps []AutoscalePolicyStep `json:"steps,omitempty"`
}
