/*
Copyright 2022.

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

var ErrUnableToConvertIngressComponent = errors.New("unable to convert to IngressComponent")

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// IngressComponentSpec defines the desired state of IngressComponent.
type IngressComponentSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +kubebuilder:validation:Optional
	// Specifies a reference to the collection to use for this workload.
	// Requires the name and namespace input to find the collection.
	// If no collection field is set, default to selecting the only
	// workload collection in the cluster, which will result in an error
	// if not exactly one collection is found.
	Collection IngressComponentCollectionSpec `json:"collection"`

	// +kubebuilder:default="nukleros-ingress-system"
	// +kubebuilder:validation:Optional
	// (Default: "nukleros-ingress-system")
	//  Namespace to use for ingress support services.
	Namespace string `json:"namespace,omitempty"`

	// +kubebuilder:validation:Optional
	ExternalDNS IngressComponentSpecExternalDNS `json:"externalDNS,omitempty"`

	ExternalDNSProvider string `json:"externalDNSProvider,omitempty"`
}

type IngressComponentCollectionSpec struct {
	// +kubebuilder:validation:Required
	// Required if specifying collection.  The name of the collection
	// within a specific collection.namespace to reference.
	Name string `json:"name"`

	// +kubebuilder:validation:Optional
	// (Default: "") The namespace where the collection exists.  Required only if
	// the collection is namespace scoped and not cluster scoped.
	Namespace string `json:"namespace"`
}

type IngressComponentSpecExternalDNS struct {
	// +kubebuilder:default="k8s.gcr.io/external-dns/external-dns"
	// +kubebuilder:validation:Optional
	// (Default: "k8s.gcr.io/external-dns/external-dns")
	//  Image repo and name to use for external-dns.
	//  Image repo and name to use for external-dns.
	Image string `json:"image,omitempty"`

	// +kubebuilder:default="v0.12.2"
	// +kubebuilder:validation:Optional
	// (Default: "v0.12.2")
	//  Version of external-dns to use.
	//  Version of external-dns to use.
	Version string `json:"version,omitempty"`
}

// IngressComponentStatus defines the observed state of IngressComponent.
type IngressComponentStatus struct {
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

// IngressComponent is the Schema for the ingresscomponents API.
type IngressComponent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IngressComponentSpec   `json:"spec,omitempty"`
	Status            IngressComponentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IngressComponentList contains a list of IngressComponent.
type IngressComponentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IngressComponent `json:"items"`
}

// interface methods

// GetReadyStatus returns the ready status for a component.
func (component *IngressComponent) GetReadyStatus() bool {
	return component.Status.Created
}

// SetReadyStatus sets the ready status for a component.
func (component *IngressComponent) SetReadyStatus(ready bool) {
	component.Status.Created = ready
}

// GetDependencyStatus returns the dependency status for a component.
func (component *IngressComponent) GetDependencyStatus() bool {
	return component.Status.DependenciesSatisfied
}

// SetDependencyStatus sets the dependency status for a component.
func (component *IngressComponent) SetDependencyStatus(dependencyStatus bool) {
	component.Status.DependenciesSatisfied = dependencyStatus
}

// GetPhaseConditions returns the phase conditions for a component.
func (component *IngressComponent) GetPhaseConditions() []*status.PhaseCondition {
	return component.Status.Conditions
}

// SetPhaseCondition sets the phase conditions for a component.
func (component *IngressComponent) SetPhaseCondition(condition *status.PhaseCondition) {
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
func (component *IngressComponent) GetChildResourceConditions() []*status.ChildResource {
	return component.Status.Resources
}

// SetResources sets the phase conditions for a component.
func (component *IngressComponent) SetChildResourceCondition(resource *status.ChildResource) {
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
func (*IngressComponent) GetDependencies() []workload.Workload {
	return []workload.Workload{}
}

// GetComponentGVK returns a GVK object for the component.
func (*IngressComponent) GetWorkloadGVK() schema.GroupVersionKind {
	return GroupVersion.WithKind("IngressComponent")
}

func init() {
	SchemeBuilder.Register(&IngressComponent{}, &IngressComponentList{})
}
