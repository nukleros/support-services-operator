/*
Copyright 2022.

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

package ingresscomponent

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	platformv1alpha1 "github.com/nukleros/support-services-operator/apis/platform/v1alpha1"
	setupv1alpha1 "github.com/nukleros/support-services-operator/apis/setup/v1alpha1"
)

// +kubebuilder:rbac:groups=core,resources=secrets,verbs=get;list;watch;create;update;patch;delete

const SecretNamespaceExternalDnsRoute53 = "external-dns-route53"

// CreateSecretNamespaceExternalDnsRoute53 creates the external-dns-route53 Secret resource.
func CreateSecretNamespaceExternalDnsRoute53(
	parent *platformv1alpha1.IngressComponent,
	collection *setupv1alpha1.SupportServices,
) ([]client.Object, error) {
	if parent.Spec.ExternalDNSProvider != "route53" {
		return []client.Object{}, nil
	}

	resourceObjs := []client.Object{}
	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:field=externalDNSProvider,value="route53",include
			"apiVersion": "v1",
			"kind":       "Secret",
			"metadata": map[string]interface{}{
				"name":      "external-dns-route53",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
			"stringData": map[string]interface{}{
				"EXTERNAL_DNS_TXT_OWNER_ID":     "external-dns-",
				"EXTERNAL_DNS_TXT_PREFIX":       "external-dns-",
				"EXTERNAL_DNS_PROVIDER":         "aws",
				"EXTERNAL_DNS_AWS_ZONE_TYPE":    "private",
				"EXTERNAL_DNS_AWS_PREFER_CNAME": "true",
				"EXTERNAL_DNS_DOMAIN_FILTER":    "mydomain.com",
				"EXTERNAL_DNS_POLICY":           "sync",
				"AWS_ACCESS_KEY_ID":             "",
				"AWS_SECRET_ACCESS_KEY":         "",
			},
		},
	}

	resourceObjs = append(resourceObjs, resourceObj)

	return resourceObjs, nil
}
