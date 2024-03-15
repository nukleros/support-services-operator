/*
Copyright 2024.

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

var ErrUnableToConvertExternalSecrets = errors.New("unable to convert to ExternalSecrets")

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ExternalSecretsSpec defines the desired state of ExternalSecrets.
type ExternalSecretsSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +kubebuilder:validation:Optional
	// Specifies a reference to the collection to use for this workload.
	// Requires the name and namespace input to find the collection.
	// If no collection field is set, default to selecting the only
	// workload collection in the cluster, which will result in an error
	// if not exactly one collection is found.
	Collection ExternalSecretsCollectionSpec `json:"collection"`

	// +kubebuilder:default="nukleros-secrets-system"
	// +kubebuilder:validation:Optional
	// (Default: "nukleros-secrets-system")
	//  Namespace to use for secrets support services.
	Namespace string `json:"namespace,omitempty"`

	// +kubebuilder:default="v0.9.11"
	// +kubebuilder:validation:Optional
	// (Default: "v0.9.11")
	//  Version of external-secrets to use.
	Version string `json:"version,omitempty"`

	// +kubebuilder:validation:Optional
	CertController ExternalSecretsSpecCertController `json:"certController,omitempty"`

	// +kubebuilder:default="ghcr.io/external-secrets/external-secrets"
	// +kubebuilder:validation:Optional
	// (Default: "ghcr.io/external-secrets/external-secrets")
	//  Image repo and name to use for external-secrets.
	Image string `json:"image,omitempty"`

	// +kubebuilder:validation:Optional
	Controller ExternalSecretsSpecController `json:"controller,omitempty"`

	// +kubebuilder:validation:Optional
	Webhook ExternalSecretsSpecWebhook `json:"webhook,omitempty"`

	// +kubebuilder:validation:Required
	//  On AWS, the IAM Role ARN that gives external-secrets access to SecretsManager
	IamRoleArn string `json:"iamRoleArn,omitempty"`
}

type ExternalSecretsCollectionSpec struct {
	// +kubebuilder:validation:Required
	// Required if specifying collection.  The name of the collection
	// within a specific collection.namespace to reference.
	Name string `json:"name"`

	// +kubebuilder:validation:Optional
	// (Default: "") The namespace where the collection exists.  Required only if
	// the collection is namespace scoped and not cluster scoped.
	Namespace string `json:"namespace"`
}

type ExternalSecretsSpecCertController struct {
	// +kubebuilder:default=1
	// +kubebuilder:validation:Optional
	// (Default: 1)
	//  Number of replicas to use for the external-secrets cert-controller deployment.
	Replicas int `json:"replicas,omitempty"`
}

type ExternalSecretsSpecController struct {
	// +kubebuilder:default=2
	// +kubebuilder:validation:Optional
	// (Default: 2)
	//  Number of replicas to use for the external-secrets controller deployment.
	Replicas int `json:"replicas,omitempty"`
}

type ExternalSecretsSpecWebhook struct {
	// +kubebuilder:default=2
	// +kubebuilder:validation:Optional
	// (Default: 2)
	//  Number of replicas to use for the external-secrets webhook deployment.
	Replicas int `json:"replicas,omitempty"`
}

// ExternalSecretsStatus defines the observed state of ExternalSecrets.
type ExternalSecretsStatus struct {
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

// ExternalSecrets is the Schema for the externalsecrets API.
type ExternalSecrets struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ExternalSecretsSpec   `json:"spec,omitempty"`
	Status            ExternalSecretsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ExternalSecretsList contains a list of ExternalSecrets.
type ExternalSecretsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExternalSecrets `json:"items"`
}

// interface methods

// GetReadyStatus returns the ready status for a component.
func (component *ExternalSecrets) GetReadyStatus() bool {
	return component.Status.Created
}

// SetReadyStatus sets the ready status for a component.
func (component *ExternalSecrets) SetReadyStatus(ready bool) {
	component.Status.Created = ready
}

// GetDependencyStatus returns the dependency status for a component.
func (component *ExternalSecrets) GetDependencyStatus() bool {
	return component.Status.DependenciesSatisfied
}

// SetDependencyStatus sets the dependency status for a component.
func (component *ExternalSecrets) SetDependencyStatus(dependencyStatus bool) {
	component.Status.DependenciesSatisfied = dependencyStatus
}

// GetPhaseConditions returns the phase conditions for a component.
func (component *ExternalSecrets) GetPhaseConditions() []*status.PhaseCondition {
	return component.Status.Conditions
}

// SetPhaseCondition sets the phase conditions for a component.
func (component *ExternalSecrets) SetPhaseCondition(condition *status.PhaseCondition) {
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
func (component *ExternalSecrets) GetChildResourceConditions() []*status.ChildResource {
	return component.Status.Resources
}

// SetResources sets the phase conditions for a component.
func (component *ExternalSecrets) SetChildResourceCondition(resource *status.ChildResource) {
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
func (*ExternalSecrets) GetDependencies() []workload.Workload {
	return []workload.Workload{}
}

// GetComponentGVK returns a GVK object for the component.
func (*ExternalSecrets) GetWorkloadGVK() schema.GroupVersionKind {
	return GroupVersion.WithKind("ExternalSecrets")
}

func init() {
	SchemeBuilder.Register(&ExternalSecrets{}, &ExternalSecretsList{})
}
