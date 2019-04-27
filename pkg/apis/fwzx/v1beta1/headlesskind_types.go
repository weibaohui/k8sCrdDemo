/*
Copyright 2019 geovis.

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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// HeadlessKindSpec defines the desired state of HeadlessKind
type HeadlessKindSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Replicas int32  `json:"replicas,omitempty"`
	Hope     string `json:"hope,omitempty"`
}

// HeadlessKindStatus defines the observed state of HeadlessKind
type HeadlessKindStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	HealthyReplicas int32  `json:"healthyReplicas,omitempty"`
	Hope            string `json:"hope,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HeadlessKind is the Schema for the headlesskinds API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type HeadlessKind struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HeadlessKindSpec   `json:"spec,omitempty"`
	Status HeadlessKindStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HeadlessKindList contains a list of HeadlessKind
type HeadlessKindList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HeadlessKind `json:"items"`
}

func init() {
	SchemeBuilder.Register(&HeadlessKind{}, &HeadlessKindList{})
}
