/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ResolverFirewallRuleObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ResolverFirewallRuleParameters struct {

	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// +kubebuilder:validation:Optional
	BlockOverrideDNSType *string `json:"blockOverrideDnsType,omitempty" tf:"block_override_dns_type,omitempty"`

	// +kubebuilder:validation:Optional
	BlockOverrideDomain *string `json:"blockOverrideDomain,omitempty" tf:"block_override_domain,omitempty"`

	// +kubebuilder:validation:Optional
	BlockOverrideTTL *int64 `json:"blockOverrideTtl,omitempty" tf:"block_override_ttl,omitempty"`

	// +kubebuilder:validation:Optional
	BlockResponse *string `json:"blockResponse,omitempty" tf:"block_response,omitempty"`

	// +kubebuilder:validation:Required
	FirewallDomainListID *string `json:"firewallDomainListId" tf:"firewall_domain_list_id,omitempty"`

	// +kubebuilder:validation:Required
	FirewallRuleGroupID *string `json:"firewallRuleGroupId" tf:"firewall_rule_group_id,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Priority *int64 `json:"priority" tf:"priority,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// ResolverFirewallRuleSpec defines the desired state of ResolverFirewallRule
type ResolverFirewallRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ResolverFirewallRuleParameters `json:"forProvider"`
}

// ResolverFirewallRuleStatus defines the observed state of ResolverFirewallRule.
type ResolverFirewallRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ResolverFirewallRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ResolverFirewallRule is the Schema for the ResolverFirewallRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type ResolverFirewallRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResolverFirewallRuleSpec   `json:"spec"`
	Status            ResolverFirewallRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResolverFirewallRuleList contains a list of ResolverFirewallRules
type ResolverFirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResolverFirewallRule `json:"items"`
}

// Repository type metadata.
var (
	ResolverFirewallRule_Kind             = "ResolverFirewallRule"
	ResolverFirewallRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ResolverFirewallRule_Kind}.String()
	ResolverFirewallRule_KindAPIVersion   = ResolverFirewallRule_Kind + "." + CRDGroupVersion.String()
	ResolverFirewallRule_GroupVersionKind = CRDGroupVersion.WithKind(ResolverFirewallRule_Kind)
)

func init() {
	SchemeBuilder.Register(&ResolverFirewallRule{}, &ResolverFirewallRuleList{})
}
