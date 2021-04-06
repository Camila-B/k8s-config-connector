// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Config Connector and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ServiceNetworkingConnectionSpec struct {
	/*  */
	NetworkRef v1alpha1.ResourceRef `json:"networkRef,omitempty"`
	/*  */
	ReservedPeeringRanges []v1alpha1.ResourceRef `json:"reservedPeeringRanges,omitempty"`
	/* Immutable. Provider peering service that is managing peering connectivity for a service provider organization. For Google services that support this functionality it is 'servicenetworking.googleapis.com'. */
	Service string `json:"service,omitempty"`
}

type ServiceNetworkingConnectionStatus struct {
	/* Conditions represent the latest available observations of the
	   ServiceNetworkingConnection's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/*  */
	Peering string `json:"peering,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceNetworkingConnection is the Schema for the servicenetworking API
// +k8s:openapi-gen=true
type ServiceNetworkingConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceNetworkingConnectionSpec   `json:"spec,omitempty"`
	Status ServiceNetworkingConnectionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServiceNetworkingConnectionList contains a list of ServiceNetworkingConnection
type ServiceNetworkingConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceNetworkingConnection `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ServiceNetworkingConnection{}, &ServiceNetworkingConnectionList{})
}