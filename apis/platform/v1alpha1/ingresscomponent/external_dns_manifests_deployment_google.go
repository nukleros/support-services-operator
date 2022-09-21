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

// +kubebuilder:rbac:groups=core,resources=deployments,verbs=get;list;watch;create;update;patch;delete

const DeploymentNamespaceExternalDnsGoogle = "external-dns-google"

// CreateDeploymentNamespaceExternalDnsGoogle creates the external-dns-google Deployment resource.
func CreateDeploymentNamespaceExternalDnsGoogle(
	parent *platformv1alpha1.IngressComponent,
	collection *setupv1alpha1.SupportServices,
) ([]client.Object, error) {
	if parent.Spec.ExternalDNSProvider != "google" {
		return []client.Object{}, nil
	}

	resourceObjs := []client.Object{}
	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:field=externalDNSProvider,value="google",include
			"apiVersion": "v1",
			"kind":       "Deployment",
			"metadata": map[string]interface{}{
				"name": "external-dns-google",
				"labels": map[string]interface{}{
					"app":                            "external-dns-google",
					"app.kubernetes.io/name":         "external-dns-google",
					"app.kubernetes.io/instance":     "external-dns",
					"platform.kubernetes.io/purpose": "dns-updates",
				},
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
			"spec": map[string]interface{}{
				"strategy": map[string]interface{}{
					"type": "Recreate",
				},
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"app": "external-dns-google",
					},
				},
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
						"labels": map[string]interface{}{
							"app":                            "external-dns-google",
							"app.kubernetes.io/name":         "external-dns-google",
							"app.kubernetes.io/instance":     "external-dns",
							"platform.kubernetes.io/purpose": "dns-updates",
						},
					},
					"spec": map[string]interface{}{
						"serviceAccountName": "external-dns",
						"containers": []interface{}{
							map[string]interface{}{
								"name": "external-dns",
								// controlled by field: externalDNS.image
								// controlled by field: externalDNS.version
								//  Image repo and name to use for external-dns.
								//  Version of external-dns to use.
								"image": "" + parent.Spec.ExternalDNS.Image + ":" + parent.Spec.ExternalDNS.Version + "",
								"args": []interface{}{
									"--source=service",
									"--source=ingress",
									"--registry=txt",
								},
								"envFrom": []interface{}{
									map[string]interface{}{
										"secretRef": map[string]interface{}{
											"name": "external-dns-google",
										},
									},
								},
								"imagePullPolicy": "IfNotPresent",
								"securityContext": map[string]interface{}{
									"allowPrivilegeEscalation": false,
									"readOnlyRootFilesystem":   true,
									"capabilities": map[string]interface{}{
										"drop": []interface{}{
											"ALL",
										},
									},
								},
								"resources": map[string]interface{}{
									"requests": map[string]interface{}{
										"cpu":    "25m",
										"memory": "32Mi",
									},
									"limits": map[string]interface{}{
										"cpu":    "50m",
										"memory": "64Mi",
									},
								},
							},
						},
						"securityContext": map[string]interface{}{
							"fsGroup":      1001,
							"runAsUser":    1001,
							"runAsGroup":   1001,
							"runAsNonRoot": true,
						},
						"nodeSelector": map[string]interface{}{
							"kubernetes.io/os": "linux",
						},
						"affinity": map[string]interface{}{
							"podAntiAffinity": map[string]interface{}{
								"preferredDuringSchedulingIgnoredDuringExecution": []interface{}{
									map[string]interface{}{
										"weight": 100,
										"podAffinityTerm": map[string]interface{}{
											"topologyKey": "kubernetes.io/hostname",
											"labelSelector": map[string]interface{}{
												"matchExpressions": []interface{}{
													map[string]interface{}{
														"key":      "app.kubernetes.io/name",
														"operator": "In",
														"values": []interface{}{
															"external-dns-google",
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
				"replicas": 1,
			},
		},
	}

	resourceObjs = append(resourceObjs, resourceObj)

	return resourceObjs, nil
}
