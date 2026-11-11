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
	NamespaceNamespace                                        = "parent.Spec.Namespace"
	ConfigMapNamespaceIstioCniConfig                          = "istio-cni-config"
	ConfigMapNamespaceValues                                  = "values"
	ConfigMapNamespaceIstio                                   = "istio"
	ConfigMapNamespaceIstioSidecarInjector                    = "istio-sidecar-injector"
	ConfigMapNamespaceIstioCniConfigGke                       = "istio-cni-config-gke"
	ConfigMapNamespaceValuesGke                               = "values-gke"
	ConfigMapNamespaceIstioGke                                = "istio-gke"
	ConfigMapNamespaceIstioSidecarInjectorGke                 = "istio-sidecar-injector-gke"
	ConfigMapNamespaceIstioCniConfigOpenshift                 = "istio-cni-config-openshift"
	ConfigMapNamespaceValuesOpenshift                         = "values-openshift"
	ConfigMapNamespaceIstioOpenshift                          = "istio-openshift"
	ConfigMapNamespaceIstioSidecarInjectorOpenshift           = "istio-sidecar-injector-openshift"
	CRDTrafficextensionsExtensionsIstioIo                     = "trafficextensions.extensions.istio.io"
	CRDWasmpluginsExtensionsIstioIo                           = "wasmplugins.extensions.istio.io"
	CRDDestinationrulesNetworkingIstioIo                      = "destinationrules.networking.istio.io"
	CRDEnvoyfiltersNetworkingIstioIo                          = "envoyfilters.networking.istio.io"
	CRDGatewaysNetworkingIstioIo                              = "gateways.networking.istio.io"
	CRDProxyconfigsNetworkingIstioIo                          = "proxyconfigs.networking.istio.io"
	CRDServiceentriesNetworkingIstioIo                        = "serviceentries.networking.istio.io"
	CRDSidecarsNetworkingIstioIo                              = "sidecars.networking.istio.io"
	CRDVirtualservicesNetworkingIstioIo                       = "virtualservices.networking.istio.io"
	CRDWorkloadentriesNetworkingIstioIo                       = "workloadentries.networking.istio.io"
	CRDWorkloadgroupsNetworkingIstioIo                        = "workloadgroups.networking.istio.io"
	CRDAuthorizationpoliciesSecurityIstioIo                   = "authorizationpolicies.security.istio.io"
	CRDPeerauthenticationsSecurityIstioIo                     = "peerauthentications.security.istio.io"
	CRDRequestauthenticationsSecurityIstioIo                  = "requestauthentications.security.istio.io"
	CRDTelemetriesTelemetryIstioIo                            = "telemetries.telemetry.istio.io"
	ServiceAccountNamespaceIstioReaderServiceAccount          = "istio-reader-service-account"
	ServiceAccountNamespaceIstioCni                           = "istio-cni"
	ServiceAccountNamespaceIstiod                             = "istiod"
	ServiceAccountNamespaceZtunnel                            = "ztunnel"
	ClusterRoleIstioCni                                       = "istio-cni"
	ClusterRoleIstioCniRepairRole                             = "istio-cni-repair-role"
	ClusterRoleIstioCniAmbient                                = "istio-cni-ambient"
	ClusterRoleBindingIstioCni                                = "istio-cni"
	ClusterRoleBindingIstioCniRepairRolebinding               = "istio-cni-repair-rolebinding"
	ClusterRoleBindingIstioCniAmbient                         = "istio-cni-ambient"
	RoleNamespaceIstiod                                       = "istiod"
	RoleBindingNamespaceIstiod                                = "istiod"
	ServiceAccountNamespaceIstioReaderServiceAccountOpenshift = "istio-reader-service-account-openshift"
	ServiceAccountNamespaceIstioCniOpenshift                  = "istio-cni-openshift"
	ServiceAccountNamespaceIstiodOpenshift                    = "istiod-openshift"
	ServiceAccountNamespaceZtunnelOpenshift                   = "ztunnel-openshift"
	ClusterRoleIstioCniOpenshift                              = "istio-cni-openshift"
	ClusterRoleIstioCniRepairRoleOpenshift                    = "istio-cni-repair-role-openshift"
	ClusterRoleIstioCniAmbientOpenshift                       = "istio-cni-ambient-openshift"
	ClusterRoleZtunnelOpenshift                               = "ztunnel-openshift"
	ClusterRoleBindingIstioCniOpenshift                       = "istio-cni-openshift"
	ClusterRoleBindingIstioCniRepairRolebindingOpenshift      = "istio-cni-repair-rolebinding-openshift"
	ClusterRoleBindingIstioCniAmbientOpenshift                = "istio-cni-ambient-openshift"
	ClusterRoleBindingZtunnelOpenshift                        = "ztunnel-openshift"
	RoleNamespaceIstiodOpenshift                              = "istiod-openshift"
	RoleBindingNamespaceIstiodOpenshift                       = "istiod-openshift"
	ServiceNamespaceIstiod                                    = "istiod"
	ValidatingWebhookIstiodDefaultValidator                   = "istiod-default-validator"
	DeploymentNamespaceIstiod                                 = "istiod"
	DeploymentNamespaceIstiodGke                              = "istiod-gke"
	DeploymentNamespaceIstiodOpenshift                        = "istiod-openshift"
	HorizontalPodAutoscalerNamespaceIstiod                    = "istiod"
	HorizontalPodAutoscalerNamespaceIstiodGke                 = "istiod-gke"
	HorizontalPodAutoscalerNamespaceIstiodOpenshift           = "istiod-openshift"
	PDBNamespaceIstiod                                        = "istiod"
	DaemonSetNamespaceIstioCniNode                            = "istio-cni-node"
	DaemonSetNamespaceZtunnel                                 = "ztunnel"
	DaemonSetNamespaceIstioCniNodeGke                         = "istio-cni-node-gke"
	DaemonSetNamespaceZtunnelGke                              = "ztunnel-gke"
	DaemonSetNamespaceIstioCniNodeOpenshift                   = "istio-cni-node-openshift"
	DaemonSetNamespaceZtunnelOpenshift                        = "ztunnel-openshift"
	NetworkPolicyNamespaceIstioCni                            = "istio-cni"
	NetworkPolicyNamespaceIstiod                              = "istiod"
	NetworkPolicyNamespaceZtunnel                             = "ztunnel"
	ResourceQuotaNamespaceIstioCniResourceQuota               = "istio-cni-resource-quota"
	ResourceQuotaNamespaceZtunnel                             = "ztunnel"
	NetworkAttachmentDefinitionDefaultIstioCni                = "istio-cni"
)
