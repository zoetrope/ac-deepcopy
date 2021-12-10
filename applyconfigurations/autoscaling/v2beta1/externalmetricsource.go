/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v2beta1

import (
	resource "k8s.io/apimachinery/pkg/api/resource"
	v1 "github.com/zoetrope/ac-deepcopy/applyconfigurations/meta/v1"
)

// ExternalMetricSourceApplyConfiguration represents an declarative configuration of the ExternalMetricSource type for use
// with apply.
type ExternalMetricSourceApplyConfiguration struct {
	MetricName         *string                             `json:"metricName,omitempty"`
	MetricSelector     *v1.LabelSelectorApplyConfiguration `json:"metricSelector,omitempty"`
	TargetValue        *resource.Quantity                  `json:"targetValue,omitempty"`
	TargetAverageValue *resource.Quantity                  `json:"targetAverageValue,omitempty"`
}

// ExternalMetricSourceApplyConfiguration constructs an declarative configuration of the ExternalMetricSource type for use with
// apply.
func ExternalMetricSource() *ExternalMetricSourceApplyConfiguration {
	return &ExternalMetricSourceApplyConfiguration{}
}

// WithMetricName sets the MetricName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MetricName field is set to the value of the last call.
func (b *ExternalMetricSourceApplyConfiguration) WithMetricName(value string) *ExternalMetricSourceApplyConfiguration {
	b.MetricName = &value
	return b
}

// WithMetricSelector sets the MetricSelector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MetricSelector field is set to the value of the last call.
func (b *ExternalMetricSourceApplyConfiguration) WithMetricSelector(value *v1.LabelSelectorApplyConfiguration) *ExternalMetricSourceApplyConfiguration {
	b.MetricSelector = value
	return b
}

// WithTargetValue sets the TargetValue field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TargetValue field is set to the value of the last call.
func (b *ExternalMetricSourceApplyConfiguration) WithTargetValue(value resource.Quantity) *ExternalMetricSourceApplyConfiguration {
	b.TargetValue = &value
	return b
}

// WithTargetAverageValue sets the TargetAverageValue field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TargetAverageValue field is set to the value of the last call.
func (b *ExternalMetricSourceApplyConfiguration) WithTargetAverageValue(value resource.Quantity) *ExternalMetricSourceApplyConfiguration {
	b.TargetAverageValue = &value
	return b
}
