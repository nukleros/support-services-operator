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

package constants

// this package includes the constants which include the resource names.  it is a standalone
// package to prevent import cycle errors when attempting to reference the names from other
// packages (e.g. mutate).
const (
	NamespaceNamespace                                                         = "parent.Spec.Namespace"
	SecretNamespaceExternalSecretsWebhook                                      = "external-secrets-webhook"
	CRDAcraccesstokensGeneratorsExternalSecretsIo                              = "acraccesstokens.generators.external-secrets.io"
	CRDBeyondtrustworkloadcredentialsdynamicsecretsGeneratorsExternalSecretsIo = "beyondtrustworkloadcredentialsdynamicsecrets.generators.external-secrets.io"
	CRDCloudsmithaccesstokensGeneratorsExternalSecretsIo                       = "cloudsmithaccesstokens.generators.external-secrets.io"
	CRDClusterexternalsecretsExternalSecretsIo                                 = "clusterexternalsecrets.external-secrets.io"
	CRDClustergeneratorsGeneratorsExternalSecretsIo                            = "clustergenerators.generators.external-secrets.io"
	CRDClusterpushsecretsExternalSecretsIo                                     = "clusterpushsecrets.external-secrets.io"
	CRDClustersecretstoresExternalSecretsIo                                    = "clustersecretstores.external-secrets.io"
	CRDEcrauthorizationtokensGeneratorsExternalSecretsIo                       = "ecrauthorizationtokens.generators.external-secrets.io"
	CRDExternalsecretsExternalSecretsIo                                        = "externalsecrets.external-secrets.io"
	CRDFakesGeneratorsExternalSecretsIo                                        = "fakes.generators.external-secrets.io"
	CRDGcraccesstokensGeneratorsExternalSecretsIo                              = "gcraccesstokens.generators.external-secrets.io"
	CRDGeneratorstatesGeneratorsExternalSecretsIo                              = "generatorstates.generators.external-secrets.io"
	CRDGithubaccesstokensGeneratorsExternalSecretsIo                           = "githubaccesstokens.generators.external-secrets.io"
	CRDGitlabdeploytokensGeneratorsExternalSecretsIo                           = "gitlabdeploytokens.generators.external-secrets.io"
	CRDGrafanasGeneratorsExternalSecretsIo                                     = "grafanas.generators.external-secrets.io"
	CRDMfasGeneratorsExternalSecretsIo                                         = "mfas.generators.external-secrets.io"
	CRDPasswordsGeneratorsExternalSecretsIo                                    = "passwords.generators.external-secrets.io"
	CRDPushsecretsExternalSecretsIo                                            = "pushsecrets.external-secrets.io"
	CRDQuayaccesstokensGeneratorsExternalSecretsIo                             = "quayaccesstokens.generators.external-secrets.io"
	CRDSecretstoresExternalSecretsIo                                           = "secretstores.external-secrets.io"
	CRDSshkeysGeneratorsExternalSecretsIo                                      = "sshkeys.generators.external-secrets.io"
	CRDStssessiontokensGeneratorsExternalSecretsIo                             = "stssessiontokens.generators.external-secrets.io"
	CRDUuidsGeneratorsExternalSecretsIo                                        = "uuids.generators.external-secrets.io"
	CRDVaultdynamicsecretsGeneratorsExternalSecretsIo                          = "vaultdynamicsecrets.generators.external-secrets.io"
	CRDWebhooksGeneratorsExternalSecretsIo                                     = "webhooks.generators.external-secrets.io"
	DeploymentNamespaceExternalSecretsCertController                           = "external-secrets-cert-controller"
	DeploymentNamespaceExternalSecrets                                         = "external-secrets"
	DeploymentNamespaceExternalSecretsWebhook                                  = "external-secrets-webhook"
	ServiceAccountNamespaceExternalSecretsCertController                       = "external-secrets-cert-controller"
	ServiceAccountNamespaceExternalSecrets                                     = "external-secrets"
	ServiceAccountNamespaceExternalSecretsWebhook                              = "external-secrets-webhook"
	ClusterRoleExternalSecretsCertController                                   = "external-secrets-cert-controller"
	ClusterRoleExternalSecretsController                                       = "external-secrets-controller"
	ClusterRoleExternalSecretsView                                             = "external-secrets-view"
	ClusterRoleExternalSecretsEdit                                             = "external-secrets-edit"
	ClusterRoleExternalSecretsServicebindings                                  = "external-secrets-servicebindings"
	ClusterRoleBindingExternalSecretsCertController                            = "external-secrets-cert-controller"
	ClusterRoleBindingExternalSecretsController                                = "external-secrets-controller"
	RoleNamespaceExternalSecretsLeaderelection                                 = "external-secrets-leaderelection"
	RoleBindingNamespaceExternalSecretsLeaderelection                          = "external-secrets-leaderelection"
	ServiceNamespaceExternalSecretsWebhook                                     = "external-secrets-webhook"
	ValidatingWebhookSecretstoreValidate                                       = "secretstore-validate"
	ValidatingWebhookExternalsecretValidate                                    = "externalsecret-validate"
	DeploymentNamespaceSecretReloader                                          = "secret-reloader"
	ServiceAccountNamespaceSecretReloader                                      = "secret-reloader"
	ClusterRoleNamespaceSecretReloader                                         = "secret-reloader"
	ClusterRoleBindingNamespaceSecretReloader                                  = "secret-reloader"
	SecretNamespaceOpenbaoInjectorCerts                                        = "openbao-injector-certs"
	ConfigMapNamespaceOpenbaoConfig                                            = "openbao-config"
	SecretNamespaceOpenbaoUnsealKeySecretName                                  = "parent.Spec.Openbao.UnsealKey.Secret.Name"
	RoleNamespaceOpenbaoInitSecretWriter                                       = "openbao-init-secret-writer"
	RoleBindingNamespaceOpenbaoInitSecretWriter                                = "openbao-init-secret-writer"
	ServiceAccountNamespaceOpenbaoAgentInjector                                = "openbao-agent-injector"
	ServiceAccountNamespaceOpenbao                                             = "openbao"
	ClusterRoleOpenbaoAgentInjectorClusterrole                                 = "openbao-agent-injector-clusterrole"
	ClusterRoleBindingOpenbaoAgentInjectorBinding                              = "openbao-agent-injector-binding"
	ClusterRoleBindingOpenbaoServerBinding                                     = "openbao-server-binding"
	RoleNamespaceOpenbaoAgentInjectorLeaderElectorRole                         = "openbao-agent-injector-leader-elector-role"
	RoleNamespaceOpenbaoDiscoveryRole                                          = "openbao-discovery-role"
	RoleBindingNamespaceOpenbaoAgentInjectorLeaderElectorBinding               = "openbao-agent-injector-leader-elector-binding"
	RoleBindingNamespaceOpenbaoDiscoveryRolebinding                            = "openbao-discovery-rolebinding"
	NetworkPolicyNamespaceOpenbao                                              = "openbao"
	PDBNamespaceOpenbao                                                        = "openbao"
	StatefulSetNamespaceOpenbao                                                = "openbao"
	DeploymentNamespaceOpenbaoAgentInjector                                    = "openbao-agent-injector"
	ServiceNamespaceOpenbaoAgentInjectorSvc                                    = "openbao-agent-injector-svc"
	ServiceNamespaceOpenbaoActive                                              = "openbao-active"
	ServiceNamespaceOpenbaoStandby                                             = "openbao-standby"
	ServiceNamespaceOpenbaoInternal                                            = "openbao-internal"
	ServiceNamespaceOpenbao                                                    = "openbao"
	ServiceNamespaceOpenbaoUi                                                  = "openbao-ui"
	MutatingWebhookOpenbaoAgentInjectorCfg                                     = "openbao-agent-injector-cfg"
)
