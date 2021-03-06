/*


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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PullSecretBindingSpec defines the desired state of PullSecretBinding
type PullSecretBindingSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Indicate which harbor server configuration is referred
	HarborServerConfig string `json:"harborServerConfig"`

	// Indicate which service account binds the pull secret
	ServiceAccount string `json:"serviceAccount"`
}

// PullSecretBindingStatus defines the observed state of PullSecretBinding
type PullSecretBindingStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// Indicate the status of binding: `binding`, `bound` and `unknown`
	Status string `json:"status"`

	// Conditions list of extracted conditions from Resource
	// Add the relate conditions of the parts involved in this binding
	// +listType:map
	// +listMapKey:type
	Conditions []Condition `json:"conditions"`
}

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories="goharbor",shortName="psb"
// +kubebuilder:printcolumn:name="Harbor Server",type=string,JSONPath=`.spec.harborServerConfig`,description="The Harbor server configuration CR reference",priority=0
// +kubebuilder:printcolumn:name="Service Account",type=string,JSONPath=`.spec.serviceAccount`,description="The service account binding the pull secret",priority=0
// +kubebuilder:printcolumn:name="Status",type=string,JSONPath=`.status.status`,description="The status of the Harbor server",priority=0

// PullSecretBinding is the Schema for the pullsecretbindings API
type PullSecretBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PullSecretBindingSpec   `json:"spec,omitempty"`
	Status PullSecretBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PullSecretBindingList contains a list of PullSecretBinding
type PullSecretBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PullSecretBinding `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PullSecretBinding{}, &PullSecretBindingList{})
}
