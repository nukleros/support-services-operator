/*
Copyright 2023.

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
	"errors"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"
	"github.com/nukleros/operator-builder-tools/pkg/status"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var ErrUnableToConvertCertManager = errors.New("unable to convert to CertManager")

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CertManagerSpec defines the desired state of CertManager.
type CertManagerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +kubebuilder:validation:Optional
	// Specifies a reference to the collection to use for this workload.
	// Requires the name and namespace input to find the collection.
	// If no collection field is set, default to selecting the only
	// workload collection in the cluster, which will result in an error
	// if not exactly one collection is found.
	Collection CertManagerCollectionSpec `json:"collection"`

	// +kubebuilder:default="nukleros-certs-system"
	// +kubebuilder:validation:Optional
	// (Default: "nukleros-certs-system")
	//  Namespace to use for certificate support services.
	Namespace string `json:"namespace,omitempty"`

	// +kubebuilder:validation:Optional
	Cainjector CertManagerSpecCainjector `json:"cainjector,omitempty"`

	// +kubebuilder:default="v1.9.1"
	// +kubebuilder:validation:Optional
	// (Default: "v1.9.1")
	//  Version of cert-manager to use.
	Version string `json:"version,omitempty"`

	// +kubebuilder:validation:Optional
	Controller CertManagerSpecController `json:"controller,omitempty"`

	// +kubebuilder:validation:Optional
	Webhook CertManagerSpecWebhook `json:"webhook,omitempty"`

	// +kubebuilder:validation:Required
	//  Contact e-mail address for receiving updates about certificates from LetsEncrypt.
	ContactEmail string `json:"contactEmail,omitempty"`
}

type CertManagerCollectionSpec struct {
	// +kubebuilder:validation:Required
	// Required if specifying collection.  The name of the collection
	// within a specific collection.namespace to reference.
	Name string `json:"name"`

	// +kubebuilder:validation:Optional
	// (Default: "") The namespace where the collection exists.  Required only if
	// the collection is namespace scoped and not cluster scoped.
	Namespace string `json:"namespace"`
}

type CertManagerSpecCainjector struct {
	// +kubebuilder:default=2
	// +kubebuilder:validation:Optional
	// (Default: 2)
	//  Number of replicas to use for the cert-manager cainjector deployment.
	Replicas int `json:"replicas,omitempty"`

	// +kubebuilder:default="quay.io/jetstack/cert-manager-cainjector"
	// +kubebuilder:validation:Optional
	// (Default: "quay.io/jetstack/cert-manager-cainjector")
	//  Image repo and name to use for cert-manager cainjector.
	Image string `json:"image,omitempty"`
}

type CertManagerSpecController struct {
	// +kubebuilder:default=2
	// +kubebuilder:validation:Optional
	// (Default: 2)
	//  Number of replicas to use for the cert-manager controller deployment.
	Replicas int `json:"replicas,omitempty"`

	// +kubebuilder:default="quay.io/jetstack/cert-manager-controller"
	// +kubebuilder:validation:Optional
	// (Default: "quay.io/jetstack/cert-manager-controller")
	//  Image repo and name to use for cert-manager controller.
	Image string `json:"image,omitempty"`
}

type CertManagerSpecWebhook struct {
	// +kubebuilder:default=2
	// +kubebuilder:validation:Optional
	// (Default: 2)
	//  Number of replicas to use for the cert-manager webhook deployment.
	Replicas int `json:"replicas,omitempty"`

	// +kubebuilder:default="quay.io/jetstack/cert-manager-webhook"
	// +kubebuilder:validation:Optional
	// (Default: "quay.io/jetstack/cert-manager-webhook")
	//  Image repo and name to use for cert-manager webhook.
	Image string `json:"image,omitempty"`
}

// CertManagerStatus defines the observed state of CertManager.
type CertManagerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Created               bool                     `json:"created,omitempty"`
	DependenciesSatisfied bool                     `json:"dependenciesSatisfied,omitempty"`
	Conditions            []*status.PhaseCondition `json:"conditions,omitempty"`
	Resources             []*status.ChildResource  `json:"resources,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster

// CertManager is the Schema for the certmanagers API.
type CertManager struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CertManagerSpec   `json:"spec,omitempty"`
	Status            CertManagerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CertManagerList contains a list of CertManager.
type CertManagerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CertManager `json:"items"`
}

// interface methods

// GetReadyStatus returns the ready status for a component.
func (component *CertManager) GetReadyStatus() bool {
	return component.Status.Created
}

// SetReadyStatus sets the ready status for a component.
func (component *CertManager) SetReadyStatus(ready bool) {
	component.Status.Created = ready
}

// GetDependencyStatus returns the dependency status for a component.
func (component *CertManager) GetDependencyStatus() bool {
	return component.Status.DependenciesSatisfied
}

// SetDependencyStatus sets the dependency status for a component.
func (component *CertManager) SetDependencyStatus(dependencyStatus bool) {
	component.Status.DependenciesSatisfied = dependencyStatus
}

// GetPhaseConditions returns the phase conditions for a component.
func (component *CertManager) GetPhaseConditions() []*status.PhaseCondition {
	return component.Status.Conditions
}

// SetPhaseCondition sets the phase conditions for a component.
func (component *CertManager) SetPhaseCondition(condition *status.PhaseCondition) {
	for i, currentCondition := range component.GetPhaseConditions() {
		if currentCondition.Phase == condition.Phase {
			component.Status.Conditions[i] = condition

			return
		}
	}

	// phase not found, lets add it to the list.
	component.Status.Conditions = append(component.Status.Conditions, condition)
}

// GetResources returns the child resource status for a component.
func (component *CertManager) GetChildResourceConditions() []*status.ChildResource {
	return component.Status.Resources
}

// SetResources sets the phase conditions for a component.
func (component *CertManager) SetChildResourceCondition(resource *status.ChildResource) {
	for i, currentResource := range component.GetChildResourceConditions() {
		if currentResource.Group == resource.Group && currentResource.Version == resource.Version && currentResource.Kind == resource.Kind {
			if currentResource.Name == resource.Name && currentResource.Namespace == resource.Namespace {
				component.Status.Resources[i] = resource

				return
			}
		}
	}

	// phase not found, lets add it to the collection
	component.Status.Resources = append(component.Status.Resources, resource)
}

// GetDependencies returns the dependencies for a component.
func (*CertManager) GetDependencies() []workload.Workload {
	return []workload.Workload{}
}

// GetComponentGVK returns a GVK object for the component.
func (*CertManager) GetWorkloadGVK() schema.GroupVersionKind {
	return GroupVersion.WithKind("CertManager")
}

func init() {
	SchemeBuilder.Register(&CertManager{}, &CertManagerList{})
}
