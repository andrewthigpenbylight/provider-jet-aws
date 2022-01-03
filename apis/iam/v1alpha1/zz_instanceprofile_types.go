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

type InstanceProfileObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	CreateDate *string `json:"createDate,omitempty" tf:"create_date,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	UniqueID *string `json:"uniqueId,omitempty" tf:"unique_id,omitempty"`
}

type InstanceProfileParameters struct {

	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// +crossplane:generate:reference:type=Role
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// +kubebuilder:validation:Optional
	RoleRef *v1.Reference `json:"roleRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	RoleSelector *v1.Selector `json:"roleSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// InstanceProfileSpec defines the desired state of InstanceProfile
type InstanceProfileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceProfileParameters `json:"forProvider"`
}

// InstanceProfileStatus defines the observed state of InstanceProfile.
type InstanceProfileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceProfileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceProfile is the Schema for the InstanceProfiles API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type InstanceProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceProfileSpec   `json:"spec"`
	Status            InstanceProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceProfileList contains a list of InstanceProfiles
type InstanceProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstanceProfile `json:"items"`
}

// Repository type metadata.
var (
	InstanceProfile_Kind             = "InstanceProfile"
	InstanceProfile_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InstanceProfile_Kind}.String()
	InstanceProfile_KindAPIVersion   = InstanceProfile_Kind + "." + CRDGroupVersion.String()
	InstanceProfile_GroupVersionKind = CRDGroupVersion.WithKind(InstanceProfile_Kind)
)

func init() {
	SchemeBuilder.Register(&InstanceProfile{}, &InstanceProfileList{})
}
