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

package constants

// this package includes the constants which include the resource names.  it is a standalone
// package to prevent import cycle errors when attempting to reference the names from other
// packages (e.g. mutate).
const (
	SecretNamespaceExternalDnsActiveDirectory            = "external-dns-active-directory"
	ConfigMapNamespaceExternalDnsActiveDirectoryKerberos = "external-dns-active-directory-kerberos"
	SecretNamespaceExternalDnsGoogle                     = "external-dns-google"
	SecretNamespaceExternalDnsRoute53                    = "external-dns-route53"
	DeploymentNamespaceExternalDnsActiveDirectory        = "external-dns-active-directory"
	DeploymentNamespaceExternalDnsGoogle                 = "external-dns-google"
	DeploymentNamespaceExternalDnsRoute53                = "external-dns-route53"
	NamespaceNamespace                                   = "parent.Spec.Namespace"
	ServiceAccountNamespaceServiceAccountName            = "parent.Spec.ServiceAccountName"
	ClusterRoleNamespaceExternalDns                      = "external-dns"
	ClusterRoleBindingExternalDnsViewer                  = "external-dns-viewer"
)
