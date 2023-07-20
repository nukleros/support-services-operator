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
	CertNuklerosGatewaySystemCertificateAuthority            = "certificate-authority"
	ClusterIssuerSelfSigned                                  = "self-signed"
	ConfigMapNuklerosGatewaySystemGlooCustomResourceConfig   = "gloo-custom-resource-config"
	ConfigMapNuklerosGatewaySystemGatewayProxyEnvoyConfig    = "gateway-proxy-envoy-config"
	CRDAuthconfigsEnterpriseGlooSoloIo                       = "authconfigs.enterprise.gloo.solo.io"
	CRDGatewaysGatewaySoloIo                                 = "gateways.gateway.solo.io"
	CRDHttpgatewaysGatewaySoloIo                             = "httpgateways.gateway.solo.io"
	CRDTcpgatewaysGatewaySoloIo                              = "tcpgateways.gateway.solo.io"
	CRDRouteoptionsGatewaySoloIo                             = "routeoptions.gateway.solo.io"
	CRDRoutetablesGatewaySoloIo                              = "routetables.gateway.solo.io"
	CRDVirtualhostoptionsGatewaySoloIo                       = "virtualhostoptions.gateway.solo.io"
	CRDVirtualservicesGatewaySoloIo                          = "virtualservices.gateway.solo.io"
	CRDProxiesGlooSoloIo                                     = "proxies.gloo.solo.io"
	CRDSettingsGlooSoloIo                                    = "settings.gloo.solo.io"
	CRDUpstreamsGlooSoloIo                                   = "upstreams.gloo.solo.io"
	CRDUpstreamgroupsGlooSoloIo                              = "upstreamgroups.gloo.solo.io"
	CRDGraphqlapisGraphqlGlooSoloIo                          = "graphqlapis.graphql.gloo.solo.io"
	CRDRatelimitconfigsRatelimitSoloIo                       = "ratelimitconfigs.ratelimit.solo.io"
	DeploymentNuklerosGatewaySystemGloo                      = "gloo"
	DeploymentNuklerosGatewaySystemDiscovery                 = "discovery"
	DeploymentNuklerosGatewaySystemGatewayProxy              = "gateway-proxy"
	GatewayNuklerosGatewaySystemGatewayProxy                 = "gateway-proxy"
	GatewayNuklerosGatewaySystemGatewayProxySsl              = "gateway-proxy-ssl"
	IssuerNuklerosGatewaySystemCertificateAuthority          = "certificate-authority"
	NamespaceNuklerosGatewaySystem                           = "nukleros-gateway-system"
	ServiceAccountNuklerosGatewaySystemGlooResourceCleanup   = "gloo-resource-cleanup"
	ServiceAccountNuklerosGatewaySystemGlooResourceMigration = "gloo-resource-migration"
	ServiceAccountNuklerosGatewaySystemGlooResourceRollout   = "gloo-resource-rollout"
	ServiceAccountNuklerosGatewaySystemCertgen               = "certgen"
	ServiceAccountNuklerosGatewaySystemGloo                  = "gloo"
	ServiceAccountNuklerosGatewaySystemDiscovery             = "discovery"
	ServiceAccountNuklerosGatewaySystemGatewayProxy          = "gateway-proxy"
	ClusterRoleKubeResourceWatcherDefault                    = "kube-resource-watcher-default"
	ClusterRoleKubeLeaderElectionDefault                     = "kube-leader-election-default"
	ClusterRoleGlooUpstreamMutatorDefault                    = "gloo-upstream-mutator-default"
	ClusterRoleGlooResourceReaderDefault                     = "gloo-resource-reader-default"
	ClusterRoleSettingsUserDefault                           = "settings-user-default"
	ClusterRoleGlooResourceMutatorDefault                    = "gloo-resource-mutator-default"
	ClusterRoleGatewayResourceReaderDefault                  = "gateway-resource-reader-default"
	ClusterRoleGlooGraphqlapiMutatorDefault                  = "gloo-graphqlapi-mutator-default"
	ClusterRoleBindingKubeResourceWatcherBindingDefault      = "kube-resource-watcher-binding-default"
	ClusterRoleBindingKubeLeaderElectionBindingDefault       = "kube-leader-election-binding-default"
	ClusterRoleBindingGlooUpstreamMutatorBindingDefault      = "gloo-upstream-mutator-binding-default"
	ClusterRoleBindingGlooResourceReaderBindingDefault       = "gloo-resource-reader-binding-default"
	ClusterRoleBindingSettingsUserBindingDefault             = "settings-user-binding-default"
	ClusterRoleBindingGlooResourceMutatorBindingDefault      = "gloo-resource-mutator-binding-default"
	ClusterRoleBindingGatewayResourceReaderBindingDefault    = "gateway-resource-reader-binding-default"
	ClusterRoleBindingGlooGraphqlapiMutatorBindingDefault    = "gloo-graphqlapi-mutator-binding-default"
	ClusterRoleGlooResourceCleanupDefault                    = "gloo-resource-cleanup-default"
	ClusterRoleGlooResourceMigrationDefault                  = "gloo-resource-migration-default"
	ClusterRoleGlooResourceRolloutDefault                    = "gloo-resource-rollout-default"
	ClusterRoleGlooGatewayVwcUpdateDefault                   = "gloo-gateway-vwc-update-default"
	ClusterRoleGlooGatewaySecretCreateDefault                = "gloo-gateway-secret-create-default"
	ClusterRoleBindingGlooResourceCleanupDefault             = "gloo-resource-cleanup-default"
	ClusterRoleBindingGlooResourceMigrationDefault           = "gloo-resource-migration-default"
	ClusterRoleBindingGlooResourceRolloutDefault             = "gloo-resource-rollout-default"
	ClusterRoleBindingGlooGatewayVwcUpdateDefault            = "gloo-gateway-vwc-update-default"
	ClusterRoleBindingGlooGatewaySecretCreateDefault         = "gloo-gateway-secret-create-default"
	RoleNuklerosGatewaySystemGlooResourceCleanup             = "gloo-resource-cleanup"
	RoleNuklerosGatewaySystemGlooResourceMigration           = "gloo-resource-migration"
	RoleNuklerosGatewaySystemGlooResourceRollout             = "gloo-resource-rollout"
	RoleBindingNuklerosGatewaySystemGlooResourceCleanup      = "gloo-resource-cleanup"
	RoleBindingNuklerosGatewaySystemGlooResourceMigration    = "gloo-resource-migration"
	RoleBindingNuklerosGatewaySystemGlooResourceRollout      = "gloo-resource-rollout"
	ServiceNuklerosGatewaySystemGloo                         = "gloo"
	ServiceNuklerosGatewaySystemGatewayProxy                 = "gateway-proxy"
	SettingsNuklerosGatewaySystemDefault                     = "default"
	ValidatingWebhookGlooGatewayValidationWebhookDefault     = "gloo-gateway-validation-webhook-default"
)
