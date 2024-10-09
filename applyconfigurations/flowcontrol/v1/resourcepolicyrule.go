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

// ResourcePolicyRuleApplyConfiguration represents an declarative configuration of the ResourcePolicyRule type for use
// with apply.
type ResourcePolicyRuleApplyConfiguration struct {
	Verbs        []string `json:"verbs,omitempty"`
	APIGroups    []string `json:"apiGroups,omitempty"`
	Resources    []string `json:"resources,omitempty"`
	ClusterScope *bool    `json:"clusterScope,omitempty"`
	Namespaces   []string `json:"namespaces,omitempty"`
}

// ResourcePolicyRuleApplyConfiguration constructs an declarative configuration of the ResourcePolicyRule type for use with
// apply.
func ResourcePolicyRule() *ResourcePolicyRuleApplyConfiguration {
	return &ResourcePolicyRuleApplyConfiguration{}
}

// WithVerbs adds the given value to the Verbs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Verbs field.
func (b *ResourcePolicyRuleApplyConfiguration) WithVerbs(values ...string) *ResourcePolicyRuleApplyConfiguration {
	for i := range values {
		b.Verbs = append(b.Verbs, values[i])
	}
	return b
}

// WithAPIGroups adds the given value to the APIGroups field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the APIGroups field.
func (b *ResourcePolicyRuleApplyConfiguration) WithAPIGroups(values ...string) *ResourcePolicyRuleApplyConfiguration {
	for i := range values {
		b.APIGroups = append(b.APIGroups, values[i])
	}
	return b
}

// WithResources adds the given value to the Resources field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Resources field.
func (b *ResourcePolicyRuleApplyConfiguration) WithResources(values ...string) *ResourcePolicyRuleApplyConfiguration {
	for i := range values {
		b.Resources = append(b.Resources, values[i])
	}
	return b
}

// WithClusterScope sets the ClusterScope field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ClusterScope field is set to the value of the last call.
func (b *ResourcePolicyRuleApplyConfiguration) WithClusterScope(value bool) *ResourcePolicyRuleApplyConfiguration {
	b.ClusterScope = &value
	return b
}

// WithNamespaces adds the given value to the Namespaces field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Namespaces field.
func (b *ResourcePolicyRuleApplyConfiguration) WithNamespaces(values ...string) *ResourcePolicyRuleApplyConfiguration {
	for i := range values {
		b.Namespaces = append(b.Namespaces, values[i])
	}
	return b
}
