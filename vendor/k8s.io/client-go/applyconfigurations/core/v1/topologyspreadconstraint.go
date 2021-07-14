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

package v1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// TopologySpreadConstraintApplyConfiguration represents an declarative configuration of the TopologySpreadConstraint type for use
// with apply.
type TopologySpreadConstraintApplyConfiguration struct {
	MaxSkew           *int32                                  `json:"maxSkew,omitempty"`
	TopologyKey       *string                                 `json:"topologyKey,omitempty"`
	WhenUnsatisfiable *v1.UnsatisfiableConstraintAction       `json:"whenUnsatisfiable,omitempty"`
	LabelSelector     *metav1.LabelSelectorApplyConfiguration `json:"labelSelector,omitempty"`
}

// TopologySpreadConstraintApplyConfiguration constructs an declarative configuration of the TopologySpreadConstraint type for use with
// apply.
func TopologySpreadConstraint() *TopologySpreadConstraintApplyConfiguration {
	return &TopologySpreadConstraintApplyConfiguration{}
}

// WithMaxSkew sets the MaxSkew field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MaxSkew field is set to the value of the last call.
func (b *TopologySpreadConstraintApplyConfiguration) WithMaxSkew(value int32) *TopologySpreadConstraintApplyConfiguration {
	b.MaxSkew = &value
	return b
}

// WithTopologyKey sets the TopologyKey field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TopologyKey field is set to the value of the last call.
func (b *TopologySpreadConstraintApplyConfiguration) WithTopologyKey(value string) *TopologySpreadConstraintApplyConfiguration {
	b.TopologyKey = &value
	return b
}

// WithWhenUnsatisfiable sets the WhenUnsatisfiable field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the WhenUnsatisfiable field is set to the value of the last call.
func (b *TopologySpreadConstraintApplyConfiguration) WithWhenUnsatisfiable(value v1.UnsatisfiableConstraintAction) *TopologySpreadConstraintApplyConfiguration {
	b.WhenUnsatisfiable = &value
	return b
}

// WithLabelSelector sets the LabelSelector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LabelSelector field is set to the value of the last call.
func (b *TopologySpreadConstraintApplyConfiguration) WithLabelSelector(value *metav1.LabelSelectorApplyConfiguration) *TopologySpreadConstraintApplyConfiguration {
	b.LabelSelector = value
	return b
}
