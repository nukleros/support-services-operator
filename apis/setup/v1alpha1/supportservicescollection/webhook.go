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

package supportservicescollection

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	setupv1alpha1 "github.com/nukleros/support-services-operator/apis/setup/v1alpha1"
	"github.com/nukleros/support-services-operator/apis/setup/v1alpha1/supportservicescollection/mutate"
)

// +kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=validatingadmissionpolicies,verbs=get;list;watch;create;update;patch;delete

// CreateValidatingAdmissionPolicySafeUpgradesGatewayNetworkingK8sIo creates the ValidatingAdmissionPolicy resource with name safe-upgrades.gateway.networking.k8s.io.
func CreateValidatingAdmissionPolicySafeUpgradesGatewayNetworkingK8sIo(
	parent *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if parent.Spec.EnableGatewayAPI != true {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:field=enableGatewayAPI,value=true,include
			"apiVersion": "admissionregistration.k8s.io/v1",
			"kind":       "ValidatingAdmissionPolicy",
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					"gateway.networking.k8s.io/bundle-version": "v1.7.0",
					"gateway.networking.k8s.io/channel":        "standard",
				},
				"name": "safe-upgrades.gateway.networking.k8s.io",
				"labels": map[string]interface{}{
					"platform.nukleros.io/category": "ingress",
					"platform.nukleros.io/project":  "gateway-api",
				},
			},
			"spec": map[string]interface{}{
				"failurePolicy": "Fail",
				"matchConstraints": map[string]interface{}{
					"resourceRules": []interface{}{
						map[string]interface{}{
							"apiGroups": []interface{}{
								"apiextensions.k8s.io",
							},
							"apiVersions": []interface{}{
								"v1",
							},
							"operations": []interface{}{
								"CREATE",
								"UPDATE",
							},
							"resources": []interface{}{
								"*",
							},
						},
					},
				},
				"validations": []interface{}{
					map[string]interface{}{
						"expression": "object.spec.group != 'gateway.networking.k8s.io' || oldObject == null || ( has(object.metadata.annotations) && object.metadata.annotations.exists(k, k == 'gateway.networking.k8s.io/channel') && object.metadata.annotations['gateway.networking.k8s.io/channel'] == 'standard' ) || ( oldObject != null && has(oldObject.metadata.annotations) && oldObject.metadata.annotations.exists(k, k == 'gateway.networking.k8s.io/channel') && oldObject.metadata.annotations['gateway.networking.k8s.io/channel'] == 'experimental' )",
						"message":    "Installing experimental CRDs on top of standard channel CRDs is prohibited by default. Uninstall ValidatingAdmissionPolicy safe-upgrades.gateway.networking.k8s.io to install experimental CRDs on top of standard channel CRDs.",
						"reason":     "Invalid",
					},
					map[string]interface{}{
						"expression": `object.spec.group != 'gateway.networking.k8s.io' || (has(object.metadata.annotations) && object.metadata.annotations.exists(k, k == 'gateway.networking.k8s.io/bundle-version') && !matches(object.metadata.annotations['gateway.networking.k8s.io/bundle-version'], 'v1.[0-6].\\d+') && !matches(object.metadata.annotations['gateway.networking.k8s.io/bundle-version'], 'v0'))`,
						"message":    "Installing CRDs with version before v1.7.0 is prohibited by default. Uninstall ValidatingAdmissionPolicy safe-upgrades.gateway.networking.k8s.io to install older versions.",
						"reason":     "Invalid",
					},
				},
			},
		},
	}

	return mutate.MutateValidatingAdmissionPolicySafeUpgradesGatewayNetworkingK8sIo(resourceObj, parent, reconciler, req)
}

// +kubebuilder:rbac:groups=admissionregistration.k8s.io,resources=validatingadmissionpolicybindings,verbs=get;list;watch;create;update;patch;delete

// CreateValidatingAdmissionPolicyBindingSafeUpgradesGatewayNetworkingK8sIo creates the ValidatingAdmissionPolicyBinding resource with name safe-upgrades.gateway.networking.k8s.io.
func CreateValidatingAdmissionPolicyBindingSafeUpgradesGatewayNetworkingK8sIo(
	parent *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if parent.Spec.EnableGatewayAPI != true {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:field=enableGatewayAPI,value=true,include
			"apiVersion": "admissionregistration.k8s.io/v1",
			"kind":       "ValidatingAdmissionPolicyBinding",
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					"gateway.networking.k8s.io/bundle-version": "v1.7.0",
					"gateway.networking.k8s.io/channel":        "standard",
				},
				"name": "safe-upgrades.gateway.networking.k8s.io",
				"labels": map[string]interface{}{
					"platform.nukleros.io/category": "ingress",
					"platform.nukleros.io/project":  "gateway-api",
				},
			},
			"spec": map[string]interface{}{
				"policyName": "safe-upgrades.gateway.networking.k8s.io",
				"validationActions": []interface{}{
					"Deny",
				},
				"matchResources": map[string]interface{}{
					"resourceRules": []interface{}{
						map[string]interface{}{
							"apiGroups": []interface{}{
								"apiextensions.k8s.io",
							},
							"apiVersions": []interface{}{
								"v1",
							},
							"resources": []interface{}{
								"customresourcedefinitions",
							},
							"operations": []interface{}{
								"CREATE",
								"UPDATE",
							},
						},
					},
				},
			},
		},
	}

	return mutate.MutateValidatingAdmissionPolicyBindingSafeUpgradesGatewayNetworkingK8sIo(resourceObj, parent, reconciler, req)
}
