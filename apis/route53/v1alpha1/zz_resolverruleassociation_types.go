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

type ResolverRuleAssociationObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ResolverRuleAssociationParameters struct {

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	ResolverRuleID *string `json:"resolverRuleId" tf:"resolver_rule_id,omitempty"`

	// +kubebuilder:validation:Required
	VPCID *string `json:"vpcId" tf:"vpc_id,omitempty"`
}

// ResolverRuleAssociationSpec defines the desired state of ResolverRuleAssociation
type ResolverRuleAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ResolverRuleAssociationParameters `json:"forProvider"`
}

// ResolverRuleAssociationStatus defines the observed state of ResolverRuleAssociation.
type ResolverRuleAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ResolverRuleAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ResolverRuleAssociation is the Schema for the ResolverRuleAssociations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type ResolverRuleAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResolverRuleAssociationSpec   `json:"spec"`
	Status            ResolverRuleAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResolverRuleAssociationList contains a list of ResolverRuleAssociations
type ResolverRuleAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResolverRuleAssociation `json:"items"`
}

// Repository type metadata.
var (
	ResolverRuleAssociation_Kind             = "ResolverRuleAssociation"
	ResolverRuleAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ResolverRuleAssociation_Kind}.String()
	ResolverRuleAssociation_KindAPIVersion   = ResolverRuleAssociation_Kind + "." + CRDGroupVersion.String()
	ResolverRuleAssociation_GroupVersionKind = CRDGroupVersion.WithKind(ResolverRuleAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&ResolverRuleAssociation{}, &ResolverRuleAssociationList{})
}
