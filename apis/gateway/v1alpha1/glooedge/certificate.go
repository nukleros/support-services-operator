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

package glooedge

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	gatewayv1alpha1 "github.com/nukleros/support-services-operator/apis/gateway/v1alpha1"
	"github.com/nukleros/support-services-operator/apis/gateway/v1alpha1/glooedge/mutate"
	orchestrationv1alpha1 "github.com/nukleros/support-services-operator/apis/orchestration/v1alpha1"
)

// +kubebuilder:rbac:groups=cert-manager.io,resources=certificates,verbs=get;list;watch;create;update;patch;delete

// CreateCertNamespaceCertificateAuthority creates the Certificate resource with name certificate-authority.
func CreateCertNamespaceCertificateAuthority(
	parent *gatewayv1alpha1.GlooEdge,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "cert-manager.io/v1",
			"kind":       "Certificate",
			"metadata": map[string]interface{}{
				"name":      "certificate-authority",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
			"spec": map[string]interface{}{
				"isCA":       true,
				"commonName": "gloo-edge",
				"secretName": "certificate-authority",
				"privateKey": map[string]interface{}{
					"algorithm": "ECDSA",
					"size":      256,
				},
				"issuerRef": map[string]interface{}{
					"name":  "self-signed",
					"kind":  "ClusterIssuer",
					"group": "cert-manager.io",
				},
			},
		},
	}

	return mutate.MutateCertNamespaceCertificateAuthority(resourceObj, parent, collection, reconciler, req)
}
