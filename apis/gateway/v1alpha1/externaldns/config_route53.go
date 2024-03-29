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

package externaldns

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	gatewayv1alpha1 "github.com/nukleros/support-services-operator/apis/gateway/v1alpha1"
	"github.com/nukleros/support-services-operator/apis/gateway/v1alpha1/externaldns/mutate"
	orchestrationv1alpha1 "github.com/nukleros/support-services-operator/apis/orchestration/v1alpha1"
)

// +kubebuilder:rbac:groups=core,resources=secrets,verbs=get;list;watch;create;update;patch;delete

// CreateSecretNamespaceExternalDnsRoute53 creates the Secret resource with name external-dns-route53.
func CreateSecretNamespaceExternalDnsRoute53(
	parent *gatewayv1alpha1.ExternalDNS,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if parent.Spec.Provider != "route53" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:field=provider,value="route53",include
			"apiVersion": "v1",
			"kind":       "Secret",
			"metadata": map[string]interface{}{
				"name":      "external-dns-route53",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"platform.nukleros.io/group":   "ingress",
					"platform.nukleros.io/project": "external-dns",
				},
			},
			"stringData": map[string]interface{}{
				"EXTERNAL_DNS_TXT_OWNER_ID": "external-dns-",
				"EXTERNAL_DNS_TXT_PREFIX":   "external-dns-",
				"EXTERNAL_DNS_PROVIDER":     "aws",
				// controlled by field: zoneType
				//  +kubebuilder:validation:Enum=private;public
				//  Type of DNS hosted zone to manage.
				"EXTERNAL_DNS_AWS_ZONE_TYPE": parent.Spec.ZoneType,
				"EXTERNAL_DNS_DOMAIN_FILTER": parent.Spec.DomainName, //  controlled by field: domainName
				"EXTERNAL_DNS_POLICY":        "sync",
				"AWS_ACCESS_KEY_ID":          "",
				"AWS_SECRET_ACCESS_KEY":      "",
			},
		},
	}

	return mutate.MutateSecretNamespaceExternalDnsRoute53(resourceObj, parent, collection, reconciler, req)
}
