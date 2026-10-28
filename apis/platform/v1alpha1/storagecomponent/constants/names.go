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
	NamespaceNamespace                                               = "parent.Spec.Namespace"
	CertNamespaceSeaweedfsOperatorWebhookCert                        = "seaweedfs-operator-webhook-cert"
	IssuerNamespaceSeaweedfsOperatorWebhookSelfsignedIssuer          = "seaweedfs-operator-webhook-selfsigned-issuer"
	CRDAdminscriptsSeaweedSeaweedfsCom                               = "adminscripts.seaweed.seaweedfs.com"
	CRDBucketlifecyclepoliciesSeaweedSeaweedfsCom                    = "bucketlifecyclepolicies.seaweed.seaweedfs.com"
	CRDBucketnotificationsSeaweedSeaweedfsCom                        = "bucketnotifications.seaweed.seaweedfs.com"
	CRDBucketsSeaweedSeaweedfsCom                                    = "buckets.seaweed.seaweedfs.com"
	CRDResourcereferencegrantsSeaweedSeaweedfsCom                    = "resourcereferencegrants.seaweed.seaweedfs.com"
	CRDS3credentialsSeaweedSeaweedfsCom                              = "s3credentials.seaweed.seaweedfs.com"
	CRDS3identitiesSeaweedSeaweedfsCom                               = "s3identities.seaweed.seaweedfs.com"
	CRDS3oidcprovidersSeaweedSeaweedfsCom                            = "s3oidcproviders.seaweed.seaweedfs.com"
	CRDS3policiesSeaweedSeaweedfsCom                                 = "s3policies.seaweed.seaweedfs.com"
	CRDS3policybindingsSeaweedSeaweedfsCom                           = "s3policybindings.seaweed.seaweedfs.com"
	CRDSeaweedbackupsSeaweedSeaweedfsCom                             = "seaweedbackups.seaweed.seaweedfs.com"
	CRDSeaweedcsidriversSeaweedSeaweedfsCom                          = "seaweedcsidrivers.seaweed.seaweedfs.com"
	CRDSeaweedrestoresSeaweedSeaweedfsCom                            = "seaweedrestores.seaweed.seaweedfs.com"
	CRDSeaweedsSeaweedSeaweedfsCom                                   = "seaweeds.seaweed.seaweedfs.com"
	DeploymentNamespaceSeaweedfsOperator                             = "seaweedfs-operator"
	ServiceAccountNamespaceSeaweedfsOperator                         = "seaweedfs-operator"
	ClusterRoleSeaweedfsOperatorManagerRole                          = "seaweedfs-operator-manager-role"
	ClusterRoleSeaweedEditorRole                                     = "seaweed-editor-role"
	ClusterRoleSeaweedViewerRole                                     = "seaweed-viewer-role"
	ClusterRoleBindingSeaweedfsOperatorManagerRolebinding            = "seaweedfs-operator-manager-rolebinding"
	RoleNamespaceSeaweedfsOperatorLeaderElectionRole                 = "seaweedfs-operator-leader-election-role"
	RoleBindingNamespaceSeaweedfsOperatorLeaderElectionRolebinding   = "seaweedfs-operator-leader-election-rolebinding"
	SeaweedNamespaceSeaweedfs                                        = "seaweedfs"
	ServiceNamespaceSeaweedfsOperator                                = "seaweedfs-operator"
	ServiceNamespaceSeaweedfsOperatorWebhook                         = "seaweedfs-operator-webhook"
	MutatingWebhookSeaweedfsOperatorMutatingWebhookConfiguration     = "seaweedfs-operator-mutating-webhook-configuration"
	ValidatingWebhookSeaweedfsOperatorValidatingWebhookConfiguration = "seaweedfs-operator-validating-webhook-configuration"
)
