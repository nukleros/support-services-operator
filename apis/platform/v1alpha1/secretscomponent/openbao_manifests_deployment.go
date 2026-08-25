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

package secretscomponent

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	platformv1alpha1 "github.com/nukleros/support-services-operator/apis/platform/v1alpha1"
	"github.com/nukleros/support-services-operator/apis/platform/v1alpha1/secretscomponent/mutate"
	setupv1alpha1 "github.com/nukleros/support-services-operator/apis/setup/v1alpha1"
)

// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete

// CreateDeploymentNamespaceOpenbaoAgentInjector creates the Deployment resource with name openbao-agent-injector.
func CreateDeploymentNamespaceOpenbaoAgentInjector(
	parent *platformv1alpha1.SecretsComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if parent.Spec.Openbao.Injector.Include != true {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:field=openbao.injector.include,value=true,include
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
			"metadata": map[string]interface{}{
				"name":      "openbao-agent-injector",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "openbao-agent-injector",
					"app.kubernetes.io/instance":    "openbao",
					"component":                     "webhook",
					"platform.nukleros.io/category": "secrets",
					"platform.nukleros.io/project":  "openbao",
				},
			},
			"spec": map[string]interface{}{
				// controlled by field: openbao.injector.replicas
				//  Number of replicas to use for the OpenBao Agent Injector deployment.
				"replicas": parent.Spec.Openbao.Injector.Replicas,
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"app.kubernetes.io/name":     "openbao-agent-injector",
						"app.kubernetes.io/instance": "openbao",
						"component":                  "webhook",
					},
				},
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
						"labels": map[string]interface{}{
							"app.kubernetes.io/name":        "openbao-agent-injector",
							"app.kubernetes.io/instance":    "openbao",
							"component":                     "webhook",
							"platform.nukleros.io/category": "secrets",
							"platform.nukleros.io/project":  "openbao",
						},
					},
					"spec": map[string]interface{}{
						"affinity": map[string]interface{}{
							"podAntiAffinity": map[string]interface{}{
								"preferredDuringSchedulingIgnoredDuringExecution": []interface{}{
									map[string]interface{}{
										"podAffinityTerm": map[string]interface{}{
											"labelSelector": map[string]interface{}{
												"matchLabels": map[string]interface{}{
													"app.kubernetes.io/instance": "openbao",
													"app.kubernetes.io/name":     "openbao-agent-injector",
													"component":                  "webhook",
												},
											},
											"topologyKey": "kubernetes.io/hostname",
										},
										"weight": 100,
									},
								},
							},
						},
						"nodeSelector": map[string]interface{}{
							"kubernetes.io/os": "linux",
						},
						"serviceAccountName": "openbao-agent-injector",
						"securityContext": map[string]interface{}{
							"seccompProfile": map[string]interface{}{
								"type": "RuntimeDefault",
							},
							"runAsNonRoot": true,
							"runAsGroup":   1000,
							"runAsUser":    100,
							"fsGroup":      1000,
						},
						"hostNetwork": false,
						"containers": []interface{}{
							map[string]interface{}{
								"name": "sidecar-injector",
								"resources": map[string]interface{}{
									"limits": map[string]interface{}{
										"memory": "256Mi",
									},
									"requests": map[string]interface{}{
										"cpu":    "250m",
										"memory": "256Mi",
									},
								},
								// controlled by field: openbao.injector.image
								// controlled by field: openbao.injector.imageVersion
								//  Image repo and name to use for the OpenBao Agent Injector sidecar.
								//  Version of the OpenBao Agent Injector sidecar image to use.
								"image":           "" + parent.Spec.Openbao.Injector.Image + ":" + parent.Spec.Openbao.Injector.ImageVersion + "",
								"imagePullPolicy": "IfNotPresent",
								"securityContext": map[string]interface{}{
									"allowPrivilegeEscalation": false,
									"capabilities": map[string]interface{}{
										"drop": []interface{}{
											"ALL",
										},
									},
								},
								"env": []interface{}{
									map[string]interface{}{
										"name":  "AGENT_INJECT_LISTEN",
										"value": ":8080",
									},
									map[string]interface{}{
										"name":  "AGENT_INJECT_LOG_LEVEL",
										"value": "info",
									},
									map[string]interface{}{
										"name":  "AGENT_INJECT_VAULT_ADDR",
										"value": "http://openbao.$(NAMESPACE).svc:8200",
									},
									map[string]interface{}{
										"name":  "AGENT_INJECT_VAULT_AUTH_PATH",
										"value": "auth/kubernetes",
									},
									map[string]interface{}{
										"name": "AGENT_INJECT_VAULT_IMAGE",
										// controlled by field: openbao.image
										// controlled by field: openbao.version
										"value": "" + parent.Spec.Openbao.Image + ":" + parent.Spec.Openbao.Version + "",
									},
									map[string]interface{}{
										"name":  "AGENT_INJECT_TLS_AUTO",
										"value": "openbao-agent-injector-cfg",
									},
									map[string]interface{}{
										"name":  "AGENT_INJECT_TLS_AUTO_HOSTS",
										"value": "openbao-agent-injector-svc,openbao-agent-injector-svc.$(NAMESPACE),openbao-agent-injector-svc.$(NAMESPACE).svc",
									},
									map[string]interface{}{
										"name":  "AGENT_INJECT_LOG_FORMAT",
										"value": "json",
									},
									map[string]interface{}{
										"name":  "AGENT_INJECT_REVOKE_ON_SHUTDOWN",
										"value": "false",
									},
									map[string]interface{}{
										"name":  "AGENT_INJECT_TELEMETRY_PATH",
										"value": "/metrics",
									},
									map[string]interface{}{
										"name":  "AGENT_INJECT_USE_LEADER_ELECTOR",
										"value": "true",
									},
									map[string]interface{}{
										"name": "NAMESPACE",
										"valueFrom": map[string]interface{}{
											"fieldRef": map[string]interface{}{
												"fieldPath": "metadata.namespace",
											},
										},
									},
									map[string]interface{}{
										"name":  "AGENT_INJECT_CPU_REQUEST",
										"value": "250m",
									},
									map[string]interface{}{
										"name":  "AGENT_INJECT_CPU_LIMIT",
										"value": "",
									},
									map[string]interface{}{
										"name":  "AGENT_INJECT_MEM_REQUEST",
										"value": "64Mi",
									},
									map[string]interface{}{
										"name":  "AGENT_INJECT_MEM_LIMIT",
										"value": "128Mi",
									},
									map[string]interface{}{
										"name":  "AGENT_INJECT_DEFAULT_TEMPLATE",
										"value": "map",
									},
									map[string]interface{}{
										"name":  "AGENT_INJECT_TEMPLATE_CONFIG_EXIT_ON_RETRY_FAILURE",
										"value": "true",
									},
									map[string]interface{}{
										"name": "POD_NAME",
										"valueFrom": map[string]interface{}{
											"fieldRef": map[string]interface{}{
												"fieldPath": "metadata.name",
											},
										},
									},
								},
								"args": []interface{}{
									"agent-inject",
									"2>&1",
								},
								"livenessProbe": map[string]interface{}{
									"httpGet": map[string]interface{}{
										"path":   "/health/ready",
										"port":   8080,
										"scheme": "HTTPS",
									},
									"failureThreshold":    2,
									"initialDelaySeconds": 5,
									"periodSeconds":       2,
									"successThreshold":    1,
									"timeoutSeconds":      5,
								},
								"readinessProbe": map[string]interface{}{
									"httpGet": map[string]interface{}{
										"path":   "/health/ready",
										"port":   8080,
										"scheme": "HTTPS",
									},
									"failureThreshold":    2,
									"initialDelaySeconds": 5,
									"periodSeconds":       2,
									"successThreshold":    1,
									"timeoutSeconds":      5,
								},
								"startupProbe": map[string]interface{}{
									"httpGet": map[string]interface{}{
										"path":   "/health/ready",
										"port":   8080,
										"scheme": "HTTPS",
									},
									"failureThreshold":    12,
									"initialDelaySeconds": 5,
									"periodSeconds":       5,
									"successThreshold":    1,
									"timeoutSeconds":      5,
								},
							},
						},
					},
				},
			},
		},
	}

	return mutate.MutateDeploymentNamespaceOpenbaoAgentInjector(resourceObj, parent, collection, reconciler, req)
}
