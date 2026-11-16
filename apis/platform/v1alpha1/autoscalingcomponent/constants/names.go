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
	NamespaceNamespace                                         = "parent.Spec.Namespace"
	CRDCloudeventsourcesEventingKedaSh                         = "cloudeventsources.eventing.keda.sh"
	CRDClustercloudeventsourcesEventingKedaSh                  = "clustercloudeventsources.eventing.keda.sh"
	CRDClustertriggerauthenticationsKedaSh                     = "clustertriggerauthentications.keda.sh"
	CRDScaledjobsKedaSh                                        = "scaledjobs.keda.sh"
	CRDScaledobjectsKedaSh                                     = "scaledobjects.keda.sh"
	CRDTriggerauthenticationsKedaSh                            = "triggerauthentications.keda.sh"
	ServiceAccountNamespaceKedaOperator                        = "keda-operator"
	ServiceAccountNamespaceKedaMetricsServer                   = "keda-metrics-server"
	ServiceAccountNamespaceKedaWebhook                         = "keda-webhook"
	ClusterRoleKedaOperator                                    = "keda-operator"
	ClusterRoleKedaOperatorMinimalClusterRole                  = "keda-operator-minimal-cluster-role"
	ClusterRoleKedaOperatorExternalMetricsReader               = "keda-operator-external-metrics-reader"
	ClusterRoleKedaOperatorWebhook                             = "keda-operator-webhook"
	ClusterRoleBindingKedaOperator                             = "keda-operator"
	ClusterRoleBindingKedaOperatorMinimal                      = "keda-operator-minimal"
	ClusterRoleBindingKedaOperatorSystemAuthDelegator          = "keda-operator-system-auth-delegator"
	ClusterRoleBindingKedaOperatorHpaControllerExternalMetrics = "keda-operator-hpa-controller-external-metrics"
	ClusterRoleBindingKedaOperatorWebhook                      = "keda-operator-webhook"
	RoleNamespaceKedaOperatorCerts                             = "keda-operator-certs"
	RoleBindingNamespaceKedaOperatorCerts                      = "keda-operator-certs"
	RoleBindingKubeSystemKedaOperatorAuthReader                = "keda-operator-auth-reader"
	ServiceNamespaceKedaOperator                               = "keda-operator"
	ServiceNamespaceKedaOperatorMetricsApiserver               = "keda-operator-metrics-apiserver"
	ServiceNamespaceKedaAdmissionWebhooks                      = "keda-admission-webhooks"
	DeploymentNamespaceKedaOperator                            = "keda-operator"
	DeploymentNamespaceKedaOperatorMetricsApiserver            = "keda-operator-metrics-apiserver"
	DeploymentNamespaceKedaAdmissionWebhooks                   = "keda-admission-webhooks"
	PDBNamespaceKedaOperator                                   = "keda-operator"
	PDBNamespaceKedaOperatorMetricsApiserver                   = "keda-operator-metrics-apiserver"
	PDBNamespaceKedaAdmissionWebhooks                          = "keda-admission-webhooks"
	NetworkPolicyNamespaceKedaOperator                         = "keda-operator"
	NetworkPolicyNamespaceKedaOperatorMetricsApiserver         = "keda-operator-metrics-apiserver"
	NetworkPolicyNamespaceKedaAdmissionWebhooks                = "keda-admission-webhooks"
	ValidatingWebhookKedaAdmission                             = "keda-admission"
	APIServiceV1beta1ExternalMetricsK8sIo                      = "v1beta1.external.metrics.k8s.io"
)
