/*
Copyright 2026.

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
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var ErrUnableToConvertServiceMeshComponent = errors.New("unable to convert to ServiceMeshComponent")

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ServiceMeshComponentSpec defines the desired state of ServiceMeshComponent.
type ServiceMeshComponentSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +kubebuilder:validation:Optional
	// Specifies a reference to the collection to use for this workload.
	// Requires the name and namespace input to find the collection.
	// If no collection field is set, default to selecting the only
	// workload collection in the cluster, which will result in an error
	// if not exactly one collection is found.
	Collection ServiceMeshComponentCollectionSpec `json:"collection"`

	// +kubebuilder:default="nukleros-service-mesh-system"
	// +kubebuilder:validation:Optional
	// (Default: "nukleros-service-mesh-system")
	//
	// Namespace to use for service mesh component support services.
	Namespace string `json:"namespace,omitempty"`

	// +kubebuilder:default={}
	// +kubebuilder:validation:Optional
	Istio ServiceMeshComponentSpecIstio `json:"istio,omitempty"`
}

type ServiceMeshComponentCollectionSpec struct {
	// +kubebuilder:validation:Required
	// Required if specifying collection.  The name of the collection
	// within a specific collection.namespace to reference.
	Name string `json:"name"`

	// +kubebuilder:validation:Optional
	// (Default: "") The namespace where the collection exists.  Required only if
	// the collection is namespace scoped and not cluster scoped.
	Namespace string `json:"namespace"`
}

type ServiceMeshComponentSpecIstio struct {
	// +kubebuilder:default="1.30.4"
	// +kubebuilder:validation:Optional
	// (Default: "1.30.4")
	//
	// Version of istio to use.  This version is used for all istio images, including
	// the proxy image used for waypoints.  Distroless images are used, so this should
	// not include the '-distroless' suffix.
	Version string `json:"version,omitempty"`

	// +kubebuilder:default={}
	// +kubebuilder:validation:Optional
	Images ServiceMeshComponentSpecIstioImages `json:"images,omitempty"`

	// +kubebuilder:default=5
	// +kubebuilder:validation:Optional
	// (Default: 5)
	//
	// Maximum number of replicas to use for the istiod deployment.  The istiod
	// deployment is autoscaled between istio.replicas and istio.maxReplicas.
	MaxReplicas int `json:"maxReplicas,omitempty"`

	// +kubebuilder:default=2
	// +kubebuilder:validation:Optional
	// (Default: 2)
	//
	// Minimum number of replicas to use for the istiod deployment.  The istiod
	// deployment is autoscaled between istio.replicas and istio.maxReplicas.
	Replicas int `json:"replicas,omitempty"`
}

type ServiceMeshComponentSpecIstioImages struct {
	// +kubebuilder:default="registry.istio.io/release/pilot"
	// +kubebuilder:validation:Optional
	// (Default: "registry.istio.io/release/pilot")
	//
	// Image repo and name to use for istiod.
	Istiod string `json:"istiod,omitempty"`

	// +kubebuilder:default="registry.istio.io/release/install-cni"
	// +kubebuilder:validation:Optional
	// (Default: "registry.istio.io/release/install-cni")
	//
	// Image repo and name to use for the istio cni node agent.
	Cni string `json:"cni,omitempty"`

	// +kubebuilder:default="registry.istio.io/release/ztunnel"
	// +kubebuilder:validation:Optional
	// (Default: "registry.istio.io/release/ztunnel")
	//
	// Image repo and name to use for ztunnel.
	Ztunnel string `json:"ztunnel,omitempty"`
}

// ServiceMeshComponentStatus defines the observed state of ServiceMeshComponent.
type ServiceMeshComponentStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +kubebuilder:validation:Optional
	Created bool `json:"created"`

	// +kubebuilder:validation:Optional
	DependenciesSatisfied bool                     `json:"dependenciesSatisfied"`
	Conditions            []*status.PhaseCondition `json:"conditions,omitempty"`
	Resources             []*status.ChildResource  `json:"resources,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster

// ServiceMeshComponent is the Schema for the servicemeshcomponents API.
type ServiceMeshComponent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceMeshComponentSpec   `json:"spec,omitempty"`
	Status            ServiceMeshComponentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceMeshComponentList contains a list of ServiceMeshComponent.
type ServiceMeshComponentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceMeshComponent `json:"items"`
}

// interface methods

// GetReadyStatus returns the ready status for a component.
func (component *ServiceMeshComponent) GetReadyStatus() bool {
	return component.Status.Created
}

// SetReadyStatus sets the ready status for a component.
func (component *ServiceMeshComponent) SetReadyStatus(ready bool) {
	component.Status.Created = ready
}

// GetDependencyStatus returns the dependency status for a component.
func (component *ServiceMeshComponent) GetDependencyStatus() bool {
	return component.Status.DependenciesSatisfied
}

// SetDependencyStatus sets the dependency status for a component.
func (component *ServiceMeshComponent) SetDependencyStatus(dependencyStatus bool) {
	component.Status.DependenciesSatisfied = dependencyStatus
}

// GetPhaseConditions returns the phase conditions for a component.
func (component *ServiceMeshComponent) GetPhaseConditions() []*status.PhaseCondition {
	return component.Status.Conditions
}

// SetPhaseCondition sets the phase conditions for a component.
func (component *ServiceMeshComponent) SetPhaseCondition(condition *status.PhaseCondition) {
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
func (component *ServiceMeshComponent) GetChildResourceConditions() []*status.ChildResource {
	return component.Status.Resources
}

// SetResources sets the phase conditions for a component.
func (component *ServiceMeshComponent) SetChildResourceCondition(resource *status.ChildResource) {
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
func (*ServiceMeshComponent) GetDependencies() []workload.Workload {
	return []workload.Workload{}
}

// GetComponentGVK returns a GVK object for the component.
func (*ServiceMeshComponent) GetWorkloadGVK() schema.GroupVersionKind {
	return GroupVersion.WithKind("ServiceMeshComponent")
}

func init() {
	SchemeBuilder.Register(func(s *runtime.Scheme) error {
		s.AddKnownTypes(SchemeGroupVersion, &ServiceMeshComponent{}, &ServiceMeshComponentList{})
		return nil
	})
}
