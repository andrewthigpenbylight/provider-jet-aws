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

type HealthCheckObservation struct {
}

type HealthCheckParameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	HealthyThreshold *int64 `json:"healthyThreshold,omitempty" tf:"healthy_threshold,omitempty"`

	// +kubebuilder:validation:Optional
	Interval *int64 `json:"interval,omitempty" tf:"interval,omitempty"`

	// +kubebuilder:validation:Optional
	Matcher *string `json:"matcher,omitempty" tf:"matcher,omitempty"`

	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// +kubebuilder:validation:Optional
	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// +kubebuilder:validation:Optional
	Timeout *int64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// +kubebuilder:validation:Optional
	UnhealthyThreshold *int64 `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold,omitempty"`
}

type LbTargetGroupObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ArnSuffix *string `json:"arnSuffix,omitempty" tf:"arn_suffix,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type LbTargetGroupParameters struct {

	// +kubebuilder:validation:Optional
	DeregistrationDelay *int64 `json:"deregistrationDelay,omitempty" tf:"deregistration_delay,omitempty"`

	// +kubebuilder:validation:Optional
	HealthCheck []HealthCheckParameters `json:"healthCheck,omitempty" tf:"health_check,omitempty"`

	// +kubebuilder:validation:Optional
	LambdaMultiValueHeadersEnabled *bool `json:"lambdaMultiValueHeadersEnabled,omitempty" tf:"lambda_multi_value_headers_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	LoadBalancingAlgorithmType *string `json:"loadBalancingAlgorithmType,omitempty" tf:"load_balancing_algorithm_type,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`

	// +kubebuilder:validation:Optional
	Port *int64 `json:"port,omitempty" tf:"port,omitempty"`

	// +kubebuilder:validation:Optional
	PreserveClientIP *string `json:"preserveClientIp,omitempty" tf:"preserve_client_ip,omitempty"`

	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// +kubebuilder:validation:Optional
	ProtocolVersion *string `json:"protocolVersion,omitempty" tf:"protocol_version,omitempty"`

	// +kubebuilder:validation:Optional
	ProxyProtocolV2 *bool `json:"proxyProtocolV2,omitempty" tf:"proxy_protocol_v2,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-,omitempty"`

	// +kubebuilder:validation:Optional
	SlowStart *int64 `json:"slowStart,omitempty" tf:"slow_start,omitempty"`

	// +kubebuilder:validation:Optional
	Stickiness []LbTargetGroupStickinessParameters `json:"stickiness,omitempty" tf:"stickiness,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	TargetType *string `json:"targetType,omitempty" tf:"target_type,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tf-aws/apis/vpc/v1alpha1.VPC
	// +kubebuilder:validation:Optional
	VpcID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// +kubebuilder:validation:Optional
	VpcIDRef *v1.Reference `json:"vpcIDRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VpcIDSelector *v1.Selector `json:"vpcIDSelector,omitempty" tf:"-"`
}

type LbTargetGroupStickinessObservation struct {
}

type LbTargetGroupStickinessParameters struct {

	// +kubebuilder:validation:Optional
	CookieDuration *int64 `json:"cookieDuration,omitempty" tf:"cookie_duration,omitempty"`

	// +kubebuilder:validation:Optional
	CookieName *string `json:"cookieName,omitempty" tf:"cookie_name,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// LbTargetGroupSpec defines the desired state of LbTargetGroup
type LbTargetGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LbTargetGroupParameters `json:"forProvider"`
}

// LbTargetGroupStatus defines the observed state of LbTargetGroup.
type LbTargetGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LbTargetGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LbTargetGroup is the Schema for the LbTargetGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type LbTargetGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LbTargetGroupSpec   `json:"spec"`
	Status            LbTargetGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LbTargetGroupList contains a list of LbTargetGroups
type LbTargetGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LbTargetGroup `json:"items"`
}

// Repository type metadata.
var (
	LbTargetGroupKind             = "LbTargetGroup"
	LbTargetGroupGroupKind        = schema.GroupKind{Group: Group, Kind: LbTargetGroupKind}.String()
	LbTargetGroupKindAPIVersion   = LbTargetGroupKind + "." + GroupVersion.String()
	LbTargetGroupGroupVersionKind = GroupVersion.WithKind(LbTargetGroupKind)
)

func init() {
	SchemeBuilder.Register(&LbTargetGroup{}, &LbTargetGroupList{})
}
