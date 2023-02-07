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

package secretscomponent

import (
	"fmt"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/yaml"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	platformv1alpha1 "github.com/nukleros/support-services-operator/apis/platform/v1alpha1"
	setupv1alpha1 "github.com/nukleros/support-services-operator/apis/setup/v1alpha1"
)

// sampleSecretsComponent is a sample containing all fields
const sampleSecretsComponent = `apiVersion: platform.addons.nukleros.io/v1alpha1
kind: SecretsComponent
metadata:
  name: secretscomponent-sample
spec:
  #collection:
    #name: "supportservices-sample"
    #namespace: ""
  namespace: "nukleros-secrets-system"
  externalSecrets:
    version: "v0.5.9"
    certController:
      replicas: 1
    image: "ghcr.io/external-secrets/external-secrets"
    controller:
      replicas: 2
    webhook:
      replicas: 2
  reloader:
    replicas: 1
    image: "stakater/reloader"
    version: "v0.0.119"
`

// sampleSecretsComponentRequired is a sample containing only required fields
const sampleSecretsComponentRequired = `apiVersion: platform.addons.nukleros.io/v1alpha1
kind: SecretsComponent
metadata:
  name: secretscomponent-sample
spec:
  #collection:
    #name: "supportservices-sample"
    #namespace: ""
`

// Sample returns the sample manifest for this custom resource.
func Sample(requiredOnly bool) string {
	if requiredOnly {
		return sampleSecretsComponentRequired
	}

	return sampleSecretsComponent
}

// Generate returns the child resources that are associated with this workload given
// appropriate structured inputs.
func Generate(
	workloadObj platformv1alpha1.SecretsComponent,
	collectionObj setupv1alpha1.SupportServices,
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
	var workloadObj platformv1alpha1.SecretsComponent
	if err := yaml.Unmarshal(workloadFile, &workloadObj); err != nil {
		return nil, fmt.Errorf("failed to unmarshal yaml into workload, %w", err)
	}

	if err := workload.Validate(&workloadObj); err != nil {
		return nil, fmt.Errorf("error validating workload yaml, %w", err)
	}

	var collectionObj setupv1alpha1.SupportServices
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
	*platformv1alpha1.SecretsComponent,
	*setupv1alpha1.SupportServices,
	workload.Reconciler,
	*workload.Request,
) ([]client.Object, error){
	CreateNamespaceNamespace,
	CreateSecretNamespaceExternalSecretsWebhook,
	CreateCRDClusterexternalsecretsExternalSecretsIo,
	CreateCRDClustersecretstoresExternalSecretsIo,
	CreateCRDExternalsecretsExternalSecretsIo,
	CreateCRDSecretstoresExternalSecretsIo,
	CreateDeploymentNamespaceExternalSecretsCertController,
	CreateDeploymentNamespaceExternalSecrets,
	CreateDeploymentNamespaceExternalSecretsWebhook,
	CreateServiceAccountNamespaceExternalSecretsCertController,
	CreateServiceAccountNamespaceExternalSecrets,
	CreateServiceAccountNamespaceExternalSecretsWebhook,
	CreateClusterRoleExternalSecretsCertController,
	CreateClusterRoleExternalSecretsController,
	CreateClusterRoleExternalSecretsView,
	CreateClusterRoleExternalSecretsEdit,
	CreateClusterRoleBindingExternalSecretsCertController,
	CreateClusterRoleBindingExternalSecretsController,
	CreateRoleNamespaceExternalSecretsLeaderelection,
	CreateRoleBindingNamespaceExternalSecretsLeaderelection,
	CreateServiceNamespaceExternalSecretsWebhook,
	CreateValidatingWebhookSecretstoreValidate,
	CreateValidatingWebhookExternalsecretValidate,
	CreateDeploymentNamespaceSecretReloader,
	CreateServiceAccountNamespaceSecretReloader,
	CreateClusterRoleNamespaceSecretReloader,
	CreateClusterRoleBindingNamespaceSecretReloader,
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
	*platformv1alpha1.SecretsComponent,
	*setupv1alpha1.SupportServices,
	workload.Reconciler,
	*workload.Request,
) ([]client.Object, error){
	CreateCRDClusterexternalsecretsExternalSecretsIo,
	CreateCRDClustersecretstoresExternalSecretsIo,
	CreateCRDExternalsecretsExternalSecretsIo,
	CreateCRDSecretstoresExternalSecretsIo,
}

func ConvertWorkload(component, collection workload.Workload) (
	*platformv1alpha1.SecretsComponent,
	*setupv1alpha1.SupportServices,
	error,
) {
	p, ok := component.(*platformv1alpha1.SecretsComponent)
	if !ok {
		return nil, nil, platformv1alpha1.ErrUnableToConvertSecretsComponent
	}

	c, ok := collection.(*setupv1alpha1.SupportServices)
	if !ok {
		return nil, nil, setupv1alpha1.ErrUnableToConvertSupportServices
	}

	return p, c, nil
}
