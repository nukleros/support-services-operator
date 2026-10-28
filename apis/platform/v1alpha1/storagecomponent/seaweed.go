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

package storagecomponent

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	platformv1alpha1 "github.com/nukleros/support-services-operator/apis/platform/v1alpha1"
	"github.com/nukleros/support-services-operator/apis/platform/v1alpha1/storagecomponent/mutate"
	setupv1alpha1 "github.com/nukleros/support-services-operator/apis/setup/v1alpha1"
)

// +kubebuilder:rbac:groups=seaweed.seaweedfs.com,resources=seaweeds,verbs=get;list;watch;create;update;patch;delete

// CreateSeaweedNamespaceSeaweedfs creates the Seaweed resource with name seaweedfs.
func CreateSeaweedNamespaceSeaweedfs(
	parent *platformv1alpha1.StorageComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "seaweed.seaweedfs.com/v1",
			"kind":       "Seaweed",
			"metadata": map[string]interface{}{
				"name":      "seaweedfs",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
			"spec": map[string]interface{}{
				// controlled by field: seaweed.image
				// controlled by field: seaweed.version
				//  Image repo and name to use for the SeaweedFS master, volume and filer servers.
				//  Version of SeaweedFS to use.
				"image":                 "" + parent.Spec.Seaweed.Image + ":" + parent.Spec.Seaweed.Version + "",
				"volumeServerDiskCount": 1,
				"master": map[string]interface{}{
					// controlled by field: seaweed.master.replicas
					//  Number of SeaweedFS master servers.  Use an odd number (3 for HA) so the masters keep a quorum.
					"replicas":          parent.Spec.Seaweed.Master.Replicas,
					"volumeSizeLimitMB": 1024,
					// controlled by field: seaweed.master.defaultReplication
					//  Default replication for new volumes, as SeaweedFS "xyz" placement (data centers, racks,
					//  servers).  "000" keeps one copy; "001" keeps a second copy on another volume server.
					"defaultReplication": parent.Spec.Seaweed.Master.DefaultReplication,
				},
				"volume": map[string]interface{}{
					// controlled by field: seaweed.volume.replicas
					//  Number of SeaweedFS volume servers.
					"replicas": parent.Spec.Seaweed.Volume.Replicas,
					"requests": map[string]interface{}{
						// controlled by field: seaweed.volume.storage
						//  Size of the persistent volume for each SeaweedFS volume server.
						"storage": parent.Spec.Seaweed.Volume.Storage,
					},
				},
				"filer": map[string]interface{}{
					// controlled by field: seaweed.filer.replicas
					//  Number of SeaweedFS filer servers.  The default leveldb2 filer store is local to each
					//  filer, so keep this at 1 unless the filer config points at a shared store.
					"replicas": parent.Spec.Seaweed.Filer.Replicas,
					"persistence": map[string]interface{}{
						"enabled": true,
					},
					"config": `[leveldb2]
enabled = true
dir = "/data/filerldb2"
`,
					"s3": map[string]interface{}{
						// controlled by field: seaweed.s3.enabled
						//  Run the S3 API gateway on the filer servers.
						"enabled": parent.Spec.Seaweed.S3.Enabled,
					},
				},
			},
		},
	}

	return mutate.MutateSeaweedNamespaceSeaweedfs(resourceObj, parent, collection, reconciler, req)
}
