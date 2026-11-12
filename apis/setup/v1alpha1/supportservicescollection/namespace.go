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

// +kubebuilder:rbac:groups=core,resources=namespaces,verbs=get;list;watch;create;update;patch;delete

// CreateNamespaceParentName creates the Namespace resource with name parent.Name.
func CreateNamespaceParentName(
	parent *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// struct: .
			// controlled by field: tier
			// controlled by field: defaultIngressController
			// controlled by field: cloud
			// controlled by field: platform
			// controlled by field: enableGatewayAPI
			//  +kubebuilder:validation:Enum=development;staging;production
			//  The tier of cluster being used.  One of: development | staging | production.
			//  +kubebuilder:validation:Enum=kong;nginx
			//  The default ingress for setting TLS certs.  One of: kong | nginx.
			//  +kubebuilder:validation:Enum=local;google;aws
			//  Cloud that this operator is running on.  One of: local | google | aws.
			//  When set to google, platform must also be set explicitly (e.g. gke), as the
			//  default platform of kind is not allowed.
			//  +kubebuilder:validation:Enum=kind;gke;openshift;eks;aks
			//  Platform that this operator is running on.  One of: kind | gke | openshift | eks | aks.
			//  Platform is set independently of cloud and is not inferred from it.  GKE clusters
			//  must set platform to gke: the generic (kind) istio-cni daemonset installs CNI
			//  plugins to /opt/cni/bin, while the GKE kubelet reads them from /home/kubernetes/bin,
			//  so pod sandboxes fail to start.  The combination of cloud=google and platform=kind
			//  is rejected.
			//  Enable Kubernetes Gateway API support.  Currently only supported for service-mesh components.
			"apiVersion": "v1",
			"kind":       "Namespace",
			"metadata": map[string]interface{}{
				"name": parent.Name, //  controlled by field:
			},
		},
	}

	return mutate.MutateNamespaceParentName(resourceObj, parent, reconciler, req)
}
