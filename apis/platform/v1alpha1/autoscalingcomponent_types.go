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

var ErrUnableToConvertAutoscalingComponent = errors.New("unable to convert to AutoscalingComponent")

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AutoscalingComponentSpec defines the desired state of AutoscalingComponent.
type AutoscalingComponentSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +kubebuilder:validation:Optional
	// Specifies a reference to the collection to use for this workload.
	// Requires the name and namespace input to find the collection.
	// If no collection field is set, default to selecting the only
	// workload collection in the cluster, which will result in an error
	// if not exactly one collection is found.
	Collection AutoscalingComponentCollectionSpec `json:"collection"`

	// +kubebuilder:default="nukleros-autoscaling-system"
	// +kubebuilder:validation:Optional
	// (Default: "nukleros-autoscaling-system")
	//
	// Namespace to use for autoscaling support services.
	Namespace string `json:"namespace,omitempty"`

	// +kubebuilder:default={}
	// +kubebuilder:validation:Optional
	Keda AutoscalingComponentSpecKeda `json:"keda,omitempty"`
}

type AutoscalingComponentCollectionSpec struct {
	// +kubebuilder:validation:Required
	// Required if specifying collection.  The name of the collection
	// within a specific collection.namespace to reference.
	Name string `json:"name"`

	// +kubebuilder:validation:Optional
	// (Default: "") The namespace where the collection exists.  Required only if
	// the collection is namespace scoped and not cluster scoped.
	Namespace string `json:"namespace"`
}

type AutoscalingComponentSpecKeda struct {
	// +kubebuilder:default="2.21.0"
	// +kubebuilder:validation:Optional
	// (Default: "2.21.0")
	//
	// Version of KEDA to use.  Applies to the operator, metrics apiserver and
	// admission webhooks, which are released together.
	Version string `json:"version,omitempty"`

	// +kubebuilder:default={}
	// +kubebuilder:validation:Optional
	Operator AutoscalingComponentSpecKedaOperator `json:"operator,omitempty"`

	// +kubebuilder:default={}
	// +kubebuilder:validation:Optional
	MetricsApiserver AutoscalingComponentSpecKedaMetricsApiserver `json:"metricsApiserver,omitempty"`

	// +kubebuilder:default={}
	// +kubebuilder:validation:Optional
	AdmissionWebhooks AutoscalingComponentSpecKedaAdmissionWebhooks `json:"admissionWebhooks,omitempty"`
}

type AutoscalingComponentSpecKedaOperator struct {
	// +kubebuilder:default=2
	// +kubebuilder:validation:Optional
	// (Default: 2)
	//
	// Number of replicas to use for the KEDA operator deployment.  With 1 replica, its
	// pod disruption budget (minAvailable: 1) blocks node drains.
	Replicas int `json:"replicas,omitempty"`

	// +kubebuilder:default="ghcr.io/kedacore/keda"
	// +kubebuilder:validation:Optional
	// (Default: "ghcr.io/kedacore/keda")
	//
	// Image repo and name to use for the KEDA operator.
	Image string `json:"image,omitempty"`
}

type AutoscalingComponentSpecKedaMetricsApiserver struct {
	// +kubebuilder:default=2
	// +kubebuilder:validation:Optional
	// (Default: 2)
	//
	// Number of replicas to use for the KEDA metrics apiserver deployment.  With 1
	// replica, its pod disruption budget (minAvailable: 1) blocks node drains.
	Replicas int `json:"replicas,omitempty"`

	// +kubebuilder:default="ghcr.io/kedacore/keda-metrics-apiserver"
	// +kubebuilder:validation:Optional
	// (Default: "ghcr.io/kedacore/keda-metrics-apiserver")
	//
	// Image repo and name to use for the KEDA metrics apiserver.
	Image string `json:"image,omitempty"`
}

type AutoscalingComponentSpecKedaAdmissionWebhooks struct {
	// +kubebuilder:default=2
	// +kubebuilder:validation:Optional
	// (Default: 2)
	//
	// Number of replicas to use for the KEDA admission webhooks deployment.  With 1
	// replica, its pod disruption budget (minAvailable: 1) blocks node drains.
	Replicas int `json:"replicas,omitempty"`

	// +kubebuilder:default="ghcr.io/kedacore/keda-admission-webhooks"
	// +kubebuilder:validation:Optional
	// (Default: "ghcr.io/kedacore/keda-admission-webhooks")
	//
	// Image repo and name to use for the KEDA admission webhooks.
	Image string `json:"image,omitempty"`
}

// AutoscalingComponentStatus defines the observed state of AutoscalingComponent.
type AutoscalingComponentStatus struct {
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

// AutoscalingComponent is the Schema for the autoscalingcomponents API.
type AutoscalingComponent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AutoscalingComponentSpec   `json:"spec,omitempty"`
	Status            AutoscalingComponentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AutoscalingComponentList contains a list of AutoscalingComponent.
type AutoscalingComponentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AutoscalingComponent `json:"items"`
}

// interface methods

// GetReadyStatus returns the ready status for a component.
func (component *AutoscalingComponent) GetReadyStatus() bool {
	return component.Status.Created
}

// SetReadyStatus sets the ready status for a component.
func (component *AutoscalingComponent) SetReadyStatus(ready bool) {
	component.Status.Created = ready
}

// GetDependencyStatus returns the dependency status for a component.
func (component *AutoscalingComponent) GetDependencyStatus() bool {
	return component.Status.DependenciesSatisfied
}

// SetDependencyStatus sets the dependency status for a component.
func (component *AutoscalingComponent) SetDependencyStatus(dependencyStatus bool) {
	component.Status.DependenciesSatisfied = dependencyStatus
}

// GetPhaseConditions returns the phase conditions for a component.
func (component *AutoscalingComponent) GetPhaseConditions() []*status.PhaseCondition {
	return component.Status.Conditions
}

// SetPhaseCondition sets the phase conditions for a component.
func (component *AutoscalingComponent) SetPhaseCondition(condition *status.PhaseCondition) {
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
func (component *AutoscalingComponent) GetChildResourceConditions() []*status.ChildResource {
	return component.Status.Resources
}

// SetResources sets the phase conditions for a component.
func (component *AutoscalingComponent) SetChildResourceCondition(resource *status.ChildResource) {
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
func (*AutoscalingComponent) GetDependencies() []workload.Workload {
	return []workload.Workload{}
}

// GetComponentGVK returns a GVK object for the component.
func (*AutoscalingComponent) GetWorkloadGVK() schema.GroupVersionKind {
	return GroupVersion.WithKind("AutoscalingComponent")
}

func init() {
	SchemeBuilder.Register(func(s *runtime.Scheme) error {
		s.AddKnownTypes(SchemeGroupVersion, &AutoscalingComponent{}, &AutoscalingComponentList{})
		return nil
	})
}
