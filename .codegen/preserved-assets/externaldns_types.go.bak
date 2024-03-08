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

	certificatesv1alpha1 "github.com/nukleros/support-services-operator/apis/certificates/v1alpha1"
)

var ErrUnableToConvertExternalDNS = errors.New("unable to convert to ExternalDNS")

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ExternalDNSSpec defines the desired state of ExternalDNS.
type ExternalDNSSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +kubebuilder:validation:Optional
	// Specifies a reference to the collection to use for this workload.
	// Requires the name and namespace input to find the collection.
	// If no collection field is set, default to selecting the only
	// workload collection in the cluster, which will result in an error
	// if not exactly one collection is found.
	Collection ExternalDNSCollectionSpec `json:"collection"`

	// +kubebuilder:default="nukleros-gateway-system"
	// +kubebuilder:validation:Optional
	// (Default: "nukleros-gateway-system")
	//  Namespace to use for ingress support services.
	Namespace string `json:"namespace,omitempty"`

	// +kubebuilder:default="private"
	// +kubebuilder:validation:Optional
	// (Default: "private")
	//  +kubebuilder:validation:Enum=private;public
	//  Type of DNS hosted zone to manage.
	ZoneType string `json:"zoneType,omitempty"`

	// +kubebuilder:validation:Required
	DomainName string `json:"domainName,omitempty"`

	// +kubebuilder:default="k8s.gcr.io/external-dns/external-dns"
	// +kubebuilder:validation:Optional
	// (Default: "k8s.gcr.io/external-dns/external-dns")
	//  Image repo and name to use for external-dns.
	Image string `json:"image,omitempty"`

	// +kubebuilder:default="v0.12.2"
	// +kubebuilder:validation:Optional
	// (Default: "v0.12.2")
	//  Version of external-dns to use.
	Version string `json:"version,omitempty"`

	// +kubebuilder:default="none"
	// +kubebuilder:validation:Optional
	// (Default: "none")
	//  +kubebuilder:validation:Enum=none;active-directory;google;route53
	//  The DNS provider to use for setting DNS records with external-dns.  One of: none | active-directory | google | route53.
	Provider string `json:"provider,omitempty"`

	// +kubebuilder:default="external-dns"
	// +kubebuilder:validation:Optional
	// (Default: "external-dns")
	//  The name of the external-dns service account which is referenced in role policy doc for AWS.
	ServiceAccountName string `json:"serviceAccountName,omitempty"`

	// +kubebuilder:validation:Required
	//  On AWS, the IAM Role ARN that gives external-dns access to Route53
	IamRoleArn string `json:"iamRoleArn,omitempty"`

	// +kubebuilder:validation:Optional
	//  Extra arguments to be passed into the External DNS container.
	ExtraArgs []string `json:"extraArgs,omitempty"`
}

type ExternalDNSCollectionSpec struct {
	// +kubebuilder:validation:Required
	// Required if specifying collection.  The name of the collection
	// within a specific collection.namespace to reference.
	Name string `json:"name"`

	// +kubebuilder:validation:Optional
	// (Default: "") The namespace where the collection exists.  Required only if
	// the collection is namespace scoped and not cluster scoped.
	Namespace string `json:"namespace"`
}

// ExternalDNSStatus defines the observed state of ExternalDNS.
type ExternalDNSStatus struct {
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

// ExternalDNS is the Schema for the externaldns API.
type ExternalDNS struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ExternalDNSSpec   `json:"spec,omitempty"`
	Status            ExternalDNSStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ExternalDNSList contains a list of ExternalDNS.
type ExternalDNSList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExternalDNS `json:"items"`
}

// interface methods

// GetReadyStatus returns the ready status for a component.
func (component *ExternalDNS) GetReadyStatus() bool {
	return component.Status.Created
}

// SetReadyStatus sets the ready status for a component.
func (component *ExternalDNS) SetReadyStatus(ready bool) {
	component.Status.Created = ready
}

// GetDependencyStatus returns the dependency status for a component.
func (component *ExternalDNS) GetDependencyStatus() bool {
	return component.Status.DependenciesSatisfied
}

// SetDependencyStatus sets the dependency status for a component.
func (component *ExternalDNS) SetDependencyStatus(dependencyStatus bool) {
	component.Status.DependenciesSatisfied = dependencyStatus
}

// GetPhaseConditions returns the phase conditions for a component.
func (component *ExternalDNS) GetPhaseConditions() []*status.PhaseCondition {
	return component.Status.Conditions
}

// SetPhaseCondition sets the phase conditions for a component.
func (component *ExternalDNS) SetPhaseCondition(condition *status.PhaseCondition) {
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
func (component *ExternalDNS) GetChildResourceConditions() []*status.ChildResource {
	return component.Status.Resources
}

// SetResources sets the phase conditions for a component.
func (component *ExternalDNS) SetChildResourceCondition(resource *status.ChildResource) {
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
func (*ExternalDNS) GetDependencies() []workload.Workload {
	return []workload.Workload{
		&certificatesv1alpha1.CertManager{},
	}
}

// GetComponentGVK returns a GVK object for the component.
func (*ExternalDNS) GetWorkloadGVK() schema.GroupVersionKind {
	return GroupVersion.WithKind("ExternalDNS")
}

func init() {
	SchemeBuilder.Register(&ExternalDNS{}, &ExternalDNSList{})
}
