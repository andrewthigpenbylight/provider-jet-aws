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

type ResolverRuleObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	ShareStatus *string `json:"shareStatus,omitempty" tf:"share_status,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ResolverRuleParameters struct {

	// +kubebuilder:validation:Required
	DomainName *string `json:"domainName" tf:"domain_name,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	ResolverEndpointID *string `json:"resolverEndpointId,omitempty" tf:"resolver_endpoint_id,omitempty"`

	// +kubebuilder:validation:Required
	RuleType *string `json:"ruleType" tf:"rule_type,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	TargetIP []TargetIPParameters `json:"targetIp,omitempty" tf:"target_ip,omitempty"`
}

type TargetIPObservation struct {
}

type TargetIPParameters struct {

	// +kubebuilder:validation:Required
	IP *string `json:"ip" tf:"ip,omitempty"`

	// +kubebuilder:validation:Optional
	Port *int64 `json:"port,omitempty" tf:"port,omitempty"`
}

// ResolverRuleSpec defines the desired state of ResolverRule
type ResolverRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ResolverRuleParameters `json:"forProvider"`
}

// ResolverRuleStatus defines the observed state of ResolverRule.
type ResolverRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ResolverRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ResolverRule is the Schema for the ResolverRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type ResolverRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResolverRuleSpec   `json:"spec"`
	Status            ResolverRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResolverRuleList contains a list of ResolverRules
type ResolverRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResolverRule `json:"items"`
}

// Repository type metadata.
var (
	ResolverRule_Kind             = "ResolverRule"
	ResolverRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ResolverRule_Kind}.String()
	ResolverRule_KindAPIVersion   = ResolverRule_Kind + "." + CRDGroupVersion.String()
	ResolverRule_GroupVersionKind = CRDGroupVersion.WithKind(ResolverRule_Kind)
)

func init() {
	SchemeBuilder.Register(&ResolverRule{}, &ResolverRuleList{})
}
