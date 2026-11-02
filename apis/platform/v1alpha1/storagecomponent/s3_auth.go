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

// +kubebuilder:rbac:groups=seaweed.seaweedfs.com,resources=s3identities,verbs=get;list;watch;create;update;patch;delete

// CreateS3IdentityNamespaceSeaweedfsAdmin creates the S3Identity resource with name seaweedfs-admin.
func CreateS3IdentityNamespaceSeaweedfsAdmin(
	parent *platformv1alpha1.StorageComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if parent.Spec.Seaweed.S3.Auth.Enabled != true {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:field=seaweed.s3.auth.enabled,value=true,include
			"apiVersion": "seaweed.seaweedfs.com/v1",
			"kind":       "S3Identity",
			"metadata": map[string]interface{}{
				// controlled by field: seaweed.s3.auth.enabled
				//  Require credentials on the S3 API.  Creates an admin identity with full access whose
				//  access key and secret key are stored in the Secret named by seaweed.s3.auth.secretName.
				//  When disabled, anything that can reach the filer on port 8333 can read and write
				//  every bucket.  Disabling it on a running cluster does not remove the identity: delete
				//  the seaweedfs-admin S3Identity, S3Credentials, S3Policy and S3PolicyBinding, then
				//  restart the filer pods.
				"name":      "seaweedfs-admin",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
			"spec": map[string]interface{}{
				"seaweedRef": map[string]interface{}{
					"name": "seaweedfs",
				},
			},
		},
	}

	return mutate.MutateS3IdentityNamespaceSeaweedfsAdmin(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=seaweed.seaweedfs.com,resources=s3credentials,verbs=get;list;watch;create;update;patch;delete

// CreateS3CredentialsNamespaceSeaweedfsAdmin creates the S3Credentials resource with name seaweedfs-admin.
func CreateS3CredentialsNamespaceSeaweedfsAdmin(
	parent *platformv1alpha1.StorageComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if parent.Spec.Seaweed.S3.Auth.Enabled != true {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:field=seaweed.s3.auth.enabled,value=true,include
			"apiVersion": "seaweed.seaweedfs.com/v1",
			"kind":       "S3Credentials",
			"metadata": map[string]interface{}{
				"name":      "seaweedfs-admin",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
			"spec": map[string]interface{}{
				"seaweedRef": map[string]interface{}{
					"name": "seaweedfs",
				},
				"identityRef": map[string]interface{}{
					"name": "seaweedfs-admin",
				},
				"secretRef": map[string]interface{}{
					// controlled by field: seaweed.s3.auth.secretName
					//  Name of the Secret, in the storage namespace, that holds the admin identity's S3 keys
					//  under accessKey and secretKey.  The keys are generated if the Secret does not exist
					//  or is empty; keys already in the Secret (for example synced from a secret store)
					//  are used as they are.
					"name": parent.Spec.Seaweed.S3.Auth.SecretName,
				},
			},
		},
	}

	return mutate.MutateS3CredentialsNamespaceSeaweedfsAdmin(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=seaweed.seaweedfs.com,resources=s3policies,verbs=get;list;watch;create;update;patch;delete

// CreateS3PolicyNamespaceSeaweedfsAdmin creates the S3Policy resource with name seaweedfs-admin.
func CreateS3PolicyNamespaceSeaweedfsAdmin(
	parent *platformv1alpha1.StorageComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if parent.Spec.Seaweed.S3.Auth.Enabled != true {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:field=seaweed.s3.auth.enabled,value=true,include
			"apiVersion": "seaweed.seaweedfs.com/v1",
			"kind":       "S3Policy",
			"metadata": map[string]interface{}{
				"name":      "seaweedfs-admin",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
			"spec": map[string]interface{}{
				"seaweedRef": map[string]interface{}{
					"name": "seaweedfs",
				},
				"statements": []interface{}{
					map[string]interface{}{
						"effect": "Allow",
						"actions": []interface{}{
							"*",
						},
						"resources": []interface{}{
							"*",
						},
					},
				},
			},
		},
	}

	return mutate.MutateS3PolicyNamespaceSeaweedfsAdmin(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=seaweed.seaweedfs.com,resources=s3policybindings,verbs=get;list;watch;create;update;patch;delete

// CreateS3PolicyBindingNamespaceSeaweedfsAdmin creates the S3PolicyBinding resource with name seaweedfs-admin.
func CreateS3PolicyBindingNamespaceSeaweedfsAdmin(
	parent *platformv1alpha1.StorageComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if parent.Spec.Seaweed.S3.Auth.Enabled != true {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:field=seaweed.s3.auth.enabled,value=true,include
			"apiVersion": "seaweed.seaweedfs.com/v1",
			"kind":       "S3PolicyBinding",
			"metadata": map[string]interface{}{
				"name":      "seaweedfs-admin",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
			"spec": map[string]interface{}{
				"seaweedRef": map[string]interface{}{
					"name": "seaweedfs",
				},
				"policyRef": map[string]interface{}{
					"name": "seaweedfs-admin",
				},
				"subjects": []interface{}{
					map[string]interface{}{
						"kind": "S3Identity",
						"name": "seaweedfs-admin",
					},
				},
			},
		},
	}

	return mutate.MutateS3PolicyBindingNamespaceSeaweedfsAdmin(resourceObj, parent, collection, reconciler, req)
}
