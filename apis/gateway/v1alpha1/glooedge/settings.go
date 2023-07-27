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

package glooedge

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	gatewayv1alpha1 "github.com/nukleros/support-services-operator/apis/gateway/v1alpha1"
	"github.com/nukleros/support-services-operator/apis/gateway/v1alpha1/glooedge/mutate"
	orchestrationv1alpha1 "github.com/nukleros/support-services-operator/apis/orchestration/v1alpha1"
)

// +kubebuilder:rbac:groups=gloo.solo.io,resources=settings,verbs=get;list;watch;create;update;patch;delete

// CreateSettingsNamespaceDefault creates the Settings resource with name default.
func CreateSettingsNamespaceDefault(
	parent *gatewayv1alpha1.GlooEdge,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "gloo.solo.io/v1",
			"kind":       "Settings",
			"metadata": map[string]interface{}{
				"labels": map[string]interface{}{
					"app":  "gloo",
					"gloo": "settings",
				},
				"name":      "default",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
			"spec": map[string]interface{}{
				"gloo": map[string]interface{}{
					"xdsBindAddr":        "0.0.0.0:9977",
					"restXdsBindAddr":    "0.0.0.0:9976",
					"proxyDebugBindAddr": "0.0.0.0:9966",
					"enableRestEds":      false,
					"invalidConfigPolicy": map[string]interface{}{
						"invalidRouteResponseBody": "Gloo Gateway has invalid configuration. Administrators should run `glooctl check` to find and fix config errors.",
						"invalidRouteResponseCode": 404,
						"replaceInvalidRoutes":     false,
					},
					"disableKubernetesDestinations": false,
					"disableProxyGarbageCollection": false,
				},
				"discoveryNamespace":       "nukleros-gateway-system",
				"kubernetesArtifactSource": map[string]interface{}{},
				"kubernetesConfigSource":   map[string]interface{}{},
				"kubernetesSecretSource":   map[string]interface{}{},
				"refreshRate":              "60s",
				"gateway": map[string]interface{}{
					"isolateVirtualHostsBySslConfig": false,
					"readGatewaysFromAllNamespaces":  false,
					"enableGatewayController":        true,
					"validation": map[string]interface{}{
						"proxyValidationServerAddr":       "gloo:9988",
						"alwaysAccept":                    true,
						"allowWarnings":                   true,
						"serverEnabled":                   true,
						"disableTransformationValidation": false,
						"warnRouteShortCircuiting":        false,
					},
				},
				"discovery": map[string]interface{}{
					"fdsMode": "WHITELIST",
				},
			},
		},
	}

	return mutate.MutateSettingsNamespaceDefault(resourceObj, parent, collection, reconciler, req)
}
