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

package certmanager

import (
	"fmt"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/yaml"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	certificatesv1alpha1 "github.com/nukleros/support-services-operator/apis/certificates/v1alpha1"
	orchestrationv1alpha1 "github.com/nukleros/support-services-operator/apis/orchestration/v1alpha1"
)

// sampleCertManager is a sample containing all fields
const sampleCertManager = `apiVersion: certificates.support-services.nukleros.io/v1alpha1
kind: CertManager
metadata:
  name: certmanager-sample
spec:
  #collection:
    #name: "supportservices-sample"
    #namespace: ""
  namespace: "nukleros-certs-system"
  cainjector:
    replicas: 2
    image: "quay.io/jetstack/cert-manager-cainjector"
  version: "v1.9.1"
  controller:
    replicas: 2
    image: "quay.io/jetstack/cert-manager-controller"
  webhook:
    replicas: 2
    image: "quay.io/jetstack/cert-manager-webhook"
  contactEmail: "admin@nukleros.io"
`

// sampleCertManagerRequired is a sample containing only required fields
const sampleCertManagerRequired = `apiVersion: certificates.support-services.nukleros.io/v1alpha1
kind: CertManager
metadata:
  name: certmanager-sample
spec:
  #collection:
    #name: "supportservices-sample"
    #namespace: ""
  contactEmail: "admin@nukleros.io"
`

// Sample returns the sample manifest for this custom resource.
func Sample(requiredOnly bool) string {
	if requiredOnly {
		return sampleCertManagerRequired
	}

	return sampleCertManager
}

// Generate returns the child resources that are associated with this workload given
// appropriate structured inputs.
func Generate(
	workloadObj certificatesv1alpha1.CertManager,
	collectionObj orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {
	resourceObjects := []client.Object{}

	for _, f := range CreateFuncs {
		resources, err := f(&workloadObj, &collectionObj, reconciler, req)

		if err != nil {
			return nil, err
		}

		resourceObjects = append(resourceObjects, resources...)
	}

	return resourceObjects, nil
}

// GenerateForCLI returns the child resources that are associated with this workload given
// appropriate YAML manifest files.
func GenerateForCLI(workloadFile []byte, collectionFile []byte) ([]client.Object, error) {
	var workloadObj certificatesv1alpha1.CertManager
	if err := yaml.Unmarshal(workloadFile, &workloadObj); err != nil {
		return nil, fmt.Errorf("failed to unmarshal yaml into workload, %w", err)
	}

	if err := workload.Validate(&workloadObj); err != nil {
		return nil, fmt.Errorf("error validating workload yaml, %w", err)
	}

	var collectionObj orchestrationv1alpha1.SupportServices
	if err := yaml.Unmarshal(collectionFile, &collectionObj); err != nil {
		return nil, fmt.Errorf("failed to unmarshal yaml into collection, %w", err)
	}

	if err := workload.Validate(&collectionObj); err != nil {
		return nil, fmt.Errorf("error validating collection yaml, %w", err)
	}

	return Generate(workloadObj, collectionObj, nil, nil)
}

// CreateFuncs is an array of functions that are called to create the child resources for the controller
// in memory during the reconciliation loop prior to persisting the changes or updates to the Kubernetes
// database.
var CreateFuncs = []func(
	*certificatesv1alpha1.CertManager,
	*orchestrationv1alpha1.SupportServices,
	workload.Reconciler,
	*workload.Request,
) ([]client.Object, error){
	CreateCRDCertificaterequestsCertManagerIo,
	CreateCRDCertificatesCertManagerIo,
	CreateCRDChallengesAcmeCertManagerIo,
	CreateCRDClusterissuersCertManagerIo,
	CreateCRDIssuersCertManagerIo,
	CreateCRDOrdersAcmeCertManagerIo,
	CreateDeploymentNamespaceCertManagerCainjector,
	CreateDeploymentNamespaceCertManager,
	CreateDeploymentNamespaceCertManagerWebhook,
	CreateClusterIssuerLetsencryptStaging,
	CreateClusterIssuerLetsencryptProduction,
	CreateNamespaceNamespace,
	CreateServiceAccountNamespaceCertManagerCainjector,
	CreateServiceAccountNamespaceCertManager,
	CreateServiceAccountNamespaceCertManagerWebhook,
	CreateClusterRoleCertManagerCainjector,
	CreateClusterRoleCertManagerControllerIssuers,
	CreateClusterRoleCertManagerControllerClusterissuers,
	CreateClusterRoleCertManagerControllerCertificates,
	CreateClusterRoleCertManagerControllerOrders,
	CreateClusterRoleCertManagerControllerChallenges,
	CreateClusterRoleCertManagerControllerIngressShim,
	CreateClusterRoleCertManagerView,
	CreateClusterRoleCertManagerEdit,
	CreateClusterRoleCertManagerControllerApproveCertManagerIo,
	CreateClusterRoleCertManagerControllerCertificatesigningrequests,
	CreateClusterRoleCertManagerWebhookSubjectaccessreviews,
	CreateClusterRoleBindingCertManagerCainjector,
	CreateClusterRoleBindingCertManagerControllerIssuers,
	CreateClusterRoleBindingCertManagerControllerClusterissuers,
	CreateClusterRoleBindingCertManagerControllerCertificates,
	CreateClusterRoleBindingCertManagerControllerOrders,
	CreateClusterRoleBindingCertManagerControllerChallenges,
	CreateClusterRoleBindingCertManagerControllerIngressShim,
	CreateClusterRoleBindingCertManagerControllerApproveCertManagerIo,
	CreateClusterRoleBindingCertManagerControllerCertificatesigningrequests,
	CreateClusterRoleBindingCertManagerWebhookSubjectaccessreviews,
	CreateRoleNamespaceCertManagerCainjectorLeaderelection,
	CreateRoleNamespaceCertManagerLeaderelection,
	CreateRoleNamespaceCertManagerWebhookDynamicServing,
	CreateRoleBindingNamespaceCertManagerCainjectorLeaderelection,
	CreateRoleBindingNamespaceCertManagerLeaderelection,
	CreateRoleBindingNamespaceCertManagerWebhookDynamicServing,
	CreateServiceNamespaceCertManager,
	CreateServiceNamespaceCertManagerWebhook,
	CreateMutatingWebhookCertManagerWebhook,
	CreateValidatingWebhookCertManagerWebhook,
}

// InitFuncs is an array of functions that are called prior to starting the controller manager.  This is
// necessary in instances which the controller needs to "own" objects which depend on resources to
// pre-exist in the cluster. A common use case for this is the need to own a custom resource.
// If the controller needs to own a custom resource type, the CRD that defines it must
// first exist. In this case, the InitFunc will create the CRD so that the controller
// can own custom resources of that type.  Without the InitFunc the controller will
// crash loop because when it tries to own a non-existent resource type during manager
// setup, it will fail.
var InitFuncs = []func(
	*certificatesv1alpha1.CertManager,
	*orchestrationv1alpha1.SupportServices,
	workload.Reconciler,
	*workload.Request,
) ([]client.Object, error){
	CreateCRDCertificaterequestsCertManagerIo,
	CreateCRDCertificatesCertManagerIo,
	CreateCRDChallengesAcmeCertManagerIo,
	CreateCRDClusterissuersCertManagerIo,
	CreateCRDIssuersCertManagerIo,
	CreateCRDOrdersAcmeCertManagerIo,
}

func ConvertWorkload(component, collection workload.Workload) (
	*certificatesv1alpha1.CertManager,
	*orchestrationv1alpha1.SupportServices,
	error,
) {
	p, ok := component.(*certificatesv1alpha1.CertManager)
	if !ok {
		return nil, nil, certificatesv1alpha1.ErrUnableToConvertCertManager
	}

	c, ok := collection.(*orchestrationv1alpha1.SupportServices)
	if !ok {
		return nil, nil, orchestrationv1alpha1.ErrUnableToConvertSupportServices
	}

	return p, c, nil
}
