// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type InstanceAttachmentInitParameters struct {

	// The resource ID of the environment.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/apigee/v1beta2.Environment
	Environment *string `json:"environment,omitempty" tf:"environment,omitempty"`

	// Reference to a Environment in apigee to populate environment.
	// +kubebuilder:validation:Optional
	EnvironmentRef *v1.Reference `json:"environmentRef,omitempty" tf:"-"`

	// Selector for a Environment in apigee to populate environment.
	// +kubebuilder:validation:Optional
	EnvironmentSelector *v1.Selector `json:"environmentSelector,omitempty" tf:"-"`

	// The Apigee instance associated with the Apigee environment,
	// in the format organizations/{{org_name}}/instances/{{instance_name}}.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/apigee/v1beta1.Instance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance in apigee to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance in apigee to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`
}

type InstanceAttachmentObservation struct {

	// The resource ID of the environment.
	Environment *string `json:"environment,omitempty" tf:"environment,omitempty"`

	// an identifier for the resource with format {{instance_id}}/attachments/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Apigee instance associated with the Apigee environment,
	// in the format organizations/{{org_name}}/instances/{{instance_name}}.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// The name of the newly created  attachment (output parameter).
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type InstanceAttachmentParameters struct {

	// The resource ID of the environment.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/apigee/v1beta2.Environment
	// +kubebuilder:validation:Optional
	Environment *string `json:"environment,omitempty" tf:"environment,omitempty"`

	// Reference to a Environment in apigee to populate environment.
	// +kubebuilder:validation:Optional
	EnvironmentRef *v1.Reference `json:"environmentRef,omitempty" tf:"-"`

	// Selector for a Environment in apigee to populate environment.
	// +kubebuilder:validation:Optional
	EnvironmentSelector *v1.Selector `json:"environmentSelector,omitempty" tf:"-"`

	// The Apigee instance associated with the Apigee environment,
	// in the format organizations/{{org_name}}/instances/{{instance_name}}.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/apigee/v1beta1.Instance
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance in apigee to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance in apigee to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`
}

// InstanceAttachmentSpec defines the desired state of InstanceAttachment
type InstanceAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceAttachmentParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider InstanceAttachmentInitParameters `json:"initProvider,omitempty"`
}

// InstanceAttachmentStatus defines the observed state of InstanceAttachment.
type InstanceAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// InstanceAttachment is the Schema for the InstanceAttachments API. An
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type InstanceAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceAttachmentSpec   `json:"spec"`
	Status            InstanceAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceAttachmentList contains a list of InstanceAttachments
type InstanceAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstanceAttachment `json:"items"`
}

// Repository type metadata.
var (
	InstanceAttachment_Kind             = "InstanceAttachment"
	InstanceAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InstanceAttachment_Kind}.String()
	InstanceAttachment_KindAPIVersion   = InstanceAttachment_Kind + "." + CRDGroupVersion.String()
	InstanceAttachment_GroupVersionKind = CRDGroupVersion.WithKind(InstanceAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&InstanceAttachment{}, &InstanceAttachmentList{})
}
