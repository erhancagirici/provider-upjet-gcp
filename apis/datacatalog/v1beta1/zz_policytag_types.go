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

type PolicyTagInitParameters struct {

	// Description of this policy tag. It must: contain only unicode characters, tabs,
	// newlines, carriage returns and page breaks; and be at most 2000 bytes long when
	// encoded in UTF-8. If not set, defaults to an empty description.
	// If not set, defaults to an empty description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// User defined name of this policy tag. It must: be unique within the parent
	// taxonomy; contain only unicode letters, numbers, underscores, dashes and spaces;
	// not start or end with spaces; and be at most 200 bytes long when encoded in UTF-8.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Resource name of this policy tag's parent policy tag.
	// If empty, it means this policy tag is a top level policy tag.
	// If not set, defaults to an empty string.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/datacatalog/v1beta1.PolicyTag
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	ParentPolicyTag *string `json:"parentPolicyTag,omitempty" tf:"parent_policy_tag,omitempty"`

	// Reference to a PolicyTag in datacatalog to populate parentPolicyTag.
	// +kubebuilder:validation:Optional
	ParentPolicyTagRef *v1.Reference `json:"parentPolicyTagRef,omitempty" tf:"-"`

	// Selector for a PolicyTag in datacatalog to populate parentPolicyTag.
	// +kubebuilder:validation:Optional
	ParentPolicyTagSelector *v1.Selector `json:"parentPolicyTagSelector,omitempty" tf:"-"`

	// Taxonomy the policy tag is associated with
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/datacatalog/v1beta1.Taxonomy
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	Taxonomy *string `json:"taxonomy,omitempty" tf:"taxonomy,omitempty"`

	// Reference to a Taxonomy in datacatalog to populate taxonomy.
	// +kubebuilder:validation:Optional
	TaxonomyRef *v1.Reference `json:"taxonomyRef,omitempty" tf:"-"`

	// Selector for a Taxonomy in datacatalog to populate taxonomy.
	// +kubebuilder:validation:Optional
	TaxonomySelector *v1.Selector `json:"taxonomySelector,omitempty" tf:"-"`
}

type PolicyTagObservation struct {

	// Resource names of child policy tags of this policy tag.
	ChildPolicyTags []*string `json:"childPolicyTags,omitempty" tf:"child_policy_tags,omitempty"`

	// Description of this policy tag. It must: contain only unicode characters, tabs,
	// newlines, carriage returns and page breaks; and be at most 2000 bytes long when
	// encoded in UTF-8. If not set, defaults to an empty description.
	// If not set, defaults to an empty description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// User defined name of this policy tag. It must: be unique within the parent
	// taxonomy; contain only unicode letters, numbers, underscores, dashes and spaces;
	// not start or end with spaces; and be at most 200 bytes long when encoded in UTF-8.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// an identifier for the resource with format {{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Resource name of this policy tag, whose format is:
	// "projects/{project}/locations/{region}/taxonomies/{taxonomy}/policyTags/{policytag}"
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Resource name of this policy tag's parent policy tag.
	// If empty, it means this policy tag is a top level policy tag.
	// If not set, defaults to an empty string.
	ParentPolicyTag *string `json:"parentPolicyTag,omitempty" tf:"parent_policy_tag,omitempty"`

	// Taxonomy the policy tag is associated with
	Taxonomy *string `json:"taxonomy,omitempty" tf:"taxonomy,omitempty"`
}

type PolicyTagParameters struct {

	// Description of this policy tag. It must: contain only unicode characters, tabs,
	// newlines, carriage returns and page breaks; and be at most 2000 bytes long when
	// encoded in UTF-8. If not set, defaults to an empty description.
	// If not set, defaults to an empty description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// User defined name of this policy tag. It must: be unique within the parent
	// taxonomy; contain only unicode letters, numbers, underscores, dashes and spaces;
	// not start or end with spaces; and be at most 200 bytes long when encoded in UTF-8.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// Resource name of this policy tag's parent policy tag.
	// If empty, it means this policy tag is a top level policy tag.
	// If not set, defaults to an empty string.
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/datacatalog/v1beta1.PolicyTag
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ParentPolicyTag *string `json:"parentPolicyTag,omitempty" tf:"parent_policy_tag,omitempty"`

	// Reference to a PolicyTag in datacatalog to populate parentPolicyTag.
	// +kubebuilder:validation:Optional
	ParentPolicyTagRef *v1.Reference `json:"parentPolicyTagRef,omitempty" tf:"-"`

	// Selector for a PolicyTag in datacatalog to populate parentPolicyTag.
	// +kubebuilder:validation:Optional
	ParentPolicyTagSelector *v1.Selector `json:"parentPolicyTagSelector,omitempty" tf:"-"`

	// Taxonomy the policy tag is associated with
	// +crossplane:generate:reference:type=github.com/upbound/provider-gcp/apis/datacatalog/v1beta1.Taxonomy
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Taxonomy *string `json:"taxonomy,omitempty" tf:"taxonomy,omitempty"`

	// Reference to a Taxonomy in datacatalog to populate taxonomy.
	// +kubebuilder:validation:Optional
	TaxonomyRef *v1.Reference `json:"taxonomyRef,omitempty" tf:"-"`

	// Selector for a Taxonomy in datacatalog to populate taxonomy.
	// +kubebuilder:validation:Optional
	TaxonomySelector *v1.Selector `json:"taxonomySelector,omitempty" tf:"-"`
}

// PolicyTagSpec defines the desired state of PolicyTag
type PolicyTagSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PolicyTagParameters `json:"forProvider"`
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
	InitProvider PolicyTagInitParameters `json:"initProvider,omitempty"`
}

// PolicyTagStatus defines the observed state of PolicyTag.
type PolicyTagStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PolicyTagObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// PolicyTag is the Schema for the PolicyTags API. Denotes one policy tag in a taxonomy.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type PolicyTag struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.displayName) || (has(self.initProvider) && has(self.initProvider.displayName))",message="spec.forProvider.displayName is a required parameter"
	Spec   PolicyTagSpec   `json:"spec"`
	Status PolicyTagStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyTagList contains a list of PolicyTags
type PolicyTagList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PolicyTag `json:"items"`
}

// Repository type metadata.
var (
	PolicyTag_Kind             = "PolicyTag"
	PolicyTag_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PolicyTag_Kind}.String()
	PolicyTag_KindAPIVersion   = PolicyTag_Kind + "." + CRDGroupVersion.String()
	PolicyTag_GroupVersionKind = CRDGroupVersion.WithKind(PolicyTag_Kind)
)

func init() {
	SchemeBuilder.Register(&PolicyTag{}, &PolicyTagList{})
}
