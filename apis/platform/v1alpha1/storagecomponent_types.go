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

var ErrUnableToConvertStorageComponent = errors.New("unable to convert to StorageComponent")

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// StorageComponentSpec defines the desired state of StorageComponent.
type StorageComponentSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +kubebuilder:validation:Optional
	// Specifies a reference to the collection to use for this workload.
	// Requires the name and namespace input to find the collection.
	// If no collection field is set, default to selecting the only
	// workload collection in the cluster, which will result in an error
	// if not exactly one collection is found.
	Collection StorageComponentCollectionSpec `json:"collection"`

	// +kubebuilder:default="nukleros-storage-system"
	// +kubebuilder:validation:Optional
	// (Default: "nukleros-storage-system")
	//
	// Namespace to use for storage support services.
	Namespace string `json:"namespace,omitempty"`

	// +kubebuilder:default={}
	// +kubebuilder:validation:Optional
	SeaweedfsOperator StorageComponentSpecSeaweedfsOperator `json:"seaweedfsOperator,omitempty"`

	// +kubebuilder:default={}
	// +kubebuilder:validation:Optional
	Seaweed StorageComponentSpecSeaweed `json:"seaweed,omitempty"`
}

type StorageComponentCollectionSpec struct {
	// +kubebuilder:validation:Required
	// Required if specifying collection.  The name of the collection
	// within a specific collection.namespace to reference.
	Name string `json:"name"`

	// +kubebuilder:validation:Optional
	// (Default: "") The namespace where the collection exists.  Required only if
	// the collection is namespace scoped and not cluster scoped.
	Namespace string `json:"namespace"`
}

type StorageComponentSpecSeaweedfsOperator struct {
	// +kubebuilder:default=2
	// +kubebuilder:validation:Optional
	// (Default: 2)
	//
	// Number of replicas to use for the seaweedfs-operator deployment.
	Replicas int `json:"replicas,omitempty"`

	// +kubebuilder:default="chrislusf/seaweedfs-operator"
	// +kubebuilder:validation:Optional
	// (Default: "chrislusf/seaweedfs-operator")
	//
	// Image repo and name to use for seaweedfs-operator.
	Image string `json:"image,omitempty"`

	// +kubebuilder:default="1.0.42"
	// +kubebuilder:validation:Optional
	// (Default: "1.0.42")
	//
	// Version of seaweedfs-operator to use.
	Version string `json:"version,omitempty"`
}

type StorageComponentSpecSeaweed struct {
	// +kubebuilder:default="chrislusf/seaweedfs"
	// +kubebuilder:validation:Optional
	// (Default: "chrislusf/seaweedfs")
	//
	// Image repo and name to use for the SeaweedFS master, volume and filer servers.
	Image string `json:"image,omitempty"`

	// +kubebuilder:default="4.51"
	// +kubebuilder:validation:Optional
	// (Default: "4.51")
	//
	// Version of SeaweedFS to use.
	Version string `json:"version,omitempty"`

	// +kubebuilder:default={}
	// +kubebuilder:validation:Optional
	Master StorageComponentSpecSeaweedMaster `json:"master,omitempty"`

	// +kubebuilder:default={}
	// +kubebuilder:validation:Optional
	Volume StorageComponentSpecSeaweedVolume `json:"volume,omitempty"`

	// +kubebuilder:default={}
	// +kubebuilder:validation:Optional
	Filer StorageComponentSpecSeaweedFiler `json:"filer,omitempty"`

	// +kubebuilder:default={}
	// +kubebuilder:validation:Optional
	S3 StorageComponentSpecSeaweedS3 `json:"s3,omitempty"`
}

type StorageComponentSpecSeaweedMaster struct {
	// +kubebuilder:default=1
	// +kubebuilder:validation:Optional
	// (Default: 1)
	//
	// Number of SeaweedFS master servers.  Use an odd number (3 for HA) so the masters
	// keep a quorum.
	Replicas int `json:"replicas,omitempty"`

	// +kubebuilder:default="000"
	// +kubebuilder:validation:Optional
	// (Default: "000")
	//
	// Default replication for new volumes, as SeaweedFS "xyz" placement (data centers,
	// racks, servers).  "000" keeps one copy; "001" keeps a second copy on another
	// volume server.
	DefaultReplication string `json:"defaultReplication,omitempty"`
}

type StorageComponentSpecSeaweedVolume struct {
	// +kubebuilder:default=1
	// +kubebuilder:validation:Optional
	// (Default: 1)
	//
	// Number of SeaweedFS volume servers.
	Replicas int `json:"replicas,omitempty"`

	// +kubebuilder:default="10Gi"
	// +kubebuilder:validation:Optional
	// (Default: "10Gi")
	//
	// Size of the persistent volume for each SeaweedFS volume server.
	Storage string `json:"storage,omitempty"`
}

type StorageComponentSpecSeaweedFiler struct {
	// +kubebuilder:default=1
	// +kubebuilder:validation:Optional
	// (Default: 1)
	//
	// Number of SeaweedFS filer servers.  The default leveldb2 filer store is local to
	// each filer, so keep this at 1 unless the filer config points at a shared store.
	Replicas int `json:"replicas,omitempty"`
}

type StorageComponentSpecSeaweedS3 struct {
	// +kubebuilder:default=true
	// +kubebuilder:validation:Optional
	// (Default: true)
	//
	// Run the S3 API gateway on the filer servers.
	Enabled bool `json:"enabled"`
}

// StorageComponentStatus defines the observed state of StorageComponent.
type StorageComponentStatus struct {
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

// StorageComponent is the Schema for the storagecomponents API.
type StorageComponent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageComponentSpec   `json:"spec,omitempty"`
	Status            StorageComponentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StorageComponentList contains a list of StorageComponent.
type StorageComponentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageComponent `json:"items"`
}

// interface methods

// GetReadyStatus returns the ready status for a component.
func (component *StorageComponent) GetReadyStatus() bool {
	return component.Status.Created
}

// SetReadyStatus sets the ready status for a component.
func (component *StorageComponent) SetReadyStatus(ready bool) {
	component.Status.Created = ready
}

// GetDependencyStatus returns the dependency status for a component.
func (component *StorageComponent) GetDependencyStatus() bool {
	return component.Status.DependenciesSatisfied
}

// SetDependencyStatus sets the dependency status for a component.
func (component *StorageComponent) SetDependencyStatus(dependencyStatus bool) {
	component.Status.DependenciesSatisfied = dependencyStatus
}

// GetPhaseConditions returns the phase conditions for a component.
func (component *StorageComponent) GetPhaseConditions() []*status.PhaseCondition {
	return component.Status.Conditions
}

// SetPhaseCondition sets the phase conditions for a component.
func (component *StorageComponent) SetPhaseCondition(condition *status.PhaseCondition) {
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
func (component *StorageComponent) GetChildResourceConditions() []*status.ChildResource {
	return component.Status.Resources
}

// SetResources sets the phase conditions for a component.
func (component *StorageComponent) SetChildResourceCondition(resource *status.ChildResource) {
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
func (*StorageComponent) GetDependencies() []workload.Workload {
	return []workload.Workload{
		&CertificatesComponent{},
	}
}

// GetComponentGVK returns a GVK object for the component.
func (*StorageComponent) GetWorkloadGVK() schema.GroupVersionKind {
	return GroupVersion.WithKind("StorageComponent")
}

func init() {
	SchemeBuilder.Register(func(s *runtime.Scheme) error {
		s.AddKnownTypes(SchemeGroupVersion, &StorageComponent{}, &StorageComponentList{})
		return nil
	})
}
