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

type ResolverFirewallDomainListObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ResolverFirewallDomainListParameters struct {

	// +kubebuilder:validation:Optional
	Domains []*string `json:"domains,omitempty" tf:"domains,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ResolverFirewallDomainListSpec defines the desired state of ResolverFirewallDomainList
type ResolverFirewallDomainListSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ResolverFirewallDomainListParameters `json:"forProvider"`
}

// ResolverFirewallDomainListStatus defines the observed state of ResolverFirewallDomainList.
type ResolverFirewallDomainListStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ResolverFirewallDomainListObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ResolverFirewallDomainList is the Schema for the ResolverFirewallDomainLists API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type ResolverFirewallDomainList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResolverFirewallDomainListSpec   `json:"spec"`
	Status            ResolverFirewallDomainListStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResolverFirewallDomainListList contains a list of ResolverFirewallDomainLists
type ResolverFirewallDomainListList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResolverFirewallDomainList `json:"items"`
}

// Repository type metadata.
var (
	ResolverFirewallDomainList_Kind             = "ResolverFirewallDomainList"
	ResolverFirewallDomainList_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ResolverFirewallDomainList_Kind}.String()
	ResolverFirewallDomainList_KindAPIVersion   = ResolverFirewallDomainList_Kind + "." + CRDGroupVersion.String()
	ResolverFirewallDomainList_GroupVersionKind = CRDGroupVersion.WithKind(ResolverFirewallDomainList_Kind)
)

func init() {
	SchemeBuilder.Register(&ResolverFirewallDomainList{}, &ResolverFirewallDomainListList{})
}
