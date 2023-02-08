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

package certificatescomponent

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	platformv1alpha1 "github.com/nukleros/support-services-operator/apis/platform/v1alpha1"
	"github.com/nukleros/support-services-operator/apis/platform/v1alpha1/certificatescomponent/mutate"
	setupv1alpha1 "github.com/nukleros/support-services-operator/apis/setup/v1alpha1"
)

// +kubebuilder:rbac:groups=cert-manager.io,resources=clusterissuers,verbs=get;list;watch;create;update;patch;delete

// CreateClusterIssuerLetsencryptStaging creates the ClusterIssuer resource with name letsencrypt-staging.
func CreateClusterIssuerLetsencryptStaging(
	parent *platformv1alpha1.CertificatesComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if collection.Spec.Tier == "production" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:collectionField=tier,value="production",include=false
			"apiVersion": "cert-manager.io/v1",
			"kind":       "ClusterIssuer",
			"metadata": map[string]interface{}{
				"name": "letsencrypt-staging",
				"annotations": map[string]interface{}{
					// Certificate provider to use.  Use one of: letsencrypt-staging or letsencrypt-production.
					"cert-provider": "letsencrypt-staging",
				},
			},
			"spec": map[string]interface{}{
				"acme": map[string]interface{}{
					"server": "https://acme-staging-v02.api.letsencrypt.org/directory",
					// controlled by field: certManager.contactEmail
					//  Contact e-mail address for receiving updates about certificates from LetsEncrypt.
					"email": parent.Spec.CertManager.ContactEmail,
					"privateKeySecretRef": map[string]interface{}{
						"name": "letsencrypt-staging",
					},
					"solvers": []interface{}{
						map[string]interface{}{
							"http01": map[string]interface{}{
								"ingress": map[string]interface{}{
									"podTemplate": map[string]interface{}{
										"metadata": map[string]interface{}{
											"creationTimestamp": nil,
											"labels": map[string]interface{}{
												"app.kubernetes.io/name": "cluster-issuer",
											},
										},
										"spec": map[string]interface{}{},
									},
									"class": "nginx",
								},
							},
						},
					},
				},
			},
		},
	}

	return mutate.MutateClusterIssuerLetsencryptStaging(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=cert-manager.io,resources=clusterissuers,verbs=get;list;watch;create;update;patch;delete

// CreateClusterIssuerLetsencryptProduction creates the ClusterIssuer resource with name letsencrypt-production.
func CreateClusterIssuerLetsencryptProduction(
	parent *platformv1alpha1.CertificatesComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if collection.Spec.Tier != "production" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:collectionField=tier,value="production",include
			"apiVersion": "cert-manager.io/v1",
			"kind":       "ClusterIssuer",
			"metadata": map[string]interface{}{
				// This issuer has low thresholds for rate limits,
				// so only use once bugs have been worked out for ingress stanzas
				"name": "letsencrypt-production",
				"annotations": map[string]interface{}{
					// Certificate provider to use.  Use one of: letsencrypt-staging or letsencrypt-production.
					"cert-provider": "letsencrypt-production",
				},
			},
			"spec": map[string]interface{}{
				"acme": map[string]interface{}{
					"server": "https://acme-v02.api.letsencrypt.org/directory",
					// controlled by field: certManager.contactEmail
					//  Contact e-mail address for receiving updates about certificates from LetsEncrypt.
					"email": parent.Spec.CertManager.ContactEmail,
					"privateKeySecretRef": map[string]interface{}{
						"name": "letsencrypt-production",
					},
					"solvers": []interface{}{
						map[string]interface{}{
							"http01": map[string]interface{}{
								"ingress": map[string]interface{}{
									"podTemplate": map[string]interface{}{
										"metadata": map[string]interface{}{
											"creationTimestamp": nil,
											"labels": map[string]interface{}{
												"app.kubernetes.io/name": "cluster-issuer",
											},
										},
										"spec": map[string]interface{}{},
									},
									"class": "nginx",
								},
							},
						},
					},
				},
			},
		},
	}

	return mutate.MutateClusterIssuerLetsencryptProduction(resourceObj, parent, collection, reconciler, req)
}
