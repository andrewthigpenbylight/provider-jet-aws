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

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type TransitGatewayVPCAttachmentObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	VPCOwnerID *string `json:"vpcOwnerId,omitempty" tf:"vpc_owner_id,omitempty"`
}

type TransitGatewayVPCAttachmentParameters struct {

	// +kubebuilder:validation:Optional
	ApplianceModeSupport *string `json:"applianceModeSupport,omitempty" tf:"appliance_mode_support,omitempty"`

	// +kubebuilder:validation:Optional
	DNSSupport *string `json:"dnsSupport,omitempty" tf:"dns_support,omitempty"`

	// +kubebuilder:validation:Optional
	IPv6Support *string `json:"ipv6Support,omitempty" tf:"ipv6_support,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIdRefs []v1.Reference `json:"subnetIdRefs,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIdSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=Subnet
	// +crossplane:generate:reference:refFieldName=SubnetIdRefs
	// +crossplane:generate:reference:selectorFieldName=SubnetIdSelector
	// +kubebuilder:validation:Optional
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	TransitGatewayDefaultRouteTableAssociation *bool `json:"transitGatewayDefaultRouteTableAssociation,omitempty" tf:"transit_gateway_default_route_table_association,omitempty"`

	// +kubebuilder:validation:Optional
	TransitGatewayDefaultRouteTablePropagation *bool `json:"transitGatewayDefaultRouteTablePropagation,omitempty" tf:"transit_gateway_default_route_table_propagation,omitempty"`

	// +crossplane:generate:reference:type=TransitGateway
	// +kubebuilder:validation:Optional
	TransitGatewayID *string `json:"transitGatewayId,omitempty" tf:"transit_gateway_id,omitempty"`

	// +kubebuilder:validation:Optional
	TransitGatewayIDRef *v1.Reference `json:"transitGatewayIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	TransitGatewayIDSelector *v1.Selector `json:"transitGatewayIdSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=VPC
	// +crossplane:generate:reference:refFieldName=VpcIdRef
	// +crossplane:generate:reference:selectorFieldName=VpcIdSelector
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// +kubebuilder:validation:Optional
	VpcIdRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VpcIdSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

// TransitGatewayVPCAttachmentSpec defines the desired state of TransitGatewayVPCAttachment
type TransitGatewayVPCAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TransitGatewayVPCAttachmentParameters `json:"forProvider"`
}

// TransitGatewayVPCAttachmentStatus defines the observed state of TransitGatewayVPCAttachment.
type TransitGatewayVPCAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TransitGatewayVPCAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayVPCAttachment is the Schema for the TransitGatewayVPCAttachments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type TransitGatewayVPCAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TransitGatewayVPCAttachmentSpec   `json:"spec"`
	Status            TransitGatewayVPCAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayVPCAttachmentList contains a list of TransitGatewayVPCAttachments
type TransitGatewayVPCAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TransitGatewayVPCAttachment `json:"items"`
}

// Repository type metadata.
var (
	TransitGatewayVPCAttachment_Kind             = "TransitGatewayVPCAttachment"
	TransitGatewayVPCAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TransitGatewayVPCAttachment_Kind}.String()
	TransitGatewayVPCAttachment_KindAPIVersion   = TransitGatewayVPCAttachment_Kind + "." + CRDGroupVersion.String()
	TransitGatewayVPCAttachment_GroupVersionKind = CRDGroupVersion.WithKind(TransitGatewayVPCAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&TransitGatewayVPCAttachment{}, &TransitGatewayVPCAttachmentList{})
}
