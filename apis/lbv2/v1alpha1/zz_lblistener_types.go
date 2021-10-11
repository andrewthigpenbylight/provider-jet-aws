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

type AuthenticateCognitoObservation struct {
}

type AuthenticateCognitoParameters struct {

	// +kubebuilder:validation:Optional
	AuthenticationRequestExtraParams map[string]*string `json:"authenticationRequestExtraParams,omitempty" tf:"authentication_request_extra_params,omitempty"`

	// +kubebuilder:validation:Optional
	OnUnauthenticatedRequest *string `json:"onUnauthenticatedRequest,omitempty" tf:"on_unauthenticated_request,omitempty"`

	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// +kubebuilder:validation:Optional
	SessionCookieName *string `json:"sessionCookieName,omitempty" tf:"session_cookie_name,omitempty"`

	// +kubebuilder:validation:Optional
	SessionTimeout *int64 `json:"sessionTimeout,omitempty" tf:"session_timeout,omitempty"`

	// +kubebuilder:validation:Required
	UserPoolArn *string `json:"userPoolArn" tf:"user_pool_arn,omitempty"`

	// +kubebuilder:validation:Required
	UserPoolClientID *string `json:"userPoolClientId" tf:"user_pool_client_id,omitempty"`

	// +kubebuilder:validation:Required
	UserPoolDomain *string `json:"userPoolDomain" tf:"user_pool_domain,omitempty"`
}

type AuthenticateOidcObservation struct {
}

type AuthenticateOidcParameters struct {

	// +kubebuilder:validation:Optional
	AuthenticationRequestExtraParams map[string]*string `json:"authenticationRequestExtraParams,omitempty" tf:"authentication_request_extra_params,omitempty"`

	// +kubebuilder:validation:Required
	AuthorizationEndpoint *string `json:"authorizationEndpoint" tf:"authorization_endpoint,omitempty"`

	// +kubebuilder:validation:Required
	ClientID *string `json:"clientId" tf:"client_id,omitempty"`

	// +kubebuilder:validation:Required
	ClientSecretSecretRef v1.SecretKeySelector `json:"clientSecretSecretRef" tf:"-,omitempty"`

	// +kubebuilder:validation:Required
	Issuer *string `json:"issuer" tf:"issuer,omitempty"`

	// +kubebuilder:validation:Optional
	OnUnauthenticatedRequest *string `json:"onUnauthenticatedRequest,omitempty" tf:"on_unauthenticated_request,omitempty"`

	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// +kubebuilder:validation:Optional
	SessionCookieName *string `json:"sessionCookieName,omitempty" tf:"session_cookie_name,omitempty"`

	// +kubebuilder:validation:Optional
	SessionTimeout *int64 `json:"sessionTimeout,omitempty" tf:"session_timeout,omitempty"`

	// +kubebuilder:validation:Required
	TokenEndpoint *string `json:"tokenEndpoint" tf:"token_endpoint,omitempty"`

	// +kubebuilder:validation:Required
	UserInfoEndpoint *string `json:"userInfoEndpoint" tf:"user_info_endpoint,omitempty"`
}

type DefaultActionObservation struct {
}

type DefaultActionParameters struct {

	// +kubebuilder:validation:Optional
	AuthenticateCognito []AuthenticateCognitoParameters `json:"authenticateCognito,omitempty" tf:"authenticate_cognito,omitempty"`

	// +kubebuilder:validation:Optional
	AuthenticateOidc []AuthenticateOidcParameters `json:"authenticateOidc,omitempty" tf:"authenticate_oidc,omitempty"`

	// +kubebuilder:validation:Optional
	FixedResponse []FixedResponseParameters `json:"fixedResponse,omitempty" tf:"fixed_response,omitempty"`

	// +kubebuilder:validation:Optional
	Forward []ForwardParameters `json:"forward,omitempty" tf:"forward,omitempty"`

	// +kubebuilder:validation:Optional
	Order *int64 `json:"order,omitempty" tf:"order,omitempty"`

	// +kubebuilder:validation:Optional
	Redirect []RedirectParameters `json:"redirect,omitempty" tf:"redirect,omitempty"`

	// +crossplane:generate:reference:type=LbTargetGroup
	// +kubebuilder:validation:Optional
	TargetGroupArn *string `json:"targetGroupArn,omitempty" tf:"target_group_arn,omitempty"`

	// +kubebuilder:validation:Optional
	TargetGroupArnRef *v1.Reference `json:"targetGroupArnRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	TargetGroupArnSelector *v1.Selector `json:"targetGroupArnSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type FixedResponseObservation struct {
}

type FixedResponseParameters struct {

	// +kubebuilder:validation:Required
	ContentType *string `json:"contentType" tf:"content_type,omitempty"`

	// +kubebuilder:validation:Optional
	MessageBody *string `json:"messageBody,omitempty" tf:"message_body,omitempty"`

	// +kubebuilder:validation:Optional
	StatusCode *string `json:"statusCode,omitempty" tf:"status_code,omitempty"`
}

type ForwardObservation struct {
}

type ForwardParameters struct {

	// +kubebuilder:validation:Optional
	Stickiness []StickinessParameters `json:"stickiness,omitempty" tf:"stickiness,omitempty"`

	// +kubebuilder:validation:Required
	TargetGroup []TargetGroupParameters `json:"targetGroup" tf:"target_group,omitempty"`
}

type LbListenerObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type LbListenerParameters struct {

	// +kubebuilder:validation:Optional
	AlpnPolicy *string `json:"alpnPolicy,omitempty" tf:"alpn_policy,omitempty"`

	// +kubebuilder:validation:Optional
	CertificateArn *string `json:"certificateArn,omitempty" tf:"certificate_arn,omitempty"`

	// +kubebuilder:validation:Required
	DefaultAction []DefaultActionParameters `json:"defaultAction" tf:"default_action,omitempty"`

	// +crossplane:generate:reference:type=Lb
	// +kubebuilder:validation:Optional
	LoadBalancerArn *string `json:"loadBalancerArn,omitempty" tf:"load_balancer_arn,omitempty"`

	// +kubebuilder:validation:Optional
	LoadBalancerArnRef *v1.Reference `json:"loadBalancerArnRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	LoadBalancerArnSelector *v1.Selector `json:"loadBalancerArnSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Port *int64 `json:"port,omitempty" tf:"port,omitempty"`

	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-,omitempty"`

	// +kubebuilder:validation:Optional
	SslPolicy *string `json:"sslPolicy,omitempty" tf:"ssl_policy,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RedirectObservation struct {
}

type RedirectParameters struct {

	// +kubebuilder:validation:Optional
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// +kubebuilder:validation:Optional
	Port *string `json:"port,omitempty" tf:"port,omitempty"`

	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// +kubebuilder:validation:Optional
	Query *string `json:"query,omitempty" tf:"query,omitempty"`

	// +kubebuilder:validation:Required
	StatusCode *string `json:"statusCode" tf:"status_code,omitempty"`
}

type StickinessObservation struct {
}

type StickinessParameters struct {

	// +kubebuilder:validation:Required
	Duration *int64 `json:"duration" tf:"duration,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type TargetGroupObservation struct {
}

type TargetGroupParameters struct {

	// +crossplane:generate:reference:type=LbTargetGroup
	// +kubebuilder:validation:Optional
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// +kubebuilder:validation:Optional
	ArnRef *v1.Reference `json:"arnRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ArnSelector *v1.Selector `json:"arnSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Weight *int64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

// LbListenerSpec defines the desired state of LbListener
type LbListenerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LbListenerParameters `json:"forProvider"`
}

// LbListenerStatus defines the observed state of LbListener.
type LbListenerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LbListenerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LbListener is the Schema for the LbListeners API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type LbListener struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LbListenerSpec   `json:"spec"`
	Status            LbListenerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LbListenerList contains a list of LbListeners
type LbListenerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LbListener `json:"items"`
}

// Repository type metadata.
var (
	LbListenerKind             = "LbListener"
	LbListenerGroupKind        = schema.GroupKind{Group: Group, Kind: LbListenerKind}.String()
	LbListenerKindAPIVersion   = LbListenerKind + "." + GroupVersion.String()
	LbListenerGroupVersionKind = GroupVersion.WithKind(LbListenerKind)
)

func init() {
	SchemeBuilder.Register(&LbListener{}, &LbListenerList{})
}
