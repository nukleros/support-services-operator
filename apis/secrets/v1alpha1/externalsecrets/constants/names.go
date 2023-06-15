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

package constants

// this package includes the constants which include the resource names.  it is a standalone
// package to prevent import cycle errors when attempting to reference the names from other
// packages (e.g. mutate).
const (
	SecretNamespaceExternalSecretsWebhook                = "external-secrets-webhook"
	CRDClusterexternalsecretsExternalSecretsIo           = "clusterexternalsecrets.external-secrets.io"
	CRDClustersecretstoresExternalSecretsIo              = "clustersecretstores.external-secrets.io"
	CRDExternalsecretsExternalSecretsIo                  = "externalsecrets.external-secrets.io"
	CRDSecretstoresExternalSecretsIo                     = "secretstores.external-secrets.io"
	DeploymentNamespaceExternalSecretsCertController     = "external-secrets-cert-controller"
	DeploymentNamespaceExternalSecrets                   = "external-secrets"
	DeploymentNamespaceExternalSecretsWebhook            = "external-secrets-webhook"
	NamespaceNamespace                                   = "parent.Spec.Namespace"
	ServiceAccountNamespaceExternalSecretsCertController = "external-secrets-cert-controller"
	ServiceAccountNamespaceExternalSecrets               = "external-secrets"
	ServiceAccountNamespaceExternalSecretsWebhook        = "external-secrets-webhook"
	ClusterRoleExternalSecretsCertController             = "external-secrets-cert-controller"
	ClusterRoleExternalSecretsController                 = "external-secrets-controller"
	ClusterRoleExternalSecretsView                       = "external-secrets-view"
	ClusterRoleExternalSecretsEdit                       = "external-secrets-edit"
	ClusterRoleBindingExternalSecretsCertController      = "external-secrets-cert-controller"
	ClusterRoleBindingExternalSecretsController          = "external-secrets-controller"
	RoleNamespaceExternalSecretsLeaderelection           = "external-secrets-leaderelection"
	RoleBindingNamespaceExternalSecretsLeaderelection    = "external-secrets-leaderelection"
	ServiceNamespaceExternalSecretsWebhook               = "external-secrets-webhook"
	ValidatingWebhookSecretstoreValidate                 = "secretstore-validate"
	ValidatingWebhookExternalsecretValidate              = "externalsecret-validate"
)
