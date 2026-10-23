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
	NamespaceNamespace                                  = "parent.Spec.Namespace"
	ConfigMapNamespacePostgresOperator                  = "postgres-operator"
	DeploymentNamespacePostgresOperator                 = "postgres-operator"
	ServiceAccountNamespacePostgresOperator             = "postgres-operator"
	ClusterRolePostgresOperator                         = "postgres-operator"
	ClusterRoleBindingPostgresOperator                  = "postgres-operator"
	ClusterRolePostgresPod                              = "postgres-pod"
	ServiceNamespacePostgresOperator                    = "postgres-operator"
	ConfigMapNamespaceCnpgDefaultMonitoring             = "cnpg-default-monitoring"
	CRDBackupsPostgresqlCnpgIo                          = "backups.postgresql.cnpg.io"
	CRDClusterimagecatalogsPostgresqlCnpgIo             = "clusterimagecatalogs.postgresql.cnpg.io"
	CRDClustersPostgresqlCnpgIo                         = "clusters.postgresql.cnpg.io"
	CRDDatabaserolesPostgresqlCnpgIo                    = "databaseroles.postgresql.cnpg.io"
	CRDDatabasesPostgresqlCnpgIo                        = "databases.postgresql.cnpg.io"
	CRDFailoverquorumsPostgresqlCnpgIo                  = "failoverquorums.postgresql.cnpg.io"
	CRDImagecatalogsPostgresqlCnpgIo                    = "imagecatalogs.postgresql.cnpg.io"
	CRDPoolersPostgresqlCnpgIo                          = "poolers.postgresql.cnpg.io"
	CRDPublicationsPostgresqlCnpgIo                     = "publications.postgresql.cnpg.io"
	CRDScheduledbackupsPostgresqlCnpgIo                 = "scheduledbackups.postgresql.cnpg.io"
	CRDSubscriptionsPostgresqlCnpgIo                    = "subscriptions.postgresql.cnpg.io"
	DeploymentNamespaceCnpgControllerManager            = "cnpg-controller-manager"
	ServiceAccountNamespaceCnpgManager                  = "cnpg-manager"
	ClusterRoleCnpgDatabaseEditorRole                   = "cnpg-database-editor-role"
	ClusterRoleCnpgDatabaseViewerRole                   = "cnpg-database-viewer-role"
	ClusterRoleCnpgManager                              = "cnpg-manager"
	ClusterRoleCnpgPublicationEditorRole                = "cnpg-publication-editor-role"
	ClusterRoleCnpgPublicationViewerRole                = "cnpg-publication-viewer-role"
	ClusterRoleCnpgSubscriptionEditorRole               = "cnpg-subscription-editor-role"
	ClusterRoleCnpgSubscriptionViewerRole               = "cnpg-subscription-viewer-role"
	ClusterRoleBindingCnpgManagerRolebinding            = "cnpg-manager-rolebinding"
	ServiceNamespaceCnpgWebhookService                  = "cnpg-webhook-service"
	MutatingWebhookCnpgMutatingWebhookConfiguration     = "cnpg-mutating-webhook-configuration"
	ValidatingWebhookCnpgValidatingWebhookConfiguration = "cnpg-validating-webhook-configuration"
)
