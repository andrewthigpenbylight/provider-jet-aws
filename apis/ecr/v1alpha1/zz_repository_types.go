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

type EncryptionConfigurationObservation struct {
}

type EncryptionConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	EncryptionType *string `json:"encryptionType,omitempty" tf:"encryption_type,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tf-aws/apis/kms/v1alpha1.Key
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-tf-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	KmsKey *string `json:"kmsKey,omitempty" tf:"kms_key,omitempty"`

	// +kubebuilder:validation:Optional
	KmsKeyRef *v1.Reference `json:"kmsKeyRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	KmsKeySelector *v1.Selector `json:"kmsKeySelector,omitempty" tf:"-"`
}

type ImageScanningConfigurationObservation struct {
}

type ImageScanningConfigurationParameters struct {

	// +kubebuilder:validation:Required
	ScanOnPush *bool `json:"scanOnPush" tf:"scan_on_push,omitempty"`
}

type RepositoryObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	RegistryID *string `json:"registryId,omitempty" tf:"registry_id,omitempty"`

	RepositoryURL *string `json:"repositoryUrl,omitempty" tf:"repository_url,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type RepositoryParameters struct {

	// +kubebuilder:validation:Optional
	EncryptionConfiguration []EncryptionConfigurationParameters `json:"encryptionConfiguration,omitempty" tf:"encryption_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	ImageScanningConfiguration []ImageScanningConfigurationParameters `json:"imageScanningConfiguration,omitempty" tf:"image_scanning_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	ImageTagMutability *string `json:"imageTagMutability,omitempty" tf:"image_tag_mutability,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// RepositorySpec defines the desired state of Repository
type RepositorySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RepositoryParameters `json:"forProvider"`
}

// RepositoryStatus defines the observed state of Repository.
type RepositoryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RepositoryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Repository is the Schema for the Repositorys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Repository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RepositorySpec   `json:"spec"`
	Status            RepositoryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RepositoryList contains a list of Repositorys
type RepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Repository `json:"items"`
}

// Repository type metadata.
var (
	RepositoryKind             = "Repository"
	RepositoryGroupKind        = schema.GroupKind{Group: Group, Kind: RepositoryKind}.String()
	RepositoryKindAPIVersion   = RepositoryKind + "." + GroupVersion.String()
	RepositoryGroupVersionKind = GroupVersion.WithKind(RepositoryKind)
)

func init() {
	SchemeBuilder.Register(&Repository{}, &RepositoryList{})
}
