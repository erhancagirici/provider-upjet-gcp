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

type CertificateMapEntryInitParameters struct {

	// A set of Certificates defines for the given hostname.
	// There can be defined up to fifteen certificates in each Certificate Map Entry.
	// Each certificate must match pattern projects//locations//certificates/*.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/certificatemanager/v1beta2.Certificate
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	Certificates []*string `json:"certificates,omitempty" tf:"certificates,omitempty"`

	// References to Certificate in certificatemanager to populate certificates.
	// +kubebuilder:validation:Optional
	CertificatesRefs []v1.Reference `json:"certificatesRefs,omitempty" tf:"-"`

	// Selector for a list of Certificate in certificatemanager to populate certificates.
	// +kubebuilder:validation:Optional
	CertificatesSelector *v1.Selector `json:"certificatesSelector,omitempty" tf:"-"`

	// A human-readable description of the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A Hostname (FQDN, e.g. example.com) or a wildcard hostname expression (*.example.com)
	// for a set of hostnames with common suffix. Used as Server Name Indication (SNI) for
	// selecting a proper certificate.
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// Set of labels associated with a Certificate Map Entry.
	// An object containing a list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// A predefined matcher for particular cases, other than SNI selection
	Matcher *string `json:"matcher,omitempty" tf:"matcher,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type CertificateMapEntryObservation struct {

	// A set of Certificates defines for the given hostname.
	// There can be defined up to fifteen certificates in each Certificate Map Entry.
	// Each certificate must match pattern projects//locations//certificates/*.
	Certificates []*string `json:"certificates,omitempty" tf:"certificates,omitempty"`

	// Creation timestamp of a Certificate Map Entry. Timestamp in RFC3339 UTC "Zulu" format,
	// with nanosecond resolution and up to nine fractional digits.
	// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// A human-readable description of the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// for all of the labels present on the resource.
	// +mapType=granular
	EffectiveLabels map[string]*string `json:"effectiveLabels,omitempty" tf:"effective_labels,omitempty"`

	// A Hostname (FQDN, e.g. example.com) or a wildcard hostname expression (*.example.com)
	// for a set of hostnames with common suffix. Used as Server Name Indication (SNI) for
	// selecting a proper certificate.
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/global/certificateMaps/{{map}}/certificateMapEntries/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Set of labels associated with a Certificate Map Entry.
	// An object containing a list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// A map entry that is inputted into the cetrificate map
	Map *string `json:"map,omitempty" tf:"map,omitempty"`

	// A predefined matcher for particular cases, other than SNI selection
	Matcher *string `json:"matcher,omitempty" tf:"matcher,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// A serving state of this Certificate Map Entry.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	// +mapType=granular
	TerraformLabels map[string]*string `json:"terraformLabels,omitempty" tf:"terraform_labels,omitempty"`

	// Update timestamp of a Certificate Map Entry. Timestamp in RFC3339 UTC "Zulu" format,
	// with nanosecond resolution and up to nine fractional digits.
	// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type CertificateMapEntryParameters struct {

	// A set of Certificates defines for the given hostname.
	// There can be defined up to fifteen certificates in each Certificate Map Entry.
	// Each certificate must match pattern projects//locations//certificates/*.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/certificatemanager/v1beta2.Certificate
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Certificates []*string `json:"certificates,omitempty" tf:"certificates,omitempty"`

	// References to Certificate in certificatemanager to populate certificates.
	// +kubebuilder:validation:Optional
	CertificatesRefs []v1.Reference `json:"certificatesRefs,omitempty" tf:"-"`

	// Selector for a list of Certificate in certificatemanager to populate certificates.
	// +kubebuilder:validation:Optional
	CertificatesSelector *v1.Selector `json:"certificatesSelector,omitempty" tf:"-"`

	// A human-readable description of the resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A Hostname (FQDN, e.g. example.com) or a wildcard hostname expression (*.example.com)
	// for a set of hostnames with common suffix. Used as Server Name Indication (SNI) for
	// selecting a proper certificate.
	// +kubebuilder:validation:Optional
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// Set of labels associated with a Certificate Map Entry.
	// An object containing a list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// A map entry that is inputted into the cetrificate map
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/certificatemanager/v1beta1.CertificateMap
	// +kubebuilder:validation:Optional
	Map *string `json:"map,omitempty" tf:"map,omitempty"`

	// Reference to a CertificateMap in certificatemanager to populate map.
	// +kubebuilder:validation:Optional
	MapRef *v1.Reference `json:"mapRef,omitempty" tf:"-"`

	// Selector for a CertificateMap in certificatemanager to populate map.
	// +kubebuilder:validation:Optional
	MapSelector *v1.Selector `json:"mapSelector,omitempty" tf:"-"`

	// A predefined matcher for particular cases, other than SNI selection
	// +kubebuilder:validation:Optional
	Matcher *string `json:"matcher,omitempty" tf:"matcher,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

// CertificateMapEntrySpec defines the desired state of CertificateMapEntry
type CertificateMapEntrySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CertificateMapEntryParameters `json:"forProvider"`
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
	InitProvider CertificateMapEntryInitParameters `json:"initProvider,omitempty"`
}

// CertificateMapEntryStatus defines the observed state of CertificateMapEntry.
type CertificateMapEntryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CertificateMapEntryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// CertificateMapEntry is the Schema for the CertificateMapEntrys API. CertificateMapEntry is a list of certificate configurations, that have been issued for a particular hostname
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type CertificateMapEntry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CertificateMapEntrySpec   `json:"spec"`
	Status            CertificateMapEntryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CertificateMapEntryList contains a list of CertificateMapEntrys
type CertificateMapEntryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CertificateMapEntry `json:"items"`
}

// Repository type metadata.
var (
	CertificateMapEntry_Kind             = "CertificateMapEntry"
	CertificateMapEntry_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CertificateMapEntry_Kind}.String()
	CertificateMapEntry_KindAPIVersion   = CertificateMapEntry_Kind + "." + CRDGroupVersion.String()
	CertificateMapEntry_GroupVersionKind = CRDGroupVersion.WithKind(CertificateMapEntry_Kind)
)

func init() {
	SchemeBuilder.Register(&CertificateMapEntry{}, &CertificateMapEntryList{})
}
