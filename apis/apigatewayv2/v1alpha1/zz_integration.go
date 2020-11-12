/*
Copyright 2020 The Crossplane Authors.

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

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	runtimev1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// IntegrationParameters defines the desired state of Integration
type IntegrationParameters struct {
	// Region is which region the Integration will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region,omitempty"`

	ConnectionID *string `json:"connectionID,omitempty"`

	ConnectionType *string `json:"connectionType,omitempty"`

	ContentHandlingStrategy *string `json:"contentHandlingStrategy,omitempty"`

	CredentialsARN *string `json:"credentialsARN,omitempty"`

	Description *string `json:"description,omitempty"`

	IntegrationMethod *string `json:"integrationMethod,omitempty"`

	IntegrationSubtype *string `json:"integrationSubtype,omitempty"`

	// +kubebuilder:validation:Required
	IntegrationType *string `json:"integrationType"`

	IntegrationURI *string `json:"integrationURI,omitempty"`

	PassthroughBehavior *string `json:"passthroughBehavior,omitempty"`

	PayloadFormatVersion *string `json:"payloadFormatVersion,omitempty"`

	RequestParameters map[string]*string `json:"requestParameters,omitempty"`

	RequestTemplates map[string]*string `json:"requestTemplates,omitempty"`

	TemplateSelectionExpression *string `json:"templateSelectionExpression,omitempty"`

	TimeoutInMillis *int64 `json:"timeoutInMillis,omitempty"`

	TLSConfig                   *TLSConfigInput `json:"tlsConfig,omitempty"`
	CustomIntegrationParameters `json:",inline"`
}

// IntegrationSpec defines the desired state of Integration
type IntegrationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  IntegrationParameters `json:"forProvider"`
}

// IntegrationObservation defines the observed state of Integration
type IntegrationObservation struct {
	APIGatewayManaged *bool `json:"apiGatewayManaged,omitempty"`

	IntegrationID *string `json:"integrationID,omitempty"`

	IntegrationResponseSelectionExpression *string `json:"integrationResponseSelectionExpression,omitempty"`
}

// IntegrationStatus defines the observed state of Integration.
type IntegrationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     IntegrationObservation `json:"atProvider"`
}

// +kubebuilder:object:root=true

// Integration is the Schema for the Integrations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Integration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IntegrationSpec   `json:"spec,omitempty"`
	Status            IntegrationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IntegrationList contains a list of Integrations
type IntegrationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Integration `json:"items"`
}

// Repository type metadata.
var (
	IntegrationKind             = "Integration"
	IntegrationGroupKind        = schema.GroupKind{Group: Group, Kind: IntegrationKind}.String()
	IntegrationKindAPIVersion   = IntegrationKind + "." + GroupVersion.String()
	IntegrationGroupVersionKind = GroupVersion.WithKind(IntegrationKind)
)

func init() {
	SchemeBuilder.Register(&Integration{}, &IntegrationList{})
}
