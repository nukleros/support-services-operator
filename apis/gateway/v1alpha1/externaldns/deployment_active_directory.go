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

// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete

// CreateDeploymentNamespaceExternalDnsActiveDirectory creates the Deployment resource with name external-dns-active-directory.
func CreateDeploymentNamespaceExternalDnsActiveDirectory(
	parent *gatewayv1alpha1.ExternalDNS,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if parent.Spec.Provider != "active-directory" {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:field=provider,value="active-directory",include
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
			"metadata": map[string]interface{}{
				"name": "external-dns-active-directory",
				"labels": map[string]interface{}{
					"app":                          "external-dns-active-directory",
					"app.kubernetes.io/name":       "external-dns-active-directory",
					"app.kubernetes.io/instance":   "external-dns",
					"platform.nukleros.io/group":   "ingress",
					"platform.nukleros.io/project": "external-dns",
				},
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
			"spec": map[string]interface{}{
				"strategy": map[string]interface{}{
					"type": "Recreate",
				},
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"app": "external-dns-active-directory",
					},
				},
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
						"labels": map[string]interface{}{
							"app":                          "external-dns-active-directory",
							"app.kubernetes.io/name":       "external-dns-active-directory",
							"app.kubernetes.io/instance":   "external-dns",
							"platform.nukleros.io/group":   "ingress",
							"platform.nukleros.io/project": "external-dns",
						},
					},
					"spec": map[string]interface{}{
						"serviceAccountName": "external-dns",
						"containers": []interface{}{
							map[string]interface{}{
								"name": "external-dns",
								// controlled by field: image
								// controlled by field: version
								//  Image repo and name to use for external-dns.
								//  Version of external-dns to use.
								"image": "" + parent.Spec.Image + ":" + parent.Spec.Version + "",
								"args": []interface{}{
									"--source=service",
									"--source=ingress",
									"--registry=txt",
								},
								"envFrom": []interface{}{
									map[string]interface{}{
										"secretRef": map[string]interface{}{
											"name": "external-dns-active-directory",
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
								"volumeMounts": []interface{}{
									map[string]interface{}{
										"name":      "external-dns-active-directory-kerberos",
										"subPath":   "krb5.conf",
										"mountPath": "/etc/krb5.conf",
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
						"volumes": []interface{}{
							map[string]interface{}{
								"name": "external-dns-active-directory-kerberos",
								"configMap": map[string]interface{}{
									"defaultMode": 420,
									"name":        "external-dns-active-directory-kerberos",
								},
							},
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
															"external-dns-active-directory",
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

	return mutate.MutateDeploymentNamespaceExternalDnsActiveDirectory(resourceObj, parent, collection, reconciler, req)
}
