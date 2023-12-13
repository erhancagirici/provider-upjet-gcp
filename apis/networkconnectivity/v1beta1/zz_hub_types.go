// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type HubInitParameters struct {

	// An optional description of the hub.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Optional labels in key:value format. For more information about labels, see Requirements for labels.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Immutable. The name of the hub. Hub names must be unique. They use the following form: projects/{project_number}/locations/global/hubs/{hub_id}
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The project for the resource
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type HubObservation struct {

	// Output only. The time the hub was created.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// An optional description of the hub.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// an identifier for the resource with format projects/{{project}}/locations/global/hubs/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Optional labels in key:value format. For more information about labels, see Requirements for labels.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Immutable. The name of the hub. Hub names must be unique. They use the following form: projects/{project_number}/locations/global/hubs/{hub_id}
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The project for the resource
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The VPC network associated with this hub's spokes. All of the VPN tunnels, VLAN attachments, and router appliance instances referenced by this hub's spokes must belong to this VPC network. This field is read-only. Network Connectivity Center automatically populates it based on the set of spokes attached to the hub.
	RoutingVpcs []RoutingVpcsObservation `json:"routingVpcs,omitempty" tf:"routing_vpcs,omitempty"`

	// Output only. The current lifecycle state of this hub. Possible values: STATE_UNSPECIFIED, CREATING, ACTIVE, DELETING
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Output only. The Google-generated UUID for the hub. This value is unique across all hub resources. If a hub is deleted and another with the same name is created, the new hub is assigned a different unique_id.
	UniqueID *string `json:"uniqueId,omitempty" tf:"unique_id,omitempty"`

	// Output only. The time the hub was last updated.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type HubParameters struct {

	// An optional description of the hub.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Optional labels in key:value format. For more information about labels, see Requirements for labels.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Immutable. The name of the hub. Hub names must be unique. They use the following form: projects/{project_number}/locations/global/hubs/{hub_id}
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The project for the resource
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type RoutingVpcsInitParameters struct {
}

type RoutingVpcsObservation struct {
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type RoutingVpcsParameters struct {
}

// HubSpec defines the desired state of Hub
type HubSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HubParameters `json:"forProvider"`
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
	InitProvider HubInitParameters `json:"initProvider,omitempty"`
}

// HubStatus defines the observed state of Hub.
type HubStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HubObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Hub is the Schema for the Hubs API. The NetworkConnectivity Hub resource
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Hub struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   HubSpec   `json:"spec"`
	Status HubStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HubList contains a list of Hubs
type HubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Hub `json:"items"`
}

// Repository type metadata.
var (
	Hub_Kind             = "Hub"
	Hub_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Hub_Kind}.String()
	Hub_KindAPIVersion   = Hub_Kind + "." + CRDGroupVersion.String()
	Hub_GroupVersionKind = CRDGroupVersion.WithKind(Hub_Kind)
)

func init() {
	SchemeBuilder.Register(&Hub{}, &HubList{})
}
