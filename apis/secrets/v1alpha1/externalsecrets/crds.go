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

package externalsecrets

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	orchestrationv1alpha1 "github.com/nukleros/support-services-operator/apis/orchestration/v1alpha1"
	secretsv1alpha1 "github.com/nukleros/support-services-operator/apis/secrets/v1alpha1"
	"github.com/nukleros/support-services-operator/apis/secrets/v1alpha1/externalsecrets/mutate"
)

// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;create;update;patch;delete

// CreateCRDAcraccesstokensGeneratorsExternalSecretsIo creates the CustomResourceDefinition resource with name acraccesstokens.generators.external-secrets.io.
func CreateCRDAcraccesstokensGeneratorsExternalSecretsIo(
	parent *secretsv1alpha1.ExternalSecrets,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// Source: external-secrets/templates/crds/acraccesstoken.yaml
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					"controller-gen.kubebuilder.io/version": "v0.13.0",
				},
				"name": "acraccesstokens.generators.external-secrets.io",
			},
			"spec": map[string]interface{}{
				"group": "generators.external-secrets.io",
				"names": map[string]interface{}{
					"categories": []interface{}{
						"acraccesstoken",
					},
					"kind":     "ACRAccessToken",
					"listKind": "ACRAccessTokenList",
					"plural":   "acraccesstokens",
					"shortNames": []interface{}{
						"acraccesstoken",
					},
					"singular": "acraccesstoken",
				},
				"scope": "Namespaced",
				"versions": []interface{}{
					map[string]interface{}{
						"name": "v1alpha1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"description": `ACRAccessToken returns a Azure Container Registry token that can be used for pushing/pulling images. Note: by default it will return an ACR Refresh Token with full access (depending on the identity). This can be scoped down to the repository level using .spec.scope. In case scope is defined it will return an ACR Access Token. 
 See docs: https://github.com/Azure/acr/blob/main/docs/AAD-OAuth.md`,
								"properties": map[string]interface{}{
									"apiVersion": map[string]interface{}{
										"description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
										"type":        "string",
									},
									"kind": map[string]interface{}{
										"description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
										"type":        "string",
									},
									"metadata": map[string]interface{}{
										"type": "object",
									},
									"spec": map[string]interface{}{
										"description": "ACRAccessTokenSpec defines how to generate the access token e.g. how to authenticate and which registry to use. see: https://github.com/Azure/acr/blob/main/docs/AAD-OAuth.md#overview",
										"properties": map[string]interface{}{
											"auth": map[string]interface{}{
												"properties": map[string]interface{}{
													"managedIdentity": map[string]interface{}{
														"description": "ManagedIdentity uses Azure Managed Identity to authenticate with Azure.",
														"properties": map[string]interface{}{
															"identityId": map[string]interface{}{
																"description": "If multiple Managed Identity is assigned to the pod, you can select the one to be used",
																"type":        "string",
															},
														},
														"type": "object",
													},
													"servicePrincipal": map[string]interface{}{
														"description": "ServicePrincipal uses Azure Service Principal credentials to authenticate with Azure.",
														"properties": map[string]interface{}{
															"secretRef": map[string]interface{}{
																"description": "Configuration used to authenticate with Azure using static credentials stored in a Kind=Secret.",
																"properties": map[string]interface{}{
																	"clientId": map[string]interface{}{
																		"description": "The Azure clientId of the service principle used for authentication.",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																				"type":        "string",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the Secret resource being referred to.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																	"clientSecret": map[string]interface{}{
																		"description": "The Azure ClientSecret of the service principle used for authentication.",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																				"type":        "string",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the Secret resource being referred to.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
														},
														"required": []interface{}{
															"secretRef",
														},
														"type": "object",
													},
													"workloadIdentity": map[string]interface{}{
														"description": "WorkloadIdentity uses Azure Workload Identity to authenticate with Azure.",
														"properties": map[string]interface{}{
															"serviceAccountRef": map[string]interface{}{
																"description": "ServiceAccountRef specified the service account that should be used when authenticating with WorkloadIdentity.",
																"properties": map[string]interface{}{
																	"audiences": map[string]interface{}{
																		"description": "Audience specifies the `aud` claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																		"items": map[string]interface{}{
																			"type": "string",
																		},
																		"type": "array",
																	},
																	"name": map[string]interface{}{
																		"description": "The name of the ServiceAccount resource being referred to.",
																		"type":        "string",
																	},
																	"namespace": map[string]interface{}{
																		"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																		"type":        "string",
																	},
																},
																"required": []interface{}{
																	"name",
																},
																"type": "object",
															},
														},
														"type": "object",
													},
												},
												"type": "object",
											},
											"environmentType": map[string]interface{}{
												"default":     "PublicCloud",
												"description": "EnvironmentType specifies the Azure cloud environment endpoints to use for connecting and authenticating with Azure. By default it points to the public cloud AAD endpoint. The following endpoints are available, also see here: https://github.com/Azure/go-autorest/blob/main/autorest/azure/environments.go#L152 PublicCloud, USGovernmentCloud, ChinaCloud, GermanCloud",
												"enum": []interface{}{
													"PublicCloud",
													"USGovernmentCloud",
													"ChinaCloud",
													"GermanCloud",
												},
												"type": "string",
											},
											"registry": map[string]interface{}{
												"description": "the domain name of the ACR registry e.g. foobarexample.azurecr.io",
												"type":        "string",
											},
											"scope": map[string]interface{}{
												"description": `Define the scope for the access token, e.g. pull/push access for a repository. if not provided it will return a refresh token that has full scope. Note: you need to pin it down to the repository level, there is no wildcard available. 
 examples: repository:my-repository:pull,push repository:my-repository:pull 
 see docs for details: https://docs.docker.com/registry/spec/auth/scope/`,
												"type": "string",
											},
											"tenantId": map[string]interface{}{
												"description": "TenantID configures the Azure Tenant to send requests to. Required for ServicePrincipal auth type.",
												"type":        "string",
											},
										},
										"required": []interface{}{
											"auth",
											"registry",
										},
										"type": "object",
									},
								},
								"type": "object",
							},
						},
						"served":  true,
						"storage": true,
						"subresources": map[string]interface{}{
							"status": map[string]interface{}{},
						},
					},
				},
				"conversion": map[string]interface{}{
					"strategy": "Webhook",
					"webhook": map[string]interface{}{
						"conversionReviewVersions": []interface{}{
							"v1",
						},
						"clientConfig": map[string]interface{}{
							"service": map[string]interface{}{
								"name":      "release-name-external-secrets-webhook",
								"namespace": "default",
								"path":      "/convert",
							},
						},
					},
				},
			},
		},
	}

	return mutate.MutateCRDAcraccesstokensGeneratorsExternalSecretsIo(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;create;update;patch;delete

// CreateCRDClusterexternalsecretsExternalSecretsIo creates the CustomResourceDefinition resource with name clusterexternalsecrets.external-secrets.io.
func CreateCRDClusterexternalsecretsExternalSecretsIo(
	parent *secretsv1alpha1.ExternalSecrets,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// Source: external-secrets/templates/crds/clusterexternalsecret.yaml
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					"controller-gen.kubebuilder.io/version": "v0.13.0",
				},
				"name": "clusterexternalsecrets.external-secrets.io",
			},
			"spec": map[string]interface{}{
				"group": "external-secrets.io",
				"names": map[string]interface{}{
					"categories": []interface{}{
						"externalsecrets",
					},
					"kind":     "ClusterExternalSecret",
					"listKind": "ClusterExternalSecretList",
					"plural":   "clusterexternalsecrets",
					"shortNames": []interface{}{
						"ces",
					},
					"singular": "clusterexternalsecret",
				},
				"scope": "Cluster",
				"versions": []interface{}{
					map[string]interface{}{
						"additionalPrinterColumns": []interface{}{
							map[string]interface{}{
								"jsonPath": ".spec.externalSecretSpec.secretStoreRef.name",
								"name":     "Store",
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".spec.refreshTime",
								"name":     "Refresh Interval",
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".status.conditions[?(@.type==\"Ready\")].status",
								"name":     "Ready",
								"type":     "string",
							},
						},
						"name": "v1beta1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"description": "ClusterExternalSecret is the Schema for the clusterexternalsecrets API.",
								"properties": map[string]interface{}{
									"apiVersion": map[string]interface{}{
										"description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
										"type":        "string",
									},
									"kind": map[string]interface{}{
										"description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
										"type":        "string",
									},
									"metadata": map[string]interface{}{
										"type": "object",
									},
									"spec": map[string]interface{}{
										"description": "ClusterExternalSecretSpec defines the desired state of ClusterExternalSecret.",
										"properties": map[string]interface{}{
											"externalSecretMetadata": map[string]interface{}{
												"description": "The metadata of the external secrets to be created",
												"properties": map[string]interface{}{
													"annotations": map[string]interface{}{
														"additionalProperties": map[string]interface{}{
															"type": "string",
														},
														"type": "object",
													},
													"labels": map[string]interface{}{
														"additionalProperties": map[string]interface{}{
															"type": "string",
														},
														"type": "object",
													},
												},
												"type": "object",
											},
											"externalSecretName": map[string]interface{}{
												"description": "The name of the external secrets to be created defaults to the name of the ClusterExternalSecret",
												"type":        "string",
											},
											"externalSecretSpec": map[string]interface{}{
												"description": "The spec for the ExternalSecrets to be created",
												"properties": map[string]interface{}{
													"data": map[string]interface{}{
														"description": "Data defines the connection between the Kubernetes Secret keys and the Provider data",
														"items": map[string]interface{}{
															"description": "ExternalSecretData defines the connection between the Kubernetes Secret key (spec.data.<key>) and the Provider data.",
															"properties": map[string]interface{}{
																"remoteRef": map[string]interface{}{
																	"description": "RemoteRef points to the remote secret and defines which secret (version/property/..) to fetch.",
																	"properties": map[string]interface{}{
																		"conversionStrategy": map[string]interface{}{
																			"default":     "Default",
																			"description": "Used to define a conversion Strategy",
																			"enum": []interface{}{
																				"Default",
																				"Unicode",
																			},
																			"type": "string",
																		},
																		"decodingStrategy": map[string]interface{}{
																			"default":     "None",
																			"description": "Used to define a decoding Strategy",
																			"enum": []interface{}{
																				"Auto",
																				"Base64",
																				"Base64URL",
																				"None",
																			},
																			"type": "string",
																		},
																		"key": map[string]interface{}{
																			"description": "Key is the key used in the Provider, mandatory",
																			"type":        "string",
																		},
																		"metadataPolicy": map[string]interface{}{
																			"default":     "None",
																			"description": "Policy for fetching tags/labels from provider secrets, possible options are Fetch, None. Defaults to None",
																			"enum": []interface{}{
																				"None",
																				"Fetch",
																			},
																			"type": "string",
																		},
																		"property": map[string]interface{}{
																			"description": "Used to select a specific property of the Provider value (if a map), if supported",
																			"type":        "string",
																		},
																		"version": map[string]interface{}{
																			"description": "Used to select a specific version of the Provider value, if supported",
																			"type":        "string",
																		},
																	},
																	"required": []interface{}{
																		"key",
																	},
																	"type": "object",
																},
																"secretKey": map[string]interface{}{
																	"description": "SecretKey defines the key in which the controller stores the value. This is the key in the Kind=Secret",
																	"type":        "string",
																},
																"sourceRef": map[string]interface{}{
																	"description":   "SourceRef allows you to override the source from which the value will pulled from.",
																	"maxProperties": 1,
																	"properties": map[string]interface{}{
																		"generatorRef": map[string]interface{}{
																			"description": `GeneratorRef points to a generator custom resource. 
 Deprecated: The generatorRef is not implemented in .data[]. this will be removed with v1.`,
																			"properties": map[string]interface{}{
																				"apiVersion": map[string]interface{}{
																					"default":     "generators.external-secrets.io/v1alpha1",
																					"description": "Specify the apiVersion of the generator resource",
																					"type":        "string",
																				},
																				"kind": map[string]interface{}{
																					"description": "Specify the Kind of the resource, e.g. Password, ACRAccessToken etc.",
																					"type":        "string",
																				},
																				"name": map[string]interface{}{
																					"description": "Specify the name of the generator resource",
																					"type":        "string",
																				},
																			},
																			"required": []interface{}{
																				"kind",
																				"name",
																			},
																			"type": "object",
																		},
																		"storeRef": map[string]interface{}{
																			"description": "SecretStoreRef defines which SecretStore to fetch the ExternalSecret data.",
																			"properties": map[string]interface{}{
																				"kind": map[string]interface{}{
																					"description": "Kind of the SecretStore resource (SecretStore or ClusterSecretStore) Defaults to `SecretStore`",
																					"type":        "string",
																				},
																				"name": map[string]interface{}{
																					"description": "Name of the SecretStore resource",
																					"type":        "string",
																				},
																			},
																			"required": []interface{}{
																				"name",
																			},
																			"type": "object",
																		},
																	},
																	"type": "object",
																},
															},
															"required": []interface{}{
																"remoteRef",
																"secretKey",
															},
															"type": "object",
														},
														"type": "array",
													},
													"dataFrom": map[string]interface{}{
														"description": "DataFrom is used to fetch all properties from a specific Provider data If multiple entries are specified, the Secret keys are merged in the specified order",
														"items": map[string]interface{}{
															"properties": map[string]interface{}{
																"extract": map[string]interface{}{
																	"description": "Used to extract multiple key/value pairs from one secret Note: Extract does not support sourceRef.Generator or sourceRef.GeneratorRef.",
																	"properties": map[string]interface{}{
																		"conversionStrategy": map[string]interface{}{
																			"default":     "Default",
																			"description": "Used to define a conversion Strategy",
																			"enum": []interface{}{
																				"Default",
																				"Unicode",
																			},
																			"type": "string",
																		},
																		"decodingStrategy": map[string]interface{}{
																			"default":     "None",
																			"description": "Used to define a decoding Strategy",
																			"enum": []interface{}{
																				"Auto",
																				"Base64",
																				"Base64URL",
																				"None",
																			},
																			"type": "string",
																		},
																		"key": map[string]interface{}{
																			"description": "Key is the key used in the Provider, mandatory",
																			"type":        "string",
																		},
																		"metadataPolicy": map[string]interface{}{
																			"default":     "None",
																			"description": "Policy for fetching tags/labels from provider secrets, possible options are Fetch, None. Defaults to None",
																			"enum": []interface{}{
																				"None",
																				"Fetch",
																			},
																			"type": "string",
																		},
																		"property": map[string]interface{}{
																			"description": "Used to select a specific property of the Provider value (if a map), if supported",
																			"type":        "string",
																		},
																		"version": map[string]interface{}{
																			"description": "Used to select a specific version of the Provider value, if supported",
																			"type":        "string",
																		},
																	},
																	"required": []interface{}{
																		"key",
																	},
																	"type": "object",
																},
																"find": map[string]interface{}{
																	"description": "Used to find secrets based on tags or regular expressions Note: Find does not support sourceRef.Generator or sourceRef.GeneratorRef.",
																	"properties": map[string]interface{}{
																		"conversionStrategy": map[string]interface{}{
																			"default":     "Default",
																			"description": "Used to define a conversion Strategy",
																			"enum": []interface{}{
																				"Default",
																				"Unicode",
																			},
																			"type": "string",
																		},
																		"decodingStrategy": map[string]interface{}{
																			"default":     "None",
																			"description": "Used to define a decoding Strategy",
																			"enum": []interface{}{
																				"Auto",
																				"Base64",
																				"Base64URL",
																				"None",
																			},
																			"type": "string",
																		},
																		"name": map[string]interface{}{
																			"description": "Finds secrets based on the name.",
																			"properties": map[string]interface{}{
																				"regexp": map[string]interface{}{
																					"description": "Finds secrets base",
																					"type":        "string",
																				},
																			},
																			"type": "object",
																		},
																		"path": map[string]interface{}{
																			"description": "A root path to start the find operations.",
																			"type":        "string",
																		},
																		"tags": map[string]interface{}{
																			"additionalProperties": map[string]interface{}{
																				"type": "string",
																			},
																			"description": "Find secrets based on tags.",
																			"type":        "object",
																		},
																	},
																	"type": "object",
																},
																"rewrite": map[string]interface{}{
																	"description": "Used to rewrite secret Keys after getting them from the secret Provider Multiple Rewrite operations can be provided. They are applied in a layered order (first to last)",
																	"items": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"regexp": map[string]interface{}{
																				"description": "Used to rewrite with regular expressions. The resulting key will be the output of a regexp.ReplaceAll operation.",
																				"properties": map[string]interface{}{
																					"source": map[string]interface{}{
																						"description": "Used to define the regular expression of a re.Compiler.",
																						"type":        "string",
																					},
																					"target": map[string]interface{}{
																						"description": "Used to define the target pattern of a ReplaceAll operation.",
																						"type":        "string",
																					},
																				},
																				"required": []interface{}{
																					"source",
																					"target",
																				},
																				"type": "object",
																			},
																			"transform": map[string]interface{}{
																				"description": "Used to apply string transformation on the secrets. The resulting key will be the output of the template applied by the operation.",
																				"properties": map[string]interface{}{
																					"template": map[string]interface{}{
																						"description": "Used to define the template to apply on the secret name. `.value ` will specify the secret name in the template.",
																						"type":        "string",
																					},
																				},
																				"required": []interface{}{
																					"template",
																				},
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																	"type": "array",
																},
																"sourceRef": map[string]interface{}{
																	"description":   "SourceRef points to a store or generator which contains secret values ready to use. Use this in combination with Extract or Find pull values out of a specific SecretStore. When sourceRef points to a generator Extract or Find is not supported. The generator returns a static map of values",
																	"maxProperties": 1,
																	"properties": map[string]interface{}{
																		"generatorRef": map[string]interface{}{
																			"description": "GeneratorRef points to a generator custom resource.",
																			"properties": map[string]interface{}{
																				"apiVersion": map[string]interface{}{
																					"default":     "generators.external-secrets.io/v1alpha1",
																					"description": "Specify the apiVersion of the generator resource",
																					"type":        "string",
																				},
																				"kind": map[string]interface{}{
																					"description": "Specify the Kind of the resource, e.g. Password, ACRAccessToken etc.",
																					"type":        "string",
																				},
																				"name": map[string]interface{}{
																					"description": "Specify the name of the generator resource",
																					"type":        "string",
																				},
																			},
																			"required": []interface{}{
																				"kind",
																				"name",
																			},
																			"type": "object",
																		},
																		"storeRef": map[string]interface{}{
																			"description": "SecretStoreRef defines which SecretStore to fetch the ExternalSecret data.",
																			"properties": map[string]interface{}{
																				"kind": map[string]interface{}{
																					"description": "Kind of the SecretStore resource (SecretStore or ClusterSecretStore) Defaults to `SecretStore`",
																					"type":        "string",
																				},
																				"name": map[string]interface{}{
																					"description": "Name of the SecretStore resource",
																					"type":        "string",
																				},
																			},
																			"required": []interface{}{
																				"name",
																			},
																			"type": "object",
																		},
																	},
																	"type": "object",
																},
															},
															"type": "object",
														},
														"type": "array",
													},
													"refreshInterval": map[string]interface{}{
														"default":     "1h",
														"description": "RefreshInterval is the amount of time before the values are read again from the SecretStore provider Valid time units are \"ns\", \"us\" (or \"µs\"), \"ms\", \"s\", \"m\", \"h\" May be set to zero to fetch and create it once. Defaults to 1h.",
														"type":        "string",
													},
													"secretStoreRef": map[string]interface{}{
														"description": "SecretStoreRef defines which SecretStore to fetch the ExternalSecret data.",
														"properties": map[string]interface{}{
															"kind": map[string]interface{}{
																"description": "Kind of the SecretStore resource (SecretStore or ClusterSecretStore) Defaults to `SecretStore`",
																"type":        "string",
															},
															"name": map[string]interface{}{
																"description": "Name of the SecretStore resource",
																"type":        "string",
															},
														},
														"required": []interface{}{
															"name",
														},
														"type": "object",
													},
													"target": map[string]interface{}{
														"default": map[string]interface{}{
															"creationPolicy": "Owner",
															"deletionPolicy": "Retain",
														},
														"description": "ExternalSecretTarget defines the Kubernetes Secret to be created There can be only one target per ExternalSecret.",
														"properties": map[string]interface{}{
															"creationPolicy": map[string]interface{}{
																"default":     "Owner",
																"description": "CreationPolicy defines rules on how to create the resulting Secret Defaults to 'Owner'",
																"enum": []interface{}{
																	"Owner",
																	"Orphan",
																	"Merge",
																	"None",
																},
																"type": "string",
															},
															"deletionPolicy": map[string]interface{}{
																"default":     "Retain",
																"description": "DeletionPolicy defines rules on how to delete the resulting Secret Defaults to 'Retain'",
																"enum": []interface{}{
																	"Delete",
																	"Merge",
																	"Retain",
																},
																"type": "string",
															},
															"immutable": map[string]interface{}{
																"description": "Immutable defines if the final secret will be immutable",
																"type":        "boolean",
															},
															"name": map[string]interface{}{
																"description": "Name defines the name of the Secret resource to be managed This field is immutable Defaults to the .metadata.name of the ExternalSecret resource",
																"type":        "string",
															},
															"template": map[string]interface{}{
																"description": "Template defines a blueprint for the created Secret resource.",
																"properties": map[string]interface{}{
																	"data": map[string]interface{}{
																		"additionalProperties": map[string]interface{}{
																			"type": "string",
																		},
																		"type": "object",
																	},
																	"engineVersion": map[string]interface{}{
																		"default":     "v2",
																		"description": "EngineVersion specifies the template engine version that should be used to compile/execute the template specified in .data and .templateFrom[].",
																		"enum": []interface{}{
																			"v1",
																			"v2",
																		},
																		"type": "string",
																	},
																	"mergePolicy": map[string]interface{}{
																		"default": "Replace",
																		"enum": []interface{}{
																			"Replace",
																			"Merge",
																		},
																		"type": "string",
																	},
																	"metadata": map[string]interface{}{
																		"description": "ExternalSecretTemplateMetadata defines metadata fields for the Secret blueprint.",
																		"properties": map[string]interface{}{
																			"annotations": map[string]interface{}{
																				"additionalProperties": map[string]interface{}{
																					"type": "string",
																				},
																				"type": "object",
																			},
																			"labels": map[string]interface{}{
																				"additionalProperties": map[string]interface{}{
																					"type": "string",
																				},
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																	"templateFrom": map[string]interface{}{
																		"items": map[string]interface{}{
																			"properties": map[string]interface{}{
																				"configMap": map[string]interface{}{
																					"properties": map[string]interface{}{
																						"items": map[string]interface{}{
																							"items": map[string]interface{}{
																								"properties": map[string]interface{}{
																									"key": map[string]interface{}{
																										"type": "string",
																									},
																									"templateAs": map[string]interface{}{
																										"default": "Values",
																										"enum": []interface{}{
																											"Values",
																											"KeysAndValues",
																										},
																										"type": "string",
																									},
																								},
																								"required": []interface{}{
																									"key",
																								},
																								"type": "object",
																							},
																							"type": "array",
																						},
																						"name": map[string]interface{}{
																							"type": "string",
																						},
																					},
																					"required": []interface{}{
																						"items",
																						"name",
																					},
																					"type": "object",
																				},
																				"literal": map[string]interface{}{
																					"type": "string",
																				},
																				"secret": map[string]interface{}{
																					"properties": map[string]interface{}{
																						"items": map[string]interface{}{
																							"items": map[string]interface{}{
																								"properties": map[string]interface{}{
																									"key": map[string]interface{}{
																										"type": "string",
																									},
																									"templateAs": map[string]interface{}{
																										"default": "Values",
																										"enum": []interface{}{
																											"Values",
																											"KeysAndValues",
																										},
																										"type": "string",
																									},
																								},
																								"required": []interface{}{
																									"key",
																								},
																								"type": "object",
																							},
																							"type": "array",
																						},
																						"name": map[string]interface{}{
																							"type": "string",
																						},
																					},
																					"required": []interface{}{
																						"items",
																						"name",
																					},
																					"type": "object",
																				},
																				"target": map[string]interface{}{
																					"default": "Data",
																					"enum": []interface{}{
																						"Data",
																						"Annotations",
																						"Labels",
																					},
																					"type": "string",
																				},
																			},
																			"type": "object",
																		},
																		"type": "array",
																	},
																	"type": map[string]interface{}{
																		"type": "string",
																	},
																},
																"type": "object",
															},
														},
														"type": "object",
													},
												},
												"type": "object",
											},
											"namespaceSelector": map[string]interface{}{
												"description": "The labels to select by to find the Namespaces to create the ExternalSecrets in.",
												"properties": map[string]interface{}{
													"matchExpressions": map[string]interface{}{
														"description": "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
														"items": map[string]interface{}{
															"description": "A label selector requirement is a selector that contains values, a key, and an operator that relates the key and values.",
															"properties": map[string]interface{}{
																"key": map[string]interface{}{
																	"description": "key is the label key that the selector applies to.",
																	"type":        "string",
																},
																"operator": map[string]interface{}{
																	"description": "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																	"type":        "string",
																},
																"values": map[string]interface{}{
																	"description": "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																	"items": map[string]interface{}{
																		"type": "string",
																	},
																	"type": "array",
																},
															},
															"required": []interface{}{
																"key",
																"operator",
															},
															"type": "object",
														},
														"type": "array",
													},
													"matchLabels": map[string]interface{}{
														"additionalProperties": map[string]interface{}{
															"type": "string",
														},
														"description": "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is \"key\", the operator is \"In\", and the values array contains only \"value\". The requirements are ANDed.",
														"type":        "object",
													},
												},
												"type":                  "object",
												"x-kubernetes-map-type": "atomic",
											},
											"namespaces": map[string]interface{}{
												"description": "Choose namespaces by name. This field is ORed with anything that NamespaceSelector ends up choosing.",
												"items": map[string]interface{}{
													"type": "string",
												},
												"type": "array",
											},
											"refreshTime": map[string]interface{}{
												"description": "The time in which the controller should reconcile its objects and recheck namespaces for labels.",
												"type":        "string",
											},
										},
										"required": []interface{}{
											"externalSecretSpec",
										},
										"type": "object",
									},
									"status": map[string]interface{}{
										"description": "ClusterExternalSecretStatus defines the observed state of ClusterExternalSecret.",
										"properties": map[string]interface{}{
											"conditions": map[string]interface{}{
												"items": map[string]interface{}{
													"properties": map[string]interface{}{
														"message": map[string]interface{}{
															"type": "string",
														},
														"status": map[string]interface{}{
															"type": "string",
														},
														"type": map[string]interface{}{
															"type": "string",
														},
													},
													"required": []interface{}{
														"status",
														"type",
													},
													"type": "object",
												},
												"type": "array",
											},
											"externalSecretName": map[string]interface{}{
												"description": "ExternalSecretName is the name of the ExternalSecrets created by the ClusterExternalSecret",
												"type":        "string",
											},
											"failedNamespaces": map[string]interface{}{
												"description": "Failed namespaces are the namespaces that failed to apply an ExternalSecret",
												"items": map[string]interface{}{
													"description": "ClusterExternalSecretNamespaceFailure represents a failed namespace deployment and it's reason.",
													"properties": map[string]interface{}{
														"namespace": map[string]interface{}{
															"description": "Namespace is the namespace that failed when trying to apply an ExternalSecret",
															"type":        "string",
														},
														"reason": map[string]interface{}{
															"description": "Reason is why the ExternalSecret failed to apply to the namespace",
															"type":        "string",
														},
													},
													"required": []interface{}{
														"namespace",
													},
													"type": "object",
												},
												"type": "array",
											},
											"provisionedNamespaces": map[string]interface{}{
												"description": "ProvisionedNamespaces are the namespaces where the ClusterExternalSecret has secrets",
												"items": map[string]interface{}{
													"type": "string",
												},
												"type": "array",
											},
										},
										"type": "object",
									},
								},
								"type": "object",
							},
						},
						"served":  true,
						"storage": true,
						"subresources": map[string]interface{}{
							"status": map[string]interface{}{},
						},
					},
				},
				"conversion": map[string]interface{}{
					"strategy": "Webhook",
					"webhook": map[string]interface{}{
						"conversionReviewVersions": []interface{}{
							"v1",
						},
						"clientConfig": map[string]interface{}{
							"service": map[string]interface{}{
								"name":      "release-name-external-secrets-webhook",
								"namespace": "default",
								"path":      "/convert",
							},
						},
					},
				},
			},
		},
	}

	return mutate.MutateCRDClusterexternalsecretsExternalSecretsIo(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;create;update;patch;delete

// CreateCRDClustersecretstoresExternalSecretsIo creates the CustomResourceDefinition resource with name clustersecretstores.external-secrets.io.
func CreateCRDClustersecretstoresExternalSecretsIo(
	parent *secretsv1alpha1.ExternalSecrets,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// Source: external-secrets/templates/crds/clustersecretstore.yaml
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					"controller-gen.kubebuilder.io/version": "v0.13.0",
				},
				"name": "clustersecretstores.external-secrets.io",
			},
			"spec": map[string]interface{}{
				"group": "external-secrets.io",
				"names": map[string]interface{}{
					"categories": []interface{}{
						"externalsecrets",
					},
					"kind":     "ClusterSecretStore",
					"listKind": "ClusterSecretStoreList",
					"plural":   "clustersecretstores",
					"shortNames": []interface{}{
						"css",
					},
					"singular": "clustersecretstore",
				},
				"scope": "Cluster",
				"versions": []interface{}{
					map[string]interface{}{
						"additionalPrinterColumns": []interface{}{
							map[string]interface{}{
								"jsonPath": ".metadata.creationTimestamp",
								"name":     "AGE",
								"type":     "date",
							},
							map[string]interface{}{
								"jsonPath": ".status.conditions[?(@.type==\"Ready\")].reason",
								"name":     "Status",
								"type":     "string",
							},
						},
						"deprecated": true,
						"name":       "v1alpha1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"description": "ClusterSecretStore represents a secure external location for storing secrets, which can be referenced as part of `storeRef` fields.",
								"properties": map[string]interface{}{
									"apiVersion": map[string]interface{}{
										"description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
										"type":        "string",
									},
									"kind": map[string]interface{}{
										"description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
										"type":        "string",
									},
									"metadata": map[string]interface{}{
										"type": "object",
									},
									"spec": map[string]interface{}{
										"description": "SecretStoreSpec defines the desired state of SecretStore.",
										"properties": map[string]interface{}{
											"controller": map[string]interface{}{
												"description": "Used to select the correct ESO controller (think: ingress.ingressClassName) The ESO controller is instantiated with a specific controller name and filters ES based on this property",
												"type":        "string",
											},
											"provider": map[string]interface{}{
												"description":   "Used to configure the provider. Only one provider may be set",
												"maxProperties": 1,
												"minProperties": 1,
												"properties": map[string]interface{}{
													"akeyless": map[string]interface{}{
														"description": "Akeyless configures this store to sync secrets using Akeyless Vault provider",
														"properties": map[string]interface{}{
															"akeylessGWApiURL": map[string]interface{}{
																"description": "Akeyless GW API Url from which the secrets to be fetched from.",
																"type":        "string",
															},
															"authSecretRef": map[string]interface{}{
																"description": "Auth configures how the operator authenticates with Akeyless.",
																"properties": map[string]interface{}{
																	"kubernetesAuth": map[string]interface{}{
																		"description": "Kubernetes authenticates with Akeyless by passing the ServiceAccount token stored in the named Secret resource.",
																		"properties": map[string]interface{}{
																			"accessID": map[string]interface{}{
																				"description": "the Akeyless Kubernetes auth-method access-id",
																				"type":        "string",
																			},
																			"k8sConfName": map[string]interface{}{
																				"description": "Kubernetes-auth configuration name in Akeyless-Gateway",
																				"type":        "string",
																			},
																			"secretRef": map[string]interface{}{
																				"description": "Optional secret field containing a Kubernetes ServiceAccount JWT used for authenticating with Akeyless. If a name is specified without a key, `token` is the default. If one is not specified, the one bound to the controller will be used.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"serviceAccountRef": map[string]interface{}{
																				"description": "Optional service account field containing the name of a kubernetes ServiceAccount. If the service account is specified, the service account secret token JWT will be used for authenticating with Akeyless. If the service account selector is not supplied, the secretRef will be used instead.",
																				"properties": map[string]interface{}{
																					"audiences": map[string]interface{}{
																						"description": "Audience specifies the `aud` claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																						"items": map[string]interface{}{
																							"type": "string",
																						},
																						"type": "array",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the ServiceAccount resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"required": []interface{}{
																					"name",
																				},
																				"type": "object",
																			},
																		},
																		"required": []interface{}{
																			"accessID",
																			"k8sConfName",
																		},
																		"type": "object",
																	},
																	"secretRef": map[string]interface{}{
																		"description": "Reference to a Secret that contains the details to authenticate with Akeyless.",
																		"properties": map[string]interface{}{
																			"accessID": map[string]interface{}{
																				"description": "The SecretAccessID is used for authentication",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"accessType": map[string]interface{}{
																				"description": "A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"accessTypeParam": map[string]interface{}{
																				"description": "A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
															"caBundle": map[string]interface{}{
																"description": "PEM/base64 encoded CA bundle used to validate Akeyless Gateway certificate. Only used if the AkeylessGWApiURL URL is using HTTPS protocol. If not set the system root certificates are used to validate the TLS connection.",
																"format":      "byte",
																"type":        "string",
															},
															"caProvider": map[string]interface{}{
																"description": "The provider for the CA bundle to use to validate Akeyless Gateway certificate.",
																"properties": map[string]interface{}{
																	"key": map[string]interface{}{
																		"description": "The key the value inside of the provider type to use, only used with \"Secret\" type",
																		"type":        "string",
																	},
																	"name": map[string]interface{}{
																		"description": "The name of the object located at the provider type.",
																		"type":        "string",
																	},
																	"namespace": map[string]interface{}{
																		"description": "The namespace the Provider type is in.",
																		"type":        "string",
																	},
																	"type": map[string]interface{}{
																		"description": "The type of provider to use such as \"Secret\", or \"ConfigMap\".",
																		"enum": []interface{}{
																			"Secret",
																			"ConfigMap",
																		},
																		"type": "string",
																	},
																},
																"required": []interface{}{
																	"name",
																	"type",
																},
																"type": "object",
															},
														},
														"required": []interface{}{
															"akeylessGWApiURL",
															"authSecretRef",
														},
														"type": "object",
													},
													"alibaba": map[string]interface{}{
														"description": "Alibaba configures this store to sync secrets using Alibaba Cloud provider",
														"properties": map[string]interface{}{
															"auth": map[string]interface{}{
																"description": "AlibabaAuth contains a secretRef for credentials.",
																"properties": map[string]interface{}{
																	"rrsa": map[string]interface{}{
																		"description": "Authenticate against Alibaba using RRSA.",
																		"properties": map[string]interface{}{
																			"oidcProviderArn": map[string]interface{}{
																				"type": "string",
																			},
																			"oidcTokenFilePath": map[string]interface{}{
																				"type": "string",
																			},
																			"roleArn": map[string]interface{}{
																				"type": "string",
																			},
																			"sessionName": map[string]interface{}{
																				"type": "string",
																			},
																		},
																		"required": []interface{}{
																			"oidcProviderArn",
																			"oidcTokenFilePath",
																			"roleArn",
																			"sessionName",
																		},
																		"type": "object",
																	},
																	"secretRef": map[string]interface{}{
																		"description": "AlibabaAuthSecretRef holds secret references for Alibaba credentials.",
																		"properties": map[string]interface{}{
																			"accessKeyIDSecretRef": map[string]interface{}{
																				"description": "The AccessKeyID is used for authentication",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"accessKeySecretSecretRef": map[string]interface{}{
																				"description": "The AccessKeySecret is used for authentication",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"required": []interface{}{
																			"accessKeyIDSecretRef",
																			"accessKeySecretSecretRef",
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
															"regionID": map[string]interface{}{
																"description": "Alibaba Region to be used for the provider",
																"type":        "string",
															},
														},
														"required": []interface{}{
															"auth",
															"regionID",
														},
														"type": "object",
													},
													"aws": map[string]interface{}{
														"description": "AWS configures this store to sync secrets using AWS Secret Manager provider",
														"properties": map[string]interface{}{
															"auth": map[string]interface{}{
																"description": "Auth defines the information necessary to authenticate against AWS if not set aws sdk will infer credentials from your environment see: https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-credentials",
																"properties": map[string]interface{}{
																	"jwt": map[string]interface{}{
																		"description": "Authenticate against AWS using service account tokens.",
																		"properties": map[string]interface{}{
																			"serviceAccountRef": map[string]interface{}{
																				"description": "A reference to a ServiceAccount resource.",
																				"properties": map[string]interface{}{
																					"audiences": map[string]interface{}{
																						"description": "Audience specifies the `aud` claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																						"items": map[string]interface{}{
																							"type": "string",
																						},
																						"type": "array",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the ServiceAccount resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"required": []interface{}{
																					"name",
																				},
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																	"secretRef": map[string]interface{}{
																		"description": "AWSAuthSecretRef holds secret references for AWS credentials both AccessKeyID and SecretAccessKey must be defined in order to properly authenticate.",
																		"properties": map[string]interface{}{
																			"accessKeyIDSecretRef": map[string]interface{}{
																				"description": "The AccessKeyID is used for authentication",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"secretAccessKeySecretRef": map[string]interface{}{
																				"description": "The SecretAccessKey is used for authentication",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
															"region": map[string]interface{}{
																"description": "AWS Region to be used for the provider",
																"type":        "string",
															},
															"role": map[string]interface{}{
																"description": "Role is a Role ARN which the SecretManager provider will assume",
																"type":        "string",
															},
															"service": map[string]interface{}{
																"description": "Service defines which service should be used to fetch the secrets",
																"enum": []interface{}{
																	"SecretsManager",
																	"ParameterStore",
																},
																"type": "string",
															},
														},
														"required": []interface{}{
															"region",
															"service",
														},
														"type": "object",
													},
													"azurekv": map[string]interface{}{
														"description": "AzureKV configures this store to sync secrets using Azure Key Vault provider",
														"properties": map[string]interface{}{
															"authSecretRef": map[string]interface{}{
																"description": "Auth configures how the operator authenticates with Azure. Required for ServicePrincipal auth type.",
																"properties": map[string]interface{}{
																	"clientId": map[string]interface{}{
																		"description": "The Azure clientId of the service principle used for authentication.",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																				"type":        "string",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the Secret resource being referred to.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																	"clientSecret": map[string]interface{}{
																		"description": "The Azure ClientSecret of the service principle used for authentication.",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																				"type":        "string",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the Secret resource being referred to.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
															"authType": map[string]interface{}{
																"default":     "ServicePrincipal",
																"description": "Auth type defines how to authenticate to the keyvault service. Valid values are: - \"ServicePrincipal\" (default): Using a service principal (tenantId, clientId, clientSecret) - \"ManagedIdentity\": Using Managed Identity assigned to the pod (see aad-pod-identity)",
																"enum": []interface{}{
																	"ServicePrincipal",
																	"ManagedIdentity",
																	"WorkloadIdentity",
																},
																"type": "string",
															},
															"identityId": map[string]interface{}{
																"description": "If multiple Managed Identity is assigned to the pod, you can select the one to be used",
																"type":        "string",
															},
															"serviceAccountRef": map[string]interface{}{
																"description": "ServiceAccountRef specified the service account that should be used when authenticating with WorkloadIdentity.",
																"properties": map[string]interface{}{
																	"audiences": map[string]interface{}{
																		"description": "Audience specifies the `aud` claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																		"items": map[string]interface{}{
																			"type": "string",
																		},
																		"type": "array",
																	},
																	"name": map[string]interface{}{
																		"description": "The name of the ServiceAccount resource being referred to.",
																		"type":        "string",
																	},
																	"namespace": map[string]interface{}{
																		"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																		"type":        "string",
																	},
																},
																"required": []interface{}{
																	"name",
																},
																"type": "object",
															},
															"tenantId": map[string]interface{}{
																"description": "TenantID configures the Azure Tenant to send requests to. Required for ServicePrincipal auth type.",
																"type":        "string",
															},
															"vaultUrl": map[string]interface{}{
																"description": "Vault Url from which the secrets to be fetched from.",
																"type":        "string",
															},
														},
														"required": []interface{}{
															"vaultUrl",
														},
														"type": "object",
													},
													"fake": map[string]interface{}{
														"description": "Fake configures a store with static key/value pairs",
														"properties": map[string]interface{}{
															"data": map[string]interface{}{
																"items": map[string]interface{}{
																	"properties": map[string]interface{}{
																		"key": map[string]interface{}{
																			"type": "string",
																		},
																		"value": map[string]interface{}{
																			"type": "string",
																		},
																		"valueMap": map[string]interface{}{
																			"additionalProperties": map[string]interface{}{
																				"type": "string",
																			},
																			"type": "object",
																		},
																		"version": map[string]interface{}{
																			"type": "string",
																		},
																	},
																	"required": []interface{}{
																		"key",
																	},
																	"type": "object",
																},
																"type": "array",
															},
														},
														"required": []interface{}{
															"data",
														},
														"type": "object",
													},
													"gcpsm": map[string]interface{}{
														"description": "GCPSM configures this store to sync secrets using Google Cloud Platform Secret Manager provider",
														"properties": map[string]interface{}{
															"auth": map[string]interface{}{
																"description": "Auth defines the information necessary to authenticate against GCP",
																"properties": map[string]interface{}{
																	"secretRef": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"secretAccessKeySecretRef": map[string]interface{}{
																				"description": "The SecretAccessKey is used for authentication",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																	"workloadIdentity": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"clusterLocation": map[string]interface{}{
																				"type": "string",
																			},
																			"clusterName": map[string]interface{}{
																				"type": "string",
																			},
																			"clusterProjectID": map[string]interface{}{
																				"type": "string",
																			},
																			"serviceAccountRef": map[string]interface{}{
																				"description": "A reference to a ServiceAccount resource.",
																				"properties": map[string]interface{}{
																					"audiences": map[string]interface{}{
																						"description": "Audience specifies the `aud` claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																						"items": map[string]interface{}{
																							"type": "string",
																						},
																						"type": "array",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the ServiceAccount resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"required": []interface{}{
																					"name",
																				},
																				"type": "object",
																			},
																		},
																		"required": []interface{}{
																			"clusterLocation",
																			"clusterName",
																			"serviceAccountRef",
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
															"projectID": map[string]interface{}{
																"description": "ProjectID project where secret is located",
																"type":        "string",
															},
														},
														"type": "object",
													},
													"gitlab": map[string]interface{}{
														"description": "GitLab configures this store to sync secrets using GitLab Variables provider",
														"properties": map[string]interface{}{
															"auth": map[string]interface{}{
																"description": "Auth configures how secret-manager authenticates with a GitLab instance.",
																"properties": map[string]interface{}{
																	"SecretRef": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"accessToken": map[string]interface{}{
																				"description": "AccessToken is used for authentication.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																},
																"required": []interface{}{
																	"SecretRef",
																},
																"type": "object",
															},
															"projectID": map[string]interface{}{
																"description": "ProjectID specifies a project where secrets are located.",
																"type":        "string",
															},
															"url": map[string]interface{}{
																"description": "URL configures the GitLab instance URL. Defaults to https://gitlab.com/.",
																"type":        "string",
															},
														},
														"required": []interface{}{
															"auth",
														},
														"type": "object",
													},
													"ibm": map[string]interface{}{
														"description": "IBM configures this store to sync secrets using IBM Cloud provider",
														"properties": map[string]interface{}{
															"auth": map[string]interface{}{
																"description": "Auth configures how secret-manager authenticates with the IBM secrets manager.",
																"properties": map[string]interface{}{
																	"secretRef": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"secretApiKeySecretRef": map[string]interface{}{
																				"description": "The SecretAccessKey is used for authentication",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																},
																"required": []interface{}{
																	"secretRef",
																},
																"type": "object",
															},
															"serviceUrl": map[string]interface{}{
																"description": "ServiceURL is the Endpoint URL that is specific to the Secrets Manager service instance",
																"type":        "string",
															},
														},
														"required": []interface{}{
															"auth",
														},
														"type": "object",
													},
													"kubernetes": map[string]interface{}{
														"description": "Kubernetes configures this store to sync secrets using a Kubernetes cluster provider",
														"properties": map[string]interface{}{
															"auth": map[string]interface{}{
																"description":   "Auth configures how secret-manager authenticates with a Kubernetes instance.",
																"maxProperties": 1,
																"minProperties": 1,
																"properties": map[string]interface{}{
																	"cert": map[string]interface{}{
																		"description": "has both clientCert and clientKey as secretKeySelector",
																		"properties": map[string]interface{}{
																			"clientCert": map[string]interface{}{
																				"description": "A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"clientKey": map[string]interface{}{
																				"description": "A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																	"serviceAccount": map[string]interface{}{
																		"description": "points to a service account that should be used for authentication",
																		"properties": map[string]interface{}{
																			"serviceAccount": map[string]interface{}{
																				"description": "A reference to a ServiceAccount resource.",
																				"properties": map[string]interface{}{
																					"audiences": map[string]interface{}{
																						"description": "Audience specifies the `aud` claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																						"items": map[string]interface{}{
																							"type": "string",
																						},
																						"type": "array",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the ServiceAccount resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"required": []interface{}{
																					"name",
																				},
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																	"token": map[string]interface{}{
																		"description": "use static token to authenticate with",
																		"properties": map[string]interface{}{
																			"bearerToken": map[string]interface{}{
																				"description": "A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
															"remoteNamespace": map[string]interface{}{
																"default":     "default",
																"description": "Remote namespace to fetch the secrets from",
																"type":        "string",
															},
															"server": map[string]interface{}{
																"description": "configures the Kubernetes server Address.",
																"properties": map[string]interface{}{
																	"caBundle": map[string]interface{}{
																		"description": "CABundle is a base64-encoded CA certificate",
																		"format":      "byte",
																		"type":        "string",
																	},
																	"caProvider": map[string]interface{}{
																		"description": "see: https://external-secrets.io/v0.4.1/spec/#external-secrets.io/v1alpha1.CAProvider",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "The key the value inside of the provider type to use, only used with \"Secret\" type",
																				"type":        "string",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the object located at the provider type.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "The namespace the Provider type is in.",
																				"type":        "string",
																			},
																			"type": map[string]interface{}{
																				"description": "The type of provider to use such as \"Secret\", or \"ConfigMap\".",
																				"enum": []interface{}{
																					"Secret",
																					"ConfigMap",
																				},
																				"type": "string",
																			},
																		},
																		"required": []interface{}{
																			"name",
																			"type",
																		},
																		"type": "object",
																	},
																	"url": map[string]interface{}{
																		"default":     "kubernetes.default",
																		"description": "configures the Kubernetes server Address.",
																		"type":        "string",
																	},
																},
																"type": "object",
															},
														},
														"required": []interface{}{
															"auth",
														},
														"type": "object",
													},
													"oracle": map[string]interface{}{
														"description": "Oracle configures this store to sync secrets using Oracle Vault provider",
														"properties": map[string]interface{}{
															"auth": map[string]interface{}{
																"description": "Auth configures how secret-manager authenticates with the Oracle Vault. If empty, instance principal is used. Optionally, the authenticating principal type and/or user data may be supplied for the use of workload identity and user principal.",
																"properties": map[string]interface{}{
																	"secretRef": map[string]interface{}{
																		"description": "SecretRef to pass through sensitive information.",
																		"properties": map[string]interface{}{
																			"fingerprint": map[string]interface{}{
																				"description": "Fingerprint is the fingerprint of the API private key.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"privatekey": map[string]interface{}{
																				"description": "PrivateKey is the user's API Signing Key in PEM format, used for authentication.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"required": []interface{}{
																			"fingerprint",
																			"privatekey",
																		},
																		"type": "object",
																	},
																	"tenancy": map[string]interface{}{
																		"description": "Tenancy is the tenancy OCID where user is located.",
																		"type":        "string",
																	},
																	"user": map[string]interface{}{
																		"description": "User is an access OCID specific to the account.",
																		"type":        "string",
																	},
																},
																"required": []interface{}{
																	"secretRef",
																	"tenancy",
																	"user",
																},
																"type": "object",
															},
															"compartment": map[string]interface{}{
																"description": "Compartment is the vault compartment OCID. Required for PushSecret",
																"type":        "string",
															},
															"encryptionKey": map[string]interface{}{
																"description": "EncryptionKey is the OCID of the encryption key within the vault. Required for PushSecret",
																"type":        "string",
															},
															"principalType": map[string]interface{}{
																"description": "The type of principal to use for authentication. If left blank, the Auth struct will determine the principal type. This optional field must be specified if using workload identity.",
																"enum": []interface{}{
																	"",
																	"UserPrincipal",
																	"InstancePrincipal",
																	"Workload",
																},
																"type": "string",
															},
															"region": map[string]interface{}{
																"description": "Region is the region where vault is located.",
																"type":        "string",
															},
															"serviceAccountRef": map[string]interface{}{
																"description": "ServiceAccountRef specified the service account that should be used when authenticating with WorkloadIdentity.",
																"properties": map[string]interface{}{
																	"audiences": map[string]interface{}{
																		"description": "Audience specifies the `aud` claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																		"items": map[string]interface{}{
																			"type": "string",
																		},
																		"type": "array",
																	},
																	"name": map[string]interface{}{
																		"description": "The name of the ServiceAccount resource being referred to.",
																		"type":        "string",
																	},
																	"namespace": map[string]interface{}{
																		"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																		"type":        "string",
																	},
																},
																"required": []interface{}{
																	"name",
																},
																"type": "object",
															},
															"vault": map[string]interface{}{
																"description": "Vault is the vault's OCID of the specific vault where secret is located.",
																"type":        "string",
															},
														},
														"required": []interface{}{
															"region",
															"vault",
														},
														"type": "object",
													},
													"vault": map[string]interface{}{
														"description": "Vault configures this store to sync secrets using Hashi provider",
														"properties": map[string]interface{}{
															"auth": map[string]interface{}{
																"description": "Auth configures how secret-manager authenticates with the Vault server.",
																"properties": map[string]interface{}{
																	"appRole": map[string]interface{}{
																		"description": "AppRole authenticates with Vault using the App Role auth mechanism, with the role and secret stored in a Kubernetes Secret resource.",
																		"properties": map[string]interface{}{
																			"path": map[string]interface{}{
																				"default":     "approle",
																				"description": "Path where the App Role authentication backend is mounted in Vault, e.g: \"approle\"",
																				"type":        "string",
																			},
																			"roleId": map[string]interface{}{
																				"description": "RoleID configured in the App Role authentication backend when setting up the authentication backend in Vault.",
																				"type":        "string",
																			},
																			"secretRef": map[string]interface{}{
																				"description": "Reference to a key in a Secret that contains the App Role secret used to authenticate with Vault. The `key` field must be specified and denotes which entry within the Secret resource is used as the app role secret.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"required": []interface{}{
																			"path",
																			"roleId",
																			"secretRef",
																		},
																		"type": "object",
																	},
																	"cert": map[string]interface{}{
																		"description": "Cert authenticates with TLS Certificates by passing client certificate, private key and ca certificate Cert authentication method",
																		"properties": map[string]interface{}{
																			"clientCert": map[string]interface{}{
																				"description": "ClientCert is a certificate to authenticate using the Cert Vault authentication method",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"secretRef": map[string]interface{}{
																				"description": "SecretRef to a key in a Secret resource containing client private key to authenticate with Vault using the Cert authentication method",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																	"jwt": map[string]interface{}{
																		"description": "Jwt authenticates with Vault by passing role and JWT token using the JWT/OIDC authentication method",
																		"properties": map[string]interface{}{
																			"kubernetesServiceAccountToken": map[string]interface{}{
																				"description": "Optional ServiceAccountToken specifies the Kubernetes service account for which to request a token for with the `TokenRequest` API.",
																				"properties": map[string]interface{}{
																					"audiences": map[string]interface{}{
																						"description": "Optional audiences field that will be used to request a temporary Kubernetes service account token for the service account referenced by `serviceAccountRef`. Defaults to a single audience `vault` it not specified.",
																						"items": map[string]interface{}{
																							"type": "string",
																						},
																						"type": "array",
																					},
																					"expirationSeconds": map[string]interface{}{
																						"description": "Optional expiration time in seconds that will be used to request a temporary Kubernetes service account token for the service account referenced by `serviceAccountRef`. Defaults to 10 minutes.",
																						"format":      "int64",
																						"type":        "integer",
																					},
																					"serviceAccountRef": map[string]interface{}{
																						"description": "Service account field containing the name of a kubernetes ServiceAccount.",
																						"properties": map[string]interface{}{
																							"audiences": map[string]interface{}{
																								"description": "Audience specifies the `aud` claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																								"items": map[string]interface{}{
																									"type": "string",
																								},
																								"type": "array",
																							},
																							"name": map[string]interface{}{
																								"description": "The name of the ServiceAccount resource being referred to.",
																								"type":        "string",
																							},
																							"namespace": map[string]interface{}{
																								"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																								"type":        "string",
																							},
																						},
																						"required": []interface{}{
																							"name",
																						},
																						"type": "object",
																					},
																				},
																				"required": []interface{}{
																					"serviceAccountRef",
																				},
																				"type": "object",
																			},
																			"path": map[string]interface{}{
																				"default":     "jwt",
																				"description": "Path where the JWT authentication backend is mounted in Vault, e.g: \"jwt\"",
																				"type":        "string",
																			},
																			"role": map[string]interface{}{
																				"description": "Role is a JWT role to authenticate using the JWT/OIDC Vault authentication method",
																				"type":        "string",
																			},
																			"secretRef": map[string]interface{}{
																				"description": "Optional SecretRef that refers to a key in a Secret resource containing JWT token to authenticate with Vault using the JWT/OIDC authentication method.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"required": []interface{}{
																			"path",
																		},
																		"type": "object",
																	},
																	"kubernetes": map[string]interface{}{
																		"description": "Kubernetes authenticates with Vault by passing the ServiceAccount token stored in the named Secret resource to the Vault server.",
																		"properties": map[string]interface{}{
																			"mountPath": map[string]interface{}{
																				"default":     "kubernetes",
																				"description": "Path where the Kubernetes authentication backend is mounted in Vault, e.g: \"kubernetes\"",
																				"type":        "string",
																			},
																			"role": map[string]interface{}{
																				"description": "A required field containing the Vault Role to assume. A Role binds a Kubernetes ServiceAccount with a set of Vault policies.",
																				"type":        "string",
																			},
																			"secretRef": map[string]interface{}{
																				"description": "Optional secret field containing a Kubernetes ServiceAccount JWT used for authenticating with Vault. If a name is specified without a key, `token` is the default. If one is not specified, the one bound to the controller will be used.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"serviceAccountRef": map[string]interface{}{
																				"description": "Optional service account field containing the name of a kubernetes ServiceAccount. If the service account is specified, the service account secret token JWT will be used for authenticating with Vault. If the service account selector is not supplied, the secretRef will be used instead.",
																				"properties": map[string]interface{}{
																					"audiences": map[string]interface{}{
																						"description": "Audience specifies the `aud` claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																						"items": map[string]interface{}{
																							"type": "string",
																						},
																						"type": "array",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the ServiceAccount resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"required": []interface{}{
																					"name",
																				},
																				"type": "object",
																			},
																		},
																		"required": []interface{}{
																			"mountPath",
																			"role",
																		},
																		"type": "object",
																	},
																	"ldap": map[string]interface{}{
																		"description": "Ldap authenticates with Vault by passing username/password pair using the LDAP authentication method",
																		"properties": map[string]interface{}{
																			"path": map[string]interface{}{
																				"default":     "ldap",
																				"description": "Path where the LDAP authentication backend is mounted in Vault, e.g: \"ldap\"",
																				"type":        "string",
																			},
																			"secretRef": map[string]interface{}{
																				"description": "SecretRef to a key in a Secret resource containing password for the LDAP user used to authenticate with Vault using the LDAP authentication method",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"username": map[string]interface{}{
																				"description": "Username is a LDAP user name used to authenticate using the LDAP Vault authentication method",
																				"type":        "string",
																			},
																		},
																		"required": []interface{}{
																			"path",
																			"username",
																		},
																		"type": "object",
																	},
																	"tokenSecretRef": map[string]interface{}{
																		"description": "TokenSecretRef authenticates with Vault by presenting a token.",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																				"type":        "string",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the Secret resource being referred to.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
															"caBundle": map[string]interface{}{
																"description": "PEM encoded CA bundle used to validate Vault server certificate. Only used if the Server URL is using HTTPS protocol. This parameter is ignored for plain HTTP protocol connection. If not set the system root certificates are used to validate the TLS connection.",
																"format":      "byte",
																"type":        "string",
															},
															"caProvider": map[string]interface{}{
																"description": "The provider for the CA bundle to use to validate Vault server certificate.",
																"properties": map[string]interface{}{
																	"key": map[string]interface{}{
																		"description": "The key the value inside of the provider type to use, only used with \"Secret\" type",
																		"type":        "string",
																	},
																	"name": map[string]interface{}{
																		"description": "The name of the object located at the provider type.",
																		"type":        "string",
																	},
																	"namespace": map[string]interface{}{
																		"description": "The namespace the Provider type is in.",
																		"type":        "string",
																	},
																	"type": map[string]interface{}{
																		"description": "The type of provider to use such as \"Secret\", or \"ConfigMap\".",
																		"enum": []interface{}{
																			"Secret",
																			"ConfigMap",
																		},
																		"type": "string",
																	},
																},
																"required": []interface{}{
																	"name",
																	"type",
																},
																"type": "object",
															},
															"forwardInconsistent": map[string]interface{}{
																"description": "ForwardInconsistent tells Vault to forward read-after-write requests to the Vault leader instead of simply retrying within a loop. This can increase performance if the option is enabled serverside. https://www.vaultproject.io/docs/configuration/replication#allow_forwarding_via_header",
																"type":        "boolean",
															},
															"namespace": map[string]interface{}{
																"description": "Name of the vault namespace. Namespaces is a set of features within Vault Enterprise that allows Vault environments to support Secure Multi-tenancy. e.g: \"ns1\". More about namespaces can be found here https://www.vaultproject.io/docs/enterprise/namespaces",
																"type":        "string",
															},
															"path": map[string]interface{}{
																"description": "Path is the mount path of the Vault KV backend endpoint, e.g: \"secret\". The v2 KV secret engine version specific \"/data\" path suffix for fetching secrets from Vault is optional and will be appended if not present in specified path.",
																"type":        "string",
															},
															"readYourWrites": map[string]interface{}{
																"description": "ReadYourWrites ensures isolated read-after-write semantics by providing discovered cluster replication states in each request. More information about eventual consistency in Vault can be found here https://www.vaultproject.io/docs/enterprise/consistency",
																"type":        "boolean",
															},
															"server": map[string]interface{}{
																"description": "Server is the connection address for the Vault server, e.g: \"https://vault.example.com:8200\".",
																"type":        "string",
															},
															"version": map[string]interface{}{
																"default":     "v2",
																"description": "Version is the Vault KV secret engine version. This can be either \"v1\" or \"v2\". Version defaults to \"v2\".",
																"enum": []interface{}{
																	"v1",
																	"v2",
																},
																"type": "string",
															},
														},
														"required": []interface{}{
															"auth",
															"server",
														},
														"type": "object",
													},
													"webhook": map[string]interface{}{
														"description": "Webhook configures this store to sync secrets using a generic templated webhook",
														"properties": map[string]interface{}{
															"body": map[string]interface{}{
																"description": "Body",
																"type":        "string",
															},
															"caBundle": map[string]interface{}{
																"description": "PEM encoded CA bundle used to validate webhook server certificate. Only used if the Server URL is using HTTPS protocol. This parameter is ignored for plain HTTP protocol connection. If not set the system root certificates are used to validate the TLS connection.",
																"format":      "byte",
																"type":        "string",
															},
															"caProvider": map[string]interface{}{
																"description": "The provider for the CA bundle to use to validate webhook server certificate.",
																"properties": map[string]interface{}{
																	"key": map[string]interface{}{
																		"description": "The key the value inside of the provider type to use, only used with \"Secret\" type",
																		"type":        "string",
																	},
																	"name": map[string]interface{}{
																		"description": "The name of the object located at the provider type.",
																		"type":        "string",
																	},
																	"namespace": map[string]interface{}{
																		"description": "The namespace the Provider type is in.",
																		"type":        "string",
																	},
																	"type": map[string]interface{}{
																		"description": "The type of provider to use such as \"Secret\", or \"ConfigMap\".",
																		"enum": []interface{}{
																			"Secret",
																			"ConfigMap",
																		},
																		"type": "string",
																	},
																},
																"required": []interface{}{
																	"name",
																	"type",
																},
																"type": "object",
															},
															"headers": map[string]interface{}{
																"additionalProperties": map[string]interface{}{
																	"type": "string",
																},
																"description": "Headers",
																"type":        "object",
															},
															"method": map[string]interface{}{
																"description": "Webhook Method",
																"type":        "string",
															},
															"result": map[string]interface{}{
																"description": "Result formatting",
																"properties": map[string]interface{}{
																	"jsonPath": map[string]interface{}{
																		"description": "Json path of return value",
																		"type":        "string",
																	},
																},
																"type": "object",
															},
															"secrets": map[string]interface{}{
																"description": "Secrets to fill in templates These secrets will be passed to the templating function as key value pairs under the given name",
																"items": map[string]interface{}{
																	"properties": map[string]interface{}{
																		"name": map[string]interface{}{
																			"description": "Name of this secret in templates",
																			"type":        "string",
																		},
																		"secretRef": map[string]interface{}{
																			"description": "Secret ref to fill in credentials",
																			"properties": map[string]interface{}{
																				"key": map[string]interface{}{
																					"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																					"type":        "string",
																				},
																				"name": map[string]interface{}{
																					"description": "The name of the Secret resource being referred to.",
																					"type":        "string",
																				},
																				"namespace": map[string]interface{}{
																					"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																					"type":        "string",
																				},
																			},
																			"type": "object",
																		},
																	},
																	"required": []interface{}{
																		"name",
																		"secretRef",
																	},
																	"type": "object",
																},
																"type": "array",
															},
															"timeout": map[string]interface{}{
																"description": "Timeout",
																"type":        "string",
															},
															"url": map[string]interface{}{
																"description": "Webhook url to call",
																"type":        "string",
															},
														},
														"required": []interface{}{
															"result",
															"url",
														},
														"type": "object",
													},
													"yandexlockbox": map[string]interface{}{
														"description": "YandexLockbox configures this store to sync secrets using Yandex Lockbox provider",
														"properties": map[string]interface{}{
															"apiEndpoint": map[string]interface{}{
																"description": "Yandex.Cloud API endpoint (e.g. 'api.cloud.yandex.net:443')",
																"type":        "string",
															},
															"auth": map[string]interface{}{
																"description": "Auth defines the information necessary to authenticate against Yandex Lockbox",
																"properties": map[string]interface{}{
																	"authorizedKeySecretRef": map[string]interface{}{
																		"description": "The authorized key used for authentication",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																				"type":        "string",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the Secret resource being referred to.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
															"caProvider": map[string]interface{}{
																"description": "The provider for the CA bundle to use to validate Yandex.Cloud server certificate.",
																"properties": map[string]interface{}{
																	"certSecretRef": map[string]interface{}{
																		"description": "A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																				"type":        "string",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the Secret resource being referred to.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
														},
														"required": []interface{}{
															"auth",
														},
														"type": "object",
													},
												},
												"type": "object",
											},
											"retrySettings": map[string]interface{}{
												"description": "Used to configure http retries if failed",
												"properties": map[string]interface{}{
													"maxRetries": map[string]interface{}{
														"format": "int32",
														"type":   "integer",
													},
													"retryInterval": map[string]interface{}{
														"type": "string",
													},
												},
												"type": "object",
											},
										},
										"required": []interface{}{
											"provider",
										},
										"type": "object",
									},
									"status": map[string]interface{}{
										"description": "SecretStoreStatus defines the observed state of the SecretStore.",
										"properties": map[string]interface{}{
											"conditions": map[string]interface{}{
												"items": map[string]interface{}{
													"properties": map[string]interface{}{
														"lastTransitionTime": map[string]interface{}{
															"format": "date-time",
															"type":   "string",
														},
														"message": map[string]interface{}{
															"type": "string",
														},
														"reason": map[string]interface{}{
															"type": "string",
														},
														"status": map[string]interface{}{
															"type": "string",
														},
														"type": map[string]interface{}{
															"type": "string",
														},
													},
													"required": []interface{}{
														"status",
														"type",
													},
													"type": "object",
												},
												"type": "array",
											},
										},
										"type": "object",
									},
								},
								"type": "object",
							},
						},
						"served":  true,
						"storage": false,
						"subresources": map[string]interface{}{
							"status": map[string]interface{}{},
						},
					},
					map[string]interface{}{
						"additionalPrinterColumns": []interface{}{
							map[string]interface{}{
								"jsonPath": ".metadata.creationTimestamp",
								"name":     "AGE",
								"type":     "date",
							},
							map[string]interface{}{
								"jsonPath": ".status.conditions[?(@.type==\"Ready\")].reason",
								"name":     "Status",
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".status.capabilities",
								"name":     "Capabilities",
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".status.conditions[?(@.type==\"Ready\")].status",
								"name":     "Ready",
								"type":     "string",
							},
						},
						"name": "v1beta1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"description": "ClusterSecretStore represents a secure external location for storing secrets, which can be referenced as part of `storeRef` fields.",
								"properties": map[string]interface{}{
									"apiVersion": map[string]interface{}{
										"description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
										"type":        "string",
									},
									"kind": map[string]interface{}{
										"description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
										"type":        "string",
									},
									"metadata": map[string]interface{}{
										"type": "object",
									},
									"spec": map[string]interface{}{
										"description": "SecretStoreSpec defines the desired state of SecretStore.",
										"properties": map[string]interface{}{
											"conditions": map[string]interface{}{
												"description": "Used to constraint a ClusterSecretStore to specific namespaces. Relevant only to ClusterSecretStore",
												"items": map[string]interface{}{
													"description": "ClusterSecretStoreCondition describes a condition by which to choose namespaces to process ExternalSecrets in for a ClusterSecretStore instance.",
													"properties": map[string]interface{}{
														"namespaceSelector": map[string]interface{}{
															"description": "Choose namespace using a labelSelector",
															"properties": map[string]interface{}{
																"matchExpressions": map[string]interface{}{
																	"description": "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																	"items": map[string]interface{}{
																		"description": "A label selector requirement is a selector that contains values, a key, and an operator that relates the key and values.",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "key is the label key that the selector applies to.",
																				"type":        "string",
																			},
																			"operator": map[string]interface{}{
																				"description": "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																				"type":        "string",
																			},
																			"values": map[string]interface{}{
																				"description": "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																				"items": map[string]interface{}{
																					"type": "string",
																				},
																				"type": "array",
																			},
																		},
																		"required": []interface{}{
																			"key",
																			"operator",
																		},
																		"type": "object",
																	},
																	"type": "array",
																},
																"matchLabels": map[string]interface{}{
																	"additionalProperties": map[string]interface{}{
																		"type": "string",
																	},
																	"description": "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is \"key\", the operator is \"In\", and the values array contains only \"value\". The requirements are ANDed.",
																	"type":        "object",
																},
															},
															"type":                  "object",
															"x-kubernetes-map-type": "atomic",
														},
														"namespaces": map[string]interface{}{
															"description": "Choose namespaces by name",
															"items": map[string]interface{}{
																"type": "string",
															},
															"type": "array",
														},
													},
													"type": "object",
												},
												"type": "array",
											},
											"controller": map[string]interface{}{
												"description": "Used to select the correct ESO controller (think: ingress.ingressClassName) The ESO controller is instantiated with a specific controller name and filters ES based on this property",
												"type":        "string",
											},
											"provider": map[string]interface{}{
												"description":   "Used to configure the provider. Only one provider may be set",
												"maxProperties": 1,
												"minProperties": 1,
												"properties": map[string]interface{}{
													"akeyless": map[string]interface{}{
														"description": "Akeyless configures this store to sync secrets using Akeyless Vault provider",
														"properties": map[string]interface{}{
															"akeylessGWApiURL": map[string]interface{}{
																"description": "Akeyless GW API Url from which the secrets to be fetched from.",
																"type":        "string",
															},
															"authSecretRef": map[string]interface{}{
																"description": "Auth configures how the operator authenticates with Akeyless.",
																"properties": map[string]interface{}{
																	"kubernetesAuth": map[string]interface{}{
																		"description": "Kubernetes authenticates with Akeyless by passing the ServiceAccount token stored in the named Secret resource.",
																		"properties": map[string]interface{}{
																			"accessID": map[string]interface{}{
																				"description": "the Akeyless Kubernetes auth-method access-id",
																				"type":        "string",
																			},
																			"k8sConfName": map[string]interface{}{
																				"description": "Kubernetes-auth configuration name in Akeyless-Gateway",
																				"type":        "string",
																			},
																			"secretRef": map[string]interface{}{
																				"description": "Optional secret field containing a Kubernetes ServiceAccount JWT used for authenticating with Akeyless. If a name is specified without a key, `token` is the default. If one is not specified, the one bound to the controller will be used.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"serviceAccountRef": map[string]interface{}{
																				"description": "Optional service account field containing the name of a kubernetes ServiceAccount. If the service account is specified, the service account secret token JWT will be used for authenticating with Akeyless. If the service account selector is not supplied, the secretRef will be used instead.",
																				"properties": map[string]interface{}{
																					"audiences": map[string]interface{}{
																						"description": "Audience specifies the `aud` claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																						"items": map[string]interface{}{
																							"type": "string",
																						},
																						"type": "array",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the ServiceAccount resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"required": []interface{}{
																					"name",
																				},
																				"type": "object",
																			},
																		},
																		"required": []interface{}{
																			"accessID",
																			"k8sConfName",
																		},
																		"type": "object",
																	},
																	"secretRef": map[string]interface{}{
																		"description": "Reference to a Secret that contains the details to authenticate with Akeyless.",
																		"properties": map[string]interface{}{
																			"accessID": map[string]interface{}{
																				"description": "The SecretAccessID is used for authentication",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"accessType": map[string]interface{}{
																				"description": "A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"accessTypeParam": map[string]interface{}{
																				"description": "A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
															"caBundle": map[string]interface{}{
																"description": "PEM/base64 encoded CA bundle used to validate Akeyless Gateway certificate. Only used if the AkeylessGWApiURL URL is using HTTPS protocol. If not set the system root certificates are used to validate the TLS connection.",
																"format":      "byte",
																"type":        "string",
															},
															"caProvider": map[string]interface{}{
																"description": "The provider for the CA bundle to use to validate Akeyless Gateway certificate.",
																"properties": map[string]interface{}{
																	"key": map[string]interface{}{
																		"description": "The key where the CA certificate can be found in the Secret or ConfigMap.",
																		"type":        "string",
																	},
																	"name": map[string]interface{}{
																		"description": "The name of the object located at the provider type.",
																		"type":        "string",
																	},
																	"namespace": map[string]interface{}{
																		"description": "The namespace the Provider type is in. Can only be defined when used in a ClusterSecretStore.",
																		"type":        "string",
																	},
																	"type": map[string]interface{}{
																		"description": "The type of provider to use such as \"Secret\", or \"ConfigMap\".",
																		"enum": []interface{}{
																			"Secret",
																			"ConfigMap",
																		},
																		"type": "string",
																	},
																},
																"required": []interface{}{
																	"name",
																	"type",
																},
																"type": "object",
															},
														},
														"required": []interface{}{
															"akeylessGWApiURL",
															"authSecretRef",
														},
														"type": "object",
													},
													"alibaba": map[string]interface{}{
														"description": "Alibaba configures this store to sync secrets using Alibaba Cloud provider",
														"properties": map[string]interface{}{
															"auth": map[string]interface{}{
																"description": "AlibabaAuth contains a secretRef for credentials.",
																"properties": map[string]interface{}{
																	"rrsa": map[string]interface{}{
																		"description": "Authenticate against Alibaba using RRSA.",
																		"properties": map[string]interface{}{
																			"oidcProviderArn": map[string]interface{}{
																				"type": "string",
																			},
																			"oidcTokenFilePath": map[string]interface{}{
																				"type": "string",
																			},
																			"roleArn": map[string]interface{}{
																				"type": "string",
																			},
																			"sessionName": map[string]interface{}{
																				"type": "string",
																			},
																		},
																		"required": []interface{}{
																			"oidcProviderArn",
																			"oidcTokenFilePath",
																			"roleArn",
																			"sessionName",
																		},
																		"type": "object",
																	},
																	"secretRef": map[string]interface{}{
																		"description": "AlibabaAuthSecretRef holds secret references for Alibaba credentials.",
																		"properties": map[string]interface{}{
																			"accessKeyIDSecretRef": map[string]interface{}{
																				"description": "The AccessKeyID is used for authentication",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"accessKeySecretSecretRef": map[string]interface{}{
																				"description": "The AccessKeySecret is used for authentication",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"required": []interface{}{
																			"accessKeyIDSecretRef",
																			"accessKeySecretSecretRef",
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
															"regionID": map[string]interface{}{
																"description": "Alibaba Region to be used for the provider",
																"type":        "string",
															},
														},
														"required": []interface{}{
															"auth",
															"regionID",
														},
														"type": "object",
													},
													"aws": map[string]interface{}{
														"description": "AWS configures this store to sync secrets using AWS Secret Manager provider",
														"properties": map[string]interface{}{
															"additionalRoles": map[string]interface{}{
																"description": "AdditionalRoles is a chained list of Role ARNs which the provider will sequentially assume before assuming the Role",
																"items": map[string]interface{}{
																	"type": "string",
																},
																"type": "array",
															},
															"auth": map[string]interface{}{
																"description": "Auth defines the information necessary to authenticate against AWS if not set aws sdk will infer credentials from your environment see: https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-credentials",
																"properties": map[string]interface{}{
																	"jwt": map[string]interface{}{
																		"description": "Authenticate against AWS using service account tokens.",
																		"properties": map[string]interface{}{
																			"serviceAccountRef": map[string]interface{}{
																				"description": "A reference to a ServiceAccount resource.",
																				"properties": map[string]interface{}{
																					"audiences": map[string]interface{}{
																						"description": "Audience specifies the `aud` claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																						"items": map[string]interface{}{
																							"type": "string",
																						},
																						"type": "array",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the ServiceAccount resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"required": []interface{}{
																					"name",
																				},
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																	"secretRef": map[string]interface{}{
																		"description": "AWSAuthSecretRef holds secret references for AWS credentials both AccessKeyID and SecretAccessKey must be defined in order to properly authenticate.",
																		"properties": map[string]interface{}{
																			"accessKeyIDSecretRef": map[string]interface{}{
																				"description": "The AccessKeyID is used for authentication",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"secretAccessKeySecretRef": map[string]interface{}{
																				"description": "The SecretAccessKey is used for authentication",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"sessionTokenSecretRef": map[string]interface{}{
																				"description": "The SessionToken used for authentication This must be defined if AccessKeyID and SecretAccessKey are temporary credentials see: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_use-resources.html",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
															"externalID": map[string]interface{}{
																"description": "AWS External ID set on assumed IAM roles",
																"type":        "string",
															},
															"region": map[string]interface{}{
																"description": "AWS Region to be used for the provider",
																"type":        "string",
															},
															"role": map[string]interface{}{
																"description": "Role is a Role ARN which the provider will assume",
																"type":        "string",
															},
															"secretsManager": map[string]interface{}{
																"description": "SecretsManager defines how the provider behaves when interacting with AWS SecretsManager",
																"properties": map[string]interface{}{
																	"forceDeleteWithoutRecovery": map[string]interface{}{
																		"description": "Specifies whether to delete the secret without any recovery window. You can't use both this parameter and RecoveryWindowInDays in the same call. If you don't use either, then by default Secrets Manager uses a 30 day recovery window. see: https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_DeleteSecret.html#SecretsManager-DeleteSecret-request-ForceDeleteWithoutRecovery",
																		"type":        "boolean",
																	},
																	"recoveryWindowInDays": map[string]interface{}{
																		"description": "The number of days from 7 to 30 that Secrets Manager waits before permanently deleting the secret. You can't use both this parameter and ForceDeleteWithoutRecovery in the same call. If you don't use either, then by default Secrets Manager uses a 30 day recovery window. see: https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_DeleteSecret.html#SecretsManager-DeleteSecret-request-RecoveryWindowInDays",
																		"format":      "int64",
																		"type":        "integer",
																	},
																},
																"type": "object",
															},
															"service": map[string]interface{}{
																"description": "Service defines which service should be used to fetch the secrets",
																"enum": []interface{}{
																	"SecretsManager",
																	"ParameterStore",
																},
																"type": "string",
															},
															"sessionTags": map[string]interface{}{
																"description": "AWS STS assume role session tags",
																"items": map[string]interface{}{
																	"properties": map[string]interface{}{
																		"key": map[string]interface{}{
																			"type": "string",
																		},
																		"value": map[string]interface{}{
																			"type": "string",
																		},
																	},
																	"required": []interface{}{
																		"key",
																		"value",
																	},
																	"type": "object",
																},
																"type": "array",
															},
															"transitiveTagKeys": map[string]interface{}{
																"description": "AWS STS assume role transitive session tags. Required when multiple rules are used with the provider",
																"items": map[string]interface{}{
																	"type": "string",
																},
																"type": "array",
															},
														},
														"required": []interface{}{
															"region",
															"service",
														},
														"type": "object",
													},
													"azurekv": map[string]interface{}{
														"description": "AzureKV configures this store to sync secrets using Azure Key Vault provider",
														"properties": map[string]interface{}{
															"authSecretRef": map[string]interface{}{
																"description": "Auth configures how the operator authenticates with Azure. Required for ServicePrincipal auth type.",
																"properties": map[string]interface{}{
																	"clientId": map[string]interface{}{
																		"description": "The Azure clientId of the service principle used for authentication.",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																				"type":        "string",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the Secret resource being referred to.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																	"clientSecret": map[string]interface{}{
																		"description": "The Azure ClientSecret of the service principle used for authentication.",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																				"type":        "string",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the Secret resource being referred to.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
															"authType": map[string]interface{}{
																"default":     "ServicePrincipal",
																"description": "Auth type defines how to authenticate to the keyvault service. Valid values are: - \"ServicePrincipal\" (default): Using a service principal (tenantId, clientId, clientSecret) - \"ManagedIdentity\": Using Managed Identity assigned to the pod (see aad-pod-identity)",
																"enum": []interface{}{
																	"ServicePrincipal",
																	"ManagedIdentity",
																	"WorkloadIdentity",
																},
																"type": "string",
															},
															"environmentType": map[string]interface{}{
																"default":     "PublicCloud",
																"description": "EnvironmentType specifies the Azure cloud environment endpoints to use for connecting and authenticating with Azure. By default it points to the public cloud AAD endpoint. The following endpoints are available, also see here: https://github.com/Azure/go-autorest/blob/main/autorest/azure/environments.go#L152 PublicCloud, USGovernmentCloud, ChinaCloud, GermanCloud",
																"enum": []interface{}{
																	"PublicCloud",
																	"USGovernmentCloud",
																	"ChinaCloud",
																	"GermanCloud",
																},
																"type": "string",
															},
															"identityId": map[string]interface{}{
																"description": "If multiple Managed Identity is assigned to the pod, you can select the one to be used",
																"type":        "string",
															},
															"serviceAccountRef": map[string]interface{}{
																"description": "ServiceAccountRef specified the service account that should be used when authenticating with WorkloadIdentity.",
																"properties": map[string]interface{}{
																	"audiences": map[string]interface{}{
																		"description": "Audience specifies the `aud` claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																		"items": map[string]interface{}{
																			"type": "string",
																		},
																		"type": "array",
																	},
																	"name": map[string]interface{}{
																		"description": "The name of the ServiceAccount resource being referred to.",
																		"type":        "string",
																	},
																	"namespace": map[string]interface{}{
																		"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																		"type":        "string",
																	},
																},
																"required": []interface{}{
																	"name",
																},
																"type": "object",
															},
															"tenantId": map[string]interface{}{
																"description": "TenantID configures the Azure Tenant to send requests to. Required for ServicePrincipal auth type.",
																"type":        "string",
															},
															"vaultUrl": map[string]interface{}{
																"description": "Vault Url from which the secrets to be fetched from.",
																"type":        "string",
															},
														},
														"required": []interface{}{
															"vaultUrl",
														},
														"type": "object",
													},
													"conjur": map[string]interface{}{
														"description": "Conjur configures this store to sync secrets using conjur provider",
														"properties": map[string]interface{}{
															"auth": map[string]interface{}{
																"properties": map[string]interface{}{
																	"apikey": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"account": map[string]interface{}{
																				"type": "string",
																			},
																			"apiKeyRef": map[string]interface{}{
																				"description": "A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"userRef": map[string]interface{}{
																				"description": "A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"required": []interface{}{
																			"account",
																			"apiKeyRef",
																			"userRef",
																		},
																		"type": "object",
																	},
																	"jwt": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"account": map[string]interface{}{
																				"type": "string",
																			},
																			"secretRef": map[string]interface{}{
																				"description": "Optional SecretRef that refers to a key in a Secret resource containing JWT token to authenticate with Conjur using the JWT authentication method.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"serviceAccountRef": map[string]interface{}{
																				"description": "Optional ServiceAccountRef specifies the Kubernetes service account for which to request a token for with the `TokenRequest` API.",
																				"properties": map[string]interface{}{
																					"audiences": map[string]interface{}{
																						"description": "Audience specifies the `aud` claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																						"items": map[string]interface{}{
																							"type": "string",
																						},
																						"type": "array",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the ServiceAccount resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"required": []interface{}{
																					"name",
																				},
																				"type": "object",
																			},
																			"serviceID": map[string]interface{}{
																				"description": "The conjur authn jwt webservice id",
																				"type":        "string",
																			},
																		},
																		"required": []interface{}{
																			"account",
																			"serviceID",
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
															"caBundle": map[string]interface{}{
																"type": "string",
															},
															"caProvider": map[string]interface{}{
																"description": "Used to provide custom certificate authority (CA) certificates for a secret store. The CAProvider points to a Secret or ConfigMap resource that contains a PEM-encoded certificate.",
																"properties": map[string]interface{}{
																	"key": map[string]interface{}{
																		"description": "The key where the CA certificate can be found in the Secret or ConfigMap.",
																		"type":        "string",
																	},
																	"name": map[string]interface{}{
																		"description": "The name of the object located at the provider type.",
																		"type":        "string",
																	},
																	"namespace": map[string]interface{}{
																		"description": "The namespace the Provider type is in. Can only be defined when used in a ClusterSecretStore.",
																		"type":        "string",
																	},
																	"type": map[string]interface{}{
																		"description": "The type of provider to use such as \"Secret\", or \"ConfigMap\".",
																		"enum": []interface{}{
																			"Secret",
																			"ConfigMap",
																		},
																		"type": "string",
																	},
																},
																"required": []interface{}{
																	"name",
																	"type",
																},
																"type": "object",
															},
															"url": map[string]interface{}{
																"type": "string",
															},
														},
														"required": []interface{}{
															"auth",
															"url",
														},
														"type": "object",
													},
													"delinea": map[string]interface{}{
														"description": "Delinea DevOps Secrets Vault https://docs.delinea.com/online-help/products/devops-secrets-vault/current",
														"properties": map[string]interface{}{
															"clientId": map[string]interface{}{
																"description": "ClientID is the non-secret part of the credential.",
																"properties": map[string]interface{}{
																	"secretRef": map[string]interface{}{
																		"description": "SecretRef references a key in a secret that will be used as value.",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																				"type":        "string",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the Secret resource being referred to.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																	"value": map[string]interface{}{
																		"description": "Value can be specified directly to set a value without using a secret.",
																		"type":        "string",
																	},
																},
																"type": "object",
															},
															"clientSecret": map[string]interface{}{
																"description": "ClientSecret is the secret part of the credential.",
																"properties": map[string]interface{}{
																	"secretRef": map[string]interface{}{
																		"description": "SecretRef references a key in a secret that will be used as value.",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																				"type":        "string",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the Secret resource being referred to.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																	"value": map[string]interface{}{
																		"description": "Value can be specified directly to set a value without using a secret.",
																		"type":        "string",
																	},
																},
																"type": "object",
															},
															"tenant": map[string]interface{}{
																"description": "Tenant is the chosen hostname / site name.",
																"type":        "string",
															},
															"tld": map[string]interface{}{
																"description": "TLD is based on the server location that was chosen during provisioning. If unset, defaults to \"com\".",
																"type":        "string",
															},
															"urlTemplate": map[string]interface{}{
																"description": "URLTemplate If unset, defaults to \"https://%s.secretsvaultcloud.%s/v1/%s%s\".",
																"type":        "string",
															},
														},
														"required": []interface{}{
															"clientId",
															"clientSecret",
															"tenant",
														},
														"type": "object",
													},
													"doppler": map[string]interface{}{
														"description": "Doppler configures this store to sync secrets using the Doppler provider",
														"properties": map[string]interface{}{
															"auth": map[string]interface{}{
																"description": "Auth configures how the Operator authenticates with the Doppler API",
																"properties": map[string]interface{}{
																	"secretRef": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"dopplerToken": map[string]interface{}{
																				"description": "The DopplerToken is used for authentication. See https://docs.doppler.com/reference/api#authentication for auth token types. The Key attribute defaults to dopplerToken if not specified.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"required": []interface{}{
																			"dopplerToken",
																		},
																		"type": "object",
																	},
																},
																"required": []interface{}{
																	"secretRef",
																},
																"type": "object",
															},
															"config": map[string]interface{}{
																"description": "Doppler config (required if not using a Service Token)",
																"type":        "string",
															},
															"format": map[string]interface{}{
																"description": "Format enables the downloading of secrets as a file (string)",
																"enum": []interface{}{
																	"json",
																	"dotnet-json",
																	"env",
																	"yaml",
																	"docker",
																},
																"type": "string",
															},
															"nameTransformer": map[string]interface{}{
																"description": "Environment variable compatible name transforms that change secret names to a different format",
																"enum": []interface{}{
																	"upper-camel",
																	"camel",
																	"lower-snake",
																	"tf-var",
																	"dotnet-env",
																	"lower-kebab",
																},
																"type": "string",
															},
															"project": map[string]interface{}{
																"description": "Doppler project (required if not using a Service Token)",
																"type":        "string",
															},
														},
														"required": []interface{}{
															"auth",
														},
														"type": "object",
													},
													"fake": map[string]interface{}{
														"description": "Fake configures a store with static key/value pairs",
														"properties": map[string]interface{}{
															"data": map[string]interface{}{
																"items": map[string]interface{}{
																	"properties": map[string]interface{}{
																		"key": map[string]interface{}{
																			"type": "string",
																		},
																		"value": map[string]interface{}{
																			"type": "string",
																		},
																		"valueMap": map[string]interface{}{
																			"additionalProperties": map[string]interface{}{
																				"type": "string",
																			},
																			"description": "Deprecated: ValueMap is deprecated and is intended to be removed in the future, use the `value` field instead.",
																			"type":        "object",
																		},
																		"version": map[string]interface{}{
																			"type": "string",
																		},
																	},
																	"required": []interface{}{
																		"key",
																	},
																	"type": "object",
																},
																"type": "array",
															},
														},
														"required": []interface{}{
															"data",
														},
														"type": "object",
													},
													"gcpsm": map[string]interface{}{
														"description": "GCPSM configures this store to sync secrets using Google Cloud Platform Secret Manager provider",
														"properties": map[string]interface{}{
															"auth": map[string]interface{}{
																"description": "Auth defines the information necessary to authenticate against GCP",
																"properties": map[string]interface{}{
																	"secretRef": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"secretAccessKeySecretRef": map[string]interface{}{
																				"description": "The SecretAccessKey is used for authentication",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																	"workloadIdentity": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"clusterLocation": map[string]interface{}{
																				"type": "string",
																			},
																			"clusterName": map[string]interface{}{
																				"type": "string",
																			},
																			"clusterProjectID": map[string]interface{}{
																				"type": "string",
																			},
																			"serviceAccountRef": map[string]interface{}{
																				"description": "A reference to a ServiceAccount resource.",
																				"properties": map[string]interface{}{
																					"audiences": map[string]interface{}{
																						"description": "Audience specifies the `aud` claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																						"items": map[string]interface{}{
																							"type": "string",
																						},
																						"type": "array",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the ServiceAccount resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"required": []interface{}{
																					"name",
																				},
																				"type": "object",
																			},
																		},
																		"required": []interface{}{
																			"clusterLocation",
																			"clusterName",
																			"serviceAccountRef",
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
															"projectID": map[string]interface{}{
																"description": "ProjectID project where secret is located",
																"type":        "string",
															},
														},
														"type": "object",
													},
													"gitlab": map[string]interface{}{
														"description": "GitLab configures this store to sync secrets using GitLab Variables provider",
														"properties": map[string]interface{}{
															"auth": map[string]interface{}{
																"description": "Auth configures how secret-manager authenticates with a GitLab instance.",
																"properties": map[string]interface{}{
																	"SecretRef": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"accessToken": map[string]interface{}{
																				"description": "AccessToken is used for authentication.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																},
																"required": []interface{}{
																	"SecretRef",
																},
																"type": "object",
															},
															"environment": map[string]interface{}{
																"description": "Environment environment_scope of gitlab CI/CD variables (Please see https://docs.gitlab.com/ee/ci/environments/#create-a-static-environment on how to create environments)",
																"type":        "string",
															},
															"groupIDs": map[string]interface{}{
																"description": "GroupIDs specify, which gitlab groups to pull secrets from. Group secrets are read from left to right followed by the project variables.",
																"items": map[string]interface{}{
																	"type": "string",
																},
																"type": "array",
															},
															"inheritFromGroups": map[string]interface{}{
																"description": "InheritFromGroups specifies whether parent groups should be discovered and checked for secrets.",
																"type":        "boolean",
															},
															"projectID": map[string]interface{}{
																"description": "ProjectID specifies a project where secrets are located.",
																"type":        "string",
															},
															"url": map[string]interface{}{
																"description": "URL configures the GitLab instance URL. Defaults to https://gitlab.com/.",
																"type":        "string",
															},
														},
														"required": []interface{}{
															"auth",
														},
														"type": "object",
													},
													"ibm": map[string]interface{}{
														"description": "IBM configures this store to sync secrets using IBM Cloud provider",
														"properties": map[string]interface{}{
															"auth": map[string]interface{}{
																"description":   "Auth configures how secret-manager authenticates with the IBM secrets manager.",
																"maxProperties": 1,
																"minProperties": 1,
																"properties": map[string]interface{}{
																	"containerAuth": map[string]interface{}{
																		"description": "IBM Container-based auth with IAM Trusted Profile.",
																		"properties": map[string]interface{}{
																			"iamEndpoint": map[string]interface{}{
																				"type": "string",
																			},
																			"profile": map[string]interface{}{
																				"description": "the IBM Trusted Profile",
																				"type":        "string",
																			},
																			"tokenLocation": map[string]interface{}{
																				"description": "Location the token is mounted on the pod",
																				"type":        "string",
																			},
																		},
																		"required": []interface{}{
																			"profile",
																		},
																		"type": "object",
																	},
																	"secretRef": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"secretApiKeySecretRef": map[string]interface{}{
																				"description": "The SecretAccessKey is used for authentication",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
															"serviceUrl": map[string]interface{}{
																"description": "ServiceURL is the Endpoint URL that is specific to the Secrets Manager service instance",
																"type":        "string",
															},
														},
														"required": []interface{}{
															"auth",
														},
														"type": "object",
													},
													"keepersecurity": map[string]interface{}{
														"description": "KeeperSecurity configures this store to sync secrets using the KeeperSecurity provider",
														"properties": map[string]interface{}{
															"authRef": map[string]interface{}{
																"description": "A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.",
																"properties": map[string]interface{}{
																	"key": map[string]interface{}{
																		"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																		"type":        "string",
																	},
																	"name": map[string]interface{}{
																		"description": "The name of the Secret resource being referred to.",
																		"type":        "string",
																	},
																	"namespace": map[string]interface{}{
																		"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																		"type":        "string",
																	},
																},
																"type": "object",
															},
															"folderID": map[string]interface{}{
																"type": "string",
															},
														},
														"required": []interface{}{
															"authRef",
															"folderID",
														},
														"type": "object",
													},
													"kubernetes": map[string]interface{}{
														"description": "Kubernetes configures this store to sync secrets using a Kubernetes cluster provider",
														"properties": map[string]interface{}{
															"auth": map[string]interface{}{
																"description":   "Auth configures how secret-manager authenticates with a Kubernetes instance.",
																"maxProperties": 1,
																"minProperties": 1,
																"properties": map[string]interface{}{
																	"cert": map[string]interface{}{
																		"description": "has both clientCert and clientKey as secretKeySelector",
																		"properties": map[string]interface{}{
																			"clientCert": map[string]interface{}{
																				"description": "A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"clientKey": map[string]interface{}{
																				"description": "A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																	"serviceAccount": map[string]interface{}{
																		"description": "points to a service account that should be used for authentication",
																		"properties": map[string]interface{}{
																			"audiences": map[string]interface{}{
																				"description": "Audience specifies the `aud` claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																				"items": map[string]interface{}{
																					"type": "string",
																				},
																				"type": "array",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the ServiceAccount resource being referred to.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																				"type":        "string",
																			},
																		},
																		"required": []interface{}{
																			"name",
																		},
																		"type": "object",
																	},
																	"token": map[string]interface{}{
																		"description": "use static token to authenticate with",
																		"properties": map[string]interface{}{
																			"bearerToken": map[string]interface{}{
																				"description": "A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
															"remoteNamespace": map[string]interface{}{
																"default":     "default",
																"description": "Remote namespace to fetch the secrets from",
																"type":        "string",
															},
															"server": map[string]interface{}{
																"description": "configures the Kubernetes server Address.",
																"properties": map[string]interface{}{
																	"caBundle": map[string]interface{}{
																		"description": "CABundle is a base64-encoded CA certificate",
																		"format":      "byte",
																		"type":        "string",
																	},
																	"caProvider": map[string]interface{}{
																		"description": "see: https://external-secrets.io/v0.4.1/spec/#external-secrets.io/v1alpha1.CAProvider",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "The key where the CA certificate can be found in the Secret or ConfigMap.",
																				"type":        "string",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the object located at the provider type.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "The namespace the Provider type is in. Can only be defined when used in a ClusterSecretStore.",
																				"type":        "string",
																			},
																			"type": map[string]interface{}{
																				"description": "The type of provider to use such as \"Secret\", or \"ConfigMap\".",
																				"enum": []interface{}{
																					"Secret",
																					"ConfigMap",
																				},
																				"type": "string",
																			},
																		},
																		"required": []interface{}{
																			"name",
																			"type",
																		},
																		"type": "object",
																	},
																	"url": map[string]interface{}{
																		"default":     "kubernetes.default",
																		"description": "configures the Kubernetes server Address.",
																		"type":        "string",
																	},
																},
																"type": "object",
															},
														},
														"required": []interface{}{
															"auth",
														},
														"type": "object",
													},
													"onepassword": map[string]interface{}{
														"description": "OnePassword configures this store to sync secrets using the 1Password Cloud provider",
														"properties": map[string]interface{}{
															"auth": map[string]interface{}{
																"description": "Auth defines the information necessary to authenticate against OnePassword Connect Server",
																"properties": map[string]interface{}{
																	"secretRef": map[string]interface{}{
																		"description": "OnePasswordAuthSecretRef holds secret references for 1Password credentials.",
																		"properties": map[string]interface{}{
																			"connectTokenSecretRef": map[string]interface{}{
																				"description": "The ConnectToken is used for authentication to a 1Password Connect Server.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"required": []interface{}{
																			"connectTokenSecretRef",
																		},
																		"type": "object",
																	},
																},
																"required": []interface{}{
																	"secretRef",
																},
																"type": "object",
															},
															"connectHost": map[string]interface{}{
																"description": "ConnectHost defines the OnePassword Connect Server to connect to",
																"type":        "string",
															},
															"vaults": map[string]interface{}{
																"additionalProperties": map[string]interface{}{
																	"type": "integer",
																},
																"description": "Vaults defines which OnePassword vaults to search in which order",
																"type":        "object",
															},
														},
														"required": []interface{}{
															"auth",
															"connectHost",
															"vaults",
														},
														"type": "object",
													},
													"oracle": map[string]interface{}{
														"description": "Oracle configures this store to sync secrets using Oracle Vault provider",
														"properties": map[string]interface{}{
															"auth": map[string]interface{}{
																"description": "Auth configures how secret-manager authenticates with the Oracle Vault. If empty, use the instance principal, otherwise the user credentials specified in Auth.",
																"properties": map[string]interface{}{
																	"secretRef": map[string]interface{}{
																		"description": "SecretRef to pass through sensitive information.",
																		"properties": map[string]interface{}{
																			"fingerprint": map[string]interface{}{
																				"description": "Fingerprint is the fingerprint of the API private key.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"privatekey": map[string]interface{}{
																				"description": "PrivateKey is the user's API Signing Key in PEM format, used for authentication.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"required": []interface{}{
																			"fingerprint",
																			"privatekey",
																		},
																		"type": "object",
																	},
																	"tenancy": map[string]interface{}{
																		"description": "Tenancy is the tenancy OCID where user is located.",
																		"type":        "string",
																	},
																	"user": map[string]interface{}{
																		"description": "User is an access OCID specific to the account.",
																		"type":        "string",
																	},
																},
																"required": []interface{}{
																	"secretRef",
																	"tenancy",
																	"user",
																},
																"type": "object",
															},
															"compartment": map[string]interface{}{
																"description": "Compartment is the vault compartment OCID. Required for PushSecret",
																"type":        "string",
															},
															"encryptionKey": map[string]interface{}{
																"description": "EncryptionKey is the OCID of the encryption key within the vault. Required for PushSecret",
																"type":        "string",
															},
															"principalType": map[string]interface{}{
																"description": "The type of principal to use for authentication. If left blank, the Auth struct will determine the principal type. This optional field must be specified if using workload identity.",
																"enum": []interface{}{
																	"",
																	"UserPrincipal",
																	"InstancePrincipal",
																	"Workload",
																},
																"type": "string",
															},
															"region": map[string]interface{}{
																"description": "Region is the region where vault is located.",
																"type":        "string",
															},
															"serviceAccountRef": map[string]interface{}{
																"description": "ServiceAccountRef specified the service account that should be used when authenticating with WorkloadIdentity.",
																"properties": map[string]interface{}{
																	"audiences": map[string]interface{}{
																		"description": "Audience specifies the `aud` claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																		"items": map[string]interface{}{
																			"type": "string",
																		},
																		"type": "array",
																	},
																	"name": map[string]interface{}{
																		"description": "The name of the ServiceAccount resource being referred to.",
																		"type":        "string",
																	},
																	"namespace": map[string]interface{}{
																		"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																		"type":        "string",
																	},
																},
																"required": []interface{}{
																	"name",
																},
																"type": "object",
															},
															"vault": map[string]interface{}{
																"description": "Vault is the vault's OCID of the specific vault where secret is located.",
																"type":        "string",
															},
														},
														"required": []interface{}{
															"region",
															"vault",
														},
														"type": "object",
													},
													"scaleway": map[string]interface{}{
														"description": "Scaleway",
														"properties": map[string]interface{}{
															"accessKey": map[string]interface{}{
																"description": "AccessKey is the non-secret part of the api key.",
																"properties": map[string]interface{}{
																	"secretRef": map[string]interface{}{
																		"description": "SecretRef references a key in a secret that will be used as value.",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																				"type":        "string",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the Secret resource being referred to.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																	"value": map[string]interface{}{
																		"description": "Value can be specified directly to set a value without using a secret.",
																		"type":        "string",
																	},
																},
																"type": "object",
															},
															"apiUrl": map[string]interface{}{
																"description": "APIURL is the url of the api to use. Defaults to https://api.scaleway.com",
																"type":        "string",
															},
															"projectId": map[string]interface{}{
																"description": "ProjectID is the id of your project, which you can find in the console: https://console.scaleway.com/project/settings",
																"type":        "string",
															},
															"region": map[string]interface{}{
																"description": "Region where your secrets are located: https://developers.scaleway.com/en/quickstart/#region-and-zone",
																"type":        "string",
															},
															"secretKey": map[string]interface{}{
																"description": "SecretKey is the non-secret part of the api key.",
																"properties": map[string]interface{}{
																	"secretRef": map[string]interface{}{
																		"description": "SecretRef references a key in a secret that will be used as value.",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																				"type":        "string",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the Secret resource being referred to.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																	"value": map[string]interface{}{
																		"description": "Value can be specified directly to set a value without using a secret.",
																		"type":        "string",
																	},
																},
																"type": "object",
															},
														},
														"required": []interface{}{
															"accessKey",
															"projectId",
															"region",
															"secretKey",
														},
														"type": "object",
													},
													"senhasegura": map[string]interface{}{
														"description": "Senhasegura configures this store to sync secrets using senhasegura provider",
														"properties": map[string]interface{}{
															"auth": map[string]interface{}{
																"description": "Auth defines parameters to authenticate in senhasegura",
																"properties": map[string]interface{}{
																	"clientId": map[string]interface{}{
																		"type": "string",
																	},
																	"clientSecretSecretRef": map[string]interface{}{
																		"description": "A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																				"type":        "string",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the Secret resource being referred to.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																},
																"required": []interface{}{
																	"clientId",
																	"clientSecretSecretRef",
																},
																"type": "object",
															},
															"ignoreSslCertificate": map[string]interface{}{
																"default":     false,
																"description": "IgnoreSslCertificate defines if SSL certificate must be ignored",
																"type":        "boolean",
															},
															"module": map[string]interface{}{
																"description": "Module defines which senhasegura module should be used to get secrets",
																"type":        "string",
															},
															"url": map[string]interface{}{
																"description": "URL of senhasegura",
																"type":        "string",
															},
														},
														"required": []interface{}{
															"auth",
															"module",
															"url",
														},
														"type": "object",
													},
													"vault": map[string]interface{}{
														"description": "Vault configures this store to sync secrets using Hashi provider",
														"properties": map[string]interface{}{
															"auth": map[string]interface{}{
																"description": "Auth configures how secret-manager authenticates with the Vault server.",
																"properties": map[string]interface{}{
																	"appRole": map[string]interface{}{
																		"description": "AppRole authenticates with Vault using the App Role auth mechanism, with the role and secret stored in a Kubernetes Secret resource.",
																		"properties": map[string]interface{}{
																			"path": map[string]interface{}{
																				"default":     "approle",
																				"description": "Path where the App Role authentication backend is mounted in Vault, e.g: \"approle\"",
																				"type":        "string",
																			},
																			"roleId": map[string]interface{}{
																				"description": "RoleID configured in the App Role authentication backend when setting up the authentication backend in Vault.",
																				"type":        "string",
																			},
																			"roleRef": map[string]interface{}{
																				"description": "Reference to a key in a Secret that contains the App Role ID used to authenticate with Vault. The `key` field must be specified and denotes which entry within the Secret resource is used as the app role id.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"secretRef": map[string]interface{}{
																				"description": "Reference to a key in a Secret that contains the App Role secret used to authenticate with Vault. The `key` field must be specified and denotes which entry within the Secret resource is used as the app role secret.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"required": []interface{}{
																			"path",
																			"secretRef",
																		},
																		"type": "object",
																	},
																	"cert": map[string]interface{}{
																		"description": "Cert authenticates with TLS Certificates by passing client certificate, private key and ca certificate Cert authentication method",
																		"properties": map[string]interface{}{
																			"clientCert": map[string]interface{}{
																				"description": "ClientCert is a certificate to authenticate using the Cert Vault authentication method",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"secretRef": map[string]interface{}{
																				"description": "SecretRef to a key in a Secret resource containing client private key to authenticate with Vault using the Cert authentication method",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																	"iam": map[string]interface{}{
																		"description": "Iam authenticates with vault by passing a special AWS request signed with AWS IAM credentials AWS IAM authentication method",
																		"properties": map[string]interface{}{
																			"externalID": map[string]interface{}{
																				"description": "AWS External ID set on assumed IAM roles",
																				"type":        "string",
																			},
																			"jwt": map[string]interface{}{
																				"description": "Specify a service account with IRSA enabled",
																				"properties": map[string]interface{}{
																					"serviceAccountRef": map[string]interface{}{
																						"description": "A reference to a ServiceAccount resource.",
																						"properties": map[string]interface{}{
																							"audiences": map[string]interface{}{
																								"description": "Audience specifies the `aud` claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																								"items": map[string]interface{}{
																									"type": "string",
																								},
																								"type": "array",
																							},
																							"name": map[string]interface{}{
																								"description": "The name of the ServiceAccount resource being referred to.",
																								"type":        "string",
																							},
																							"namespace": map[string]interface{}{
																								"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																								"type":        "string",
																							},
																						},
																						"required": []interface{}{
																							"name",
																						},
																						"type": "object",
																					},
																				},
																				"type": "object",
																			},
																			"path": map[string]interface{}{
																				"description": "Path where the AWS auth method is enabled in Vault, e.g: \"aws\"",
																				"type":        "string",
																			},
																			"region": map[string]interface{}{
																				"description": "AWS region",
																				"type":        "string",
																			},
																			"role": map[string]interface{}{
																				"description": "This is the AWS role to be assumed before talking to vault",
																				"type":        "string",
																			},
																			"secretRef": map[string]interface{}{
																				"description": "Specify credentials in a Secret object",
																				"properties": map[string]interface{}{
																					"accessKeyIDSecretRef": map[string]interface{}{
																						"description": "The AccessKeyID is used for authentication",
																						"properties": map[string]interface{}{
																							"key": map[string]interface{}{
																								"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																								"type":        "string",
																							},
																							"name": map[string]interface{}{
																								"description": "The name of the Secret resource being referred to.",
																								"type":        "string",
																							},
																							"namespace": map[string]interface{}{
																								"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																								"type":        "string",
																							},
																						},
																						"type": "object",
																					},
																					"secretAccessKeySecretRef": map[string]interface{}{
																						"description": "The SecretAccessKey is used for authentication",
																						"properties": map[string]interface{}{
																							"key": map[string]interface{}{
																								"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																								"type":        "string",
																							},
																							"name": map[string]interface{}{
																								"description": "The name of the Secret resource being referred to.",
																								"type":        "string",
																							},
																							"namespace": map[string]interface{}{
																								"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																								"type":        "string",
																							},
																						},
																						"type": "object",
																					},
																					"sessionTokenSecretRef": map[string]interface{}{
																						"description": "The SessionToken used for authentication This must be defined if AccessKeyID and SecretAccessKey are temporary credentials see: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_use-resources.html",
																						"properties": map[string]interface{}{
																							"key": map[string]interface{}{
																								"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																								"type":        "string",
																							},
																							"name": map[string]interface{}{
																								"description": "The name of the Secret resource being referred to.",
																								"type":        "string",
																							},
																							"namespace": map[string]interface{}{
																								"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																								"type":        "string",
																							},
																						},
																						"type": "object",
																					},
																				},
																				"type": "object",
																			},
																			"vaultAwsIamServerID": map[string]interface{}{
																				"description": "X-Vault-AWS-IAM-Server-ID is an additional header used by Vault IAM auth method to mitigate against different types of replay attacks. More details here: https://developer.hashicorp.com/vault/docs/auth/aws",
																				"type":        "string",
																			},
																			"vaultRole": map[string]interface{}{
																				"description": "Vault Role. In vault, a role describes an identity with a set of permissions, groups, or policies you want to attach a user of the secrets engine",
																				"type":        "string",
																			},
																		},
																		"required": []interface{}{
																			"vaultRole",
																		},
																		"type": "object",
																	},
																	"jwt": map[string]interface{}{
																		"description": "Jwt authenticates with Vault by passing role and JWT token using the JWT/OIDC authentication method",
																		"properties": map[string]interface{}{
																			"kubernetesServiceAccountToken": map[string]interface{}{
																				"description": "Optional ServiceAccountToken specifies the Kubernetes service account for which to request a token for with the `TokenRequest` API.",
																				"properties": map[string]interface{}{
																					"audiences": map[string]interface{}{
																						"description": "Optional audiences field that will be used to request a temporary Kubernetes service account token for the service account referenced by `serviceAccountRef`. Defaults to a single audience `vault` it not specified. Deprecated: use serviceAccountRef.Audiences instead",
																						"items": map[string]interface{}{
																							"type": "string",
																						},
																						"type": "array",
																					},
																					"expirationSeconds": map[string]interface{}{
																						"description": "Optional expiration time in seconds that will be used to request a temporary Kubernetes service account token for the service account referenced by `serviceAccountRef`. Deprecated: this will be removed in the future. Defaults to 10 minutes.",
																						"format":      "int64",
																						"type":        "integer",
																					},
																					"serviceAccountRef": map[string]interface{}{
																						"description": "Service account field containing the name of a kubernetes ServiceAccount.",
																						"properties": map[string]interface{}{
																							"audiences": map[string]interface{}{
																								"description": "Audience specifies the `aud` claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																								"items": map[string]interface{}{
																									"type": "string",
																								},
																								"type": "array",
																							},
																							"name": map[string]interface{}{
																								"description": "The name of the ServiceAccount resource being referred to.",
																								"type":        "string",
																							},
																							"namespace": map[string]interface{}{
																								"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																								"type":        "string",
																							},
																						},
																						"required": []interface{}{
																							"name",
																						},
																						"type": "object",
																					},
																				},
																				"required": []interface{}{
																					"serviceAccountRef",
																				},
																				"type": "object",
																			},
																			"path": map[string]interface{}{
																				"default":     "jwt",
																				"description": "Path where the JWT authentication backend is mounted in Vault, e.g: \"jwt\"",
																				"type":        "string",
																			},
																			"role": map[string]interface{}{
																				"description": "Role is a JWT role to authenticate using the JWT/OIDC Vault authentication method",
																				"type":        "string",
																			},
																			"secretRef": map[string]interface{}{
																				"description": "Optional SecretRef that refers to a key in a Secret resource containing JWT token to authenticate with Vault using the JWT/OIDC authentication method.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"required": []interface{}{
																			"path",
																		},
																		"type": "object",
																	},
																	"kubernetes": map[string]interface{}{
																		"description": "Kubernetes authenticates with Vault by passing the ServiceAccount token stored in the named Secret resource to the Vault server.",
																		"properties": map[string]interface{}{
																			"mountPath": map[string]interface{}{
																				"default":     "kubernetes",
																				"description": "Path where the Kubernetes authentication backend is mounted in Vault, e.g: \"kubernetes\"",
																				"type":        "string",
																			},
																			"role": map[string]interface{}{
																				"description": "A required field containing the Vault Role to assume. A Role binds a Kubernetes ServiceAccount with a set of Vault policies.",
																				"type":        "string",
																			},
																			"secretRef": map[string]interface{}{
																				"description": "Optional secret field containing a Kubernetes ServiceAccount JWT used for authenticating with Vault. If a name is specified without a key, `token` is the default. If one is not specified, the one bound to the controller will be used.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"serviceAccountRef": map[string]interface{}{
																				"description": "Optional service account field containing the name of a kubernetes ServiceAccount. If the service account is specified, the service account secret token JWT will be used for authenticating with Vault. If the service account selector is not supplied, the secretRef will be used instead.",
																				"properties": map[string]interface{}{
																					"audiences": map[string]interface{}{
																						"description": "Audience specifies the `aud` claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																						"items": map[string]interface{}{
																							"type": "string",
																						},
																						"type": "array",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the ServiceAccount resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"required": []interface{}{
																					"name",
																				},
																				"type": "object",
																			},
																		},
																		"required": []interface{}{
																			"mountPath",
																			"role",
																		},
																		"type": "object",
																	},
																	"ldap": map[string]interface{}{
																		"description": "Ldap authenticates with Vault by passing username/password pair using the LDAP authentication method",
																		"properties": map[string]interface{}{
																			"path": map[string]interface{}{
																				"default":     "ldap",
																				"description": "Path where the LDAP authentication backend is mounted in Vault, e.g: \"ldap\"",
																				"type":        "string",
																			},
																			"secretRef": map[string]interface{}{
																				"description": "SecretRef to a key in a Secret resource containing password for the LDAP user used to authenticate with Vault using the LDAP authentication method",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"username": map[string]interface{}{
																				"description": "Username is a LDAP user name used to authenticate using the LDAP Vault authentication method",
																				"type":        "string",
																			},
																		},
																		"required": []interface{}{
																			"path",
																			"username",
																		},
																		"type": "object",
																	},
																	"tokenSecretRef": map[string]interface{}{
																		"description": "TokenSecretRef authenticates with Vault by presenting a token.",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																				"type":        "string",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the Secret resource being referred to.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																	"userPass": map[string]interface{}{
																		"description": "UserPass authenticates with Vault by passing username/password pair",
																		"properties": map[string]interface{}{
																			"path": map[string]interface{}{
																				"default":     "user",
																				"description": "Path where the UserPassword authentication backend is mounted in Vault, e.g: \"user\"",
																				"type":        "string",
																			},
																			"secretRef": map[string]interface{}{
																				"description": "SecretRef to a key in a Secret resource containing password for the user used to authenticate with Vault using the UserPass authentication method",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"username": map[string]interface{}{
																				"description": "Username is a user name used to authenticate using the UserPass Vault authentication method",
																				"type":        "string",
																			},
																		},
																		"required": []interface{}{
																			"path",
																			"username",
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
															"caBundle": map[string]interface{}{
																"description": "PEM encoded CA bundle used to validate Vault server certificate. Only used if the Server URL is using HTTPS protocol. This parameter is ignored for plain HTTP protocol connection. If not set the system root certificates are used to validate the TLS connection.",
																"format":      "byte",
																"type":        "string",
															},
															"caProvider": map[string]interface{}{
																"description": "The provider for the CA bundle to use to validate Vault server certificate.",
																"properties": map[string]interface{}{
																	"key": map[string]interface{}{
																		"description": "The key where the CA certificate can be found in the Secret or ConfigMap.",
																		"type":        "string",
																	},
																	"name": map[string]interface{}{
																		"description": "The name of the object located at the provider type.",
																		"type":        "string",
																	},
																	"namespace": map[string]interface{}{
																		"description": "The namespace the Provider type is in. Can only be defined when used in a ClusterSecretStore.",
																		"type":        "string",
																	},
																	"type": map[string]interface{}{
																		"description": "The type of provider to use such as \"Secret\", or \"ConfigMap\".",
																		"enum": []interface{}{
																			"Secret",
																			"ConfigMap",
																		},
																		"type": "string",
																	},
																},
																"required": []interface{}{
																	"name",
																	"type",
																},
																"type": "object",
															},
															"forwardInconsistent": map[string]interface{}{
																"description": "ForwardInconsistent tells Vault to forward read-after-write requests to the Vault leader instead of simply retrying within a loop. This can increase performance if the option is enabled serverside. https://www.vaultproject.io/docs/configuration/replication#allow_forwarding_via_header",
																"type":        "boolean",
															},
															"namespace": map[string]interface{}{
																"description": "Name of the vault namespace. Namespaces is a set of features within Vault Enterprise that allows Vault environments to support Secure Multi-tenancy. e.g: \"ns1\". More about namespaces can be found here https://www.vaultproject.io/docs/enterprise/namespaces",
																"type":        "string",
															},
															"path": map[string]interface{}{
																"description": "Path is the mount path of the Vault KV backend endpoint, e.g: \"secret\". The v2 KV secret engine version specific \"/data\" path suffix for fetching secrets from Vault is optional and will be appended if not present in specified path.",
																"type":        "string",
															},
															"readYourWrites": map[string]interface{}{
																"description": "ReadYourWrites ensures isolated read-after-write semantics by providing discovered cluster replication states in each request. More information about eventual consistency in Vault can be found here https://www.vaultproject.io/docs/enterprise/consistency",
																"type":        "boolean",
															},
															"server": map[string]interface{}{
																"description": "Server is the connection address for the Vault server, e.g: \"https://vault.example.com:8200\".",
																"type":        "string",
															},
															"version": map[string]interface{}{
																"default":     "v2",
																"description": "Version is the Vault KV secret engine version. This can be either \"v1\" or \"v2\". Version defaults to \"v2\".",
																"enum": []interface{}{
																	"v1",
																	"v2",
																},
																"type": "string",
															},
														},
														"required": []interface{}{
															"auth",
															"server",
														},
														"type": "object",
													},
													"webhook": map[string]interface{}{
														"description": "Webhook configures this store to sync secrets using a generic templated webhook",
														"properties": map[string]interface{}{
															"body": map[string]interface{}{
																"description": "Body",
																"type":        "string",
															},
															"caBundle": map[string]interface{}{
																"description": "PEM encoded CA bundle used to validate webhook server certificate. Only used if the Server URL is using HTTPS protocol. This parameter is ignored for plain HTTP protocol connection. If not set the system root certificates are used to validate the TLS connection.",
																"format":      "byte",
																"type":        "string",
															},
															"caProvider": map[string]interface{}{
																"description": "The provider for the CA bundle to use to validate webhook server certificate.",
																"properties": map[string]interface{}{
																	"key": map[string]interface{}{
																		"description": "The key the value inside of the provider type to use, only used with \"Secret\" type",
																		"type":        "string",
																	},
																	"name": map[string]interface{}{
																		"description": "The name of the object located at the provider type.",
																		"type":        "string",
																	},
																	"namespace": map[string]interface{}{
																		"description": "The namespace the Provider type is in.",
																		"type":        "string",
																	},
																	"type": map[string]interface{}{
																		"description": "The type of provider to use such as \"Secret\", or \"ConfigMap\".",
																		"enum": []interface{}{
																			"Secret",
																			"ConfigMap",
																		},
																		"type": "string",
																	},
																},
																"required": []interface{}{
																	"name",
																	"type",
																},
																"type": "object",
															},
															"headers": map[string]interface{}{
																"additionalProperties": map[string]interface{}{
																	"type": "string",
																},
																"description": "Headers",
																"type":        "object",
															},
															"method": map[string]interface{}{
																"description": "Webhook Method",
																"type":        "string",
															},
															"result": map[string]interface{}{
																"description": "Result formatting",
																"properties": map[string]interface{}{
																	"jsonPath": map[string]interface{}{
																		"description": "Json path of return value",
																		"type":        "string",
																	},
																},
																"type": "object",
															},
															"secrets": map[string]interface{}{
																"description": "Secrets to fill in templates These secrets will be passed to the templating function as key value pairs under the given name",
																"items": map[string]interface{}{
																	"properties": map[string]interface{}{
																		"name": map[string]interface{}{
																			"description": "Name of this secret in templates",
																			"type":        "string",
																		},
																		"secretRef": map[string]interface{}{
																			"description": "Secret ref to fill in credentials",
																			"properties": map[string]interface{}{
																				"key": map[string]interface{}{
																					"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																					"type":        "string",
																				},
																				"name": map[string]interface{}{
																					"description": "The name of the Secret resource being referred to.",
																					"type":        "string",
																				},
																				"namespace": map[string]interface{}{
																					"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																					"type":        "string",
																				},
																			},
																			"type": "object",
																		},
																	},
																	"required": []interface{}{
																		"name",
																		"secretRef",
																	},
																	"type": "object",
																},
																"type": "array",
															},
															"timeout": map[string]interface{}{
																"description": "Timeout",
																"type":        "string",
															},
															"url": map[string]interface{}{
																"description": "Webhook url to call",
																"type":        "string",
															},
														},
														"required": []interface{}{
															"result",
															"url",
														},
														"type": "object",
													},
													"yandexcertificatemanager": map[string]interface{}{
														"description": "YandexCertificateManager configures this store to sync secrets using Yandex Certificate Manager provider",
														"properties": map[string]interface{}{
															"apiEndpoint": map[string]interface{}{
																"description": "Yandex.Cloud API endpoint (e.g. 'api.cloud.yandex.net:443')",
																"type":        "string",
															},
															"auth": map[string]interface{}{
																"description": "Auth defines the information necessary to authenticate against Yandex Certificate Manager",
																"properties": map[string]interface{}{
																	"authorizedKeySecretRef": map[string]interface{}{
																		"description": "The authorized key used for authentication",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																				"type":        "string",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the Secret resource being referred to.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
															"caProvider": map[string]interface{}{
																"description": "The provider for the CA bundle to use to validate Yandex.Cloud server certificate.",
																"properties": map[string]interface{}{
																	"certSecretRef": map[string]interface{}{
																		"description": "A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																				"type":        "string",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the Secret resource being referred to.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
														},
														"required": []interface{}{
															"auth",
														},
														"type": "object",
													},
													"yandexlockbox": map[string]interface{}{
														"description": "YandexLockbox configures this store to sync secrets using Yandex Lockbox provider",
														"properties": map[string]interface{}{
															"apiEndpoint": map[string]interface{}{
																"description": "Yandex.Cloud API endpoint (e.g. 'api.cloud.yandex.net:443')",
																"type":        "string",
															},
															"auth": map[string]interface{}{
																"description": "Auth defines the information necessary to authenticate against Yandex Lockbox",
																"properties": map[string]interface{}{
																	"authorizedKeySecretRef": map[string]interface{}{
																		"description": "The authorized key used for authentication",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																				"type":        "string",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the Secret resource being referred to.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
															"caProvider": map[string]interface{}{
																"description": "The provider for the CA bundle to use to validate Yandex.Cloud server certificate.",
																"properties": map[string]interface{}{
																	"certSecretRef": map[string]interface{}{
																		"description": "A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																				"type":        "string",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the Secret resource being referred to.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
														},
														"required": []interface{}{
															"auth",
														},
														"type": "object",
													},
												},
												"type": "object",
											},
											"refreshInterval": map[string]interface{}{
												"description": "Used to configure store refresh interval in seconds. Empty or 0 will default to the controller config.",
												"type":        "integer",
											},
											"retrySettings": map[string]interface{}{
												"description": "Used to configure http retries if failed",
												"properties": map[string]interface{}{
													"maxRetries": map[string]interface{}{
														"format": "int32",
														"type":   "integer",
													},
													"retryInterval": map[string]interface{}{
														"type": "string",
													},
												},
												"type": "object",
											},
										},
										"required": []interface{}{
											"provider",
										},
										"type": "object",
									},
									"status": map[string]interface{}{
										"description": "SecretStoreStatus defines the observed state of the SecretStore.",
										"properties": map[string]interface{}{
											"capabilities": map[string]interface{}{
												"description": "SecretStoreCapabilities defines the possible operations a SecretStore can do.",
												"type":        "string",
											},
											"conditions": map[string]interface{}{
												"items": map[string]interface{}{
													"properties": map[string]interface{}{
														"lastTransitionTime": map[string]interface{}{
															"format": "date-time",
															"type":   "string",
														},
														"message": map[string]interface{}{
															"type": "string",
														},
														"reason": map[string]interface{}{
															"type": "string",
														},
														"status": map[string]interface{}{
															"type": "string",
														},
														"type": map[string]interface{}{
															"type": "string",
														},
													},
													"required": []interface{}{
														"status",
														"type",
													},
													"type": "object",
												},
												"type": "array",
											},
										},
										"type": "object",
									},
								},
								"type": "object",
							},
						},
						"served":  true,
						"storage": true,
						"subresources": map[string]interface{}{
							"status": map[string]interface{}{},
						},
					},
				},
				"conversion": map[string]interface{}{
					"strategy": "Webhook",
					"webhook": map[string]interface{}{
						"conversionReviewVersions": []interface{}{
							"v1",
						},
						"clientConfig": map[string]interface{}{
							"service": map[string]interface{}{
								"name":      "release-name-external-secrets-webhook",
								"namespace": "default",
								"path":      "/convert",
							},
						},
					},
				},
			},
		},
	}

	return mutate.MutateCRDClustersecretstoresExternalSecretsIo(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;create;update;patch;delete

// CreateCRDEcrauthorizationtokensGeneratorsExternalSecretsIo creates the CustomResourceDefinition resource with name ecrauthorizationtokens.generators.external-secrets.io.
func CreateCRDEcrauthorizationtokensGeneratorsExternalSecretsIo(
	parent *secretsv1alpha1.ExternalSecrets,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// Source: external-secrets/templates/crds/ecrauthorizationtoken.yaml
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					"controller-gen.kubebuilder.io/version": "v0.13.0",
				},
				"name": "ecrauthorizationtokens.generators.external-secrets.io",
			},
			"spec": map[string]interface{}{
				"group": "generators.external-secrets.io",
				"names": map[string]interface{}{
					"categories": []interface{}{
						"ecrauthorizationtoken",
					},
					"kind":     "ECRAuthorizationToken",
					"listKind": "ECRAuthorizationTokenList",
					"plural":   "ecrauthorizationtokens",
					"shortNames": []interface{}{
						"ecrauthorizationtoken",
					},
					"singular": "ecrauthorizationtoken",
				},
				"scope": "Namespaced",
				"versions": []interface{}{
					map[string]interface{}{
						"name": "v1alpha1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"description": "ECRAuthorizationTokenSpec uses the GetAuthorizationToken API to retrieve an authorization token. The authorization token is valid for 12 hours. The authorizationToken returned is a base64 encoded string that can be decoded and used in a docker login command to authenticate to a registry. For more information, see Registry authentication (https://docs.aws.amazon.com/AmazonECR/latest/userguide/Registries.html#registry_auth) in the Amazon Elastic Container Registry User Guide.",
								"properties": map[string]interface{}{
									"apiVersion": map[string]interface{}{
										"description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
										"type":        "string",
									},
									"kind": map[string]interface{}{
										"description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
										"type":        "string",
									},
									"metadata": map[string]interface{}{
										"type": "object",
									},
									"spec": map[string]interface{}{
										"properties": map[string]interface{}{
											"auth": map[string]interface{}{
												"description": "Auth defines how to authenticate with AWS",
												"properties": map[string]interface{}{
													"jwt": map[string]interface{}{
														"description": "Authenticate against AWS using service account tokens.",
														"properties": map[string]interface{}{
															"serviceAccountRef": map[string]interface{}{
																"description": "A reference to a ServiceAccount resource.",
																"properties": map[string]interface{}{
																	"audiences": map[string]interface{}{
																		"description": "Audience specifies the `aud` claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																		"items": map[string]interface{}{
																			"type": "string",
																		},
																		"type": "array",
																	},
																	"name": map[string]interface{}{
																		"description": "The name of the ServiceAccount resource being referred to.",
																		"type":        "string",
																	},
																	"namespace": map[string]interface{}{
																		"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																		"type":        "string",
																	},
																},
																"required": []interface{}{
																	"name",
																},
																"type": "object",
															},
														},
														"type": "object",
													},
													"secretRef": map[string]interface{}{
														"description": "AWSAuthSecretRef holds secret references for AWS credentials both AccessKeyID and SecretAccessKey must be defined in order to properly authenticate.",
														"properties": map[string]interface{}{
															"accessKeyIDSecretRef": map[string]interface{}{
																"description": "The AccessKeyID is used for authentication",
																"properties": map[string]interface{}{
																	"key": map[string]interface{}{
																		"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																		"type":        "string",
																	},
																	"name": map[string]interface{}{
																		"description": "The name of the Secret resource being referred to.",
																		"type":        "string",
																	},
																	"namespace": map[string]interface{}{
																		"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																		"type":        "string",
																	},
																},
																"type": "object",
															},
															"secretAccessKeySecretRef": map[string]interface{}{
																"description": "The SecretAccessKey is used for authentication",
																"properties": map[string]interface{}{
																	"key": map[string]interface{}{
																		"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																		"type":        "string",
																	},
																	"name": map[string]interface{}{
																		"description": "The name of the Secret resource being referred to.",
																		"type":        "string",
																	},
																	"namespace": map[string]interface{}{
																		"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																		"type":        "string",
																	},
																},
																"type": "object",
															},
															"sessionTokenSecretRef": map[string]interface{}{
																"description": "The SessionToken used for authentication This must be defined if AccessKeyID and SecretAccessKey are temporary credentials see: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_use-resources.html",
																"properties": map[string]interface{}{
																	"key": map[string]interface{}{
																		"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																		"type":        "string",
																	},
																	"name": map[string]interface{}{
																		"description": "The name of the Secret resource being referred to.",
																		"type":        "string",
																	},
																	"namespace": map[string]interface{}{
																		"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																		"type":        "string",
																	},
																},
																"type": "object",
															},
														},
														"type": "object",
													},
												},
												"type": "object",
											},
											"region": map[string]interface{}{
												"description": "Region specifies the region to operate in.",
												"type":        "string",
											},
											"role": map[string]interface{}{
												"description": "You can assume a role before making calls to the desired AWS service.",
												"type":        "string",
											},
										},
										"required": []interface{}{
											"region",
										},
										"type": "object",
									},
								},
								"type": "object",
							},
						},
						"served":  true,
						"storage": true,
						"subresources": map[string]interface{}{
							"status": map[string]interface{}{},
						},
					},
				},
				"conversion": map[string]interface{}{
					"strategy": "Webhook",
					"webhook": map[string]interface{}{
						"conversionReviewVersions": []interface{}{
							"v1",
						},
						"clientConfig": map[string]interface{}{
							"service": map[string]interface{}{
								"name":      "release-name-external-secrets-webhook",
								"namespace": "default",
								"path":      "/convert",
							},
						},
					},
				},
			},
		},
	}

	return mutate.MutateCRDEcrauthorizationtokensGeneratorsExternalSecretsIo(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;create;update;patch;delete

// CreateCRDExternalsecretsExternalSecretsIo creates the CustomResourceDefinition resource with name externalsecrets.external-secrets.io.
func CreateCRDExternalsecretsExternalSecretsIo(
	parent *secretsv1alpha1.ExternalSecrets,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// Source: external-secrets/templates/crds/externalsecret.yaml
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					"controller-gen.kubebuilder.io/version": "v0.13.0",
				},
				"name": "externalsecrets.external-secrets.io",
			},
			"spec": map[string]interface{}{
				"group": "external-secrets.io",
				"names": map[string]interface{}{
					"categories": []interface{}{
						"externalsecrets",
					},
					"kind":     "ExternalSecret",
					"listKind": "ExternalSecretList",
					"plural":   "externalsecrets",
					"shortNames": []interface{}{
						"es",
					},
					"singular": "externalsecret",
				},
				"scope": "Namespaced",
				"versions": []interface{}{
					map[string]interface{}{
						"additionalPrinterColumns": []interface{}{
							map[string]interface{}{
								"jsonPath": ".spec.secretStoreRef.name",
								"name":     "Store",
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".spec.refreshInterval",
								"name":     "Refresh Interval",
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".status.conditions[?(@.type==\"Ready\")].reason",
								"name":     "Status",
								"type":     "string",
							},
						},
						"deprecated": true,
						"name":       "v1alpha1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"description": "ExternalSecret is the Schema for the external-secrets API.",
								"properties": map[string]interface{}{
									"apiVersion": map[string]interface{}{
										"description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
										"type":        "string",
									},
									"kind": map[string]interface{}{
										"description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
										"type":        "string",
									},
									"metadata": map[string]interface{}{
										"type": "object",
									},
									"spec": map[string]interface{}{
										"description": "ExternalSecretSpec defines the desired state of ExternalSecret.",
										"properties": map[string]interface{}{
											"data": map[string]interface{}{
												"description": "Data defines the connection between the Kubernetes Secret keys and the Provider data",
												"items": map[string]interface{}{
													"description": "ExternalSecretData defines the connection between the Kubernetes Secret key (spec.data.<key>) and the Provider data.",
													"properties": map[string]interface{}{
														"remoteRef": map[string]interface{}{
															"description": "ExternalSecretDataRemoteRef defines Provider data location.",
															"properties": map[string]interface{}{
																"conversionStrategy": map[string]interface{}{
																	"default":     "Default",
																	"description": "Used to define a conversion Strategy",
																	"enum": []interface{}{
																		"Default",
																		"Unicode",
																	},
																	"type": "string",
																},
																"key": map[string]interface{}{
																	"description": "Key is the key used in the Provider, mandatory",
																	"type":        "string",
																},
																"property": map[string]interface{}{
																	"description": "Used to select a specific property of the Provider value (if a map), if supported",
																	"type":        "string",
																},
																"version": map[string]interface{}{
																	"description": "Used to select a specific version of the Provider value, if supported",
																	"type":        "string",
																},
															},
															"required": []interface{}{
																"key",
															},
															"type": "object",
														},
														"secretKey": map[string]interface{}{
															"type": "string",
														},
													},
													"required": []interface{}{
														"remoteRef",
														"secretKey",
													},
													"type": "object",
												},
												"type": "array",
											},
											"dataFrom": map[string]interface{}{
												"description": "DataFrom is used to fetch all properties from a specific Provider data If multiple entries are specified, the Secret keys are merged in the specified order",
												"items": map[string]interface{}{
													"description": "ExternalSecretDataRemoteRef defines Provider data location.",
													"properties": map[string]interface{}{
														"conversionStrategy": map[string]interface{}{
															"default":     "Default",
															"description": "Used to define a conversion Strategy",
															"enum": []interface{}{
																"Default",
																"Unicode",
															},
															"type": "string",
														},
														"key": map[string]interface{}{
															"description": "Key is the key used in the Provider, mandatory",
															"type":        "string",
														},
														"property": map[string]interface{}{
															"description": "Used to select a specific property of the Provider value (if a map), if supported",
															"type":        "string",
														},
														"version": map[string]interface{}{
															"description": "Used to select a specific version of the Provider value, if supported",
															"type":        "string",
														},
													},
													"required": []interface{}{
														"key",
													},
													"type": "object",
												},
												"type": "array",
											},
											"refreshInterval": map[string]interface{}{
												"default":     "1h",
												"description": "RefreshInterval is the amount of time before the values are read again from the SecretStore provider Valid time units are \"ns\", \"us\" (or \"µs\"), \"ms\", \"s\", \"m\", \"h\" May be set to zero to fetch and create it once. Defaults to 1h.",
												"type":        "string",
											},
											"secretStoreRef": map[string]interface{}{
												"description": "SecretStoreRef defines which SecretStore to fetch the ExternalSecret data.",
												"properties": map[string]interface{}{
													"kind": map[string]interface{}{
														"description": "Kind of the SecretStore resource (SecretStore or ClusterSecretStore) Defaults to `SecretStore`",
														"type":        "string",
													},
													"name": map[string]interface{}{
														"description": "Name of the SecretStore resource",
														"type":        "string",
													},
												},
												"required": []interface{}{
													"name",
												},
												"type": "object",
											},
											"target": map[string]interface{}{
												"description": "ExternalSecretTarget defines the Kubernetes Secret to be created There can be only one target per ExternalSecret.",
												"properties": map[string]interface{}{
													"creationPolicy": map[string]interface{}{
														"default":     "Owner",
														"description": "CreationPolicy defines rules on how to create the resulting Secret Defaults to 'Owner'",
														"enum": []interface{}{
															"Owner",
															"Merge",
															"None",
														},
														"type": "string",
													},
													"immutable": map[string]interface{}{
														"description": "Immutable defines if the final secret will be immutable",
														"type":        "boolean",
													},
													"name": map[string]interface{}{
														"description": "Name defines the name of the Secret resource to be managed This field is immutable Defaults to the .metadata.name of the ExternalSecret resource",
														"type":        "string",
													},
													"template": map[string]interface{}{
														"description": "Template defines a blueprint for the created Secret resource.",
														"properties": map[string]interface{}{
															"data": map[string]interface{}{
																"additionalProperties": map[string]interface{}{
																	"type": "string",
																},
																"type": "object",
															},
															"engineVersion": map[string]interface{}{
																"default":     "v1",
																"description": "EngineVersion specifies the template engine version that should be used to compile/execute the template specified in .data and .templateFrom[].",
																"enum": []interface{}{
																	"v1",
																	"v2",
																},
																"type": "string",
															},
															"metadata": map[string]interface{}{
																"description": "ExternalSecretTemplateMetadata defines metadata fields for the Secret blueprint.",
																"properties": map[string]interface{}{
																	"annotations": map[string]interface{}{
																		"additionalProperties": map[string]interface{}{
																			"type": "string",
																		},
																		"type": "object",
																	},
																	"labels": map[string]interface{}{
																		"additionalProperties": map[string]interface{}{
																			"type": "string",
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
															"templateFrom": map[string]interface{}{
																"items": map[string]interface{}{
																	"maxProperties": 1,
																	"minProperties": 1,
																	"properties": map[string]interface{}{
																		"configMap": map[string]interface{}{
																			"properties": map[string]interface{}{
																				"items": map[string]interface{}{
																					"items": map[string]interface{}{
																						"properties": map[string]interface{}{
																							"key": map[string]interface{}{
																								"type": "string",
																							},
																						},
																						"required": []interface{}{
																							"key",
																						},
																						"type": "object",
																					},
																					"type": "array",
																				},
																				"name": map[string]interface{}{
																					"type": "string",
																				},
																			},
																			"required": []interface{}{
																				"items",
																				"name",
																			},
																			"type": "object",
																		},
																		"secret": map[string]interface{}{
																			"properties": map[string]interface{}{
																				"items": map[string]interface{}{
																					"items": map[string]interface{}{
																						"properties": map[string]interface{}{
																							"key": map[string]interface{}{
																								"type": "string",
																							},
																						},
																						"required": []interface{}{
																							"key",
																						},
																						"type": "object",
																					},
																					"type": "array",
																				},
																				"name": map[string]interface{}{
																					"type": "string",
																				},
																			},
																			"required": []interface{}{
																				"items",
																				"name",
																			},
																			"type": "object",
																		},
																	},
																	"type": "object",
																},
																"type": "array",
															},
															"type": map[string]interface{}{
																"type": "string",
															},
														},
														"type": "object",
													},
												},
												"type": "object",
											},
										},
										"required": []interface{}{
											"secretStoreRef",
											"target",
										},
										"type": "object",
									},
									"status": map[string]interface{}{
										"properties": map[string]interface{}{
											"binding": map[string]interface{}{
												"description": "Binding represents a servicebinding.io Provisioned Service reference to the secret",
												"properties": map[string]interface{}{
													"name": map[string]interface{}{
														"description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														"type":        "string",
													},
												},
												"type":                  "object",
												"x-kubernetes-map-type": "atomic",
											},
											"conditions": map[string]interface{}{
												"items": map[string]interface{}{
													"properties": map[string]interface{}{
														"lastTransitionTime": map[string]interface{}{
															"format": "date-time",
															"type":   "string",
														},
														"message": map[string]interface{}{
															"type": "string",
														},
														"reason": map[string]interface{}{
															"type": "string",
														},
														"status": map[string]interface{}{
															"type": "string",
														},
														"type": map[string]interface{}{
															"type": "string",
														},
													},
													"required": []interface{}{
														"status",
														"type",
													},
													"type": "object",
												},
												"type": "array",
											},
											"refreshTime": map[string]interface{}{
												"description": "refreshTime is the time and date the external secret was fetched and the target secret updated",
												"format":      "date-time",
												"nullable":    true,
												"type":        "string",
											},
											"syncedResourceVersion": map[string]interface{}{
												"description": "SyncedResourceVersion keeps track of the last synced version",
												"type":        "string",
											},
										},
										"type": "object",
									},
								},
								"type": "object",
							},
						},
						"served":  true,
						"storage": false,
						"subresources": map[string]interface{}{
							"status": map[string]interface{}{},
						},
					},
					map[string]interface{}{
						"additionalPrinterColumns": []interface{}{
							map[string]interface{}{
								"jsonPath": ".spec.secretStoreRef.name",
								"name":     "Store",
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".spec.refreshInterval",
								"name":     "Refresh Interval",
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".status.conditions[?(@.type==\"Ready\")].reason",
								"name":     "Status",
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".status.conditions[?(@.type==\"Ready\")].status",
								"name":     "Ready",
								"type":     "string",
							},
						},
						"name": "v1beta1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"description": "ExternalSecret is the Schema for the external-secrets API.",
								"properties": map[string]interface{}{
									"apiVersion": map[string]interface{}{
										"description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
										"type":        "string",
									},
									"kind": map[string]interface{}{
										"description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
										"type":        "string",
									},
									"metadata": map[string]interface{}{
										"type": "object",
									},
									"spec": map[string]interface{}{
										"description": "ExternalSecretSpec defines the desired state of ExternalSecret.",
										"properties": map[string]interface{}{
											"data": map[string]interface{}{
												"description": "Data defines the connection between the Kubernetes Secret keys and the Provider data",
												"items": map[string]interface{}{
													"description": "ExternalSecretData defines the connection between the Kubernetes Secret key (spec.data.<key>) and the Provider data.",
													"properties": map[string]interface{}{
														"remoteRef": map[string]interface{}{
															"description": "RemoteRef points to the remote secret and defines which secret (version/property/..) to fetch.",
															"properties": map[string]interface{}{
																"conversionStrategy": map[string]interface{}{
																	"default":     "Default",
																	"description": "Used to define a conversion Strategy",
																	"enum": []interface{}{
																		"Default",
																		"Unicode",
																	},
																	"type": "string",
																},
																"decodingStrategy": map[string]interface{}{
																	"default":     "None",
																	"description": "Used to define a decoding Strategy",
																	"enum": []interface{}{
																		"Auto",
																		"Base64",
																		"Base64URL",
																		"None",
																	},
																	"type": "string",
																},
																"key": map[string]interface{}{
																	"description": "Key is the key used in the Provider, mandatory",
																	"type":        "string",
																},
																"metadataPolicy": map[string]interface{}{
																	"default":     "None",
																	"description": "Policy for fetching tags/labels from provider secrets, possible options are Fetch, None. Defaults to None",
																	"enum": []interface{}{
																		"None",
																		"Fetch",
																	},
																	"type": "string",
																},
																"property": map[string]interface{}{
																	"description": "Used to select a specific property of the Provider value (if a map), if supported",
																	"type":        "string",
																},
																"version": map[string]interface{}{
																	"description": "Used to select a specific version of the Provider value, if supported",
																	"type":        "string",
																},
															},
															"required": []interface{}{
																"key",
															},
															"type": "object",
														},
														"secretKey": map[string]interface{}{
															"description": "SecretKey defines the key in which the controller stores the value. This is the key in the Kind=Secret",
															"type":        "string",
														},
														"sourceRef": map[string]interface{}{
															"description":   "SourceRef allows you to override the source from which the value will pulled from.",
															"maxProperties": 1,
															"properties": map[string]interface{}{
																"generatorRef": map[string]interface{}{
																	"description": `GeneratorRef points to a generator custom resource. 
 Deprecated: The generatorRef is not implemented in .data[]. this will be removed with v1.`,
																	"properties": map[string]interface{}{
																		"apiVersion": map[string]interface{}{
																			"default":     "generators.external-secrets.io/v1alpha1",
																			"description": "Specify the apiVersion of the generator resource",
																			"type":        "string",
																		},
																		"kind": map[string]interface{}{
																			"description": "Specify the Kind of the resource, e.g. Password, ACRAccessToken etc.",
																			"type":        "string",
																		},
																		"name": map[string]interface{}{
																			"description": "Specify the name of the generator resource",
																			"type":        "string",
																		},
																	},
																	"required": []interface{}{
																		"kind",
																		"name",
																	},
																	"type": "object",
																},
																"storeRef": map[string]interface{}{
																	"description": "SecretStoreRef defines which SecretStore to fetch the ExternalSecret data.",
																	"properties": map[string]interface{}{
																		"kind": map[string]interface{}{
																			"description": "Kind of the SecretStore resource (SecretStore or ClusterSecretStore) Defaults to `SecretStore`",
																			"type":        "string",
																		},
																		"name": map[string]interface{}{
																			"description": "Name of the SecretStore resource",
																			"type":        "string",
																		},
																	},
																	"required": []interface{}{
																		"name",
																	},
																	"type": "object",
																},
															},
															"type": "object",
														},
													},
													"required": []interface{}{
														"remoteRef",
														"secretKey",
													},
													"type": "object",
												},
												"type": "array",
											},
											"dataFrom": map[string]interface{}{
												"description": "DataFrom is used to fetch all properties from a specific Provider data If multiple entries are specified, the Secret keys are merged in the specified order",
												"items": map[string]interface{}{
													"properties": map[string]interface{}{
														"extract": map[string]interface{}{
															"description": "Used to extract multiple key/value pairs from one secret Note: Extract does not support sourceRef.Generator or sourceRef.GeneratorRef.",
															"properties": map[string]interface{}{
																"conversionStrategy": map[string]interface{}{
																	"default":     "Default",
																	"description": "Used to define a conversion Strategy",
																	"enum": []interface{}{
																		"Default",
																		"Unicode",
																	},
																	"type": "string",
																},
																"decodingStrategy": map[string]interface{}{
																	"default":     "None",
																	"description": "Used to define a decoding Strategy",
																	"enum": []interface{}{
																		"Auto",
																		"Base64",
																		"Base64URL",
																		"None",
																	},
																	"type": "string",
																},
																"key": map[string]interface{}{
																	"description": "Key is the key used in the Provider, mandatory",
																	"type":        "string",
																},
																"metadataPolicy": map[string]interface{}{
																	"default":     "None",
																	"description": "Policy for fetching tags/labels from provider secrets, possible options are Fetch, None. Defaults to None",
																	"enum": []interface{}{
																		"None",
																		"Fetch",
																	},
																	"type": "string",
																},
																"property": map[string]interface{}{
																	"description": "Used to select a specific property of the Provider value (if a map), if supported",
																	"type":        "string",
																},
																"version": map[string]interface{}{
																	"description": "Used to select a specific version of the Provider value, if supported",
																	"type":        "string",
																},
															},
															"required": []interface{}{
																"key",
															},
															"type": "object",
														},
														"find": map[string]interface{}{
															"description": "Used to find secrets based on tags or regular expressions Note: Find does not support sourceRef.Generator or sourceRef.GeneratorRef.",
															"properties": map[string]interface{}{
																"conversionStrategy": map[string]interface{}{
																	"default":     "Default",
																	"description": "Used to define a conversion Strategy",
																	"enum": []interface{}{
																		"Default",
																		"Unicode",
																	},
																	"type": "string",
																},
																"decodingStrategy": map[string]interface{}{
																	"default":     "None",
																	"description": "Used to define a decoding Strategy",
																	"enum": []interface{}{
																		"Auto",
																		"Base64",
																		"Base64URL",
																		"None",
																	},
																	"type": "string",
																},
																"name": map[string]interface{}{
																	"description": "Finds secrets based on the name.",
																	"properties": map[string]interface{}{
																		"regexp": map[string]interface{}{
																			"description": "Finds secrets base",
																			"type":        "string",
																		},
																	},
																	"type": "object",
																},
																"path": map[string]interface{}{
																	"description": "A root path to start the find operations.",
																	"type":        "string",
																},
																"tags": map[string]interface{}{
																	"additionalProperties": map[string]interface{}{
																		"type": "string",
																	},
																	"description": "Find secrets based on tags.",
																	"type":        "object",
																},
															},
															"type": "object",
														},
														"rewrite": map[string]interface{}{
															"description": "Used to rewrite secret Keys after getting them from the secret Provider Multiple Rewrite operations can be provided. They are applied in a layered order (first to last)",
															"items": map[string]interface{}{
																"properties": map[string]interface{}{
																	"regexp": map[string]interface{}{
																		"description": "Used to rewrite with regular expressions. The resulting key will be the output of a regexp.ReplaceAll operation.",
																		"properties": map[string]interface{}{
																			"source": map[string]interface{}{
																				"description": "Used to define the regular expression of a re.Compiler.",
																				"type":        "string",
																			},
																			"target": map[string]interface{}{
																				"description": "Used to define the target pattern of a ReplaceAll operation.",
																				"type":        "string",
																			},
																		},
																		"required": []interface{}{
																			"source",
																			"target",
																		},
																		"type": "object",
																	},
																	"transform": map[string]interface{}{
																		"description": "Used to apply string transformation on the secrets. The resulting key will be the output of the template applied by the operation.",
																		"properties": map[string]interface{}{
																			"template": map[string]interface{}{
																				"description": "Used to define the template to apply on the secret name. `.value ` will specify the secret name in the template.",
																				"type":        "string",
																			},
																		},
																		"required": []interface{}{
																			"template",
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
															"type": "array",
														},
														"sourceRef": map[string]interface{}{
															"description":   "SourceRef points to a store or generator which contains secret values ready to use. Use this in combination with Extract or Find pull values out of a specific SecretStore. When sourceRef points to a generator Extract or Find is not supported. The generator returns a static map of values",
															"maxProperties": 1,
															"properties": map[string]interface{}{
																"generatorRef": map[string]interface{}{
																	"description": "GeneratorRef points to a generator custom resource.",
																	"properties": map[string]interface{}{
																		"apiVersion": map[string]interface{}{
																			"default":     "generators.external-secrets.io/v1alpha1",
																			"description": "Specify the apiVersion of the generator resource",
																			"type":        "string",
																		},
																		"kind": map[string]interface{}{
																			"description": "Specify the Kind of the resource, e.g. Password, ACRAccessToken etc.",
																			"type":        "string",
																		},
																		"name": map[string]interface{}{
																			"description": "Specify the name of the generator resource",
																			"type":        "string",
																		},
																	},
																	"required": []interface{}{
																		"kind",
																		"name",
																	},
																	"type": "object",
																},
																"storeRef": map[string]interface{}{
																	"description": "SecretStoreRef defines which SecretStore to fetch the ExternalSecret data.",
																	"properties": map[string]interface{}{
																		"kind": map[string]interface{}{
																			"description": "Kind of the SecretStore resource (SecretStore or ClusterSecretStore) Defaults to `SecretStore`",
																			"type":        "string",
																		},
																		"name": map[string]interface{}{
																			"description": "Name of the SecretStore resource",
																			"type":        "string",
																		},
																	},
																	"required": []interface{}{
																		"name",
																	},
																	"type": "object",
																},
															},
															"type": "object",
														},
													},
													"type": "object",
												},
												"type": "array",
											},
											"refreshInterval": map[string]interface{}{
												"default":     "1h",
												"description": "RefreshInterval is the amount of time before the values are read again from the SecretStore provider Valid time units are \"ns\", \"us\" (or \"µs\"), \"ms\", \"s\", \"m\", \"h\" May be set to zero to fetch and create it once. Defaults to 1h.",
												"type":        "string",
											},
											"secretStoreRef": map[string]interface{}{
												"description": "SecretStoreRef defines which SecretStore to fetch the ExternalSecret data.",
												"properties": map[string]interface{}{
													"kind": map[string]interface{}{
														"description": "Kind of the SecretStore resource (SecretStore or ClusterSecretStore) Defaults to `SecretStore`",
														"type":        "string",
													},
													"name": map[string]interface{}{
														"description": "Name of the SecretStore resource",
														"type":        "string",
													},
												},
												"required": []interface{}{
													"name",
												},
												"type": "object",
											},
											"target": map[string]interface{}{
												"default": map[string]interface{}{
													"creationPolicy": "Owner",
													"deletionPolicy": "Retain",
												},
												"description": "ExternalSecretTarget defines the Kubernetes Secret to be created There can be only one target per ExternalSecret.",
												"properties": map[string]interface{}{
													"creationPolicy": map[string]interface{}{
														"default":     "Owner",
														"description": "CreationPolicy defines rules on how to create the resulting Secret Defaults to 'Owner'",
														"enum": []interface{}{
															"Owner",
															"Orphan",
															"Merge",
															"None",
														},
														"type": "string",
													},
													"deletionPolicy": map[string]interface{}{
														"default":     "Retain",
														"description": "DeletionPolicy defines rules on how to delete the resulting Secret Defaults to 'Retain'",
														"enum": []interface{}{
															"Delete",
															"Merge",
															"Retain",
														},
														"type": "string",
													},
													"immutable": map[string]interface{}{
														"description": "Immutable defines if the final secret will be immutable",
														"type":        "boolean",
													},
													"name": map[string]interface{}{
														"description": "Name defines the name of the Secret resource to be managed This field is immutable Defaults to the .metadata.name of the ExternalSecret resource",
														"type":        "string",
													},
													"template": map[string]interface{}{
														"description": "Template defines a blueprint for the created Secret resource.",
														"properties": map[string]interface{}{
															"data": map[string]interface{}{
																"additionalProperties": map[string]interface{}{
																	"type": "string",
																},
																"type": "object",
															},
															"engineVersion": map[string]interface{}{
																"default":     "v2",
																"description": "EngineVersion specifies the template engine version that should be used to compile/execute the template specified in .data and .templateFrom[].",
																"enum": []interface{}{
																	"v1",
																	"v2",
																},
																"type": "string",
															},
															"mergePolicy": map[string]interface{}{
																"default": "Replace",
																"enum": []interface{}{
																	"Replace",
																	"Merge",
																},
																"type": "string",
															},
															"metadata": map[string]interface{}{
																"description": "ExternalSecretTemplateMetadata defines metadata fields for the Secret blueprint.",
																"properties": map[string]interface{}{
																	"annotations": map[string]interface{}{
																		"additionalProperties": map[string]interface{}{
																			"type": "string",
																		},
																		"type": "object",
																	},
																	"labels": map[string]interface{}{
																		"additionalProperties": map[string]interface{}{
																			"type": "string",
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
															"templateFrom": map[string]interface{}{
																"items": map[string]interface{}{
																	"properties": map[string]interface{}{
																		"configMap": map[string]interface{}{
																			"properties": map[string]interface{}{
																				"items": map[string]interface{}{
																					"items": map[string]interface{}{
																						"properties": map[string]interface{}{
																							"key": map[string]interface{}{
																								"type": "string",
																							},
																							"templateAs": map[string]interface{}{
																								"default": "Values",
																								"enum": []interface{}{
																									"Values",
																									"KeysAndValues",
																								},
																								"type": "string",
																							},
																						},
																						"required": []interface{}{
																							"key",
																						},
																						"type": "object",
																					},
																					"type": "array",
																				},
																				"name": map[string]interface{}{
																					"type": "string",
																				},
																			},
																			"required": []interface{}{
																				"items",
																				"name",
																			},
																			"type": "object",
																		},
																		"literal": map[string]interface{}{
																			"type": "string",
																		},
																		"secret": map[string]interface{}{
																			"properties": map[string]interface{}{
																				"items": map[string]interface{}{
																					"items": map[string]interface{}{
																						"properties": map[string]interface{}{
																							"key": map[string]interface{}{
																								"type": "string",
																							},
																							"templateAs": map[string]interface{}{
																								"default": "Values",
																								"enum": []interface{}{
																									"Values",
																									"KeysAndValues",
																								},
																								"type": "string",
																							},
																						},
																						"required": []interface{}{
																							"key",
																						},
																						"type": "object",
																					},
																					"type": "array",
																				},
																				"name": map[string]interface{}{
																					"type": "string",
																				},
																			},
																			"required": []interface{}{
																				"items",
																				"name",
																			},
																			"type": "object",
																		},
																		"target": map[string]interface{}{
																			"default": "Data",
																			"enum": []interface{}{
																				"Data",
																				"Annotations",
																				"Labels",
																			},
																			"type": "string",
																		},
																	},
																	"type": "object",
																},
																"type": "array",
															},
															"type": map[string]interface{}{
																"type": "string",
															},
														},
														"type": "object",
													},
												},
												"type": "object",
											},
										},
										"type": "object",
									},
									"status": map[string]interface{}{
										"properties": map[string]interface{}{
											"binding": map[string]interface{}{
												"description": "Binding represents a servicebinding.io Provisioned Service reference to the secret",
												"properties": map[string]interface{}{
													"name": map[string]interface{}{
														"description": "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														"type":        "string",
													},
												},
												"type":                  "object",
												"x-kubernetes-map-type": "atomic",
											},
											"conditions": map[string]interface{}{
												"items": map[string]interface{}{
													"properties": map[string]interface{}{
														"lastTransitionTime": map[string]interface{}{
															"format": "date-time",
															"type":   "string",
														},
														"message": map[string]interface{}{
															"type": "string",
														},
														"reason": map[string]interface{}{
															"type": "string",
														},
														"status": map[string]interface{}{
															"type": "string",
														},
														"type": map[string]interface{}{
															"type": "string",
														},
													},
													"required": []interface{}{
														"status",
														"type",
													},
													"type": "object",
												},
												"type": "array",
											},
											"refreshTime": map[string]interface{}{
												"description": "refreshTime is the time and date the external secret was fetched and the target secret updated",
												"format":      "date-time",
												"nullable":    true,
												"type":        "string",
											},
											"syncedResourceVersion": map[string]interface{}{
												"description": "SyncedResourceVersion keeps track of the last synced version",
												"type":        "string",
											},
										},
										"type": "object",
									},
								},
								"type": "object",
							},
						},
						"served":  true,
						"storage": true,
						"subresources": map[string]interface{}{
							"status": map[string]interface{}{},
						},
					},
				},
				"conversion": map[string]interface{}{
					"strategy": "Webhook",
					"webhook": map[string]interface{}{
						"conversionReviewVersions": []interface{}{
							"v1",
						},
						"clientConfig": map[string]interface{}{
							"service": map[string]interface{}{
								"name":      "release-name-external-secrets-webhook",
								"namespace": "default",
								"path":      "/convert",
							},
						},
					},
				},
			},
		},
	}

	return mutate.MutateCRDExternalsecretsExternalSecretsIo(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;create;update;patch;delete

// CreateCRDFakesGeneratorsExternalSecretsIo creates the CustomResourceDefinition resource with name fakes.generators.external-secrets.io.
func CreateCRDFakesGeneratorsExternalSecretsIo(
	parent *secretsv1alpha1.ExternalSecrets,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// Source: external-secrets/templates/crds/fake.yaml
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					"controller-gen.kubebuilder.io/version": "v0.13.0",
				},
				"name": "fakes.generators.external-secrets.io",
			},
			"spec": map[string]interface{}{
				"group": "generators.external-secrets.io",
				"names": map[string]interface{}{
					"categories": []interface{}{
						"fake",
					},
					"kind":     "Fake",
					"listKind": "FakeList",
					"plural":   "fakes",
					"shortNames": []interface{}{
						"fake",
					},
					"singular": "fake",
				},
				"scope": "Namespaced",
				"versions": []interface{}{
					map[string]interface{}{
						"name": "v1alpha1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"description": "Fake generator is used for testing. It lets you define a static set of credentials that is always returned.",
								"properties": map[string]interface{}{
									"apiVersion": map[string]interface{}{
										"description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
										"type":        "string",
									},
									"kind": map[string]interface{}{
										"description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
										"type":        "string",
									},
									"metadata": map[string]interface{}{
										"type": "object",
									},
									"spec": map[string]interface{}{
										"description": "FakeSpec contains the static data.",
										"properties": map[string]interface{}{
											"controller": map[string]interface{}{
												"description": "Used to select the correct ESO controller (think: ingress.ingressClassName) The ESO controller is instantiated with a specific controller name and filters VDS based on this property",
												"type":        "string",
											},
											"data": map[string]interface{}{
												"additionalProperties": map[string]interface{}{
													"type": "string",
												},
												"description": "Data defines the static data returned by this generator.",
												"type":        "object",
											},
										},
										"type": "object",
									},
								},
								"type": "object",
							},
						},
						"served":  true,
						"storage": true,
						"subresources": map[string]interface{}{
							"status": map[string]interface{}{},
						},
					},
				},
				"conversion": map[string]interface{}{
					"strategy": "Webhook",
					"webhook": map[string]interface{}{
						"conversionReviewVersions": []interface{}{
							"v1",
						},
						"clientConfig": map[string]interface{}{
							"service": map[string]interface{}{
								"name":      "release-name-external-secrets-webhook",
								"namespace": "default",
								"path":      "/convert",
							},
						},
					},
				},
			},
		},
	}

	return mutate.MutateCRDFakesGeneratorsExternalSecretsIo(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;create;update;patch;delete

// CreateCRDGcraccesstokensGeneratorsExternalSecretsIo creates the CustomResourceDefinition resource with name gcraccesstokens.generators.external-secrets.io.
func CreateCRDGcraccesstokensGeneratorsExternalSecretsIo(
	parent *secretsv1alpha1.ExternalSecrets,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// Source: external-secrets/templates/crds/gcraccesstoken.yaml
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					"controller-gen.kubebuilder.io/version": "v0.13.0",
				},
				"name": "gcraccesstokens.generators.external-secrets.io",
			},
			"spec": map[string]interface{}{
				"group": "generators.external-secrets.io",
				"names": map[string]interface{}{
					"categories": []interface{}{
						"gcraccesstoken",
					},
					"kind":     "GCRAccessToken",
					"listKind": "GCRAccessTokenList",
					"plural":   "gcraccesstokens",
					"shortNames": []interface{}{
						"gcraccesstoken",
					},
					"singular": "gcraccesstoken",
				},
				"scope": "Namespaced",
				"versions": []interface{}{
					map[string]interface{}{
						"name": "v1alpha1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"description": "GCRAccessToken generates an GCP access token that can be used to authenticate with GCR.",
								"properties": map[string]interface{}{
									"apiVersion": map[string]interface{}{
										"description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
										"type":        "string",
									},
									"kind": map[string]interface{}{
										"description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
										"type":        "string",
									},
									"metadata": map[string]interface{}{
										"type": "object",
									},
									"spec": map[string]interface{}{
										"properties": map[string]interface{}{
											"auth": map[string]interface{}{
												"description": "Auth defines the means for authenticating with GCP",
												"properties": map[string]interface{}{
													"secretRef": map[string]interface{}{
														"properties": map[string]interface{}{
															"secretAccessKeySecretRef": map[string]interface{}{
																"description": "The SecretAccessKey is used for authentication",
																"properties": map[string]interface{}{
																	"key": map[string]interface{}{
																		"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																		"type":        "string",
																	},
																	"name": map[string]interface{}{
																		"description": "The name of the Secret resource being referred to.",
																		"type":        "string",
																	},
																	"namespace": map[string]interface{}{
																		"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																		"type":        "string",
																	},
																},
																"type": "object",
															},
														},
														"type": "object",
													},
													"workloadIdentity": map[string]interface{}{
														"properties": map[string]interface{}{
															"clusterLocation": map[string]interface{}{
																"type": "string",
															},
															"clusterName": map[string]interface{}{
																"type": "string",
															},
															"clusterProjectID": map[string]interface{}{
																"type": "string",
															},
															"serviceAccountRef": map[string]interface{}{
																"description": "A reference to a ServiceAccount resource.",
																"properties": map[string]interface{}{
																	"audiences": map[string]interface{}{
																		"description": "Audience specifies the `aud` claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																		"items": map[string]interface{}{
																			"type": "string",
																		},
																		"type": "array",
																	},
																	"name": map[string]interface{}{
																		"description": "The name of the ServiceAccount resource being referred to.",
																		"type":        "string",
																	},
																	"namespace": map[string]interface{}{
																		"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																		"type":        "string",
																	},
																},
																"required": []interface{}{
																	"name",
																},
																"type": "object",
															},
														},
														"required": []interface{}{
															"clusterLocation",
															"clusterName",
															"serviceAccountRef",
														},
														"type": "object",
													},
												},
												"type": "object",
											},
											"projectID": map[string]interface{}{
												"description": "ProjectID defines which project to use to authenticate with",
												"type":        "string",
											},
										},
										"required": []interface{}{
											"auth",
											"projectID",
										},
										"type": "object",
									},
								},
								"type": "object",
							},
						},
						"served":  true,
						"storage": true,
						"subresources": map[string]interface{}{
							"status": map[string]interface{}{},
						},
					},
				},
				"conversion": map[string]interface{}{
					"strategy": "Webhook",
					"webhook": map[string]interface{}{
						"conversionReviewVersions": []interface{}{
							"v1",
						},
						"clientConfig": map[string]interface{}{
							"service": map[string]interface{}{
								"name":      "release-name-external-secrets-webhook",
								"namespace": "default",
								"path":      "/convert",
							},
						},
					},
				},
			},
		},
	}

	return mutate.MutateCRDGcraccesstokensGeneratorsExternalSecretsIo(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;create;update;patch;delete

// CreateCRDPasswordsGeneratorsExternalSecretsIo creates the CustomResourceDefinition resource with name passwords.generators.external-secrets.io.
func CreateCRDPasswordsGeneratorsExternalSecretsIo(
	parent *secretsv1alpha1.ExternalSecrets,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// Source: external-secrets/templates/crds/password.yaml
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					"controller-gen.kubebuilder.io/version": "v0.13.0",
				},
				"name": "passwords.generators.external-secrets.io",
			},
			"spec": map[string]interface{}{
				"group": "generators.external-secrets.io",
				"names": map[string]interface{}{
					"categories": []interface{}{
						"password",
					},
					"kind":     "Password",
					"listKind": "PasswordList",
					"plural":   "passwords",
					"shortNames": []interface{}{
						"password",
					},
					"singular": "password",
				},
				"scope": "Namespaced",
				"versions": []interface{}{
					map[string]interface{}{
						"name": "v1alpha1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"description": "Password generates a random password based on the configuration parameters in spec. You can specify the length, characterset and other attributes.",
								"properties": map[string]interface{}{
									"apiVersion": map[string]interface{}{
										"description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
										"type":        "string",
									},
									"kind": map[string]interface{}{
										"description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
										"type":        "string",
									},
									"metadata": map[string]interface{}{
										"type": "object",
									},
									"spec": map[string]interface{}{
										"description": "PasswordSpec controls the behavior of the password generator.",
										"properties": map[string]interface{}{
											"allowRepeat": map[string]interface{}{
												"default":     false,
												"description": "set AllowRepeat to true to allow repeating characters.",
												"type":        "boolean",
											},
											"digits": map[string]interface{}{
												"description": "Digits specifies the number of digits in the generated password. If omitted it defaults to 25% of the length of the password",
												"type":        "integer",
											},
											"length": map[string]interface{}{
												"default":     24,
												"description": "Length of the password to be generated. Defaults to 24",
												"type":        "integer",
											},
											"noUpper": map[string]interface{}{
												"default":     false,
												"description": "Set NoUpper to disable uppercase characters",
												"type":        "boolean",
											},
											"symbolCharacters": map[string]interface{}{
												"description": "SymbolCharacters specifies the special characters that should be used in the generated password.",
												"type":        "string",
											},
											"symbols": map[string]interface{}{
												"description": "Symbols specifies the number of symbol characters in the generated password. If omitted it defaults to 25% of the length of the password",
												"type":        "integer",
											},
										},
										"required": []interface{}{
											"allowRepeat",
											"length",
											"noUpper",
										},
										"type": "object",
									},
								},
								"type": "object",
							},
						},
						"served":  true,
						"storage": true,
						"subresources": map[string]interface{}{
							"status": map[string]interface{}{},
						},
					},
				},
				"conversion": map[string]interface{}{
					"strategy": "Webhook",
					"webhook": map[string]interface{}{
						"conversionReviewVersions": []interface{}{
							"v1",
						},
						"clientConfig": map[string]interface{}{
							"service": map[string]interface{}{
								"name":      "release-name-external-secrets-webhook",
								"namespace": "default",
								"path":      "/convert",
							},
						},
					},
				},
			},
		},
	}

	return mutate.MutateCRDPasswordsGeneratorsExternalSecretsIo(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;create;update;patch;delete

// CreateCRDPushsecretsExternalSecretsIo creates the CustomResourceDefinition resource with name pushsecrets.external-secrets.io.
func CreateCRDPushsecretsExternalSecretsIo(
	parent *secretsv1alpha1.ExternalSecrets,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// Source: external-secrets/templates/crds/pushsecret.yaml
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					"controller-gen.kubebuilder.io/version": "v0.13.0",
				},
				"name": "pushsecrets.external-secrets.io",
			},
			"spec": map[string]interface{}{
				"group": "external-secrets.io",
				"names": map[string]interface{}{
					"categories": []interface{}{
						"pushsecrets",
					},
					"kind":     "PushSecret",
					"listKind": "PushSecretList",
					"plural":   "pushsecrets",
					"singular": "pushsecret",
				},
				"scope": "Namespaced",
				"versions": []interface{}{
					map[string]interface{}{
						"additionalPrinterColumns": []interface{}{
							map[string]interface{}{
								"jsonPath": ".metadata.creationTimestamp",
								"name":     "AGE",
								"type":     "date",
							},
							map[string]interface{}{
								"jsonPath": ".status.conditions[?(@.type==\"Ready\")].reason",
								"name":     "Status",
								"type":     "string",
							},
						},
						"name": "v1alpha1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"properties": map[string]interface{}{
									"apiVersion": map[string]interface{}{
										"description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
										"type":        "string",
									},
									"kind": map[string]interface{}{
										"description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
										"type":        "string",
									},
									"metadata": map[string]interface{}{
										"type": "object",
									},
									"spec": map[string]interface{}{
										"description": "PushSecretSpec configures the behavior of the PushSecret.",
										"properties": map[string]interface{}{
											"data": map[string]interface{}{
												"description": "Secret Data that should be pushed to providers",
												"items": map[string]interface{}{
													"properties": map[string]interface{}{
														"match": map[string]interface{}{
															"description": "Match a given Secret Key to be pushed to the provider.",
															"properties": map[string]interface{}{
																"remoteRef": map[string]interface{}{
																	"description": "Remote Refs to push to providers.",
																	"properties": map[string]interface{}{
																		"property": map[string]interface{}{
																			"description": "Name of the property in the resulting secret",
																			"type":        "string",
																		},
																		"remoteKey": map[string]interface{}{
																			"description": "Name of the resulting provider secret.",
																			"type":        "string",
																		},
																	},
																	"required": []interface{}{
																		"remoteKey",
																	},
																	"type": "object",
																},
																"secretKey": map[string]interface{}{
																	"description": "Secret Key to be pushed",
																	"type":        "string",
																},
															},
															"required": []interface{}{
																"remoteRef",
															},
															"type": "object",
														},
														"metadata": map[string]interface{}{
															"description":                          "Metadata is metadata attached to the secret. The structure of metadata is provider specific, please look it up in the provider documentation.",
															"x-kubernetes-preserve-unknown-fields": true,
														},
													},
													"required": []interface{}{
														"match",
													},
													"type": "object",
												},
												"type": "array",
											},
											"deletionPolicy": map[string]interface{}{
												"default":     "None",
												"description": "Deletion Policy to handle Secrets in the provider. Possible Values: \"Delete/None\". Defaults to \"None\".",
												"enum": []interface{}{
													"Delete",
													"None",
												},
												"type": "string",
											},
											"refreshInterval": map[string]interface{}{
												"description": "The Interval to which External Secrets will try to push a secret definition",
												"type":        "string",
											},
											"secretStoreRefs": map[string]interface{}{
												"items": map[string]interface{}{
													"properties": map[string]interface{}{
														"kind": map[string]interface{}{
															"default":     "SecretStore",
															"description": "Kind of the SecretStore resource (SecretStore or ClusterSecretStore) Defaults to `SecretStore`",
															"type":        "string",
														},
														"labelSelector": map[string]interface{}{
															"description": "Optionally, sync to secret stores with label selector",
															"properties": map[string]interface{}{
																"matchExpressions": map[string]interface{}{
																	"description": "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																	"items": map[string]interface{}{
																		"description": "A label selector requirement is a selector that contains values, a key, and an operator that relates the key and values.",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "key is the label key that the selector applies to.",
																				"type":        "string",
																			},
																			"operator": map[string]interface{}{
																				"description": "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																				"type":        "string",
																			},
																			"values": map[string]interface{}{
																				"description": "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																				"items": map[string]interface{}{
																					"type": "string",
																				},
																				"type": "array",
																			},
																		},
																		"required": []interface{}{
																			"key",
																			"operator",
																		},
																		"type": "object",
																	},
																	"type": "array",
																},
																"matchLabels": map[string]interface{}{
																	"additionalProperties": map[string]interface{}{
																		"type": "string",
																	},
																	"description": "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is \"key\", the operator is \"In\", and the values array contains only \"value\". The requirements are ANDed.",
																	"type":        "object",
																},
															},
															"type":                  "object",
															"x-kubernetes-map-type": "atomic",
														},
														"name": map[string]interface{}{
															"description": "Optionally, sync to the SecretStore of the given name",
															"type":        "string",
														},
													},
													"type": "object",
												},
												"type": "array",
											},
											"selector": map[string]interface{}{
												"description": "The Secret Selector (k8s source) for the Push Secret",
												"properties": map[string]interface{}{
													"secret": map[string]interface{}{
														"description": "Select a Secret to Push.",
														"properties": map[string]interface{}{
															"name": map[string]interface{}{
																"description": "Name of the Secret. The Secret must exist in the same namespace as the PushSecret manifest.",
																"type":        "string",
															},
														},
														"required": []interface{}{
															"name",
														},
														"type": "object",
													},
												},
												"required": []interface{}{
													"secret",
												},
												"type": "object",
											},
											"template": map[string]interface{}{
												"description": "Template defines a blueprint for the created Secret resource.",
												"properties": map[string]interface{}{
													"data": map[string]interface{}{
														"additionalProperties": map[string]interface{}{
															"type": "string",
														},
														"type": "object",
													},
													"engineVersion": map[string]interface{}{
														"default":     "v2",
														"description": "EngineVersion specifies the template engine version that should be used to compile/execute the template specified in .data and .templateFrom[].",
														"enum": []interface{}{
															"v1",
															"v2",
														},
														"type": "string",
													},
													"mergePolicy": map[string]interface{}{
														"default": "Replace",
														"enum": []interface{}{
															"Replace",
															"Merge",
														},
														"type": "string",
													},
													"metadata": map[string]interface{}{
														"description": "ExternalSecretTemplateMetadata defines metadata fields for the Secret blueprint.",
														"properties": map[string]interface{}{
															"annotations": map[string]interface{}{
																"additionalProperties": map[string]interface{}{
																	"type": "string",
																},
																"type": "object",
															},
															"labels": map[string]interface{}{
																"additionalProperties": map[string]interface{}{
																	"type": "string",
																},
																"type": "object",
															},
														},
														"type": "object",
													},
													"templateFrom": map[string]interface{}{
														"items": map[string]interface{}{
															"properties": map[string]interface{}{
																"configMap": map[string]interface{}{
																	"properties": map[string]interface{}{
																		"items": map[string]interface{}{
																			"items": map[string]interface{}{
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"type": "string",
																					},
																					"templateAs": map[string]interface{}{
																						"default": "Values",
																						"enum": []interface{}{
																							"Values",
																							"KeysAndValues",
																						},
																						"type": "string",
																					},
																				},
																				"required": []interface{}{
																					"key",
																				},
																				"type": "object",
																			},
																			"type": "array",
																		},
																		"name": map[string]interface{}{
																			"type": "string",
																		},
																	},
																	"required": []interface{}{
																		"items",
																		"name",
																	},
																	"type": "object",
																},
																"literal": map[string]interface{}{
																	"type": "string",
																},
																"secret": map[string]interface{}{
																	"properties": map[string]interface{}{
																		"items": map[string]interface{}{
																			"items": map[string]interface{}{
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"type": "string",
																					},
																					"templateAs": map[string]interface{}{
																						"default": "Values",
																						"enum": []interface{}{
																							"Values",
																							"KeysAndValues",
																						},
																						"type": "string",
																					},
																				},
																				"required": []interface{}{
																					"key",
																				},
																				"type": "object",
																			},
																			"type": "array",
																		},
																		"name": map[string]interface{}{
																			"type": "string",
																		},
																	},
																	"required": []interface{}{
																		"items",
																		"name",
																	},
																	"type": "object",
																},
																"target": map[string]interface{}{
																	"default": "Data",
																	"enum": []interface{}{
																		"Data",
																		"Annotations",
																		"Labels",
																	},
																	"type": "string",
																},
															},
															"type": "object",
														},
														"type": "array",
													},
													"type": map[string]interface{}{
														"type": "string",
													},
												},
												"type": "object",
											},
										},
										"required": []interface{}{
											"secretStoreRefs",
											"selector",
										},
										"type": "object",
									},
									"status": map[string]interface{}{
										"description": "PushSecretStatus indicates the history of the status of PushSecret.",
										"properties": map[string]interface{}{
											"conditions": map[string]interface{}{
												"items": map[string]interface{}{
													"description": "PushSecretStatusCondition indicates the status of the PushSecret.",
													"properties": map[string]interface{}{
														"lastTransitionTime": map[string]interface{}{
															"format": "date-time",
															"type":   "string",
														},
														"message": map[string]interface{}{
															"type": "string",
														},
														"reason": map[string]interface{}{
															"type": "string",
														},
														"status": map[string]interface{}{
															"type": "string",
														},
														"type": map[string]interface{}{
															"description": "PushSecretConditionType indicates the condition of the PushSecret.",
															"type":        "string",
														},
													},
													"required": []interface{}{
														"status",
														"type",
													},
													"type": "object",
												},
												"type": "array",
											},
											"refreshTime": map[string]interface{}{
												"description": "refreshTime is the time and date the external secret was fetched and the target secret updated",
												"format":      "date-time",
												"nullable":    true,
												"type":        "string",
											},
											"syncedPushSecrets": map[string]interface{}{
												"additionalProperties": map[string]interface{}{
													"additionalProperties": map[string]interface{}{
														"properties": map[string]interface{}{
															"match": map[string]interface{}{
																"description": "Match a given Secret Key to be pushed to the provider.",
																"properties": map[string]interface{}{
																	"remoteRef": map[string]interface{}{
																		"description": "Remote Refs to push to providers.",
																		"properties": map[string]interface{}{
																			"property": map[string]interface{}{
																				"description": "Name of the property in the resulting secret",
																				"type":        "string",
																			},
																			"remoteKey": map[string]interface{}{
																				"description": "Name of the resulting provider secret.",
																				"type":        "string",
																			},
																		},
																		"required": []interface{}{
																			"remoteKey",
																		},
																		"type": "object",
																	},
																	"secretKey": map[string]interface{}{
																		"description": "Secret Key to be pushed",
																		"type":        "string",
																	},
																},
																"required": []interface{}{
																	"remoteRef",
																},
																"type": "object",
															},
															"metadata": map[string]interface{}{
																"description":                          "Metadata is metadata attached to the secret. The structure of metadata is provider specific, please look it up in the provider documentation.",
																"x-kubernetes-preserve-unknown-fields": true,
															},
														},
														"required": []interface{}{
															"match",
														},
														"type": "object",
													},
													"type": "object",
												},
												"description": "Synced Push Secrets for later deletion. Matches Secret Stores to PushSecretData that was stored to that secretStore.",
												"type":        "object",
											},
											"syncedResourceVersion": map[string]interface{}{
												"description": "SyncedResourceVersion keeps track of the last synced version.",
												"type":        "string",
											},
										},
										"type": "object",
									},
								},
								"type": "object",
							},
						},
						"served":  true,
						"storage": true,
						"subresources": map[string]interface{}{
							"status": map[string]interface{}{},
						},
					},
				},
				"conversion": map[string]interface{}{
					"strategy": "Webhook",
					"webhook": map[string]interface{}{
						"conversionReviewVersions": []interface{}{
							"v1",
						},
						"clientConfig": map[string]interface{}{
							"service": map[string]interface{}{
								"name":      "release-name-external-secrets-webhook",
								"namespace": "default",
								"path":      "/convert",
							},
						},
					},
				},
			},
		},
	}

	return mutate.MutateCRDPushsecretsExternalSecretsIo(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;create;update;patch;delete

// CreateCRDSecretstoresExternalSecretsIo creates the CustomResourceDefinition resource with name secretstores.external-secrets.io.
func CreateCRDSecretstoresExternalSecretsIo(
	parent *secretsv1alpha1.ExternalSecrets,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// Source: external-secrets/templates/crds/secretstore.yaml
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					"controller-gen.kubebuilder.io/version": "v0.13.0",
				},
				"name": "secretstores.external-secrets.io",
			},
			"spec": map[string]interface{}{
				"group": "external-secrets.io",
				"names": map[string]interface{}{
					"categories": []interface{}{
						"externalsecrets",
					},
					"kind":     "SecretStore",
					"listKind": "SecretStoreList",
					"plural":   "secretstores",
					"shortNames": []interface{}{
						"ss",
					},
					"singular": "secretstore",
				},
				"scope": "Namespaced",
				"versions": []interface{}{
					map[string]interface{}{
						"additionalPrinterColumns": []interface{}{
							map[string]interface{}{
								"jsonPath": ".metadata.creationTimestamp",
								"name":     "AGE",
								"type":     "date",
							},
							map[string]interface{}{
								"jsonPath": ".status.conditions[?(@.type==\"Ready\")].reason",
								"name":     "Status",
								"type":     "string",
							},
						},
						"deprecated": true,
						"name":       "v1alpha1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"description": "SecretStore represents a secure external location for storing secrets, which can be referenced as part of `storeRef` fields.",
								"properties": map[string]interface{}{
									"apiVersion": map[string]interface{}{
										"description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
										"type":        "string",
									},
									"kind": map[string]interface{}{
										"description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
										"type":        "string",
									},
									"metadata": map[string]interface{}{
										"type": "object",
									},
									"spec": map[string]interface{}{
										"description": "SecretStoreSpec defines the desired state of SecretStore.",
										"properties": map[string]interface{}{
											"controller": map[string]interface{}{
												"description": "Used to select the correct ESO controller (think: ingress.ingressClassName) The ESO controller is instantiated with a specific controller name and filters ES based on this property",
												"type":        "string",
											},
											"provider": map[string]interface{}{
												"description":   "Used to configure the provider. Only one provider may be set",
												"maxProperties": 1,
												"minProperties": 1,
												"properties": map[string]interface{}{
													"akeyless": map[string]interface{}{
														"description": "Akeyless configures this store to sync secrets using Akeyless Vault provider",
														"properties": map[string]interface{}{
															"akeylessGWApiURL": map[string]interface{}{
																"description": "Akeyless GW API Url from which the secrets to be fetched from.",
																"type":        "string",
															},
															"authSecretRef": map[string]interface{}{
																"description": "Auth configures how the operator authenticates with Akeyless.",
																"properties": map[string]interface{}{
																	"kubernetesAuth": map[string]interface{}{
																		"description": "Kubernetes authenticates with Akeyless by passing the ServiceAccount token stored in the named Secret resource.",
																		"properties": map[string]interface{}{
																			"accessID": map[string]interface{}{
																				"description": "the Akeyless Kubernetes auth-method access-id",
																				"type":        "string",
																			},
																			"k8sConfName": map[string]interface{}{
																				"description": "Kubernetes-auth configuration name in Akeyless-Gateway",
																				"type":        "string",
																			},
																			"secretRef": map[string]interface{}{
																				"description": "Optional secret field containing a Kubernetes ServiceAccount JWT used for authenticating with Akeyless. If a name is specified without a key, `token` is the default. If one is not specified, the one bound to the controller will be used.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"serviceAccountRef": map[string]interface{}{
																				"description": "Optional service account field containing the name of a kubernetes ServiceAccount. If the service account is specified, the service account secret token JWT will be used for authenticating with Akeyless. If the service account selector is not supplied, the secretRef will be used instead.",
																				"properties": map[string]interface{}{
																					"audiences": map[string]interface{}{
																						"description": "Audience specifies the `aud` claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																						"items": map[string]interface{}{
																							"type": "string",
																						},
																						"type": "array",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the ServiceAccount resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"required": []interface{}{
																					"name",
																				},
																				"type": "object",
																			},
																		},
																		"required": []interface{}{
																			"accessID",
																			"k8sConfName",
																		},
																		"type": "object",
																	},
																	"secretRef": map[string]interface{}{
																		"description": "Reference to a Secret that contains the details to authenticate with Akeyless.",
																		"properties": map[string]interface{}{
																			"accessID": map[string]interface{}{
																				"description": "The SecretAccessID is used for authentication",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"accessType": map[string]interface{}{
																				"description": "A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"accessTypeParam": map[string]interface{}{
																				"description": "A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
															"caBundle": map[string]interface{}{
																"description": "PEM/base64 encoded CA bundle used to validate Akeyless Gateway certificate. Only used if the AkeylessGWApiURL URL is using HTTPS protocol. If not set the system root certificates are used to validate the TLS connection.",
																"format":      "byte",
																"type":        "string",
															},
															"caProvider": map[string]interface{}{
																"description": "The provider for the CA bundle to use to validate Akeyless Gateway certificate.",
																"properties": map[string]interface{}{
																	"key": map[string]interface{}{
																		"description": "The key the value inside of the provider type to use, only used with \"Secret\" type",
																		"type":        "string",
																	},
																	"name": map[string]interface{}{
																		"description": "The name of the object located at the provider type.",
																		"type":        "string",
																	},
																	"namespace": map[string]interface{}{
																		"description": "The namespace the Provider type is in.",
																		"type":        "string",
																	},
																	"type": map[string]interface{}{
																		"description": "The type of provider to use such as \"Secret\", or \"ConfigMap\".",
																		"enum": []interface{}{
																			"Secret",
																			"ConfigMap",
																		},
																		"type": "string",
																	},
																},
																"required": []interface{}{
																	"name",
																	"type",
																},
																"type": "object",
															},
														},
														"required": []interface{}{
															"akeylessGWApiURL",
															"authSecretRef",
														},
														"type": "object",
													},
													"alibaba": map[string]interface{}{
														"description": "Alibaba configures this store to sync secrets using Alibaba Cloud provider",
														"properties": map[string]interface{}{
															"auth": map[string]interface{}{
																"description": "AlibabaAuth contains a secretRef for credentials.",
																"properties": map[string]interface{}{
																	"rrsa": map[string]interface{}{
																		"description": "Authenticate against Alibaba using RRSA.",
																		"properties": map[string]interface{}{
																			"oidcProviderArn": map[string]interface{}{
																				"type": "string",
																			},
																			"oidcTokenFilePath": map[string]interface{}{
																				"type": "string",
																			},
																			"roleArn": map[string]interface{}{
																				"type": "string",
																			},
																			"sessionName": map[string]interface{}{
																				"type": "string",
																			},
																		},
																		"required": []interface{}{
																			"oidcProviderArn",
																			"oidcTokenFilePath",
																			"roleArn",
																			"sessionName",
																		},
																		"type": "object",
																	},
																	"secretRef": map[string]interface{}{
																		"description": "AlibabaAuthSecretRef holds secret references for Alibaba credentials.",
																		"properties": map[string]interface{}{
																			"accessKeyIDSecretRef": map[string]interface{}{
																				"description": "The AccessKeyID is used for authentication",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"accessKeySecretSecretRef": map[string]interface{}{
																				"description": "The AccessKeySecret is used for authentication",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"required": []interface{}{
																			"accessKeyIDSecretRef",
																			"accessKeySecretSecretRef",
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
															"regionID": map[string]interface{}{
																"description": "Alibaba Region to be used for the provider",
																"type":        "string",
															},
														},
														"required": []interface{}{
															"auth",
															"regionID",
														},
														"type": "object",
													},
													"aws": map[string]interface{}{
														"description": "AWS configures this store to sync secrets using AWS Secret Manager provider",
														"properties": map[string]interface{}{
															"auth": map[string]interface{}{
																"description": "Auth defines the information necessary to authenticate against AWS if not set aws sdk will infer credentials from your environment see: https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-credentials",
																"properties": map[string]interface{}{
																	"jwt": map[string]interface{}{
																		"description": "Authenticate against AWS using service account tokens.",
																		"properties": map[string]interface{}{
																			"serviceAccountRef": map[string]interface{}{
																				"description": "A reference to a ServiceAccount resource.",
																				"properties": map[string]interface{}{
																					"audiences": map[string]interface{}{
																						"description": "Audience specifies the `aud` claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																						"items": map[string]interface{}{
																							"type": "string",
																						},
																						"type": "array",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the ServiceAccount resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"required": []interface{}{
																					"name",
																				},
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																	"secretRef": map[string]interface{}{
																		"description": "AWSAuthSecretRef holds secret references for AWS credentials both AccessKeyID and SecretAccessKey must be defined in order to properly authenticate.",
																		"properties": map[string]interface{}{
																			"accessKeyIDSecretRef": map[string]interface{}{
																				"description": "The AccessKeyID is used for authentication",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"secretAccessKeySecretRef": map[string]interface{}{
																				"description": "The SecretAccessKey is used for authentication",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
															"region": map[string]interface{}{
																"description": "AWS Region to be used for the provider",
																"type":        "string",
															},
															"role": map[string]interface{}{
																"description": "Role is a Role ARN which the SecretManager provider will assume",
																"type":        "string",
															},
															"service": map[string]interface{}{
																"description": "Service defines which service should be used to fetch the secrets",
																"enum": []interface{}{
																	"SecretsManager",
																	"ParameterStore",
																},
																"type": "string",
															},
														},
														"required": []interface{}{
															"region",
															"service",
														},
														"type": "object",
													},
													"azurekv": map[string]interface{}{
														"description": "AzureKV configures this store to sync secrets using Azure Key Vault provider",
														"properties": map[string]interface{}{
															"authSecretRef": map[string]interface{}{
																"description": "Auth configures how the operator authenticates with Azure. Required for ServicePrincipal auth type.",
																"properties": map[string]interface{}{
																	"clientId": map[string]interface{}{
																		"description": "The Azure clientId of the service principle used for authentication.",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																				"type":        "string",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the Secret resource being referred to.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																	"clientSecret": map[string]interface{}{
																		"description": "The Azure ClientSecret of the service principle used for authentication.",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																				"type":        "string",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the Secret resource being referred to.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
															"authType": map[string]interface{}{
																"default":     "ServicePrincipal",
																"description": "Auth type defines how to authenticate to the keyvault service. Valid values are: - \"ServicePrincipal\" (default): Using a service principal (tenantId, clientId, clientSecret) - \"ManagedIdentity\": Using Managed Identity assigned to the pod (see aad-pod-identity)",
																"enum": []interface{}{
																	"ServicePrincipal",
																	"ManagedIdentity",
																	"WorkloadIdentity",
																},
																"type": "string",
															},
															"identityId": map[string]interface{}{
																"description": "If multiple Managed Identity is assigned to the pod, you can select the one to be used",
																"type":        "string",
															},
															"serviceAccountRef": map[string]interface{}{
																"description": "ServiceAccountRef specified the service account that should be used when authenticating with WorkloadIdentity.",
																"properties": map[string]interface{}{
																	"audiences": map[string]interface{}{
																		"description": "Audience specifies the `aud` claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																		"items": map[string]interface{}{
																			"type": "string",
																		},
																		"type": "array",
																	},
																	"name": map[string]interface{}{
																		"description": "The name of the ServiceAccount resource being referred to.",
																		"type":        "string",
																	},
																	"namespace": map[string]interface{}{
																		"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																		"type":        "string",
																	},
																},
																"required": []interface{}{
																	"name",
																},
																"type": "object",
															},
															"tenantId": map[string]interface{}{
																"description": "TenantID configures the Azure Tenant to send requests to. Required for ServicePrincipal auth type.",
																"type":        "string",
															},
															"vaultUrl": map[string]interface{}{
																"description": "Vault Url from which the secrets to be fetched from.",
																"type":        "string",
															},
														},
														"required": []interface{}{
															"vaultUrl",
														},
														"type": "object",
													},
													"fake": map[string]interface{}{
														"description": "Fake configures a store with static key/value pairs",
														"properties": map[string]interface{}{
															"data": map[string]interface{}{
																"items": map[string]interface{}{
																	"properties": map[string]interface{}{
																		"key": map[string]interface{}{
																			"type": "string",
																		},
																		"value": map[string]interface{}{
																			"type": "string",
																		},
																		"valueMap": map[string]interface{}{
																			"additionalProperties": map[string]interface{}{
																				"type": "string",
																			},
																			"type": "object",
																		},
																		"version": map[string]interface{}{
																			"type": "string",
																		},
																	},
																	"required": []interface{}{
																		"key",
																	},
																	"type": "object",
																},
																"type": "array",
															},
														},
														"required": []interface{}{
															"data",
														},
														"type": "object",
													},
													"gcpsm": map[string]interface{}{
														"description": "GCPSM configures this store to sync secrets using Google Cloud Platform Secret Manager provider",
														"properties": map[string]interface{}{
															"auth": map[string]interface{}{
																"description": "Auth defines the information necessary to authenticate against GCP",
																"properties": map[string]interface{}{
																	"secretRef": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"secretAccessKeySecretRef": map[string]interface{}{
																				"description": "The SecretAccessKey is used for authentication",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																	"workloadIdentity": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"clusterLocation": map[string]interface{}{
																				"type": "string",
																			},
																			"clusterName": map[string]interface{}{
																				"type": "string",
																			},
																			"clusterProjectID": map[string]interface{}{
																				"type": "string",
																			},
																			"serviceAccountRef": map[string]interface{}{
																				"description": "A reference to a ServiceAccount resource.",
																				"properties": map[string]interface{}{
																					"audiences": map[string]interface{}{
																						"description": "Audience specifies the `aud` claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																						"items": map[string]interface{}{
																							"type": "string",
																						},
																						"type": "array",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the ServiceAccount resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"required": []interface{}{
																					"name",
																				},
																				"type": "object",
																			},
																		},
																		"required": []interface{}{
																			"clusterLocation",
																			"clusterName",
																			"serviceAccountRef",
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
															"projectID": map[string]interface{}{
																"description": "ProjectID project where secret is located",
																"type":        "string",
															},
														},
														"type": "object",
													},
													"gitlab": map[string]interface{}{
														"description": "GitLab configures this store to sync secrets using GitLab Variables provider",
														"properties": map[string]interface{}{
															"auth": map[string]interface{}{
																"description": "Auth configures how secret-manager authenticates with a GitLab instance.",
																"properties": map[string]interface{}{
																	"SecretRef": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"accessToken": map[string]interface{}{
																				"description": "AccessToken is used for authentication.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																},
																"required": []interface{}{
																	"SecretRef",
																},
																"type": "object",
															},
															"projectID": map[string]interface{}{
																"description": "ProjectID specifies a project where secrets are located.",
																"type":        "string",
															},
															"url": map[string]interface{}{
																"description": "URL configures the GitLab instance URL. Defaults to https://gitlab.com/.",
																"type":        "string",
															},
														},
														"required": []interface{}{
															"auth",
														},
														"type": "object",
													},
													"ibm": map[string]interface{}{
														"description": "IBM configures this store to sync secrets using IBM Cloud provider",
														"properties": map[string]interface{}{
															"auth": map[string]interface{}{
																"description": "Auth configures how secret-manager authenticates with the IBM secrets manager.",
																"properties": map[string]interface{}{
																	"secretRef": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"secretApiKeySecretRef": map[string]interface{}{
																				"description": "The SecretAccessKey is used for authentication",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																},
																"required": []interface{}{
																	"secretRef",
																},
																"type": "object",
															},
															"serviceUrl": map[string]interface{}{
																"description": "ServiceURL is the Endpoint URL that is specific to the Secrets Manager service instance",
																"type":        "string",
															},
														},
														"required": []interface{}{
															"auth",
														},
														"type": "object",
													},
													"kubernetes": map[string]interface{}{
														"description": "Kubernetes configures this store to sync secrets using a Kubernetes cluster provider",
														"properties": map[string]interface{}{
															"auth": map[string]interface{}{
																"description":   "Auth configures how secret-manager authenticates with a Kubernetes instance.",
																"maxProperties": 1,
																"minProperties": 1,
																"properties": map[string]interface{}{
																	"cert": map[string]interface{}{
																		"description": "has both clientCert and clientKey as secretKeySelector",
																		"properties": map[string]interface{}{
																			"clientCert": map[string]interface{}{
																				"description": "A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"clientKey": map[string]interface{}{
																				"description": "A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																	"serviceAccount": map[string]interface{}{
																		"description": "points to a service account that should be used for authentication",
																		"properties": map[string]interface{}{
																			"serviceAccount": map[string]interface{}{
																				"description": "A reference to a ServiceAccount resource.",
																				"properties": map[string]interface{}{
																					"audiences": map[string]interface{}{
																						"description": "Audience specifies the `aud` claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																						"items": map[string]interface{}{
																							"type": "string",
																						},
																						"type": "array",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the ServiceAccount resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"required": []interface{}{
																					"name",
																				},
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																	"token": map[string]interface{}{
																		"description": "use static token to authenticate with",
																		"properties": map[string]interface{}{
																			"bearerToken": map[string]interface{}{
																				"description": "A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
															"remoteNamespace": map[string]interface{}{
																"default":     "default",
																"description": "Remote namespace to fetch the secrets from",
																"type":        "string",
															},
															"server": map[string]interface{}{
																"description": "configures the Kubernetes server Address.",
																"properties": map[string]interface{}{
																	"caBundle": map[string]interface{}{
																		"description": "CABundle is a base64-encoded CA certificate",
																		"format":      "byte",
																		"type":        "string",
																	},
																	"caProvider": map[string]interface{}{
																		"description": "see: https://external-secrets.io/v0.4.1/spec/#external-secrets.io/v1alpha1.CAProvider",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "The key the value inside of the provider type to use, only used with \"Secret\" type",
																				"type":        "string",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the object located at the provider type.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "The namespace the Provider type is in.",
																				"type":        "string",
																			},
																			"type": map[string]interface{}{
																				"description": "The type of provider to use such as \"Secret\", or \"ConfigMap\".",
																				"enum": []interface{}{
																					"Secret",
																					"ConfigMap",
																				},
																				"type": "string",
																			},
																		},
																		"required": []interface{}{
																			"name",
																			"type",
																		},
																		"type": "object",
																	},
																	"url": map[string]interface{}{
																		"default":     "kubernetes.default",
																		"description": "configures the Kubernetes server Address.",
																		"type":        "string",
																	},
																},
																"type": "object",
															},
														},
														"required": []interface{}{
															"auth",
														},
														"type": "object",
													},
													"oracle": map[string]interface{}{
														"description": "Oracle configures this store to sync secrets using Oracle Vault provider",
														"properties": map[string]interface{}{
															"auth": map[string]interface{}{
																"description": "Auth configures how secret-manager authenticates with the Oracle Vault. If empty, instance principal is used. Optionally, the authenticating principal type and/or user data may be supplied for the use of workload identity and user principal.",
																"properties": map[string]interface{}{
																	"secretRef": map[string]interface{}{
																		"description": "SecretRef to pass through sensitive information.",
																		"properties": map[string]interface{}{
																			"fingerprint": map[string]interface{}{
																				"description": "Fingerprint is the fingerprint of the API private key.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"privatekey": map[string]interface{}{
																				"description": "PrivateKey is the user's API Signing Key in PEM format, used for authentication.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"required": []interface{}{
																			"fingerprint",
																			"privatekey",
																		},
																		"type": "object",
																	},
																	"tenancy": map[string]interface{}{
																		"description": "Tenancy is the tenancy OCID where user is located.",
																		"type":        "string",
																	},
																	"user": map[string]interface{}{
																		"description": "User is an access OCID specific to the account.",
																		"type":        "string",
																	},
																},
																"required": []interface{}{
																	"secretRef",
																	"tenancy",
																	"user",
																},
																"type": "object",
															},
															"compartment": map[string]interface{}{
																"description": "Compartment is the vault compartment OCID. Required for PushSecret",
																"type":        "string",
															},
															"encryptionKey": map[string]interface{}{
																"description": "EncryptionKey is the OCID of the encryption key within the vault. Required for PushSecret",
																"type":        "string",
															},
															"principalType": map[string]interface{}{
																"description": "The type of principal to use for authentication. If left blank, the Auth struct will determine the principal type. This optional field must be specified if using workload identity.",
																"enum": []interface{}{
																	"",
																	"UserPrincipal",
																	"InstancePrincipal",
																	"Workload",
																},
																"type": "string",
															},
															"region": map[string]interface{}{
																"description": "Region is the region where vault is located.",
																"type":        "string",
															},
															"serviceAccountRef": map[string]interface{}{
																"description": "ServiceAccountRef specified the service account that should be used when authenticating with WorkloadIdentity.",
																"properties": map[string]interface{}{
																	"audiences": map[string]interface{}{
																		"description": "Audience specifies the `aud` claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																		"items": map[string]interface{}{
																			"type": "string",
																		},
																		"type": "array",
																	},
																	"name": map[string]interface{}{
																		"description": "The name of the ServiceAccount resource being referred to.",
																		"type":        "string",
																	},
																	"namespace": map[string]interface{}{
																		"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																		"type":        "string",
																	},
																},
																"required": []interface{}{
																	"name",
																},
																"type": "object",
															},
															"vault": map[string]interface{}{
																"description": "Vault is the vault's OCID of the specific vault where secret is located.",
																"type":        "string",
															},
														},
														"required": []interface{}{
															"region",
															"vault",
														},
														"type": "object",
													},
													"vault": map[string]interface{}{
														"description": "Vault configures this store to sync secrets using Hashi provider",
														"properties": map[string]interface{}{
															"auth": map[string]interface{}{
																"description": "Auth configures how secret-manager authenticates with the Vault server.",
																"properties": map[string]interface{}{
																	"appRole": map[string]interface{}{
																		"description": "AppRole authenticates with Vault using the App Role auth mechanism, with the role and secret stored in a Kubernetes Secret resource.",
																		"properties": map[string]interface{}{
																			"path": map[string]interface{}{
																				"default":     "approle",
																				"description": "Path where the App Role authentication backend is mounted in Vault, e.g: \"approle\"",
																				"type":        "string",
																			},
																			"roleId": map[string]interface{}{
																				"description": "RoleID configured in the App Role authentication backend when setting up the authentication backend in Vault.",
																				"type":        "string",
																			},
																			"secretRef": map[string]interface{}{
																				"description": "Reference to a key in a Secret that contains the App Role secret used to authenticate with Vault. The `key` field must be specified and denotes which entry within the Secret resource is used as the app role secret.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"required": []interface{}{
																			"path",
																			"roleId",
																			"secretRef",
																		},
																		"type": "object",
																	},
																	"cert": map[string]interface{}{
																		"description": "Cert authenticates with TLS Certificates by passing client certificate, private key and ca certificate Cert authentication method",
																		"properties": map[string]interface{}{
																			"clientCert": map[string]interface{}{
																				"description": "ClientCert is a certificate to authenticate using the Cert Vault authentication method",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"secretRef": map[string]interface{}{
																				"description": "SecretRef to a key in a Secret resource containing client private key to authenticate with Vault using the Cert authentication method",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																	"jwt": map[string]interface{}{
																		"description": "Jwt authenticates with Vault by passing role and JWT token using the JWT/OIDC authentication method",
																		"properties": map[string]interface{}{
																			"kubernetesServiceAccountToken": map[string]interface{}{
																				"description": "Optional ServiceAccountToken specifies the Kubernetes service account for which to request a token for with the `TokenRequest` API.",
																				"properties": map[string]interface{}{
																					"audiences": map[string]interface{}{
																						"description": "Optional audiences field that will be used to request a temporary Kubernetes service account token for the service account referenced by `serviceAccountRef`. Defaults to a single audience `vault` it not specified.",
																						"items": map[string]interface{}{
																							"type": "string",
																						},
																						"type": "array",
																					},
																					"expirationSeconds": map[string]interface{}{
																						"description": "Optional expiration time in seconds that will be used to request a temporary Kubernetes service account token for the service account referenced by `serviceAccountRef`. Defaults to 10 minutes.",
																						"format":      "int64",
																						"type":        "integer",
																					},
																					"serviceAccountRef": map[string]interface{}{
																						"description": "Service account field containing the name of a kubernetes ServiceAccount.",
																						"properties": map[string]interface{}{
																							"audiences": map[string]interface{}{
																								"description": "Audience specifies the `aud` claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																								"items": map[string]interface{}{
																									"type": "string",
																								},
																								"type": "array",
																							},
																							"name": map[string]interface{}{
																								"description": "The name of the ServiceAccount resource being referred to.",
																								"type":        "string",
																							},
																							"namespace": map[string]interface{}{
																								"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																								"type":        "string",
																							},
																						},
																						"required": []interface{}{
																							"name",
																						},
																						"type": "object",
																					},
																				},
																				"required": []interface{}{
																					"serviceAccountRef",
																				},
																				"type": "object",
																			},
																			"path": map[string]interface{}{
																				"default":     "jwt",
																				"description": "Path where the JWT authentication backend is mounted in Vault, e.g: \"jwt\"",
																				"type":        "string",
																			},
																			"role": map[string]interface{}{
																				"description": "Role is a JWT role to authenticate using the JWT/OIDC Vault authentication method",
																				"type":        "string",
																			},
																			"secretRef": map[string]interface{}{
																				"description": "Optional SecretRef that refers to a key in a Secret resource containing JWT token to authenticate with Vault using the JWT/OIDC authentication method.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"required": []interface{}{
																			"path",
																		},
																		"type": "object",
																	},
																	"kubernetes": map[string]interface{}{
																		"description": "Kubernetes authenticates with Vault by passing the ServiceAccount token stored in the named Secret resource to the Vault server.",
																		"properties": map[string]interface{}{
																			"mountPath": map[string]interface{}{
																				"default":     "kubernetes",
																				"description": "Path where the Kubernetes authentication backend is mounted in Vault, e.g: \"kubernetes\"",
																				"type":        "string",
																			},
																			"role": map[string]interface{}{
																				"description": "A required field containing the Vault Role to assume. A Role binds a Kubernetes ServiceAccount with a set of Vault policies.",
																				"type":        "string",
																			},
																			"secretRef": map[string]interface{}{
																				"description": "Optional secret field containing a Kubernetes ServiceAccount JWT used for authenticating with Vault. If a name is specified without a key, `token` is the default. If one is not specified, the one bound to the controller will be used.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"serviceAccountRef": map[string]interface{}{
																				"description": "Optional service account field containing the name of a kubernetes ServiceAccount. If the service account is specified, the service account secret token JWT will be used for authenticating with Vault. If the service account selector is not supplied, the secretRef will be used instead.",
																				"properties": map[string]interface{}{
																					"audiences": map[string]interface{}{
																						"description": "Audience specifies the `aud` claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																						"items": map[string]interface{}{
																							"type": "string",
																						},
																						"type": "array",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the ServiceAccount resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"required": []interface{}{
																					"name",
																				},
																				"type": "object",
																			},
																		},
																		"required": []interface{}{
																			"mountPath",
																			"role",
																		},
																		"type": "object",
																	},
																	"ldap": map[string]interface{}{
																		"description": "Ldap authenticates with Vault by passing username/password pair using the LDAP authentication method",
																		"properties": map[string]interface{}{
																			"path": map[string]interface{}{
																				"default":     "ldap",
																				"description": "Path where the LDAP authentication backend is mounted in Vault, e.g: \"ldap\"",
																				"type":        "string",
																			},
																			"secretRef": map[string]interface{}{
																				"description": "SecretRef to a key in a Secret resource containing password for the LDAP user used to authenticate with Vault using the LDAP authentication method",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"username": map[string]interface{}{
																				"description": "Username is a LDAP user name used to authenticate using the LDAP Vault authentication method",
																				"type":        "string",
																			},
																		},
																		"required": []interface{}{
																			"path",
																			"username",
																		},
																		"type": "object",
																	},
																	"tokenSecretRef": map[string]interface{}{
																		"description": "TokenSecretRef authenticates with Vault by presenting a token.",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																				"type":        "string",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the Secret resource being referred to.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
															"caBundle": map[string]interface{}{
																"description": "PEM encoded CA bundle used to validate Vault server certificate. Only used if the Server URL is using HTTPS protocol. This parameter is ignored for plain HTTP protocol connection. If not set the system root certificates are used to validate the TLS connection.",
																"format":      "byte",
																"type":        "string",
															},
															"caProvider": map[string]interface{}{
																"description": "The provider for the CA bundle to use to validate Vault server certificate.",
																"properties": map[string]interface{}{
																	"key": map[string]interface{}{
																		"description": "The key the value inside of the provider type to use, only used with \"Secret\" type",
																		"type":        "string",
																	},
																	"name": map[string]interface{}{
																		"description": "The name of the object located at the provider type.",
																		"type":        "string",
																	},
																	"namespace": map[string]interface{}{
																		"description": "The namespace the Provider type is in.",
																		"type":        "string",
																	},
																	"type": map[string]interface{}{
																		"description": "The type of provider to use such as \"Secret\", or \"ConfigMap\".",
																		"enum": []interface{}{
																			"Secret",
																			"ConfigMap",
																		},
																		"type": "string",
																	},
																},
																"required": []interface{}{
																	"name",
																	"type",
																},
																"type": "object",
															},
															"forwardInconsistent": map[string]interface{}{
																"description": "ForwardInconsistent tells Vault to forward read-after-write requests to the Vault leader instead of simply retrying within a loop. This can increase performance if the option is enabled serverside. https://www.vaultproject.io/docs/configuration/replication#allow_forwarding_via_header",
																"type":        "boolean",
															},
															"namespace": map[string]interface{}{
																"description": "Name of the vault namespace. Namespaces is a set of features within Vault Enterprise that allows Vault environments to support Secure Multi-tenancy. e.g: \"ns1\". More about namespaces can be found here https://www.vaultproject.io/docs/enterprise/namespaces",
																"type":        "string",
															},
															"path": map[string]interface{}{
																"description": "Path is the mount path of the Vault KV backend endpoint, e.g: \"secret\". The v2 KV secret engine version specific \"/data\" path suffix for fetching secrets from Vault is optional and will be appended if not present in specified path.",
																"type":        "string",
															},
															"readYourWrites": map[string]interface{}{
																"description": "ReadYourWrites ensures isolated read-after-write semantics by providing discovered cluster replication states in each request. More information about eventual consistency in Vault can be found here https://www.vaultproject.io/docs/enterprise/consistency",
																"type":        "boolean",
															},
															"server": map[string]interface{}{
																"description": "Server is the connection address for the Vault server, e.g: \"https://vault.example.com:8200\".",
																"type":        "string",
															},
															"version": map[string]interface{}{
																"default":     "v2",
																"description": "Version is the Vault KV secret engine version. This can be either \"v1\" or \"v2\". Version defaults to \"v2\".",
																"enum": []interface{}{
																	"v1",
																	"v2",
																},
																"type": "string",
															},
														},
														"required": []interface{}{
															"auth",
															"server",
														},
														"type": "object",
													},
													"webhook": map[string]interface{}{
														"description": "Webhook configures this store to sync secrets using a generic templated webhook",
														"properties": map[string]interface{}{
															"body": map[string]interface{}{
																"description": "Body",
																"type":        "string",
															},
															"caBundle": map[string]interface{}{
																"description": "PEM encoded CA bundle used to validate webhook server certificate. Only used if the Server URL is using HTTPS protocol. This parameter is ignored for plain HTTP protocol connection. If not set the system root certificates are used to validate the TLS connection.",
																"format":      "byte",
																"type":        "string",
															},
															"caProvider": map[string]interface{}{
																"description": "The provider for the CA bundle to use to validate webhook server certificate.",
																"properties": map[string]interface{}{
																	"key": map[string]interface{}{
																		"description": "The key the value inside of the provider type to use, only used with \"Secret\" type",
																		"type":        "string",
																	},
																	"name": map[string]interface{}{
																		"description": "The name of the object located at the provider type.",
																		"type":        "string",
																	},
																	"namespace": map[string]interface{}{
																		"description": "The namespace the Provider type is in.",
																		"type":        "string",
																	},
																	"type": map[string]interface{}{
																		"description": "The type of provider to use such as \"Secret\", or \"ConfigMap\".",
																		"enum": []interface{}{
																			"Secret",
																			"ConfigMap",
																		},
																		"type": "string",
																	},
																},
																"required": []interface{}{
																	"name",
																	"type",
																},
																"type": "object",
															},
															"headers": map[string]interface{}{
																"additionalProperties": map[string]interface{}{
																	"type": "string",
																},
																"description": "Headers",
																"type":        "object",
															},
															"method": map[string]interface{}{
																"description": "Webhook Method",
																"type":        "string",
															},
															"result": map[string]interface{}{
																"description": "Result formatting",
																"properties": map[string]interface{}{
																	"jsonPath": map[string]interface{}{
																		"description": "Json path of return value",
																		"type":        "string",
																	},
																},
																"type": "object",
															},
															"secrets": map[string]interface{}{
																"description": "Secrets to fill in templates These secrets will be passed to the templating function as key value pairs under the given name",
																"items": map[string]interface{}{
																	"properties": map[string]interface{}{
																		"name": map[string]interface{}{
																			"description": "Name of this secret in templates",
																			"type":        "string",
																		},
																		"secretRef": map[string]interface{}{
																			"description": "Secret ref to fill in credentials",
																			"properties": map[string]interface{}{
																				"key": map[string]interface{}{
																					"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																					"type":        "string",
																				},
																				"name": map[string]interface{}{
																					"description": "The name of the Secret resource being referred to.",
																					"type":        "string",
																				},
																				"namespace": map[string]interface{}{
																					"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																					"type":        "string",
																				},
																			},
																			"type": "object",
																		},
																	},
																	"required": []interface{}{
																		"name",
																		"secretRef",
																	},
																	"type": "object",
																},
																"type": "array",
															},
															"timeout": map[string]interface{}{
																"description": "Timeout",
																"type":        "string",
															},
															"url": map[string]interface{}{
																"description": "Webhook url to call",
																"type":        "string",
															},
														},
														"required": []interface{}{
															"result",
															"url",
														},
														"type": "object",
													},
													"yandexlockbox": map[string]interface{}{
														"description": "YandexLockbox configures this store to sync secrets using Yandex Lockbox provider",
														"properties": map[string]interface{}{
															"apiEndpoint": map[string]interface{}{
																"description": "Yandex.Cloud API endpoint (e.g. 'api.cloud.yandex.net:443')",
																"type":        "string",
															},
															"auth": map[string]interface{}{
																"description": "Auth defines the information necessary to authenticate against Yandex Lockbox",
																"properties": map[string]interface{}{
																	"authorizedKeySecretRef": map[string]interface{}{
																		"description": "The authorized key used for authentication",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																				"type":        "string",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the Secret resource being referred to.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
															"caProvider": map[string]interface{}{
																"description": "The provider for the CA bundle to use to validate Yandex.Cloud server certificate.",
																"properties": map[string]interface{}{
																	"certSecretRef": map[string]interface{}{
																		"description": "A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																				"type":        "string",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the Secret resource being referred to.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
														},
														"required": []interface{}{
															"auth",
														},
														"type": "object",
													},
												},
												"type": "object",
											},
											"retrySettings": map[string]interface{}{
												"description": "Used to configure http retries if failed",
												"properties": map[string]interface{}{
													"maxRetries": map[string]interface{}{
														"format": "int32",
														"type":   "integer",
													},
													"retryInterval": map[string]interface{}{
														"type": "string",
													},
												},
												"type": "object",
											},
										},
										"required": []interface{}{
											"provider",
										},
										"type": "object",
									},
									"status": map[string]interface{}{
										"description": "SecretStoreStatus defines the observed state of the SecretStore.",
										"properties": map[string]interface{}{
											"conditions": map[string]interface{}{
												"items": map[string]interface{}{
													"properties": map[string]interface{}{
														"lastTransitionTime": map[string]interface{}{
															"format": "date-time",
															"type":   "string",
														},
														"message": map[string]interface{}{
															"type": "string",
														},
														"reason": map[string]interface{}{
															"type": "string",
														},
														"status": map[string]interface{}{
															"type": "string",
														},
														"type": map[string]interface{}{
															"type": "string",
														},
													},
													"required": []interface{}{
														"status",
														"type",
													},
													"type": "object",
												},
												"type": "array",
											},
										},
										"type": "object",
									},
								},
								"type": "object",
							},
						},
						"served":  true,
						"storage": false,
						"subresources": map[string]interface{}{
							"status": map[string]interface{}{},
						},
					},
					map[string]interface{}{
						"additionalPrinterColumns": []interface{}{
							map[string]interface{}{
								"jsonPath": ".metadata.creationTimestamp",
								"name":     "AGE",
								"type":     "date",
							},
							map[string]interface{}{
								"jsonPath": ".status.conditions[?(@.type==\"Ready\")].reason",
								"name":     "Status",
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".status.capabilities",
								"name":     "Capabilities",
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".status.conditions[?(@.type==\"Ready\")].status",
								"name":     "Ready",
								"type":     "string",
							},
						},
						"name": "v1beta1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"description": "SecretStore represents a secure external location for storing secrets, which can be referenced as part of `storeRef` fields.",
								"properties": map[string]interface{}{
									"apiVersion": map[string]interface{}{
										"description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
										"type":        "string",
									},
									"kind": map[string]interface{}{
										"description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
										"type":        "string",
									},
									"metadata": map[string]interface{}{
										"type": "object",
									},
									"spec": map[string]interface{}{
										"description": "SecretStoreSpec defines the desired state of SecretStore.",
										"properties": map[string]interface{}{
											"conditions": map[string]interface{}{
												"description": "Used to constraint a ClusterSecretStore to specific namespaces. Relevant only to ClusterSecretStore",
												"items": map[string]interface{}{
													"description": "ClusterSecretStoreCondition describes a condition by which to choose namespaces to process ExternalSecrets in for a ClusterSecretStore instance.",
													"properties": map[string]interface{}{
														"namespaceSelector": map[string]interface{}{
															"description": "Choose namespace using a labelSelector",
															"properties": map[string]interface{}{
																"matchExpressions": map[string]interface{}{
																	"description": "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																	"items": map[string]interface{}{
																		"description": "A label selector requirement is a selector that contains values, a key, and an operator that relates the key and values.",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "key is the label key that the selector applies to.",
																				"type":        "string",
																			},
																			"operator": map[string]interface{}{
																				"description": "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																				"type":        "string",
																			},
																			"values": map[string]interface{}{
																				"description": "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																				"items": map[string]interface{}{
																					"type": "string",
																				},
																				"type": "array",
																			},
																		},
																		"required": []interface{}{
																			"key",
																			"operator",
																		},
																		"type": "object",
																	},
																	"type": "array",
																},
																"matchLabels": map[string]interface{}{
																	"additionalProperties": map[string]interface{}{
																		"type": "string",
																	},
																	"description": "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is \"key\", the operator is \"In\", and the values array contains only \"value\". The requirements are ANDed.",
																	"type":        "object",
																},
															},
															"type":                  "object",
															"x-kubernetes-map-type": "atomic",
														},
														"namespaces": map[string]interface{}{
															"description": "Choose namespaces by name",
															"items": map[string]interface{}{
																"type": "string",
															},
															"type": "array",
														},
													},
													"type": "object",
												},
												"type": "array",
											},
											"controller": map[string]interface{}{
												"description": "Used to select the correct ESO controller (think: ingress.ingressClassName) The ESO controller is instantiated with a specific controller name and filters ES based on this property",
												"type":        "string",
											},
											"provider": map[string]interface{}{
												"description":   "Used to configure the provider. Only one provider may be set",
												"maxProperties": 1,
												"minProperties": 1,
												"properties": map[string]interface{}{
													"akeyless": map[string]interface{}{
														"description": "Akeyless configures this store to sync secrets using Akeyless Vault provider",
														"properties": map[string]interface{}{
															"akeylessGWApiURL": map[string]interface{}{
																"description": "Akeyless GW API Url from which the secrets to be fetched from.",
																"type":        "string",
															},
															"authSecretRef": map[string]interface{}{
																"description": "Auth configures how the operator authenticates with Akeyless.",
																"properties": map[string]interface{}{
																	"kubernetesAuth": map[string]interface{}{
																		"description": "Kubernetes authenticates with Akeyless by passing the ServiceAccount token stored in the named Secret resource.",
																		"properties": map[string]interface{}{
																			"accessID": map[string]interface{}{
																				"description": "the Akeyless Kubernetes auth-method access-id",
																				"type":        "string",
																			},
																			"k8sConfName": map[string]interface{}{
																				"description": "Kubernetes-auth configuration name in Akeyless-Gateway",
																				"type":        "string",
																			},
																			"secretRef": map[string]interface{}{
																				"description": "Optional secret field containing a Kubernetes ServiceAccount JWT used for authenticating with Akeyless. If a name is specified without a key, `token` is the default. If one is not specified, the one bound to the controller will be used.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"serviceAccountRef": map[string]interface{}{
																				"description": "Optional service account field containing the name of a kubernetes ServiceAccount. If the service account is specified, the service account secret token JWT will be used for authenticating with Akeyless. If the service account selector is not supplied, the secretRef will be used instead.",
																				"properties": map[string]interface{}{
																					"audiences": map[string]interface{}{
																						"description": "Audience specifies the `aud` claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																						"items": map[string]interface{}{
																							"type": "string",
																						},
																						"type": "array",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the ServiceAccount resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"required": []interface{}{
																					"name",
																				},
																				"type": "object",
																			},
																		},
																		"required": []interface{}{
																			"accessID",
																			"k8sConfName",
																		},
																		"type": "object",
																	},
																	"secretRef": map[string]interface{}{
																		"description": "Reference to a Secret that contains the details to authenticate with Akeyless.",
																		"properties": map[string]interface{}{
																			"accessID": map[string]interface{}{
																				"description": "The SecretAccessID is used for authentication",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"accessType": map[string]interface{}{
																				"description": "A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"accessTypeParam": map[string]interface{}{
																				"description": "A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
															"caBundle": map[string]interface{}{
																"description": "PEM/base64 encoded CA bundle used to validate Akeyless Gateway certificate. Only used if the AkeylessGWApiURL URL is using HTTPS protocol. If not set the system root certificates are used to validate the TLS connection.",
																"format":      "byte",
																"type":        "string",
															},
															"caProvider": map[string]interface{}{
																"description": "The provider for the CA bundle to use to validate Akeyless Gateway certificate.",
																"properties": map[string]interface{}{
																	"key": map[string]interface{}{
																		"description": "The key where the CA certificate can be found in the Secret or ConfigMap.",
																		"type":        "string",
																	},
																	"name": map[string]interface{}{
																		"description": "The name of the object located at the provider type.",
																		"type":        "string",
																	},
																	"namespace": map[string]interface{}{
																		"description": "The namespace the Provider type is in. Can only be defined when used in a ClusterSecretStore.",
																		"type":        "string",
																	},
																	"type": map[string]interface{}{
																		"description": "The type of provider to use such as \"Secret\", or \"ConfigMap\".",
																		"enum": []interface{}{
																			"Secret",
																			"ConfigMap",
																		},
																		"type": "string",
																	},
																},
																"required": []interface{}{
																	"name",
																	"type",
																},
																"type": "object",
															},
														},
														"required": []interface{}{
															"akeylessGWApiURL",
															"authSecretRef",
														},
														"type": "object",
													},
													"alibaba": map[string]interface{}{
														"description": "Alibaba configures this store to sync secrets using Alibaba Cloud provider",
														"properties": map[string]interface{}{
															"auth": map[string]interface{}{
																"description": "AlibabaAuth contains a secretRef for credentials.",
																"properties": map[string]interface{}{
																	"rrsa": map[string]interface{}{
																		"description": "Authenticate against Alibaba using RRSA.",
																		"properties": map[string]interface{}{
																			"oidcProviderArn": map[string]interface{}{
																				"type": "string",
																			},
																			"oidcTokenFilePath": map[string]interface{}{
																				"type": "string",
																			},
																			"roleArn": map[string]interface{}{
																				"type": "string",
																			},
																			"sessionName": map[string]interface{}{
																				"type": "string",
																			},
																		},
																		"required": []interface{}{
																			"oidcProviderArn",
																			"oidcTokenFilePath",
																			"roleArn",
																			"sessionName",
																		},
																		"type": "object",
																	},
																	"secretRef": map[string]interface{}{
																		"description": "AlibabaAuthSecretRef holds secret references for Alibaba credentials.",
																		"properties": map[string]interface{}{
																			"accessKeyIDSecretRef": map[string]interface{}{
																				"description": "The AccessKeyID is used for authentication",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"accessKeySecretSecretRef": map[string]interface{}{
																				"description": "The AccessKeySecret is used for authentication",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"required": []interface{}{
																			"accessKeyIDSecretRef",
																			"accessKeySecretSecretRef",
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
															"regionID": map[string]interface{}{
																"description": "Alibaba Region to be used for the provider",
																"type":        "string",
															},
														},
														"required": []interface{}{
															"auth",
															"regionID",
														},
														"type": "object",
													},
													"aws": map[string]interface{}{
														"description": "AWS configures this store to sync secrets using AWS Secret Manager provider",
														"properties": map[string]interface{}{
															"additionalRoles": map[string]interface{}{
																"description": "AdditionalRoles is a chained list of Role ARNs which the provider will sequentially assume before assuming the Role",
																"items": map[string]interface{}{
																	"type": "string",
																},
																"type": "array",
															},
															"auth": map[string]interface{}{
																"description": "Auth defines the information necessary to authenticate against AWS if not set aws sdk will infer credentials from your environment see: https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-credentials",
																"properties": map[string]interface{}{
																	"jwt": map[string]interface{}{
																		"description": "Authenticate against AWS using service account tokens.",
																		"properties": map[string]interface{}{
																			"serviceAccountRef": map[string]interface{}{
																				"description": "A reference to a ServiceAccount resource.",
																				"properties": map[string]interface{}{
																					"audiences": map[string]interface{}{
																						"description": "Audience specifies the `aud` claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																						"items": map[string]interface{}{
																							"type": "string",
																						},
																						"type": "array",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the ServiceAccount resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"required": []interface{}{
																					"name",
																				},
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																	"secretRef": map[string]interface{}{
																		"description": "AWSAuthSecretRef holds secret references for AWS credentials both AccessKeyID and SecretAccessKey must be defined in order to properly authenticate.",
																		"properties": map[string]interface{}{
																			"accessKeyIDSecretRef": map[string]interface{}{
																				"description": "The AccessKeyID is used for authentication",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"secretAccessKeySecretRef": map[string]interface{}{
																				"description": "The SecretAccessKey is used for authentication",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"sessionTokenSecretRef": map[string]interface{}{
																				"description": "The SessionToken used for authentication This must be defined if AccessKeyID and SecretAccessKey are temporary credentials see: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_use-resources.html",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
															"externalID": map[string]interface{}{
																"description": "AWS External ID set on assumed IAM roles",
																"type":        "string",
															},
															"region": map[string]interface{}{
																"description": "AWS Region to be used for the provider",
																"type":        "string",
															},
															"role": map[string]interface{}{
																"description": "Role is a Role ARN which the provider will assume",
																"type":        "string",
															},
															"secretsManager": map[string]interface{}{
																"description": "SecretsManager defines how the provider behaves when interacting with AWS SecretsManager",
																"properties": map[string]interface{}{
																	"forceDeleteWithoutRecovery": map[string]interface{}{
																		"description": "Specifies whether to delete the secret without any recovery window. You can't use both this parameter and RecoveryWindowInDays in the same call. If you don't use either, then by default Secrets Manager uses a 30 day recovery window. see: https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_DeleteSecret.html#SecretsManager-DeleteSecret-request-ForceDeleteWithoutRecovery",
																		"type":        "boolean",
																	},
																	"recoveryWindowInDays": map[string]interface{}{
																		"description": "The number of days from 7 to 30 that Secrets Manager waits before permanently deleting the secret. You can't use both this parameter and ForceDeleteWithoutRecovery in the same call. If you don't use either, then by default Secrets Manager uses a 30 day recovery window. see: https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_DeleteSecret.html#SecretsManager-DeleteSecret-request-RecoveryWindowInDays",
																		"format":      "int64",
																		"type":        "integer",
																	},
																},
																"type": "object",
															},
															"service": map[string]interface{}{
																"description": "Service defines which service should be used to fetch the secrets",
																"enum": []interface{}{
																	"SecretsManager",
																	"ParameterStore",
																},
																"type": "string",
															},
															"sessionTags": map[string]interface{}{
																"description": "AWS STS assume role session tags",
																"items": map[string]interface{}{
																	"properties": map[string]interface{}{
																		"key": map[string]interface{}{
																			"type": "string",
																		},
																		"value": map[string]interface{}{
																			"type": "string",
																		},
																	},
																	"required": []interface{}{
																		"key",
																		"value",
																	},
																	"type": "object",
																},
																"type": "array",
															},
															"transitiveTagKeys": map[string]interface{}{
																"description": "AWS STS assume role transitive session tags. Required when multiple rules are used with the provider",
																"items": map[string]interface{}{
																	"type": "string",
																},
																"type": "array",
															},
														},
														"required": []interface{}{
															"region",
															"service",
														},
														"type": "object",
													},
													"azurekv": map[string]interface{}{
														"description": "AzureKV configures this store to sync secrets using Azure Key Vault provider",
														"properties": map[string]interface{}{
															"authSecretRef": map[string]interface{}{
																"description": "Auth configures how the operator authenticates with Azure. Required for ServicePrincipal auth type.",
																"properties": map[string]interface{}{
																	"clientId": map[string]interface{}{
																		"description": "The Azure clientId of the service principle used for authentication.",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																				"type":        "string",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the Secret resource being referred to.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																	"clientSecret": map[string]interface{}{
																		"description": "The Azure ClientSecret of the service principle used for authentication.",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																				"type":        "string",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the Secret resource being referred to.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
															"authType": map[string]interface{}{
																"default":     "ServicePrincipal",
																"description": "Auth type defines how to authenticate to the keyvault service. Valid values are: - \"ServicePrincipal\" (default): Using a service principal (tenantId, clientId, clientSecret) - \"ManagedIdentity\": Using Managed Identity assigned to the pod (see aad-pod-identity)",
																"enum": []interface{}{
																	"ServicePrincipal",
																	"ManagedIdentity",
																	"WorkloadIdentity",
																},
																"type": "string",
															},
															"environmentType": map[string]interface{}{
																"default":     "PublicCloud",
																"description": "EnvironmentType specifies the Azure cloud environment endpoints to use for connecting and authenticating with Azure. By default it points to the public cloud AAD endpoint. The following endpoints are available, also see here: https://github.com/Azure/go-autorest/blob/main/autorest/azure/environments.go#L152 PublicCloud, USGovernmentCloud, ChinaCloud, GermanCloud",
																"enum": []interface{}{
																	"PublicCloud",
																	"USGovernmentCloud",
																	"ChinaCloud",
																	"GermanCloud",
																},
																"type": "string",
															},
															"identityId": map[string]interface{}{
																"description": "If multiple Managed Identity is assigned to the pod, you can select the one to be used",
																"type":        "string",
															},
															"serviceAccountRef": map[string]interface{}{
																"description": "ServiceAccountRef specified the service account that should be used when authenticating with WorkloadIdentity.",
																"properties": map[string]interface{}{
																	"audiences": map[string]interface{}{
																		"description": "Audience specifies the `aud` claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																		"items": map[string]interface{}{
																			"type": "string",
																		},
																		"type": "array",
																	},
																	"name": map[string]interface{}{
																		"description": "The name of the ServiceAccount resource being referred to.",
																		"type":        "string",
																	},
																	"namespace": map[string]interface{}{
																		"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																		"type":        "string",
																	},
																},
																"required": []interface{}{
																	"name",
																},
																"type": "object",
															},
															"tenantId": map[string]interface{}{
																"description": "TenantID configures the Azure Tenant to send requests to. Required for ServicePrincipal auth type.",
																"type":        "string",
															},
															"vaultUrl": map[string]interface{}{
																"description": "Vault Url from which the secrets to be fetched from.",
																"type":        "string",
															},
														},
														"required": []interface{}{
															"vaultUrl",
														},
														"type": "object",
													},
													"conjur": map[string]interface{}{
														"description": "Conjur configures this store to sync secrets using conjur provider",
														"properties": map[string]interface{}{
															"auth": map[string]interface{}{
																"properties": map[string]interface{}{
																	"apikey": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"account": map[string]interface{}{
																				"type": "string",
																			},
																			"apiKeyRef": map[string]interface{}{
																				"description": "A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"userRef": map[string]interface{}{
																				"description": "A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"required": []interface{}{
																			"account",
																			"apiKeyRef",
																			"userRef",
																		},
																		"type": "object",
																	},
																	"jwt": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"account": map[string]interface{}{
																				"type": "string",
																			},
																			"secretRef": map[string]interface{}{
																				"description": "Optional SecretRef that refers to a key in a Secret resource containing JWT token to authenticate with Conjur using the JWT authentication method.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"serviceAccountRef": map[string]interface{}{
																				"description": "Optional ServiceAccountRef specifies the Kubernetes service account for which to request a token for with the `TokenRequest` API.",
																				"properties": map[string]interface{}{
																					"audiences": map[string]interface{}{
																						"description": "Audience specifies the `aud` claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																						"items": map[string]interface{}{
																							"type": "string",
																						},
																						"type": "array",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the ServiceAccount resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"required": []interface{}{
																					"name",
																				},
																				"type": "object",
																			},
																			"serviceID": map[string]interface{}{
																				"description": "The conjur authn jwt webservice id",
																				"type":        "string",
																			},
																		},
																		"required": []interface{}{
																			"account",
																			"serviceID",
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
															"caBundle": map[string]interface{}{
																"type": "string",
															},
															"caProvider": map[string]interface{}{
																"description": "Used to provide custom certificate authority (CA) certificates for a secret store. The CAProvider points to a Secret or ConfigMap resource that contains a PEM-encoded certificate.",
																"properties": map[string]interface{}{
																	"key": map[string]interface{}{
																		"description": "The key where the CA certificate can be found in the Secret or ConfigMap.",
																		"type":        "string",
																	},
																	"name": map[string]interface{}{
																		"description": "The name of the object located at the provider type.",
																		"type":        "string",
																	},
																	"namespace": map[string]interface{}{
																		"description": "The namespace the Provider type is in. Can only be defined when used in a ClusterSecretStore.",
																		"type":        "string",
																	},
																	"type": map[string]interface{}{
																		"description": "The type of provider to use such as \"Secret\", or \"ConfigMap\".",
																		"enum": []interface{}{
																			"Secret",
																			"ConfigMap",
																		},
																		"type": "string",
																	},
																},
																"required": []interface{}{
																	"name",
																	"type",
																},
																"type": "object",
															},
															"url": map[string]interface{}{
																"type": "string",
															},
														},
														"required": []interface{}{
															"auth",
															"url",
														},
														"type": "object",
													},
													"delinea": map[string]interface{}{
														"description": "Delinea DevOps Secrets Vault https://docs.delinea.com/online-help/products/devops-secrets-vault/current",
														"properties": map[string]interface{}{
															"clientId": map[string]interface{}{
																"description": "ClientID is the non-secret part of the credential.",
																"properties": map[string]interface{}{
																	"secretRef": map[string]interface{}{
																		"description": "SecretRef references a key in a secret that will be used as value.",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																				"type":        "string",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the Secret resource being referred to.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																	"value": map[string]interface{}{
																		"description": "Value can be specified directly to set a value without using a secret.",
																		"type":        "string",
																	},
																},
																"type": "object",
															},
															"clientSecret": map[string]interface{}{
																"description": "ClientSecret is the secret part of the credential.",
																"properties": map[string]interface{}{
																	"secretRef": map[string]interface{}{
																		"description": "SecretRef references a key in a secret that will be used as value.",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																				"type":        "string",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the Secret resource being referred to.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																	"value": map[string]interface{}{
																		"description": "Value can be specified directly to set a value without using a secret.",
																		"type":        "string",
																	},
																},
																"type": "object",
															},
															"tenant": map[string]interface{}{
																"description": "Tenant is the chosen hostname / site name.",
																"type":        "string",
															},
															"tld": map[string]interface{}{
																"description": "TLD is based on the server location that was chosen during provisioning. If unset, defaults to \"com\".",
																"type":        "string",
															},
															"urlTemplate": map[string]interface{}{
																"description": "URLTemplate If unset, defaults to \"https://%s.secretsvaultcloud.%s/v1/%s%s\".",
																"type":        "string",
															},
														},
														"required": []interface{}{
															"clientId",
															"clientSecret",
															"tenant",
														},
														"type": "object",
													},
													"doppler": map[string]interface{}{
														"description": "Doppler configures this store to sync secrets using the Doppler provider",
														"properties": map[string]interface{}{
															"auth": map[string]interface{}{
																"description": "Auth configures how the Operator authenticates with the Doppler API",
																"properties": map[string]interface{}{
																	"secretRef": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"dopplerToken": map[string]interface{}{
																				"description": "The DopplerToken is used for authentication. See https://docs.doppler.com/reference/api#authentication for auth token types. The Key attribute defaults to dopplerToken if not specified.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"required": []interface{}{
																			"dopplerToken",
																		},
																		"type": "object",
																	},
																},
																"required": []interface{}{
																	"secretRef",
																},
																"type": "object",
															},
															"config": map[string]interface{}{
																"description": "Doppler config (required if not using a Service Token)",
																"type":        "string",
															},
															"format": map[string]interface{}{
																"description": "Format enables the downloading of secrets as a file (string)",
																"enum": []interface{}{
																	"json",
																	"dotnet-json",
																	"env",
																	"yaml",
																	"docker",
																},
																"type": "string",
															},
															"nameTransformer": map[string]interface{}{
																"description": "Environment variable compatible name transforms that change secret names to a different format",
																"enum": []interface{}{
																	"upper-camel",
																	"camel",
																	"lower-snake",
																	"tf-var",
																	"dotnet-env",
																	"lower-kebab",
																},
																"type": "string",
															},
															"project": map[string]interface{}{
																"description": "Doppler project (required if not using a Service Token)",
																"type":        "string",
															},
														},
														"required": []interface{}{
															"auth",
														},
														"type": "object",
													},
													"fake": map[string]interface{}{
														"description": "Fake configures a store with static key/value pairs",
														"properties": map[string]interface{}{
															"data": map[string]interface{}{
																"items": map[string]interface{}{
																	"properties": map[string]interface{}{
																		"key": map[string]interface{}{
																			"type": "string",
																		},
																		"value": map[string]interface{}{
																			"type": "string",
																		},
																		"valueMap": map[string]interface{}{
																			"additionalProperties": map[string]interface{}{
																				"type": "string",
																			},
																			"description": "Deprecated: ValueMap is deprecated and is intended to be removed in the future, use the `value` field instead.",
																			"type":        "object",
																		},
																		"version": map[string]interface{}{
																			"type": "string",
																		},
																	},
																	"required": []interface{}{
																		"key",
																	},
																	"type": "object",
																},
																"type": "array",
															},
														},
														"required": []interface{}{
															"data",
														},
														"type": "object",
													},
													"gcpsm": map[string]interface{}{
														"description": "GCPSM configures this store to sync secrets using Google Cloud Platform Secret Manager provider",
														"properties": map[string]interface{}{
															"auth": map[string]interface{}{
																"description": "Auth defines the information necessary to authenticate against GCP",
																"properties": map[string]interface{}{
																	"secretRef": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"secretAccessKeySecretRef": map[string]interface{}{
																				"description": "The SecretAccessKey is used for authentication",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																	"workloadIdentity": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"clusterLocation": map[string]interface{}{
																				"type": "string",
																			},
																			"clusterName": map[string]interface{}{
																				"type": "string",
																			},
																			"clusterProjectID": map[string]interface{}{
																				"type": "string",
																			},
																			"serviceAccountRef": map[string]interface{}{
																				"description": "A reference to a ServiceAccount resource.",
																				"properties": map[string]interface{}{
																					"audiences": map[string]interface{}{
																						"description": "Audience specifies the `aud` claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																						"items": map[string]interface{}{
																							"type": "string",
																						},
																						"type": "array",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the ServiceAccount resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"required": []interface{}{
																					"name",
																				},
																				"type": "object",
																			},
																		},
																		"required": []interface{}{
																			"clusterLocation",
																			"clusterName",
																			"serviceAccountRef",
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
															"projectID": map[string]interface{}{
																"description": "ProjectID project where secret is located",
																"type":        "string",
															},
														},
														"type": "object",
													},
													"gitlab": map[string]interface{}{
														"description": "GitLab configures this store to sync secrets using GitLab Variables provider",
														"properties": map[string]interface{}{
															"auth": map[string]interface{}{
																"description": "Auth configures how secret-manager authenticates with a GitLab instance.",
																"properties": map[string]interface{}{
																	"SecretRef": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"accessToken": map[string]interface{}{
																				"description": "AccessToken is used for authentication.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																},
																"required": []interface{}{
																	"SecretRef",
																},
																"type": "object",
															},
															"environment": map[string]interface{}{
																"description": "Environment environment_scope of gitlab CI/CD variables (Please see https://docs.gitlab.com/ee/ci/environments/#create-a-static-environment on how to create environments)",
																"type":        "string",
															},
															"groupIDs": map[string]interface{}{
																"description": "GroupIDs specify, which gitlab groups to pull secrets from. Group secrets are read from left to right followed by the project variables.",
																"items": map[string]interface{}{
																	"type": "string",
																},
																"type": "array",
															},
															"inheritFromGroups": map[string]interface{}{
																"description": "InheritFromGroups specifies whether parent groups should be discovered and checked for secrets.",
																"type":        "boolean",
															},
															"projectID": map[string]interface{}{
																"description": "ProjectID specifies a project where secrets are located.",
																"type":        "string",
															},
															"url": map[string]interface{}{
																"description": "URL configures the GitLab instance URL. Defaults to https://gitlab.com/.",
																"type":        "string",
															},
														},
														"required": []interface{}{
															"auth",
														},
														"type": "object",
													},
													"ibm": map[string]interface{}{
														"description": "IBM configures this store to sync secrets using IBM Cloud provider",
														"properties": map[string]interface{}{
															"auth": map[string]interface{}{
																"description":   "Auth configures how secret-manager authenticates with the IBM secrets manager.",
																"maxProperties": 1,
																"minProperties": 1,
																"properties": map[string]interface{}{
																	"containerAuth": map[string]interface{}{
																		"description": "IBM Container-based auth with IAM Trusted Profile.",
																		"properties": map[string]interface{}{
																			"iamEndpoint": map[string]interface{}{
																				"type": "string",
																			},
																			"profile": map[string]interface{}{
																				"description": "the IBM Trusted Profile",
																				"type":        "string",
																			},
																			"tokenLocation": map[string]interface{}{
																				"description": "Location the token is mounted on the pod",
																				"type":        "string",
																			},
																		},
																		"required": []interface{}{
																			"profile",
																		},
																		"type": "object",
																	},
																	"secretRef": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"secretApiKeySecretRef": map[string]interface{}{
																				"description": "The SecretAccessKey is used for authentication",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
															"serviceUrl": map[string]interface{}{
																"description": "ServiceURL is the Endpoint URL that is specific to the Secrets Manager service instance",
																"type":        "string",
															},
														},
														"required": []interface{}{
															"auth",
														},
														"type": "object",
													},
													"keepersecurity": map[string]interface{}{
														"description": "KeeperSecurity configures this store to sync secrets using the KeeperSecurity provider",
														"properties": map[string]interface{}{
															"authRef": map[string]interface{}{
																"description": "A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.",
																"properties": map[string]interface{}{
																	"key": map[string]interface{}{
																		"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																		"type":        "string",
																	},
																	"name": map[string]interface{}{
																		"description": "The name of the Secret resource being referred to.",
																		"type":        "string",
																	},
																	"namespace": map[string]interface{}{
																		"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																		"type":        "string",
																	},
																},
																"type": "object",
															},
															"folderID": map[string]interface{}{
																"type": "string",
															},
														},
														"required": []interface{}{
															"authRef",
															"folderID",
														},
														"type": "object",
													},
													"kubernetes": map[string]interface{}{
														"description": "Kubernetes configures this store to sync secrets using a Kubernetes cluster provider",
														"properties": map[string]interface{}{
															"auth": map[string]interface{}{
																"description":   "Auth configures how secret-manager authenticates with a Kubernetes instance.",
																"maxProperties": 1,
																"minProperties": 1,
																"properties": map[string]interface{}{
																	"cert": map[string]interface{}{
																		"description": "has both clientCert and clientKey as secretKeySelector",
																		"properties": map[string]interface{}{
																			"clientCert": map[string]interface{}{
																				"description": "A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"clientKey": map[string]interface{}{
																				"description": "A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																	"serviceAccount": map[string]interface{}{
																		"description": "points to a service account that should be used for authentication",
																		"properties": map[string]interface{}{
																			"audiences": map[string]interface{}{
																				"description": "Audience specifies the `aud` claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																				"items": map[string]interface{}{
																					"type": "string",
																				},
																				"type": "array",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the ServiceAccount resource being referred to.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																				"type":        "string",
																			},
																		},
																		"required": []interface{}{
																			"name",
																		},
																		"type": "object",
																	},
																	"token": map[string]interface{}{
																		"description": "use static token to authenticate with",
																		"properties": map[string]interface{}{
																			"bearerToken": map[string]interface{}{
																				"description": "A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
															"remoteNamespace": map[string]interface{}{
																"default":     "default",
																"description": "Remote namespace to fetch the secrets from",
																"type":        "string",
															},
															"server": map[string]interface{}{
																"description": "configures the Kubernetes server Address.",
																"properties": map[string]interface{}{
																	"caBundle": map[string]interface{}{
																		"description": "CABundle is a base64-encoded CA certificate",
																		"format":      "byte",
																		"type":        "string",
																	},
																	"caProvider": map[string]interface{}{
																		"description": "see: https://external-secrets.io/v0.4.1/spec/#external-secrets.io/v1alpha1.CAProvider",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "The key where the CA certificate can be found in the Secret or ConfigMap.",
																				"type":        "string",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the object located at the provider type.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "The namespace the Provider type is in. Can only be defined when used in a ClusterSecretStore.",
																				"type":        "string",
																			},
																			"type": map[string]interface{}{
																				"description": "The type of provider to use such as \"Secret\", or \"ConfigMap\".",
																				"enum": []interface{}{
																					"Secret",
																					"ConfigMap",
																				},
																				"type": "string",
																			},
																		},
																		"required": []interface{}{
																			"name",
																			"type",
																		},
																		"type": "object",
																	},
																	"url": map[string]interface{}{
																		"default":     "kubernetes.default",
																		"description": "configures the Kubernetes server Address.",
																		"type":        "string",
																	},
																},
																"type": "object",
															},
														},
														"required": []interface{}{
															"auth",
														},
														"type": "object",
													},
													"onepassword": map[string]interface{}{
														"description": "OnePassword configures this store to sync secrets using the 1Password Cloud provider",
														"properties": map[string]interface{}{
															"auth": map[string]interface{}{
																"description": "Auth defines the information necessary to authenticate against OnePassword Connect Server",
																"properties": map[string]interface{}{
																	"secretRef": map[string]interface{}{
																		"description": "OnePasswordAuthSecretRef holds secret references for 1Password credentials.",
																		"properties": map[string]interface{}{
																			"connectTokenSecretRef": map[string]interface{}{
																				"description": "The ConnectToken is used for authentication to a 1Password Connect Server.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"required": []interface{}{
																			"connectTokenSecretRef",
																		},
																		"type": "object",
																	},
																},
																"required": []interface{}{
																	"secretRef",
																},
																"type": "object",
															},
															"connectHost": map[string]interface{}{
																"description": "ConnectHost defines the OnePassword Connect Server to connect to",
																"type":        "string",
															},
															"vaults": map[string]interface{}{
																"additionalProperties": map[string]interface{}{
																	"type": "integer",
																},
																"description": "Vaults defines which OnePassword vaults to search in which order",
																"type":        "object",
															},
														},
														"required": []interface{}{
															"auth",
															"connectHost",
															"vaults",
														},
														"type": "object",
													},
													"oracle": map[string]interface{}{
														"description": "Oracle configures this store to sync secrets using Oracle Vault provider",
														"properties": map[string]interface{}{
															"auth": map[string]interface{}{
																"description": "Auth configures how secret-manager authenticates with the Oracle Vault. If empty, use the instance principal, otherwise the user credentials specified in Auth.",
																"properties": map[string]interface{}{
																	"secretRef": map[string]interface{}{
																		"description": "SecretRef to pass through sensitive information.",
																		"properties": map[string]interface{}{
																			"fingerprint": map[string]interface{}{
																				"description": "Fingerprint is the fingerprint of the API private key.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"privatekey": map[string]interface{}{
																				"description": "PrivateKey is the user's API Signing Key in PEM format, used for authentication.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"required": []interface{}{
																			"fingerprint",
																			"privatekey",
																		},
																		"type": "object",
																	},
																	"tenancy": map[string]interface{}{
																		"description": "Tenancy is the tenancy OCID where user is located.",
																		"type":        "string",
																	},
																	"user": map[string]interface{}{
																		"description": "User is an access OCID specific to the account.",
																		"type":        "string",
																	},
																},
																"required": []interface{}{
																	"secretRef",
																	"tenancy",
																	"user",
																},
																"type": "object",
															},
															"compartment": map[string]interface{}{
																"description": "Compartment is the vault compartment OCID. Required for PushSecret",
																"type":        "string",
															},
															"encryptionKey": map[string]interface{}{
																"description": "EncryptionKey is the OCID of the encryption key within the vault. Required for PushSecret",
																"type":        "string",
															},
															"principalType": map[string]interface{}{
																"description": "The type of principal to use for authentication. If left blank, the Auth struct will determine the principal type. This optional field must be specified if using workload identity.",
																"enum": []interface{}{
																	"",
																	"UserPrincipal",
																	"InstancePrincipal",
																	"Workload",
																},
																"type": "string",
															},
															"region": map[string]interface{}{
																"description": "Region is the region where vault is located.",
																"type":        "string",
															},
															"serviceAccountRef": map[string]interface{}{
																"description": "ServiceAccountRef specified the service account that should be used when authenticating with WorkloadIdentity.",
																"properties": map[string]interface{}{
																	"audiences": map[string]interface{}{
																		"description": "Audience specifies the `aud` claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																		"items": map[string]interface{}{
																			"type": "string",
																		},
																		"type": "array",
																	},
																	"name": map[string]interface{}{
																		"description": "The name of the ServiceAccount resource being referred to.",
																		"type":        "string",
																	},
																	"namespace": map[string]interface{}{
																		"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																		"type":        "string",
																	},
																},
																"required": []interface{}{
																	"name",
																},
																"type": "object",
															},
															"vault": map[string]interface{}{
																"description": "Vault is the vault's OCID of the specific vault where secret is located.",
																"type":        "string",
															},
														},
														"required": []interface{}{
															"region",
															"vault",
														},
														"type": "object",
													},
													"scaleway": map[string]interface{}{
														"description": "Scaleway",
														"properties": map[string]interface{}{
															"accessKey": map[string]interface{}{
																"description": "AccessKey is the non-secret part of the api key.",
																"properties": map[string]interface{}{
																	"secretRef": map[string]interface{}{
																		"description": "SecretRef references a key in a secret that will be used as value.",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																				"type":        "string",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the Secret resource being referred to.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																	"value": map[string]interface{}{
																		"description": "Value can be specified directly to set a value without using a secret.",
																		"type":        "string",
																	},
																},
																"type": "object",
															},
															"apiUrl": map[string]interface{}{
																"description": "APIURL is the url of the api to use. Defaults to https://api.scaleway.com",
																"type":        "string",
															},
															"projectId": map[string]interface{}{
																"description": "ProjectID is the id of your project, which you can find in the console: https://console.scaleway.com/project/settings",
																"type":        "string",
															},
															"region": map[string]interface{}{
																"description": "Region where your secrets are located: https://developers.scaleway.com/en/quickstart/#region-and-zone",
																"type":        "string",
															},
															"secretKey": map[string]interface{}{
																"description": "SecretKey is the non-secret part of the api key.",
																"properties": map[string]interface{}{
																	"secretRef": map[string]interface{}{
																		"description": "SecretRef references a key in a secret that will be used as value.",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																				"type":        "string",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the Secret resource being referred to.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																	"value": map[string]interface{}{
																		"description": "Value can be specified directly to set a value without using a secret.",
																		"type":        "string",
																	},
																},
																"type": "object",
															},
														},
														"required": []interface{}{
															"accessKey",
															"projectId",
															"region",
															"secretKey",
														},
														"type": "object",
													},
													"senhasegura": map[string]interface{}{
														"description": "Senhasegura configures this store to sync secrets using senhasegura provider",
														"properties": map[string]interface{}{
															"auth": map[string]interface{}{
																"description": "Auth defines parameters to authenticate in senhasegura",
																"properties": map[string]interface{}{
																	"clientId": map[string]interface{}{
																		"type": "string",
																	},
																	"clientSecretSecretRef": map[string]interface{}{
																		"description": "A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																				"type":        "string",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the Secret resource being referred to.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																},
																"required": []interface{}{
																	"clientId",
																	"clientSecretSecretRef",
																},
																"type": "object",
															},
															"ignoreSslCertificate": map[string]interface{}{
																"default":     false,
																"description": "IgnoreSslCertificate defines if SSL certificate must be ignored",
																"type":        "boolean",
															},
															"module": map[string]interface{}{
																"description": "Module defines which senhasegura module should be used to get secrets",
																"type":        "string",
															},
															"url": map[string]interface{}{
																"description": "URL of senhasegura",
																"type":        "string",
															},
														},
														"required": []interface{}{
															"auth",
															"module",
															"url",
														},
														"type": "object",
													},
													"vault": map[string]interface{}{
														"description": "Vault configures this store to sync secrets using Hashi provider",
														"properties": map[string]interface{}{
															"auth": map[string]interface{}{
																"description": "Auth configures how secret-manager authenticates with the Vault server.",
																"properties": map[string]interface{}{
																	"appRole": map[string]interface{}{
																		"description": "AppRole authenticates with Vault using the App Role auth mechanism, with the role and secret stored in a Kubernetes Secret resource.",
																		"properties": map[string]interface{}{
																			"path": map[string]interface{}{
																				"default":     "approle",
																				"description": "Path where the App Role authentication backend is mounted in Vault, e.g: \"approle\"",
																				"type":        "string",
																			},
																			"roleId": map[string]interface{}{
																				"description": "RoleID configured in the App Role authentication backend when setting up the authentication backend in Vault.",
																				"type":        "string",
																			},
																			"roleRef": map[string]interface{}{
																				"description": "Reference to a key in a Secret that contains the App Role ID used to authenticate with Vault. The `key` field must be specified and denotes which entry within the Secret resource is used as the app role id.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"secretRef": map[string]interface{}{
																				"description": "Reference to a key in a Secret that contains the App Role secret used to authenticate with Vault. The `key` field must be specified and denotes which entry within the Secret resource is used as the app role secret.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"required": []interface{}{
																			"path",
																			"secretRef",
																		},
																		"type": "object",
																	},
																	"cert": map[string]interface{}{
																		"description": "Cert authenticates with TLS Certificates by passing client certificate, private key and ca certificate Cert authentication method",
																		"properties": map[string]interface{}{
																			"clientCert": map[string]interface{}{
																				"description": "ClientCert is a certificate to authenticate using the Cert Vault authentication method",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"secretRef": map[string]interface{}{
																				"description": "SecretRef to a key in a Secret resource containing client private key to authenticate with Vault using the Cert authentication method",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																	"iam": map[string]interface{}{
																		"description": "Iam authenticates with vault by passing a special AWS request signed with AWS IAM credentials AWS IAM authentication method",
																		"properties": map[string]interface{}{
																			"externalID": map[string]interface{}{
																				"description": "AWS External ID set on assumed IAM roles",
																				"type":        "string",
																			},
																			"jwt": map[string]interface{}{
																				"description": "Specify a service account with IRSA enabled",
																				"properties": map[string]interface{}{
																					"serviceAccountRef": map[string]interface{}{
																						"description": "A reference to a ServiceAccount resource.",
																						"properties": map[string]interface{}{
																							"audiences": map[string]interface{}{
																								"description": "Audience specifies the `aud` claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																								"items": map[string]interface{}{
																									"type": "string",
																								},
																								"type": "array",
																							},
																							"name": map[string]interface{}{
																								"description": "The name of the ServiceAccount resource being referred to.",
																								"type":        "string",
																							},
																							"namespace": map[string]interface{}{
																								"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																								"type":        "string",
																							},
																						},
																						"required": []interface{}{
																							"name",
																						},
																						"type": "object",
																					},
																				},
																				"type": "object",
																			},
																			"path": map[string]interface{}{
																				"description": "Path where the AWS auth method is enabled in Vault, e.g: \"aws\"",
																				"type":        "string",
																			},
																			"region": map[string]interface{}{
																				"description": "AWS region",
																				"type":        "string",
																			},
																			"role": map[string]interface{}{
																				"description": "This is the AWS role to be assumed before talking to vault",
																				"type":        "string",
																			},
																			"secretRef": map[string]interface{}{
																				"description": "Specify credentials in a Secret object",
																				"properties": map[string]interface{}{
																					"accessKeyIDSecretRef": map[string]interface{}{
																						"description": "The AccessKeyID is used for authentication",
																						"properties": map[string]interface{}{
																							"key": map[string]interface{}{
																								"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																								"type":        "string",
																							},
																							"name": map[string]interface{}{
																								"description": "The name of the Secret resource being referred to.",
																								"type":        "string",
																							},
																							"namespace": map[string]interface{}{
																								"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																								"type":        "string",
																							},
																						},
																						"type": "object",
																					},
																					"secretAccessKeySecretRef": map[string]interface{}{
																						"description": "The SecretAccessKey is used for authentication",
																						"properties": map[string]interface{}{
																							"key": map[string]interface{}{
																								"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																								"type":        "string",
																							},
																							"name": map[string]interface{}{
																								"description": "The name of the Secret resource being referred to.",
																								"type":        "string",
																							},
																							"namespace": map[string]interface{}{
																								"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																								"type":        "string",
																							},
																						},
																						"type": "object",
																					},
																					"sessionTokenSecretRef": map[string]interface{}{
																						"description": "The SessionToken used for authentication This must be defined if AccessKeyID and SecretAccessKey are temporary credentials see: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_use-resources.html",
																						"properties": map[string]interface{}{
																							"key": map[string]interface{}{
																								"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																								"type":        "string",
																							},
																							"name": map[string]interface{}{
																								"description": "The name of the Secret resource being referred to.",
																								"type":        "string",
																							},
																							"namespace": map[string]interface{}{
																								"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																								"type":        "string",
																							},
																						},
																						"type": "object",
																					},
																				},
																				"type": "object",
																			},
																			"vaultAwsIamServerID": map[string]interface{}{
																				"description": "X-Vault-AWS-IAM-Server-ID is an additional header used by Vault IAM auth method to mitigate against different types of replay attacks. More details here: https://developer.hashicorp.com/vault/docs/auth/aws",
																				"type":        "string",
																			},
																			"vaultRole": map[string]interface{}{
																				"description": "Vault Role. In vault, a role describes an identity with a set of permissions, groups, or policies you want to attach a user of the secrets engine",
																				"type":        "string",
																			},
																		},
																		"required": []interface{}{
																			"vaultRole",
																		},
																		"type": "object",
																	},
																	"jwt": map[string]interface{}{
																		"description": "Jwt authenticates with Vault by passing role and JWT token using the JWT/OIDC authentication method",
																		"properties": map[string]interface{}{
																			"kubernetesServiceAccountToken": map[string]interface{}{
																				"description": "Optional ServiceAccountToken specifies the Kubernetes service account for which to request a token for with the `TokenRequest` API.",
																				"properties": map[string]interface{}{
																					"audiences": map[string]interface{}{
																						"description": "Optional audiences field that will be used to request a temporary Kubernetes service account token for the service account referenced by `serviceAccountRef`. Defaults to a single audience `vault` it not specified. Deprecated: use serviceAccountRef.Audiences instead",
																						"items": map[string]interface{}{
																							"type": "string",
																						},
																						"type": "array",
																					},
																					"expirationSeconds": map[string]interface{}{
																						"description": "Optional expiration time in seconds that will be used to request a temporary Kubernetes service account token for the service account referenced by `serviceAccountRef`. Deprecated: this will be removed in the future. Defaults to 10 minutes.",
																						"format":      "int64",
																						"type":        "integer",
																					},
																					"serviceAccountRef": map[string]interface{}{
																						"description": "Service account field containing the name of a kubernetes ServiceAccount.",
																						"properties": map[string]interface{}{
																							"audiences": map[string]interface{}{
																								"description": "Audience specifies the `aud` claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																								"items": map[string]interface{}{
																									"type": "string",
																								},
																								"type": "array",
																							},
																							"name": map[string]interface{}{
																								"description": "The name of the ServiceAccount resource being referred to.",
																								"type":        "string",
																							},
																							"namespace": map[string]interface{}{
																								"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																								"type":        "string",
																							},
																						},
																						"required": []interface{}{
																							"name",
																						},
																						"type": "object",
																					},
																				},
																				"required": []interface{}{
																					"serviceAccountRef",
																				},
																				"type": "object",
																			},
																			"path": map[string]interface{}{
																				"default":     "jwt",
																				"description": "Path where the JWT authentication backend is mounted in Vault, e.g: \"jwt\"",
																				"type":        "string",
																			},
																			"role": map[string]interface{}{
																				"description": "Role is a JWT role to authenticate using the JWT/OIDC Vault authentication method",
																				"type":        "string",
																			},
																			"secretRef": map[string]interface{}{
																				"description": "Optional SecretRef that refers to a key in a Secret resource containing JWT token to authenticate with Vault using the JWT/OIDC authentication method.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"required": []interface{}{
																			"path",
																		},
																		"type": "object",
																	},
																	"kubernetes": map[string]interface{}{
																		"description": "Kubernetes authenticates with Vault by passing the ServiceAccount token stored in the named Secret resource to the Vault server.",
																		"properties": map[string]interface{}{
																			"mountPath": map[string]interface{}{
																				"default":     "kubernetes",
																				"description": "Path where the Kubernetes authentication backend is mounted in Vault, e.g: \"kubernetes\"",
																				"type":        "string",
																			},
																			"role": map[string]interface{}{
																				"description": "A required field containing the Vault Role to assume. A Role binds a Kubernetes ServiceAccount with a set of Vault policies.",
																				"type":        "string",
																			},
																			"secretRef": map[string]interface{}{
																				"description": "Optional secret field containing a Kubernetes ServiceAccount JWT used for authenticating with Vault. If a name is specified without a key, `token` is the default. If one is not specified, the one bound to the controller will be used.",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"serviceAccountRef": map[string]interface{}{
																				"description": "Optional service account field containing the name of a kubernetes ServiceAccount. If the service account is specified, the service account secret token JWT will be used for authenticating with Vault. If the service account selector is not supplied, the secretRef will be used instead.",
																				"properties": map[string]interface{}{
																					"audiences": map[string]interface{}{
																						"description": "Audience specifies the `aud` claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																						"items": map[string]interface{}{
																							"type": "string",
																						},
																						"type": "array",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the ServiceAccount resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"required": []interface{}{
																					"name",
																				},
																				"type": "object",
																			},
																		},
																		"required": []interface{}{
																			"mountPath",
																			"role",
																		},
																		"type": "object",
																	},
																	"ldap": map[string]interface{}{
																		"description": "Ldap authenticates with Vault by passing username/password pair using the LDAP authentication method",
																		"properties": map[string]interface{}{
																			"path": map[string]interface{}{
																				"default":     "ldap",
																				"description": "Path where the LDAP authentication backend is mounted in Vault, e.g: \"ldap\"",
																				"type":        "string",
																			},
																			"secretRef": map[string]interface{}{
																				"description": "SecretRef to a key in a Secret resource containing password for the LDAP user used to authenticate with Vault using the LDAP authentication method",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"username": map[string]interface{}{
																				"description": "Username is a LDAP user name used to authenticate using the LDAP Vault authentication method",
																				"type":        "string",
																			},
																		},
																		"required": []interface{}{
																			"path",
																			"username",
																		},
																		"type": "object",
																	},
																	"tokenSecretRef": map[string]interface{}{
																		"description": "TokenSecretRef authenticates with Vault by presenting a token.",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																				"type":        "string",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the Secret resource being referred to.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																	"userPass": map[string]interface{}{
																		"description": "UserPass authenticates with Vault by passing username/password pair",
																		"properties": map[string]interface{}{
																			"path": map[string]interface{}{
																				"default":     "user",
																				"description": "Path where the UserPassword authentication backend is mounted in Vault, e.g: \"user\"",
																				"type":        "string",
																			},
																			"secretRef": map[string]interface{}{
																				"description": "SecretRef to a key in a Secret resource containing password for the user used to authenticate with Vault using the UserPass authentication method",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"username": map[string]interface{}{
																				"description": "Username is a user name used to authenticate using the UserPass Vault authentication method",
																				"type":        "string",
																			},
																		},
																		"required": []interface{}{
																			"path",
																			"username",
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
															"caBundle": map[string]interface{}{
																"description": "PEM encoded CA bundle used to validate Vault server certificate. Only used if the Server URL is using HTTPS protocol. This parameter is ignored for plain HTTP protocol connection. If not set the system root certificates are used to validate the TLS connection.",
																"format":      "byte",
																"type":        "string",
															},
															"caProvider": map[string]interface{}{
																"description": "The provider for the CA bundle to use to validate Vault server certificate.",
																"properties": map[string]interface{}{
																	"key": map[string]interface{}{
																		"description": "The key where the CA certificate can be found in the Secret or ConfigMap.",
																		"type":        "string",
																	},
																	"name": map[string]interface{}{
																		"description": "The name of the object located at the provider type.",
																		"type":        "string",
																	},
																	"namespace": map[string]interface{}{
																		"description": "The namespace the Provider type is in. Can only be defined when used in a ClusterSecretStore.",
																		"type":        "string",
																	},
																	"type": map[string]interface{}{
																		"description": "The type of provider to use such as \"Secret\", or \"ConfigMap\".",
																		"enum": []interface{}{
																			"Secret",
																			"ConfigMap",
																		},
																		"type": "string",
																	},
																},
																"required": []interface{}{
																	"name",
																	"type",
																},
																"type": "object",
															},
															"forwardInconsistent": map[string]interface{}{
																"description": "ForwardInconsistent tells Vault to forward read-after-write requests to the Vault leader instead of simply retrying within a loop. This can increase performance if the option is enabled serverside. https://www.vaultproject.io/docs/configuration/replication#allow_forwarding_via_header",
																"type":        "boolean",
															},
															"namespace": map[string]interface{}{
																"description": "Name of the vault namespace. Namespaces is a set of features within Vault Enterprise that allows Vault environments to support Secure Multi-tenancy. e.g: \"ns1\". More about namespaces can be found here https://www.vaultproject.io/docs/enterprise/namespaces",
																"type":        "string",
															},
															"path": map[string]interface{}{
																"description": "Path is the mount path of the Vault KV backend endpoint, e.g: \"secret\". The v2 KV secret engine version specific \"/data\" path suffix for fetching secrets from Vault is optional and will be appended if not present in specified path.",
																"type":        "string",
															},
															"readYourWrites": map[string]interface{}{
																"description": "ReadYourWrites ensures isolated read-after-write semantics by providing discovered cluster replication states in each request. More information about eventual consistency in Vault can be found here https://www.vaultproject.io/docs/enterprise/consistency",
																"type":        "boolean",
															},
															"server": map[string]interface{}{
																"description": "Server is the connection address for the Vault server, e.g: \"https://vault.example.com:8200\".",
																"type":        "string",
															},
															"version": map[string]interface{}{
																"default":     "v2",
																"description": "Version is the Vault KV secret engine version. This can be either \"v1\" or \"v2\". Version defaults to \"v2\".",
																"enum": []interface{}{
																	"v1",
																	"v2",
																},
																"type": "string",
															},
														},
														"required": []interface{}{
															"auth",
															"server",
														},
														"type": "object",
													},
													"webhook": map[string]interface{}{
														"description": "Webhook configures this store to sync secrets using a generic templated webhook",
														"properties": map[string]interface{}{
															"body": map[string]interface{}{
																"description": "Body",
																"type":        "string",
															},
															"caBundle": map[string]interface{}{
																"description": "PEM encoded CA bundle used to validate webhook server certificate. Only used if the Server URL is using HTTPS protocol. This parameter is ignored for plain HTTP protocol connection. If not set the system root certificates are used to validate the TLS connection.",
																"format":      "byte",
																"type":        "string",
															},
															"caProvider": map[string]interface{}{
																"description": "The provider for the CA bundle to use to validate webhook server certificate.",
																"properties": map[string]interface{}{
																	"key": map[string]interface{}{
																		"description": "The key the value inside of the provider type to use, only used with \"Secret\" type",
																		"type":        "string",
																	},
																	"name": map[string]interface{}{
																		"description": "The name of the object located at the provider type.",
																		"type":        "string",
																	},
																	"namespace": map[string]interface{}{
																		"description": "The namespace the Provider type is in.",
																		"type":        "string",
																	},
																	"type": map[string]interface{}{
																		"description": "The type of provider to use such as \"Secret\", or \"ConfigMap\".",
																		"enum": []interface{}{
																			"Secret",
																			"ConfigMap",
																		},
																		"type": "string",
																	},
																},
																"required": []interface{}{
																	"name",
																	"type",
																},
																"type": "object",
															},
															"headers": map[string]interface{}{
																"additionalProperties": map[string]interface{}{
																	"type": "string",
																},
																"description": "Headers",
																"type":        "object",
															},
															"method": map[string]interface{}{
																"description": "Webhook Method",
																"type":        "string",
															},
															"result": map[string]interface{}{
																"description": "Result formatting",
																"properties": map[string]interface{}{
																	"jsonPath": map[string]interface{}{
																		"description": "Json path of return value",
																		"type":        "string",
																	},
																},
																"type": "object",
															},
															"secrets": map[string]interface{}{
																"description": "Secrets to fill in templates These secrets will be passed to the templating function as key value pairs under the given name",
																"items": map[string]interface{}{
																	"properties": map[string]interface{}{
																		"name": map[string]interface{}{
																			"description": "Name of this secret in templates",
																			"type":        "string",
																		},
																		"secretRef": map[string]interface{}{
																			"description": "Secret ref to fill in credentials",
																			"properties": map[string]interface{}{
																				"key": map[string]interface{}{
																					"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																					"type":        "string",
																				},
																				"name": map[string]interface{}{
																					"description": "The name of the Secret resource being referred to.",
																					"type":        "string",
																				},
																				"namespace": map[string]interface{}{
																					"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																					"type":        "string",
																				},
																			},
																			"type": "object",
																		},
																	},
																	"required": []interface{}{
																		"name",
																		"secretRef",
																	},
																	"type": "object",
																},
																"type": "array",
															},
															"timeout": map[string]interface{}{
																"description": "Timeout",
																"type":        "string",
															},
															"url": map[string]interface{}{
																"description": "Webhook url to call",
																"type":        "string",
															},
														},
														"required": []interface{}{
															"result",
															"url",
														},
														"type": "object",
													},
													"yandexcertificatemanager": map[string]interface{}{
														"description": "YandexCertificateManager configures this store to sync secrets using Yandex Certificate Manager provider",
														"properties": map[string]interface{}{
															"apiEndpoint": map[string]interface{}{
																"description": "Yandex.Cloud API endpoint (e.g. 'api.cloud.yandex.net:443')",
																"type":        "string",
															},
															"auth": map[string]interface{}{
																"description": "Auth defines the information necessary to authenticate against Yandex Certificate Manager",
																"properties": map[string]interface{}{
																	"authorizedKeySecretRef": map[string]interface{}{
																		"description": "The authorized key used for authentication",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																				"type":        "string",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the Secret resource being referred to.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
															"caProvider": map[string]interface{}{
																"description": "The provider for the CA bundle to use to validate Yandex.Cloud server certificate.",
																"properties": map[string]interface{}{
																	"certSecretRef": map[string]interface{}{
																		"description": "A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																				"type":        "string",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the Secret resource being referred to.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
														},
														"required": []interface{}{
															"auth",
														},
														"type": "object",
													},
													"yandexlockbox": map[string]interface{}{
														"description": "YandexLockbox configures this store to sync secrets using Yandex Lockbox provider",
														"properties": map[string]interface{}{
															"apiEndpoint": map[string]interface{}{
																"description": "Yandex.Cloud API endpoint (e.g. 'api.cloud.yandex.net:443')",
																"type":        "string",
															},
															"auth": map[string]interface{}{
																"description": "Auth defines the information necessary to authenticate against Yandex Lockbox",
																"properties": map[string]interface{}{
																	"authorizedKeySecretRef": map[string]interface{}{
																		"description": "The authorized key used for authentication",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																				"type":        "string",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the Secret resource being referred to.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
															"caProvider": map[string]interface{}{
																"description": "The provider for the CA bundle to use to validate Yandex.Cloud server certificate.",
																"properties": map[string]interface{}{
																	"certSecretRef": map[string]interface{}{
																		"description": "A reference to a specific 'key' within a Secret resource, In some instances, `key` is a required field.",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																				"type":        "string",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the Secret resource being referred to.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
														},
														"required": []interface{}{
															"auth",
														},
														"type": "object",
													},
												},
												"type": "object",
											},
											"refreshInterval": map[string]interface{}{
												"description": "Used to configure store refresh interval in seconds. Empty or 0 will default to the controller config.",
												"type":        "integer",
											},
											"retrySettings": map[string]interface{}{
												"description": "Used to configure http retries if failed",
												"properties": map[string]interface{}{
													"maxRetries": map[string]interface{}{
														"format": "int32",
														"type":   "integer",
													},
													"retryInterval": map[string]interface{}{
														"type": "string",
													},
												},
												"type": "object",
											},
										},
										"required": []interface{}{
											"provider",
										},
										"type": "object",
									},
									"status": map[string]interface{}{
										"description": "SecretStoreStatus defines the observed state of the SecretStore.",
										"properties": map[string]interface{}{
											"capabilities": map[string]interface{}{
												"description": "SecretStoreCapabilities defines the possible operations a SecretStore can do.",
												"type":        "string",
											},
											"conditions": map[string]interface{}{
												"items": map[string]interface{}{
													"properties": map[string]interface{}{
														"lastTransitionTime": map[string]interface{}{
															"format": "date-time",
															"type":   "string",
														},
														"message": map[string]interface{}{
															"type": "string",
														},
														"reason": map[string]interface{}{
															"type": "string",
														},
														"status": map[string]interface{}{
															"type": "string",
														},
														"type": map[string]interface{}{
															"type": "string",
														},
													},
													"required": []interface{}{
														"status",
														"type",
													},
													"type": "object",
												},
												"type": "array",
											},
										},
										"type": "object",
									},
								},
								"type": "object",
							},
						},
						"served":  true,
						"storage": true,
						"subresources": map[string]interface{}{
							"status": map[string]interface{}{},
						},
					},
				},
				"conversion": map[string]interface{}{
					"strategy": "Webhook",
					"webhook": map[string]interface{}{
						"conversionReviewVersions": []interface{}{
							"v1",
						},
						"clientConfig": map[string]interface{}{
							"service": map[string]interface{}{
								"name":      "release-name-external-secrets-webhook",
								"namespace": "default",
								"path":      "/convert",
							},
						},
					},
				},
			},
		},
	}

	return mutate.MutateCRDSecretstoresExternalSecretsIo(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;create;update;patch;delete

// CreateCRDVaultdynamicsecretsGeneratorsExternalSecretsIo creates the CustomResourceDefinition resource with name vaultdynamicsecrets.generators.external-secrets.io.
func CreateCRDVaultdynamicsecretsGeneratorsExternalSecretsIo(
	parent *secretsv1alpha1.ExternalSecrets,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// Source: external-secrets/templates/crds/vaultdynamicsecret.yaml
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					"controller-gen.kubebuilder.io/version": "v0.13.0",
				},
				"name": "vaultdynamicsecrets.generators.external-secrets.io",
			},
			"spec": map[string]interface{}{
				"group": "generators.external-secrets.io",
				"names": map[string]interface{}{
					"categories": []interface{}{
						"vaultdynamicsecret",
					},
					"kind":     "VaultDynamicSecret",
					"listKind": "VaultDynamicSecretList",
					"plural":   "vaultdynamicsecrets",
					"shortNames": []interface{}{
						"vaultdynamicsecret",
					},
					"singular": "vaultdynamicsecret",
				},
				"scope": "Namespaced",
				"versions": []interface{}{
					map[string]interface{}{
						"name": "v1alpha1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"properties": map[string]interface{}{
									"apiVersion": map[string]interface{}{
										"description": "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
										"type":        "string",
									},
									"kind": map[string]interface{}{
										"description": "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
										"type":        "string",
									},
									"metadata": map[string]interface{}{
										"type": "object",
									},
									"spec": map[string]interface{}{
										"properties": map[string]interface{}{
											"controller": map[string]interface{}{
												"description": "Used to select the correct ESO controller (think: ingress.ingressClassName) The ESO controller is instantiated with a specific controller name and filters VDS based on this property",
												"type":        "string",
											},
											"method": map[string]interface{}{
												"description": "Vault API method to use (GET/POST/other)",
												"type":        "string",
											},
											"parameters": map[string]interface{}{
												"description":                          "Parameters to pass to Vault write (for non-GET methods)",
												"x-kubernetes-preserve-unknown-fields": true,
											},
											"path": map[string]interface{}{
												"description": "Vault path to obtain the dynamic secret from",
												"type":        "string",
											},
											"provider": map[string]interface{}{
												"description": "Vault provider common spec",
												"properties": map[string]interface{}{
													"auth": map[string]interface{}{
														"description": "Auth configures how secret-manager authenticates with the Vault server.",
														"properties": map[string]interface{}{
															"appRole": map[string]interface{}{
																"description": "AppRole authenticates with Vault using the App Role auth mechanism, with the role and secret stored in a Kubernetes Secret resource.",
																"properties": map[string]interface{}{
																	"path": map[string]interface{}{
																		"default":     "approle",
																		"description": "Path where the App Role authentication backend is mounted in Vault, e.g: \"approle\"",
																		"type":        "string",
																	},
																	"roleId": map[string]interface{}{
																		"description": "RoleID configured in the App Role authentication backend when setting up the authentication backend in Vault.",
																		"type":        "string",
																	},
																	"roleRef": map[string]interface{}{
																		"description": "Reference to a key in a Secret that contains the App Role ID used to authenticate with Vault. The `key` field must be specified and denotes which entry within the Secret resource is used as the app role id.",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																				"type":        "string",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the Secret resource being referred to.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																	"secretRef": map[string]interface{}{
																		"description": "Reference to a key in a Secret that contains the App Role secret used to authenticate with Vault. The `key` field must be specified and denotes which entry within the Secret resource is used as the app role secret.",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																				"type":        "string",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the Secret resource being referred to.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																},
																"required": []interface{}{
																	"path",
																	"secretRef",
																},
																"type": "object",
															},
															"cert": map[string]interface{}{
																"description": "Cert authenticates with TLS Certificates by passing client certificate, private key and ca certificate Cert authentication method",
																"properties": map[string]interface{}{
																	"clientCert": map[string]interface{}{
																		"description": "ClientCert is a certificate to authenticate using the Cert Vault authentication method",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																				"type":        "string",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the Secret resource being referred to.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																	"secretRef": map[string]interface{}{
																		"description": "SecretRef to a key in a Secret resource containing client private key to authenticate with Vault using the Cert authentication method",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																				"type":        "string",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the Secret resource being referred to.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
															"iam": map[string]interface{}{
																"description": "Iam authenticates with vault by passing a special AWS request signed with AWS IAM credentials AWS IAM authentication method",
																"properties": map[string]interface{}{
																	"externalID": map[string]interface{}{
																		"description": "AWS External ID set on assumed IAM roles",
																		"type":        "string",
																	},
																	"jwt": map[string]interface{}{
																		"description": "Specify a service account with IRSA enabled",
																		"properties": map[string]interface{}{
																			"serviceAccountRef": map[string]interface{}{
																				"description": "A reference to a ServiceAccount resource.",
																				"properties": map[string]interface{}{
																					"audiences": map[string]interface{}{
																						"description": "Audience specifies the `aud` claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																						"items": map[string]interface{}{
																							"type": "string",
																						},
																						"type": "array",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the ServiceAccount resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"required": []interface{}{
																					"name",
																				},
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																	"path": map[string]interface{}{
																		"description": "Path where the AWS auth method is enabled in Vault, e.g: \"aws\"",
																		"type":        "string",
																	},
																	"region": map[string]interface{}{
																		"description": "AWS region",
																		"type":        "string",
																	},
																	"role": map[string]interface{}{
																		"description": "This is the AWS role to be assumed before talking to vault",
																		"type":        "string",
																	},
																	"secretRef": map[string]interface{}{
																		"description": "Specify credentials in a Secret object",
																		"properties": map[string]interface{}{
																			"accessKeyIDSecretRef": map[string]interface{}{
																				"description": "The AccessKeyID is used for authentication",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"secretAccessKeySecretRef": map[string]interface{}{
																				"description": "The SecretAccessKey is used for authentication",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"sessionTokenSecretRef": map[string]interface{}{
																				"description": "The SessionToken used for authentication This must be defined if AccessKeyID and SecretAccessKey are temporary credentials see: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_use-resources.html",
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																						"type":        "string",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the Secret resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																	"vaultAwsIamServerID": map[string]interface{}{
																		"description": "X-Vault-AWS-IAM-Server-ID is an additional header used by Vault IAM auth method to mitigate against different types of replay attacks. More details here: https://developer.hashicorp.com/vault/docs/auth/aws",
																		"type":        "string",
																	},
																	"vaultRole": map[string]interface{}{
																		"description": "Vault Role. In vault, a role describes an identity with a set of permissions, groups, or policies you want to attach a user of the secrets engine",
																		"type":        "string",
																	},
																},
																"required": []interface{}{
																	"vaultRole",
																},
																"type": "object",
															},
															"jwt": map[string]interface{}{
																"description": "Jwt authenticates with Vault by passing role and JWT token using the JWT/OIDC authentication method",
																"properties": map[string]interface{}{
																	"kubernetesServiceAccountToken": map[string]interface{}{
																		"description": "Optional ServiceAccountToken specifies the Kubernetes service account for which to request a token for with the `TokenRequest` API.",
																		"properties": map[string]interface{}{
																			"audiences": map[string]interface{}{
																				"description": "Optional audiences field that will be used to request a temporary Kubernetes service account token for the service account referenced by `serviceAccountRef`. Defaults to a single audience `vault` it not specified. Deprecated: use serviceAccountRef.Audiences instead",
																				"items": map[string]interface{}{
																					"type": "string",
																				},
																				"type": "array",
																			},
																			"expirationSeconds": map[string]interface{}{
																				"description": "Optional expiration time in seconds that will be used to request a temporary Kubernetes service account token for the service account referenced by `serviceAccountRef`. Deprecated: this will be removed in the future. Defaults to 10 minutes.",
																				"format":      "int64",
																				"type":        "integer",
																			},
																			"serviceAccountRef": map[string]interface{}{
																				"description": "Service account field containing the name of a kubernetes ServiceAccount.",
																				"properties": map[string]interface{}{
																					"audiences": map[string]interface{}{
																						"description": "Audience specifies the `aud` claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																						"items": map[string]interface{}{
																							"type": "string",
																						},
																						"type": "array",
																					},
																					"name": map[string]interface{}{
																						"description": "The name of the ServiceAccount resource being referred to.",
																						"type":        "string",
																					},
																					"namespace": map[string]interface{}{
																						"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																						"type":        "string",
																					},
																				},
																				"required": []interface{}{
																					"name",
																				},
																				"type": "object",
																			},
																		},
																		"required": []interface{}{
																			"serviceAccountRef",
																		},
																		"type": "object",
																	},
																	"path": map[string]interface{}{
																		"default":     "jwt",
																		"description": "Path where the JWT authentication backend is mounted in Vault, e.g: \"jwt\"",
																		"type":        "string",
																	},
																	"role": map[string]interface{}{
																		"description": "Role is a JWT role to authenticate using the JWT/OIDC Vault authentication method",
																		"type":        "string",
																	},
																	"secretRef": map[string]interface{}{
																		"description": "Optional SecretRef that refers to a key in a Secret resource containing JWT token to authenticate with Vault using the JWT/OIDC authentication method.",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																				"type":        "string",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the Secret resource being referred to.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																},
																"required": []interface{}{
																	"path",
																},
																"type": "object",
															},
															"kubernetes": map[string]interface{}{
																"description": "Kubernetes authenticates with Vault by passing the ServiceAccount token stored in the named Secret resource to the Vault server.",
																"properties": map[string]interface{}{
																	"mountPath": map[string]interface{}{
																		"default":     "kubernetes",
																		"description": "Path where the Kubernetes authentication backend is mounted in Vault, e.g: \"kubernetes\"",
																		"type":        "string",
																	},
																	"role": map[string]interface{}{
																		"description": "A required field containing the Vault Role to assume. A Role binds a Kubernetes ServiceAccount with a set of Vault policies.",
																		"type":        "string",
																	},
																	"secretRef": map[string]interface{}{
																		"description": "Optional secret field containing a Kubernetes ServiceAccount JWT used for authenticating with Vault. If a name is specified without a key, `token` is the default. If one is not specified, the one bound to the controller will be used.",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																				"type":        "string",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the Secret resource being referred to.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																	"serviceAccountRef": map[string]interface{}{
																		"description": "Optional service account field containing the name of a kubernetes ServiceAccount. If the service account is specified, the service account secret token JWT will be used for authenticating with Vault. If the service account selector is not supplied, the secretRef will be used instead.",
																		"properties": map[string]interface{}{
																			"audiences": map[string]interface{}{
																				"description": "Audience specifies the `aud` claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																				"items": map[string]interface{}{
																					"type": "string",
																				},
																				"type": "array",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the ServiceAccount resource being referred to.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																				"type":        "string",
																			},
																		},
																		"required": []interface{}{
																			"name",
																		},
																		"type": "object",
																	},
																},
																"required": []interface{}{
																	"mountPath",
																	"role",
																},
																"type": "object",
															},
															"ldap": map[string]interface{}{
																"description": "Ldap authenticates with Vault by passing username/password pair using the LDAP authentication method",
																"properties": map[string]interface{}{
																	"path": map[string]interface{}{
																		"default":     "ldap",
																		"description": "Path where the LDAP authentication backend is mounted in Vault, e.g: \"ldap\"",
																		"type":        "string",
																	},
																	"secretRef": map[string]interface{}{
																		"description": "SecretRef to a key in a Secret resource containing password for the LDAP user used to authenticate with Vault using the LDAP authentication method",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																				"type":        "string",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the Secret resource being referred to.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																	"username": map[string]interface{}{
																		"description": "Username is a LDAP user name used to authenticate using the LDAP Vault authentication method",
																		"type":        "string",
																	},
																},
																"required": []interface{}{
																	"path",
																	"username",
																},
																"type": "object",
															},
															"tokenSecretRef": map[string]interface{}{
																"description": "TokenSecretRef authenticates with Vault by presenting a token.",
																"properties": map[string]interface{}{
																	"key": map[string]interface{}{
																		"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																		"type":        "string",
																	},
																	"name": map[string]interface{}{
																		"description": "The name of the Secret resource being referred to.",
																		"type":        "string",
																	},
																	"namespace": map[string]interface{}{
																		"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																		"type":        "string",
																	},
																},
																"type": "object",
															},
															"userPass": map[string]interface{}{
																"description": "UserPass authenticates with Vault by passing username/password pair",
																"properties": map[string]interface{}{
																	"path": map[string]interface{}{
																		"default":     "user",
																		"description": "Path where the UserPassword authentication backend is mounted in Vault, e.g: \"user\"",
																		"type":        "string",
																	},
																	"secretRef": map[string]interface{}{
																		"description": "SecretRef to a key in a Secret resource containing password for the user used to authenticate with Vault using the UserPass authentication method",
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"description": "The key of the entry in the Secret resource's `data` field to be used. Some instances of this field may be defaulted, in others it may be required.",
																				"type":        "string",
																			},
																			"name": map[string]interface{}{
																				"description": "The name of the Secret resource being referred to.",
																				"type":        "string",
																			},
																			"namespace": map[string]interface{}{
																				"description": "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped. cluster-scoped defaults to the namespace of the referent.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																	"username": map[string]interface{}{
																		"description": "Username is a user name used to authenticate using the UserPass Vault authentication method",
																		"type":        "string",
																	},
																},
																"required": []interface{}{
																	"path",
																	"username",
																},
																"type": "object",
															},
														},
														"type": "object",
													},
													"caBundle": map[string]interface{}{
														"description": "PEM encoded CA bundle used to validate Vault server certificate. Only used if the Server URL is using HTTPS protocol. This parameter is ignored for plain HTTP protocol connection. If not set the system root certificates are used to validate the TLS connection.",
														"format":      "byte",
														"type":        "string",
													},
													"caProvider": map[string]interface{}{
														"description": "The provider for the CA bundle to use to validate Vault server certificate.",
														"properties": map[string]interface{}{
															"key": map[string]interface{}{
																"description": "The key where the CA certificate can be found in the Secret or ConfigMap.",
																"type":        "string",
															},
															"name": map[string]interface{}{
																"description": "The name of the object located at the provider type.",
																"type":        "string",
															},
															"namespace": map[string]interface{}{
																"description": "The namespace the Provider type is in. Can only be defined when used in a ClusterSecretStore.",
																"type":        "string",
															},
															"type": map[string]interface{}{
																"description": "The type of provider to use such as \"Secret\", or \"ConfigMap\".",
																"enum": []interface{}{
																	"Secret",
																	"ConfigMap",
																},
																"type": "string",
															},
														},
														"required": []interface{}{
															"name",
															"type",
														},
														"type": "object",
													},
													"forwardInconsistent": map[string]interface{}{
														"description": "ForwardInconsistent tells Vault to forward read-after-write requests to the Vault leader instead of simply retrying within a loop. This can increase performance if the option is enabled serverside. https://www.vaultproject.io/docs/configuration/replication#allow_forwarding_via_header",
														"type":        "boolean",
													},
													"namespace": map[string]interface{}{
														"description": "Name of the vault namespace. Namespaces is a set of features within Vault Enterprise that allows Vault environments to support Secure Multi-tenancy. e.g: \"ns1\". More about namespaces can be found here https://www.vaultproject.io/docs/enterprise/namespaces",
														"type":        "string",
													},
													"path": map[string]interface{}{
														"description": "Path is the mount path of the Vault KV backend endpoint, e.g: \"secret\". The v2 KV secret engine version specific \"/data\" path suffix for fetching secrets from Vault is optional and will be appended if not present in specified path.",
														"type":        "string",
													},
													"readYourWrites": map[string]interface{}{
														"description": "ReadYourWrites ensures isolated read-after-write semantics by providing discovered cluster replication states in each request. More information about eventual consistency in Vault can be found here https://www.vaultproject.io/docs/enterprise/consistency",
														"type":        "boolean",
													},
													"server": map[string]interface{}{
														"description": "Server is the connection address for the Vault server, e.g: \"https://vault.example.com:8200\".",
														"type":        "string",
													},
													"version": map[string]interface{}{
														"default":     "v2",
														"description": "Version is the Vault KV secret engine version. This can be either \"v1\" or \"v2\". Version defaults to \"v2\".",
														"enum": []interface{}{
															"v1",
															"v2",
														},
														"type": "string",
													},
												},
												"required": []interface{}{
													"auth",
													"server",
												},
												"type": "object",
											},
											"resultType": map[string]interface{}{
												"default":     "Data",
												"description": "Result type defines which data is returned from the generator. By default it is the \"data\" section of the Vault API response. When using e.g. /auth/token/create the \"data\" section is empty but the \"auth\" section contains the generated token. Please refer to the vault docs regarding the result data structure.",
												"enum": []interface{}{
													"Data",
													"Auth",
												},
												"type": "string",
											},
										},
										"required": []interface{}{
											"path",
											"provider",
										},
										"type": "object",
									},
								},
								"type": "object",
							},
						},
						"served":  true,
						"storage": true,
						"subresources": map[string]interface{}{
							"status": map[string]interface{}{},
						},
					},
				},
				"conversion": map[string]interface{}{
					"strategy": "Webhook",
					"webhook": map[string]interface{}{
						"conversionReviewVersions": []interface{}{
							"v1",
						},
						"clientConfig": map[string]interface{}{
							"service": map[string]interface{}{
								"name":      "release-name-external-secrets-webhook",
								"namespace": "default",
								"path":      "/convert",
							},
						},
					},
				},
			},
		},
	}

	return mutate.MutateCRDVaultdynamicsecretsGeneratorsExternalSecretsIo(resourceObj, parent, collection, reconciler, req)
}
