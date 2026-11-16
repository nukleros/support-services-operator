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

package autoscalingcomponent

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	platformv1alpha1 "github.com/nukleros/support-services-operator/apis/platform/v1alpha1"
	"github.com/nukleros/support-services-operator/apis/platform/v1alpha1/autoscalingcomponent/mutate"
	setupv1alpha1 "github.com/nukleros/support-services-operator/apis/setup/v1alpha1"
)

// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;create;update;patch;delete

// CreateCRDCloudeventsourcesEventingKedaSh creates the CustomResourceDefinition resource with name cloudeventsources.eventing.keda.sh.
func CreateCRDCloudeventsourcesEventingKedaSh(
	parent *platformv1alpha1.AutoscalingComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					"controller-gen.kubebuilder.io/version": "v0.21.0",
				},
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "keda-operator",
					"app.kubernetes.io/component":   "operator",
					"app.kubernetes.io/part-of":     "keda-operator",
					"app.kubernetes.io/version":     parent.Spec.Keda.Version, //  controlled by field: keda.version
					"platform.nukleros.io/category": "autoscaling",
					"platform.nukleros.io/project":  "keda",
				},
				"name": "cloudeventsources.eventing.keda.sh",
			},
			"spec": map[string]interface{}{
				"group": "eventing.keda.sh",
				"names": map[string]interface{}{
					"kind":     "CloudEventSource",
					"listKind": "CloudEventSourceList",
					"plural":   "cloudeventsources",
					"singular": "cloudeventsource",
				},
				"scope": "Namespaced",
				"versions": []interface{}{
					map[string]interface{}{
						"additionalPrinterColumns": []interface{}{
							map[string]interface{}{
								"jsonPath": ".status.conditions[?(@.type==\"Active\")].status",
								"name":     "Active",
								"type":     "string",
							},
						},
						"name": "v1alpha1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"description": "CloudEventSource defines how a KEDA event will be sent to event sink",
								"properties": map[string]interface{}{
									"apiVersion": map[string]interface{}{
										"description": `APIVersion defines the versioned schema of this representation of an object.
Servers should convert recognized schemas to the latest internal value, and
may reject unrecognized values.
More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources`,
										"type": "string",
									},
									"kind": map[string]interface{}{
										"description": `Kind is a string value representing the REST resource this object represents.
Servers may infer this from the endpoint the client submits requests to.
Cannot be updated.
In CamelCase.
More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds`,
										"type": "string",
									},
									"metadata": map[string]interface{}{
										"type": "object",
									},
									"spec": map[string]interface{}{
										"description": "CloudEventSourceSpec defines the spec of CloudEventSource",
										"properties": map[string]interface{}{
											"authenticationRef": map[string]interface{}{
												"description": `AuthenticationRef points to the TriggerAuthentication or ClusterTriggerAuthentication object that
is used to authenticate the scaler with the environment`,
												"properties": map[string]interface{}{
													"kind": map[string]interface{}{
														"description": "Kind of the resource being referred to. Defaults to TriggerAuthentication.",
														"enum": []interface{}{
															"TriggerAuthentication",
															"ClusterTriggerAuthentication",
														},
														"type": "string",
													},
													"name": map[string]interface{}{
														"type": "string",
													},
												},
												"required": []interface{}{
													"name",
												},
												"type": "object",
											},
											"clusterName": map[string]interface{}{
												"type": "string",
											},
											"destination": map[string]interface{}{
												"description": "Destination defines the various ways to emit events",
												"properties": map[string]interface{}{
													"azureEventGridTopic": map[string]interface{}{
														"properties": map[string]interface{}{
															"endpoint": map[string]interface{}{
																"type": "string",
															},
														},
														"required": []interface{}{
															"endpoint",
														},
														"type": "object",
													},
													"http": map[string]interface{}{
														"properties": map[string]interface{}{
															"uri": map[string]interface{}{
																"type": "string",
															},
														},
														"required": []interface{}{
															"uri",
														},
														"type": "object",
													},
												},
												"type": "object",
											},
											"eventSubscription": map[string]interface{}{
												"description": "EventSubscription defines filters for events",
												"properties": map[string]interface{}{
													"excludedEventTypes": map[string]interface{}{
														"items": map[string]interface{}{
															"enum": []interface{}{
																"keda.scaledobject.ready.v1",
																"keda.scaledobject.failed.v1",
																"keda.scaledobject.removed.v1",
																"keda.scaledobject.paused.v1",
																"keda.scaledobject.unpaused.v1",
																"keda.scaledjob.ready.v1",
																"keda.scaledjob.failed.v1",
																"keda.scaledjob.removed.v1",
																"keda.scaledjob.paused.v1",
																"keda.scaledjob.unpaused.v1",
																"keda.scaledjob.rolloutcleanup.started.v1",
																"keda.scaledjob.rolloutcleanup.completed.v1",
																"keda.scaledjob.rolloutcleanup.failed.v1",
																"keda.authentication.triggerauthentication.created.v1",
																"keda.authentication.triggerauthentication.updated.v1",
																"keda.authentication.triggerauthentication.removed.v1",
																"keda.authentication.clustertriggerauthentication.created.v1",
																"keda.authentication.clustertriggerauthentication.updated.v1",
																"keda.authentication.clustertriggerauthentication.removed.v1",
															},
															"type": "string",
														},
														"type": "array",
													},
													"includedEventTypes": map[string]interface{}{
														"items": map[string]interface{}{
															"enum": []interface{}{
																"keda.scaledobject.ready.v1",
																"keda.scaledobject.failed.v1",
																"keda.scaledobject.removed.v1",
																"keda.scaledobject.paused.v1",
																"keda.scaledobject.unpaused.v1",
																"keda.scaledjob.ready.v1",
																"keda.scaledjob.failed.v1",
																"keda.scaledjob.removed.v1",
																"keda.scaledjob.paused.v1",
																"keda.scaledjob.unpaused.v1",
																"keda.scaledjob.rolloutcleanup.started.v1",
																"keda.scaledjob.rolloutcleanup.completed.v1",
																"keda.scaledjob.rolloutcleanup.failed.v1",
																"keda.authentication.triggerauthentication.created.v1",
																"keda.authentication.triggerauthentication.updated.v1",
																"keda.authentication.triggerauthentication.removed.v1",
																"keda.authentication.clustertriggerauthentication.created.v1",
																"keda.authentication.clustertriggerauthentication.updated.v1",
																"keda.authentication.clustertriggerauthentication.removed.v1",
															},
															"type": "string",
														},
														"type": "array",
													},
												},
												"type": "object",
											},
										},
										"required": []interface{}{
											"destination",
										},
										"type": "object",
									},
									"status": map[string]interface{}{
										"description": "CloudEventSourceStatus defines the observed state of CloudEventSource",
										"properties": map[string]interface{}{
											"conditions": map[string]interface{}{
												"description": "Conditions an array representation to store multiple Conditions",
												"items": map[string]interface{}{
													"description": "Condition to store the condition state",
													"properties": map[string]interface{}{
														"message": map[string]interface{}{
															"description": "A human readable message indicating details about the transition.",
															"type":        "string",
														},
														"reason": map[string]interface{}{
															"description": "The reason for the condition's last transition.",
															"type":        "string",
														},
														"status": map[string]interface{}{
															"description": "Status of the condition, one of True, False, Unknown.",
															"type":        "string",
														},
														"type": map[string]interface{}{
															"description": "Type of condition",
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
										},
										"type": "object",
									},
								},
								"required": []interface{}{
									"spec",
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
			},
		},
	}

	return mutate.MutateCRDCloudeventsourcesEventingKedaSh(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;create;update;patch;delete

// CreateCRDClustercloudeventsourcesEventingKedaSh creates the CustomResourceDefinition resource with name clustercloudeventsources.eventing.keda.sh.
func CreateCRDClustercloudeventsourcesEventingKedaSh(
	parent *platformv1alpha1.AutoscalingComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					"controller-gen.kubebuilder.io/version": "v0.21.0",
				},
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "keda-operator",
					"app.kubernetes.io/component":   "operator",
					"app.kubernetes.io/part-of":     "keda-operator",
					"app.kubernetes.io/version":     parent.Spec.Keda.Version, //  controlled by field: keda.version
					"platform.nukleros.io/category": "autoscaling",
					"platform.nukleros.io/project":  "keda",
				},
				"name": "clustercloudeventsources.eventing.keda.sh",
			},
			"spec": map[string]interface{}{
				"group": "eventing.keda.sh",
				"names": map[string]interface{}{
					"kind":     "ClusterCloudEventSource",
					"listKind": "ClusterCloudEventSourceList",
					"plural":   "clustercloudeventsources",
					"singular": "clustercloudeventsource",
				},
				"scope": "Cluster",
				"versions": []interface{}{
					map[string]interface{}{
						"additionalPrinterColumns": []interface{}{
							map[string]interface{}{
								"jsonPath": ".status.conditions[?(@.type==\"Active\")].status",
								"name":     "Active",
								"type":     "string",
							},
						},
						"name": "v1alpha1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"properties": map[string]interface{}{
									"apiVersion": map[string]interface{}{
										"description": `APIVersion defines the versioned schema of this representation of an object.
Servers should convert recognized schemas to the latest internal value, and
may reject unrecognized values.
More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources`,
										"type": "string",
									},
									"kind": map[string]interface{}{
										"description": `Kind is a string value representing the REST resource this object represents.
Servers may infer this from the endpoint the client submits requests to.
Cannot be updated.
In CamelCase.
More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds`,
										"type": "string",
									},
									"metadata": map[string]interface{}{
										"type": "object",
									},
									"spec": map[string]interface{}{
										"description": "CloudEventSourceSpec defines the spec of CloudEventSource",
										"properties": map[string]interface{}{
											"authenticationRef": map[string]interface{}{
												"description": `AuthenticationRef points to the TriggerAuthentication or ClusterTriggerAuthentication object that
is used to authenticate the scaler with the environment`,
												"properties": map[string]interface{}{
													"kind": map[string]interface{}{
														"description": "Kind of the resource being referred to. Defaults to TriggerAuthentication.",
														"enum": []interface{}{
															"TriggerAuthentication",
															"ClusterTriggerAuthentication",
														},
														"type": "string",
													},
													"name": map[string]interface{}{
														"type": "string",
													},
												},
												"required": []interface{}{
													"name",
												},
												"type": "object",
											},
											"clusterName": map[string]interface{}{
												"type": "string",
											},
											"destination": map[string]interface{}{
												"description": "Destination defines the various ways to emit events",
												"properties": map[string]interface{}{
													"azureEventGridTopic": map[string]interface{}{
														"properties": map[string]interface{}{
															"endpoint": map[string]interface{}{
																"type": "string",
															},
														},
														"required": []interface{}{
															"endpoint",
														},
														"type": "object",
													},
													"http": map[string]interface{}{
														"properties": map[string]interface{}{
															"uri": map[string]interface{}{
																"type": "string",
															},
														},
														"required": []interface{}{
															"uri",
														},
														"type": "object",
													},
												},
												"type": "object",
											},
											"eventSubscription": map[string]interface{}{
												"description": "EventSubscription defines filters for events",
												"properties": map[string]interface{}{
													"excludedEventTypes": map[string]interface{}{
														"items": map[string]interface{}{
															"enum": []interface{}{
																"keda.scaledobject.ready.v1",
																"keda.scaledobject.failed.v1",
																"keda.scaledobject.removed.v1",
																"keda.scaledobject.paused.v1",
																"keda.scaledobject.unpaused.v1",
																"keda.scaledjob.ready.v1",
																"keda.scaledjob.failed.v1",
																"keda.scaledjob.removed.v1",
																"keda.scaledjob.paused.v1",
																"keda.scaledjob.unpaused.v1",
																"keda.scaledjob.rolloutcleanup.started.v1",
																"keda.scaledjob.rolloutcleanup.completed.v1",
																"keda.scaledjob.rolloutcleanup.failed.v1",
																"keda.authentication.triggerauthentication.created.v1",
																"keda.authentication.triggerauthentication.updated.v1",
																"keda.authentication.triggerauthentication.removed.v1",
																"keda.authentication.clustertriggerauthentication.created.v1",
																"keda.authentication.clustertriggerauthentication.updated.v1",
																"keda.authentication.clustertriggerauthentication.removed.v1",
															},
															"type": "string",
														},
														"type": "array",
													},
													"includedEventTypes": map[string]interface{}{
														"items": map[string]interface{}{
															"enum": []interface{}{
																"keda.scaledobject.ready.v1",
																"keda.scaledobject.failed.v1",
																"keda.scaledobject.removed.v1",
																"keda.scaledobject.paused.v1",
																"keda.scaledobject.unpaused.v1",
																"keda.scaledjob.ready.v1",
																"keda.scaledjob.failed.v1",
																"keda.scaledjob.removed.v1",
																"keda.scaledjob.paused.v1",
																"keda.scaledjob.unpaused.v1",
																"keda.scaledjob.rolloutcleanup.started.v1",
																"keda.scaledjob.rolloutcleanup.completed.v1",
																"keda.scaledjob.rolloutcleanup.failed.v1",
																"keda.authentication.triggerauthentication.created.v1",
																"keda.authentication.triggerauthentication.updated.v1",
																"keda.authentication.triggerauthentication.removed.v1",
																"keda.authentication.clustertriggerauthentication.created.v1",
																"keda.authentication.clustertriggerauthentication.updated.v1",
																"keda.authentication.clustertriggerauthentication.removed.v1",
															},
															"type": "string",
														},
														"type": "array",
													},
												},
												"type": "object",
											},
										},
										"required": []interface{}{
											"destination",
										},
										"type": "object",
									},
									"status": map[string]interface{}{
										"description": "CloudEventSourceStatus defines the observed state of CloudEventSource",
										"properties": map[string]interface{}{
											"conditions": map[string]interface{}{
												"description": "Conditions an array representation to store multiple Conditions",
												"items": map[string]interface{}{
													"description": "Condition to store the condition state",
													"properties": map[string]interface{}{
														"message": map[string]interface{}{
															"description": "A human readable message indicating details about the transition.",
															"type":        "string",
														},
														"reason": map[string]interface{}{
															"description": "The reason for the condition's last transition.",
															"type":        "string",
														},
														"status": map[string]interface{}{
															"description": "Status of the condition, one of True, False, Unknown.",
															"type":        "string",
														},
														"type": map[string]interface{}{
															"description": "Type of condition",
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
										},
										"type": "object",
									},
								},
								"required": []interface{}{
									"spec",
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
			},
		},
	}

	return mutate.MutateCRDClustercloudeventsourcesEventingKedaSh(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;create;update;patch;delete

// CreateCRDClustertriggerauthenticationsKedaSh creates the CustomResourceDefinition resource with name clustertriggerauthentications.keda.sh.
func CreateCRDClustertriggerauthenticationsKedaSh(
	parent *platformv1alpha1.AutoscalingComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					"controller-gen.kubebuilder.io/version": "v0.21.0",
				},
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "keda-operator",
					"app.kubernetes.io/component":   "operator",
					"app.kubernetes.io/part-of":     "keda-operator",
					"app.kubernetes.io/version":     parent.Spec.Keda.Version, //  controlled by field: keda.version
					"platform.nukleros.io/category": "autoscaling",
					"platform.nukleros.io/project":  "keda",
				},
				"name": "clustertriggerauthentications.keda.sh",
			},
			"spec": map[string]interface{}{
				"group": "keda.sh",
				"names": map[string]interface{}{
					"kind":     "ClusterTriggerAuthentication",
					"listKind": "ClusterTriggerAuthenticationList",
					"plural":   "clustertriggerauthentications",
					"shortNames": []interface{}{
						"cta",
						"clustertriggerauth",
					},
					"singular": "clustertriggerauthentication",
				},
				"scope": "Cluster",
				"versions": []interface{}{
					map[string]interface{}{
						"additionalPrinterColumns": []interface{}{
							map[string]interface{}{
								"jsonPath": ".spec.podIdentity.provider",
								"name":     "PodIdentity",
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".spec.secretTargetRef[*].name",
								"name":     "Secret",
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".spec.env[*].name",
								"name":     "Env",
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".spec.filePath",
								"name":     "FilePath",
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".spec.hashiCorpVault.address",
								"name":     "VaultAddress",
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".status.scaledobjects",
								"name":     "ScaledObjects",
								"priority": 1,
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".status.scaledjobs",
								"name":     "ScaledJobs",
								"priority": 1,
								"type":     "string",
							},
						},
						"name": "v1alpha1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"description": "ClusterTriggerAuthentication defines how a trigger can authenticate globally",
								"properties": map[string]interface{}{
									"apiVersion": map[string]interface{}{
										"description": `APIVersion defines the versioned schema of this representation of an object.
Servers should convert recognized schemas to the latest internal value, and
may reject unrecognized values.
More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources`,
										"type": "string",
									},
									"kind": map[string]interface{}{
										"description": `Kind is a string value representing the REST resource this object represents.
Servers may infer this from the endpoint the client submits requests to.
Cannot be updated.
In CamelCase.
More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds`,
										"type": "string",
									},
									"metadata": map[string]interface{}{
										"type": "object",
									},
									"spec": map[string]interface{}{
										"description": "TriggerAuthenticationSpec defines the various ways to authenticate",
										"properties": map[string]interface{}{
											"awsSecretManager": map[string]interface{}{
												"description": "AwsSecretManager is used to authenticate using AwsSecretManager",
												"properties": map[string]interface{}{
													"credentials": map[string]interface{}{
														"properties": map[string]interface{}{
															"accessKey": map[string]interface{}{
																"properties": map[string]interface{}{
																	"valueFrom": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"secretKeyRef": map[string]interface{}{
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"type": "string",
																					},
																					"name": map[string]interface{}{
																						"type": "string",
																					},
																				},
																				"required": []interface{}{
																					"key",
																					"name",
																				},
																				"type": "object",
																			},
																		},
																		"required": []interface{}{
																			"secretKeyRef",
																		},
																		"type": "object",
																	},
																},
																"required": []interface{}{
																	"valueFrom",
																},
																"type": "object",
															},
															"accessSecretKey": map[string]interface{}{
																"properties": map[string]interface{}{
																	"valueFrom": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"secretKeyRef": map[string]interface{}{
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"type": "string",
																					},
																					"name": map[string]interface{}{
																						"type": "string",
																					},
																				},
																				"required": []interface{}{
																					"key",
																					"name",
																				},
																				"type": "object",
																			},
																		},
																		"required": []interface{}{
																			"secretKeyRef",
																		},
																		"type": "object",
																	},
																},
																"required": []interface{}{
																	"valueFrom",
																},
																"type": "object",
															},
															"accessToken": map[string]interface{}{
																"properties": map[string]interface{}{
																	"valueFrom": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"secretKeyRef": map[string]interface{}{
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"type": "string",
																					},
																					"name": map[string]interface{}{
																						"type": "string",
																					},
																				},
																				"required": []interface{}{
																					"key",
																					"name",
																				},
																				"type": "object",
																			},
																		},
																		"required": []interface{}{
																			"secretKeyRef",
																		},
																		"type": "object",
																	},
																},
																"required": []interface{}{
																	"valueFrom",
																},
																"type": "object",
															},
														},
														"required": []interface{}{
															"accessKey",
															"accessSecretKey",
														},
														"type": "object",
													},
													"podIdentity": map[string]interface{}{
														"description": `AuthPodIdentity allows users to select the platform native identity
mechanism`,
														"properties": map[string]interface{}{
															"externalID": map[string]interface{}{
																"description": "ExternalID sets the External ID to be used when assuming an identity. This is only applicable when using AWS pod identity with a RoleArn.",
																"type":        "string",
															},
															"identityAuthorityHost": map[string]interface{}{
																"description": "Set identityAuthorityHost to override the default Azure authority host. If this is set, then the IdentityTenantID must also be set",
																"type":        "string",
															},
															"identityId": map[string]interface{}{
																"type": "string",
															},
															"identityOwner": map[string]interface{}{
																"description": "IdentityOwner configures which identity has to be used during auto discovery, keda or the scaled workload. Mutually exclusive with roleArn",
																"enum": []interface{}{
																	"keda",
																	"workload",
																},
																"type": "string",
															},
															"identityTenantId": map[string]interface{}{
																"description": "Set identityTenantId to override the default Azure tenant id. If this is set, then the IdentityID must also be set",
																"type":        "string",
															},
															"provider": map[string]interface{}{
																"description": "PodIdentityProvider contains the list of providers",
																"enum": []interface{}{
																	"azure-workload",
																	"gcp",
																	"aws",
																	"aws-eks",
																	"none",
																},
																"type": "string",
															},
															"roleArn": map[string]interface{}{
																"description": "RoleArn sets the AWS RoleArn to be used. Mutually exclusive with IdentityOwner",
																"type":        "string",
															},
														},
														"required": []interface{}{
															"provider",
														},
														"type": "object",
													},
													"region": map[string]interface{}{
														"type": "string",
													},
													"secrets": map[string]interface{}{
														"items": map[string]interface{}{
															"properties": map[string]interface{}{
																"name": map[string]interface{}{
																	"type": "string",
																},
																"parameter": map[string]interface{}{
																	"type": "string",
																},
																"secretKey": map[string]interface{}{
																	"type": "string",
																},
																"versionId": map[string]interface{}{
																	"type": "string",
																},
																"versionStage": map[string]interface{}{
																	"type": "string",
																},
															},
															"required": []interface{}{
																"name",
																"parameter",
															},
															"type": "object",
														},
														"minItems": 1,
														"type":     "array",
													},
												},
												"required": []interface{}{
													"secrets",
												},
												"type": "object",
											},
											"azureKeyVault": map[string]interface{}{
												"description": "AzureKeyVault is used to authenticate using Azure Key Vault",
												"properties": map[string]interface{}{
													"cloud": map[string]interface{}{
														"properties": map[string]interface{}{
															"activeDirectoryEndpoint": map[string]interface{}{
																"type": "string",
															},
															"keyVaultResourceURL": map[string]interface{}{
																"type": "string",
															},
															"type": map[string]interface{}{
																"type": "string",
															},
														},
														"required": []interface{}{
															"type",
														},
														"type": "object",
													},
													"credentials": map[string]interface{}{
														"properties": map[string]interface{}{
															"clientId": map[string]interface{}{
																"type": "string",
															},
															"clientSecret": map[string]interface{}{
																"properties": map[string]interface{}{
																	"valueFrom": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"secretKeyRef": map[string]interface{}{
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"type": "string",
																					},
																					"name": map[string]interface{}{
																						"type": "string",
																					},
																				},
																				"required": []interface{}{
																					"key",
																					"name",
																				},
																				"type": "object",
																			},
																		},
																		"required": []interface{}{
																			"secretKeyRef",
																		},
																		"type": "object",
																	},
																},
																"required": []interface{}{
																	"valueFrom",
																},
																"type": "object",
															},
															"tenantId": map[string]interface{}{
																"type": "string",
															},
														},
														"required": []interface{}{
															"clientId",
															"clientSecret",
															"tenantId",
														},
														"type": "object",
													},
													"podIdentity": map[string]interface{}{
														"description": `AuthPodIdentity allows users to select the platform native identity
mechanism`,
														"properties": map[string]interface{}{
															"externalID": map[string]interface{}{
																"description": "ExternalID sets the External ID to be used when assuming an identity. This is only applicable when using AWS pod identity with a RoleArn.",
																"type":        "string",
															},
															"identityAuthorityHost": map[string]interface{}{
																"description": "Set identityAuthorityHost to override the default Azure authority host. If this is set, then the IdentityTenantID must also be set",
																"type":        "string",
															},
															"identityId": map[string]interface{}{
																"type": "string",
															},
															"identityOwner": map[string]interface{}{
																"description": "IdentityOwner configures which identity has to be used during auto discovery, keda or the scaled workload. Mutually exclusive with roleArn",
																"enum": []interface{}{
																	"keda",
																	"workload",
																},
																"type": "string",
															},
															"identityTenantId": map[string]interface{}{
																"description": "Set identityTenantId to override the default Azure tenant id. If this is set, then the IdentityID must also be set",
																"type":        "string",
															},
															"provider": map[string]interface{}{
																"description": "PodIdentityProvider contains the list of providers",
																"enum": []interface{}{
																	"azure-workload",
																	"gcp",
																	"aws",
																	"aws-eks",
																	"none",
																},
																"type": "string",
															},
															"roleArn": map[string]interface{}{
																"description": "RoleArn sets the AWS RoleArn to be used. Mutually exclusive with IdentityOwner",
																"type":        "string",
															},
														},
														"required": []interface{}{
															"provider",
														},
														"type": "object",
													},
													"secrets": map[string]interface{}{
														"items": map[string]interface{}{
															"properties": map[string]interface{}{
																"name": map[string]interface{}{
																	"type": "string",
																},
																"parameter": map[string]interface{}{
																	"type": "string",
																},
																"version": map[string]interface{}{
																	"type": "string",
																},
															},
															"required": []interface{}{
																"name",
																"parameter",
															},
															"type": "object",
														},
														"minItems": 1,
														"type":     "array",
													},
													"vaultUri": map[string]interface{}{
														"type": "string",
													},
												},
												"required": []interface{}{
													"secrets",
													"vaultUri",
												},
												"type": "object",
											},
											"azureServicePrincipal": map[string]interface{}{
												"description": `AzureServicePrincipal defines Microsoft Entra service principal credentials
that can be shared by Azure scalers.`,
												"properties": map[string]interface{}{
													"activeDirectoryEndpoint": map[string]interface{}{
														"description": "ActiveDirectoryEndpoint is required when cloud is set to Private.",
														"type":        "string",
													},
													"clientCertificate": map[string]interface{}{
														"description": `AzureServicePrincipalCredential references credential data stored in a
Kubernetes Secret.`,
														"properties": map[string]interface{}{
															"valueFrom": map[string]interface{}{
																"properties": map[string]interface{}{
																	"secretKeyRef": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"type": "string",
																			},
																			"name": map[string]interface{}{
																				"type": "string",
																			},
																		},
																		"required": []interface{}{
																			"key",
																			"name",
																		},
																		"type": "object",
																	},
																},
																"required": []interface{}{
																	"secretKeyRef",
																},
																"type": "object",
															},
														},
														"required": []interface{}{
															"valueFrom",
														},
														"type": "object",
													},
													"clientCertificatePassword": map[string]interface{}{
														"description": `AzureServicePrincipalCredential references credential data stored in a
Kubernetes Secret.`,
														"properties": map[string]interface{}{
															"valueFrom": map[string]interface{}{
																"properties": map[string]interface{}{
																	"secretKeyRef": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"type": "string",
																			},
																			"name": map[string]interface{}{
																				"type": "string",
																			},
																		},
																		"required": []interface{}{
																			"key",
																			"name",
																		},
																		"type": "object",
																	},
																},
																"required": []interface{}{
																	"secretKeyRef",
																},
																"type": "object",
															},
														},
														"required": []interface{}{
															"valueFrom",
														},
														"type": "object",
													},
													"clientId": map[string]interface{}{
														"type": "string",
													},
													"clientSecret": map[string]interface{}{
														"description": `AzureServicePrincipalCredential references credential data stored in a
Kubernetes Secret.`,
														"properties": map[string]interface{}{
															"valueFrom": map[string]interface{}{
																"properties": map[string]interface{}{
																	"secretKeyRef": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"type": "string",
																			},
																			"name": map[string]interface{}{
																				"type": "string",
																			},
																		},
																		"required": []interface{}{
																			"key",
																			"name",
																		},
																		"type": "object",
																	},
																},
																"required": []interface{}{
																	"secretKeyRef",
																},
																"type": "object",
															},
														},
														"required": []interface{}{
															"valueFrom",
														},
														"type": "object",
													},
													"cloud": map[string]interface{}{
														"description": "Cloud specifies the Azure cloud environment. Defaults to AzurePublicCloud.",
														"type":        "string",
													},
													"tenantId": map[string]interface{}{
														"type": "string",
													},
												},
												"required": []interface{}{
													"clientId",
													"tenantId",
												},
												"type": "object",
											},
											"boundServiceAccountToken": map[string]interface{}{
												"items": map[string]interface{}{
													"properties": map[string]interface{}{
														"parameter": map[string]interface{}{
															"type": "string",
														},
														"serviceAccountName": map[string]interface{}{
															"type": "string",
														},
													},
													"required": []interface{}{
														"parameter",
														"serviceAccountName",
													},
													"type": "object",
												},
												"type": "array",
											},
											"configMapTargetRef": map[string]interface{}{
												"items": map[string]interface{}{
													"description": "AuthConfigMapTargetRef is used to authenticate using a reference to a config map",
													"properties": map[string]interface{}{
														"key": map[string]interface{}{
															"type": "string",
														},
														"name": map[string]interface{}{
															"type": "string",
														},
														"parameter": map[string]interface{}{
															"type": "string",
														},
													},
													"required": []interface{}{
														"key",
														"name",
														"parameter",
													},
													"type": "object",
												},
												"type": "array",
											},
											"env": map[string]interface{}{
												"items": map[string]interface{}{
													"description": `AuthEnvironment is used to authenticate using environment variables
in the destination ScaleTarget spec`,
													"properties": map[string]interface{}{
														"containerName": map[string]interface{}{
															"type": "string",
														},
														"name": map[string]interface{}{
															"type": "string",
														},
														"parameter": map[string]interface{}{
															"type": "string",
														},
													},
													"required": []interface{}{
														"name",
														"parameter",
													},
													"type": "object",
												},
												"type": "array",
											},
											"filePath": map[string]interface{}{
												"description": `FilePath specifies a file containing auth parameters as JSON map[string]string.
When set, auth params are read directly from this file instead.`,
												"type": "string",
											},
											"gcpSecretManager": map[string]interface{}{
												"properties": map[string]interface{}{
													"credentials": map[string]interface{}{
														"properties": map[string]interface{}{
															"clientSecret": map[string]interface{}{
																"properties": map[string]interface{}{
																	"valueFrom": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"secretKeyRef": map[string]interface{}{
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"type": "string",
																					},
																					"name": map[string]interface{}{
																						"type": "string",
																					},
																				},
																				"required": []interface{}{
																					"key",
																					"name",
																				},
																				"type": "object",
																			},
																		},
																		"required": []interface{}{
																			"secretKeyRef",
																		},
																		"type": "object",
																	},
																},
																"required": []interface{}{
																	"valueFrom",
																},
																"type": "object",
															},
														},
														"required": []interface{}{
															"clientSecret",
														},
														"type": "object",
													},
													"podIdentity": map[string]interface{}{
														"description": `AuthPodIdentity allows users to select the platform native identity
mechanism`,
														"properties": map[string]interface{}{
															"externalID": map[string]interface{}{
																"description": "ExternalID sets the External ID to be used when assuming an identity. This is only applicable when using AWS pod identity with a RoleArn.",
																"type":        "string",
															},
															"identityAuthorityHost": map[string]interface{}{
																"description": "Set identityAuthorityHost to override the default Azure authority host. If this is set, then the IdentityTenantID must also be set",
																"type":        "string",
															},
															"identityId": map[string]interface{}{
																"type": "string",
															},
															"identityOwner": map[string]interface{}{
																"description": "IdentityOwner configures which identity has to be used during auto discovery, keda or the scaled workload. Mutually exclusive with roleArn",
																"enum": []interface{}{
																	"keda",
																	"workload",
																},
																"type": "string",
															},
															"identityTenantId": map[string]interface{}{
																"description": "Set identityTenantId to override the default Azure tenant id. If this is set, then the IdentityID must also be set",
																"type":        "string",
															},
															"provider": map[string]interface{}{
																"description": "PodIdentityProvider contains the list of providers",
																"enum": []interface{}{
																	"azure-workload",
																	"gcp",
																	"aws",
																	"aws-eks",
																	"none",
																},
																"type": "string",
															},
															"roleArn": map[string]interface{}{
																"description": "RoleArn sets the AWS RoleArn to be used. Mutually exclusive with IdentityOwner",
																"type":        "string",
															},
														},
														"required": []interface{}{
															"provider",
														},
														"type": "object",
													},
													"secrets": map[string]interface{}{
														"items": map[string]interface{}{
															"properties": map[string]interface{}{
																"id": map[string]interface{}{
																	"type": "string",
																},
																"parameter": map[string]interface{}{
																	"type": "string",
																},
																"version": map[string]interface{}{
																	"type": "string",
																},
															},
															"required": []interface{}{
																"id",
																"parameter",
															},
															"type": "object",
														},
														"minItems": 1,
														"type":     "array",
													},
												},
												"required": []interface{}{
													"secrets",
												},
												"type": "object",
											},
											"hashiCorpVault": map[string]interface{}{
												"description": "HashiCorpVault is used to authenticate using Hashicorp Vault",
												"properties": map[string]interface{}{
													"address": map[string]interface{}{
														"type": "string",
													},
													"authentication": map[string]interface{}{
														"description": "VaultAuthentication contains the list of Hashicorp Vault authentication methods",
														"enum": []interface{}{
															"token",
															"kubernetes",
														},
														"type": "string",
													},
													"credential": map[string]interface{}{
														"description": "Credential defines the Hashicorp Vault credentials depending on the authentication method",
														"properties": map[string]interface{}{
															"serviceAccount": map[string]interface{}{
																"type": "string",
															},
															"serviceAccountName": map[string]interface{}{
																"type": "string",
															},
															"token": map[string]interface{}{
																"type": "string",
															},
														},
														"type": "object",
													},
													"mount": map[string]interface{}{
														"type": "string",
													},
													"namespace": map[string]interface{}{
														"type": "string",
													},
													"role": map[string]interface{}{
														"type": "string",
													},
													"secrets": map[string]interface{}{
														"items": map[string]interface{}{
															"description": "VaultSecret defines the mapping between the path of the secret in Vault to the parameter",
															"properties": map[string]interface{}{
																"key": map[string]interface{}{
																	"type": "string",
																},
																"parameter": map[string]interface{}{
																	"type": "string",
																},
																"path": map[string]interface{}{
																	"type": "string",
																},
																"pkiData": map[string]interface{}{
																	"properties": map[string]interface{}{
																		"altNames": map[string]interface{}{
																			"type": "string",
																		},
																		"commonName": map[string]interface{}{
																			"type": "string",
																		},
																		"format": map[string]interface{}{
																			"type": "string",
																		},
																		"ipSans": map[string]interface{}{
																			"type": "string",
																		},
																		"otherSans": map[string]interface{}{
																			"type": "string",
																		},
																		"ttl": map[string]interface{}{
																			"type": "string",
																		},
																		"uriSans": map[string]interface{}{
																			"type": "string",
																		},
																	},
																	"type": "object",
																},
																"type": map[string]interface{}{
																	"description": "VaultSecretType defines the type of vault secret",
																	"enum": []interface{}{
																		"",
																		"secretV2",
																		"secret",
																		"pki",
																	},
																	"type": "string",
																},
															},
															"required": []interface{}{
																"key",
																"parameter",
																"path",
															},
															"type": "object",
														},
														"minItems": 1,
														"type":     "array",
													},
												},
												"required": []interface{}{
													"address",
													"authentication",
													"secrets",
												},
												"type": "object",
											},
											"oauth2": map[string]interface{}{
												"properties": map[string]interface{}{
													"clientId": map[string]interface{}{
														"type": "string",
													},
													"clientSecret": map[string]interface{}{
														"properties": map[string]interface{}{
															"valueFrom": map[string]interface{}{
																"properties": map[string]interface{}{
																	"secretKeyRef": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"type": "string",
																			},
																			"name": map[string]interface{}{
																				"type": "string",
																			},
																		},
																		"required": []interface{}{
																			"key",
																			"name",
																		},
																		"type": "object",
																	},
																},
																"required": []interface{}{
																	"secretKeyRef",
																},
																"type": "object",
															},
														},
														"required": []interface{}{
															"valueFrom",
														},
														"type": "object",
													},
													"scopes": map[string]interface{}{
														"items": map[string]interface{}{
															"type": "string",
														},
														"type": "array",
													},
													"tokenUrl": map[string]interface{}{
														"type": "string",
													},
													"tokenUrlParams": map[string]interface{}{
														"additionalProperties": map[string]interface{}{
															"type": "string",
														},
														"type": "object",
													},
													"type": map[string]interface{}{
														"default": "clientCredentials",
														"enum": []interface{}{
															"clientCredentials",
														},
														"type": "string",
													},
												},
												"required": []interface{}{
													"clientId",
													"tokenUrl",
													"type",
												},
												"type": "object",
											},
											"podIdentity": map[string]interface{}{
												"description": `AuthPodIdentity allows users to select the platform native identity
mechanism`,
												"properties": map[string]interface{}{
													"externalID": map[string]interface{}{
														"description": "ExternalID sets the External ID to be used when assuming an identity. This is only applicable when using AWS pod identity with a RoleArn.",
														"type":        "string",
													},
													"identityAuthorityHost": map[string]interface{}{
														"description": "Set identityAuthorityHost to override the default Azure authority host. If this is set, then the IdentityTenantID must also be set",
														"type":        "string",
													},
													"identityId": map[string]interface{}{
														"type": "string",
													},
													"identityOwner": map[string]interface{}{
														"description": "IdentityOwner configures which identity has to be used during auto discovery, keda or the scaled workload. Mutually exclusive with roleArn",
														"enum": []interface{}{
															"keda",
															"workload",
														},
														"type": "string",
													},
													"identityTenantId": map[string]interface{}{
														"description": "Set identityTenantId to override the default Azure tenant id. If this is set, then the IdentityID must also be set",
														"type":        "string",
													},
													"provider": map[string]interface{}{
														"description": "PodIdentityProvider contains the list of providers",
														"enum": []interface{}{
															"azure-workload",
															"gcp",
															"aws",
															"aws-eks",
															"none",
														},
														"type": "string",
													},
													"roleArn": map[string]interface{}{
														"description": "RoleArn sets the AWS RoleArn to be used. Mutually exclusive with IdentityOwner",
														"type":        "string",
													},
												},
												"required": []interface{}{
													"provider",
												},
												"type": "object",
											},
											"secretTargetRef": map[string]interface{}{
												"items": map[string]interface{}{
													"description": "AuthSecretTargetRef is used to authenticate using a reference to a secret",
													"properties": map[string]interface{}{
														"key": map[string]interface{}{
															"type": "string",
														},
														"name": map[string]interface{}{
															"type": "string",
														},
														"parameter": map[string]interface{}{
															"type": "string",
														},
													},
													"required": []interface{}{
														"key",
														"name",
														"parameter",
													},
													"type": "object",
												},
												"type": "array",
											},
										},
										"type": "object",
									},
									"status": map[string]interface{}{
										"description": "TriggerAuthenticationStatus defines the observed state of TriggerAuthentication",
										"properties": map[string]interface{}{
											"scaledjobs": map[string]interface{}{
												"type": "string",
											},
											"scaledobjects": map[string]interface{}{
												"type": "string",
											},
										},
										"type": "object",
									},
								},
								"required": []interface{}{
									"spec",
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
			},
		},
	}

	return mutate.MutateCRDClustertriggerauthenticationsKedaSh(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;create;update;patch;delete

// CreateCRDScaledjobsKedaSh creates the CustomResourceDefinition resource with name scaledjobs.keda.sh.
func CreateCRDScaledjobsKedaSh(
	parent *platformv1alpha1.AutoscalingComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					"controller-gen.kubebuilder.io/version": "v0.21.0",
				},
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "keda-operator",
					"app.kubernetes.io/component":   "operator",
					"app.kubernetes.io/part-of":     "keda-operator",
					"app.kubernetes.io/version":     parent.Spec.Keda.Version, //  controlled by field: keda.version
					"platform.nukleros.io/category": "autoscaling",
					"platform.nukleros.io/project":  "keda",
				},
				"name": "scaledjobs.keda.sh",
			},
			"spec": map[string]interface{}{
				"group": "keda.sh",
				"names": map[string]interface{}{
					"kind":     "ScaledJob",
					"listKind": "ScaledJobList",
					"plural":   "scaledjobs",
					"shortNames": []interface{}{
						"sj",
					},
					"singular": "scaledjob",
				},
				"scope": "Namespaced",
				"versions": []interface{}{
					map[string]interface{}{
						"additionalPrinterColumns": []interface{}{
							map[string]interface{}{
								"jsonPath": ".spec.minReplicaCount",
								"name":     "Min",
								"type":     "integer",
							},
							map[string]interface{}{
								"jsonPath": ".spec.maxReplicaCount",
								"name":     "Max",
								"type":     "integer",
							},
							map[string]interface{}{
								"jsonPath": ".status.conditions[?(@.type==\"Ready\")].status",
								"name":     "Ready",
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".status.conditions[?(@.type==\"Active\")].status",
								"name":     "Active",
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".status.conditions[?(@.type==\"Paused\")].status",
								"name":     "Paused",
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".status.triggersTypes",
								"name":     "Triggers",
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".status.authenticationsTypes",
								"name":     "Authentications",
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".metadata.creationTimestamp",
								"name":     "Age",
								"type":     "date",
							},
						},
						"name": "v1alpha1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"description": "ScaledJob is the Schema for the scaledjobs API",
								"properties": map[string]interface{}{
									"apiVersion": map[string]interface{}{
										"description": `APIVersion defines the versioned schema of this representation of an object.
Servers should convert recognized schemas to the latest internal value, and
may reject unrecognized values.
More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources`,
										"type": "string",
									},
									"kind": map[string]interface{}{
										"description": `Kind is a string value representing the REST resource this object represents.
Servers may infer this from the endpoint the client submits requests to.
Cannot be updated.
In CamelCase.
More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds`,
										"type": "string",
									},
									"metadata": map[string]interface{}{
										"type": "object",
									},
									"spec": map[string]interface{}{
										"description": "ScaledJobSpec defines the desired state of ScaledJob",
										"properties": map[string]interface{}{
											"envSourceContainerName": map[string]interface{}{
												"type": "string",
											},
											"failedJobsHistoryLimit": map[string]interface{}{
												"format":  "int32",
												"minimum": 0,
												"type":    "integer",
											},
											"jobTargetRef": map[string]interface{}{
												"description": "JobSpec describes how the job execution will look like.",
												"properties": map[string]interface{}{
													"activeDeadlineSeconds": map[string]interface{}{
														"description": `Specifies the duration in seconds relative to the startTime that the job
may be continuously active before the system tries to terminate it; value
must be positive integer. If a Job is suspended (at creation or through an
update), this timer will effectively be stopped and reset when the Job is
resumed again.`,
														"format": "int64",
														"type":   "integer",
													},
													"backoffLimit": map[string]interface{}{
														"description": `Specifies the number of retries before marking this job failed.
Defaults to 6, unless backoffLimitPerIndex (only Indexed Job) is specified.
When backoffLimitPerIndex is specified, backoffLimit defaults to 2147483647.`,
														"format": "int32",
														"type":   "integer",
													},
													"backoffLimitPerIndex": map[string]interface{}{
														"description": `Specifies the limit for the number of retries within an
index before marking this index as failed. When enabled the number of
failures per index is kept in the pod's
batch.kubernetes.io/job-index-failure-count annotation. It can only
be set when Job's completionMode=Indexed, and the Pod's restart
policy is Never. The field is immutable.`,
														"format": "int32",
														"type":   "integer",
													},
													"completionMode": map[string]interface{}{
														"description": `completionMode specifies how Pod completions are tracked. It can be
` + "`" + `NonIndexed` + "`" + ` (default) or ` + "`" + `Indexed` + "`" + `.

` + "`" + `NonIndexed` + "`" + ` means that the Job is considered complete when there have
been .spec.completions successfully completed Pods. Each Pod completion is
homologous to each other.

` + "`" + `Indexed` + "`" + ` means that the Pods of a
Job get an associated completion index from 0 to (.spec.completions - 1),
available in the annotation batch.kubernetes.io/job-completion-index.
The Job is considered complete when there is one successfully completed Pod
for each index.
When value is ` + "`" + `Indexed` + "`" + `, .spec.completions must be specified and
` + "`" + `.spec.parallelism` + "`" + ` must be less than or equal to 10^5.
In addition, The Pod name takes the form
` + "`" + `$(job-name)-$(index)-$(random-string)` + "`" + `,
the Pod hostname takes the form ` + "`" + `$(job-name)-$(index)` + "`" + `.

More completion modes can be added in the future.
If the Job controller observes a mode that it doesn't recognize, which
is possible during upgrades due to version skew, the controller
skips updates for the Job.`,
														"type": "string",
													},
													"completions": map[string]interface{}{
														"description": `Specifies the desired number of successfully finished pods the
job should be run with.  Setting to null means that the success of any
pod signals the success of all pods, and allows parallelism to have any positive
value.  Setting to 1 means that parallelism is limited to 1 and the success of that
pod signals the success of the job.
More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/`,
														"format": "int32",
														"type":   "integer",
													},
													"managedBy": map[string]interface{}{
														"description": `ManagedBy field indicates the controller that manages a Job. The k8s Job
controller reconciles jobs which don't have this field at all or the field
value is the reserved string ` + "`" + `kubernetes.io/job-controller` + "`" + `, but skips
reconciling Jobs with a custom value for this field.
The value must be a valid domain-prefixed path (e.g. acme.io/foo) -
all characters before the first "/" must be a valid subdomain as defined
by RFC 1123. All characters trailing the first "/" must be valid HTTP Path
characters as defined by RFC 3986. The value cannot exceed 63 characters.
This field is immutable.`,
														"type": "string",
													},
													"manualSelector": map[string]interface{}{
														"description": `manualSelector controls generation of pod labels and pod selectors.
Leave ` + "`" + `manualSelector` + "`" + ` unset unless you are certain what you are doing.
When false or unset, the system pick labels unique to this job
and appends those labels to the pod template.  When true,
the user is responsible for picking unique labels and specifying
the selector.  Failure to pick a unique label may cause this
and other jobs to not function correctly.  However, You may see
` + "`" + `manualSelector=true` + "`" + ` in jobs that were created with the old ` + "`" + `extensions/v1beta1` + "`" + `
API.
More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/#specifying-your-own-pod-selector`,
														"type": "boolean",
													},
													"maxFailedIndexes": map[string]interface{}{
														"description": `Specifies the maximal number of failed indexes before marking the Job as
failed, when backoffLimitPerIndex is set. Once the number of failed
indexes exceeds this number the entire Job is marked as Failed and its
execution is terminated. When left as null the job continues execution of
all of its indexes and is marked with the ` + "`" + `Complete` + "`" + ` Job condition.
It can only be specified when backoffLimitPerIndex is set.
It can be null or up to completions. It is required and must be
less than or equal to 10^4 when is completions greater than 10^5.`,
														"format": "int32",
														"type":   "integer",
													},
													"parallelism": map[string]interface{}{
														"description": `Specifies the maximum desired number of pods the job should
run at any given time. The actual number of pods running in steady state will
be less than this number when ((.spec.completions - .status.successful) < .spec.parallelism),
i.e. when the work left to do is less than max parallelism.
More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/`,
														"format": "int32",
														"type":   "integer",
													},
													"podFailurePolicy": map[string]interface{}{
														"description": `Specifies the policy of handling failed pods. In particular, it allows to
specify the set of actions and conditions which need to be
satisfied to take the associated action.
If empty, the default behaviour applies - the counter of failed pods,
represented by the jobs's .status.failed field, is incremented and it is
checked against the backoffLimit. This field cannot be used in combination
with restartPolicy=OnFailure.`,
														"properties": map[string]interface{}{
															"rules": map[string]interface{}{
																"description": `A list of pod failure policy rules. The rules are evaluated in order.
Once a rule matches a Pod failure, the remaining of the rules are ignored.
When no rule matches the Pod failure, the default handling applies - the
counter of pod failures is incremented and it is checked against
the backoffLimit. At most 20 elements are allowed.`,
																"items": map[string]interface{}{
																	"description": `PodFailurePolicyRule describes how a pod failure is handled when the requirements are met.
One of onExitCodes and onPodConditions, but not both, can be used in each rule.`,
																	"properties": map[string]interface{}{
																		"action": map[string]interface{}{
																			"description": `Specifies the action taken on a pod failure when the requirements are satisfied.
Possible values are:

- FailJob: indicates that the pod's job is marked as Failed and all
  running pods are terminated.
- FailIndex: indicates that the pod's index is marked as Failed and will
  not be restarted.
- Ignore: indicates that the counter towards the .backoffLimit is not
  incremented and a replacement pod is created.
- Count: indicates that the pod is handled in the default way - the
  counter towards the .backoffLimit is incremented.
Additional values are considered to be added in the future. Clients should
react to an unknown action by skipping the rule.`,
																			"type": "string",
																		},
																		"onExitCodes": map[string]interface{}{
																			"description": "Represents the requirement on the container exit codes.",
																			"properties": map[string]interface{}{
																				"containerName": map[string]interface{}{
																					"description": `Restricts the check for exit codes to the container with the
specified name. When null, the rule applies to all containers.
When specified, it should match one the container or initContainer
names in the pod template.`,
																					"type": "string",
																				},
																				"operator": map[string]interface{}{
																					"description": `Represents the relationship between the container exit code(s) and the
specified values. Containers completed with success (exit code 0) are
excluded from the requirement check. Possible values are:

- In: the requirement is satisfied if at least one container exit code
  (might be multiple if there are multiple containers not restricted
  by the 'containerName' field) is in the set of specified values.
- NotIn: the requirement is satisfied if at least one container exit code
  (might be multiple if there are multiple containers not restricted
  by the 'containerName' field) is not in the set of specified values.
Additional values are considered to be added in the future. Clients should
react to an unknown operator by assuming the requirement is not satisfied.`,
																					"type": "string",
																				},
																				"values": map[string]interface{}{
																					"description": `Specifies the set of values. Each returned container exit code (might be
multiple in case of multiple containers) is checked against this set of
values with respect to the operator. The list of values must be ordered
and must not contain duplicates. Value '0' cannot be used for the In operator.
At least one element is required. At most 255 elements are allowed.`,
																					"items": map[string]interface{}{
																						"format": "int32",
																						"type":   "integer",
																					},
																					"type":                   "array",
																					"x-kubernetes-list-type": "set",
																				},
																			},
																			"required": []interface{}{
																				"operator",
																				"values",
																			},
																			"type": "object",
																		},
																		"onPodConditions": map[string]interface{}{
																			"description": `Represents the requirement on the pod conditions. The requirement is represented
as a list of pod condition patterns. The requirement is satisfied if at
least one pattern matches an actual pod condition. At most 20 elements are allowed.`,
																			"items": map[string]interface{}{
																				"description": `PodFailurePolicyOnPodConditionsPattern describes a pattern for matching
an actual pod condition type.`,
																				"properties": map[string]interface{}{
																					"status": map[string]interface{}{
																						"description": `Specifies the required Pod condition status. To match a pod condition
it is required that the specified status equals the pod condition status.
Defaults to True.`,
																						"type": "string",
																					},
																					"type": map[string]interface{}{
																						"description": `Specifies the required Pod condition type. To match a pod condition
it is required that specified type equals the pod condition type.`,
																						"type": "string",
																					},
																				},
																				"required": []interface{}{
																					"type",
																				},
																				"type": "object",
																			},
																			"type":                   "array",
																			"x-kubernetes-list-type": "atomic",
																		},
																	},
																	"required": []interface{}{
																		"action",
																	},
																	"type": "object",
																},
																"type":                   "array",
																"x-kubernetes-list-type": "atomic",
															},
														},
														"required": []interface{}{
															"rules",
														},
														"type": "object",
													},
													"podReplacementPolicy": map[string]interface{}{
														"description": `podReplacementPolicy specifies when to create replacement Pods.
Possible values are:
- TerminatingOrFailed means that we recreate pods
  when they are terminating (has a metadata.deletionTimestamp) or failed.
- Failed means to wait until a previously created Pod is fully terminated (has phase
  Failed or Succeeded) before creating a replacement Pod.

When using podFailurePolicy, Failed is the the only allowed value.
TerminatingOrFailed and Failed are allowed values when podFailurePolicy is not in use.`,
														"type": "string",
													},
													"selector": map[string]interface{}{
														"description": `A label query over pods that should match the pod count.
Normally, the system sets this field for you.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors`,
														"properties": map[string]interface{}{
															"matchExpressions": map[string]interface{}{
																"description": "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																"items": map[string]interface{}{
																	"description": `A label selector requirement is a selector that contains values, a key, and an operator that
relates the key and values.`,
																	"properties": map[string]interface{}{
																		"key": map[string]interface{}{
																			"description": "key is the label key that the selector applies to.",
																			"type":        "string",
																		},
																		"operator": map[string]interface{}{
																			"description": `operator represents a key's relationship to a set of values.
Valid operators are In, NotIn, Exists and DoesNotExist.`,
																			"type": "string",
																		},
																		"values": map[string]interface{}{
																			"description": `values is an array of string values. If the operator is In or NotIn,
the values array must be non-empty. If the operator is Exists or DoesNotExist,
the values array must be empty. This array is replaced during a strategic
merge patch.`,
																			"items": map[string]interface{}{
																				"type": "string",
																			},
																			"type":                   "array",
																			"x-kubernetes-list-type": "atomic",
																		},
																	},
																	"required": []interface{}{
																		"key",
																		"operator",
																	},
																	"type": "object",
																},
																"type":                   "array",
																"x-kubernetes-list-type": "atomic",
															},
															"matchLabels": map[string]interface{}{
																"additionalProperties": map[string]interface{}{
																	"type": "string",
																},
																"description": `matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels
map is equivalent to an element of matchExpressions, whose key field is "key", the
operator is "In", and the values array contains only "value". The requirements are ANDed.`,
																"type": "object",
															},
														},
														"type":                  "object",
														"x-kubernetes-map-type": "atomic",
													},
													"successPolicy": map[string]interface{}{
														"description": `successPolicy specifies the policy when the Job can be declared as succeeded.
If empty, the default behavior applies - the Job is declared as succeeded
only when the number of succeeded pods equals to the completions.
When the field is specified, it must be immutable and works only for the Indexed Jobs.
Once the Job meets the SuccessPolicy, the lingering pods are terminated.`,
														"properties": map[string]interface{}{
															"rules": map[string]interface{}{
																"description": `rules represents the list of alternative rules for the declaring the Jobs
as successful before ` + "`" + `.status.succeeded >= .spec.completions` + "`" + `. Once any of the rules are met,
the "SuccessCriteriaMet" condition is added, and the lingering pods are removed.
The terminal state for such a Job has the "Complete" condition.
Additionally, these rules are evaluated in order; Once the Job meets one of the rules,
other rules are ignored. At most 20 elements are allowed.`,
																"items": map[string]interface{}{
																	"description": `SuccessPolicyRule describes rule for declaring a Job as succeeded.
Each rule must have at least one of the "succeededIndexes" or "succeededCount" specified.`,
																	"properties": map[string]interface{}{
																		"succeededCount": map[string]interface{}{
																			"description": `succeededCount specifies the minimal required size of the actual set of the succeeded indexes
for the Job. When succeededCount is used along with succeededIndexes, the check is
constrained only to the set of indexes specified by succeededIndexes.
For example, given that succeededIndexes is "1-4", succeededCount is "3",
and completed indexes are "1", "3", and "5", the Job isn't declared as succeeded
because only "1" and "3" indexes are considered in that rules.
When this field is null, this doesn't default to any value and
is never evaluated at any time.
When specified it needs to be a positive integer.`,
																			"format": "int32",
																			"type":   "integer",
																		},
																		"succeededIndexes": map[string]interface{}{
																			"description": `succeededIndexes specifies the set of indexes
which need to be contained in the actual set of the succeeded indexes for the Job.
The list of indexes must be within 0 to ".spec.completions-1" and
must not contain duplicates. At least one element is required.
The indexes are represented as intervals separated by commas.
The intervals can be a decimal integer or a pair of decimal integers separated by a hyphen.
The number are listed in represented by the first and last element of the series,
separated by a hyphen.
For example, if the completed indexes are 1, 3, 4, 5 and 7, they are
represented as "1,3-5,7".
When this field is null, this field doesn't default to any value
and is never evaluated at any time.`,
																			"type": "string",
																		},
																	},
																	"type": "object",
																},
																"type":                   "array",
																"x-kubernetes-list-type": "atomic",
															},
														},
														"required": []interface{}{
															"rules",
														},
														"type": "object",
													},
													"suspend": map[string]interface{}{
														"description": `suspend specifies whether the Job controller should create Pods or not. If
a Job is created with suspend set to true, no Pods are created by the Job
controller. If a Job is suspended after creation (i.e. the flag goes from
false to true), the Job controller will delete all active Pods associated
with this Job. Users must design their workload to gracefully handle this.
Suspending a Job will reset the StartTime field of the Job, effectively
resetting the ActiveDeadlineSeconds timer too. Defaults to false.`,
														"type": "boolean",
													},
													"template": map[string]interface{}{
														"description": `Describes the pod that will be created when executing a job.
The only allowed template.spec.restartPolicy values are "Never" or "OnFailure".
More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/`,
														"properties": map[string]interface{}{
															"metadata": map[string]interface{}{
																"description": `Standard object's metadata.
More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata`,
																"properties": map[string]interface{}{
																	"annotations": map[string]interface{}{
																		"additionalProperties": map[string]interface{}{
																			"type": "string",
																		},
																		"type": "object",
																	},
																	"finalizers": map[string]interface{}{
																		"items": map[string]interface{}{
																			"type": "string",
																		},
																		"type": "array",
																	},
																	"labels": map[string]interface{}{
																		"additionalProperties": map[string]interface{}{
																			"type": "string",
																		},
																		"type": "object",
																	},
																	"name": map[string]interface{}{
																		"type": "string",
																	},
																	"namespace": map[string]interface{}{
																		"type": "string",
																	},
																},
																"type":                                 "object",
																"x-kubernetes-preserve-unknown-fields": true,
															},
															"spec": map[string]interface{}{
																"description": `Specification of the desired behavior of the pod.
More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status`,
																"properties": map[string]interface{}{
																	"activeDeadlineSeconds": map[string]interface{}{
																		"description": `Optional duration in seconds the pod may be active on the node relative to
StartTime before the system will actively try to mark it failed and kill associated containers.
Value must be a positive integer.`,
																		"format": "int64",
																		"type":   "integer",
																	},
																	"affinity": map[string]interface{}{
																		"description": "If specified, the pod's scheduling constraints",
																		"properties": map[string]interface{}{
																			"nodeAffinity": map[string]interface{}{
																				"description": "Describes node affinity scheduling rules for the pod.",
																				"properties": map[string]interface{}{
																					"preferredDuringSchedulingIgnoredDuringExecution": map[string]interface{}{
																						"description": `The scheduler will prefer to schedule pods to nodes that satisfy
the affinity expressions specified by this field, but it may choose
a node that violates one or more of the expressions. The node that is
most preferred is the one with the greatest sum of weights, i.e.
for each node that meets all of the scheduling requirements (resource
request, requiredDuringScheduling affinity expressions, etc.),
compute a sum by iterating through the elements of this field and adding
"weight" to the sum if the node matches the corresponding matchExpressions; the
node(s) with the highest sum are the most preferred.`,
																						"items": map[string]interface{}{
																							"description": `An empty preferred scheduling term matches all objects with implicit weight 0
(i.e. it's a no-op). A null preferred scheduling term matches no objects (i.e. is also a no-op).`,
																							"properties": map[string]interface{}{
																								"preference": map[string]interface{}{
																									"description": "A node selector term, associated with the corresponding weight.",
																									"properties": map[string]interface{}{
																										"matchExpressions": map[string]interface{}{
																											"description": "A list of node selector requirements by node's labels.",
																											"items": map[string]interface{}{
																												"description": `A node selector requirement is a selector that contains values, a key, and an operator
that relates the key and values.`,
																												"properties": map[string]interface{}{
																													"key": map[string]interface{}{
																														"description": "The label key that the selector applies to.",
																														"type":        "string",
																													},
																													"operator": map[string]interface{}{
																														"description": `Represents a key's relationship to a set of values.
Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.`,
																														"type": "string",
																													},
																													"values": map[string]interface{}{
																														"description": `An array of string values. If the operator is In or NotIn,
the values array must be non-empty. If the operator is Exists or DoesNotExist,
the values array must be empty. If the operator is Gt or Lt, the values
array must have a single element, which will be interpreted as an integer.
This array is replaced during a strategic merge patch.`,
																														"items": map[string]interface{}{
																															"type": "string",
																														},
																														"type":                   "array",
																														"x-kubernetes-list-type": "atomic",
																													},
																												},
																												"required": []interface{}{
																													"key",
																													"operator",
																												},
																												"type": "object",
																											},
																											"type":                   "array",
																											"x-kubernetes-list-type": "atomic",
																										},
																										"matchFields": map[string]interface{}{
																											"description": "A list of node selector requirements by node's fields.",
																											"items": map[string]interface{}{
																												"description": `A node selector requirement is a selector that contains values, a key, and an operator
that relates the key and values.`,
																												"properties": map[string]interface{}{
																													"key": map[string]interface{}{
																														"description": "The label key that the selector applies to.",
																														"type":        "string",
																													},
																													"operator": map[string]interface{}{
																														"description": `Represents a key's relationship to a set of values.
Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.`,
																														"type": "string",
																													},
																													"values": map[string]interface{}{
																														"description": `An array of string values. If the operator is In or NotIn,
the values array must be non-empty. If the operator is Exists or DoesNotExist,
the values array must be empty. If the operator is Gt or Lt, the values
array must have a single element, which will be interpreted as an integer.
This array is replaced during a strategic merge patch.`,
																														"items": map[string]interface{}{
																															"type": "string",
																														},
																														"type":                   "array",
																														"x-kubernetes-list-type": "atomic",
																													},
																												},
																												"required": []interface{}{
																													"key",
																													"operator",
																												},
																												"type": "object",
																											},
																											"type":                   "array",
																											"x-kubernetes-list-type": "atomic",
																										},
																									},
																									"type":                  "object",
																									"x-kubernetes-map-type": "atomic",
																								},
																								"weight": map[string]interface{}{
																									"description": "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",
																									"format":      "int32",
																									"type":        "integer",
																								},
																							},
																							"required": []interface{}{
																								"preference",
																								"weight",
																							},
																							"type": "object",
																						},
																						"type":                   "array",
																						"x-kubernetes-list-type": "atomic",
																					},
																					"requiredDuringSchedulingIgnoredDuringExecution": map[string]interface{}{
																						"description": `If the affinity requirements specified by this field are not met at
scheduling time, the pod will not be scheduled onto the node.
If the affinity requirements specified by this field cease to be met
at some point during pod execution (e.g. due to an update), the system
may or may not try to eventually evict the pod from its node.`,
																						"properties": map[string]interface{}{
																							"nodeSelectorTerms": map[string]interface{}{
																								"description": "Required. A list of node selector terms. The terms are ORed.",
																								"items": map[string]interface{}{
																									"description": `A null or empty node selector term matches no objects. The requirements of
them are ANDed.
The TopologySelectorTerm type implements a subset of the NodeSelectorTerm.`,
																									"properties": map[string]interface{}{
																										"matchExpressions": map[string]interface{}{
																											"description": "A list of node selector requirements by node's labels.",
																											"items": map[string]interface{}{
																												"description": `A node selector requirement is a selector that contains values, a key, and an operator
that relates the key and values.`,
																												"properties": map[string]interface{}{
																													"key": map[string]interface{}{
																														"description": "The label key that the selector applies to.",
																														"type":        "string",
																													},
																													"operator": map[string]interface{}{
																														"description": `Represents a key's relationship to a set of values.
Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.`,
																														"type": "string",
																													},
																													"values": map[string]interface{}{
																														"description": `An array of string values. If the operator is In or NotIn,
the values array must be non-empty. If the operator is Exists or DoesNotExist,
the values array must be empty. If the operator is Gt or Lt, the values
array must have a single element, which will be interpreted as an integer.
This array is replaced during a strategic merge patch.`,
																														"items": map[string]interface{}{
																															"type": "string",
																														},
																														"type":                   "array",
																														"x-kubernetes-list-type": "atomic",
																													},
																												},
																												"required": []interface{}{
																													"key",
																													"operator",
																												},
																												"type": "object",
																											},
																											"type":                   "array",
																											"x-kubernetes-list-type": "atomic",
																										},
																										"matchFields": map[string]interface{}{
																											"description": "A list of node selector requirements by node's fields.",
																											"items": map[string]interface{}{
																												"description": `A node selector requirement is a selector that contains values, a key, and an operator
that relates the key and values.`,
																												"properties": map[string]interface{}{
																													"key": map[string]interface{}{
																														"description": "The label key that the selector applies to.",
																														"type":        "string",
																													},
																													"operator": map[string]interface{}{
																														"description": `Represents a key's relationship to a set of values.
Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.`,
																														"type": "string",
																													},
																													"values": map[string]interface{}{
																														"description": `An array of string values. If the operator is In or NotIn,
the values array must be non-empty. If the operator is Exists or DoesNotExist,
the values array must be empty. If the operator is Gt or Lt, the values
array must have a single element, which will be interpreted as an integer.
This array is replaced during a strategic merge patch.`,
																														"items": map[string]interface{}{
																															"type": "string",
																														},
																														"type":                   "array",
																														"x-kubernetes-list-type": "atomic",
																													},
																												},
																												"required": []interface{}{
																													"key",
																													"operator",
																												},
																												"type": "object",
																											},
																											"type":                   "array",
																											"x-kubernetes-list-type": "atomic",
																										},
																									},
																									"type":                  "object",
																									"x-kubernetes-map-type": "atomic",
																								},
																								"type":                   "array",
																								"x-kubernetes-list-type": "atomic",
																							},
																						},
																						"required": []interface{}{
																							"nodeSelectorTerms",
																						},
																						"type":                  "object",
																						"x-kubernetes-map-type": "atomic",
																					},
																				},
																				"type": "object",
																			},
																			"podAffinity": map[string]interface{}{
																				"description": "Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).",
																				"properties": map[string]interface{}{
																					"preferredDuringSchedulingIgnoredDuringExecution": map[string]interface{}{
																						"description": `The scheduler will prefer to schedule pods to nodes that satisfy
the affinity expressions specified by this field, but it may choose
a node that violates one or more of the expressions. The node that is
most preferred is the one with the greatest sum of weights, i.e.
for each node that meets all of the scheduling requirements (resource
request, requiredDuringScheduling affinity expressions, etc.),
compute a sum by iterating through the elements of this field and adding
"weight" to the sum if the node has pods which matches the corresponding podAffinityTerm; the
node(s) with the highest sum are the most preferred.`,
																						"items": map[string]interface{}{
																							"description": "The weights of all of the matched WeightedPodAffinityTerm fields are added per-node to find the most preferred node(s)",
																							"properties": map[string]interface{}{
																								"podAffinityTerm": map[string]interface{}{
																									"description": "Required. A pod affinity term, associated with the corresponding weight.",
																									"properties": map[string]interface{}{
																										"labelSelector": map[string]interface{}{
																											"description": `A label query over a set of resources, in this case pods.
If it's null, this PodAffinityTerm matches with no Pods.`,
																											"properties": map[string]interface{}{
																												"matchExpressions": map[string]interface{}{
																													"description": "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																													"items": map[string]interface{}{
																														"description": `A label selector requirement is a selector that contains values, a key, and an operator that
relates the key and values.`,
																														"properties": map[string]interface{}{
																															"key": map[string]interface{}{
																																"description": "key is the label key that the selector applies to.",
																																"type":        "string",
																															},
																															"operator": map[string]interface{}{
																																"description": `operator represents a key's relationship to a set of values.
Valid operators are In, NotIn, Exists and DoesNotExist.`,
																																"type": "string",
																															},
																															"values": map[string]interface{}{
																																"description": `values is an array of string values. If the operator is In or NotIn,
the values array must be non-empty. If the operator is Exists or DoesNotExist,
the values array must be empty. This array is replaced during a strategic
merge patch.`,
																																"items": map[string]interface{}{
																																	"type": "string",
																																},
																																"type":                   "array",
																																"x-kubernetes-list-type": "atomic",
																															},
																														},
																														"required": []interface{}{
																															"key",
																															"operator",
																														},
																														"type": "object",
																													},
																													"type":                   "array",
																													"x-kubernetes-list-type": "atomic",
																												},
																												"matchLabels": map[string]interface{}{
																													"additionalProperties": map[string]interface{}{
																														"type": "string",
																													},
																													"description": `matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels
map is equivalent to an element of matchExpressions, whose key field is "key", the
operator is "In", and the values array contains only "value". The requirements are ANDed.`,
																													"type": "object",
																												},
																											},
																											"type":                  "object",
																											"x-kubernetes-map-type": "atomic",
																										},
																										"matchLabelKeys": map[string]interface{}{
																											"description": `MatchLabelKeys is a set of pod label keys to select which pods will
be taken into consideration. The keys are used to lookup values from the
incoming pod labels, those key-value labels are merged with ` + "`" + `labelSelector` + "`" + ` as ` + "`" + `key in (value)` + "`" + `
to select the group of existing pods which pods will be taken into consideration
for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming
pod labels will be ignored. The default value is empty.
The same key is forbidden to exist in both matchLabelKeys and labelSelector.
Also, matchLabelKeys cannot be set when labelSelector isn't set.`,
																											"items": map[string]interface{}{
																												"type": "string",
																											},
																											"type":                   "array",
																											"x-kubernetes-list-type": "atomic",
																										},
																										"mismatchLabelKeys": map[string]interface{}{
																											"description": `MismatchLabelKeys is a set of pod label keys to select which pods will
be taken into consideration. The keys are used to lookup values from the
incoming pod labels, those key-value labels are merged with ` + "`" + `labelSelector` + "`" + ` as ` + "`" + `key notin (value)` + "`" + `
to select the group of existing pods which pods will be taken into consideration
for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming
pod labels will be ignored. The default value is empty.
The same key is forbidden to exist in both mismatchLabelKeys and labelSelector.
Also, mismatchLabelKeys cannot be set when labelSelector isn't set.`,
																											"items": map[string]interface{}{
																												"type": "string",
																											},
																											"type":                   "array",
																											"x-kubernetes-list-type": "atomic",
																										},
																										"namespaceSelector": map[string]interface{}{
																											"description": `A label query over the set of namespaces that the term applies to.
The term is applied to the union of the namespaces selected by this field
and the ones listed in the namespaces field.
null selector and null or empty namespaces list means "this pod's namespace".
An empty selector ({}) matches all namespaces.`,
																											"properties": map[string]interface{}{
																												"matchExpressions": map[string]interface{}{
																													"description": "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																													"items": map[string]interface{}{
																														"description": `A label selector requirement is a selector that contains values, a key, and an operator that
relates the key and values.`,
																														"properties": map[string]interface{}{
																															"key": map[string]interface{}{
																																"description": "key is the label key that the selector applies to.",
																																"type":        "string",
																															},
																															"operator": map[string]interface{}{
																																"description": `operator represents a key's relationship to a set of values.
Valid operators are In, NotIn, Exists and DoesNotExist.`,
																																"type": "string",
																															},
																															"values": map[string]interface{}{
																																"description": `values is an array of string values. If the operator is In or NotIn,
the values array must be non-empty. If the operator is Exists or DoesNotExist,
the values array must be empty. This array is replaced during a strategic
merge patch.`,
																																"items": map[string]interface{}{
																																	"type": "string",
																																},
																																"type":                   "array",
																																"x-kubernetes-list-type": "atomic",
																															},
																														},
																														"required": []interface{}{
																															"key",
																															"operator",
																														},
																														"type": "object",
																													},
																													"type":                   "array",
																													"x-kubernetes-list-type": "atomic",
																												},
																												"matchLabels": map[string]interface{}{
																													"additionalProperties": map[string]interface{}{
																														"type": "string",
																													},
																													"description": `matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels
map is equivalent to an element of matchExpressions, whose key field is "key", the
operator is "In", and the values array contains only "value". The requirements are ANDed.`,
																													"type": "object",
																												},
																											},
																											"type":                  "object",
																											"x-kubernetes-map-type": "atomic",
																										},
																										"namespaces": map[string]interface{}{
																											"description": `namespaces specifies a static list of namespace names that the term applies to.
The term is applied to the union of the namespaces listed in this field
and the ones selected by namespaceSelector.
null or empty namespaces list and null namespaceSelector means "this pod's namespace".`,
																											"items": map[string]interface{}{
																												"type": "string",
																											},
																											"type":                   "array",
																											"x-kubernetes-list-type": "atomic",
																										},
																										"topologyKey": map[string]interface{}{
																											"description": `This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching
the labelSelector in the specified namespaces, where co-located is defined as running on a node
whose value of the label with key topologyKey matches that of any node on which any of the
selected pods is running.
Empty topologyKey is not allowed.`,
																											"type": "string",
																										},
																									},
																									"required": []interface{}{
																										"topologyKey",
																									},
																									"type": "object",
																								},
																								"weight": map[string]interface{}{
																									"description": `weight associated with matching the corresponding podAffinityTerm,
in the range 1-100.`,
																									"format": "int32",
																									"type":   "integer",
																								},
																							},
																							"required": []interface{}{
																								"podAffinityTerm",
																								"weight",
																							},
																							"type": "object",
																						},
																						"type":                   "array",
																						"x-kubernetes-list-type": "atomic",
																					},
																					"requiredDuringSchedulingIgnoredDuringExecution": map[string]interface{}{
																						"description": `If the affinity requirements specified by this field are not met at
scheduling time, the pod will not be scheduled onto the node.
If the affinity requirements specified by this field cease to be met
at some point during pod execution (e.g. due to a pod label update), the
system may or may not try to eventually evict the pod from its node.
When there are multiple elements, the lists of nodes corresponding to each
podAffinityTerm are intersected, i.e. all terms must be satisfied.`,
																						"items": map[string]interface{}{
																							"description": `Defines a set of pods (namely those matching the labelSelector
relative to the given namespace(s)) that this pod should be
co-located (affinity) or not co-located (anti-affinity) with,
where co-located is defined as running on a node whose value of
the label with key <topologyKey> matches that of any node on which
a pod of the set of pods is running`,
																							"properties": map[string]interface{}{
																								"labelSelector": map[string]interface{}{
																									"description": `A label query over a set of resources, in this case pods.
If it's null, this PodAffinityTerm matches with no Pods.`,
																									"properties": map[string]interface{}{
																										"matchExpressions": map[string]interface{}{
																											"description": "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																											"items": map[string]interface{}{
																												"description": `A label selector requirement is a selector that contains values, a key, and an operator that
relates the key and values.`,
																												"properties": map[string]interface{}{
																													"key": map[string]interface{}{
																														"description": "key is the label key that the selector applies to.",
																														"type":        "string",
																													},
																													"operator": map[string]interface{}{
																														"description": `operator represents a key's relationship to a set of values.
Valid operators are In, NotIn, Exists and DoesNotExist.`,
																														"type": "string",
																													},
																													"values": map[string]interface{}{
																														"description": `values is an array of string values. If the operator is In or NotIn,
the values array must be non-empty. If the operator is Exists or DoesNotExist,
the values array must be empty. This array is replaced during a strategic
merge patch.`,
																														"items": map[string]interface{}{
																															"type": "string",
																														},
																														"type":                   "array",
																														"x-kubernetes-list-type": "atomic",
																													},
																												},
																												"required": []interface{}{
																													"key",
																													"operator",
																												},
																												"type": "object",
																											},
																											"type":                   "array",
																											"x-kubernetes-list-type": "atomic",
																										},
																										"matchLabels": map[string]interface{}{
																											"additionalProperties": map[string]interface{}{
																												"type": "string",
																											},
																											"description": `matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels
map is equivalent to an element of matchExpressions, whose key field is "key", the
operator is "In", and the values array contains only "value". The requirements are ANDed.`,
																											"type": "object",
																										},
																									},
																									"type":                  "object",
																									"x-kubernetes-map-type": "atomic",
																								},
																								"matchLabelKeys": map[string]interface{}{
																									"description": `MatchLabelKeys is a set of pod label keys to select which pods will
be taken into consideration. The keys are used to lookup values from the
incoming pod labels, those key-value labels are merged with ` + "`" + `labelSelector` + "`" + ` as ` + "`" + `key in (value)` + "`" + `
to select the group of existing pods which pods will be taken into consideration
for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming
pod labels will be ignored. The default value is empty.
The same key is forbidden to exist in both matchLabelKeys and labelSelector.
Also, matchLabelKeys cannot be set when labelSelector isn't set.`,
																									"items": map[string]interface{}{
																										"type": "string",
																									},
																									"type":                   "array",
																									"x-kubernetes-list-type": "atomic",
																								},
																								"mismatchLabelKeys": map[string]interface{}{
																									"description": `MismatchLabelKeys is a set of pod label keys to select which pods will
be taken into consideration. The keys are used to lookup values from the
incoming pod labels, those key-value labels are merged with ` + "`" + `labelSelector` + "`" + ` as ` + "`" + `key notin (value)` + "`" + `
to select the group of existing pods which pods will be taken into consideration
for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming
pod labels will be ignored. The default value is empty.
The same key is forbidden to exist in both mismatchLabelKeys and labelSelector.
Also, mismatchLabelKeys cannot be set when labelSelector isn't set.`,
																									"items": map[string]interface{}{
																										"type": "string",
																									},
																									"type":                   "array",
																									"x-kubernetes-list-type": "atomic",
																								},
																								"namespaceSelector": map[string]interface{}{
																									"description": `A label query over the set of namespaces that the term applies to.
The term is applied to the union of the namespaces selected by this field
and the ones listed in the namespaces field.
null selector and null or empty namespaces list means "this pod's namespace".
An empty selector ({}) matches all namespaces.`,
																									"properties": map[string]interface{}{
																										"matchExpressions": map[string]interface{}{
																											"description": "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																											"items": map[string]interface{}{
																												"description": `A label selector requirement is a selector that contains values, a key, and an operator that
relates the key and values.`,
																												"properties": map[string]interface{}{
																													"key": map[string]interface{}{
																														"description": "key is the label key that the selector applies to.",
																														"type":        "string",
																													},
																													"operator": map[string]interface{}{
																														"description": `operator represents a key's relationship to a set of values.
Valid operators are In, NotIn, Exists and DoesNotExist.`,
																														"type": "string",
																													},
																													"values": map[string]interface{}{
																														"description": `values is an array of string values. If the operator is In or NotIn,
the values array must be non-empty. If the operator is Exists or DoesNotExist,
the values array must be empty. This array is replaced during a strategic
merge patch.`,
																														"items": map[string]interface{}{
																															"type": "string",
																														},
																														"type":                   "array",
																														"x-kubernetes-list-type": "atomic",
																													},
																												},
																												"required": []interface{}{
																													"key",
																													"operator",
																												},
																												"type": "object",
																											},
																											"type":                   "array",
																											"x-kubernetes-list-type": "atomic",
																										},
																										"matchLabels": map[string]interface{}{
																											"additionalProperties": map[string]interface{}{
																												"type": "string",
																											},
																											"description": `matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels
map is equivalent to an element of matchExpressions, whose key field is "key", the
operator is "In", and the values array contains only "value". The requirements are ANDed.`,
																											"type": "object",
																										},
																									},
																									"type":                  "object",
																									"x-kubernetes-map-type": "atomic",
																								},
																								"namespaces": map[string]interface{}{
																									"description": `namespaces specifies a static list of namespace names that the term applies to.
The term is applied to the union of the namespaces listed in this field
and the ones selected by namespaceSelector.
null or empty namespaces list and null namespaceSelector means "this pod's namespace".`,
																									"items": map[string]interface{}{
																										"type": "string",
																									},
																									"type":                   "array",
																									"x-kubernetes-list-type": "atomic",
																								},
																								"topologyKey": map[string]interface{}{
																									"description": `This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching
the labelSelector in the specified namespaces, where co-located is defined as running on a node
whose value of the label with key topologyKey matches that of any node on which any of the
selected pods is running.
Empty topologyKey is not allowed.`,
																									"type": "string",
																								},
																							},
																							"required": []interface{}{
																								"topologyKey",
																							},
																							"type": "object",
																						},
																						"type":                   "array",
																						"x-kubernetes-list-type": "atomic",
																					},
																				},
																				"type": "object",
																			},
																			"podAntiAffinity": map[string]interface{}{
																				"description": "Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).",
																				"properties": map[string]interface{}{
																					"preferredDuringSchedulingIgnoredDuringExecution": map[string]interface{}{
																						"description": `The scheduler will prefer to schedule pods to nodes that satisfy
the anti-affinity expressions specified by this field, but it may choose
a node that violates one or more of the expressions. The node that is
most preferred is the one with the greatest sum of weights, i.e.
for each node that meets all of the scheduling requirements (resource
request, requiredDuringScheduling anti-affinity expressions, etc.),
compute a sum by iterating through the elements of this field and subtracting
"weight" from the sum if the node has pods which matches the corresponding podAffinityTerm; the
node(s) with the highest sum are the most preferred.`,
																						"items": map[string]interface{}{
																							"description": "The weights of all of the matched WeightedPodAffinityTerm fields are added per-node to find the most preferred node(s)",
																							"properties": map[string]interface{}{
																								"podAffinityTerm": map[string]interface{}{
																									"description": "Required. A pod affinity term, associated with the corresponding weight.",
																									"properties": map[string]interface{}{
																										"labelSelector": map[string]interface{}{
																											"description": `A label query over a set of resources, in this case pods.
If it's null, this PodAffinityTerm matches with no Pods.`,
																											"properties": map[string]interface{}{
																												"matchExpressions": map[string]interface{}{
																													"description": "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																													"items": map[string]interface{}{
																														"description": `A label selector requirement is a selector that contains values, a key, and an operator that
relates the key and values.`,
																														"properties": map[string]interface{}{
																															"key": map[string]interface{}{
																																"description": "key is the label key that the selector applies to.",
																																"type":        "string",
																															},
																															"operator": map[string]interface{}{
																																"description": `operator represents a key's relationship to a set of values.
Valid operators are In, NotIn, Exists and DoesNotExist.`,
																																"type": "string",
																															},
																															"values": map[string]interface{}{
																																"description": `values is an array of string values. If the operator is In or NotIn,
the values array must be non-empty. If the operator is Exists or DoesNotExist,
the values array must be empty. This array is replaced during a strategic
merge patch.`,
																																"items": map[string]interface{}{
																																	"type": "string",
																																},
																																"type":                   "array",
																																"x-kubernetes-list-type": "atomic",
																															},
																														},
																														"required": []interface{}{
																															"key",
																															"operator",
																														},
																														"type": "object",
																													},
																													"type":                   "array",
																													"x-kubernetes-list-type": "atomic",
																												},
																												"matchLabels": map[string]interface{}{
																													"additionalProperties": map[string]interface{}{
																														"type": "string",
																													},
																													"description": `matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels
map is equivalent to an element of matchExpressions, whose key field is "key", the
operator is "In", and the values array contains only "value". The requirements are ANDed.`,
																													"type": "object",
																												},
																											},
																											"type":                  "object",
																											"x-kubernetes-map-type": "atomic",
																										},
																										"matchLabelKeys": map[string]interface{}{
																											"description": `MatchLabelKeys is a set of pod label keys to select which pods will
be taken into consideration. The keys are used to lookup values from the
incoming pod labels, those key-value labels are merged with ` + "`" + `labelSelector` + "`" + ` as ` + "`" + `key in (value)` + "`" + `
to select the group of existing pods which pods will be taken into consideration
for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming
pod labels will be ignored. The default value is empty.
The same key is forbidden to exist in both matchLabelKeys and labelSelector.
Also, matchLabelKeys cannot be set when labelSelector isn't set.`,
																											"items": map[string]interface{}{
																												"type": "string",
																											},
																											"type":                   "array",
																											"x-kubernetes-list-type": "atomic",
																										},
																										"mismatchLabelKeys": map[string]interface{}{
																											"description": `MismatchLabelKeys is a set of pod label keys to select which pods will
be taken into consideration. The keys are used to lookup values from the
incoming pod labels, those key-value labels are merged with ` + "`" + `labelSelector` + "`" + ` as ` + "`" + `key notin (value)` + "`" + `
to select the group of existing pods which pods will be taken into consideration
for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming
pod labels will be ignored. The default value is empty.
The same key is forbidden to exist in both mismatchLabelKeys and labelSelector.
Also, mismatchLabelKeys cannot be set when labelSelector isn't set.`,
																											"items": map[string]interface{}{
																												"type": "string",
																											},
																											"type":                   "array",
																											"x-kubernetes-list-type": "atomic",
																										},
																										"namespaceSelector": map[string]interface{}{
																											"description": `A label query over the set of namespaces that the term applies to.
The term is applied to the union of the namespaces selected by this field
and the ones listed in the namespaces field.
null selector and null or empty namespaces list means "this pod's namespace".
An empty selector ({}) matches all namespaces.`,
																											"properties": map[string]interface{}{
																												"matchExpressions": map[string]interface{}{
																													"description": "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																													"items": map[string]interface{}{
																														"description": `A label selector requirement is a selector that contains values, a key, and an operator that
relates the key and values.`,
																														"properties": map[string]interface{}{
																															"key": map[string]interface{}{
																																"description": "key is the label key that the selector applies to.",
																																"type":        "string",
																															},
																															"operator": map[string]interface{}{
																																"description": `operator represents a key's relationship to a set of values.
Valid operators are In, NotIn, Exists and DoesNotExist.`,
																																"type": "string",
																															},
																															"values": map[string]interface{}{
																																"description": `values is an array of string values. If the operator is In or NotIn,
the values array must be non-empty. If the operator is Exists or DoesNotExist,
the values array must be empty. This array is replaced during a strategic
merge patch.`,
																																"items": map[string]interface{}{
																																	"type": "string",
																																},
																																"type":                   "array",
																																"x-kubernetes-list-type": "atomic",
																															},
																														},
																														"required": []interface{}{
																															"key",
																															"operator",
																														},
																														"type": "object",
																													},
																													"type":                   "array",
																													"x-kubernetes-list-type": "atomic",
																												},
																												"matchLabels": map[string]interface{}{
																													"additionalProperties": map[string]interface{}{
																														"type": "string",
																													},
																													"description": `matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels
map is equivalent to an element of matchExpressions, whose key field is "key", the
operator is "In", and the values array contains only "value". The requirements are ANDed.`,
																													"type": "object",
																												},
																											},
																											"type":                  "object",
																											"x-kubernetes-map-type": "atomic",
																										},
																										"namespaces": map[string]interface{}{
																											"description": `namespaces specifies a static list of namespace names that the term applies to.
The term is applied to the union of the namespaces listed in this field
and the ones selected by namespaceSelector.
null or empty namespaces list and null namespaceSelector means "this pod's namespace".`,
																											"items": map[string]interface{}{
																												"type": "string",
																											},
																											"type":                   "array",
																											"x-kubernetes-list-type": "atomic",
																										},
																										"topologyKey": map[string]interface{}{
																											"description": `This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching
the labelSelector in the specified namespaces, where co-located is defined as running on a node
whose value of the label with key topologyKey matches that of any node on which any of the
selected pods is running.
Empty topologyKey is not allowed.`,
																											"type": "string",
																										},
																									},
																									"required": []interface{}{
																										"topologyKey",
																									},
																									"type": "object",
																								},
																								"weight": map[string]interface{}{
																									"description": `weight associated with matching the corresponding podAffinityTerm,
in the range 1-100.`,
																									"format": "int32",
																									"type":   "integer",
																								},
																							},
																							"required": []interface{}{
																								"podAffinityTerm",
																								"weight",
																							},
																							"type": "object",
																						},
																						"type":                   "array",
																						"x-kubernetes-list-type": "atomic",
																					},
																					"requiredDuringSchedulingIgnoredDuringExecution": map[string]interface{}{
																						"description": `If the anti-affinity requirements specified by this field are not met at
scheduling time, the pod will not be scheduled onto the node.
If the anti-affinity requirements specified by this field cease to be met
at some point during pod execution (e.g. due to a pod label update), the
system may or may not try to eventually evict the pod from its node.
When there are multiple elements, the lists of nodes corresponding to each
podAffinityTerm are intersected, i.e. all terms must be satisfied.`,
																						"items": map[string]interface{}{
																							"description": `Defines a set of pods (namely those matching the labelSelector
relative to the given namespace(s)) that this pod should be
co-located (affinity) or not co-located (anti-affinity) with,
where co-located is defined as running on a node whose value of
the label with key <topologyKey> matches that of any node on which
a pod of the set of pods is running`,
																							"properties": map[string]interface{}{
																								"labelSelector": map[string]interface{}{
																									"description": `A label query over a set of resources, in this case pods.
If it's null, this PodAffinityTerm matches with no Pods.`,
																									"properties": map[string]interface{}{
																										"matchExpressions": map[string]interface{}{
																											"description": "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																											"items": map[string]interface{}{
																												"description": `A label selector requirement is a selector that contains values, a key, and an operator that
relates the key and values.`,
																												"properties": map[string]interface{}{
																													"key": map[string]interface{}{
																														"description": "key is the label key that the selector applies to.",
																														"type":        "string",
																													},
																													"operator": map[string]interface{}{
																														"description": `operator represents a key's relationship to a set of values.
Valid operators are In, NotIn, Exists and DoesNotExist.`,
																														"type": "string",
																													},
																													"values": map[string]interface{}{
																														"description": `values is an array of string values. If the operator is In or NotIn,
the values array must be non-empty. If the operator is Exists or DoesNotExist,
the values array must be empty. This array is replaced during a strategic
merge patch.`,
																														"items": map[string]interface{}{
																															"type": "string",
																														},
																														"type":                   "array",
																														"x-kubernetes-list-type": "atomic",
																													},
																												},
																												"required": []interface{}{
																													"key",
																													"operator",
																												},
																												"type": "object",
																											},
																											"type":                   "array",
																											"x-kubernetes-list-type": "atomic",
																										},
																										"matchLabels": map[string]interface{}{
																											"additionalProperties": map[string]interface{}{
																												"type": "string",
																											},
																											"description": `matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels
map is equivalent to an element of matchExpressions, whose key field is "key", the
operator is "In", and the values array contains only "value". The requirements are ANDed.`,
																											"type": "object",
																										},
																									},
																									"type":                  "object",
																									"x-kubernetes-map-type": "atomic",
																								},
																								"matchLabelKeys": map[string]interface{}{
																									"description": `MatchLabelKeys is a set of pod label keys to select which pods will
be taken into consideration. The keys are used to lookup values from the
incoming pod labels, those key-value labels are merged with ` + "`" + `labelSelector` + "`" + ` as ` + "`" + `key in (value)` + "`" + `
to select the group of existing pods which pods will be taken into consideration
for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming
pod labels will be ignored. The default value is empty.
The same key is forbidden to exist in both matchLabelKeys and labelSelector.
Also, matchLabelKeys cannot be set when labelSelector isn't set.`,
																									"items": map[string]interface{}{
																										"type": "string",
																									},
																									"type":                   "array",
																									"x-kubernetes-list-type": "atomic",
																								},
																								"mismatchLabelKeys": map[string]interface{}{
																									"description": `MismatchLabelKeys is a set of pod label keys to select which pods will
be taken into consideration. The keys are used to lookup values from the
incoming pod labels, those key-value labels are merged with ` + "`" + `labelSelector` + "`" + ` as ` + "`" + `key notin (value)` + "`" + `
to select the group of existing pods which pods will be taken into consideration
for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming
pod labels will be ignored. The default value is empty.
The same key is forbidden to exist in both mismatchLabelKeys and labelSelector.
Also, mismatchLabelKeys cannot be set when labelSelector isn't set.`,
																									"items": map[string]interface{}{
																										"type": "string",
																									},
																									"type":                   "array",
																									"x-kubernetes-list-type": "atomic",
																								},
																								"namespaceSelector": map[string]interface{}{
																									"description": `A label query over the set of namespaces that the term applies to.
The term is applied to the union of the namespaces selected by this field
and the ones listed in the namespaces field.
null selector and null or empty namespaces list means "this pod's namespace".
An empty selector ({}) matches all namespaces.`,
																									"properties": map[string]interface{}{
																										"matchExpressions": map[string]interface{}{
																											"description": "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																											"items": map[string]interface{}{
																												"description": `A label selector requirement is a selector that contains values, a key, and an operator that
relates the key and values.`,
																												"properties": map[string]interface{}{
																													"key": map[string]interface{}{
																														"description": "key is the label key that the selector applies to.",
																														"type":        "string",
																													},
																													"operator": map[string]interface{}{
																														"description": `operator represents a key's relationship to a set of values.
Valid operators are In, NotIn, Exists and DoesNotExist.`,
																														"type": "string",
																													},
																													"values": map[string]interface{}{
																														"description": `values is an array of string values. If the operator is In or NotIn,
the values array must be non-empty. If the operator is Exists or DoesNotExist,
the values array must be empty. This array is replaced during a strategic
merge patch.`,
																														"items": map[string]interface{}{
																															"type": "string",
																														},
																														"type":                   "array",
																														"x-kubernetes-list-type": "atomic",
																													},
																												},
																												"required": []interface{}{
																													"key",
																													"operator",
																												},
																												"type": "object",
																											},
																											"type":                   "array",
																											"x-kubernetes-list-type": "atomic",
																										},
																										"matchLabels": map[string]interface{}{
																											"additionalProperties": map[string]interface{}{
																												"type": "string",
																											},
																											"description": `matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels
map is equivalent to an element of matchExpressions, whose key field is "key", the
operator is "In", and the values array contains only "value". The requirements are ANDed.`,
																											"type": "object",
																										},
																									},
																									"type":                  "object",
																									"x-kubernetes-map-type": "atomic",
																								},
																								"namespaces": map[string]interface{}{
																									"description": `namespaces specifies a static list of namespace names that the term applies to.
The term is applied to the union of the namespaces listed in this field
and the ones selected by namespaceSelector.
null or empty namespaces list and null namespaceSelector means "this pod's namespace".`,
																									"items": map[string]interface{}{
																										"type": "string",
																									},
																									"type":                   "array",
																									"x-kubernetes-list-type": "atomic",
																								},
																								"topologyKey": map[string]interface{}{
																									"description": `This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching
the labelSelector in the specified namespaces, where co-located is defined as running on a node
whose value of the label with key topologyKey matches that of any node on which any of the
selected pods is running.
Empty topologyKey is not allowed.`,
																									"type": "string",
																								},
																							},
																							"required": []interface{}{
																								"topologyKey",
																							},
																							"type": "object",
																						},
																						"type":                   "array",
																						"x-kubernetes-list-type": "atomic",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																	"automountServiceAccountToken": map[string]interface{}{
																		"description": "AutomountServiceAccountToken indicates whether a service account token should be automatically mounted.",
																		"type":        "boolean",
																	},
																	"containers": map[string]interface{}{
																		"description": `List of containers belonging to the pod.
Containers cannot currently be added or removed.
There must be at least one container in a Pod.
Cannot be updated.`,
																		"items": map[string]interface{}{
																			"description": "A single application container that you want to run within a pod.",
																			"properties": map[string]interface{}{
																				"args": map[string]interface{}{
																					"description": `Arguments to the entrypoint.
The container image's CMD is used if this is not provided.
Variable references $(VAR_NAME) are expanded using the container's environment. If a variable
cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced
to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. "$$(VAR_NAME)" will
produce the string literal "$(VAR_NAME)". Escaped references will never be expanded, regardless
of whether the variable exists or not. Cannot be updated.
More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell`,
																					"items": map[string]interface{}{
																						"type": "string",
																					},
																					"type":                   "array",
																					"x-kubernetes-list-type": "atomic",
																				},
																				"command": map[string]interface{}{
																					"description": `Entrypoint array. Not executed within a shell.
The container image's ENTRYPOINT is used if this is not provided.
Variable references $(VAR_NAME) are expanded using the container's environment. If a variable
cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced
to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. "$$(VAR_NAME)" will
produce the string literal "$(VAR_NAME)". Escaped references will never be expanded, regardless
of whether the variable exists or not. Cannot be updated.
More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell`,
																					"items": map[string]interface{}{
																						"type": "string",
																					},
																					"type":                   "array",
																					"x-kubernetes-list-type": "atomic",
																				},
																				"env": map[string]interface{}{
																					"description": `List of environment variables to set in the container.
Cannot be updated.`,
																					"items": map[string]interface{}{
																						"description": "EnvVar represents an environment variable present in a Container.",
																						"properties": map[string]interface{}{
																							"name": map[string]interface{}{
																								"description": `Name of the environment variable.
May consist of any printable ASCII characters except '='.`,
																								"type": "string",
																							},
																							"value": map[string]interface{}{
																								"description": `Variable references $(VAR_NAME) are expanded
using the previously defined environment variables in the container and
any service environment variables. If a variable cannot be resolved,
the reference in the input string will be unchanged. Double $$ are reduced
to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e.
"$$(VAR_NAME)" will produce the string literal "$(VAR_NAME)".
Escaped references will never be expanded, regardless of whether the variable
exists or not.
Defaults to "".`,
																								"type": "string",
																							},
																							"valueFrom": map[string]interface{}{
																								"description": "Source for the environment variable's value. Cannot be used if value is not empty.",
																								"properties": map[string]interface{}{
																									"configMapKeyRef": map[string]interface{}{
																										"description": "Selects a key of a ConfigMap.",
																										"properties": map[string]interface{}{
																											"key": map[string]interface{}{
																												"description": "The key to select.",
																												"type":        "string",
																											},
																											"name": map[string]interface{}{
																												"default": "",
																												"description": `Name of the referent.
This field is effectively required, but due to backwards compatibility is
allowed to be empty. Instances of this type with an empty value here are
almost certainly wrong.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names`,
																												"type": "string",
																											},
																											"optional": map[string]interface{}{
																												"description": "Specify whether the ConfigMap or its key must be defined",
																												"type":        "boolean",
																											},
																										},
																										"required": []interface{}{
																											"key",
																										},
																										"type":                  "object",
																										"x-kubernetes-map-type": "atomic",
																									},
																									"fieldRef": map[string]interface{}{
																										"description": `Selects a field of the pod: supports metadata.name, metadata.namespace, ` + "`" + `metadata.labels['<KEY>']` + "`" + `, ` + "`" + `metadata.annotations['<KEY>']` + "`" + `,
spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.`,
																										"properties": map[string]interface{}{
																											"apiVersion": map[string]interface{}{
																												"description": "Version of the schema the FieldPath is written in terms of, defaults to \"v1\".",
																												"type":        "string",
																											},
																											"fieldPath": map[string]interface{}{
																												"description": "Path of the field to select in the specified API version.",
																												"type":        "string",
																											},
																										},
																										"required": []interface{}{
																											"fieldPath",
																										},
																										"type":                  "object",
																										"x-kubernetes-map-type": "atomic",
																									},
																									"fileKeyRef": map[string]interface{}{
																										"description": `FileKeyRef selects a key of the env file.
Requires the EnvFiles feature gate to be enabled.`,
																										"properties": map[string]interface{}{
																											"key": map[string]interface{}{
																												"description": `The key within the env file. An invalid key will prevent the pod from starting.
The keys defined within a source may consist of any printable ASCII characters except '='.
During Alpha stage of the EnvFiles feature gate, the key size is limited to 128 characters.`,
																												"type": "string",
																											},
																											"optional": map[string]interface{}{
																												"default": false,
																												"description": `Specify whether the file or its key must be defined. If the file or key
does not exist, then the env var is not published.
If optional is set to true and the specified key does not exist,
the environment variable will not be set in the Pod's containers.

If optional is set to false and the specified key does not exist,
an error will be returned during Pod creation.`,
																												"type": "boolean",
																											},
																											"path": map[string]interface{}{
																												"description": `The path within the volume from which to select the file.
Must be relative and may not contain the '..' path or start with '..'.`,
																												"type": "string",
																											},
																											"volumeName": map[string]interface{}{
																												"description": "The name of the volume mount containing the env file.",
																												"type":        "string",
																											},
																										},
																										"required": []interface{}{
																											"key",
																											"path",
																											"volumeName",
																										},
																										"type":                  "object",
																										"x-kubernetes-map-type": "atomic",
																									},
																									"resourceFieldRef": map[string]interface{}{
																										"description": `Selects a resource of the container: only resources limits and requests
(limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.`,
																										"properties": map[string]interface{}{
																											"containerName": map[string]interface{}{
																												"description": "Container name: required for volumes, optional for env vars",
																												"type":        "string",
																											},
																											"divisor": map[string]interface{}{
																												"anyOf": []interface{}{
																													map[string]interface{}{
																														"type": "integer",
																													},
																													map[string]interface{}{
																														"type": "string",
																													},
																												},
																												"description":                "Specifies the output format of the exposed resources, defaults to \"1\"",
																												"pattern":                    `^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$`,
																												"x-kubernetes-int-or-string": true,
																											},
																											"resource": map[string]interface{}{
																												"description": "Required: resource to select",
																												"type":        "string",
																											},
																										},
																										"required": []interface{}{
																											"resource",
																										},
																										"type":                  "object",
																										"x-kubernetes-map-type": "atomic",
																									},
																									"secretKeyRef": map[string]interface{}{
																										"description": "Selects a key of a secret in the pod's namespace",
																										"properties": map[string]interface{}{
																											"key": map[string]interface{}{
																												"description": "The key of the secret to select from.  Must be a valid secret key.",
																												"type":        "string",
																											},
																											"name": map[string]interface{}{
																												"default": "",
																												"description": `Name of the referent.
This field is effectively required, but due to backwards compatibility is
allowed to be empty. Instances of this type with an empty value here are
almost certainly wrong.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names`,
																												"type": "string",
																											},
																											"optional": map[string]interface{}{
																												"description": "Specify whether the Secret or its key must be defined",
																												"type":        "boolean",
																											},
																										},
																										"required": []interface{}{
																											"key",
																										},
																										"type":                  "object",
																										"x-kubernetes-map-type": "atomic",
																									},
																								},
																								"type": "object",
																							},
																						},
																						"required": []interface{}{
																							"name",
																						},
																						"type": "object",
																					},
																					"type": "array",
																					"x-kubernetes-list-map-keys": []interface{}{
																						"name",
																					},
																					"x-kubernetes-list-type": "map",
																				},
																				"envFrom": map[string]interface{}{
																					"description": `List of sources to populate environment variables in the container.
The keys defined within a source may consist of any printable ASCII characters except '='.
When a key exists in multiple
sources, the value associated with the last source will take precedence.
Values defined by an Env with a duplicate key will take precedence.
Cannot be updated.`,
																					"items": map[string]interface{}{
																						"description": "EnvFromSource represents the source of a set of ConfigMaps or Secrets",
																						"properties": map[string]interface{}{
																							"configMapRef": map[string]interface{}{
																								"description": "The ConfigMap to select from",
																								"properties": map[string]interface{}{
																									"name": map[string]interface{}{
																										"default": "",
																										"description": `Name of the referent.
This field is effectively required, but due to backwards compatibility is
allowed to be empty. Instances of this type with an empty value here are
almost certainly wrong.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names`,
																										"type": "string",
																									},
																									"optional": map[string]interface{}{
																										"description": "Specify whether the ConfigMap must be defined",
																										"type":        "boolean",
																									},
																								},
																								"type":                  "object",
																								"x-kubernetes-map-type": "atomic",
																							},
																							"prefix": map[string]interface{}{
																								"description": `Optional text to prepend to the name of each environment variable.
May consist of any printable ASCII characters except '='.`,
																								"type": "string",
																							},
																							"secretRef": map[string]interface{}{
																								"description": "The Secret to select from",
																								"properties": map[string]interface{}{
																									"name": map[string]interface{}{
																										"default": "",
																										"description": `Name of the referent.
This field is effectively required, but due to backwards compatibility is
allowed to be empty. Instances of this type with an empty value here are
almost certainly wrong.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names`,
																										"type": "string",
																									},
																									"optional": map[string]interface{}{
																										"description": "Specify whether the Secret must be defined",
																										"type":        "boolean",
																									},
																								},
																								"type":                  "object",
																								"x-kubernetes-map-type": "atomic",
																							},
																						},
																						"type": "object",
																					},
																					"type":                   "array",
																					"x-kubernetes-list-type": "atomic",
																				},
																				"image": map[string]interface{}{
																					"description": `Container image name.
More info: https://kubernetes.io/docs/concepts/containers/images
This field is optional to allow higher level config management to default or override
container images in workload controllers like Deployments and StatefulSets.`,
																					"type": "string",
																				},
																				"imagePullPolicy": map[string]interface{}{
																					"description": `Image pull policy.
One of Always, Never, IfNotPresent.
Defaults to Always if :latest tag is specified, or IfNotPresent otherwise.
Cannot be updated.
More info: https://kubernetes.io/docs/concepts/containers/images#updating-images`,
																					"type": "string",
																				},
																				"lifecycle": map[string]interface{}{
																					"description": `Actions that the management system should take in response to container lifecycle events.
Cannot be updated.`,
																					"properties": map[string]interface{}{
																						"postStart": map[string]interface{}{
																							"description": `PostStart is called immediately after a container is created. If the handler fails,
the container is terminated and restarted according to its restart policy.
Other management of the container blocks until the hook completes.
More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks`,
																							"properties": map[string]interface{}{
																								"exec": map[string]interface{}{
																									"description": "Exec specifies a command to execute in the container.",
																									"properties": map[string]interface{}{
																										"command": map[string]interface{}{
																											"description": `Command is the command line to execute inside the container, the working directory for the
command  is root ('/') in the container's filesystem. The command is simply exec'd, it is
not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use
a shell, you need to explicitly call out to that shell.
Exit status of 0 is treated as live/healthy and non-zero is unhealthy.`,
																											"items": map[string]interface{}{
																												"type": "string",
																											},
																											"type":                   "array",
																											"x-kubernetes-list-type": "atomic",
																										},
																									},
																									"type": "object",
																								},
																								"httpGet": map[string]interface{}{
																									"description": "HTTPGet specifies an HTTP GET request to perform.",
																									"properties": map[string]interface{}{
																										"host": map[string]interface{}{
																											"description": `Host name to connect to, defaults to the pod IP. You probably want to set
"Host" in httpHeaders instead.`,
																											"type": "string",
																										},
																										"httpHeaders": map[string]interface{}{
																											"description": "Custom headers to set in the request. HTTP allows repeated headers.",
																											"items": map[string]interface{}{
																												"description": "HTTPHeader describes a custom header to be used in HTTP probes",
																												"properties": map[string]interface{}{
																													"name": map[string]interface{}{
																														"description": `The header field name.
This will be canonicalized upon output, so case-variant names will be understood as the same header.`,
																														"type": "string",
																													},
																													"value": map[string]interface{}{
																														"description": "The header field value",
																														"type":        "string",
																													},
																												},
																												"required": []interface{}{
																													"name",
																													"value",
																												},
																												"type": "object",
																											},
																											"type":                   "array",
																											"x-kubernetes-list-type": "atomic",
																										},
																										"path": map[string]interface{}{
																											"description": "Path to access on the HTTP server.",
																											"type":        "string",
																										},
																										"port": map[string]interface{}{
																											"anyOf": []interface{}{
																												map[string]interface{}{
																													"type": "integer",
																												},
																												map[string]interface{}{
																													"type": "string",
																												},
																											},
																											"description": `Name or number of the port to access on the container.
Number must be in the range 1 to 65535.
Name must be an IANA_SVC_NAME.`,
																											"x-kubernetes-int-or-string": true,
																										},
																										"scheme": map[string]interface{}{
																											"description": `Scheme to use for connecting to the host.
Defaults to HTTP.`,
																											"type": "string",
																										},
																									},
																									"required": []interface{}{
																										"port",
																									},
																									"type": "object",
																								},
																								"sleep": map[string]interface{}{
																									"description": "Sleep represents a duration that the container should sleep.",
																									"properties": map[string]interface{}{
																										"seconds": map[string]interface{}{
																											"description": "Seconds is the number of seconds to sleep.",
																											"format":      "int64",
																											"type":        "integer",
																										},
																									},
																									"required": []interface{}{
																										"seconds",
																									},
																									"type": "object",
																								},
																								"tcpSocket": map[string]interface{}{
																									"description": `Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept
for backward compatibility. There is no validation of this field and
lifecycle hooks will fail at runtime when it is specified.`,
																									"properties": map[string]interface{}{
																										"host": map[string]interface{}{
																											"description": "Optional: Host name to connect to, defaults to the pod IP.",
																											"type":        "string",
																										},
																										"port": map[string]interface{}{
																											"anyOf": []interface{}{
																												map[string]interface{}{
																													"type": "integer",
																												},
																												map[string]interface{}{
																													"type": "string",
																												},
																											},
																											"description": `Number or name of the port to access on the container.
Number must be in the range 1 to 65535.
Name must be an IANA_SVC_NAME.`,
																											"x-kubernetes-int-or-string": true,
																										},
																									},
																									"required": []interface{}{
																										"port",
																									},
																									"type": "object",
																								},
																							},
																							"type": "object",
																						},
																						"preStop": map[string]interface{}{
																							"description": `PreStop is called immediately before a container is terminated due to an
API request or management event such as liveness/startup probe failure,
preemption, resource contention, etc. The handler is not called if the
container crashes or exits. The Pod's termination grace period countdown begins before the
PreStop hook is executed. Regardless of the outcome of the handler, the
container will eventually terminate within the Pod's termination grace
period (unless delayed by finalizers). Other management of the container blocks until the hook completes
or until the termination grace period is reached.
More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks`,
																							"properties": map[string]interface{}{
																								"exec": map[string]interface{}{
																									"description": "Exec specifies a command to execute in the container.",
																									"properties": map[string]interface{}{
																										"command": map[string]interface{}{
																											"description": `Command is the command line to execute inside the container, the working directory for the
command  is root ('/') in the container's filesystem. The command is simply exec'd, it is
not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use
a shell, you need to explicitly call out to that shell.
Exit status of 0 is treated as live/healthy and non-zero is unhealthy.`,
																											"items": map[string]interface{}{
																												"type": "string",
																											},
																											"type":                   "array",
																											"x-kubernetes-list-type": "atomic",
																										},
																									},
																									"type": "object",
																								},
																								"httpGet": map[string]interface{}{
																									"description": "HTTPGet specifies an HTTP GET request to perform.",
																									"properties": map[string]interface{}{
																										"host": map[string]interface{}{
																											"description": `Host name to connect to, defaults to the pod IP. You probably want to set
"Host" in httpHeaders instead.`,
																											"type": "string",
																										},
																										"httpHeaders": map[string]interface{}{
																											"description": "Custom headers to set in the request. HTTP allows repeated headers.",
																											"items": map[string]interface{}{
																												"description": "HTTPHeader describes a custom header to be used in HTTP probes",
																												"properties": map[string]interface{}{
																													"name": map[string]interface{}{
																														"description": `The header field name.
This will be canonicalized upon output, so case-variant names will be understood as the same header.`,
																														"type": "string",
																													},
																													"value": map[string]interface{}{
																														"description": "The header field value",
																														"type":        "string",
																													},
																												},
																												"required": []interface{}{
																													"name",
																													"value",
																												},
																												"type": "object",
																											},
																											"type":                   "array",
																											"x-kubernetes-list-type": "atomic",
																										},
																										"path": map[string]interface{}{
																											"description": "Path to access on the HTTP server.",
																											"type":        "string",
																										},
																										"port": map[string]interface{}{
																											"anyOf": []interface{}{
																												map[string]interface{}{
																													"type": "integer",
																												},
																												map[string]interface{}{
																													"type": "string",
																												},
																											},
																											"description": `Name or number of the port to access on the container.
Number must be in the range 1 to 65535.
Name must be an IANA_SVC_NAME.`,
																											"x-kubernetes-int-or-string": true,
																										},
																										"scheme": map[string]interface{}{
																											"description": `Scheme to use for connecting to the host.
Defaults to HTTP.`,
																											"type": "string",
																										},
																									},
																									"required": []interface{}{
																										"port",
																									},
																									"type": "object",
																								},
																								"sleep": map[string]interface{}{
																									"description": "Sleep represents a duration that the container should sleep.",
																									"properties": map[string]interface{}{
																										"seconds": map[string]interface{}{
																											"description": "Seconds is the number of seconds to sleep.",
																											"format":      "int64",
																											"type":        "integer",
																										},
																									},
																									"required": []interface{}{
																										"seconds",
																									},
																									"type": "object",
																								},
																								"tcpSocket": map[string]interface{}{
																									"description": `Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept
for backward compatibility. There is no validation of this field and
lifecycle hooks will fail at runtime when it is specified.`,
																									"properties": map[string]interface{}{
																										"host": map[string]interface{}{
																											"description": "Optional: Host name to connect to, defaults to the pod IP.",
																											"type":        "string",
																										},
																										"port": map[string]interface{}{
																											"anyOf": []interface{}{
																												map[string]interface{}{
																													"type": "integer",
																												},
																												map[string]interface{}{
																													"type": "string",
																												},
																											},
																											"description": `Number or name of the port to access on the container.
Number must be in the range 1 to 65535.
Name must be an IANA_SVC_NAME.`,
																											"x-kubernetes-int-or-string": true,
																										},
																									},
																									"required": []interface{}{
																										"port",
																									},
																									"type": "object",
																								},
																							},
																							"type": "object",
																						},
																						"stopSignal": map[string]interface{}{
																							"description": `StopSignal defines which signal will be sent to a container when it is being stopped.
If not specified, the default is defined by the container runtime in use.
StopSignal can only be set for Pods with a non-empty .spec.os.name`,
																							"type": "string",
																						},
																					},
																					"type": "object",
																				},
																				"livenessProbe": map[string]interface{}{
																					"description": `Periodic probe of container liveness.
Container will be restarted if the probe fails.
Cannot be updated.
More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes`,
																					"properties": map[string]interface{}{
																						"exec": map[string]interface{}{
																							"description": "Exec specifies a command to execute in the container.",
																							"properties": map[string]interface{}{
																								"command": map[string]interface{}{
																									"description": `Command is the command line to execute inside the container, the working directory for the
command  is root ('/') in the container's filesystem. The command is simply exec'd, it is
not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use
a shell, you need to explicitly call out to that shell.
Exit status of 0 is treated as live/healthy and non-zero is unhealthy.`,
																									"items": map[string]interface{}{
																										"type": "string",
																									},
																									"type":                   "array",
																									"x-kubernetes-list-type": "atomic",
																								},
																							},
																							"type": "object",
																						},
																						"failureThreshold": map[string]interface{}{
																							"description": `Minimum consecutive failures for the probe to be considered failed after having succeeded.
Defaults to 3. Minimum value is 1.`,
																							"format": "int32",
																							"type":   "integer",
																						},
																						"grpc": map[string]interface{}{
																							"description": "GRPC specifies a GRPC HealthCheckRequest.",
																							"properties": map[string]interface{}{
																								"port": map[string]interface{}{
																									"description": "Port number of the gRPC service. Number must be in the range 1 to 65535.",
																									"format":      "int32",
																									"type":        "integer",
																								},
																								"service": map[string]interface{}{
																									"default": "",
																									"description": `Service is the name of the service to place in the gRPC HealthCheckRequest
(see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).

If this is not specified, the default behavior is defined by gRPC.`,
																									"type": "string",
																								},
																							},
																							"required": []interface{}{
																								"port",
																							},
																							"type": "object",
																						},
																						"httpGet": map[string]interface{}{
																							"description": "HTTPGet specifies an HTTP GET request to perform.",
																							"properties": map[string]interface{}{
																								"host": map[string]interface{}{
																									"description": `Host name to connect to, defaults to the pod IP. You probably want to set
"Host" in httpHeaders instead.`,
																									"type": "string",
																								},
																								"httpHeaders": map[string]interface{}{
																									"description": "Custom headers to set in the request. HTTP allows repeated headers.",
																									"items": map[string]interface{}{
																										"description": "HTTPHeader describes a custom header to be used in HTTP probes",
																										"properties": map[string]interface{}{
																											"name": map[string]interface{}{
																												"description": `The header field name.
This will be canonicalized upon output, so case-variant names will be understood as the same header.`,
																												"type": "string",
																											},
																											"value": map[string]interface{}{
																												"description": "The header field value",
																												"type":        "string",
																											},
																										},
																										"required": []interface{}{
																											"name",
																											"value",
																										},
																										"type": "object",
																									},
																									"type":                   "array",
																									"x-kubernetes-list-type": "atomic",
																								},
																								"path": map[string]interface{}{
																									"description": "Path to access on the HTTP server.",
																									"type":        "string",
																								},
																								"port": map[string]interface{}{
																									"anyOf": []interface{}{
																										map[string]interface{}{
																											"type": "integer",
																										},
																										map[string]interface{}{
																											"type": "string",
																										},
																									},
																									"description": `Name or number of the port to access on the container.
Number must be in the range 1 to 65535.
Name must be an IANA_SVC_NAME.`,
																									"x-kubernetes-int-or-string": true,
																								},
																								"scheme": map[string]interface{}{
																									"description": `Scheme to use for connecting to the host.
Defaults to HTTP.`,
																									"type": "string",
																								},
																							},
																							"required": []interface{}{
																								"port",
																							},
																							"type": "object",
																						},
																						"initialDelaySeconds": map[string]interface{}{
																							"description": `Number of seconds after the container has started before liveness probes are initiated.
More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes`,
																							"format": "int32",
																							"type":   "integer",
																						},
																						"periodSeconds": map[string]interface{}{
																							"description": `How often (in seconds) to perform the probe.
Default to 10 seconds. Minimum value is 1.`,
																							"format": "int32",
																							"type":   "integer",
																						},
																						"successThreshold": map[string]interface{}{
																							"description": `Minimum consecutive successes for the probe to be considered successful after having failed.
Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.`,
																							"format": "int32",
																							"type":   "integer",
																						},
																						"tcpSocket": map[string]interface{}{
																							"description": "TCPSocket specifies a connection to a TCP port.",
																							"properties": map[string]interface{}{
																								"host": map[string]interface{}{
																									"description": "Optional: Host name to connect to, defaults to the pod IP.",
																									"type":        "string",
																								},
																								"port": map[string]interface{}{
																									"anyOf": []interface{}{
																										map[string]interface{}{
																											"type": "integer",
																										},
																										map[string]interface{}{
																											"type": "string",
																										},
																									},
																									"description": `Number or name of the port to access on the container.
Number must be in the range 1 to 65535.
Name must be an IANA_SVC_NAME.`,
																									"x-kubernetes-int-or-string": true,
																								},
																							},
																							"required": []interface{}{
																								"port",
																							},
																							"type": "object",
																						},
																						"terminationGracePeriodSeconds": map[string]interface{}{
																							"description": `Optional duration in seconds the pod needs to terminate gracefully upon probe failure.
The grace period is the duration in seconds after the processes running in the pod are sent
a termination signal and the time when the processes are forcibly halted with a kill signal.
Set this value longer than the expected cleanup time for your process.
If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this
value overrides the value provided by the pod spec.
Value must be non-negative integer. The value zero indicates stop immediately via
the kill signal (no opportunity to shut down).
This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate.
Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.`,
																							"format": "int64",
																							"type":   "integer",
																						},
																						"timeoutSeconds": map[string]interface{}{
																							"description": `Number of seconds after which the probe times out.
Defaults to 1 second. Minimum value is 1.
More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes`,
																							"format": "int32",
																							"type":   "integer",
																						},
																					},
																					"type": "object",
																				},
																				"name": map[string]interface{}{
																					"description": `Name of the container specified as a DNS_LABEL.
Each container in a pod must have a unique name (DNS_LABEL).
Cannot be updated.`,
																					"type": "string",
																				},
																				"ports": map[string]interface{}{
																					"description": `List of ports to expose from the container. Not specifying a port here
DOES NOT prevent that port from being exposed. Any port which is
listening on the default "0.0.0.0" address inside a container will be
accessible from the network.
Modifying this array with strategic merge patch may corrupt the data.
For more information See https://github.com/kubernetes/kubernetes/issues/108255.
Cannot be updated.`,
																					"items": map[string]interface{}{
																						"description": "ContainerPort represents a network port in a single container.",
																						"properties": map[string]interface{}{
																							"containerPort": map[string]interface{}{
																								"description": `Number of port to expose on the pod's IP address.
This must be a valid port number, 0 < x < 65536.`,
																								"format": "int32",
																								"type":   "integer",
																							},
																							"hostIP": map[string]interface{}{
																								"description": "What host IP to bind the external port to.",
																								"type":        "string",
																							},
																							"hostPort": map[string]interface{}{
																								"description": `Number of port to expose on the host.
If specified, this must be a valid port number, 0 < x < 65536.
If HostNetwork is specified, this must match ContainerPort.
Most containers do not need this.`,
																								"format": "int32",
																								"type":   "integer",
																							},
																							"name": map[string]interface{}{
																								"description": `If specified, this must be an IANA_SVC_NAME and unique within the pod. Each
named port in a pod must have a unique name. Name for the port that can be
referred to by services.`,
																								"type": "string",
																							},
																							"protocol": map[string]interface{}{
																								"default": "TCP",
																								"description": `Protocol for port. Must be UDP, TCP, or SCTP.
Defaults to "TCP".`,
																								"type": "string",
																							},
																						},
																						"required": []interface{}{
																							"containerPort",
																							"protocol",
																						},
																						"type": "object",
																					},
																					"type": "array",
																					"x-kubernetes-list-map-keys": []interface{}{
																						"containerPort",
																						"protocol",
																					},
																					"x-kubernetes-list-type": "map",
																				},
																				"readinessProbe": map[string]interface{}{
																					"description": `Periodic probe of container service readiness.
Container will be removed from service endpoints if the probe fails.
Cannot be updated.
More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes`,
																					"properties": map[string]interface{}{
																						"exec": map[string]interface{}{
																							"description": "Exec specifies a command to execute in the container.",
																							"properties": map[string]interface{}{
																								"command": map[string]interface{}{
																									"description": `Command is the command line to execute inside the container, the working directory for the
command  is root ('/') in the container's filesystem. The command is simply exec'd, it is
not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use
a shell, you need to explicitly call out to that shell.
Exit status of 0 is treated as live/healthy and non-zero is unhealthy.`,
																									"items": map[string]interface{}{
																										"type": "string",
																									},
																									"type":                   "array",
																									"x-kubernetes-list-type": "atomic",
																								},
																							},
																							"type": "object",
																						},
																						"failureThreshold": map[string]interface{}{
																							"description": `Minimum consecutive failures for the probe to be considered failed after having succeeded.
Defaults to 3. Minimum value is 1.`,
																							"format": "int32",
																							"type":   "integer",
																						},
																						"grpc": map[string]interface{}{
																							"description": "GRPC specifies a GRPC HealthCheckRequest.",
																							"properties": map[string]interface{}{
																								"port": map[string]interface{}{
																									"description": "Port number of the gRPC service. Number must be in the range 1 to 65535.",
																									"format":      "int32",
																									"type":        "integer",
																								},
																								"service": map[string]interface{}{
																									"default": "",
																									"description": `Service is the name of the service to place in the gRPC HealthCheckRequest
(see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).

If this is not specified, the default behavior is defined by gRPC.`,
																									"type": "string",
																								},
																							},
																							"required": []interface{}{
																								"port",
																							},
																							"type": "object",
																						},
																						"httpGet": map[string]interface{}{
																							"description": "HTTPGet specifies an HTTP GET request to perform.",
																							"properties": map[string]interface{}{
																								"host": map[string]interface{}{
																									"description": `Host name to connect to, defaults to the pod IP. You probably want to set
"Host" in httpHeaders instead.`,
																									"type": "string",
																								},
																								"httpHeaders": map[string]interface{}{
																									"description": "Custom headers to set in the request. HTTP allows repeated headers.",
																									"items": map[string]interface{}{
																										"description": "HTTPHeader describes a custom header to be used in HTTP probes",
																										"properties": map[string]interface{}{
																											"name": map[string]interface{}{
																												"description": `The header field name.
This will be canonicalized upon output, so case-variant names will be understood as the same header.`,
																												"type": "string",
																											},
																											"value": map[string]interface{}{
																												"description": "The header field value",
																												"type":        "string",
																											},
																										},
																										"required": []interface{}{
																											"name",
																											"value",
																										},
																										"type": "object",
																									},
																									"type":                   "array",
																									"x-kubernetes-list-type": "atomic",
																								},
																								"path": map[string]interface{}{
																									"description": "Path to access on the HTTP server.",
																									"type":        "string",
																								},
																								"port": map[string]interface{}{
																									"anyOf": []interface{}{
																										map[string]interface{}{
																											"type": "integer",
																										},
																										map[string]interface{}{
																											"type": "string",
																										},
																									},
																									"description": `Name or number of the port to access on the container.
Number must be in the range 1 to 65535.
Name must be an IANA_SVC_NAME.`,
																									"x-kubernetes-int-or-string": true,
																								},
																								"scheme": map[string]interface{}{
																									"description": `Scheme to use for connecting to the host.
Defaults to HTTP.`,
																									"type": "string",
																								},
																							},
																							"required": []interface{}{
																								"port",
																							},
																							"type": "object",
																						},
																						"initialDelaySeconds": map[string]interface{}{
																							"description": `Number of seconds after the container has started before liveness probes are initiated.
More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes`,
																							"format": "int32",
																							"type":   "integer",
																						},
																						"periodSeconds": map[string]interface{}{
																							"description": `How often (in seconds) to perform the probe.
Default to 10 seconds. Minimum value is 1.`,
																							"format": "int32",
																							"type":   "integer",
																						},
																						"successThreshold": map[string]interface{}{
																							"description": `Minimum consecutive successes for the probe to be considered successful after having failed.
Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.`,
																							"format": "int32",
																							"type":   "integer",
																						},
																						"tcpSocket": map[string]interface{}{
																							"description": "TCPSocket specifies a connection to a TCP port.",
																							"properties": map[string]interface{}{
																								"host": map[string]interface{}{
																									"description": "Optional: Host name to connect to, defaults to the pod IP.",
																									"type":        "string",
																								},
																								"port": map[string]interface{}{
																									"anyOf": []interface{}{
																										map[string]interface{}{
																											"type": "integer",
																										},
																										map[string]interface{}{
																											"type": "string",
																										},
																									},
																									"description": `Number or name of the port to access on the container.
Number must be in the range 1 to 65535.
Name must be an IANA_SVC_NAME.`,
																									"x-kubernetes-int-or-string": true,
																								},
																							},
																							"required": []interface{}{
																								"port",
																							},
																							"type": "object",
																						},
																						"terminationGracePeriodSeconds": map[string]interface{}{
																							"description": `Optional duration in seconds the pod needs to terminate gracefully upon probe failure.
The grace period is the duration in seconds after the processes running in the pod are sent
a termination signal and the time when the processes are forcibly halted with a kill signal.
Set this value longer than the expected cleanup time for your process.
If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this
value overrides the value provided by the pod spec.
Value must be non-negative integer. The value zero indicates stop immediately via
the kill signal (no opportunity to shut down).
This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate.
Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.`,
																							"format": "int64",
																							"type":   "integer",
																						},
																						"timeoutSeconds": map[string]interface{}{
																							"description": `Number of seconds after which the probe times out.
Defaults to 1 second. Minimum value is 1.
More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes`,
																							"format": "int32",
																							"type":   "integer",
																						},
																					},
																					"type": "object",
																				},
																				"resizePolicy": map[string]interface{}{
																					"description": `Resources resize policy for the container.
This field cannot be set on ephemeral containers.`,
																					"items": map[string]interface{}{
																						"description": "ContainerResizePolicy represents resource resize policy for the container.",
																						"properties": map[string]interface{}{
																							"resourceName": map[string]interface{}{
																								"description": `Name of the resource to which this resource resize policy applies.
Supported values: cpu, memory.`,
																								"type": "string",
																							},
																							"restartPolicy": map[string]interface{}{
																								"description": `Restart policy to apply when specified resource is resized.
If not specified, it defaults to NotRequired.`,
																								"type": "string",
																							},
																						},
																						"required": []interface{}{
																							"resourceName",
																							"restartPolicy",
																						},
																						"type": "object",
																					},
																					"type":                   "array",
																					"x-kubernetes-list-type": "atomic",
																				},
																				"resources": map[string]interface{}{
																					"description": `Compute Resources required by this container.
Cannot be updated.
More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/`,
																					"properties": map[string]interface{}{
																						"claims": map[string]interface{}{
																							"description": `Claims lists the names of resources, defined in spec.resourceClaims,
that are used by this container.

This field depends on the
DynamicResourceAllocation feature gate.

This field is immutable. It can only be set for containers.`,
																							"items": map[string]interface{}{
																								"description": "ResourceClaim references one entry in PodSpec.ResourceClaims.",
																								"properties": map[string]interface{}{
																									"name": map[string]interface{}{
																										"description": `Name must match the name of one entry in pod.spec.resourceClaims of
the Pod where this field is used. It makes that resource available
inside a container.`,
																										"type": "string",
																									},
																									"request": map[string]interface{}{
																										"description": `Request is the name chosen for a request in the referenced claim.
If empty, everything from the claim is made available, otherwise
only the result of this request.`,
																										"type": "string",
																									},
																								},
																								"required": []interface{}{
																									"name",
																								},
																								"type": "object",
																							},
																							"type": "array",
																							"x-kubernetes-list-map-keys": []interface{}{
																								"name",
																							},
																							"x-kubernetes-list-type": "map",
																						},
																						"limits": map[string]interface{}{
																							"additionalProperties": map[string]interface{}{
																								"anyOf": []interface{}{
																									map[string]interface{}{
																										"type": "integer",
																									},
																									map[string]interface{}{
																										"type": "string",
																									},
																								},
																								"pattern":                    `^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$`,
																								"x-kubernetes-int-or-string": true,
																							},
																							"description": `Limits describes the maximum amount of compute resources allowed.
More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/`,
																							"type": "object",
																						},
																						"requests": map[string]interface{}{
																							"additionalProperties": map[string]interface{}{
																								"anyOf": []interface{}{
																									map[string]interface{}{
																										"type": "integer",
																									},
																									map[string]interface{}{
																										"type": "string",
																									},
																								},
																								"pattern":                    `^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$`,
																								"x-kubernetes-int-or-string": true,
																							},
																							"description": `Requests describes the minimum amount of compute resources required.
If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,
otherwise to an implementation-defined value. Requests cannot exceed Limits.
More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/`,
																							"type": "object",
																						},
																					},
																					"type": "object",
																				},
																				"restartPolicy": map[string]interface{}{
																					"description": `RestartPolicy defines the restart behavior of individual containers in a pod.
This overrides the pod-level restart policy. When this field is not specified,
the restart behavior is defined by the Pod's restart policy and the container type.
Additionally, setting the RestartPolicy as "Always" for the init container will
have the following effect:
this init container will be continually restarted on
exit until all regular containers have terminated. Once all regular
containers have completed, all init containers with restartPolicy "Always"
will be shut down. This lifecycle differs from normal init containers and
is often referred to as a "sidecar" container. Although this init
container still starts in the init container sequence, it does not wait
for the container to complete before proceeding to the next init
container. Instead, the next init container starts immediately after this
init container is started, or after any startupProbe has successfully
completed.`,
																					"type": "string",
																				},
																				"restartPolicyRules": map[string]interface{}{
																					"description": `Represents a list of rules to be checked to determine if the
container should be restarted on exit. The rules are evaluated in
order. Once a rule matches a container exit condition, the remaining
rules are ignored. If no rule matches the container exit condition,
the Container-level restart policy determines the whether the container
is restarted or not. Constraints on the rules:
- At most 20 rules are allowed.
- Rules can have the same action.
- Identical rules are not forbidden in validations.
When rules are specified, container MUST set RestartPolicy explicitly
even it if matches the Pod's RestartPolicy.`,
																					"items": map[string]interface{}{
																						"description": "ContainerRestartRule describes how a container exit is handled.",
																						"properties": map[string]interface{}{
																							"action": map[string]interface{}{
																								"description": `Specifies the action taken on a container exit if the requirements
are satisfied. The only possible value is "Restart" to restart the
container.`,
																								"type": "string",
																							},
																							"exitCodes": map[string]interface{}{
																								"description": "Represents the exit codes to check on container exits.",
																								"properties": map[string]interface{}{
																									"operator": map[string]interface{}{
																										"description": `Represents the relationship between the container exit code(s) and the
specified values. Possible values are:
- In: the requirement is satisfied if the container exit code is in the
  set of specified values.
- NotIn: the requirement is satisfied if the container exit code is
  not in the set of specified values.`,
																										"type": "string",
																									},
																									"values": map[string]interface{}{
																										"description": `Specifies the set of values to check for container exit codes.
At most 255 elements are allowed.`,
																										"items": map[string]interface{}{
																											"format": "int32",
																											"type":   "integer",
																										},
																										"type":                   "array",
																										"x-kubernetes-list-type": "set",
																									},
																								},
																								"required": []interface{}{
																									"operator",
																								},
																								"type": "object",
																							},
																						},
																						"required": []interface{}{
																							"action",
																						},
																						"type": "object",
																					},
																					"type":                   "array",
																					"x-kubernetes-list-type": "atomic",
																				},
																				"securityContext": map[string]interface{}{
																					"description": `SecurityContext defines the security options the container should be run with.
If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext.
More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/`,
																					"properties": map[string]interface{}{
																						"allowPrivilegeEscalation": map[string]interface{}{
																							"description": `AllowPrivilegeEscalation controls whether a process can gain more
privileges than its parent process. This bool directly controls if
the no_new_privs flag will be set on the container process.
AllowPrivilegeEscalation is true always when the container is:
1) run as Privileged
2) has CAP_SYS_ADMIN
Note that this field cannot be set when spec.os.name is windows.`,
																							"type": "boolean",
																						},
																						"appArmorProfile": map[string]interface{}{
																							"description": `appArmorProfile is the AppArmor options to use by this container. If set, this profile
overrides the pod's appArmorProfile.
Note that this field cannot be set when spec.os.name is windows.`,
																							"properties": map[string]interface{}{
																								"localhostProfile": map[string]interface{}{
																									"description": `localhostProfile indicates a profile loaded on the node that should be used.
The profile must be preconfigured on the node to work.
Must match the loaded name of the profile.
Must be set if and only if type is "Localhost".`,
																									"type": "string",
																								},
																								"type": map[string]interface{}{
																									"description": `type indicates which kind of AppArmor profile will be applied.
Valid options are:
  Localhost - a profile pre-loaded on the node.
  RuntimeDefault - the container runtime's default profile.
  Unconfined - no AppArmor enforcement.`,
																									"type": "string",
																								},
																							},
																							"required": []interface{}{
																								"type",
																							},
																							"type": "object",
																						},
																						"capabilities": map[string]interface{}{
																							"description": `The capabilities to add/drop when running containers.
Defaults to the default set of capabilities granted by the container runtime.
Note that this field cannot be set when spec.os.name is windows.`,
																							"properties": map[string]interface{}{
																								"add": map[string]interface{}{
																									"description": "Added capabilities",
																									"items": map[string]interface{}{
																										"description": "Capability represent POSIX capabilities type",
																										"type":        "string",
																									},
																									"type":                   "array",
																									"x-kubernetes-list-type": "atomic",
																								},
																								"drop": map[string]interface{}{
																									"description": "Removed capabilities",
																									"items": map[string]interface{}{
																										"description": "Capability represent POSIX capabilities type",
																										"type":        "string",
																									},
																									"type":                   "array",
																									"x-kubernetes-list-type": "atomic",
																								},
																							},
																							"type": "object",
																						},
																						"privileged": map[string]interface{}{
																							"description": `Run container in privileged mode.
Processes in privileged containers are essentially equivalent to root on the host.
Defaults to false.
Note that this field cannot be set when spec.os.name is windows.`,
																							"type": "boolean",
																						},
																						"procMount": map[string]interface{}{
																							"description": `procMount denotes the type of proc mount to use for the containers.
The default value is Default which uses the container runtime defaults for
readonly paths and masked paths.
Note that this field cannot be set when spec.os.name is windows.`,
																							"type": "string",
																						},
																						"readOnlyRootFilesystem": map[string]interface{}{
																							"description": `Whether this container has a read-only root filesystem.
Default is false.
Note that this field cannot be set when spec.os.name is windows.`,
																							"type": "boolean",
																						},
																						"runAsGroup": map[string]interface{}{
																							"description": `The GID to run the entrypoint of the container process.
Uses runtime default if unset.
May also be set in PodSecurityContext.  If set in both SecurityContext and
PodSecurityContext, the value specified in SecurityContext takes precedence.
Note that this field cannot be set when spec.os.name is windows.`,
																							"format": "int64",
																							"type":   "integer",
																						},
																						"runAsNonRoot": map[string]interface{}{
																							"description": `Indicates that the container must run as a non-root user.
If true, the Kubelet will validate the image at runtime to ensure that it
does not run as UID 0 (root) and fail to start the container if it does.
If unset or false, no such validation will be performed.
May also be set in PodSecurityContext.  If set in both SecurityContext and
PodSecurityContext, the value specified in SecurityContext takes precedence.`,
																							"type": "boolean",
																						},
																						"runAsUser": map[string]interface{}{
																							"description": `The UID to run the entrypoint of the container process.
Defaults to user specified in image metadata if unspecified.
May also be set in PodSecurityContext.  If set in both SecurityContext and
PodSecurityContext, the value specified in SecurityContext takes precedence.
Note that this field cannot be set when spec.os.name is windows.`,
																							"format": "int64",
																							"type":   "integer",
																						},
																						"seLinuxOptions": map[string]interface{}{
																							"description": `The SELinux context to be applied to the container.
If unspecified, the container runtime will allocate a random SELinux context for each
container.  May also be set in PodSecurityContext.  If set in both SecurityContext and
PodSecurityContext, the value specified in SecurityContext takes precedence.
Note that this field cannot be set when spec.os.name is windows.`,
																							"properties": map[string]interface{}{
																								"level": map[string]interface{}{
																									"description": "Level is SELinux level label that applies to the container.",
																									"type":        "string",
																								},
																								"role": map[string]interface{}{
																									"description": "Role is a SELinux role label that applies to the container.",
																									"type":        "string",
																								},
																								"type": map[string]interface{}{
																									"description": "Type is a SELinux type label that applies to the container.",
																									"type":        "string",
																								},
																								"user": map[string]interface{}{
																									"description": "User is a SELinux user label that applies to the container.",
																									"type":        "string",
																								},
																							},
																							"type": "object",
																						},
																						"seccompProfile": map[string]interface{}{
																							"description": `The seccomp options to use by this container. If seccomp options are
provided at both the pod & container level, the container options
override the pod options.
Note that this field cannot be set when spec.os.name is windows.`,
																							"properties": map[string]interface{}{
																								"localhostProfile": map[string]interface{}{
																									"description": `localhostProfile indicates a profile defined in a file on the node should be used.
The profile must be preconfigured on the node to work.
Must be a descending path, relative to the kubelet's configured seccomp profile location.
Must be set if type is "Localhost". Must NOT be set for any other type.`,
																									"type": "string",
																								},
																								"type": map[string]interface{}{
																									"description": `type indicates which kind of seccomp profile will be applied.
Valid options are:

Localhost - a profile defined in a file on the node should be used.
RuntimeDefault - the container runtime default profile should be used.
Unconfined - no profile should be applied.`,
																									"type": "string",
																								},
																							},
																							"required": []interface{}{
																								"type",
																							},
																							"type": "object",
																						},
																						"windowsOptions": map[string]interface{}{
																							"description": `The Windows specific settings applied to all containers.
If unspecified, the options from the PodSecurityContext will be used.
If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
Note that this field cannot be set when spec.os.name is linux.`,
																							"properties": map[string]interface{}{
																								"gmsaCredentialSpec": map[string]interface{}{
																									"description": `GMSACredentialSpec is where the GMSA admission webhook
(https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the
GMSA credential spec named by the GMSACredentialSpecName field.`,
																									"type": "string",
																								},
																								"gmsaCredentialSpecName": map[string]interface{}{
																									"description": "GMSACredentialSpecName is the name of the GMSA credential spec to use.",
																									"type":        "string",
																								},
																								"hostProcess": map[string]interface{}{
																									"description": `HostProcess determines if a container should be run as a 'Host Process' container.
All of a Pod's containers must have the same effective HostProcess value
(it is not allowed to have a mix of HostProcess containers and non-HostProcess containers).
In addition, if HostProcess is true then HostNetwork must also be set to true.`,
																									"type": "boolean",
																								},
																								"runAsUserName": map[string]interface{}{
																									"description": `The UserName in Windows to run the entrypoint of the container process.
Defaults to the user specified in image metadata if unspecified.
May also be set in PodSecurityContext. If set in both SecurityContext and
PodSecurityContext, the value specified in SecurityContext takes precedence.`,
																									"type": "string",
																								},
																							},
																							"type": "object",
																						},
																					},
																					"type": "object",
																				},
																				"startupProbe": map[string]interface{}{
																					"description": `StartupProbe indicates that the Pod has successfully initialized.
If specified, no other probes are executed until this completes successfully.
If this probe fails, the Pod will be restarted, just as if the livenessProbe failed.
This can be used to provide different probe parameters at the beginning of a Pod's lifecycle,
when it might take a long time to load data or warm a cache, than during steady-state operation.
This cannot be updated.
More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes`,
																					"properties": map[string]interface{}{
																						"exec": map[string]interface{}{
																							"description": "Exec specifies a command to execute in the container.",
																							"properties": map[string]interface{}{
																								"command": map[string]interface{}{
																									"description": `Command is the command line to execute inside the container, the working directory for the
command  is root ('/') in the container's filesystem. The command is simply exec'd, it is
not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use
a shell, you need to explicitly call out to that shell.
Exit status of 0 is treated as live/healthy and non-zero is unhealthy.`,
																									"items": map[string]interface{}{
																										"type": "string",
																									},
																									"type":                   "array",
																									"x-kubernetes-list-type": "atomic",
																								},
																							},
																							"type": "object",
																						},
																						"failureThreshold": map[string]interface{}{
																							"description": `Minimum consecutive failures for the probe to be considered failed after having succeeded.
Defaults to 3. Minimum value is 1.`,
																							"format": "int32",
																							"type":   "integer",
																						},
																						"grpc": map[string]interface{}{
																							"description": "GRPC specifies a GRPC HealthCheckRequest.",
																							"properties": map[string]interface{}{
																								"port": map[string]interface{}{
																									"description": "Port number of the gRPC service. Number must be in the range 1 to 65535.",
																									"format":      "int32",
																									"type":        "integer",
																								},
																								"service": map[string]interface{}{
																									"default": "",
																									"description": `Service is the name of the service to place in the gRPC HealthCheckRequest
(see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).

If this is not specified, the default behavior is defined by gRPC.`,
																									"type": "string",
																								},
																							},
																							"required": []interface{}{
																								"port",
																							},
																							"type": "object",
																						},
																						"httpGet": map[string]interface{}{
																							"description": "HTTPGet specifies an HTTP GET request to perform.",
																							"properties": map[string]interface{}{
																								"host": map[string]interface{}{
																									"description": `Host name to connect to, defaults to the pod IP. You probably want to set
"Host" in httpHeaders instead.`,
																									"type": "string",
																								},
																								"httpHeaders": map[string]interface{}{
																									"description": "Custom headers to set in the request. HTTP allows repeated headers.",
																									"items": map[string]interface{}{
																										"description": "HTTPHeader describes a custom header to be used in HTTP probes",
																										"properties": map[string]interface{}{
																											"name": map[string]interface{}{
																												"description": `The header field name.
This will be canonicalized upon output, so case-variant names will be understood as the same header.`,
																												"type": "string",
																											},
																											"value": map[string]interface{}{
																												"description": "The header field value",
																												"type":        "string",
																											},
																										},
																										"required": []interface{}{
																											"name",
																											"value",
																										},
																										"type": "object",
																									},
																									"type":                   "array",
																									"x-kubernetes-list-type": "atomic",
																								},
																								"path": map[string]interface{}{
																									"description": "Path to access on the HTTP server.",
																									"type":        "string",
																								},
																								"port": map[string]interface{}{
																									"anyOf": []interface{}{
																										map[string]interface{}{
																											"type": "integer",
																										},
																										map[string]interface{}{
																											"type": "string",
																										},
																									},
																									"description": `Name or number of the port to access on the container.
Number must be in the range 1 to 65535.
Name must be an IANA_SVC_NAME.`,
																									"x-kubernetes-int-or-string": true,
																								},
																								"scheme": map[string]interface{}{
																									"description": `Scheme to use for connecting to the host.
Defaults to HTTP.`,
																									"type": "string",
																								},
																							},
																							"required": []interface{}{
																								"port",
																							},
																							"type": "object",
																						},
																						"initialDelaySeconds": map[string]interface{}{
																							"description": `Number of seconds after the container has started before liveness probes are initiated.
More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes`,
																							"format": "int32",
																							"type":   "integer",
																						},
																						"periodSeconds": map[string]interface{}{
																							"description": `How often (in seconds) to perform the probe.
Default to 10 seconds. Minimum value is 1.`,
																							"format": "int32",
																							"type":   "integer",
																						},
																						"successThreshold": map[string]interface{}{
																							"description": `Minimum consecutive successes for the probe to be considered successful after having failed.
Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.`,
																							"format": "int32",
																							"type":   "integer",
																						},
																						"tcpSocket": map[string]interface{}{
																							"description": "TCPSocket specifies a connection to a TCP port.",
																							"properties": map[string]interface{}{
																								"host": map[string]interface{}{
																									"description": "Optional: Host name to connect to, defaults to the pod IP.",
																									"type":        "string",
																								},
																								"port": map[string]interface{}{
																									"anyOf": []interface{}{
																										map[string]interface{}{
																											"type": "integer",
																										},
																										map[string]interface{}{
																											"type": "string",
																										},
																									},
																									"description": `Number or name of the port to access on the container.
Number must be in the range 1 to 65535.
Name must be an IANA_SVC_NAME.`,
																									"x-kubernetes-int-or-string": true,
																								},
																							},
																							"required": []interface{}{
																								"port",
																							},
																							"type": "object",
																						},
																						"terminationGracePeriodSeconds": map[string]interface{}{
																							"description": `Optional duration in seconds the pod needs to terminate gracefully upon probe failure.
The grace period is the duration in seconds after the processes running in the pod are sent
a termination signal and the time when the processes are forcibly halted with a kill signal.
Set this value longer than the expected cleanup time for your process.
If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this
value overrides the value provided by the pod spec.
Value must be non-negative integer. The value zero indicates stop immediately via
the kill signal (no opportunity to shut down).
This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate.
Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.`,
																							"format": "int64",
																							"type":   "integer",
																						},
																						"timeoutSeconds": map[string]interface{}{
																							"description": `Number of seconds after which the probe times out.
Defaults to 1 second. Minimum value is 1.
More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes`,
																							"format": "int32",
																							"type":   "integer",
																						},
																					},
																					"type": "object",
																				},
																				"stdin": map[string]interface{}{
																					"description": `Whether this container should allocate a buffer for stdin in the container runtime. If this
is not set, reads from stdin in the container will always result in EOF.
Default is false.`,
																					"type": "boolean",
																				},
																				"stdinOnce": map[string]interface{}{
																					"description": `Whether the container runtime should close the stdin channel after it has been opened by
a single attach. When stdin is true the stdin stream will remain open across multiple attach
sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the
first client attaches to stdin, and then remains open and accepts data until the client disconnects,
at which time stdin is closed and remains closed until the container is restarted. If this
flag is false, a container processes that reads from stdin will never receive an EOF.
Default is false`,
																					"type": "boolean",
																				},
																				"terminationMessagePath": map[string]interface{}{
																					"description": `Optional: Path at which the file to which the container's termination message
will be written is mounted into the container's filesystem.
Message written is intended to be brief final status, such as an assertion failure message.
Will be truncated by the node if greater than 4096 bytes. The total message length across
all containers will be limited to 12kb.
Defaults to /dev/termination-log.
Cannot be updated.`,
																					"type": "string",
																				},
																				"terminationMessagePolicy": map[string]interface{}{
																					"description": `Indicate how the termination message should be populated. File will use the contents of
terminationMessagePath to populate the container status message on both success and failure.
FallbackToLogsOnError will use the last chunk of container log output if the termination
message file is empty and the container exited with an error.
The log output is limited to 2048 bytes or 80 lines, whichever is smaller.
Defaults to File.
Cannot be updated.`,
																					"type": "string",
																				},
																				"tty": map[string]interface{}{
																					"description": `Whether this container should allocate a TTY for itself, also requires 'stdin' to be true.
Default is false.`,
																					"type": "boolean",
																				},
																				"volumeDevices": map[string]interface{}{
																					"description": "volumeDevices is the list of block devices to be used by the container.",
																					"items": map[string]interface{}{
																						"description": "volumeDevice describes a mapping of a raw block device within a container.",
																						"properties": map[string]interface{}{
																							"devicePath": map[string]interface{}{
																								"description": "devicePath is the path inside of the container that the device will be mapped to.",
																								"type":        "string",
																							},
																							"name": map[string]interface{}{
																								"description": "name must match the name of a persistentVolumeClaim in the pod",
																								"type":        "string",
																							},
																						},
																						"required": []interface{}{
																							"devicePath",
																							"name",
																						},
																						"type": "object",
																					},
																					"type": "array",
																					"x-kubernetes-list-map-keys": []interface{}{
																						"devicePath",
																					},
																					"x-kubernetes-list-type": "map",
																				},
																				"volumeMounts": map[string]interface{}{
																					"description": `Pod volumes to mount into the container's filesystem.
Cannot be updated.`,
																					"items": map[string]interface{}{
																						"description": "VolumeMount describes a mounting of a Volume within a container.",
																						"properties": map[string]interface{}{
																							"mountPath": map[string]interface{}{
																								"description": `Path within the container at which the volume should be mounted.  Must
not contain ':'.`,
																								"type": "string",
																							},
																							"mountPropagation": map[string]interface{}{
																								"description": `mountPropagation determines how mounts are propagated from the host
to container and the other way around.
When not set, MountPropagationNone is used.
This field is beta in 1.10.
When RecursiveReadOnly is set to IfPossible or to Enabled, MountPropagation must be None or unspecified
(which defaults to None).`,
																								"type": "string",
																							},
																							"name": map[string]interface{}{
																								"description": "This must match the Name of a Volume.",
																								"type":        "string",
																							},
																							"readOnly": map[string]interface{}{
																								"description": `Mounted read-only if true, read-write otherwise (false or unspecified).
Defaults to false.`,
																								"type": "boolean",
																							},
																							"recursiveReadOnly": map[string]interface{}{
																								"description": `RecursiveReadOnly specifies whether read-only mounts should be handled
recursively.

If ReadOnly is false, this field has no meaning and must be unspecified.

If ReadOnly is true, and this field is set to Disabled, the mount is not made
recursively read-only.  If this field is set to IfPossible, the mount is made
recursively read-only, if it is supported by the container runtime.  If this
field is set to Enabled, the mount is made recursively read-only if it is
supported by the container runtime, otherwise the pod will not be started and
an error will be generated to indicate the reason.

If this field is set to IfPossible or Enabled, MountPropagation must be set to
None (or be unspecified, which defaults to None).

If this field is not specified, it is treated as an equivalent of Disabled.`,
																								"type": "string",
																							},
																							"subPath": map[string]interface{}{
																								"description": `Path within the volume from which the container's volume should be mounted.
Defaults to "" (volume's root).`,
																								"type": "string",
																							},
																							"subPathExpr": map[string]interface{}{
																								"description": `Expanded path within the volume from which the container's volume should be mounted.
Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment.
Defaults to "" (volume's root).
SubPathExpr and SubPath are mutually exclusive.`,
																								"type": "string",
																							},
																						},
																						"required": []interface{}{
																							"mountPath",
																							"name",
																						},
																						"type": "object",
																					},
																					"type": "array",
																					"x-kubernetes-list-map-keys": []interface{}{
																						"mountPath",
																					},
																					"x-kubernetes-list-type": "map",
																				},
																				"workingDir": map[string]interface{}{
																					"description": `Container's working directory.
If not specified, the container runtime's default will be used, which
might be configured in the container image.
Cannot be updated.`,
																					"type": "string",
																				},
																			},
																			"required": []interface{}{
																				"name",
																			},
																			"type": "object",
																		},
																		"type": "array",
																		"x-kubernetes-list-map-keys": []interface{}{
																			"name",
																		},
																		"x-kubernetes-list-type": "map",
																	},
																	"dnsConfig": map[string]interface{}{
																		"description": `Specifies the DNS parameters of a pod.
Parameters specified here will be merged to the generated DNS
configuration based on DNSPolicy.`,
																		"properties": map[string]interface{}{
																			"nameservers": map[string]interface{}{
																				"description": `A list of DNS name server IP addresses.
This will be appended to the base nameservers generated from DNSPolicy.
Duplicated nameservers will be removed.`,
																				"items": map[string]interface{}{
																					"type": "string",
																				},
																				"type":                   "array",
																				"x-kubernetes-list-type": "atomic",
																			},
																			"options": map[string]interface{}{
																				"description": `A list of DNS resolver options.
This will be merged with the base options generated from DNSPolicy.
Duplicated entries will be removed. Resolution options given in Options
will override those that appear in the base DNSPolicy.`,
																				"items": map[string]interface{}{
																					"description": "PodDNSConfigOption defines DNS resolver options of a pod.",
																					"properties": map[string]interface{}{
																						"name": map[string]interface{}{
																							"description": `Name is this DNS resolver option's name.
Required.`,
																							"type": "string",
																						},
																						"value": map[string]interface{}{
																							"description": "Value is this DNS resolver option's value.",
																							"type":        "string",
																						},
																					},
																					"type": "object",
																				},
																				"type":                   "array",
																				"x-kubernetes-list-type": "atomic",
																			},
																			"searches": map[string]interface{}{
																				"description": `A list of DNS search domains for host-name lookup.
This will be appended to the base search paths generated from DNSPolicy.
Duplicated search paths will be removed.`,
																				"items": map[string]interface{}{
																					"type": "string",
																				},
																				"type":                   "array",
																				"x-kubernetes-list-type": "atomic",
																			},
																		},
																		"type": "object",
																	},
																	"dnsPolicy": map[string]interface{}{
																		"description": `Set DNS policy for the pod.
Defaults to "ClusterFirst".
Valid values are 'ClusterFirstWithHostNet', 'ClusterFirst', 'Default' or 'None'.
DNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy.
To have DNS options set along with hostNetwork, you have to specify DNS policy
explicitly to 'ClusterFirstWithHostNet'.`,
																		"type": "string",
																	},
																	"enableServiceLinks": map[string]interface{}{
																		"description": `EnableServiceLinks indicates whether information about services should be injected into pod's
environment variables, matching the syntax of Docker links.
Optional: Defaults to true.`,
																		"type": "boolean",
																	},
																	"ephemeralContainers": map[string]interface{}{
																		"description": `List of ephemeral containers run in this pod. Ephemeral containers may be run in an existing
pod to perform user-initiated actions such as debugging. This list cannot be specified when
creating a pod, and it cannot be modified by updating the pod spec. In order to add an
ephemeral container to an existing pod, use the pod's ephemeralcontainers subresource.`,
																		"items": map[string]interface{}{
																			"description": `An EphemeralContainer is a temporary container that you may add to an existing Pod for
user-initiated activities such as debugging. Ephemeral containers have no resource or
scheduling guarantees, and they will not be restarted when they exit or when a Pod is
removed or restarted. The kubelet may evict a Pod if an ephemeral container causes the
Pod to exceed its resource allocation.

To add an ephemeral container, use the ephemeralcontainers subresource of an existing
Pod. Ephemeral containers may not be removed or restarted.`,
																			"properties": map[string]interface{}{
																				"args": map[string]interface{}{
																					"description": `Arguments to the entrypoint.
The image's CMD is used if this is not provided.
Variable references $(VAR_NAME) are expanded using the container's environment. If a variable
cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced
to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. "$$(VAR_NAME)" will
produce the string literal "$(VAR_NAME)". Escaped references will never be expanded, regardless
of whether the variable exists or not. Cannot be updated.
More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell`,
																					"items": map[string]interface{}{
																						"type": "string",
																					},
																					"type":                   "array",
																					"x-kubernetes-list-type": "atomic",
																				},
																				"command": map[string]interface{}{
																					"description": `Entrypoint array. Not executed within a shell.
The image's ENTRYPOINT is used if this is not provided.
Variable references $(VAR_NAME) are expanded using the container's environment. If a variable
cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced
to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. "$$(VAR_NAME)" will
produce the string literal "$(VAR_NAME)". Escaped references will never be expanded, regardless
of whether the variable exists or not. Cannot be updated.
More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell`,
																					"items": map[string]interface{}{
																						"type": "string",
																					},
																					"type":                   "array",
																					"x-kubernetes-list-type": "atomic",
																				},
																				"env": map[string]interface{}{
																					"description": `List of environment variables to set in the container.
Cannot be updated.`,
																					"items": map[string]interface{}{
																						"description": "EnvVar represents an environment variable present in a Container.",
																						"properties": map[string]interface{}{
																							"name": map[string]interface{}{
																								"description": `Name of the environment variable.
May consist of any printable ASCII characters except '='.`,
																								"type": "string",
																							},
																							"value": map[string]interface{}{
																								"description": `Variable references $(VAR_NAME) are expanded
using the previously defined environment variables in the container and
any service environment variables. If a variable cannot be resolved,
the reference in the input string will be unchanged. Double $$ are reduced
to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e.
"$$(VAR_NAME)" will produce the string literal "$(VAR_NAME)".
Escaped references will never be expanded, regardless of whether the variable
exists or not.
Defaults to "".`,
																								"type": "string",
																							},
																							"valueFrom": map[string]interface{}{
																								"description": "Source for the environment variable's value. Cannot be used if value is not empty.",
																								"properties": map[string]interface{}{
																									"configMapKeyRef": map[string]interface{}{
																										"description": "Selects a key of a ConfigMap.",
																										"properties": map[string]interface{}{
																											"key": map[string]interface{}{
																												"description": "The key to select.",
																												"type":        "string",
																											},
																											"name": map[string]interface{}{
																												"default": "",
																												"description": `Name of the referent.
This field is effectively required, but due to backwards compatibility is
allowed to be empty. Instances of this type with an empty value here are
almost certainly wrong.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names`,
																												"type": "string",
																											},
																											"optional": map[string]interface{}{
																												"description": "Specify whether the ConfigMap or its key must be defined",
																												"type":        "boolean",
																											},
																										},
																										"required": []interface{}{
																											"key",
																										},
																										"type":                  "object",
																										"x-kubernetes-map-type": "atomic",
																									},
																									"fieldRef": map[string]interface{}{
																										"description": `Selects a field of the pod: supports metadata.name, metadata.namespace, ` + "`" + `metadata.labels['<KEY>']` + "`" + `, ` + "`" + `metadata.annotations['<KEY>']` + "`" + `,
spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.`,
																										"properties": map[string]interface{}{
																											"apiVersion": map[string]interface{}{
																												"description": "Version of the schema the FieldPath is written in terms of, defaults to \"v1\".",
																												"type":        "string",
																											},
																											"fieldPath": map[string]interface{}{
																												"description": "Path of the field to select in the specified API version.",
																												"type":        "string",
																											},
																										},
																										"required": []interface{}{
																											"fieldPath",
																										},
																										"type":                  "object",
																										"x-kubernetes-map-type": "atomic",
																									},
																									"fileKeyRef": map[string]interface{}{
																										"description": `FileKeyRef selects a key of the env file.
Requires the EnvFiles feature gate to be enabled.`,
																										"properties": map[string]interface{}{
																											"key": map[string]interface{}{
																												"description": `The key within the env file. An invalid key will prevent the pod from starting.
The keys defined within a source may consist of any printable ASCII characters except '='.
During Alpha stage of the EnvFiles feature gate, the key size is limited to 128 characters.`,
																												"type": "string",
																											},
																											"optional": map[string]interface{}{
																												"default": false,
																												"description": `Specify whether the file or its key must be defined. If the file or key
does not exist, then the env var is not published.
If optional is set to true and the specified key does not exist,
the environment variable will not be set in the Pod's containers.

If optional is set to false and the specified key does not exist,
an error will be returned during Pod creation.`,
																												"type": "boolean",
																											},
																											"path": map[string]interface{}{
																												"description": `The path within the volume from which to select the file.
Must be relative and may not contain the '..' path or start with '..'.`,
																												"type": "string",
																											},
																											"volumeName": map[string]interface{}{
																												"description": "The name of the volume mount containing the env file.",
																												"type":        "string",
																											},
																										},
																										"required": []interface{}{
																											"key",
																											"path",
																											"volumeName",
																										},
																										"type":                  "object",
																										"x-kubernetes-map-type": "atomic",
																									},
																									"resourceFieldRef": map[string]interface{}{
																										"description": `Selects a resource of the container: only resources limits and requests
(limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.`,
																										"properties": map[string]interface{}{
																											"containerName": map[string]interface{}{
																												"description": "Container name: required for volumes, optional for env vars",
																												"type":        "string",
																											},
																											"divisor": map[string]interface{}{
																												"anyOf": []interface{}{
																													map[string]interface{}{
																														"type": "integer",
																													},
																													map[string]interface{}{
																														"type": "string",
																													},
																												},
																												"description":                "Specifies the output format of the exposed resources, defaults to \"1\"",
																												"pattern":                    `^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$`,
																												"x-kubernetes-int-or-string": true,
																											},
																											"resource": map[string]interface{}{
																												"description": "Required: resource to select",
																												"type":        "string",
																											},
																										},
																										"required": []interface{}{
																											"resource",
																										},
																										"type":                  "object",
																										"x-kubernetes-map-type": "atomic",
																									},
																									"secretKeyRef": map[string]interface{}{
																										"description": "Selects a key of a secret in the pod's namespace",
																										"properties": map[string]interface{}{
																											"key": map[string]interface{}{
																												"description": "The key of the secret to select from.  Must be a valid secret key.",
																												"type":        "string",
																											},
																											"name": map[string]interface{}{
																												"default": "",
																												"description": `Name of the referent.
This field is effectively required, but due to backwards compatibility is
allowed to be empty. Instances of this type with an empty value here are
almost certainly wrong.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names`,
																												"type": "string",
																											},
																											"optional": map[string]interface{}{
																												"description": "Specify whether the Secret or its key must be defined",
																												"type":        "boolean",
																											},
																										},
																										"required": []interface{}{
																											"key",
																										},
																										"type":                  "object",
																										"x-kubernetes-map-type": "atomic",
																									},
																								},
																								"type": "object",
																							},
																						},
																						"required": []interface{}{
																							"name",
																						},
																						"type": "object",
																					},
																					"type": "array",
																					"x-kubernetes-list-map-keys": []interface{}{
																						"name",
																					},
																					"x-kubernetes-list-type": "map",
																				},
																				"envFrom": map[string]interface{}{
																					"description": `List of sources to populate environment variables in the container.
The keys defined within a source may consist of any printable ASCII characters except '='.
When a key exists in multiple
sources, the value associated with the last source will take precedence.
Values defined by an Env with a duplicate key will take precedence.
Cannot be updated.`,
																					"items": map[string]interface{}{
																						"description": "EnvFromSource represents the source of a set of ConfigMaps or Secrets",
																						"properties": map[string]interface{}{
																							"configMapRef": map[string]interface{}{
																								"description": "The ConfigMap to select from",
																								"properties": map[string]interface{}{
																									"name": map[string]interface{}{
																										"default": "",
																										"description": `Name of the referent.
This field is effectively required, but due to backwards compatibility is
allowed to be empty. Instances of this type with an empty value here are
almost certainly wrong.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names`,
																										"type": "string",
																									},
																									"optional": map[string]interface{}{
																										"description": "Specify whether the ConfigMap must be defined",
																										"type":        "boolean",
																									},
																								},
																								"type":                  "object",
																								"x-kubernetes-map-type": "atomic",
																							},
																							"prefix": map[string]interface{}{
																								"description": `Optional text to prepend to the name of each environment variable.
May consist of any printable ASCII characters except '='.`,
																								"type": "string",
																							},
																							"secretRef": map[string]interface{}{
																								"description": "The Secret to select from",
																								"properties": map[string]interface{}{
																									"name": map[string]interface{}{
																										"default": "",
																										"description": `Name of the referent.
This field is effectively required, but due to backwards compatibility is
allowed to be empty. Instances of this type with an empty value here are
almost certainly wrong.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names`,
																										"type": "string",
																									},
																									"optional": map[string]interface{}{
																										"description": "Specify whether the Secret must be defined",
																										"type":        "boolean",
																									},
																								},
																								"type":                  "object",
																								"x-kubernetes-map-type": "atomic",
																							},
																						},
																						"type": "object",
																					},
																					"type":                   "array",
																					"x-kubernetes-list-type": "atomic",
																				},
																				"image": map[string]interface{}{
																					"description": `Container image name.
More info: https://kubernetes.io/docs/concepts/containers/images`,
																					"type": "string",
																				},
																				"imagePullPolicy": map[string]interface{}{
																					"description": `Image pull policy.
One of Always, Never, IfNotPresent.
Defaults to Always if :latest tag is specified, or IfNotPresent otherwise.
Cannot be updated.
More info: https://kubernetes.io/docs/concepts/containers/images#updating-images`,
																					"type": "string",
																				},
																				"lifecycle": map[string]interface{}{
																					"description": "Lifecycle is not allowed for ephemeral containers.",
																					"properties": map[string]interface{}{
																						"postStart": map[string]interface{}{
																							"description": `PostStart is called immediately after a container is created. If the handler fails,
the container is terminated and restarted according to its restart policy.
Other management of the container blocks until the hook completes.
More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks`,
																							"properties": map[string]interface{}{
																								"exec": map[string]interface{}{
																									"description": "Exec specifies a command to execute in the container.",
																									"properties": map[string]interface{}{
																										"command": map[string]interface{}{
																											"description": `Command is the command line to execute inside the container, the working directory for the
command  is root ('/') in the container's filesystem. The command is simply exec'd, it is
not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use
a shell, you need to explicitly call out to that shell.
Exit status of 0 is treated as live/healthy and non-zero is unhealthy.`,
																											"items": map[string]interface{}{
																												"type": "string",
																											},
																											"type":                   "array",
																											"x-kubernetes-list-type": "atomic",
																										},
																									},
																									"type": "object",
																								},
																								"httpGet": map[string]interface{}{
																									"description": "HTTPGet specifies an HTTP GET request to perform.",
																									"properties": map[string]interface{}{
																										"host": map[string]interface{}{
																											"description": `Host name to connect to, defaults to the pod IP. You probably want to set
"Host" in httpHeaders instead.`,
																											"type": "string",
																										},
																										"httpHeaders": map[string]interface{}{
																											"description": "Custom headers to set in the request. HTTP allows repeated headers.",
																											"items": map[string]interface{}{
																												"description": "HTTPHeader describes a custom header to be used in HTTP probes",
																												"properties": map[string]interface{}{
																													"name": map[string]interface{}{
																														"description": `The header field name.
This will be canonicalized upon output, so case-variant names will be understood as the same header.`,
																														"type": "string",
																													},
																													"value": map[string]interface{}{
																														"description": "The header field value",
																														"type":        "string",
																													},
																												},
																												"required": []interface{}{
																													"name",
																													"value",
																												},
																												"type": "object",
																											},
																											"type":                   "array",
																											"x-kubernetes-list-type": "atomic",
																										},
																										"path": map[string]interface{}{
																											"description": "Path to access on the HTTP server.",
																											"type":        "string",
																										},
																										"port": map[string]interface{}{
																											"anyOf": []interface{}{
																												map[string]interface{}{
																													"type": "integer",
																												},
																												map[string]interface{}{
																													"type": "string",
																												},
																											},
																											"description": `Name or number of the port to access on the container.
Number must be in the range 1 to 65535.
Name must be an IANA_SVC_NAME.`,
																											"x-kubernetes-int-or-string": true,
																										},
																										"scheme": map[string]interface{}{
																											"description": `Scheme to use for connecting to the host.
Defaults to HTTP.`,
																											"type": "string",
																										},
																									},
																									"required": []interface{}{
																										"port",
																									},
																									"type": "object",
																								},
																								"sleep": map[string]interface{}{
																									"description": "Sleep represents a duration that the container should sleep.",
																									"properties": map[string]interface{}{
																										"seconds": map[string]interface{}{
																											"description": "Seconds is the number of seconds to sleep.",
																											"format":      "int64",
																											"type":        "integer",
																										},
																									},
																									"required": []interface{}{
																										"seconds",
																									},
																									"type": "object",
																								},
																								"tcpSocket": map[string]interface{}{
																									"description": `Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept
for backward compatibility. There is no validation of this field and
lifecycle hooks will fail at runtime when it is specified.`,
																									"properties": map[string]interface{}{
																										"host": map[string]interface{}{
																											"description": "Optional: Host name to connect to, defaults to the pod IP.",
																											"type":        "string",
																										},
																										"port": map[string]interface{}{
																											"anyOf": []interface{}{
																												map[string]interface{}{
																													"type": "integer",
																												},
																												map[string]interface{}{
																													"type": "string",
																												},
																											},
																											"description": `Number or name of the port to access on the container.
Number must be in the range 1 to 65535.
Name must be an IANA_SVC_NAME.`,
																											"x-kubernetes-int-or-string": true,
																										},
																									},
																									"required": []interface{}{
																										"port",
																									},
																									"type": "object",
																								},
																							},
																							"type": "object",
																						},
																						"preStop": map[string]interface{}{
																							"description": `PreStop is called immediately before a container is terminated due to an
API request or management event such as liveness/startup probe failure,
preemption, resource contention, etc. The handler is not called if the
container crashes or exits. The Pod's termination grace period countdown begins before the
PreStop hook is executed. Regardless of the outcome of the handler, the
container will eventually terminate within the Pod's termination grace
period (unless delayed by finalizers). Other management of the container blocks until the hook completes
or until the termination grace period is reached.
More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks`,
																							"properties": map[string]interface{}{
																								"exec": map[string]interface{}{
																									"description": "Exec specifies a command to execute in the container.",
																									"properties": map[string]interface{}{
																										"command": map[string]interface{}{
																											"description": `Command is the command line to execute inside the container, the working directory for the
command  is root ('/') in the container's filesystem. The command is simply exec'd, it is
not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use
a shell, you need to explicitly call out to that shell.
Exit status of 0 is treated as live/healthy and non-zero is unhealthy.`,
																											"items": map[string]interface{}{
																												"type": "string",
																											},
																											"type":                   "array",
																											"x-kubernetes-list-type": "atomic",
																										},
																									},
																									"type": "object",
																								},
																								"httpGet": map[string]interface{}{
																									"description": "HTTPGet specifies an HTTP GET request to perform.",
																									"properties": map[string]interface{}{
																										"host": map[string]interface{}{
																											"description": `Host name to connect to, defaults to the pod IP. You probably want to set
"Host" in httpHeaders instead.`,
																											"type": "string",
																										},
																										"httpHeaders": map[string]interface{}{
																											"description": "Custom headers to set in the request. HTTP allows repeated headers.",
																											"items": map[string]interface{}{
																												"description": "HTTPHeader describes a custom header to be used in HTTP probes",
																												"properties": map[string]interface{}{
																													"name": map[string]interface{}{
																														"description": `The header field name.
This will be canonicalized upon output, so case-variant names will be understood as the same header.`,
																														"type": "string",
																													},
																													"value": map[string]interface{}{
																														"description": "The header field value",
																														"type":        "string",
																													},
																												},
																												"required": []interface{}{
																													"name",
																													"value",
																												},
																												"type": "object",
																											},
																											"type":                   "array",
																											"x-kubernetes-list-type": "atomic",
																										},
																										"path": map[string]interface{}{
																											"description": "Path to access on the HTTP server.",
																											"type":        "string",
																										},
																										"port": map[string]interface{}{
																											"anyOf": []interface{}{
																												map[string]interface{}{
																													"type": "integer",
																												},
																												map[string]interface{}{
																													"type": "string",
																												},
																											},
																											"description": `Name or number of the port to access on the container.
Number must be in the range 1 to 65535.
Name must be an IANA_SVC_NAME.`,
																											"x-kubernetes-int-or-string": true,
																										},
																										"scheme": map[string]interface{}{
																											"description": `Scheme to use for connecting to the host.
Defaults to HTTP.`,
																											"type": "string",
																										},
																									},
																									"required": []interface{}{
																										"port",
																									},
																									"type": "object",
																								},
																								"sleep": map[string]interface{}{
																									"description": "Sleep represents a duration that the container should sleep.",
																									"properties": map[string]interface{}{
																										"seconds": map[string]interface{}{
																											"description": "Seconds is the number of seconds to sleep.",
																											"format":      "int64",
																											"type":        "integer",
																										},
																									},
																									"required": []interface{}{
																										"seconds",
																									},
																									"type": "object",
																								},
																								"tcpSocket": map[string]interface{}{
																									"description": `Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept
for backward compatibility. There is no validation of this field and
lifecycle hooks will fail at runtime when it is specified.`,
																									"properties": map[string]interface{}{
																										"host": map[string]interface{}{
																											"description": "Optional: Host name to connect to, defaults to the pod IP.",
																											"type":        "string",
																										},
																										"port": map[string]interface{}{
																											"anyOf": []interface{}{
																												map[string]interface{}{
																													"type": "integer",
																												},
																												map[string]interface{}{
																													"type": "string",
																												},
																											},
																											"description": `Number or name of the port to access on the container.
Number must be in the range 1 to 65535.
Name must be an IANA_SVC_NAME.`,
																											"x-kubernetes-int-or-string": true,
																										},
																									},
																									"required": []interface{}{
																										"port",
																									},
																									"type": "object",
																								},
																							},
																							"type": "object",
																						},
																						"stopSignal": map[string]interface{}{
																							"description": `StopSignal defines which signal will be sent to a container when it is being stopped.
If not specified, the default is defined by the container runtime in use.
StopSignal can only be set for Pods with a non-empty .spec.os.name`,
																							"type": "string",
																						},
																					},
																					"type": "object",
																				},
																				"livenessProbe": map[string]interface{}{
																					"description": "Probes are not allowed for ephemeral containers.",
																					"properties": map[string]interface{}{
																						"exec": map[string]interface{}{
																							"description": "Exec specifies a command to execute in the container.",
																							"properties": map[string]interface{}{
																								"command": map[string]interface{}{
																									"description": `Command is the command line to execute inside the container, the working directory for the
command  is root ('/') in the container's filesystem. The command is simply exec'd, it is
not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use
a shell, you need to explicitly call out to that shell.
Exit status of 0 is treated as live/healthy and non-zero is unhealthy.`,
																									"items": map[string]interface{}{
																										"type": "string",
																									},
																									"type":                   "array",
																									"x-kubernetes-list-type": "atomic",
																								},
																							},
																							"type": "object",
																						},
																						"failureThreshold": map[string]interface{}{
																							"description": `Minimum consecutive failures for the probe to be considered failed after having succeeded.
Defaults to 3. Minimum value is 1.`,
																							"format": "int32",
																							"type":   "integer",
																						},
																						"grpc": map[string]interface{}{
																							"description": "GRPC specifies a GRPC HealthCheckRequest.",
																							"properties": map[string]interface{}{
																								"port": map[string]interface{}{
																									"description": "Port number of the gRPC service. Number must be in the range 1 to 65535.",
																									"format":      "int32",
																									"type":        "integer",
																								},
																								"service": map[string]interface{}{
																									"default": "",
																									"description": `Service is the name of the service to place in the gRPC HealthCheckRequest
(see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).

If this is not specified, the default behavior is defined by gRPC.`,
																									"type": "string",
																								},
																							},
																							"required": []interface{}{
																								"port",
																							},
																							"type": "object",
																						},
																						"httpGet": map[string]interface{}{
																							"description": "HTTPGet specifies an HTTP GET request to perform.",
																							"properties": map[string]interface{}{
																								"host": map[string]interface{}{
																									"description": `Host name to connect to, defaults to the pod IP. You probably want to set
"Host" in httpHeaders instead.`,
																									"type": "string",
																								},
																								"httpHeaders": map[string]interface{}{
																									"description": "Custom headers to set in the request. HTTP allows repeated headers.",
																									"items": map[string]interface{}{
																										"description": "HTTPHeader describes a custom header to be used in HTTP probes",
																										"properties": map[string]interface{}{
																											"name": map[string]interface{}{
																												"description": `The header field name.
This will be canonicalized upon output, so case-variant names will be understood as the same header.`,
																												"type": "string",
																											},
																											"value": map[string]interface{}{
																												"description": "The header field value",
																												"type":        "string",
																											},
																										},
																										"required": []interface{}{
																											"name",
																											"value",
																										},
																										"type": "object",
																									},
																									"type":                   "array",
																									"x-kubernetes-list-type": "atomic",
																								},
																								"path": map[string]interface{}{
																									"description": "Path to access on the HTTP server.",
																									"type":        "string",
																								},
																								"port": map[string]interface{}{
																									"anyOf": []interface{}{
																										map[string]interface{}{
																											"type": "integer",
																										},
																										map[string]interface{}{
																											"type": "string",
																										},
																									},
																									"description": `Name or number of the port to access on the container.
Number must be in the range 1 to 65535.
Name must be an IANA_SVC_NAME.`,
																									"x-kubernetes-int-or-string": true,
																								},
																								"scheme": map[string]interface{}{
																									"description": `Scheme to use for connecting to the host.
Defaults to HTTP.`,
																									"type": "string",
																								},
																							},
																							"required": []interface{}{
																								"port",
																							},
																							"type": "object",
																						},
																						"initialDelaySeconds": map[string]interface{}{
																							"description": `Number of seconds after the container has started before liveness probes are initiated.
More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes`,
																							"format": "int32",
																							"type":   "integer",
																						},
																						"periodSeconds": map[string]interface{}{
																							"description": `How often (in seconds) to perform the probe.
Default to 10 seconds. Minimum value is 1.`,
																							"format": "int32",
																							"type":   "integer",
																						},
																						"successThreshold": map[string]interface{}{
																							"description": `Minimum consecutive successes for the probe to be considered successful after having failed.
Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.`,
																							"format": "int32",
																							"type":   "integer",
																						},
																						"tcpSocket": map[string]interface{}{
																							"description": "TCPSocket specifies a connection to a TCP port.",
																							"properties": map[string]interface{}{
																								"host": map[string]interface{}{
																									"description": "Optional: Host name to connect to, defaults to the pod IP.",
																									"type":        "string",
																								},
																								"port": map[string]interface{}{
																									"anyOf": []interface{}{
																										map[string]interface{}{
																											"type": "integer",
																										},
																										map[string]interface{}{
																											"type": "string",
																										},
																									},
																									"description": `Number or name of the port to access on the container.
Number must be in the range 1 to 65535.
Name must be an IANA_SVC_NAME.`,
																									"x-kubernetes-int-or-string": true,
																								},
																							},
																							"required": []interface{}{
																								"port",
																							},
																							"type": "object",
																						},
																						"terminationGracePeriodSeconds": map[string]interface{}{
																							"description": `Optional duration in seconds the pod needs to terminate gracefully upon probe failure.
The grace period is the duration in seconds after the processes running in the pod are sent
a termination signal and the time when the processes are forcibly halted with a kill signal.
Set this value longer than the expected cleanup time for your process.
If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this
value overrides the value provided by the pod spec.
Value must be non-negative integer. The value zero indicates stop immediately via
the kill signal (no opportunity to shut down).
This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate.
Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.`,
																							"format": "int64",
																							"type":   "integer",
																						},
																						"timeoutSeconds": map[string]interface{}{
																							"description": `Number of seconds after which the probe times out.
Defaults to 1 second. Minimum value is 1.
More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes`,
																							"format": "int32",
																							"type":   "integer",
																						},
																					},
																					"type": "object",
																				},
																				"name": map[string]interface{}{
																					"description": `Name of the ephemeral container specified as a DNS_LABEL.
This name must be unique among all containers, init containers and ephemeral containers.`,
																					"type": "string",
																				},
																				"ports": map[string]interface{}{
																					"description": "Ports are not allowed for ephemeral containers.",
																					"items": map[string]interface{}{
																						"description": "ContainerPort represents a network port in a single container.",
																						"properties": map[string]interface{}{
																							"containerPort": map[string]interface{}{
																								"description": `Number of port to expose on the pod's IP address.
This must be a valid port number, 0 < x < 65536.`,
																								"format": "int32",
																								"type":   "integer",
																							},
																							"hostIP": map[string]interface{}{
																								"description": "What host IP to bind the external port to.",
																								"type":        "string",
																							},
																							"hostPort": map[string]interface{}{
																								"description": `Number of port to expose on the host.
If specified, this must be a valid port number, 0 < x < 65536.
If HostNetwork is specified, this must match ContainerPort.
Most containers do not need this.`,
																								"format": "int32",
																								"type":   "integer",
																							},
																							"name": map[string]interface{}{
																								"description": `If specified, this must be an IANA_SVC_NAME and unique within the pod. Each
named port in a pod must have a unique name. Name for the port that can be
referred to by services.`,
																								"type": "string",
																							},
																							"protocol": map[string]interface{}{
																								"default": "TCP",
																								"description": `Protocol for port. Must be UDP, TCP, or SCTP.
Defaults to "TCP".`,
																								"type": "string",
																							},
																						},
																						"required": []interface{}{
																							"containerPort",
																							"protocol",
																						},
																						"type": "object",
																					},
																					"type": "array",
																					"x-kubernetes-list-map-keys": []interface{}{
																						"containerPort",
																						"protocol",
																					},
																					"x-kubernetes-list-type": "map",
																				},
																				"readinessProbe": map[string]interface{}{
																					"description": "Probes are not allowed for ephemeral containers.",
																					"properties": map[string]interface{}{
																						"exec": map[string]interface{}{
																							"description": "Exec specifies a command to execute in the container.",
																							"properties": map[string]interface{}{
																								"command": map[string]interface{}{
																									"description": `Command is the command line to execute inside the container, the working directory for the
command  is root ('/') in the container's filesystem. The command is simply exec'd, it is
not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use
a shell, you need to explicitly call out to that shell.
Exit status of 0 is treated as live/healthy and non-zero is unhealthy.`,
																									"items": map[string]interface{}{
																										"type": "string",
																									},
																									"type":                   "array",
																									"x-kubernetes-list-type": "atomic",
																								},
																							},
																							"type": "object",
																						},
																						"failureThreshold": map[string]interface{}{
																							"description": `Minimum consecutive failures for the probe to be considered failed after having succeeded.
Defaults to 3. Minimum value is 1.`,
																							"format": "int32",
																							"type":   "integer",
																						},
																						"grpc": map[string]interface{}{
																							"description": "GRPC specifies a GRPC HealthCheckRequest.",
																							"properties": map[string]interface{}{
																								"port": map[string]interface{}{
																									"description": "Port number of the gRPC service. Number must be in the range 1 to 65535.",
																									"format":      "int32",
																									"type":        "integer",
																								},
																								"service": map[string]interface{}{
																									"default": "",
																									"description": `Service is the name of the service to place in the gRPC HealthCheckRequest
(see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).

If this is not specified, the default behavior is defined by gRPC.`,
																									"type": "string",
																								},
																							},
																							"required": []interface{}{
																								"port",
																							},
																							"type": "object",
																						},
																						"httpGet": map[string]interface{}{
																							"description": "HTTPGet specifies an HTTP GET request to perform.",
																							"properties": map[string]interface{}{
																								"host": map[string]interface{}{
																									"description": `Host name to connect to, defaults to the pod IP. You probably want to set
"Host" in httpHeaders instead.`,
																									"type": "string",
																								},
																								"httpHeaders": map[string]interface{}{
																									"description": "Custom headers to set in the request. HTTP allows repeated headers.",
																									"items": map[string]interface{}{
																										"description": "HTTPHeader describes a custom header to be used in HTTP probes",
																										"properties": map[string]interface{}{
																											"name": map[string]interface{}{
																												"description": `The header field name.
This will be canonicalized upon output, so case-variant names will be understood as the same header.`,
																												"type": "string",
																											},
																											"value": map[string]interface{}{
																												"description": "The header field value",
																												"type":        "string",
																											},
																										},
																										"required": []interface{}{
																											"name",
																											"value",
																										},
																										"type": "object",
																									},
																									"type":                   "array",
																									"x-kubernetes-list-type": "atomic",
																								},
																								"path": map[string]interface{}{
																									"description": "Path to access on the HTTP server.",
																									"type":        "string",
																								},
																								"port": map[string]interface{}{
																									"anyOf": []interface{}{
																										map[string]interface{}{
																											"type": "integer",
																										},
																										map[string]interface{}{
																											"type": "string",
																										},
																									},
																									"description": `Name or number of the port to access on the container.
Number must be in the range 1 to 65535.
Name must be an IANA_SVC_NAME.`,
																									"x-kubernetes-int-or-string": true,
																								},
																								"scheme": map[string]interface{}{
																									"description": `Scheme to use for connecting to the host.
Defaults to HTTP.`,
																									"type": "string",
																								},
																							},
																							"required": []interface{}{
																								"port",
																							},
																							"type": "object",
																						},
																						"initialDelaySeconds": map[string]interface{}{
																							"description": `Number of seconds after the container has started before liveness probes are initiated.
More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes`,
																							"format": "int32",
																							"type":   "integer",
																						},
																						"periodSeconds": map[string]interface{}{
																							"description": `How often (in seconds) to perform the probe.
Default to 10 seconds. Minimum value is 1.`,
																							"format": "int32",
																							"type":   "integer",
																						},
																						"successThreshold": map[string]interface{}{
																							"description": `Minimum consecutive successes for the probe to be considered successful after having failed.
Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.`,
																							"format": "int32",
																							"type":   "integer",
																						},
																						"tcpSocket": map[string]interface{}{
																							"description": "TCPSocket specifies a connection to a TCP port.",
																							"properties": map[string]interface{}{
																								"host": map[string]interface{}{
																									"description": "Optional: Host name to connect to, defaults to the pod IP.",
																									"type":        "string",
																								},
																								"port": map[string]interface{}{
																									"anyOf": []interface{}{
																										map[string]interface{}{
																											"type": "integer",
																										},
																										map[string]interface{}{
																											"type": "string",
																										},
																									},
																									"description": `Number or name of the port to access on the container.
Number must be in the range 1 to 65535.
Name must be an IANA_SVC_NAME.`,
																									"x-kubernetes-int-or-string": true,
																								},
																							},
																							"required": []interface{}{
																								"port",
																							},
																							"type": "object",
																						},
																						"terminationGracePeriodSeconds": map[string]interface{}{
																							"description": `Optional duration in seconds the pod needs to terminate gracefully upon probe failure.
The grace period is the duration in seconds after the processes running in the pod are sent
a termination signal and the time when the processes are forcibly halted with a kill signal.
Set this value longer than the expected cleanup time for your process.
If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this
value overrides the value provided by the pod spec.
Value must be non-negative integer. The value zero indicates stop immediately via
the kill signal (no opportunity to shut down).
This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate.
Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.`,
																							"format": "int64",
																							"type":   "integer",
																						},
																						"timeoutSeconds": map[string]interface{}{
																							"description": `Number of seconds after which the probe times out.
Defaults to 1 second. Minimum value is 1.
More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes`,
																							"format": "int32",
																							"type":   "integer",
																						},
																					},
																					"type": "object",
																				},
																				"resizePolicy": map[string]interface{}{
																					"description": "Resources resize policy for the container.",
																					"items": map[string]interface{}{
																						"description": "ContainerResizePolicy represents resource resize policy for the container.",
																						"properties": map[string]interface{}{
																							"resourceName": map[string]interface{}{
																								"description": `Name of the resource to which this resource resize policy applies.
Supported values: cpu, memory.`,
																								"type": "string",
																							},
																							"restartPolicy": map[string]interface{}{
																								"description": `Restart policy to apply when specified resource is resized.
If not specified, it defaults to NotRequired.`,
																								"type": "string",
																							},
																						},
																						"required": []interface{}{
																							"resourceName",
																							"restartPolicy",
																						},
																						"type": "object",
																					},
																					"type":                   "array",
																					"x-kubernetes-list-type": "atomic",
																				},
																				"resources": map[string]interface{}{
																					"description": `Resources are not allowed for ephemeral containers. Ephemeral containers use spare resources
already allocated to the pod.`,
																					"properties": map[string]interface{}{
																						"claims": map[string]interface{}{
																							"description": `Claims lists the names of resources, defined in spec.resourceClaims,
that are used by this container.

This field depends on the
DynamicResourceAllocation feature gate.

This field is immutable. It can only be set for containers.`,
																							"items": map[string]interface{}{
																								"description": "ResourceClaim references one entry in PodSpec.ResourceClaims.",
																								"properties": map[string]interface{}{
																									"name": map[string]interface{}{
																										"description": `Name must match the name of one entry in pod.spec.resourceClaims of
the Pod where this field is used. It makes that resource available
inside a container.`,
																										"type": "string",
																									},
																									"request": map[string]interface{}{
																										"description": `Request is the name chosen for a request in the referenced claim.
If empty, everything from the claim is made available, otherwise
only the result of this request.`,
																										"type": "string",
																									},
																								},
																								"required": []interface{}{
																									"name",
																								},
																								"type": "object",
																							},
																							"type": "array",
																							"x-kubernetes-list-map-keys": []interface{}{
																								"name",
																							},
																							"x-kubernetes-list-type": "map",
																						},
																						"limits": map[string]interface{}{
																							"additionalProperties": map[string]interface{}{
																								"anyOf": []interface{}{
																									map[string]interface{}{
																										"type": "integer",
																									},
																									map[string]interface{}{
																										"type": "string",
																									},
																								},
																								"pattern":                    `^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$`,
																								"x-kubernetes-int-or-string": true,
																							},
																							"description": `Limits describes the maximum amount of compute resources allowed.
More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/`,
																							"type": "object",
																						},
																						"requests": map[string]interface{}{
																							"additionalProperties": map[string]interface{}{
																								"anyOf": []interface{}{
																									map[string]interface{}{
																										"type": "integer",
																									},
																									map[string]interface{}{
																										"type": "string",
																									},
																								},
																								"pattern":                    `^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$`,
																								"x-kubernetes-int-or-string": true,
																							},
																							"description": `Requests describes the minimum amount of compute resources required.
If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,
otherwise to an implementation-defined value. Requests cannot exceed Limits.
More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/`,
																							"type": "object",
																						},
																					},
																					"type": "object",
																				},
																				"restartPolicy": map[string]interface{}{
																					"description": `Restart policy for the container to manage the restart behavior of each
container within a pod.
You cannot set this field on ephemeral containers.`,
																					"type": "string",
																				},
																				"restartPolicyRules": map[string]interface{}{
																					"description": `Represents a list of rules to be checked to determine if the
container should be restarted on exit. You cannot set this field on
ephemeral containers.`,
																					"items": map[string]interface{}{
																						"description": "ContainerRestartRule describes how a container exit is handled.",
																						"properties": map[string]interface{}{
																							"action": map[string]interface{}{
																								"description": `Specifies the action taken on a container exit if the requirements
are satisfied. The only possible value is "Restart" to restart the
container.`,
																								"type": "string",
																							},
																							"exitCodes": map[string]interface{}{
																								"description": "Represents the exit codes to check on container exits.",
																								"properties": map[string]interface{}{
																									"operator": map[string]interface{}{
																										"description": `Represents the relationship between the container exit code(s) and the
specified values. Possible values are:
- In: the requirement is satisfied if the container exit code is in the
  set of specified values.
- NotIn: the requirement is satisfied if the container exit code is
  not in the set of specified values.`,
																										"type": "string",
																									},
																									"values": map[string]interface{}{
																										"description": `Specifies the set of values to check for container exit codes.
At most 255 elements are allowed.`,
																										"items": map[string]interface{}{
																											"format": "int32",
																											"type":   "integer",
																										},
																										"type":                   "array",
																										"x-kubernetes-list-type": "set",
																									},
																								},
																								"required": []interface{}{
																									"operator",
																								},
																								"type": "object",
																							},
																						},
																						"required": []interface{}{
																							"action",
																						},
																						"type": "object",
																					},
																					"type":                   "array",
																					"x-kubernetes-list-type": "atomic",
																				},
																				"securityContext": map[string]interface{}{
																					"description": `Optional: SecurityContext defines the security options the ephemeral container should be run with.
If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext.`,
																					"properties": map[string]interface{}{
																						"allowPrivilegeEscalation": map[string]interface{}{
																							"description": `AllowPrivilegeEscalation controls whether a process can gain more
privileges than its parent process. This bool directly controls if
the no_new_privs flag will be set on the container process.
AllowPrivilegeEscalation is true always when the container is:
1) run as Privileged
2) has CAP_SYS_ADMIN
Note that this field cannot be set when spec.os.name is windows.`,
																							"type": "boolean",
																						},
																						"appArmorProfile": map[string]interface{}{
																							"description": `appArmorProfile is the AppArmor options to use by this container. If set, this profile
overrides the pod's appArmorProfile.
Note that this field cannot be set when spec.os.name is windows.`,
																							"properties": map[string]interface{}{
																								"localhostProfile": map[string]interface{}{
																									"description": `localhostProfile indicates a profile loaded on the node that should be used.
The profile must be preconfigured on the node to work.
Must match the loaded name of the profile.
Must be set if and only if type is "Localhost".`,
																									"type": "string",
																								},
																								"type": map[string]interface{}{
																									"description": `type indicates which kind of AppArmor profile will be applied.
Valid options are:
  Localhost - a profile pre-loaded on the node.
  RuntimeDefault - the container runtime's default profile.
  Unconfined - no AppArmor enforcement.`,
																									"type": "string",
																								},
																							},
																							"required": []interface{}{
																								"type",
																							},
																							"type": "object",
																						},
																						"capabilities": map[string]interface{}{
																							"description": `The capabilities to add/drop when running containers.
Defaults to the default set of capabilities granted by the container runtime.
Note that this field cannot be set when spec.os.name is windows.`,
																							"properties": map[string]interface{}{
																								"add": map[string]interface{}{
																									"description": "Added capabilities",
																									"items": map[string]interface{}{
																										"description": "Capability represent POSIX capabilities type",
																										"type":        "string",
																									},
																									"type":                   "array",
																									"x-kubernetes-list-type": "atomic",
																								},
																								"drop": map[string]interface{}{
																									"description": "Removed capabilities",
																									"items": map[string]interface{}{
																										"description": "Capability represent POSIX capabilities type",
																										"type":        "string",
																									},
																									"type":                   "array",
																									"x-kubernetes-list-type": "atomic",
																								},
																							},
																							"type": "object",
																						},
																						"privileged": map[string]interface{}{
																							"description": `Run container in privileged mode.
Processes in privileged containers are essentially equivalent to root on the host.
Defaults to false.
Note that this field cannot be set when spec.os.name is windows.`,
																							"type": "boolean",
																						},
																						"procMount": map[string]interface{}{
																							"description": `procMount denotes the type of proc mount to use for the containers.
The default value is Default which uses the container runtime defaults for
readonly paths and masked paths.
Note that this field cannot be set when spec.os.name is windows.`,
																							"type": "string",
																						},
																						"readOnlyRootFilesystem": map[string]interface{}{
																							"description": `Whether this container has a read-only root filesystem.
Default is false.
Note that this field cannot be set when spec.os.name is windows.`,
																							"type": "boolean",
																						},
																						"runAsGroup": map[string]interface{}{
																							"description": `The GID to run the entrypoint of the container process.
Uses runtime default if unset.
May also be set in PodSecurityContext.  If set in both SecurityContext and
PodSecurityContext, the value specified in SecurityContext takes precedence.
Note that this field cannot be set when spec.os.name is windows.`,
																							"format": "int64",
																							"type":   "integer",
																						},
																						"runAsNonRoot": map[string]interface{}{
																							"description": `Indicates that the container must run as a non-root user.
If true, the Kubelet will validate the image at runtime to ensure that it
does not run as UID 0 (root) and fail to start the container if it does.
If unset or false, no such validation will be performed.
May also be set in PodSecurityContext.  If set in both SecurityContext and
PodSecurityContext, the value specified in SecurityContext takes precedence.`,
																							"type": "boolean",
																						},
																						"runAsUser": map[string]interface{}{
																							"description": `The UID to run the entrypoint of the container process.
Defaults to user specified in image metadata if unspecified.
May also be set in PodSecurityContext.  If set in both SecurityContext and
PodSecurityContext, the value specified in SecurityContext takes precedence.
Note that this field cannot be set when spec.os.name is windows.`,
																							"format": "int64",
																							"type":   "integer",
																						},
																						"seLinuxOptions": map[string]interface{}{
																							"description": `The SELinux context to be applied to the container.
If unspecified, the container runtime will allocate a random SELinux context for each
container.  May also be set in PodSecurityContext.  If set in both SecurityContext and
PodSecurityContext, the value specified in SecurityContext takes precedence.
Note that this field cannot be set when spec.os.name is windows.`,
																							"properties": map[string]interface{}{
																								"level": map[string]interface{}{
																									"description": "Level is SELinux level label that applies to the container.",
																									"type":        "string",
																								},
																								"role": map[string]interface{}{
																									"description": "Role is a SELinux role label that applies to the container.",
																									"type":        "string",
																								},
																								"type": map[string]interface{}{
																									"description": "Type is a SELinux type label that applies to the container.",
																									"type":        "string",
																								},
																								"user": map[string]interface{}{
																									"description": "User is a SELinux user label that applies to the container.",
																									"type":        "string",
																								},
																							},
																							"type": "object",
																						},
																						"seccompProfile": map[string]interface{}{
																							"description": `The seccomp options to use by this container. If seccomp options are
provided at both the pod & container level, the container options
override the pod options.
Note that this field cannot be set when spec.os.name is windows.`,
																							"properties": map[string]interface{}{
																								"localhostProfile": map[string]interface{}{
																									"description": `localhostProfile indicates a profile defined in a file on the node should be used.
The profile must be preconfigured on the node to work.
Must be a descending path, relative to the kubelet's configured seccomp profile location.
Must be set if type is "Localhost". Must NOT be set for any other type.`,
																									"type": "string",
																								},
																								"type": map[string]interface{}{
																									"description": `type indicates which kind of seccomp profile will be applied.
Valid options are:

Localhost - a profile defined in a file on the node should be used.
RuntimeDefault - the container runtime default profile should be used.
Unconfined - no profile should be applied.`,
																									"type": "string",
																								},
																							},
																							"required": []interface{}{
																								"type",
																							},
																							"type": "object",
																						},
																						"windowsOptions": map[string]interface{}{
																							"description": `The Windows specific settings applied to all containers.
If unspecified, the options from the PodSecurityContext will be used.
If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
Note that this field cannot be set when spec.os.name is linux.`,
																							"properties": map[string]interface{}{
																								"gmsaCredentialSpec": map[string]interface{}{
																									"description": `GMSACredentialSpec is where the GMSA admission webhook
(https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the
GMSA credential spec named by the GMSACredentialSpecName field.`,
																									"type": "string",
																								},
																								"gmsaCredentialSpecName": map[string]interface{}{
																									"description": "GMSACredentialSpecName is the name of the GMSA credential spec to use.",
																									"type":        "string",
																								},
																								"hostProcess": map[string]interface{}{
																									"description": `HostProcess determines if a container should be run as a 'Host Process' container.
All of a Pod's containers must have the same effective HostProcess value
(it is not allowed to have a mix of HostProcess containers and non-HostProcess containers).
In addition, if HostProcess is true then HostNetwork must also be set to true.`,
																									"type": "boolean",
																								},
																								"runAsUserName": map[string]interface{}{
																									"description": `The UserName in Windows to run the entrypoint of the container process.
Defaults to the user specified in image metadata if unspecified.
May also be set in PodSecurityContext. If set in both SecurityContext and
PodSecurityContext, the value specified in SecurityContext takes precedence.`,
																									"type": "string",
																								},
																							},
																							"type": "object",
																						},
																					},
																					"type": "object",
																				},
																				"startupProbe": map[string]interface{}{
																					"description": "Probes are not allowed for ephemeral containers.",
																					"properties": map[string]interface{}{
																						"exec": map[string]interface{}{
																							"description": "Exec specifies a command to execute in the container.",
																							"properties": map[string]interface{}{
																								"command": map[string]interface{}{
																									"description": `Command is the command line to execute inside the container, the working directory for the
command  is root ('/') in the container's filesystem. The command is simply exec'd, it is
not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use
a shell, you need to explicitly call out to that shell.
Exit status of 0 is treated as live/healthy and non-zero is unhealthy.`,
																									"items": map[string]interface{}{
																										"type": "string",
																									},
																									"type":                   "array",
																									"x-kubernetes-list-type": "atomic",
																								},
																							},
																							"type": "object",
																						},
																						"failureThreshold": map[string]interface{}{
																							"description": `Minimum consecutive failures for the probe to be considered failed after having succeeded.
Defaults to 3. Minimum value is 1.`,
																							"format": "int32",
																							"type":   "integer",
																						},
																						"grpc": map[string]interface{}{
																							"description": "GRPC specifies a GRPC HealthCheckRequest.",
																							"properties": map[string]interface{}{
																								"port": map[string]interface{}{
																									"description": "Port number of the gRPC service. Number must be in the range 1 to 65535.",
																									"format":      "int32",
																									"type":        "integer",
																								},
																								"service": map[string]interface{}{
																									"default": "",
																									"description": `Service is the name of the service to place in the gRPC HealthCheckRequest
(see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).

If this is not specified, the default behavior is defined by gRPC.`,
																									"type": "string",
																								},
																							},
																							"required": []interface{}{
																								"port",
																							},
																							"type": "object",
																						},
																						"httpGet": map[string]interface{}{
																							"description": "HTTPGet specifies an HTTP GET request to perform.",
																							"properties": map[string]interface{}{
																								"host": map[string]interface{}{
																									"description": `Host name to connect to, defaults to the pod IP. You probably want to set
"Host" in httpHeaders instead.`,
																									"type": "string",
																								},
																								"httpHeaders": map[string]interface{}{
																									"description": "Custom headers to set in the request. HTTP allows repeated headers.",
																									"items": map[string]interface{}{
																										"description": "HTTPHeader describes a custom header to be used in HTTP probes",
																										"properties": map[string]interface{}{
																											"name": map[string]interface{}{
																												"description": `The header field name.
This will be canonicalized upon output, so case-variant names will be understood as the same header.`,
																												"type": "string",
																											},
																											"value": map[string]interface{}{
																												"description": "The header field value",
																												"type":        "string",
																											},
																										},
																										"required": []interface{}{
																											"name",
																											"value",
																										},
																										"type": "object",
																									},
																									"type":                   "array",
																									"x-kubernetes-list-type": "atomic",
																								},
																								"path": map[string]interface{}{
																									"description": "Path to access on the HTTP server.",
																									"type":        "string",
																								},
																								"port": map[string]interface{}{
																									"anyOf": []interface{}{
																										map[string]interface{}{
																											"type": "integer",
																										},
																										map[string]interface{}{
																											"type": "string",
																										},
																									},
																									"description": `Name or number of the port to access on the container.
Number must be in the range 1 to 65535.
Name must be an IANA_SVC_NAME.`,
																									"x-kubernetes-int-or-string": true,
																								},
																								"scheme": map[string]interface{}{
																									"description": `Scheme to use for connecting to the host.
Defaults to HTTP.`,
																									"type": "string",
																								},
																							},
																							"required": []interface{}{
																								"port",
																							},
																							"type": "object",
																						},
																						"initialDelaySeconds": map[string]interface{}{
																							"description": `Number of seconds after the container has started before liveness probes are initiated.
More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes`,
																							"format": "int32",
																							"type":   "integer",
																						},
																						"periodSeconds": map[string]interface{}{
																							"description": `How often (in seconds) to perform the probe.
Default to 10 seconds. Minimum value is 1.`,
																							"format": "int32",
																							"type":   "integer",
																						},
																						"successThreshold": map[string]interface{}{
																							"description": `Minimum consecutive successes for the probe to be considered successful after having failed.
Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.`,
																							"format": "int32",
																							"type":   "integer",
																						},
																						"tcpSocket": map[string]interface{}{
																							"description": "TCPSocket specifies a connection to a TCP port.",
																							"properties": map[string]interface{}{
																								"host": map[string]interface{}{
																									"description": "Optional: Host name to connect to, defaults to the pod IP.",
																									"type":        "string",
																								},
																								"port": map[string]interface{}{
																									"anyOf": []interface{}{
																										map[string]interface{}{
																											"type": "integer",
																										},
																										map[string]interface{}{
																											"type": "string",
																										},
																									},
																									"description": `Number or name of the port to access on the container.
Number must be in the range 1 to 65535.
Name must be an IANA_SVC_NAME.`,
																									"x-kubernetes-int-or-string": true,
																								},
																							},
																							"required": []interface{}{
																								"port",
																							},
																							"type": "object",
																						},
																						"terminationGracePeriodSeconds": map[string]interface{}{
																							"description": `Optional duration in seconds the pod needs to terminate gracefully upon probe failure.
The grace period is the duration in seconds after the processes running in the pod are sent
a termination signal and the time when the processes are forcibly halted with a kill signal.
Set this value longer than the expected cleanup time for your process.
If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this
value overrides the value provided by the pod spec.
Value must be non-negative integer. The value zero indicates stop immediately via
the kill signal (no opportunity to shut down).
This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate.
Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.`,
																							"format": "int64",
																							"type":   "integer",
																						},
																						"timeoutSeconds": map[string]interface{}{
																							"description": `Number of seconds after which the probe times out.
Defaults to 1 second. Minimum value is 1.
More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes`,
																							"format": "int32",
																							"type":   "integer",
																						},
																					},
																					"type": "object",
																				},
																				"stdin": map[string]interface{}{
																					"description": `Whether this container should allocate a buffer for stdin in the container runtime. If this
is not set, reads from stdin in the container will always result in EOF.
Default is false.`,
																					"type": "boolean",
																				},
																				"stdinOnce": map[string]interface{}{
																					"description": `Whether the container runtime should close the stdin channel after it has been opened by
a single attach. When stdin is true the stdin stream will remain open across multiple attach
sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the
first client attaches to stdin, and then remains open and accepts data until the client disconnects,
at which time stdin is closed and remains closed until the container is restarted. If this
flag is false, a container processes that reads from stdin will never receive an EOF.
Default is false`,
																					"type": "boolean",
																				},
																				"targetContainerName": map[string]interface{}{
																					"description": `If set, the name of the container from PodSpec that this ephemeral container targets.
The ephemeral container will be run in the namespaces (IPC, PID, etc) of this container.
If not set then the ephemeral container uses the namespaces configured in the Pod spec.

The container runtime must implement support for this feature. If the runtime does not
support namespace targeting then the result of setting this field is undefined.`,
																					"type": "string",
																				},
																				"terminationMessagePath": map[string]interface{}{
																					"description": `Optional: Path at which the file to which the container's termination message
will be written is mounted into the container's filesystem.
Message written is intended to be brief final status, such as an assertion failure message.
Will be truncated by the node if greater than 4096 bytes. The total message length across
all containers will be limited to 12kb.
Defaults to /dev/termination-log.
Cannot be updated.`,
																					"type": "string",
																				},
																				"terminationMessagePolicy": map[string]interface{}{
																					"description": `Indicate how the termination message should be populated. File will use the contents of
terminationMessagePath to populate the container status message on both success and failure.
FallbackToLogsOnError will use the last chunk of container log output if the termination
message file is empty and the container exited with an error.
The log output is limited to 2048 bytes or 80 lines, whichever is smaller.
Defaults to File.
Cannot be updated.`,
																					"type": "string",
																				},
																				"tty": map[string]interface{}{
																					"description": `Whether this container should allocate a TTY for itself, also requires 'stdin' to be true.
Default is false.`,
																					"type": "boolean",
																				},
																				"volumeDevices": map[string]interface{}{
																					"description": "volumeDevices is the list of block devices to be used by the container.",
																					"items": map[string]interface{}{
																						"description": "volumeDevice describes a mapping of a raw block device within a container.",
																						"properties": map[string]interface{}{
																							"devicePath": map[string]interface{}{
																								"description": "devicePath is the path inside of the container that the device will be mapped to.",
																								"type":        "string",
																							},
																							"name": map[string]interface{}{
																								"description": "name must match the name of a persistentVolumeClaim in the pod",
																								"type":        "string",
																							},
																						},
																						"required": []interface{}{
																							"devicePath",
																							"name",
																						},
																						"type": "object",
																					},
																					"type": "array",
																					"x-kubernetes-list-map-keys": []interface{}{
																						"devicePath",
																					},
																					"x-kubernetes-list-type": "map",
																				},
																				"volumeMounts": map[string]interface{}{
																					"description": `Pod volumes to mount into the container's filesystem. Subpath mounts are not allowed for ephemeral containers.
Cannot be updated.`,
																					"items": map[string]interface{}{
																						"description": "VolumeMount describes a mounting of a Volume within a container.",
																						"properties": map[string]interface{}{
																							"mountPath": map[string]interface{}{
																								"description": `Path within the container at which the volume should be mounted.  Must
not contain ':'.`,
																								"type": "string",
																							},
																							"mountPropagation": map[string]interface{}{
																								"description": `mountPropagation determines how mounts are propagated from the host
to container and the other way around.
When not set, MountPropagationNone is used.
This field is beta in 1.10.
When RecursiveReadOnly is set to IfPossible or to Enabled, MountPropagation must be None or unspecified
(which defaults to None).`,
																								"type": "string",
																							},
																							"name": map[string]interface{}{
																								"description": "This must match the Name of a Volume.",
																								"type":        "string",
																							},
																							"readOnly": map[string]interface{}{
																								"description": `Mounted read-only if true, read-write otherwise (false or unspecified).
Defaults to false.`,
																								"type": "boolean",
																							},
																							"recursiveReadOnly": map[string]interface{}{
																								"description": `RecursiveReadOnly specifies whether read-only mounts should be handled
recursively.

If ReadOnly is false, this field has no meaning and must be unspecified.

If ReadOnly is true, and this field is set to Disabled, the mount is not made
recursively read-only.  If this field is set to IfPossible, the mount is made
recursively read-only, if it is supported by the container runtime.  If this
field is set to Enabled, the mount is made recursively read-only if it is
supported by the container runtime, otherwise the pod will not be started and
an error will be generated to indicate the reason.

If this field is set to IfPossible or Enabled, MountPropagation must be set to
None (or be unspecified, which defaults to None).

If this field is not specified, it is treated as an equivalent of Disabled.`,
																								"type": "string",
																							},
																							"subPath": map[string]interface{}{
																								"description": `Path within the volume from which the container's volume should be mounted.
Defaults to "" (volume's root).`,
																								"type": "string",
																							},
																							"subPathExpr": map[string]interface{}{
																								"description": `Expanded path within the volume from which the container's volume should be mounted.
Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment.
Defaults to "" (volume's root).
SubPathExpr and SubPath are mutually exclusive.`,
																								"type": "string",
																							},
																						},
																						"required": []interface{}{
																							"mountPath",
																							"name",
																						},
																						"type": "object",
																					},
																					"type": "array",
																					"x-kubernetes-list-map-keys": []interface{}{
																						"mountPath",
																					},
																					"x-kubernetes-list-type": "map",
																				},
																				"workingDir": map[string]interface{}{
																					"description": `Container's working directory.
If not specified, the container runtime's default will be used, which
might be configured in the container image.
Cannot be updated.`,
																					"type": "string",
																				},
																			},
																			"required": []interface{}{
																				"name",
																			},
																			"type": "object",
																		},
																		"type": "array",
																		"x-kubernetes-list-map-keys": []interface{}{
																			"name",
																		},
																		"x-kubernetes-list-type": "map",
																	},
																	"hostAliases": map[string]interface{}{
																		"description": `HostAliases is an optional list of hosts and IPs that will be injected into the pod's hosts
file if specified.`,
																		"items": map[string]interface{}{
																			"description": `HostAlias holds the mapping between IP and hostnames that will be injected as an entry in the
pod's hosts file.`,
																			"properties": map[string]interface{}{
																				"hostnames": map[string]interface{}{
																					"description": "Hostnames for the above IP address.",
																					"items": map[string]interface{}{
																						"type": "string",
																					},
																					"type":                   "array",
																					"x-kubernetes-list-type": "atomic",
																				},
																				"ip": map[string]interface{}{
																					"description": "IP address of the host file entry.",
																					"type":        "string",
																				},
																			},
																			"required": []interface{}{
																				"ip",
																			},
																			"type": "object",
																		},
																		"type": "array",
																		"x-kubernetes-list-map-keys": []interface{}{
																			"ip",
																		},
																		"x-kubernetes-list-type": "map",
																	},
																	"hostIPC": map[string]interface{}{
																		"description": `Use the host's ipc namespace.
Optional: Default to false.`,
																		"type": "boolean",
																	},
																	"hostNetwork": map[string]interface{}{
																		"description": `Host networking requested for this pod. Use the host's network namespace.
When using HostNetwork you should specify ports so the scheduler is aware.
When ` + "`" + `hostNetwork` + "`" + ` is true, specified ` + "`" + `hostPort` + "`" + ` fields in port definitions must match ` + "`" + `containerPort` + "`" + `,
and unspecified ` + "`" + `hostPort` + "`" + ` fields in port definitions are defaulted to match ` + "`" + `containerPort` + "`" + `.
Default to false.`,
																		"type": "boolean",
																	},
																	"hostPID": map[string]interface{}{
																		"description": `Use the host's pid namespace.
Optional: Default to false.`,
																		"type": "boolean",
																	},
																	"hostUsers": map[string]interface{}{
																		"description": `Use the host's user namespace.
Optional: Default to true.
If set to true or not present, the pod will be run in the host user namespace, useful
for when the pod needs a feature only available to the host user namespace, such as
loading a kernel module with CAP_SYS_MODULE.
When set to false, a new userns is created for the pod. Setting false is useful for
mitigating container breakout vulnerabilities even allowing users to run their
containers as root without actually having root privileges on the host.`,
																		"type": "boolean",
																	},
																	"hostname": map[string]interface{}{
																		"description": `Specifies the hostname of the Pod
If not specified, the pod's hostname will be set to a system-defined value.`,
																		"type": "string",
																	},
																	"hostnameOverride": map[string]interface{}{
																		"description": `HostnameOverride specifies an explicit override for the pod's hostname as perceived by the pod.
This field only specifies the pod's hostname and does not affect its DNS records.
When this field is set to a non-empty string:
- It takes precedence over the values set in ` + "`" + `hostname` + "`" + ` and ` + "`" + `subdomain` + "`" + `.
- The Pod's hostname will be set to this value.
- ` + "`" + `setHostnameAsFQDN` + "`" + ` must be nil or set to false.
- ` + "`" + `hostNetwork` + "`" + ` must be set to false.

This field must be a valid DNS subdomain as defined in RFC 1123 and contain at most 64 characters.
Requires the HostnameOverride feature gate to be enabled.`,
																		"type": "string",
																	},
																	"imagePullSecrets": map[string]interface{}{
																		"description": `ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.
If specified, these secrets will be passed to individual puller implementations for them to use.
More info: https://kubernetes.io/docs/concepts/containers/images#specifying-imagepullsecrets-on-a-pod`,
																		"items": map[string]interface{}{
																			"description": `LocalObjectReference contains enough information to let you locate the
referenced object inside the same namespace.`,
																			"properties": map[string]interface{}{
																				"name": map[string]interface{}{
																					"default": "",
																					"description": `Name of the referent.
This field is effectively required, but due to backwards compatibility is
allowed to be empty. Instances of this type with an empty value here are
almost certainly wrong.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names`,
																					"type": "string",
																				},
																			},
																			"type":                  "object",
																			"x-kubernetes-map-type": "atomic",
																		},
																		"type": "array",
																		"x-kubernetes-list-map-keys": []interface{}{
																			"name",
																		},
																		"x-kubernetes-list-type": "map",
																	},
																	"initContainers": map[string]interface{}{
																		"description": `List of initialization containers belonging to the pod.
Init containers are executed in order prior to containers being started. If any
init container fails, the pod is considered to have failed and is handled according
to its restartPolicy. The name for an init container or normal container must be
unique among all containers.
Init containers may not have Lifecycle actions, Readiness probes, Liveness probes, or Startup probes.
The resourceRequirements of an init container are taken into account during scheduling
by finding the highest request/limit for each resource type, and then using the max of
that value or the sum of the normal containers. Limits are applied to init containers
in a similar fashion.
Init containers cannot currently be added or removed.
Cannot be updated.
More info: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/`,
																		"items": map[string]interface{}{
																			"description": "A single application container that you want to run within a pod.",
																			"properties": map[string]interface{}{
																				"args": map[string]interface{}{
																					"description": `Arguments to the entrypoint.
The container image's CMD is used if this is not provided.
Variable references $(VAR_NAME) are expanded using the container's environment. If a variable
cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced
to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. "$$(VAR_NAME)" will
produce the string literal "$(VAR_NAME)". Escaped references will never be expanded, regardless
of whether the variable exists or not. Cannot be updated.
More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell`,
																					"items": map[string]interface{}{
																						"type": "string",
																					},
																					"type":                   "array",
																					"x-kubernetes-list-type": "atomic",
																				},
																				"command": map[string]interface{}{
																					"description": `Entrypoint array. Not executed within a shell.
The container image's ENTRYPOINT is used if this is not provided.
Variable references $(VAR_NAME) are expanded using the container's environment. If a variable
cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced
to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. "$$(VAR_NAME)" will
produce the string literal "$(VAR_NAME)". Escaped references will never be expanded, regardless
of whether the variable exists or not. Cannot be updated.
More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell`,
																					"items": map[string]interface{}{
																						"type": "string",
																					},
																					"type":                   "array",
																					"x-kubernetes-list-type": "atomic",
																				},
																				"env": map[string]interface{}{
																					"description": `List of environment variables to set in the container.
Cannot be updated.`,
																					"items": map[string]interface{}{
																						"description": "EnvVar represents an environment variable present in a Container.",
																						"properties": map[string]interface{}{
																							"name": map[string]interface{}{
																								"description": `Name of the environment variable.
May consist of any printable ASCII characters except '='.`,
																								"type": "string",
																							},
																							"value": map[string]interface{}{
																								"description": `Variable references $(VAR_NAME) are expanded
using the previously defined environment variables in the container and
any service environment variables. If a variable cannot be resolved,
the reference in the input string will be unchanged. Double $$ are reduced
to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e.
"$$(VAR_NAME)" will produce the string literal "$(VAR_NAME)".
Escaped references will never be expanded, regardless of whether the variable
exists or not.
Defaults to "".`,
																								"type": "string",
																							},
																							"valueFrom": map[string]interface{}{
																								"description": "Source for the environment variable's value. Cannot be used if value is not empty.",
																								"properties": map[string]interface{}{
																									"configMapKeyRef": map[string]interface{}{
																										"description": "Selects a key of a ConfigMap.",
																										"properties": map[string]interface{}{
																											"key": map[string]interface{}{
																												"description": "The key to select.",
																												"type":        "string",
																											},
																											"name": map[string]interface{}{
																												"default": "",
																												"description": `Name of the referent.
This field is effectively required, but due to backwards compatibility is
allowed to be empty. Instances of this type with an empty value here are
almost certainly wrong.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names`,
																												"type": "string",
																											},
																											"optional": map[string]interface{}{
																												"description": "Specify whether the ConfigMap or its key must be defined",
																												"type":        "boolean",
																											},
																										},
																										"required": []interface{}{
																											"key",
																										},
																										"type":                  "object",
																										"x-kubernetes-map-type": "atomic",
																									},
																									"fieldRef": map[string]interface{}{
																										"description": `Selects a field of the pod: supports metadata.name, metadata.namespace, ` + "`" + `metadata.labels['<KEY>']` + "`" + `, ` + "`" + `metadata.annotations['<KEY>']` + "`" + `,
spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.`,
																										"properties": map[string]interface{}{
																											"apiVersion": map[string]interface{}{
																												"description": "Version of the schema the FieldPath is written in terms of, defaults to \"v1\".",
																												"type":        "string",
																											},
																											"fieldPath": map[string]interface{}{
																												"description": "Path of the field to select in the specified API version.",
																												"type":        "string",
																											},
																										},
																										"required": []interface{}{
																											"fieldPath",
																										},
																										"type":                  "object",
																										"x-kubernetes-map-type": "atomic",
																									},
																									"fileKeyRef": map[string]interface{}{
																										"description": `FileKeyRef selects a key of the env file.
Requires the EnvFiles feature gate to be enabled.`,
																										"properties": map[string]interface{}{
																											"key": map[string]interface{}{
																												"description": `The key within the env file. An invalid key will prevent the pod from starting.
The keys defined within a source may consist of any printable ASCII characters except '='.
During Alpha stage of the EnvFiles feature gate, the key size is limited to 128 characters.`,
																												"type": "string",
																											},
																											"optional": map[string]interface{}{
																												"default": false,
																												"description": `Specify whether the file or its key must be defined. If the file or key
does not exist, then the env var is not published.
If optional is set to true and the specified key does not exist,
the environment variable will not be set in the Pod's containers.

If optional is set to false and the specified key does not exist,
an error will be returned during Pod creation.`,
																												"type": "boolean",
																											},
																											"path": map[string]interface{}{
																												"description": `The path within the volume from which to select the file.
Must be relative and may not contain the '..' path or start with '..'.`,
																												"type": "string",
																											},
																											"volumeName": map[string]interface{}{
																												"description": "The name of the volume mount containing the env file.",
																												"type":        "string",
																											},
																										},
																										"required": []interface{}{
																											"key",
																											"path",
																											"volumeName",
																										},
																										"type":                  "object",
																										"x-kubernetes-map-type": "atomic",
																									},
																									"resourceFieldRef": map[string]interface{}{
																										"description": `Selects a resource of the container: only resources limits and requests
(limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.`,
																										"properties": map[string]interface{}{
																											"containerName": map[string]interface{}{
																												"description": "Container name: required for volumes, optional for env vars",
																												"type":        "string",
																											},
																											"divisor": map[string]interface{}{
																												"anyOf": []interface{}{
																													map[string]interface{}{
																														"type": "integer",
																													},
																													map[string]interface{}{
																														"type": "string",
																													},
																												},
																												"description":                "Specifies the output format of the exposed resources, defaults to \"1\"",
																												"pattern":                    `^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$`,
																												"x-kubernetes-int-or-string": true,
																											},
																											"resource": map[string]interface{}{
																												"description": "Required: resource to select",
																												"type":        "string",
																											},
																										},
																										"required": []interface{}{
																											"resource",
																										},
																										"type":                  "object",
																										"x-kubernetes-map-type": "atomic",
																									},
																									"secretKeyRef": map[string]interface{}{
																										"description": "Selects a key of a secret in the pod's namespace",
																										"properties": map[string]interface{}{
																											"key": map[string]interface{}{
																												"description": "The key of the secret to select from.  Must be a valid secret key.",
																												"type":        "string",
																											},
																											"name": map[string]interface{}{
																												"default": "",
																												"description": `Name of the referent.
This field is effectively required, but due to backwards compatibility is
allowed to be empty. Instances of this type with an empty value here are
almost certainly wrong.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names`,
																												"type": "string",
																											},
																											"optional": map[string]interface{}{
																												"description": "Specify whether the Secret or its key must be defined",
																												"type":        "boolean",
																											},
																										},
																										"required": []interface{}{
																											"key",
																										},
																										"type":                  "object",
																										"x-kubernetes-map-type": "atomic",
																									},
																								},
																								"type": "object",
																							},
																						},
																						"required": []interface{}{
																							"name",
																						},
																						"type": "object",
																					},
																					"type": "array",
																					"x-kubernetes-list-map-keys": []interface{}{
																						"name",
																					},
																					"x-kubernetes-list-type": "map",
																				},
																				"envFrom": map[string]interface{}{
																					"description": `List of sources to populate environment variables in the container.
The keys defined within a source may consist of any printable ASCII characters except '='.
When a key exists in multiple
sources, the value associated with the last source will take precedence.
Values defined by an Env with a duplicate key will take precedence.
Cannot be updated.`,
																					"items": map[string]interface{}{
																						"description": "EnvFromSource represents the source of a set of ConfigMaps or Secrets",
																						"properties": map[string]interface{}{
																							"configMapRef": map[string]interface{}{
																								"description": "The ConfigMap to select from",
																								"properties": map[string]interface{}{
																									"name": map[string]interface{}{
																										"default": "",
																										"description": `Name of the referent.
This field is effectively required, but due to backwards compatibility is
allowed to be empty. Instances of this type with an empty value here are
almost certainly wrong.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names`,
																										"type": "string",
																									},
																									"optional": map[string]interface{}{
																										"description": "Specify whether the ConfigMap must be defined",
																										"type":        "boolean",
																									},
																								},
																								"type":                  "object",
																								"x-kubernetes-map-type": "atomic",
																							},
																							"prefix": map[string]interface{}{
																								"description": `Optional text to prepend to the name of each environment variable.
May consist of any printable ASCII characters except '='.`,
																								"type": "string",
																							},
																							"secretRef": map[string]interface{}{
																								"description": "The Secret to select from",
																								"properties": map[string]interface{}{
																									"name": map[string]interface{}{
																										"default": "",
																										"description": `Name of the referent.
This field is effectively required, but due to backwards compatibility is
allowed to be empty. Instances of this type with an empty value here are
almost certainly wrong.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names`,
																										"type": "string",
																									},
																									"optional": map[string]interface{}{
																										"description": "Specify whether the Secret must be defined",
																										"type":        "boolean",
																									},
																								},
																								"type":                  "object",
																								"x-kubernetes-map-type": "atomic",
																							},
																						},
																						"type": "object",
																					},
																					"type":                   "array",
																					"x-kubernetes-list-type": "atomic",
																				},
																				"image": map[string]interface{}{
																					"description": `Container image name.
More info: https://kubernetes.io/docs/concepts/containers/images
This field is optional to allow higher level config management to default or override
container images in workload controllers like Deployments and StatefulSets.`,
																					"type": "string",
																				},
																				"imagePullPolicy": map[string]interface{}{
																					"description": `Image pull policy.
One of Always, Never, IfNotPresent.
Defaults to Always if :latest tag is specified, or IfNotPresent otherwise.
Cannot be updated.
More info: https://kubernetes.io/docs/concepts/containers/images#updating-images`,
																					"type": "string",
																				},
																				"lifecycle": map[string]interface{}{
																					"description": `Actions that the management system should take in response to container lifecycle events.
Cannot be updated.`,
																					"properties": map[string]interface{}{
																						"postStart": map[string]interface{}{
																							"description": `PostStart is called immediately after a container is created. If the handler fails,
the container is terminated and restarted according to its restart policy.
Other management of the container blocks until the hook completes.
More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks`,
																							"properties": map[string]interface{}{
																								"exec": map[string]interface{}{
																									"description": "Exec specifies a command to execute in the container.",
																									"properties": map[string]interface{}{
																										"command": map[string]interface{}{
																											"description": `Command is the command line to execute inside the container, the working directory for the
command  is root ('/') in the container's filesystem. The command is simply exec'd, it is
not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use
a shell, you need to explicitly call out to that shell.
Exit status of 0 is treated as live/healthy and non-zero is unhealthy.`,
																											"items": map[string]interface{}{
																												"type": "string",
																											},
																											"type":                   "array",
																											"x-kubernetes-list-type": "atomic",
																										},
																									},
																									"type": "object",
																								},
																								"httpGet": map[string]interface{}{
																									"description": "HTTPGet specifies an HTTP GET request to perform.",
																									"properties": map[string]interface{}{
																										"host": map[string]interface{}{
																											"description": `Host name to connect to, defaults to the pod IP. You probably want to set
"Host" in httpHeaders instead.`,
																											"type": "string",
																										},
																										"httpHeaders": map[string]interface{}{
																											"description": "Custom headers to set in the request. HTTP allows repeated headers.",
																											"items": map[string]interface{}{
																												"description": "HTTPHeader describes a custom header to be used in HTTP probes",
																												"properties": map[string]interface{}{
																													"name": map[string]interface{}{
																														"description": `The header field name.
This will be canonicalized upon output, so case-variant names will be understood as the same header.`,
																														"type": "string",
																													},
																													"value": map[string]interface{}{
																														"description": "The header field value",
																														"type":        "string",
																													},
																												},
																												"required": []interface{}{
																													"name",
																													"value",
																												},
																												"type": "object",
																											},
																											"type":                   "array",
																											"x-kubernetes-list-type": "atomic",
																										},
																										"path": map[string]interface{}{
																											"description": "Path to access on the HTTP server.",
																											"type":        "string",
																										},
																										"port": map[string]interface{}{
																											"anyOf": []interface{}{
																												map[string]interface{}{
																													"type": "integer",
																												},
																												map[string]interface{}{
																													"type": "string",
																												},
																											},
																											"description": `Name or number of the port to access on the container.
Number must be in the range 1 to 65535.
Name must be an IANA_SVC_NAME.`,
																											"x-kubernetes-int-or-string": true,
																										},
																										"scheme": map[string]interface{}{
																											"description": `Scheme to use for connecting to the host.
Defaults to HTTP.`,
																											"type": "string",
																										},
																									},
																									"required": []interface{}{
																										"port",
																									},
																									"type": "object",
																								},
																								"sleep": map[string]interface{}{
																									"description": "Sleep represents a duration that the container should sleep.",
																									"properties": map[string]interface{}{
																										"seconds": map[string]interface{}{
																											"description": "Seconds is the number of seconds to sleep.",
																											"format":      "int64",
																											"type":        "integer",
																										},
																									},
																									"required": []interface{}{
																										"seconds",
																									},
																									"type": "object",
																								},
																								"tcpSocket": map[string]interface{}{
																									"description": `Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept
for backward compatibility. There is no validation of this field and
lifecycle hooks will fail at runtime when it is specified.`,
																									"properties": map[string]interface{}{
																										"host": map[string]interface{}{
																											"description": "Optional: Host name to connect to, defaults to the pod IP.",
																											"type":        "string",
																										},
																										"port": map[string]interface{}{
																											"anyOf": []interface{}{
																												map[string]interface{}{
																													"type": "integer",
																												},
																												map[string]interface{}{
																													"type": "string",
																												},
																											},
																											"description": `Number or name of the port to access on the container.
Number must be in the range 1 to 65535.
Name must be an IANA_SVC_NAME.`,
																											"x-kubernetes-int-or-string": true,
																										},
																									},
																									"required": []interface{}{
																										"port",
																									},
																									"type": "object",
																								},
																							},
																							"type": "object",
																						},
																						"preStop": map[string]interface{}{
																							"description": `PreStop is called immediately before a container is terminated due to an
API request or management event such as liveness/startup probe failure,
preemption, resource contention, etc. The handler is not called if the
container crashes or exits. The Pod's termination grace period countdown begins before the
PreStop hook is executed. Regardless of the outcome of the handler, the
container will eventually terminate within the Pod's termination grace
period (unless delayed by finalizers). Other management of the container blocks until the hook completes
or until the termination grace period is reached.
More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks`,
																							"properties": map[string]interface{}{
																								"exec": map[string]interface{}{
																									"description": "Exec specifies a command to execute in the container.",
																									"properties": map[string]interface{}{
																										"command": map[string]interface{}{
																											"description": `Command is the command line to execute inside the container, the working directory for the
command  is root ('/') in the container's filesystem. The command is simply exec'd, it is
not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use
a shell, you need to explicitly call out to that shell.
Exit status of 0 is treated as live/healthy and non-zero is unhealthy.`,
																											"items": map[string]interface{}{
																												"type": "string",
																											},
																											"type":                   "array",
																											"x-kubernetes-list-type": "atomic",
																										},
																									},
																									"type": "object",
																								},
																								"httpGet": map[string]interface{}{
																									"description": "HTTPGet specifies an HTTP GET request to perform.",
																									"properties": map[string]interface{}{
																										"host": map[string]interface{}{
																											"description": `Host name to connect to, defaults to the pod IP. You probably want to set
"Host" in httpHeaders instead.`,
																											"type": "string",
																										},
																										"httpHeaders": map[string]interface{}{
																											"description": "Custom headers to set in the request. HTTP allows repeated headers.",
																											"items": map[string]interface{}{
																												"description": "HTTPHeader describes a custom header to be used in HTTP probes",
																												"properties": map[string]interface{}{
																													"name": map[string]interface{}{
																														"description": `The header field name.
This will be canonicalized upon output, so case-variant names will be understood as the same header.`,
																														"type": "string",
																													},
																													"value": map[string]interface{}{
																														"description": "The header field value",
																														"type":        "string",
																													},
																												},
																												"required": []interface{}{
																													"name",
																													"value",
																												},
																												"type": "object",
																											},
																											"type":                   "array",
																											"x-kubernetes-list-type": "atomic",
																										},
																										"path": map[string]interface{}{
																											"description": "Path to access on the HTTP server.",
																											"type":        "string",
																										},
																										"port": map[string]interface{}{
																											"anyOf": []interface{}{
																												map[string]interface{}{
																													"type": "integer",
																												},
																												map[string]interface{}{
																													"type": "string",
																												},
																											},
																											"description": `Name or number of the port to access on the container.
Number must be in the range 1 to 65535.
Name must be an IANA_SVC_NAME.`,
																											"x-kubernetes-int-or-string": true,
																										},
																										"scheme": map[string]interface{}{
																											"description": `Scheme to use for connecting to the host.
Defaults to HTTP.`,
																											"type": "string",
																										},
																									},
																									"required": []interface{}{
																										"port",
																									},
																									"type": "object",
																								},
																								"sleep": map[string]interface{}{
																									"description": "Sleep represents a duration that the container should sleep.",
																									"properties": map[string]interface{}{
																										"seconds": map[string]interface{}{
																											"description": "Seconds is the number of seconds to sleep.",
																											"format":      "int64",
																											"type":        "integer",
																										},
																									},
																									"required": []interface{}{
																										"seconds",
																									},
																									"type": "object",
																								},
																								"tcpSocket": map[string]interface{}{
																									"description": `Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept
for backward compatibility. There is no validation of this field and
lifecycle hooks will fail at runtime when it is specified.`,
																									"properties": map[string]interface{}{
																										"host": map[string]interface{}{
																											"description": "Optional: Host name to connect to, defaults to the pod IP.",
																											"type":        "string",
																										},
																										"port": map[string]interface{}{
																											"anyOf": []interface{}{
																												map[string]interface{}{
																													"type": "integer",
																												},
																												map[string]interface{}{
																													"type": "string",
																												},
																											},
																											"description": `Number or name of the port to access on the container.
Number must be in the range 1 to 65535.
Name must be an IANA_SVC_NAME.`,
																											"x-kubernetes-int-or-string": true,
																										},
																									},
																									"required": []interface{}{
																										"port",
																									},
																									"type": "object",
																								},
																							},
																							"type": "object",
																						},
																						"stopSignal": map[string]interface{}{
																							"description": `StopSignal defines which signal will be sent to a container when it is being stopped.
If not specified, the default is defined by the container runtime in use.
StopSignal can only be set for Pods with a non-empty .spec.os.name`,
																							"type": "string",
																						},
																					},
																					"type": "object",
																				},
																				"livenessProbe": map[string]interface{}{
																					"description": `Periodic probe of container liveness.
Container will be restarted if the probe fails.
Cannot be updated.
More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes`,
																					"properties": map[string]interface{}{
																						"exec": map[string]interface{}{
																							"description": "Exec specifies a command to execute in the container.",
																							"properties": map[string]interface{}{
																								"command": map[string]interface{}{
																									"description": `Command is the command line to execute inside the container, the working directory for the
command  is root ('/') in the container's filesystem. The command is simply exec'd, it is
not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use
a shell, you need to explicitly call out to that shell.
Exit status of 0 is treated as live/healthy and non-zero is unhealthy.`,
																									"items": map[string]interface{}{
																										"type": "string",
																									},
																									"type":                   "array",
																									"x-kubernetes-list-type": "atomic",
																								},
																							},
																							"type": "object",
																						},
																						"failureThreshold": map[string]interface{}{
																							"description": `Minimum consecutive failures for the probe to be considered failed after having succeeded.
Defaults to 3. Minimum value is 1.`,
																							"format": "int32",
																							"type":   "integer",
																						},
																						"grpc": map[string]interface{}{
																							"description": "GRPC specifies a GRPC HealthCheckRequest.",
																							"properties": map[string]interface{}{
																								"port": map[string]interface{}{
																									"description": "Port number of the gRPC service. Number must be in the range 1 to 65535.",
																									"format":      "int32",
																									"type":        "integer",
																								},
																								"service": map[string]interface{}{
																									"default": "",
																									"description": `Service is the name of the service to place in the gRPC HealthCheckRequest
(see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).

If this is not specified, the default behavior is defined by gRPC.`,
																									"type": "string",
																								},
																							},
																							"required": []interface{}{
																								"port",
																							},
																							"type": "object",
																						},
																						"httpGet": map[string]interface{}{
																							"description": "HTTPGet specifies an HTTP GET request to perform.",
																							"properties": map[string]interface{}{
																								"host": map[string]interface{}{
																									"description": `Host name to connect to, defaults to the pod IP. You probably want to set
"Host" in httpHeaders instead.`,
																									"type": "string",
																								},
																								"httpHeaders": map[string]interface{}{
																									"description": "Custom headers to set in the request. HTTP allows repeated headers.",
																									"items": map[string]interface{}{
																										"description": "HTTPHeader describes a custom header to be used in HTTP probes",
																										"properties": map[string]interface{}{
																											"name": map[string]interface{}{
																												"description": `The header field name.
This will be canonicalized upon output, so case-variant names will be understood as the same header.`,
																												"type": "string",
																											},
																											"value": map[string]interface{}{
																												"description": "The header field value",
																												"type":        "string",
																											},
																										},
																										"required": []interface{}{
																											"name",
																											"value",
																										},
																										"type": "object",
																									},
																									"type":                   "array",
																									"x-kubernetes-list-type": "atomic",
																								},
																								"path": map[string]interface{}{
																									"description": "Path to access on the HTTP server.",
																									"type":        "string",
																								},
																								"port": map[string]interface{}{
																									"anyOf": []interface{}{
																										map[string]interface{}{
																											"type": "integer",
																										},
																										map[string]interface{}{
																											"type": "string",
																										},
																									},
																									"description": `Name or number of the port to access on the container.
Number must be in the range 1 to 65535.
Name must be an IANA_SVC_NAME.`,
																									"x-kubernetes-int-or-string": true,
																								},
																								"scheme": map[string]interface{}{
																									"description": `Scheme to use for connecting to the host.
Defaults to HTTP.`,
																									"type": "string",
																								},
																							},
																							"required": []interface{}{
																								"port",
																							},
																							"type": "object",
																						},
																						"initialDelaySeconds": map[string]interface{}{
																							"description": `Number of seconds after the container has started before liveness probes are initiated.
More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes`,
																							"format": "int32",
																							"type":   "integer",
																						},
																						"periodSeconds": map[string]interface{}{
																							"description": `How often (in seconds) to perform the probe.
Default to 10 seconds. Minimum value is 1.`,
																							"format": "int32",
																							"type":   "integer",
																						},
																						"successThreshold": map[string]interface{}{
																							"description": `Minimum consecutive successes for the probe to be considered successful after having failed.
Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.`,
																							"format": "int32",
																							"type":   "integer",
																						},
																						"tcpSocket": map[string]interface{}{
																							"description": "TCPSocket specifies a connection to a TCP port.",
																							"properties": map[string]interface{}{
																								"host": map[string]interface{}{
																									"description": "Optional: Host name to connect to, defaults to the pod IP.",
																									"type":        "string",
																								},
																								"port": map[string]interface{}{
																									"anyOf": []interface{}{
																										map[string]interface{}{
																											"type": "integer",
																										},
																										map[string]interface{}{
																											"type": "string",
																										},
																									},
																									"description": `Number or name of the port to access on the container.
Number must be in the range 1 to 65535.
Name must be an IANA_SVC_NAME.`,
																									"x-kubernetes-int-or-string": true,
																								},
																							},
																							"required": []interface{}{
																								"port",
																							},
																							"type": "object",
																						},
																						"terminationGracePeriodSeconds": map[string]interface{}{
																							"description": `Optional duration in seconds the pod needs to terminate gracefully upon probe failure.
The grace period is the duration in seconds after the processes running in the pod are sent
a termination signal and the time when the processes are forcibly halted with a kill signal.
Set this value longer than the expected cleanup time for your process.
If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this
value overrides the value provided by the pod spec.
Value must be non-negative integer. The value zero indicates stop immediately via
the kill signal (no opportunity to shut down).
This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate.
Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.`,
																							"format": "int64",
																							"type":   "integer",
																						},
																						"timeoutSeconds": map[string]interface{}{
																							"description": `Number of seconds after which the probe times out.
Defaults to 1 second. Minimum value is 1.
More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes`,
																							"format": "int32",
																							"type":   "integer",
																						},
																					},
																					"type": "object",
																				},
																				"name": map[string]interface{}{
																					"description": `Name of the container specified as a DNS_LABEL.
Each container in a pod must have a unique name (DNS_LABEL).
Cannot be updated.`,
																					"type": "string",
																				},
																				"ports": map[string]interface{}{
																					"description": `List of ports to expose from the container. Not specifying a port here
DOES NOT prevent that port from being exposed. Any port which is
listening on the default "0.0.0.0" address inside a container will be
accessible from the network.
Modifying this array with strategic merge patch may corrupt the data.
For more information See https://github.com/kubernetes/kubernetes/issues/108255.
Cannot be updated.`,
																					"items": map[string]interface{}{
																						"description": "ContainerPort represents a network port in a single container.",
																						"properties": map[string]interface{}{
																							"containerPort": map[string]interface{}{
																								"description": `Number of port to expose on the pod's IP address.
This must be a valid port number, 0 < x < 65536.`,
																								"format": "int32",
																								"type":   "integer",
																							},
																							"hostIP": map[string]interface{}{
																								"description": "What host IP to bind the external port to.",
																								"type":        "string",
																							},
																							"hostPort": map[string]interface{}{
																								"description": `Number of port to expose on the host.
If specified, this must be a valid port number, 0 < x < 65536.
If HostNetwork is specified, this must match ContainerPort.
Most containers do not need this.`,
																								"format": "int32",
																								"type":   "integer",
																							},
																							"name": map[string]interface{}{
																								"description": `If specified, this must be an IANA_SVC_NAME and unique within the pod. Each
named port in a pod must have a unique name. Name for the port that can be
referred to by services.`,
																								"type": "string",
																							},
																							"protocol": map[string]interface{}{
																								"default": "TCP",
																								"description": `Protocol for port. Must be UDP, TCP, or SCTP.
Defaults to "TCP".`,
																								"type": "string",
																							},
																						},
																						"required": []interface{}{
																							"containerPort",
																							"protocol",
																						},
																						"type": "object",
																					},
																					"type": "array",
																					"x-kubernetes-list-map-keys": []interface{}{
																						"containerPort",
																						"protocol",
																					},
																					"x-kubernetes-list-type": "map",
																				},
																				"readinessProbe": map[string]interface{}{
																					"description": `Periodic probe of container service readiness.
Container will be removed from service endpoints if the probe fails.
Cannot be updated.
More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes`,
																					"properties": map[string]interface{}{
																						"exec": map[string]interface{}{
																							"description": "Exec specifies a command to execute in the container.",
																							"properties": map[string]interface{}{
																								"command": map[string]interface{}{
																									"description": `Command is the command line to execute inside the container, the working directory for the
command  is root ('/') in the container's filesystem. The command is simply exec'd, it is
not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use
a shell, you need to explicitly call out to that shell.
Exit status of 0 is treated as live/healthy and non-zero is unhealthy.`,
																									"items": map[string]interface{}{
																										"type": "string",
																									},
																									"type":                   "array",
																									"x-kubernetes-list-type": "atomic",
																								},
																							},
																							"type": "object",
																						},
																						"failureThreshold": map[string]interface{}{
																							"description": `Minimum consecutive failures for the probe to be considered failed after having succeeded.
Defaults to 3. Minimum value is 1.`,
																							"format": "int32",
																							"type":   "integer",
																						},
																						"grpc": map[string]interface{}{
																							"description": "GRPC specifies a GRPC HealthCheckRequest.",
																							"properties": map[string]interface{}{
																								"port": map[string]interface{}{
																									"description": "Port number of the gRPC service. Number must be in the range 1 to 65535.",
																									"format":      "int32",
																									"type":        "integer",
																								},
																								"service": map[string]interface{}{
																									"default": "",
																									"description": `Service is the name of the service to place in the gRPC HealthCheckRequest
(see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).

If this is not specified, the default behavior is defined by gRPC.`,
																									"type": "string",
																								},
																							},
																							"required": []interface{}{
																								"port",
																							},
																							"type": "object",
																						},
																						"httpGet": map[string]interface{}{
																							"description": "HTTPGet specifies an HTTP GET request to perform.",
																							"properties": map[string]interface{}{
																								"host": map[string]interface{}{
																									"description": `Host name to connect to, defaults to the pod IP. You probably want to set
"Host" in httpHeaders instead.`,
																									"type": "string",
																								},
																								"httpHeaders": map[string]interface{}{
																									"description": "Custom headers to set in the request. HTTP allows repeated headers.",
																									"items": map[string]interface{}{
																										"description": "HTTPHeader describes a custom header to be used in HTTP probes",
																										"properties": map[string]interface{}{
																											"name": map[string]interface{}{
																												"description": `The header field name.
This will be canonicalized upon output, so case-variant names will be understood as the same header.`,
																												"type": "string",
																											},
																											"value": map[string]interface{}{
																												"description": "The header field value",
																												"type":        "string",
																											},
																										},
																										"required": []interface{}{
																											"name",
																											"value",
																										},
																										"type": "object",
																									},
																									"type":                   "array",
																									"x-kubernetes-list-type": "atomic",
																								},
																								"path": map[string]interface{}{
																									"description": "Path to access on the HTTP server.",
																									"type":        "string",
																								},
																								"port": map[string]interface{}{
																									"anyOf": []interface{}{
																										map[string]interface{}{
																											"type": "integer",
																										},
																										map[string]interface{}{
																											"type": "string",
																										},
																									},
																									"description": `Name or number of the port to access on the container.
Number must be in the range 1 to 65535.
Name must be an IANA_SVC_NAME.`,
																									"x-kubernetes-int-or-string": true,
																								},
																								"scheme": map[string]interface{}{
																									"description": `Scheme to use for connecting to the host.
Defaults to HTTP.`,
																									"type": "string",
																								},
																							},
																							"required": []interface{}{
																								"port",
																							},
																							"type": "object",
																						},
																						"initialDelaySeconds": map[string]interface{}{
																							"description": `Number of seconds after the container has started before liveness probes are initiated.
More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes`,
																							"format": "int32",
																							"type":   "integer",
																						},
																						"periodSeconds": map[string]interface{}{
																							"description": `How often (in seconds) to perform the probe.
Default to 10 seconds. Minimum value is 1.`,
																							"format": "int32",
																							"type":   "integer",
																						},
																						"successThreshold": map[string]interface{}{
																							"description": `Minimum consecutive successes for the probe to be considered successful after having failed.
Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.`,
																							"format": "int32",
																							"type":   "integer",
																						},
																						"tcpSocket": map[string]interface{}{
																							"description": "TCPSocket specifies a connection to a TCP port.",
																							"properties": map[string]interface{}{
																								"host": map[string]interface{}{
																									"description": "Optional: Host name to connect to, defaults to the pod IP.",
																									"type":        "string",
																								},
																								"port": map[string]interface{}{
																									"anyOf": []interface{}{
																										map[string]interface{}{
																											"type": "integer",
																										},
																										map[string]interface{}{
																											"type": "string",
																										},
																									},
																									"description": `Number or name of the port to access on the container.
Number must be in the range 1 to 65535.
Name must be an IANA_SVC_NAME.`,
																									"x-kubernetes-int-or-string": true,
																								},
																							},
																							"required": []interface{}{
																								"port",
																							},
																							"type": "object",
																						},
																						"terminationGracePeriodSeconds": map[string]interface{}{
																							"description": `Optional duration in seconds the pod needs to terminate gracefully upon probe failure.
The grace period is the duration in seconds after the processes running in the pod are sent
a termination signal and the time when the processes are forcibly halted with a kill signal.
Set this value longer than the expected cleanup time for your process.
If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this
value overrides the value provided by the pod spec.
Value must be non-negative integer. The value zero indicates stop immediately via
the kill signal (no opportunity to shut down).
This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate.
Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.`,
																							"format": "int64",
																							"type":   "integer",
																						},
																						"timeoutSeconds": map[string]interface{}{
																							"description": `Number of seconds after which the probe times out.
Defaults to 1 second. Minimum value is 1.
More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes`,
																							"format": "int32",
																							"type":   "integer",
																						},
																					},
																					"type": "object",
																				},
																				"resizePolicy": map[string]interface{}{
																					"description": `Resources resize policy for the container.
This field cannot be set on ephemeral containers.`,
																					"items": map[string]interface{}{
																						"description": "ContainerResizePolicy represents resource resize policy for the container.",
																						"properties": map[string]interface{}{
																							"resourceName": map[string]interface{}{
																								"description": `Name of the resource to which this resource resize policy applies.
Supported values: cpu, memory.`,
																								"type": "string",
																							},
																							"restartPolicy": map[string]interface{}{
																								"description": `Restart policy to apply when specified resource is resized.
If not specified, it defaults to NotRequired.`,
																								"type": "string",
																							},
																						},
																						"required": []interface{}{
																							"resourceName",
																							"restartPolicy",
																						},
																						"type": "object",
																					},
																					"type":                   "array",
																					"x-kubernetes-list-type": "atomic",
																				},
																				"resources": map[string]interface{}{
																					"description": `Compute Resources required by this container.
Cannot be updated.
More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/`,
																					"properties": map[string]interface{}{
																						"claims": map[string]interface{}{
																							"description": `Claims lists the names of resources, defined in spec.resourceClaims,
that are used by this container.

This field depends on the
DynamicResourceAllocation feature gate.

This field is immutable. It can only be set for containers.`,
																							"items": map[string]interface{}{
																								"description": "ResourceClaim references one entry in PodSpec.ResourceClaims.",
																								"properties": map[string]interface{}{
																									"name": map[string]interface{}{
																										"description": `Name must match the name of one entry in pod.spec.resourceClaims of
the Pod where this field is used. It makes that resource available
inside a container.`,
																										"type": "string",
																									},
																									"request": map[string]interface{}{
																										"description": `Request is the name chosen for a request in the referenced claim.
If empty, everything from the claim is made available, otherwise
only the result of this request.`,
																										"type": "string",
																									},
																								},
																								"required": []interface{}{
																									"name",
																								},
																								"type": "object",
																							},
																							"type": "array",
																							"x-kubernetes-list-map-keys": []interface{}{
																								"name",
																							},
																							"x-kubernetes-list-type": "map",
																						},
																						"limits": map[string]interface{}{
																							"additionalProperties": map[string]interface{}{
																								"anyOf": []interface{}{
																									map[string]interface{}{
																										"type": "integer",
																									},
																									map[string]interface{}{
																										"type": "string",
																									},
																								},
																								"pattern":                    `^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$`,
																								"x-kubernetes-int-or-string": true,
																							},
																							"description": `Limits describes the maximum amount of compute resources allowed.
More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/`,
																							"type": "object",
																						},
																						"requests": map[string]interface{}{
																							"additionalProperties": map[string]interface{}{
																								"anyOf": []interface{}{
																									map[string]interface{}{
																										"type": "integer",
																									},
																									map[string]interface{}{
																										"type": "string",
																									},
																								},
																								"pattern":                    `^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$`,
																								"x-kubernetes-int-or-string": true,
																							},
																							"description": `Requests describes the minimum amount of compute resources required.
If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,
otherwise to an implementation-defined value. Requests cannot exceed Limits.
More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/`,
																							"type": "object",
																						},
																					},
																					"type": "object",
																				},
																				"restartPolicy": map[string]interface{}{
																					"description": `RestartPolicy defines the restart behavior of individual containers in a pod.
This overrides the pod-level restart policy. When this field is not specified,
the restart behavior is defined by the Pod's restart policy and the container type.
Additionally, setting the RestartPolicy as "Always" for the init container will
have the following effect:
this init container will be continually restarted on
exit until all regular containers have terminated. Once all regular
containers have completed, all init containers with restartPolicy "Always"
will be shut down. This lifecycle differs from normal init containers and
is often referred to as a "sidecar" container. Although this init
container still starts in the init container sequence, it does not wait
for the container to complete before proceeding to the next init
container. Instead, the next init container starts immediately after this
init container is started, or after any startupProbe has successfully
completed.`,
																					"type": "string",
																				},
																				"restartPolicyRules": map[string]interface{}{
																					"description": `Represents a list of rules to be checked to determine if the
container should be restarted on exit. The rules are evaluated in
order. Once a rule matches a container exit condition, the remaining
rules are ignored. If no rule matches the container exit condition,
the Container-level restart policy determines the whether the container
is restarted or not. Constraints on the rules:
- At most 20 rules are allowed.
- Rules can have the same action.
- Identical rules are not forbidden in validations.
When rules are specified, container MUST set RestartPolicy explicitly
even it if matches the Pod's RestartPolicy.`,
																					"items": map[string]interface{}{
																						"description": "ContainerRestartRule describes how a container exit is handled.",
																						"properties": map[string]interface{}{
																							"action": map[string]interface{}{
																								"description": `Specifies the action taken on a container exit if the requirements
are satisfied. The only possible value is "Restart" to restart the
container.`,
																								"type": "string",
																							},
																							"exitCodes": map[string]interface{}{
																								"description": "Represents the exit codes to check on container exits.",
																								"properties": map[string]interface{}{
																									"operator": map[string]interface{}{
																										"description": `Represents the relationship between the container exit code(s) and the
specified values. Possible values are:
- In: the requirement is satisfied if the container exit code is in the
  set of specified values.
- NotIn: the requirement is satisfied if the container exit code is
  not in the set of specified values.`,
																										"type": "string",
																									},
																									"values": map[string]interface{}{
																										"description": `Specifies the set of values to check for container exit codes.
At most 255 elements are allowed.`,
																										"items": map[string]interface{}{
																											"format": "int32",
																											"type":   "integer",
																										},
																										"type":                   "array",
																										"x-kubernetes-list-type": "set",
																									},
																								},
																								"required": []interface{}{
																									"operator",
																								},
																								"type": "object",
																							},
																						},
																						"required": []interface{}{
																							"action",
																						},
																						"type": "object",
																					},
																					"type":                   "array",
																					"x-kubernetes-list-type": "atomic",
																				},
																				"securityContext": map[string]interface{}{
																					"description": `SecurityContext defines the security options the container should be run with.
If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext.
More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/`,
																					"properties": map[string]interface{}{
																						"allowPrivilegeEscalation": map[string]interface{}{
																							"description": `AllowPrivilegeEscalation controls whether a process can gain more
privileges than its parent process. This bool directly controls if
the no_new_privs flag will be set on the container process.
AllowPrivilegeEscalation is true always when the container is:
1) run as Privileged
2) has CAP_SYS_ADMIN
Note that this field cannot be set when spec.os.name is windows.`,
																							"type": "boolean",
																						},
																						"appArmorProfile": map[string]interface{}{
																							"description": `appArmorProfile is the AppArmor options to use by this container. If set, this profile
overrides the pod's appArmorProfile.
Note that this field cannot be set when spec.os.name is windows.`,
																							"properties": map[string]interface{}{
																								"localhostProfile": map[string]interface{}{
																									"description": `localhostProfile indicates a profile loaded on the node that should be used.
The profile must be preconfigured on the node to work.
Must match the loaded name of the profile.
Must be set if and only if type is "Localhost".`,
																									"type": "string",
																								},
																								"type": map[string]interface{}{
																									"description": `type indicates which kind of AppArmor profile will be applied.
Valid options are:
  Localhost - a profile pre-loaded on the node.
  RuntimeDefault - the container runtime's default profile.
  Unconfined - no AppArmor enforcement.`,
																									"type": "string",
																								},
																							},
																							"required": []interface{}{
																								"type",
																							},
																							"type": "object",
																						},
																						"capabilities": map[string]interface{}{
																							"description": `The capabilities to add/drop when running containers.
Defaults to the default set of capabilities granted by the container runtime.
Note that this field cannot be set when spec.os.name is windows.`,
																							"properties": map[string]interface{}{
																								"add": map[string]interface{}{
																									"description": "Added capabilities",
																									"items": map[string]interface{}{
																										"description": "Capability represent POSIX capabilities type",
																										"type":        "string",
																									},
																									"type":                   "array",
																									"x-kubernetes-list-type": "atomic",
																								},
																								"drop": map[string]interface{}{
																									"description": "Removed capabilities",
																									"items": map[string]interface{}{
																										"description": "Capability represent POSIX capabilities type",
																										"type":        "string",
																									},
																									"type":                   "array",
																									"x-kubernetes-list-type": "atomic",
																								},
																							},
																							"type": "object",
																						},
																						"privileged": map[string]interface{}{
																							"description": `Run container in privileged mode.
Processes in privileged containers are essentially equivalent to root on the host.
Defaults to false.
Note that this field cannot be set when spec.os.name is windows.`,
																							"type": "boolean",
																						},
																						"procMount": map[string]interface{}{
																							"description": `procMount denotes the type of proc mount to use for the containers.
The default value is Default which uses the container runtime defaults for
readonly paths and masked paths.
Note that this field cannot be set when spec.os.name is windows.`,
																							"type": "string",
																						},
																						"readOnlyRootFilesystem": map[string]interface{}{
																							"description": `Whether this container has a read-only root filesystem.
Default is false.
Note that this field cannot be set when spec.os.name is windows.`,
																							"type": "boolean",
																						},
																						"runAsGroup": map[string]interface{}{
																							"description": `The GID to run the entrypoint of the container process.
Uses runtime default if unset.
May also be set in PodSecurityContext.  If set in both SecurityContext and
PodSecurityContext, the value specified in SecurityContext takes precedence.
Note that this field cannot be set when spec.os.name is windows.`,
																							"format": "int64",
																							"type":   "integer",
																						},
																						"runAsNonRoot": map[string]interface{}{
																							"description": `Indicates that the container must run as a non-root user.
If true, the Kubelet will validate the image at runtime to ensure that it
does not run as UID 0 (root) and fail to start the container if it does.
If unset or false, no such validation will be performed.
May also be set in PodSecurityContext.  If set in both SecurityContext and
PodSecurityContext, the value specified in SecurityContext takes precedence.`,
																							"type": "boolean",
																						},
																						"runAsUser": map[string]interface{}{
																							"description": `The UID to run the entrypoint of the container process.
Defaults to user specified in image metadata if unspecified.
May also be set in PodSecurityContext.  If set in both SecurityContext and
PodSecurityContext, the value specified in SecurityContext takes precedence.
Note that this field cannot be set when spec.os.name is windows.`,
																							"format": "int64",
																							"type":   "integer",
																						},
																						"seLinuxOptions": map[string]interface{}{
																							"description": `The SELinux context to be applied to the container.
If unspecified, the container runtime will allocate a random SELinux context for each
container.  May also be set in PodSecurityContext.  If set in both SecurityContext and
PodSecurityContext, the value specified in SecurityContext takes precedence.
Note that this field cannot be set when spec.os.name is windows.`,
																							"properties": map[string]interface{}{
																								"level": map[string]interface{}{
																									"description": "Level is SELinux level label that applies to the container.",
																									"type":        "string",
																								},
																								"role": map[string]interface{}{
																									"description": "Role is a SELinux role label that applies to the container.",
																									"type":        "string",
																								},
																								"type": map[string]interface{}{
																									"description": "Type is a SELinux type label that applies to the container.",
																									"type":        "string",
																								},
																								"user": map[string]interface{}{
																									"description": "User is a SELinux user label that applies to the container.",
																									"type":        "string",
																								},
																							},
																							"type": "object",
																						},
																						"seccompProfile": map[string]interface{}{
																							"description": `The seccomp options to use by this container. If seccomp options are
provided at both the pod & container level, the container options
override the pod options.
Note that this field cannot be set when spec.os.name is windows.`,
																							"properties": map[string]interface{}{
																								"localhostProfile": map[string]interface{}{
																									"description": `localhostProfile indicates a profile defined in a file on the node should be used.
The profile must be preconfigured on the node to work.
Must be a descending path, relative to the kubelet's configured seccomp profile location.
Must be set if type is "Localhost". Must NOT be set for any other type.`,
																									"type": "string",
																								},
																								"type": map[string]interface{}{
																									"description": `type indicates which kind of seccomp profile will be applied.
Valid options are:

Localhost - a profile defined in a file on the node should be used.
RuntimeDefault - the container runtime default profile should be used.
Unconfined - no profile should be applied.`,
																									"type": "string",
																								},
																							},
																							"required": []interface{}{
																								"type",
																							},
																							"type": "object",
																						},
																						"windowsOptions": map[string]interface{}{
																							"description": `The Windows specific settings applied to all containers.
If unspecified, the options from the PodSecurityContext will be used.
If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
Note that this field cannot be set when spec.os.name is linux.`,
																							"properties": map[string]interface{}{
																								"gmsaCredentialSpec": map[string]interface{}{
																									"description": `GMSACredentialSpec is where the GMSA admission webhook
(https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the
GMSA credential spec named by the GMSACredentialSpecName field.`,
																									"type": "string",
																								},
																								"gmsaCredentialSpecName": map[string]interface{}{
																									"description": "GMSACredentialSpecName is the name of the GMSA credential spec to use.",
																									"type":        "string",
																								},
																								"hostProcess": map[string]interface{}{
																									"description": `HostProcess determines if a container should be run as a 'Host Process' container.
All of a Pod's containers must have the same effective HostProcess value
(it is not allowed to have a mix of HostProcess containers and non-HostProcess containers).
In addition, if HostProcess is true then HostNetwork must also be set to true.`,
																									"type": "boolean",
																								},
																								"runAsUserName": map[string]interface{}{
																									"description": `The UserName in Windows to run the entrypoint of the container process.
Defaults to the user specified in image metadata if unspecified.
May also be set in PodSecurityContext. If set in both SecurityContext and
PodSecurityContext, the value specified in SecurityContext takes precedence.`,
																									"type": "string",
																								},
																							},
																							"type": "object",
																						},
																					},
																					"type": "object",
																				},
																				"startupProbe": map[string]interface{}{
																					"description": `StartupProbe indicates that the Pod has successfully initialized.
If specified, no other probes are executed until this completes successfully.
If this probe fails, the Pod will be restarted, just as if the livenessProbe failed.
This can be used to provide different probe parameters at the beginning of a Pod's lifecycle,
when it might take a long time to load data or warm a cache, than during steady-state operation.
This cannot be updated.
More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes`,
																					"properties": map[string]interface{}{
																						"exec": map[string]interface{}{
																							"description": "Exec specifies a command to execute in the container.",
																							"properties": map[string]interface{}{
																								"command": map[string]interface{}{
																									"description": `Command is the command line to execute inside the container, the working directory for the
command  is root ('/') in the container's filesystem. The command is simply exec'd, it is
not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use
a shell, you need to explicitly call out to that shell.
Exit status of 0 is treated as live/healthy and non-zero is unhealthy.`,
																									"items": map[string]interface{}{
																										"type": "string",
																									},
																									"type":                   "array",
																									"x-kubernetes-list-type": "atomic",
																								},
																							},
																							"type": "object",
																						},
																						"failureThreshold": map[string]interface{}{
																							"description": `Minimum consecutive failures for the probe to be considered failed after having succeeded.
Defaults to 3. Minimum value is 1.`,
																							"format": "int32",
																							"type":   "integer",
																						},
																						"grpc": map[string]interface{}{
																							"description": "GRPC specifies a GRPC HealthCheckRequest.",
																							"properties": map[string]interface{}{
																								"port": map[string]interface{}{
																									"description": "Port number of the gRPC service. Number must be in the range 1 to 65535.",
																									"format":      "int32",
																									"type":        "integer",
																								},
																								"service": map[string]interface{}{
																									"default": "",
																									"description": `Service is the name of the service to place in the gRPC HealthCheckRequest
(see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).

If this is not specified, the default behavior is defined by gRPC.`,
																									"type": "string",
																								},
																							},
																							"required": []interface{}{
																								"port",
																							},
																							"type": "object",
																						},
																						"httpGet": map[string]interface{}{
																							"description": "HTTPGet specifies an HTTP GET request to perform.",
																							"properties": map[string]interface{}{
																								"host": map[string]interface{}{
																									"description": `Host name to connect to, defaults to the pod IP. You probably want to set
"Host" in httpHeaders instead.`,
																									"type": "string",
																								},
																								"httpHeaders": map[string]interface{}{
																									"description": "Custom headers to set in the request. HTTP allows repeated headers.",
																									"items": map[string]interface{}{
																										"description": "HTTPHeader describes a custom header to be used in HTTP probes",
																										"properties": map[string]interface{}{
																											"name": map[string]interface{}{
																												"description": `The header field name.
This will be canonicalized upon output, so case-variant names will be understood as the same header.`,
																												"type": "string",
																											},
																											"value": map[string]interface{}{
																												"description": "The header field value",
																												"type":        "string",
																											},
																										},
																										"required": []interface{}{
																											"name",
																											"value",
																										},
																										"type": "object",
																									},
																									"type":                   "array",
																									"x-kubernetes-list-type": "atomic",
																								},
																								"path": map[string]interface{}{
																									"description": "Path to access on the HTTP server.",
																									"type":        "string",
																								},
																								"port": map[string]interface{}{
																									"anyOf": []interface{}{
																										map[string]interface{}{
																											"type": "integer",
																										},
																										map[string]interface{}{
																											"type": "string",
																										},
																									},
																									"description": `Name or number of the port to access on the container.
Number must be in the range 1 to 65535.
Name must be an IANA_SVC_NAME.`,
																									"x-kubernetes-int-or-string": true,
																								},
																								"scheme": map[string]interface{}{
																									"description": `Scheme to use for connecting to the host.
Defaults to HTTP.`,
																									"type": "string",
																								},
																							},
																							"required": []interface{}{
																								"port",
																							},
																							"type": "object",
																						},
																						"initialDelaySeconds": map[string]interface{}{
																							"description": `Number of seconds after the container has started before liveness probes are initiated.
More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes`,
																							"format": "int32",
																							"type":   "integer",
																						},
																						"periodSeconds": map[string]interface{}{
																							"description": `How often (in seconds) to perform the probe.
Default to 10 seconds. Minimum value is 1.`,
																							"format": "int32",
																							"type":   "integer",
																						},
																						"successThreshold": map[string]interface{}{
																							"description": `Minimum consecutive successes for the probe to be considered successful after having failed.
Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.`,
																							"format": "int32",
																							"type":   "integer",
																						},
																						"tcpSocket": map[string]interface{}{
																							"description": "TCPSocket specifies a connection to a TCP port.",
																							"properties": map[string]interface{}{
																								"host": map[string]interface{}{
																									"description": "Optional: Host name to connect to, defaults to the pod IP.",
																									"type":        "string",
																								},
																								"port": map[string]interface{}{
																									"anyOf": []interface{}{
																										map[string]interface{}{
																											"type": "integer",
																										},
																										map[string]interface{}{
																											"type": "string",
																										},
																									},
																									"description": `Number or name of the port to access on the container.
Number must be in the range 1 to 65535.
Name must be an IANA_SVC_NAME.`,
																									"x-kubernetes-int-or-string": true,
																								},
																							},
																							"required": []interface{}{
																								"port",
																							},
																							"type": "object",
																						},
																						"terminationGracePeriodSeconds": map[string]interface{}{
																							"description": `Optional duration in seconds the pod needs to terminate gracefully upon probe failure.
The grace period is the duration in seconds after the processes running in the pod are sent
a termination signal and the time when the processes are forcibly halted with a kill signal.
Set this value longer than the expected cleanup time for your process.
If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this
value overrides the value provided by the pod spec.
Value must be non-negative integer. The value zero indicates stop immediately via
the kill signal (no opportunity to shut down).
This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate.
Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.`,
																							"format": "int64",
																							"type":   "integer",
																						},
																						"timeoutSeconds": map[string]interface{}{
																							"description": `Number of seconds after which the probe times out.
Defaults to 1 second. Minimum value is 1.
More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes`,
																							"format": "int32",
																							"type":   "integer",
																						},
																					},
																					"type": "object",
																				},
																				"stdin": map[string]interface{}{
																					"description": `Whether this container should allocate a buffer for stdin in the container runtime. If this
is not set, reads from stdin in the container will always result in EOF.
Default is false.`,
																					"type": "boolean",
																				},
																				"stdinOnce": map[string]interface{}{
																					"description": `Whether the container runtime should close the stdin channel after it has been opened by
a single attach. When stdin is true the stdin stream will remain open across multiple attach
sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the
first client attaches to stdin, and then remains open and accepts data until the client disconnects,
at which time stdin is closed and remains closed until the container is restarted. If this
flag is false, a container processes that reads from stdin will never receive an EOF.
Default is false`,
																					"type": "boolean",
																				},
																				"terminationMessagePath": map[string]interface{}{
																					"description": `Optional: Path at which the file to which the container's termination message
will be written is mounted into the container's filesystem.
Message written is intended to be brief final status, such as an assertion failure message.
Will be truncated by the node if greater than 4096 bytes. The total message length across
all containers will be limited to 12kb.
Defaults to /dev/termination-log.
Cannot be updated.`,
																					"type": "string",
																				},
																				"terminationMessagePolicy": map[string]interface{}{
																					"description": `Indicate how the termination message should be populated. File will use the contents of
terminationMessagePath to populate the container status message on both success and failure.
FallbackToLogsOnError will use the last chunk of container log output if the termination
message file is empty and the container exited with an error.
The log output is limited to 2048 bytes or 80 lines, whichever is smaller.
Defaults to File.
Cannot be updated.`,
																					"type": "string",
																				},
																				"tty": map[string]interface{}{
																					"description": `Whether this container should allocate a TTY for itself, also requires 'stdin' to be true.
Default is false.`,
																					"type": "boolean",
																				},
																				"volumeDevices": map[string]interface{}{
																					"description": "volumeDevices is the list of block devices to be used by the container.",
																					"items": map[string]interface{}{
																						"description": "volumeDevice describes a mapping of a raw block device within a container.",
																						"properties": map[string]interface{}{
																							"devicePath": map[string]interface{}{
																								"description": "devicePath is the path inside of the container that the device will be mapped to.",
																								"type":        "string",
																							},
																							"name": map[string]interface{}{
																								"description": "name must match the name of a persistentVolumeClaim in the pod",
																								"type":        "string",
																							},
																						},
																						"required": []interface{}{
																							"devicePath",
																							"name",
																						},
																						"type": "object",
																					},
																					"type": "array",
																					"x-kubernetes-list-map-keys": []interface{}{
																						"devicePath",
																					},
																					"x-kubernetes-list-type": "map",
																				},
																				"volumeMounts": map[string]interface{}{
																					"description": `Pod volumes to mount into the container's filesystem.
Cannot be updated.`,
																					"items": map[string]interface{}{
																						"description": "VolumeMount describes a mounting of a Volume within a container.",
																						"properties": map[string]interface{}{
																							"mountPath": map[string]interface{}{
																								"description": `Path within the container at which the volume should be mounted.  Must
not contain ':'.`,
																								"type": "string",
																							},
																							"mountPropagation": map[string]interface{}{
																								"description": `mountPropagation determines how mounts are propagated from the host
to container and the other way around.
When not set, MountPropagationNone is used.
This field is beta in 1.10.
When RecursiveReadOnly is set to IfPossible or to Enabled, MountPropagation must be None or unspecified
(which defaults to None).`,
																								"type": "string",
																							},
																							"name": map[string]interface{}{
																								"description": "This must match the Name of a Volume.",
																								"type":        "string",
																							},
																							"readOnly": map[string]interface{}{
																								"description": `Mounted read-only if true, read-write otherwise (false or unspecified).
Defaults to false.`,
																								"type": "boolean",
																							},
																							"recursiveReadOnly": map[string]interface{}{
																								"description": `RecursiveReadOnly specifies whether read-only mounts should be handled
recursively.

If ReadOnly is false, this field has no meaning and must be unspecified.

If ReadOnly is true, and this field is set to Disabled, the mount is not made
recursively read-only.  If this field is set to IfPossible, the mount is made
recursively read-only, if it is supported by the container runtime.  If this
field is set to Enabled, the mount is made recursively read-only if it is
supported by the container runtime, otherwise the pod will not be started and
an error will be generated to indicate the reason.

If this field is set to IfPossible or Enabled, MountPropagation must be set to
None (or be unspecified, which defaults to None).

If this field is not specified, it is treated as an equivalent of Disabled.`,
																								"type": "string",
																							},
																							"subPath": map[string]interface{}{
																								"description": `Path within the volume from which the container's volume should be mounted.
Defaults to "" (volume's root).`,
																								"type": "string",
																							},
																							"subPathExpr": map[string]interface{}{
																								"description": `Expanded path within the volume from which the container's volume should be mounted.
Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment.
Defaults to "" (volume's root).
SubPathExpr and SubPath are mutually exclusive.`,
																								"type": "string",
																							},
																						},
																						"required": []interface{}{
																							"mountPath",
																							"name",
																						},
																						"type": "object",
																					},
																					"type": "array",
																					"x-kubernetes-list-map-keys": []interface{}{
																						"mountPath",
																					},
																					"x-kubernetes-list-type": "map",
																				},
																				"workingDir": map[string]interface{}{
																					"description": `Container's working directory.
If not specified, the container runtime's default will be used, which
might be configured in the container image.
Cannot be updated.`,
																					"type": "string",
																				},
																			},
																			"required": []interface{}{
																				"name",
																			},
																			"type": "object",
																		},
																		"type": "array",
																		"x-kubernetes-list-map-keys": []interface{}{
																			"name",
																		},
																		"x-kubernetes-list-type": "map",
																	},
																	"nodeName": map[string]interface{}{
																		"description": `NodeName indicates in which node this pod is scheduled.
If empty, this pod is a candidate for scheduling by the scheduler defined in schedulerName.
Once this field is set, the kubelet for this node becomes responsible for the lifecycle of this pod.
This field should not be used to express a desire for the pod to be scheduled on a specific node.
https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#nodename`,
																		"type": "string",
																	},
																	"nodeSelector": map[string]interface{}{
																		"additionalProperties": map[string]interface{}{
																			"type": "string",
																		},
																		"description": `NodeSelector is a selector which must be true for the pod to fit on a node.
Selector which must match a node's labels for the pod to be scheduled on that node.
More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/`,
																		"type":                  "object",
																		"x-kubernetes-map-type": "atomic",
																	},
																	"os": map[string]interface{}{
																		"description": `Specifies the OS of the containers in the pod.
Some pod and container fields are restricted if this is set.

If the OS field is set to linux, the following fields must be unset:
-securityContext.windowsOptions

If the OS field is set to windows, following fields must be unset:
- spec.hostPID
- spec.hostIPC
- spec.hostUsers
- spec.resources
- spec.securityContext.appArmorProfile
- spec.securityContext.seLinuxOptions
- spec.securityContext.seccompProfile
- spec.securityContext.fsGroup
- spec.securityContext.fsGroupChangePolicy
- spec.securityContext.sysctls
- spec.shareProcessNamespace
- spec.securityContext.runAsUser
- spec.securityContext.runAsGroup
- spec.securityContext.supplementalGroups
- spec.securityContext.supplementalGroupsPolicy
- spec.containers[*].securityContext.appArmorProfile
- spec.containers[*].securityContext.seLinuxOptions
- spec.containers[*].securityContext.seccompProfile
- spec.containers[*].securityContext.capabilities
- spec.containers[*].securityContext.readOnlyRootFilesystem
- spec.containers[*].securityContext.privileged
- spec.containers[*].securityContext.allowPrivilegeEscalation
- spec.containers[*].securityContext.procMount
- spec.containers[*].securityContext.runAsUser
- spec.containers[*].securityContext.runAsGroup`,
																		"properties": map[string]interface{}{
																			"name": map[string]interface{}{
																				"description": `Name is the name of the operating system. The currently supported values are linux and windows.
Additional value may be defined in future and can be one of:
https://github.com/opencontainers/runtime-spec/blob/master/config.md#platform-specific-configuration
Clients should expect to handle additional values and treat unrecognized values in this field as os: null`,
																				"type": "string",
																			},
																		},
																		"required": []interface{}{
																			"name",
																		},
																		"type": "object",
																	},
																	"overhead": map[string]interface{}{
																		"additionalProperties": map[string]interface{}{
																			"anyOf": []interface{}{
																				map[string]interface{}{
																					"type": "integer",
																				},
																				map[string]interface{}{
																					"type": "string",
																				},
																			},
																			"pattern":                    `^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$`,
																			"x-kubernetes-int-or-string": true,
																		},
																		"description": `Overhead represents the resource overhead associated with running a pod for a given RuntimeClass.
This field will be autopopulated at admission time by the RuntimeClass admission controller. If
the RuntimeClass admission controller is enabled, overhead must not be set in Pod create requests.
The RuntimeClass admission controller will reject Pod create requests which have the overhead already
set. If RuntimeClass is configured and selected in the PodSpec, Overhead will be set to the value
defined in the corresponding RuntimeClass, otherwise it will remain unset and treated as zero.
More info: https://git.k8s.io/enhancements/keps/sig-node/688-pod-overhead/README.md`,
																		"type": "object",
																	},
																	"preemptionPolicy": map[string]interface{}{
																		"description": `PreemptionPolicy is the Policy for preempting pods with lower priority.
One of Never, PreemptLowerPriority.
Defaults to PreemptLowerPriority if unset.`,
																		"type": "string",
																	},
																	"priority": map[string]interface{}{
																		"description": `The priority value. Various system components use this field to find the
priority of the pod. When Priority Admission Controller is enabled, it
prevents users from setting this field. The admission controller populates
this field from PriorityClassName.
The higher the value, the higher the priority.`,
																		"format": "int32",
																		"type":   "integer",
																	},
																	"priorityClassName": map[string]interface{}{
																		"description": `If specified, indicates the pod's priority. "system-node-critical" and
"system-cluster-critical" are two special keywords which indicate the
highest priorities with the former being the highest priority. Any other
name must be defined by creating a PriorityClass object with that name.
If not specified, the pod priority will be default or zero if there is no
default.`,
																		"type": "string",
																	},
																	"readinessGates": map[string]interface{}{
																		"description": `If specified, all readiness gates will be evaluated for pod readiness.
A pod is ready when all its containers are ready AND
all conditions specified in the readiness gates have status equal to "True"
More info: https://git.k8s.io/enhancements/keps/sig-network/580-pod-readiness-gates`,
																		"items": map[string]interface{}{
																			"description": "PodReadinessGate contains the reference to a pod condition",
																			"properties": map[string]interface{}{
																				"conditionType": map[string]interface{}{
																					"description": "ConditionType refers to a condition in the pod's condition list with matching type.",
																					"type":        "string",
																				},
																			},
																			"required": []interface{}{
																				"conditionType",
																			},
																			"type": "object",
																		},
																		"type":                   "array",
																		"x-kubernetes-list-type": "atomic",
																	},
																	"resourceClaims": map[string]interface{}{
																		"description": `ResourceClaims defines which ResourceClaims must be allocated
and reserved before the Pod is allowed to start. The resources
will be made available to those containers which consume them
by name.

This is a stable field but requires that the
DynamicResourceAllocation feature gate is enabled.

This field is immutable.`,
																		"items": map[string]interface{}{
																			"description": `PodResourceClaim references exactly one ResourceClaim, either directly
or by naming a ResourceClaimTemplate which is then turned into a ResourceClaim
for the pod.

It adds a name to it that uniquely identifies the ResourceClaim inside the Pod.
Containers that need access to the ResourceClaim reference it with this name.

When the DRAWorkloadResourceClaims feature gate is enabled and this Pod
belongs to a PodGroup, a PodResourceClaim is matched to a
PodGroupResourceClaim if all of their fields are equal (Name,
ResourceClaimName, and ResourceClaimTemplateName). A matched claim references
a single ResourceClaim shared across all Pods in the PodGroup, reserved for
the PodGroup in ResourceClaimStatus.ReservedFor rather than for individual
Pods.`,
																			"properties": map[string]interface{}{
																				"name": map[string]interface{}{
																					"description": `Name uniquely identifies this resource claim inside the pod.
This must be a DNS_LABEL.`,
																					"type": "string",
																				},
																				"resourceClaimName": map[string]interface{}{
																					"description": `ResourceClaimName is the name of a ResourceClaim object in the same
namespace as this pod.

Exactly one of ResourceClaimName and ResourceClaimTemplateName must
be set.`,
																					"type": "string",
																				},
																				"resourceClaimTemplateName": map[string]interface{}{
																					"description": `ResourceClaimTemplateName is the name of a ResourceClaimTemplate
object in the same namespace as this pod.

The template will be used to create a new ResourceClaim, which will
be bound to this pod. When this pod is deleted, the ResourceClaim
will also be deleted. The pod name and resource name, along with a
generated component, will be used to form a unique name for the
ResourceClaim, which will be recorded in pod.status.resourceClaimStatuses.

When the DRAWorkloadResourceClaims feature gate is enabled and the pod
belongs to a PodGroup that defines a PodGroupResourceClaim with the same
Name and ResourceClaimTemplateName, this PodResourceClaim resolves to the
ResourceClaim generated for the PodGroup. All pods in the group that
define an equivalent PodResourceClaim matching the
PodGroupResourceClaim's Name and ResourceClaimTemplateName share the same
generated ResourceClaim. ResourceClaims generated for a PodGroup are
owned by the PodGroup and their lifecycles are tied to the PodGroup
instead of any individual pod.

This field is immutable and no changes will be made to the
corresponding ResourceClaim by the control plane after creating the
ResourceClaim.

Exactly one of ResourceClaimName and ResourceClaimTemplateName must
be set.`,
																					"type": "string",
																				},
																			},
																			"required": []interface{}{
																				"name",
																			},
																			"type": "object",
																		},
																		"type": "array",
																		"x-kubernetes-list-map-keys": []interface{}{
																			"name",
																		},
																		"x-kubernetes-list-type": "map",
																	},
																	"resources": map[string]interface{}{
																		"description": `Resources is the total amount of CPU and Memory resources required by all
containers in the pod. It supports specifying Requests and Limits for
"cpu", "memory" and "hugepages-" resource names only. ResourceClaims are not supported.

This field enables fine-grained control over resource allocation for the
entire pod, allowing resource sharing among containers in a pod.

This is an alpha field and requires enabling the PodLevelResources feature
gate.`,
																		"properties": map[string]interface{}{
																			"claims": map[string]interface{}{
																				"description": `Claims lists the names of resources, defined in spec.resourceClaims,
that are used by this container.

This field depends on the
DynamicResourceAllocation feature gate.

This field is immutable. It can only be set for containers.`,
																				"items": map[string]interface{}{
																					"description": "ResourceClaim references one entry in PodSpec.ResourceClaims.",
																					"properties": map[string]interface{}{
																						"name": map[string]interface{}{
																							"description": `Name must match the name of one entry in pod.spec.resourceClaims of
the Pod where this field is used. It makes that resource available
inside a container.`,
																							"type": "string",
																						},
																						"request": map[string]interface{}{
																							"description": `Request is the name chosen for a request in the referenced claim.
If empty, everything from the claim is made available, otherwise
only the result of this request.`,
																							"type": "string",
																						},
																					},
																					"required": []interface{}{
																						"name",
																					},
																					"type": "object",
																				},
																				"type": "array",
																				"x-kubernetes-list-map-keys": []interface{}{
																					"name",
																				},
																				"x-kubernetes-list-type": "map",
																			},
																			"limits": map[string]interface{}{
																				"additionalProperties": map[string]interface{}{
																					"anyOf": []interface{}{
																						map[string]interface{}{
																							"type": "integer",
																						},
																						map[string]interface{}{
																							"type": "string",
																						},
																					},
																					"pattern":                    `^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$`,
																					"x-kubernetes-int-or-string": true,
																				},
																				"description": `Limits describes the maximum amount of compute resources allowed.
More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/`,
																				"type": "object",
																			},
																			"requests": map[string]interface{}{
																				"additionalProperties": map[string]interface{}{
																					"anyOf": []interface{}{
																						map[string]interface{}{
																							"type": "integer",
																						},
																						map[string]interface{}{
																							"type": "string",
																						},
																					},
																					"pattern":                    `^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$`,
																					"x-kubernetes-int-or-string": true,
																				},
																				"description": `Requests describes the minimum amount of compute resources required.
If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,
otherwise to an implementation-defined value. Requests cannot exceed Limits.
More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/`,
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																	"restartPolicy": map[string]interface{}{
																		"description": `Restart policy for all containers within the pod.
One of Always, OnFailure, Never. In some contexts, only a subset of those values may be permitted.
Default to Always.
More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy`,
																		"type": "string",
																	},
																	"runtimeClassName": map[string]interface{}{
																		"description": `RuntimeClassName refers to a RuntimeClass object in the node.k8s.io group, which should be used
to run this pod.  If no RuntimeClass resource matches the named class, the pod will not be run.
If unset or empty, the "legacy" RuntimeClass will be used, which is an implicit class with an
empty definition that uses the default runtime handler.
More info: https://git.k8s.io/enhancements/keps/sig-node/585-runtime-class`,
																		"type": "string",
																	},
																	"schedulerName": map[string]interface{}{
																		"description": `If specified, the pod will be dispatched by specified scheduler.
If not specified, the pod will be dispatched by default scheduler.`,
																		"type": "string",
																	},
																	"schedulingGates": map[string]interface{}{
																		"description": `SchedulingGates is an opaque list of values that if specified will block scheduling the pod.
If schedulingGates is not empty, the pod will stay in the SchedulingGated state and the
scheduler will not attempt to schedule the pod.

SchedulingGates can only be set at pod creation time, and be removed only afterwards.`,
																		"items": map[string]interface{}{
																			"description": "PodSchedulingGate is associated to a Pod to guard its scheduling.",
																			"properties": map[string]interface{}{
																				"name": map[string]interface{}{
																					"description": `Name of the scheduling gate.
Each scheduling gate must have a unique name field.`,
																					"type": "string",
																				},
																			},
																			"required": []interface{}{
																				"name",
																			},
																			"type": "object",
																		},
																		"type": "array",
																		"x-kubernetes-list-map-keys": []interface{}{
																			"name",
																		},
																		"x-kubernetes-list-type": "map",
																	},
																	"schedulingGroup": map[string]interface{}{
																		"description": `SchedulingGroup provides a reference to the immediate scheduling runtime
grouping object that this Pod belongs to.
This field is used by the scheduler to identify the group and apply the
correct group scheduling policies. The association with a group also
impacts other lifecycle aspects of a Pod that are relevant in a wider context
of scheduling like preemption, resource attachment, etc. If not specified,
the Pod is treated as a single unit in all of these aspects.
The group object referenced by this field may not exist at the time the
Pod is created.
This field is immutable, but a group object with the same name may be
recreated with different policies. Doing this during pod scheduling
may result in the placement not conforming to the expected policies.`,
																		"properties": map[string]interface{}{
																			"podGroupName": map[string]interface{}{
																				"description": `PodGroupName specifies the name of the standalone PodGroup object
that represents the runtime instance of this group.
Must be a DNS subdomain.`,
																				"type": "string",
																			},
																		},
																		"type": "object",
																	},
																	"securityContext": map[string]interface{}{
																		"description": `SecurityContext holds pod-level security attributes and common container settings.
Optional: Defaults to empty.  See type description for default values of each field.`,
																		"properties": map[string]interface{}{
																			"appArmorProfile": map[string]interface{}{
																				"description": `appArmorProfile is the AppArmor options to use by the containers in this pod.
Note that this field cannot be set when spec.os.name is windows.`,
																				"properties": map[string]interface{}{
																					"localhostProfile": map[string]interface{}{
																						"description": `localhostProfile indicates a profile loaded on the node that should be used.
The profile must be preconfigured on the node to work.
Must match the loaded name of the profile.
Must be set if and only if type is "Localhost".`,
																						"type": "string",
																					},
																					"type": map[string]interface{}{
																						"description": `type indicates which kind of AppArmor profile will be applied.
Valid options are:
  Localhost - a profile pre-loaded on the node.
  RuntimeDefault - the container runtime's default profile.
  Unconfined - no AppArmor enforcement.`,
																						"type": "string",
																					},
																				},
																				"required": []interface{}{
																					"type",
																				},
																				"type": "object",
																			},
																			"fsGroup": map[string]interface{}{
																				"description": `A special supplemental group that applies to all containers in a pod.
Some volume types allow the Kubelet to change the ownership of that volume
to be owned by the pod:

1. The owning GID will be the FSGroup
2. The setgid bit is set (new files created in the volume will be owned by FSGroup)
3. The permission bits are OR'd with rw-rw----

If unset, the Kubelet will not modify the ownership and permissions of any volume.
Note that this field cannot be set when spec.os.name is windows.`,
																				"format": "int64",
																				"type":   "integer",
																			},
																			"fsGroupChangePolicy": map[string]interface{}{
																				"description": `fsGroupChangePolicy defines behavior of changing ownership and permission of the volume
before being exposed inside Pod. This field will only apply to
volume types which support fsGroup based ownership(and permissions).
It will have no effect on ephemeral volume types such as: secret, configmaps
and emptydir.
Valid values are "OnRootMismatch" and "Always". If not specified, "Always" is used.
Note that this field cannot be set when spec.os.name is windows.`,
																				"type": "string",
																			},
																			"runAsGroup": map[string]interface{}{
																				"description": `The GID to run the entrypoint of the container process.
Uses runtime default if unset.
May also be set in SecurityContext.  If set in both SecurityContext and
PodSecurityContext, the value specified in SecurityContext takes precedence
for that container.
Note that this field cannot be set when spec.os.name is windows.`,
																				"format": "int64",
																				"type":   "integer",
																			},
																			"runAsNonRoot": map[string]interface{}{
																				"description": `Indicates that the container must run as a non-root user.
If true, the Kubelet will validate the image at runtime to ensure that it
does not run as UID 0 (root) and fail to start the container if it does.
If unset or false, no such validation will be performed.
May also be set in SecurityContext.  If set in both SecurityContext and
PodSecurityContext, the value specified in SecurityContext takes precedence.`,
																				"type": "boolean",
																			},
																			"runAsUser": map[string]interface{}{
																				"description": `The UID to run the entrypoint of the container process.
Defaults to user specified in image metadata if unspecified.
May also be set in SecurityContext.  If set in both SecurityContext and
PodSecurityContext, the value specified in SecurityContext takes precedence
for that container.
Note that this field cannot be set when spec.os.name is windows.`,
																				"format": "int64",
																				"type":   "integer",
																			},
																			"seLinuxChangePolicy": map[string]interface{}{
																				"description": `seLinuxChangePolicy defines how the container's SELinux label is applied to all volumes used by the Pod.
It has no effect on nodes that do not support SELinux or to volumes does not support SELinux.
Valid values are "MountOption" and "Recursive".

"Recursive" means relabeling of all files on all Pod volumes by the container runtime.
This may be slow for large volumes, but allows mixing privileged and unprivileged Pods sharing the same volume on the same node.

"MountOption" mounts all eligible Pod volumes with ` + "`" + `-o context` + "`" + ` mount option.
This requires all Pods that share the same volume to use the same SELinux label.
It is not possible to share the same volume among privileged and unprivileged Pods.
Eligible volumes are in-tree FibreChannel and iSCSI volumes, and all CSI volumes
whose CSI driver announces SELinux support by setting spec.seLinuxMount: true in their
CSIDriver instance. Other volumes are always re-labelled recursively.
"MountOption" value is allowed only when SELinuxMount feature gate is enabled.

If not specified and SELinuxMount feature gate is enabled, "MountOption" is used.
If not specified and SELinuxMount feature gate is disabled, "MountOption" is used for ReadWriteOncePod volumes
and "Recursive" for all other volumes.

This field affects only Pods that have SELinux label set, either in PodSecurityContext or in SecurityContext of all containers.

All Pods that use the same volume should use the same seLinuxChangePolicy, otherwise some pods can get stuck in ContainerCreating state.
Note that this field cannot be set when spec.os.name is windows.`,
																				"type": "string",
																			},
																			"seLinuxOptions": map[string]interface{}{
																				"description": `The SELinux context to be applied to all containers.
If unspecified, the container runtime will allocate a random SELinux context for each
container.  May also be set in SecurityContext.  If set in
both SecurityContext and PodSecurityContext, the value specified in SecurityContext
takes precedence for that container.
Note that this field cannot be set when spec.os.name is windows.`,
																				"properties": map[string]interface{}{
																					"level": map[string]interface{}{
																						"description": "Level is SELinux level label that applies to the container.",
																						"type":        "string",
																					},
																					"role": map[string]interface{}{
																						"description": "Role is a SELinux role label that applies to the container.",
																						"type":        "string",
																					},
																					"type": map[string]interface{}{
																						"description": "Type is a SELinux type label that applies to the container.",
																						"type":        "string",
																					},
																					"user": map[string]interface{}{
																						"description": "User is a SELinux user label that applies to the container.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"seccompProfile": map[string]interface{}{
																				"description": `The seccomp options to use by the containers in this pod.
Note that this field cannot be set when spec.os.name is windows.`,
																				"properties": map[string]interface{}{
																					"localhostProfile": map[string]interface{}{
																						"description": `localhostProfile indicates a profile defined in a file on the node should be used.
The profile must be preconfigured on the node to work.
Must be a descending path, relative to the kubelet's configured seccomp profile location.
Must be set if type is "Localhost". Must NOT be set for any other type.`,
																						"type": "string",
																					},
																					"type": map[string]interface{}{
																						"description": `type indicates which kind of seccomp profile will be applied.
Valid options are:

Localhost - a profile defined in a file on the node should be used.
RuntimeDefault - the container runtime default profile should be used.
Unconfined - no profile should be applied.`,
																						"type": "string",
																					},
																				},
																				"required": []interface{}{
																					"type",
																				},
																				"type": "object",
																			},
																			"supplementalGroups": map[string]interface{}{
																				"description": `A list of groups applied to the first process run in each container, in
addition to the container's primary GID and fsGroup (if specified).  If
the SupplementalGroupsPolicy feature is enabled, the
supplementalGroupsPolicy field determines whether these are in addition
to or instead of any group memberships defined in the container image.
If unspecified, no additional groups are added, though group memberships
defined in the container image may still be used, depending on the
supplementalGroupsPolicy field.
Note that this field cannot be set when spec.os.name is windows.`,
																				"items": map[string]interface{}{
																					"format": "int64",
																					"type":   "integer",
																				},
																				"type":                   "array",
																				"x-kubernetes-list-type": "atomic",
																			},
																			"supplementalGroupsPolicy": map[string]interface{}{
																				"description": `Defines how supplemental groups of the first container processes are calculated.
Valid values are "Merge" and "Strict". If not specified, "Merge" is used.
(Alpha) Using the field requires the SupplementalGroupsPolicy feature gate to be enabled
and the container runtime must implement support for this feature.
Note that this field cannot be set when spec.os.name is windows.`,
																				"type": "string",
																			},
																			"sysctls": map[string]interface{}{
																				"description": `Sysctls hold a list of namespaced sysctls used for the pod. Pods with unsupported
sysctls (by the container runtime) might fail to launch.
Note that this field cannot be set when spec.os.name is windows.`,
																				"items": map[string]interface{}{
																					"description": "Sysctl defines a kernel parameter to be set",
																					"properties": map[string]interface{}{
																						"name": map[string]interface{}{
																							"description": "Name of a property to set",
																							"type":        "string",
																						},
																						"value": map[string]interface{}{
																							"description": "Value of a property to set",
																							"type":        "string",
																						},
																					},
																					"required": []interface{}{
																						"name",
																						"value",
																					},
																					"type": "object",
																				},
																				"type":                   "array",
																				"x-kubernetes-list-type": "atomic",
																			},
																			"windowsOptions": map[string]interface{}{
																				"description": `The Windows specific settings applied to all containers.
If unspecified, the options within a container's SecurityContext will be used.
If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
Note that this field cannot be set when spec.os.name is linux.`,
																				"properties": map[string]interface{}{
																					"gmsaCredentialSpec": map[string]interface{}{
																						"description": `GMSACredentialSpec is where the GMSA admission webhook
(https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the
GMSA credential spec named by the GMSACredentialSpecName field.`,
																						"type": "string",
																					},
																					"gmsaCredentialSpecName": map[string]interface{}{
																						"description": "GMSACredentialSpecName is the name of the GMSA credential spec to use.",
																						"type":        "string",
																					},
																					"hostProcess": map[string]interface{}{
																						"description": `HostProcess determines if a container should be run as a 'Host Process' container.
All of a Pod's containers must have the same effective HostProcess value
(it is not allowed to have a mix of HostProcess containers and non-HostProcess containers).
In addition, if HostProcess is true then HostNetwork must also be set to true.`,
																						"type": "boolean",
																					},
																					"runAsUserName": map[string]interface{}{
																						"description": `The UserName in Windows to run the entrypoint of the container process.
Defaults to the user specified in image metadata if unspecified.
May also be set in PodSecurityContext. If set in both SecurityContext and
PodSecurityContext, the value specified in SecurityContext takes precedence.`,
																						"type": "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																	"serviceAccount": map[string]interface{}{
																		"description": `DeprecatedServiceAccount is a deprecated alias for ServiceAccountName.
Deprecated: Use serviceAccountName instead.`,
																		"type": "string",
																	},
																	"serviceAccountName": map[string]interface{}{
																		"description": `ServiceAccountName is the name of the ServiceAccount to use to run this pod.
More info: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/`,
																		"type": "string",
																	},
																	"setHostnameAsFQDN": map[string]interface{}{
																		"description": `If true the pod's hostname will be configured as the pod's FQDN, rather than the leaf name (the default).
In Linux containers, this means setting the FQDN in the hostname field of the kernel (the nodename field of struct utsname).
In Windows containers, this means setting the registry value of hostname for the registry key HKEY_LOCAL_MACHINE\\SYSTEM\\CurrentControlSet\\Services\\Tcpip\\Parameters to FQDN.
If a pod does not have FQDN, this has no effect.
Default to false.`,
																		"type": "boolean",
																	},
																	"shareProcessNamespace": map[string]interface{}{
																		"description": `Share a single process namespace between all of the containers in a pod.
When this is set containers will be able to view and signal processes from other containers
in the same pod, and the first process in each container will not be assigned PID 1.
HostPID and ShareProcessNamespace cannot both be set.
Optional: Default to false.`,
																		"type": "boolean",
																	},
																	"subdomain": map[string]interface{}{
																		"description": `If specified, the fully qualified Pod hostname will be "<hostname>.<subdomain>.<pod namespace>.svc.<cluster domain>".
If not specified, the pod will not have a domainname at all.`,
																		"type": "string",
																	},
																	"terminationGracePeriodSeconds": map[string]interface{}{
																		"description": `Optional duration in seconds the pod needs to terminate gracefully. May be decreased in delete request.
Value must be non-negative integer. The value zero indicates stop immediately via
the kill signal (no opportunity to shut down).
If this value is nil, the default grace period will be used instead.
The grace period is the duration in seconds after the processes running in the pod are sent
a termination signal and the time when the processes are forcibly halted with a kill signal.
Set this value longer than the expected cleanup time for your process.
Defaults to 30 seconds.`,
																		"format": "int64",
																		"type":   "integer",
																	},
																	"tolerations": map[string]interface{}{
																		"description": "If specified, the pod's tolerations.",
																		"items": map[string]interface{}{
																			"description": `The pod this Toleration is attached to tolerates any taint that matches
the triple <key,value,effect> using the matching operator <operator>.`,
																			"properties": map[string]interface{}{
																				"effect": map[string]interface{}{
																					"description": `Effect indicates the taint effect to match. Empty means match all taint effects.
When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.`,
																					"type": "string",
																				},
																				"key": map[string]interface{}{
																					"description": `Key is the taint key that the toleration applies to. Empty means match all taint keys.
If the key is empty, operator must be Exists; this combination means to match all values and all keys.`,
																					"type": "string",
																				},
																				"operator": map[string]interface{}{
																					"description": `Operator represents a key's relationship to the value.
Valid operators are Exists, Equal, Lt, and Gt. Defaults to Equal.
Exists is equivalent to wildcard for value, so that a pod can
tolerate all taints of a particular category.
Lt and Gt perform numeric comparisons (requires feature gate TaintTolerationComparisonOperators).`,
																					"type": "string",
																				},
																				"tolerationSeconds": map[string]interface{}{
																					"description": `TolerationSeconds represents the period of time the toleration (which must be
of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default,
it is not set, which means tolerate the taint forever (do not evict). Zero and
negative values will be treated as 0 (evict immediately) by the system.`,
																					"format": "int64",
																					"type":   "integer",
																				},
																				"value": map[string]interface{}{
																					"description": `Value is the taint value the toleration matches to.
If the operator is Exists, the value should be empty, otherwise just a regular string.`,
																					"type": "string",
																				},
																			},
																			"type": "object",
																		},
																		"type":                   "array",
																		"x-kubernetes-list-type": "atomic",
																	},
																	"topologySpreadConstraints": map[string]interface{}{
																		"description": `TopologySpreadConstraints describes how a group of pods ought to spread across topology
domains. Scheduler will schedule pods in a way which abides by the constraints.
All topologySpreadConstraints are ANDed.`,
																		"items": map[string]interface{}{
																			"description": "TopologySpreadConstraint specifies how to spread matching pods among the given topology.",
																			"properties": map[string]interface{}{
																				"labelSelector": map[string]interface{}{
																					"description": `LabelSelector is used to find matching pods.
Pods that match this label selector are counted to determine the number of pods
in their corresponding topology domain.`,
																					"properties": map[string]interface{}{
																						"matchExpressions": map[string]interface{}{
																							"description": "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																							"items": map[string]interface{}{
																								"description": `A label selector requirement is a selector that contains values, a key, and an operator that
relates the key and values.`,
																								"properties": map[string]interface{}{
																									"key": map[string]interface{}{
																										"description": "key is the label key that the selector applies to.",
																										"type":        "string",
																									},
																									"operator": map[string]interface{}{
																										"description": `operator represents a key's relationship to a set of values.
Valid operators are In, NotIn, Exists and DoesNotExist.`,
																										"type": "string",
																									},
																									"values": map[string]interface{}{
																										"description": `values is an array of string values. If the operator is In or NotIn,
the values array must be non-empty. If the operator is Exists or DoesNotExist,
the values array must be empty. This array is replaced during a strategic
merge patch.`,
																										"items": map[string]interface{}{
																											"type": "string",
																										},
																										"type":                   "array",
																										"x-kubernetes-list-type": "atomic",
																									},
																								},
																								"required": []interface{}{
																									"key",
																									"operator",
																								},
																								"type": "object",
																							},
																							"type":                   "array",
																							"x-kubernetes-list-type": "atomic",
																						},
																						"matchLabels": map[string]interface{}{
																							"additionalProperties": map[string]interface{}{
																								"type": "string",
																							},
																							"description": `matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels
map is equivalent to an element of matchExpressions, whose key field is "key", the
operator is "In", and the values array contains only "value". The requirements are ANDed.`,
																							"type": "object",
																						},
																					},
																					"type":                  "object",
																					"x-kubernetes-map-type": "atomic",
																				},
																				"matchLabelKeys": map[string]interface{}{
																					"description": `MatchLabelKeys is a set of pod label keys to select the pods over which
spreading will be calculated. The keys are used to lookup values from the
incoming pod labels, those key-value labels are ANDed with labelSelector
to select the group of existing pods over which spreading will be calculated
for the incoming pod. The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.
MatchLabelKeys cannot be set when LabelSelector isn't set.
Keys that don't exist in the incoming pod labels will
be ignored. A null or empty list means only match against labelSelector.

This is a beta field and requires the MatchLabelKeysInPodTopologySpread feature gate to be enabled (enabled by default).`,
																					"items": map[string]interface{}{
																						"type": "string",
																					},
																					"type":                   "array",
																					"x-kubernetes-list-type": "atomic",
																				},
																				"maxSkew": map[string]interface{}{
																					"description": `MaxSkew describes the degree to which pods may be unevenly distributed.
When ` + "`" + `whenUnsatisfiable=DoNotSchedule` + "`" + `, it is the maximum permitted difference
between the number of matching pods in the target topology and the global minimum.
The global minimum is the minimum number of matching pods in an eligible domain
or zero if the number of eligible domains is less than MinDomains.
For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same
labelSelector spread as 2/2/1:
In this case, the global minimum is 1.
| zone1 | zone2 | zone3 |
|  P P  |  P P  |   P   |
- if MaxSkew is 1, incoming pod can only be scheduled to zone3 to become 2/2/2;
scheduling it onto zone1(zone2) would make the ActualSkew(3-1) on zone1(zone2)
violate MaxSkew(1).
- if MaxSkew is 2, incoming pod can be scheduled onto any zone.
When ` + "`" + `whenUnsatisfiable=ScheduleAnyway` + "`" + `, it is used to give higher precedence
to topologies that satisfy it.
It's a required field. Default value is 1 and 0 is not allowed.`,
																					"format": "int32",
																					"type":   "integer",
																				},
																				"minDomains": map[string]interface{}{
																					"description": `MinDomains indicates a minimum number of eligible domains.
When the number of eligible domains with matching topology keys is less than minDomains,
Pod Topology Spread treats "global minimum" as 0, and then the calculation of Skew is performed.
And when the number of eligible domains with matching topology keys equals or greater than minDomains,
this value has no effect on scheduling.
As a result, when the number of eligible domains is less than minDomains,
scheduler won't schedule more than maxSkew Pods to those domains.
If value is nil, the constraint behaves as if MinDomains is equal to 1.
Valid values are integers greater than 0.
When value is not nil, WhenUnsatisfiable must be DoNotSchedule.

For example, in a 3-zone cluster, MaxSkew is set to 2, MinDomains is set to 5 and pods with the same
labelSelector spread as 2/2/2:
| zone1 | zone2 | zone3 |
|  P P  |  P P  |  P P  |
The number of domains is less than 5(MinDomains), so "global minimum" is treated as 0.
In this situation, new pod with the same labelSelector cannot be scheduled,
because computed skew will be 3(3 - 0) if new Pod is scheduled to any of the three zones,
it will violate MaxSkew.`,
																					"format": "int32",
																					"type":   "integer",
																				},
																				"nodeAffinityPolicy": map[string]interface{}{
																					"description": `NodeAffinityPolicy indicates how we will treat Pod's nodeAffinity/nodeSelector
when calculating pod topology spread skew. Options are:
- Honor: only nodes matching nodeAffinity/nodeSelector are included in the calculations.
- Ignore: nodeAffinity/nodeSelector are ignored. All nodes are included in the calculations.

If this value is nil, the behavior is equivalent to the Honor policy.`,
																					"type": "string",
																				},
																				"nodeTaintsPolicy": map[string]interface{}{
																					"description": `NodeTaintsPolicy indicates how we will treat node taints when calculating
pod topology spread skew. Options are:
- Honor: nodes without taints, along with tainted nodes for which the incoming pod
has a toleration, are included.
- Ignore: node taints are ignored. All nodes are included.

If this value is nil, the behavior is equivalent to the Ignore policy.`,
																					"type": "string",
																				},
																				"topologyKey": map[string]interface{}{
																					"description": `TopologyKey is the key of node labels. Nodes that have a label with this key
and identical values are considered to be in the same topology.
We consider each <key, value> as a "bucket", and try to put balanced number
of pods into each bucket.
We define a domain as a particular instance of a topology.
Also, we define an eligible domain as a domain whose nodes meet the requirements of
nodeAffinityPolicy and nodeTaintsPolicy.
e.g. If TopologyKey is "kubernetes.io/hostname", each Node is a domain of that topology.
And, if TopologyKey is "topology.kubernetes.io/zone", each zone is a domain of that topology.
It's a required field.`,
																					"type": "string",
																				},
																				"whenUnsatisfiable": map[string]interface{}{
																					"description": `WhenUnsatisfiable indicates how to deal with a pod if it doesn't satisfy
the spread constraint.
- DoNotSchedule (default) tells the scheduler not to schedule it.
- ScheduleAnyway tells the scheduler to schedule the pod in any location,
  but giving higher precedence to topologies that would help reduce the
  skew.
A constraint is considered "Unsatisfiable" for an incoming pod
if and only if every possible node assignment for that pod would violate
"MaxSkew" on some topology.
For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same
labelSelector spread as 3/1/1:
| zone1 | zone2 | zone3 |
| P P P |   P   |   P   |
If WhenUnsatisfiable is set to DoNotSchedule, incoming pod can only be scheduled
to zone2(zone3) to become 3/2/1(3/1/2) as ActualSkew(2-1) on zone2(zone3) satisfies
MaxSkew(1). In other words, the cluster can still be imbalanced, but scheduler
won't make it *more* imbalanced.
It's a required field.`,
																					"type": "string",
																				},
																			},
																			"required": []interface{}{
																				"maxSkew",
																				"topologyKey",
																				"whenUnsatisfiable",
																			},
																			"type": "object",
																		},
																		"type": "array",
																		"x-kubernetes-list-map-keys": []interface{}{
																			"topologyKey",
																			"whenUnsatisfiable",
																		},
																		"x-kubernetes-list-type": "map",
																	},
																	"volumes": map[string]interface{}{
																		"description": `List of volumes that can be mounted by containers belonging to the pod.
More info: https://kubernetes.io/docs/concepts/storage/volumes`,
																		"items": map[string]interface{}{
																			"description": "Volume represents a named volume in a pod that may be accessed by any container in the pod.",
																			"properties": map[string]interface{}{
																				"awsElasticBlockStore": map[string]interface{}{
																					"description": `awsElasticBlockStore represents an AWS Disk resource that is attached to a
kubelet's host machine and then exposed to the pod.
Deprecated: AWSElasticBlockStore is deprecated. All operations for the in-tree
awsElasticBlockStore type are redirected to the ebs.csi.aws.com CSI driver.
More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore`,
																					"properties": map[string]interface{}{
																						"fsType": map[string]interface{}{
																							"description": `fsType is the filesystem type of the volume that you want to mount.
Tip: Ensure that the filesystem type is supported by the host operating system.
Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore`,
																							"type": "string",
																						},
																						"partition": map[string]interface{}{
																							"description": `partition is the partition in the volume that you want to mount.
If omitted, the default is to mount by volume name.
Examples: For volume /dev/sda1, you specify the partition as "1".
Similarly, the volume partition for /dev/sda is "0" (or you can leave the property empty).`,
																							"format": "int32",
																							"type":   "integer",
																						},
																						"readOnly": map[string]interface{}{
																							"description": `readOnly value true will force the readOnly setting in VolumeMounts.
More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore`,
																							"type": "boolean",
																						},
																						"volumeID": map[string]interface{}{
																							"description": `volumeID is unique ID of the persistent disk resource in AWS (Amazon EBS volume).
More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore`,
																							"type": "string",
																						},
																					},
																					"required": []interface{}{
																						"volumeID",
																					},
																					"type": "object",
																				},
																				"azureDisk": map[string]interface{}{
																					"description": `azureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.
Deprecated: AzureDisk is deprecated. All operations for the in-tree azureDisk type
are redirected to the disk.csi.azure.com CSI driver.`,
																					"properties": map[string]interface{}{
																						"cachingMode": map[string]interface{}{
																							"description": "cachingMode is the Host Caching mode: None, Read Only, Read Write.",
																							"type":        "string",
																						},
																						"diskName": map[string]interface{}{
																							"description": "diskName is the Name of the data disk in the blob storage",
																							"type":        "string",
																						},
																						"diskURI": map[string]interface{}{
																							"description": "diskURI is the URI of data disk in the blob storage",
																							"type":        "string",
																						},
																						"fsType": map[string]interface{}{
																							"default": "ext4",
																							"description": `fsType is Filesystem type to mount.
Must be a filesystem type supported by the host operating system.
Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.`,
																							"type": "string",
																						},
																						"kind": map[string]interface{}{
																							"description": "kind expected values are Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared",
																							"type":        "string",
																						},
																						"readOnly": map[string]interface{}{
																							"default": false,
																							"description": `readOnly Defaults to false (read/write). ReadOnly here will force
the ReadOnly setting in VolumeMounts.`,
																							"type": "boolean",
																						},
																					},
																					"required": []interface{}{
																						"diskName",
																						"diskURI",
																					},
																					"type": "object",
																				},
																				"azureFile": map[string]interface{}{
																					"description": `azureFile represents an Azure File Service mount on the host and bind mount to the pod.
Deprecated: AzureFile is deprecated. All operations for the in-tree azureFile type
are redirected to the file.csi.azure.com CSI driver.`,
																					"properties": map[string]interface{}{
																						"readOnly": map[string]interface{}{
																							"description": `readOnly defaults to false (read/write). ReadOnly here will force
the ReadOnly setting in VolumeMounts.`,
																							"type": "boolean",
																						},
																						"secretName": map[string]interface{}{
																							"description": "secretName is the  name of secret that contains Azure Storage Account Name and Key",
																							"type":        "string",
																						},
																						"shareName": map[string]interface{}{
																							"description": "shareName is the azure share Name",
																							"type":        "string",
																						},
																					},
																					"required": []interface{}{
																						"secretName",
																						"shareName",
																					},
																					"type": "object",
																				},
																				"cephfs": map[string]interface{}{
																					"description": `cephFS represents a Ceph FS mount on the host that shares a pod's lifetime.
Deprecated: CephFS is deprecated and the in-tree cephfs type is no longer supported.`,
																					"properties": map[string]interface{}{
																						"monitors": map[string]interface{}{
																							"description": `monitors is Required: Monitors is a collection of Ceph monitors
More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it`,
																							"items": map[string]interface{}{
																								"type": "string",
																							},
																							"type":                   "array",
																							"x-kubernetes-list-type": "atomic",
																						},
																						"path": map[string]interface{}{
																							"description": "path is Optional: Used as the mounted root, rather than the full Ceph tree, default is /",
																							"type":        "string",
																						},
																						"readOnly": map[string]interface{}{
																							"description": `readOnly is Optional: Defaults to false (read/write). ReadOnly here will force
the ReadOnly setting in VolumeMounts.
More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it`,
																							"type": "boolean",
																						},
																						"secretFile": map[string]interface{}{
																							"description": `secretFile is Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret
More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it`,
																							"type": "string",
																						},
																						"secretRef": map[string]interface{}{
																							"description": `secretRef is Optional: SecretRef is reference to the authentication secret for User, default is empty.
More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it`,
																							"properties": map[string]interface{}{
																								"name": map[string]interface{}{
																									"default": "",
																									"description": `Name of the referent.
This field is effectively required, but due to backwards compatibility is
allowed to be empty. Instances of this type with an empty value here are
almost certainly wrong.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names`,
																									"type": "string",
																								},
																							},
																							"type":                  "object",
																							"x-kubernetes-map-type": "atomic",
																						},
																						"user": map[string]interface{}{
																							"description": `user is optional: User is the rados user name, default is admin
More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it`,
																							"type": "string",
																						},
																					},
																					"required": []interface{}{
																						"monitors",
																					},
																					"type": "object",
																				},
																				"cinder": map[string]interface{}{
																					"description": `cinder represents a cinder volume attached and mounted on kubelets host machine.
Deprecated: Cinder is deprecated. All operations for the in-tree cinder type
are redirected to the cinder.csi.openstack.org CSI driver.
More info: https://examples.k8s.io/mysql-cinder-pd/README.md`,
																					"properties": map[string]interface{}{
																						"fsType": map[string]interface{}{
																							"description": `fsType is the filesystem type to mount.
Must be a filesystem type supported by the host operating system.
Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
More info: https://examples.k8s.io/mysql-cinder-pd/README.md`,
																							"type": "string",
																						},
																						"readOnly": map[string]interface{}{
																							"description": `readOnly defaults to false (read/write). ReadOnly here will force
the ReadOnly setting in VolumeMounts.
More info: https://examples.k8s.io/mysql-cinder-pd/README.md`,
																							"type": "boolean",
																						},
																						"secretRef": map[string]interface{}{
																							"description": `secretRef is optional: points to a secret object containing parameters used to connect
to OpenStack.`,
																							"properties": map[string]interface{}{
																								"name": map[string]interface{}{
																									"default": "",
																									"description": `Name of the referent.
This field is effectively required, but due to backwards compatibility is
allowed to be empty. Instances of this type with an empty value here are
almost certainly wrong.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names`,
																									"type": "string",
																								},
																							},
																							"type":                  "object",
																							"x-kubernetes-map-type": "atomic",
																						},
																						"volumeID": map[string]interface{}{
																							"description": `volumeID used to identify the volume in cinder.
More info: https://examples.k8s.io/mysql-cinder-pd/README.md`,
																							"type": "string",
																						},
																					},
																					"required": []interface{}{
																						"volumeID",
																					},
																					"type": "object",
																				},
																				"configMap": map[string]interface{}{
																					"description": "configMap represents a configMap that should populate this volume",
																					"properties": map[string]interface{}{
																						"defaultMode": map[string]interface{}{
																							"description": `defaultMode is optional: mode bits used to set permissions on created files by default.
Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.
YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.
Defaults to 0644.
Directories within the path are not affected by this setting.
This might be in conflict with other options that affect the file
mode, like fsGroup, and the result can be other mode bits set.`,
																							"format": "int32",
																							"type":   "integer",
																						},
																						"items": map[string]interface{}{
																							"description": `items if unspecified, each key-value pair in the Data field of the referenced
ConfigMap will be projected into the volume as a file whose name is the
key and content is the value. If specified, the listed keys will be
projected into the specified paths, and unlisted keys will not be
present. If a key is specified which is not present in the ConfigMap,
the volume setup will error unless it is marked optional. Paths must be
relative and may not contain the '..' path or start with '..'.`,
																							"items": map[string]interface{}{
																								"description": "Maps a string key to a path within a volume.",
																								"properties": map[string]interface{}{
																									"key": map[string]interface{}{
																										"description": "key is the key to project.",
																										"type":        "string",
																									},
																									"mode": map[string]interface{}{
																										"description": `mode is Optional: mode bits used to set permissions on this file.
Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.
YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.
If not specified, the volume defaultMode will be used.
This might be in conflict with other options that affect the file
mode, like fsGroup, and the result can be other mode bits set.`,
																										"format": "int32",
																										"type":   "integer",
																									},
																									"path": map[string]interface{}{
																										"description": `path is the relative path of the file to map the key to.
May not be an absolute path.
May not contain the path element '..'.
May not start with the string '..'.`,
																										"type": "string",
																									},
																								},
																								"required": []interface{}{
																									"key",
																									"path",
																								},
																								"type": "object",
																							},
																							"type":                   "array",
																							"x-kubernetes-list-type": "atomic",
																						},
																						"name": map[string]interface{}{
																							"default": "",
																							"description": `Name of the referent.
This field is effectively required, but due to backwards compatibility is
allowed to be empty. Instances of this type with an empty value here are
almost certainly wrong.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names`,
																							"type": "string",
																						},
																						"optional": map[string]interface{}{
																							"description": "optional specify whether the ConfigMap or its keys must be defined",
																							"type":        "boolean",
																						},
																					},
																					"type":                  "object",
																					"x-kubernetes-map-type": "atomic",
																				},
																				"csi": map[string]interface{}{
																					"description": "csi (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers.",
																					"properties": map[string]interface{}{
																						"driver": map[string]interface{}{
																							"description": `driver is the name of the CSI driver that handles this volume.
Consult with your admin for the correct name as registered in the cluster.`,
																							"type": "string",
																						},
																						"fsType": map[string]interface{}{
																							"description": `fsType to mount. Ex. "ext4", "xfs", "ntfs".
If not provided, the empty value is passed to the associated CSI driver
which will determine the default filesystem to apply.`,
																							"type": "string",
																						},
																						"nodePublishSecretRef": map[string]interface{}{
																							"description": `nodePublishSecretRef is a reference to the secret object containing
sensitive information to pass to the CSI driver to complete the CSI
NodePublishVolume and NodeUnpublishVolume calls.
This field is optional, and  may be empty if no secret is required. If the
secret object contains more than one secret, all secret references are passed.`,
																							"properties": map[string]interface{}{
																								"name": map[string]interface{}{
																									"default": "",
																									"description": `Name of the referent.
This field is effectively required, but due to backwards compatibility is
allowed to be empty. Instances of this type with an empty value here are
almost certainly wrong.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names`,
																									"type": "string",
																								},
																							},
																							"type":                  "object",
																							"x-kubernetes-map-type": "atomic",
																						},
																						"readOnly": map[string]interface{}{
																							"description": `readOnly specifies a read-only configuration for the volume.
Defaults to false (read/write).`,
																							"type": "boolean",
																						},
																						"volumeAttributes": map[string]interface{}{
																							"additionalProperties": map[string]interface{}{
																								"type": "string",
																							},
																							"description": `volumeAttributes stores driver-specific properties that are passed to the CSI
driver. Consult your driver's documentation for supported values.`,
																							"type": "object",
																						},
																					},
																					"required": []interface{}{
																						"driver",
																					},
																					"type": "object",
																				},
																				"downwardAPI": map[string]interface{}{
																					"description": "downwardAPI represents downward API about the pod that should populate this volume",
																					"properties": map[string]interface{}{
																						"defaultMode": map[string]interface{}{
																							"description": `Optional: mode bits to use on created files by default. Must be a
Optional: mode bits used to set permissions on created files by default.
Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.
YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.
Defaults to 0644.
Directories within the path are not affected by this setting.
This might be in conflict with other options that affect the file
mode, like fsGroup, and the result can be other mode bits set.`,
																							"format": "int32",
																							"type":   "integer",
																						},
																						"items": map[string]interface{}{
																							"description": "Items is a list of downward API volume file",
																							"items": map[string]interface{}{
																								"description": "DownwardAPIVolumeFile represents information to create the file containing the pod field",
																								"properties": map[string]interface{}{
																									"fieldRef": map[string]interface{}{
																										"description": "Required: Selects a field of the pod: only annotations, labels, name, namespace and uid are supported.",
																										"properties": map[string]interface{}{
																											"apiVersion": map[string]interface{}{
																												"description": "Version of the schema the FieldPath is written in terms of, defaults to \"v1\".",
																												"type":        "string",
																											},
																											"fieldPath": map[string]interface{}{
																												"description": "Path of the field to select in the specified API version.",
																												"type":        "string",
																											},
																										},
																										"required": []interface{}{
																											"fieldPath",
																										},
																										"type":                  "object",
																										"x-kubernetes-map-type": "atomic",
																									},
																									"mode": map[string]interface{}{
																										"description": `Optional: mode bits used to set permissions on this file, must be an octal value
between 0000 and 0777 or a decimal value between 0 and 511.
YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.
If not specified, the volume defaultMode will be used.
This might be in conflict with other options that affect the file
mode, like fsGroup, and the result can be other mode bits set.`,
																										"format": "int32",
																										"type":   "integer",
																									},
																									"path": map[string]interface{}{
																										"description": "Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",
																										"type":        "string",
																									},
																									"resourceFieldRef": map[string]interface{}{
																										"description": `Selects a resource of the container: only resources limits and requests
(limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.`,
																										"properties": map[string]interface{}{
																											"containerName": map[string]interface{}{
																												"description": "Container name: required for volumes, optional for env vars",
																												"type":        "string",
																											},
																											"divisor": map[string]interface{}{
																												"anyOf": []interface{}{
																													map[string]interface{}{
																														"type": "integer",
																													},
																													map[string]interface{}{
																														"type": "string",
																													},
																												},
																												"description":                "Specifies the output format of the exposed resources, defaults to \"1\"",
																												"pattern":                    `^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$`,
																												"x-kubernetes-int-or-string": true,
																											},
																											"resource": map[string]interface{}{
																												"description": "Required: resource to select",
																												"type":        "string",
																											},
																										},
																										"required": []interface{}{
																											"resource",
																										},
																										"type":                  "object",
																										"x-kubernetes-map-type": "atomic",
																									},
																								},
																								"required": []interface{}{
																									"path",
																								},
																								"type": "object",
																							},
																							"type":                   "array",
																							"x-kubernetes-list-type": "atomic",
																						},
																					},
																					"type": "object",
																				},
																				"emptyDir": map[string]interface{}{
																					"description": `emptyDir represents a temporary directory that shares a pod's lifetime.
More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir`,
																					"properties": map[string]interface{}{
																						"medium": map[string]interface{}{
																							"description": `medium represents what type of storage medium should back this directory.
The default is "" which means to use the node's default medium.
Must be an empty string (default) or Memory.
More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir`,
																							"type": "string",
																						},
																						"sizeLimit": map[string]interface{}{
																							"anyOf": []interface{}{
																								map[string]interface{}{
																									"type": "integer",
																								},
																								map[string]interface{}{
																									"type": "string",
																								},
																							},
																							"description": `sizeLimit is the total amount of local storage required for this EmptyDir volume.
The size limit is also applicable for memory medium.
The maximum usage on memory medium EmptyDir would be the minimum value between
the SizeLimit specified here and the sum of memory limits of all containers in a pod.
The default is nil which means that the limit is undefined.
More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir`,
																							"pattern":                    `^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$`,
																							"x-kubernetes-int-or-string": true,
																						},
																					},
																					"type": "object",
																				},
																				"ephemeral": map[string]interface{}{
																					"description": `ephemeral represents a volume that is handled by a cluster storage driver.
The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts,
and deleted when the pod is removed.

Use this if:
a) the volume is only needed while the pod runs,
b) features of normal volumes like restoring from snapshot or capacity
   tracking are needed,
c) the storage driver is specified through a storage class, and
d) the storage driver supports dynamic volume provisioning through
   a PersistentVolumeClaim (see EphemeralVolumeSource for more
   information on the connection between this volume type
   and PersistentVolumeClaim).

Use PersistentVolumeClaim or one of the vendor-specific
APIs for volumes that persist for longer than the lifecycle
of an individual pod.

Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to
be used that way - see the documentation of the driver for
more information.

A pod can use both types of ephemeral volumes and
persistent volumes at the same time.`,
																					"properties": map[string]interface{}{
																						"volumeClaimTemplate": map[string]interface{}{
																							"description": `Will be used to create a stand-alone PVC to provision the volume.
The pod in which this EphemeralVolumeSource is embedded will be the
owner of the PVC, i.e. the PVC will be deleted together with the
pod.  The name of the PVC will be ` + "`" + `<pod name>-<volume name>` + "`" + ` where
` + "`" + `<volume name>` + "`" + ` is the name from the ` + "`" + `PodSpec.Volumes` + "`" + ` array
entry. Pod validation will reject the pod if the concatenated name
is not valid for a PVC (for example, too long).

An existing PVC with that name that is not owned by the pod
will *not* be used for the pod to avoid using an unrelated
volume by mistake. Starting the pod is then blocked until
the unrelated PVC is removed. If such a pre-created PVC is
meant to be used by the pod, the PVC has to updated with an
owner reference to the pod once the pod exists. Normally
this should not be necessary, but it may be useful when
manually reconstructing a broken cluster.

This field is read-only and no changes will be made by Kubernetes
to the PVC after it has been created.

Required, must not be nil.`,
																							"properties": map[string]interface{}{
																								"metadata": map[string]interface{}{
																									"description": `May contain labels and annotations that will be copied into the PVC
when creating it. No other fields are allowed and will be rejected during
validation.`,
																									"properties": map[string]interface{}{
																										"annotations": map[string]interface{}{
																											"additionalProperties": map[string]interface{}{
																												"type": "string",
																											},
																											"type": "object",
																										},
																										"finalizers": map[string]interface{}{
																											"items": map[string]interface{}{
																												"type": "string",
																											},
																											"type": "array",
																										},
																										"labels": map[string]interface{}{
																											"additionalProperties": map[string]interface{}{
																												"type": "string",
																											},
																											"type": "object",
																										},
																										"name": map[string]interface{}{
																											"type": "string",
																										},
																										"namespace": map[string]interface{}{
																											"type": "string",
																										},
																									},
																									"type": "object",
																								},
																								"spec": map[string]interface{}{
																									"description": `The specification for the PersistentVolumeClaim. The entire content is
copied unchanged into the PVC that gets created from this
template. The same fields as in a PersistentVolumeClaim
are also valid here.`,
																									"properties": map[string]interface{}{
																										"accessModes": map[string]interface{}{
																											"description": `accessModes contains the desired access modes the volume should have.
More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1`,
																											"items": map[string]interface{}{
																												"type": "string",
																											},
																											"type":                   "array",
																											"x-kubernetes-list-type": "atomic",
																										},
																										"dataSource": map[string]interface{}{
																											"description": `dataSource field can be used to specify either:
* An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot)
* An existing PVC (PersistentVolumeClaim)
If the provisioner or an external controller can support the specified data source,
it will create a new volume based on the contents of the specified data source.
When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef,
and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified.
If the namespace is specified, then dataSourceRef will not be copied to dataSource.`,
																											"properties": map[string]interface{}{
																												"apiGroup": map[string]interface{}{
																													"description": `APIGroup is the group for the resource being referenced.
If APIGroup is not specified, the specified Kind must be in the core API group.
For any other third-party types, APIGroup is required.`,
																													"type": "string",
																												},
																												"kind": map[string]interface{}{
																													"description": "Kind is the type of resource being referenced",
																													"type":        "string",
																												},
																												"name": map[string]interface{}{
																													"description": "Name is the name of resource being referenced",
																													"type":        "string",
																												},
																											},
																											"required": []interface{}{
																												"kind",
																												"name",
																											},
																											"type":                  "object",
																											"x-kubernetes-map-type": "atomic",
																										},
																										"dataSourceRef": map[string]interface{}{
																											"description": `dataSourceRef specifies the object from which to populate the volume with data, if a non-empty
volume is desired. This may be any object from a non-empty API group (non
core object) or a PersistentVolumeClaim object.
When this field is specified, volume binding will only succeed if the type of
the specified object matches some installed volume populator or dynamic
provisioner.
This field will replace the functionality of the dataSource field and as such
if both fields are non-empty, they must have the same value. For backwards
compatibility, when namespace isn't specified in dataSourceRef,
both fields (dataSource and dataSourceRef) will be set to the same
value automatically if one of them is empty and the other is non-empty.
When namespace is specified in dataSourceRef,
dataSource isn't set to the same value and must be empty.
There are three important differences between dataSource and dataSourceRef:
* While dataSource only allows two specific types of objects, dataSourceRef
  allows any non-core object, as well as PersistentVolumeClaim objects.
* While dataSource ignores disallowed values (dropping them), dataSourceRef
  preserves all values, and generates an error if a disallowed value is
  specified.
* While dataSource only allows local objects, dataSourceRef allows objects
  in any namespaces.
(Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled.
(Alpha) Using the namespace field of dataSourceRef requires the CrossNamespaceVolumeDataSource feature gate to be enabled.`,
																											"properties": map[string]interface{}{
																												"apiGroup": map[string]interface{}{
																													"description": `APIGroup is the group for the resource being referenced.
If APIGroup is not specified, the specified Kind must be in the core API group.
For any other third-party types, APIGroup is required.`,
																													"type": "string",
																												},
																												"kind": map[string]interface{}{
																													"description": "Kind is the type of resource being referenced",
																													"type":        "string",
																												},
																												"name": map[string]interface{}{
																													"description": "Name is the name of resource being referenced",
																													"type":        "string",
																												},
																												"namespace": map[string]interface{}{
																													"description": `Namespace is the namespace of resource being referenced
Note that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details.
(Alpha) This field requires the CrossNamespaceVolumeDataSource feature gate to be enabled.`,
																													"type": "string",
																												},
																											},
																											"required": []interface{}{
																												"kind",
																												"name",
																											},
																											"type": "object",
																										},
																										"resources": map[string]interface{}{
																											"description": `resources represents the minimum resources the volume should have.
Users are allowed to specify resource requirements
that are lower than previous value but must still be higher than capacity recorded in the
status field of the claim.
More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources`,
																											"properties": map[string]interface{}{
																												"limits": map[string]interface{}{
																													"additionalProperties": map[string]interface{}{
																														"anyOf": []interface{}{
																															map[string]interface{}{
																																"type": "integer",
																															},
																															map[string]interface{}{
																																"type": "string",
																															},
																														},
																														"pattern":                    `^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$`,
																														"x-kubernetes-int-or-string": true,
																													},
																													"description": `Limits describes the maximum amount of compute resources allowed.
More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/`,
																													"type": "object",
																												},
																												"requests": map[string]interface{}{
																													"additionalProperties": map[string]interface{}{
																														"anyOf": []interface{}{
																															map[string]interface{}{
																																"type": "integer",
																															},
																															map[string]interface{}{
																																"type": "string",
																															},
																														},
																														"pattern":                    `^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$`,
																														"x-kubernetes-int-or-string": true,
																													},
																													"description": `Requests describes the minimum amount of compute resources required.
If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,
otherwise to an implementation-defined value. Requests cannot exceed Limits.
More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/`,
																													"type": "object",
																												},
																											},
																											"type": "object",
																										},
																										"selector": map[string]interface{}{
																											"description": "selector is a label query over volumes to consider for binding.",
																											"properties": map[string]interface{}{
																												"matchExpressions": map[string]interface{}{
																													"description": "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																													"items": map[string]interface{}{
																														"description": `A label selector requirement is a selector that contains values, a key, and an operator that
relates the key and values.`,
																														"properties": map[string]interface{}{
																															"key": map[string]interface{}{
																																"description": "key is the label key that the selector applies to.",
																																"type":        "string",
																															},
																															"operator": map[string]interface{}{
																																"description": `operator represents a key's relationship to a set of values.
Valid operators are In, NotIn, Exists and DoesNotExist.`,
																																"type": "string",
																															},
																															"values": map[string]interface{}{
																																"description": `values is an array of string values. If the operator is In or NotIn,
the values array must be non-empty. If the operator is Exists or DoesNotExist,
the values array must be empty. This array is replaced during a strategic
merge patch.`,
																																"items": map[string]interface{}{
																																	"type": "string",
																																},
																																"type":                   "array",
																																"x-kubernetes-list-type": "atomic",
																															},
																														},
																														"required": []interface{}{
																															"key",
																															"operator",
																														},
																														"type": "object",
																													},
																													"type":                   "array",
																													"x-kubernetes-list-type": "atomic",
																												},
																												"matchLabels": map[string]interface{}{
																													"additionalProperties": map[string]interface{}{
																														"type": "string",
																													},
																													"description": `matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels
map is equivalent to an element of matchExpressions, whose key field is "key", the
operator is "In", and the values array contains only "value". The requirements are ANDed.`,
																													"type": "object",
																												},
																											},
																											"type":                  "object",
																											"x-kubernetes-map-type": "atomic",
																										},
																										"storageClassName": map[string]interface{}{
																											"description": `storageClassName is the name of the StorageClass required by the claim.
More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1`,
																											"type": "string",
																										},
																										"volumeAttributesClassName": map[string]interface{}{
																											"description": `volumeAttributesClassName may be used to set the VolumeAttributesClass used by this claim.
If specified, the CSI driver will create or update the volume with the attributes defined
in the corresponding VolumeAttributesClass. This has a different purpose than storageClassName,
it can be changed after the claim is created. An empty string or nil value indicates that no
VolumeAttributesClass will be applied to the claim. If the claim enters an Infeasible error state,
this field can be reset to its previous value (including nil) to cancel the modification.
If the resource referred to by volumeAttributesClass does not exist, this PersistentVolumeClaim will be
set to a Pending state, as reflected by the modifyVolumeStatus field, until such as a resource
exists.
More info: https://kubernetes.io/docs/concepts/storage/volume-attributes-classes/`,
																											"type": "string",
																										},
																										"volumeMode": map[string]interface{}{
																											"description": `volumeMode defines what type of volume is required by the claim.
Value of Filesystem is implied when not included in claim spec.`,
																											"type": "string",
																										},
																										"volumeName": map[string]interface{}{
																											"description": "volumeName is the binding reference to the PersistentVolume backing this claim.",
																											"type":        "string",
																										},
																									},
																									"type": "object",
																								},
																							},
																							"required": []interface{}{
																								"spec",
																							},
																							"type": "object",
																						},
																					},
																					"type": "object",
																				},
																				"fc": map[string]interface{}{
																					"description": "fc represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.",
																					"properties": map[string]interface{}{
																						"fsType": map[string]interface{}{
																							"description": `fsType is the filesystem type to mount.
Must be a filesystem type supported by the host operating system.
Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.`,
																							"type": "string",
																						},
																						"lun": map[string]interface{}{
																							"description": "lun is Optional: FC target lun number",
																							"format":      "int32",
																							"type":        "integer",
																						},
																						"readOnly": map[string]interface{}{
																							"description": `readOnly is Optional: Defaults to false (read/write). ReadOnly here will force
the ReadOnly setting in VolumeMounts.`,
																							"type": "boolean",
																						},
																						"targetWWNs": map[string]interface{}{
																							"description": "targetWWNs is Optional: FC target worldwide names (WWNs)",
																							"items": map[string]interface{}{
																								"type": "string",
																							},
																							"type":                   "array",
																							"x-kubernetes-list-type": "atomic",
																						},
																						"wwids": map[string]interface{}{
																							"description": `wwids Optional: FC volume world wide identifiers (wwids)
Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.`,
																							"items": map[string]interface{}{
																								"type": "string",
																							},
																							"type":                   "array",
																							"x-kubernetes-list-type": "atomic",
																						},
																					},
																					"type": "object",
																				},
																				"flexVolume": map[string]interface{}{
																					"description": `flexVolume represents a generic volume resource that is
provisioned/attached using an exec based plugin.
Deprecated: FlexVolume is deprecated. Consider using a CSIDriver instead.`,
																					"properties": map[string]interface{}{
																						"driver": map[string]interface{}{
																							"description": "driver is the name of the driver to use for this volume.",
																							"type":        "string",
																						},
																						"fsType": map[string]interface{}{
																							"description": `fsType is the filesystem type to mount.
Must be a filesystem type supported by the host operating system.
Ex. "ext4", "xfs", "ntfs". The default filesystem depends on FlexVolume script.`,
																							"type": "string",
																						},
																						"options": map[string]interface{}{
																							"additionalProperties": map[string]interface{}{
																								"type": "string",
																							},
																							"description": "options is Optional: this field holds extra command options if any.",
																							"type":        "object",
																						},
																						"readOnly": map[string]interface{}{
																							"description": `readOnly is Optional: defaults to false (read/write). ReadOnly here will force
the ReadOnly setting in VolumeMounts.`,
																							"type": "boolean",
																						},
																						"secretRef": map[string]interface{}{
																							"description": `secretRef is Optional: secretRef is reference to the secret object containing
sensitive information to pass to the plugin scripts. This may be
empty if no secret object is specified. If the secret object
contains more than one secret, all secrets are passed to the plugin
scripts.`,
																							"properties": map[string]interface{}{
																								"name": map[string]interface{}{
																									"default": "",
																									"description": `Name of the referent.
This field is effectively required, but due to backwards compatibility is
allowed to be empty. Instances of this type with an empty value here are
almost certainly wrong.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names`,
																									"type": "string",
																								},
																							},
																							"type":                  "object",
																							"x-kubernetes-map-type": "atomic",
																						},
																					},
																					"required": []interface{}{
																						"driver",
																					},
																					"type": "object",
																				},
																				"flocker": map[string]interface{}{
																					"description": `flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running.
Deprecated: Flocker is deprecated and the in-tree flocker type is no longer supported.`,
																					"properties": map[string]interface{}{
																						"datasetName": map[string]interface{}{
																							"description": `datasetName is Name of the dataset stored as metadata -> name on the dataset for Flocker
should be considered as deprecated`,
																							"type": "string",
																						},
																						"datasetUUID": map[string]interface{}{
																							"description": "datasetUUID is the UUID of the dataset. This is unique identifier of a Flocker dataset",
																							"type":        "string",
																						},
																					},
																					"type": "object",
																				},
																				"gcePersistentDisk": map[string]interface{}{
																					"description": `gcePersistentDisk represents a GCE Disk resource that is attached to a
kubelet's host machine and then exposed to the pod.
Deprecated: GCEPersistentDisk is deprecated. All operations for the in-tree
gcePersistentDisk type are redirected to the pd.csi.storage.gke.io CSI driver.
More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk`,
																					"properties": map[string]interface{}{
																						"fsType": map[string]interface{}{
																							"description": `fsType is filesystem type of the volume that you want to mount.
Tip: Ensure that the filesystem type is supported by the host operating system.
Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk`,
																							"type": "string",
																						},
																						"partition": map[string]interface{}{
																							"description": `partition is the partition in the volume that you want to mount.
If omitted, the default is to mount by volume name.
Examples: For volume /dev/sda1, you specify the partition as "1".
Similarly, the volume partition for /dev/sda is "0" (or you can leave the property empty).
More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk`,
																							"format": "int32",
																							"type":   "integer",
																						},
																						"pdName": map[string]interface{}{
																							"description": `pdName is unique name of the PD resource in GCE. Used to identify the disk in GCE.
More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk`,
																							"type": "string",
																						},
																						"readOnly": map[string]interface{}{
																							"description": `readOnly here will force the ReadOnly setting in VolumeMounts.
Defaults to false.
More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk`,
																							"type": "boolean",
																						},
																					},
																					"required": []interface{}{
																						"pdName",
																					},
																					"type": "object",
																				},
																				"gitRepo": map[string]interface{}{
																					"description": `gitRepo represents a git repository at a particular revision.
Deprecated: GitRepo is deprecated. To provision a container with a git repo, mount an
EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir
into the Pod's container.`,
																					"properties": map[string]interface{}{
																						"directory": map[string]interface{}{
																							"description": `directory is the target directory name.
Must not contain or start with '..'.  If '.' is supplied, the volume directory will be the
git repository.  Otherwise, if specified, the volume will contain the git repository in
the subdirectory with the given name.`,
																							"type": "string",
																						},
																						"repository": map[string]interface{}{
																							"description": "repository is the URL",
																							"type":        "string",
																						},
																						"revision": map[string]interface{}{
																							"description": "revision is the commit hash for the specified revision.",
																							"type":        "string",
																						},
																					},
																					"required": []interface{}{
																						"repository",
																					},
																					"type": "object",
																				},
																				"glusterfs": map[string]interface{}{
																					"description": `glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime.
Deprecated: Glusterfs is deprecated and the in-tree glusterfs type is no longer supported.`,
																					"properties": map[string]interface{}{
																						"endpoints": map[string]interface{}{
																							"description": "endpoints is the endpoint name that details Glusterfs topology.",
																							"type":        "string",
																						},
																						"path": map[string]interface{}{
																							"description": `path is the Glusterfs volume path.
More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod`,
																							"type": "string",
																						},
																						"readOnly": map[string]interface{}{
																							"description": `readOnly here will force the Glusterfs volume to be mounted with read-only permissions.
Defaults to false.
More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod`,
																							"type": "boolean",
																						},
																					},
																					"required": []interface{}{
																						"endpoints",
																						"path",
																					},
																					"type": "object",
																				},
																				"hostPath": map[string]interface{}{
																					"description": `hostPath represents a pre-existing file or directory on the host
machine that is directly exposed to the container. This is generally
used for system agents or other privileged things that are allowed
to see the host machine. Most containers will NOT need this.
More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath`,
																					"properties": map[string]interface{}{
																						"path": map[string]interface{}{
																							"description": `path of the directory on the host.
If the path is a symlink, it will follow the link to the real path.
More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath`,
																							"type": "string",
																						},
																						"type": map[string]interface{}{
																							"description": `type for HostPath Volume
Defaults to ""
More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath`,
																							"type": "string",
																						},
																					},
																					"required": []interface{}{
																						"path",
																					},
																					"type": "object",
																				},
																				"image": map[string]interface{}{
																					"description": `image represents an OCI object (a container image or artifact) pulled and mounted on the kubelet's host machine.
The volume is resolved at pod startup depending on which PullPolicy value is provided:

- Always: the kubelet always attempts to pull the reference. Container creation will fail If the pull fails.
- Never: the kubelet never pulls the reference and only uses a local image or artifact. Container creation will fail if the reference isn't present.
- IfNotPresent: the kubelet pulls if the reference isn't already present on disk. Container creation will fail if the reference isn't present and the pull fails.

The volume gets re-resolved if the pod gets deleted and recreated, which means that new remote content will become available on pod recreation.
A failure to resolve or pull the image during pod startup will block containers from starting and may add significant latency. Failures will be retried using normal volume backoff and will be reported on the pod reason and message.
The types of objects that may be mounted by this volume are defined by the container runtime implementation on a host machine and at minimum must include all valid types supported by the container image field.
The OCI object gets mounted in a single directory (spec.containers[*].volumeMounts.mountPath) by merging the manifest layers in the same way as for container images.
The volume will be mounted read-only (ro).
Sub path mounts for containers are not supported (spec.containers[*].volumeMounts.subpath) before 1.33.
The field spec.securityContext.fsGroupChangePolicy has no effect on this volume type.`,
																					"properties": map[string]interface{}{
																						"pullPolicy": map[string]interface{}{
																							"description": `Policy for pulling OCI objects. Possible values are:
Always: the kubelet always attempts to pull the reference. Container creation will fail If the pull fails.
Never: the kubelet never pulls the reference and only uses a local image or artifact. Container creation will fail if the reference isn't present.
IfNotPresent: the kubelet pulls if the reference isn't already present on disk. Container creation will fail if the reference isn't present and the pull fails.
Defaults to Always if :latest tag is specified, or IfNotPresent otherwise.`,
																							"type": "string",
																						},
																						"reference": map[string]interface{}{
																							"description": `Required: Image or artifact reference to be used.
Behaves in the same way as pod.spec.containers[*].image.
Pull secrets will be assembled in the same way as for the container image by looking up node credentials, SA image pull secrets, and pod spec image pull secrets.
More info: https://kubernetes.io/docs/concepts/containers/images
This field is optional to allow higher level config management to default or override
container images in workload controllers like Deployments and StatefulSets.`,
																							"type": "string",
																						},
																					},
																					"type": "object",
																				},
																				"iscsi": map[string]interface{}{
																					"description": `iscsi represents an ISCSI Disk resource that is attached to a
kubelet's host machine and then exposed to the pod.
More info: https://kubernetes.io/docs/concepts/storage/volumes/#iscsi`,
																					"properties": map[string]interface{}{
																						"chapAuthDiscovery": map[string]interface{}{
																							"description": "chapAuthDiscovery defines whether support iSCSI Discovery CHAP authentication",
																							"type":        "boolean",
																						},
																						"chapAuthSession": map[string]interface{}{
																							"description": "chapAuthSession defines whether support iSCSI Session CHAP authentication",
																							"type":        "boolean",
																						},
																						"fsType": map[string]interface{}{
																							"description": `fsType is the filesystem type of the volume that you want to mount.
Tip: Ensure that the filesystem type is supported by the host operating system.
Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi`,
																							"type": "string",
																						},
																						"initiatorName": map[string]interface{}{
																							"description": `initiatorName is the custom iSCSI Initiator Name.
If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface
<target portal>:<volume name> will be created for the connection.`,
																							"type": "string",
																						},
																						"iqn": map[string]interface{}{
																							"description": "iqn is the target iSCSI Qualified Name.",
																							"type":        "string",
																						},
																						"iscsiInterface": map[string]interface{}{
																							"default": "default",
																							"description": `iscsiInterface is the interface Name that uses an iSCSI transport.
Defaults to 'default' (tcp).`,
																							"type": "string",
																						},
																						"lun": map[string]interface{}{
																							"description": "lun represents iSCSI Target Lun number.",
																							"format":      "int32",
																							"type":        "integer",
																						},
																						"portals": map[string]interface{}{
																							"description": `portals is the iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port
is other than default (typically TCP ports 860 and 3260).`,
																							"items": map[string]interface{}{
																								"type": "string",
																							},
																							"type":                   "array",
																							"x-kubernetes-list-type": "atomic",
																						},
																						"readOnly": map[string]interface{}{
																							"description": `readOnly here will force the ReadOnly setting in VolumeMounts.
Defaults to false.`,
																							"type": "boolean",
																						},
																						"secretRef": map[string]interface{}{
																							"description": "secretRef is the CHAP Secret for iSCSI target and initiator authentication",
																							"properties": map[string]interface{}{
																								"name": map[string]interface{}{
																									"default": "",
																									"description": `Name of the referent.
This field is effectively required, but due to backwards compatibility is
allowed to be empty. Instances of this type with an empty value here are
almost certainly wrong.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names`,
																									"type": "string",
																								},
																							},
																							"type":                  "object",
																							"x-kubernetes-map-type": "atomic",
																						},
																						"targetPortal": map[string]interface{}{
																							"description": `targetPortal is iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port
is other than default (typically TCP ports 860 and 3260).`,
																							"type": "string",
																						},
																					},
																					"required": []interface{}{
																						"iqn",
																						"lun",
																						"targetPortal",
																					},
																					"type": "object",
																				},
																				"name": map[string]interface{}{
																					"description": `name of the volume.
Must be a DNS_LABEL and unique within the pod.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names`,
																					"type": "string",
																				},
																				"nfs": map[string]interface{}{
																					"description": `nfs represents an NFS mount on the host that shares a pod's lifetime
More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs`,
																					"properties": map[string]interface{}{
																						"path": map[string]interface{}{
																							"description": `path that is exported by the NFS server.
More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs`,
																							"type": "string",
																						},
																						"readOnly": map[string]interface{}{
																							"description": `readOnly here will force the NFS export to be mounted with read-only permissions.
Defaults to false.
More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs`,
																							"type": "boolean",
																						},
																						"server": map[string]interface{}{
																							"description": `server is the hostname or IP address of the NFS server.
More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs`,
																							"type": "string",
																						},
																					},
																					"required": []interface{}{
																						"path",
																						"server",
																					},
																					"type": "object",
																				},
																				"persistentVolumeClaim": map[string]interface{}{
																					"description": `persistentVolumeClaimVolumeSource represents a reference to a
PersistentVolumeClaim in the same namespace.
More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims`,
																					"properties": map[string]interface{}{
																						"claimName": map[string]interface{}{
																							"description": `claimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume.
More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims`,
																							"type": "string",
																						},
																						"readOnly": map[string]interface{}{
																							"description": `readOnly Will force the ReadOnly setting in VolumeMounts.
Default false.`,
																							"type": "boolean",
																						},
																					},
																					"required": []interface{}{
																						"claimName",
																					},
																					"type": "object",
																				},
																				"photonPersistentDisk": map[string]interface{}{
																					"description": `photonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine.
Deprecated: PhotonPersistentDisk is deprecated and the in-tree photonPersistentDisk type is no longer supported.`,
																					"properties": map[string]interface{}{
																						"fsType": map[string]interface{}{
																							"description": `fsType is the filesystem type to mount.
Must be a filesystem type supported by the host operating system.
Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.`,
																							"type": "string",
																						},
																						"pdID": map[string]interface{}{
																							"description": "pdID is the ID that identifies Photon Controller persistent disk",
																							"type":        "string",
																						},
																					},
																					"required": []interface{}{
																						"pdID",
																					},
																					"type": "object",
																				},
																				"portworxVolume": map[string]interface{}{
																					"description": `portworxVolume represents a portworx volume attached and mounted on kubelets host machine.
Deprecated: PortworxVolume is deprecated. All operations for the in-tree portworxVolume type
are redirected to the pxd.portworx.com CSI driver.`,
																					"properties": map[string]interface{}{
																						"fsType": map[string]interface{}{
																							"description": `fSType represents the filesystem type to mount
Must be a filesystem type supported by the host operating system.
Ex. "ext4", "xfs". Implicitly inferred to be "ext4" if unspecified.`,
																							"type": "string",
																						},
																						"readOnly": map[string]interface{}{
																							"description": `readOnly defaults to false (read/write). ReadOnly here will force
the ReadOnly setting in VolumeMounts.`,
																							"type": "boolean",
																						},
																						"volumeID": map[string]interface{}{
																							"description": "volumeID uniquely identifies a Portworx volume",
																							"type":        "string",
																						},
																					},
																					"required": []interface{}{
																						"volumeID",
																					},
																					"type": "object",
																				},
																				"projected": map[string]interface{}{
																					"description": "projected items for all in one resources secrets, configmaps, and downward API",
																					"properties": map[string]interface{}{
																						"defaultMode": map[string]interface{}{
																							"description": `defaultMode are the mode bits used to set permissions on created files by default.
Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.
YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.
Directories within the path are not affected by this setting.
This might be in conflict with other options that affect the file
mode, like fsGroup, and the result can be other mode bits set.`,
																							"format": "int32",
																							"type":   "integer",
																						},
																						"sources": map[string]interface{}{
																							"description": `sources is the list of volume projections. Each entry in this list
handles one source.`,
																							"items": map[string]interface{}{
																								"description": `Projection that may be projected along with other supported volume types.
Exactly one of these fields must be set.`,
																								"properties": map[string]interface{}{
																									"clusterTrustBundle": map[string]interface{}{
																										"description": `ClusterTrustBundle allows a pod to access the ` + "`" + `.spec.trustBundle` + "`" + ` field
of ClusterTrustBundle objects in an auto-updating file.

Alpha, gated by the ClusterTrustBundleProjection feature gate.

ClusterTrustBundle objects can either be selected by name, or by the
combination of signer name and a label selector.

Kubelet performs aggressive normalization of the PEM contents written
into the pod filesystem.  Esoteric PEM features such as inter-block
comments and block headers are stripped.  Certificates are deduplicated.
The ordering of certificates within the file is arbitrary, and Kubelet
may change the order over time.`,
																										"properties": map[string]interface{}{
																											"labelSelector": map[string]interface{}{
																												"description": `Select all ClusterTrustBundles that match this label selector.  Only has
effect if signerName is set.  Mutually-exclusive with name.  If unset,
interpreted as "match nothing".  If set but empty, interpreted as "match
everything".`,
																												"properties": map[string]interface{}{
																													"matchExpressions": map[string]interface{}{
																														"description": "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																														"items": map[string]interface{}{
																															"description": `A label selector requirement is a selector that contains values, a key, and an operator that
relates the key and values.`,
																															"properties": map[string]interface{}{
																																"key": map[string]interface{}{
																																	"description": "key is the label key that the selector applies to.",
																																	"type":        "string",
																																},
																																"operator": map[string]interface{}{
																																	"description": `operator represents a key's relationship to a set of values.
Valid operators are In, NotIn, Exists and DoesNotExist.`,
																																	"type": "string",
																																},
																																"values": map[string]interface{}{
																																	"description": `values is an array of string values. If the operator is In or NotIn,
the values array must be non-empty. If the operator is Exists or DoesNotExist,
the values array must be empty. This array is replaced during a strategic
merge patch.`,
																																	"items": map[string]interface{}{
																																		"type": "string",
																																	},
																																	"type":                   "array",
																																	"x-kubernetes-list-type": "atomic",
																																},
																															},
																															"required": []interface{}{
																																"key",
																																"operator",
																															},
																															"type": "object",
																														},
																														"type":                   "array",
																														"x-kubernetes-list-type": "atomic",
																													},
																													"matchLabels": map[string]interface{}{
																														"additionalProperties": map[string]interface{}{
																															"type": "string",
																														},
																														"description": `matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels
map is equivalent to an element of matchExpressions, whose key field is "key", the
operator is "In", and the values array contains only "value". The requirements are ANDed.`,
																														"type": "object",
																													},
																												},
																												"type":                  "object",
																												"x-kubernetes-map-type": "atomic",
																											},
																											"name": map[string]interface{}{
																												"description": `Select a single ClusterTrustBundle by object name.  Mutually-exclusive
with signerName and labelSelector.`,
																												"type": "string",
																											},
																											"optional": map[string]interface{}{
																												"description": `If true, don't block pod startup if the referenced ClusterTrustBundle(s)
aren't available.  If using name, then the named ClusterTrustBundle is
allowed not to exist.  If using signerName, then the combination of
signerName and labelSelector is allowed to match zero
ClusterTrustBundles.`,
																												"type": "boolean",
																											},
																											"path": map[string]interface{}{
																												"description": "Relative path from the volume root to write the bundle.",
																												"type":        "string",
																											},
																											"signerName": map[string]interface{}{
																												"description": `Select all ClusterTrustBundles that match this signer name.
Mutually-exclusive with name.  The contents of all selected
ClusterTrustBundles will be unified and deduplicated.`,
																												"type": "string",
																											},
																										},
																										"required": []interface{}{
																											"path",
																										},
																										"type": "object",
																									},
																									"configMap": map[string]interface{}{
																										"description": "configMap information about the configMap data to project",
																										"properties": map[string]interface{}{
																											"items": map[string]interface{}{
																												"description": `items if unspecified, each key-value pair in the Data field of the referenced
ConfigMap will be projected into the volume as a file whose name is the
key and content is the value. If specified, the listed keys will be
projected into the specified paths, and unlisted keys will not be
present. If a key is specified which is not present in the ConfigMap,
the volume setup will error unless it is marked optional. Paths must be
relative and may not contain the '..' path or start with '..'.`,
																												"items": map[string]interface{}{
																													"description": "Maps a string key to a path within a volume.",
																													"properties": map[string]interface{}{
																														"key": map[string]interface{}{
																															"description": "key is the key to project.",
																															"type":        "string",
																														},
																														"mode": map[string]interface{}{
																															"description": `mode is Optional: mode bits used to set permissions on this file.
Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.
YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.
If not specified, the volume defaultMode will be used.
This might be in conflict with other options that affect the file
mode, like fsGroup, and the result can be other mode bits set.`,
																															"format": "int32",
																															"type":   "integer",
																														},
																														"path": map[string]interface{}{
																															"description": `path is the relative path of the file to map the key to.
May not be an absolute path.
May not contain the path element '..'.
May not start with the string '..'.`,
																															"type": "string",
																														},
																													},
																													"required": []interface{}{
																														"key",
																														"path",
																													},
																													"type": "object",
																												},
																												"type":                   "array",
																												"x-kubernetes-list-type": "atomic",
																											},
																											"name": map[string]interface{}{
																												"default": "",
																												"description": `Name of the referent.
This field is effectively required, but due to backwards compatibility is
allowed to be empty. Instances of this type with an empty value here are
almost certainly wrong.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names`,
																												"type": "string",
																											},
																											"optional": map[string]interface{}{
																												"description": "optional specify whether the ConfigMap or its keys must be defined",
																												"type":        "boolean",
																											},
																										},
																										"type":                  "object",
																										"x-kubernetes-map-type": "atomic",
																									},
																									"downwardAPI": map[string]interface{}{
																										"description": "downwardAPI information about the downwardAPI data to project",
																										"properties": map[string]interface{}{
																											"items": map[string]interface{}{
																												"description": "Items is a list of DownwardAPIVolume file",
																												"items": map[string]interface{}{
																													"description": "DownwardAPIVolumeFile represents information to create the file containing the pod field",
																													"properties": map[string]interface{}{
																														"fieldRef": map[string]interface{}{
																															"description": "Required: Selects a field of the pod: only annotations, labels, name, namespace and uid are supported.",
																															"properties": map[string]interface{}{
																																"apiVersion": map[string]interface{}{
																																	"description": "Version of the schema the FieldPath is written in terms of, defaults to \"v1\".",
																																	"type":        "string",
																																},
																																"fieldPath": map[string]interface{}{
																																	"description": "Path of the field to select in the specified API version.",
																																	"type":        "string",
																																},
																															},
																															"required": []interface{}{
																																"fieldPath",
																															},
																															"type":                  "object",
																															"x-kubernetes-map-type": "atomic",
																														},
																														"mode": map[string]interface{}{
																															"description": `Optional: mode bits used to set permissions on this file, must be an octal value
between 0000 and 0777 or a decimal value between 0 and 511.
YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.
If not specified, the volume defaultMode will be used.
This might be in conflict with other options that affect the file
mode, like fsGroup, and the result can be other mode bits set.`,
																															"format": "int32",
																															"type":   "integer",
																														},
																														"path": map[string]interface{}{
																															"description": "Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",
																															"type":        "string",
																														},
																														"resourceFieldRef": map[string]interface{}{
																															"description": `Selects a resource of the container: only resources limits and requests
(limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.`,
																															"properties": map[string]interface{}{
																																"containerName": map[string]interface{}{
																																	"description": "Container name: required for volumes, optional for env vars",
																																	"type":        "string",
																																},
																																"divisor": map[string]interface{}{
																																	"anyOf": []interface{}{
																																		map[string]interface{}{
																																			"type": "integer",
																																		},
																																		map[string]interface{}{
																																			"type": "string",
																																		},
																																	},
																																	"description":                "Specifies the output format of the exposed resources, defaults to \"1\"",
																																	"pattern":                    `^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$`,
																																	"x-kubernetes-int-or-string": true,
																																},
																																"resource": map[string]interface{}{
																																	"description": "Required: resource to select",
																																	"type":        "string",
																																},
																															},
																															"required": []interface{}{
																																"resource",
																															},
																															"type":                  "object",
																															"x-kubernetes-map-type": "atomic",
																														},
																													},
																													"required": []interface{}{
																														"path",
																													},
																													"type": "object",
																												},
																												"type":                   "array",
																												"x-kubernetes-list-type": "atomic",
																											},
																										},
																										"type": "object",
																									},
																									"podCertificate": map[string]interface{}{
																										"description": `Projects an auto-rotating credential bundle (private key and certificate
chain) that the pod can use either as a TLS client or server.

Kubelet generates a private key and uses it to send a
PodCertificateRequest to the named signer.  Once the signer approves the
request and issues a certificate chain, Kubelet writes the key and
certificate chain to the pod filesystem.  The pod does not start until
certificates have been issued for each podCertificate projected volume
source in its spec.

Kubelet will begin trying to rotate the certificate at the time indicated
by the signer using the PodCertificateRequest.Status.BeginRefreshAt
timestamp.

Kubelet can write a single file, indicated by the credentialBundlePath
field, or separate files, indicated by the keyPath and
certificateChainPath fields.

The credential bundle is a single file in PEM format.  The first PEM
entry is the private key (in PKCS#8 format), and the remaining PEM
entries are the certificate chain issued by the signer (typically,
signers will return their certificate chain in leaf-to-root order).

Prefer using the credential bundle format, since your application code
can read it atomically.  If you use keyPath and certificateChainPath,
your application must make two separate file reads. If these coincide
with a certificate rotation, it is possible that the private key and leaf
certificate you read may not correspond to each other.  Your application
will need to check for this condition, and re-read until they are
consistent.

The named signer controls chooses the format of the certificate it
issues; consult the signer implementation's documentation to learn how to
use the certificates it issues.`,
																										"properties": map[string]interface{}{
																											"certificateChainPath": map[string]interface{}{
																												"description": `Write the certificate chain at this path in the projected volume.

Most applications should use credentialBundlePath.  When using keyPath
and certificateChainPath, your application needs to check that the key
and leaf certificate are consistent, because it is possible to read the
files mid-rotation.`,
																												"type": "string",
																											},
																											"credentialBundlePath": map[string]interface{}{
																												"description": `Write the credential bundle at this path in the projected volume.

The credential bundle is a single file that contains multiple PEM blocks.
The first PEM block is a PRIVATE KEY block, containing a PKCS#8 private
key.

The remaining blocks are CERTIFICATE blocks, containing the issued
certificate chain from the signer (leaf and any intermediates).

Using credentialBundlePath lets your Pod's application code make a single
atomic read that retrieves a consistent key and certificate chain.  If you
project them to separate files, your application code will need to
additionally check that the leaf certificate was issued to the key.`,
																												"type": "string",
																											},
																											"keyPath": map[string]interface{}{
																												"description": `Write the key at this path in the projected volume.

Most applications should use credentialBundlePath.  When using keyPath
and certificateChainPath, your application needs to check that the key
and leaf certificate are consistent, because it is possible to read the
files mid-rotation.`,
																												"type": "string",
																											},
																											"keyType": map[string]interface{}{
																												"description": `The type of keypair Kubelet will generate for the pod.

Valid values are "RSA3072", "RSA4096", "ECDSAP256", "ECDSAP384",
"ECDSAP521", and "ED25519".`,
																												"type": "string",
																											},
																											"maxExpirationSeconds": map[string]interface{}{
																												"description": `maxExpirationSeconds is the maximum lifetime permitted for the
certificate.

Kubelet copies this value verbatim into the PodCertificateRequests it
generates for this projection.

If omitted, kube-apiserver will set it to 86400(24 hours). kube-apiserver
will reject values shorter than 3600 (1 hour).  The maximum allowable
value is 7862400 (91 days).

The signer implementation is then free to issue a certificate with any
lifetime *shorter* than MaxExpirationSeconds, but no shorter than 3600
seconds (1 hour).  This constraint is enforced by kube-apiserver.
` + "`" + `kubernetes.io` + "`" + ` signers will never issue certificates with a lifetime
longer than 24 hours.`,
																												"format": "int32",
																												"type":   "integer",
																											},
																											"signerName": map[string]interface{}{
																												"description": "Kubelet's generated CSRs will be addressed to this signer.",
																												"type":        "string",
																											},
																											"userAnnotations": map[string]interface{}{
																												"additionalProperties": map[string]interface{}{
																													"type": "string",
																												},
																												"description": `userAnnotations allow pod authors to pass additional information to
the signer implementation.  Kubernetes does not restrict or validate this
metadata in any way.

These values are copied verbatim into the ` + "`" + `spec.unverifiedUserAnnotations` + "`" + ` field of
the PodCertificateRequest objects that Kubelet creates.

Entries are subject to the same validation as object metadata annotations,
with the addition that all keys must be domain-prefixed. No restrictions
are placed on values, except an overall size limitation on the entire field.

Signers should document the keys and values they support. Signers should
deny requests that contain keys they do not recognize.`,
																												"type": "object",
																											},
																										},
																										"required": []interface{}{
																											"keyType",
																											"signerName",
																										},
																										"type": "object",
																									},
																									"secret": map[string]interface{}{
																										"description": "secret information about the secret data to project",
																										"properties": map[string]interface{}{
																											"items": map[string]interface{}{
																												"description": `items if unspecified, each key-value pair in the Data field of the referenced
Secret will be projected into the volume as a file whose name is the
key and content is the value. If specified, the listed keys will be
projected into the specified paths, and unlisted keys will not be
present. If a key is specified which is not present in the Secret,
the volume setup will error unless it is marked optional. Paths must be
relative and may not contain the '..' path or start with '..'.`,
																												"items": map[string]interface{}{
																													"description": "Maps a string key to a path within a volume.",
																													"properties": map[string]interface{}{
																														"key": map[string]interface{}{
																															"description": "key is the key to project.",
																															"type":        "string",
																														},
																														"mode": map[string]interface{}{
																															"description": `mode is Optional: mode bits used to set permissions on this file.
Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.
YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.
If not specified, the volume defaultMode will be used.
This might be in conflict with other options that affect the file
mode, like fsGroup, and the result can be other mode bits set.`,
																															"format": "int32",
																															"type":   "integer",
																														},
																														"path": map[string]interface{}{
																															"description": `path is the relative path of the file to map the key to.
May not be an absolute path.
May not contain the path element '..'.
May not start with the string '..'.`,
																															"type": "string",
																														},
																													},
																													"required": []interface{}{
																														"key",
																														"path",
																													},
																													"type": "object",
																												},
																												"type":                   "array",
																												"x-kubernetes-list-type": "atomic",
																											},
																											"name": map[string]interface{}{
																												"default": "",
																												"description": `Name of the referent.
This field is effectively required, but due to backwards compatibility is
allowed to be empty. Instances of this type with an empty value here are
almost certainly wrong.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names`,
																												"type": "string",
																											},
																											"optional": map[string]interface{}{
																												"description": "optional field specify whether the Secret or its key must be defined",
																												"type":        "boolean",
																											},
																										},
																										"type":                  "object",
																										"x-kubernetes-map-type": "atomic",
																									},
																									"serviceAccountToken": map[string]interface{}{
																										"description": "serviceAccountToken is information about the serviceAccountToken data to project",
																										"properties": map[string]interface{}{
																											"audience": map[string]interface{}{
																												"description": `audience is the intended audience of the token. A recipient of a token
must identify itself with an identifier specified in the audience of the
token, and otherwise should reject the token. The audience defaults to the
identifier of the apiserver.`,
																												"type": "string",
																											},
																											"expirationSeconds": map[string]interface{}{
																												"description": `expirationSeconds is the requested duration of validity of the service
account token. As the token approaches expiration, the kubelet volume
plugin will proactively rotate the service account token. The kubelet will
start trying to rotate the token if the token is older than 80 percent of
its time to live or if the token is older than 24 hours.Defaults to 1 hour
and must be at least 10 minutes.`,
																												"format": "int64",
																												"type":   "integer",
																											},
																											"path": map[string]interface{}{
																												"description": `path is the path relative to the mount point of the file to project the
token into.`,
																												"type": "string",
																											},
																										},
																										"required": []interface{}{
																											"path",
																										},
																										"type": "object",
																									},
																								},
																								"type": "object",
																							},
																							"type":                   "array",
																							"x-kubernetes-list-type": "atomic",
																						},
																					},
																					"type": "object",
																				},
																				"quobyte": map[string]interface{}{
																					"description": `quobyte represents a Quobyte mount on the host that shares a pod's lifetime.
Deprecated: Quobyte is deprecated and the in-tree quobyte type is no longer supported.`,
																					"properties": map[string]interface{}{
																						"group": map[string]interface{}{
																							"description": `group to map volume access to
Default is no group`,
																							"type": "string",
																						},
																						"readOnly": map[string]interface{}{
																							"description": `readOnly here will force the Quobyte volume to be mounted with read-only permissions.
Defaults to false.`,
																							"type": "boolean",
																						},
																						"registry": map[string]interface{}{
																							"description": `registry represents a single or multiple Quobyte Registry services
specified as a string as host:port pair (multiple entries are separated with commas)
which acts as the central registry for volumes`,
																							"type": "string",
																						},
																						"tenant": map[string]interface{}{
																							"description": `tenant owning the given Quobyte volume in the Backend
Used with dynamically provisioned Quobyte volumes, value is set by the plugin`,
																							"type": "string",
																						},
																						"user": map[string]interface{}{
																							"description": `user to map volume access to
Defaults to serivceaccount user`,
																							"type": "string",
																						},
																						"volume": map[string]interface{}{
																							"description": "volume is a string that references an already created Quobyte volume by name.",
																							"type":        "string",
																						},
																					},
																					"required": []interface{}{
																						"registry",
																						"volume",
																					},
																					"type": "object",
																				},
																				"rbd": map[string]interface{}{
																					"description": `rbd represents a Rados Block Device mount on the host that shares a pod's lifetime.
Deprecated: RBD is deprecated and the in-tree rbd type is no longer supported.`,
																					"properties": map[string]interface{}{
																						"fsType": map[string]interface{}{
																							"description": `fsType is the filesystem type of the volume that you want to mount.
Tip: Ensure that the filesystem type is supported by the host operating system.
Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd`,
																							"type": "string",
																						},
																						"image": map[string]interface{}{
																							"description": `image is the rados image name.
More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it`,
																							"type": "string",
																						},
																						"keyring": map[string]interface{}{
																							"default": "/etc/ceph/keyring",
																							"description": `keyring is the path to key ring for RBDUser.
Default is /etc/ceph/keyring.
More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it`,
																							"type": "string",
																						},
																						"monitors": map[string]interface{}{
																							"description": `monitors is a collection of Ceph monitors.
More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it`,
																							"items": map[string]interface{}{
																								"type": "string",
																							},
																							"type":                   "array",
																							"x-kubernetes-list-type": "atomic",
																						},
																						"pool": map[string]interface{}{
																							"default": "rbd",
																							"description": `pool is the rados pool name.
Default is rbd.
More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it`,
																							"type": "string",
																						},
																						"readOnly": map[string]interface{}{
																							"description": `readOnly here will force the ReadOnly setting in VolumeMounts.
Defaults to false.
More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it`,
																							"type": "boolean",
																						},
																						"secretRef": map[string]interface{}{
																							"description": `secretRef is name of the authentication secret for RBDUser. If provided
overrides keyring.
Default is nil.
More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it`,
																							"properties": map[string]interface{}{
																								"name": map[string]interface{}{
																									"default": "",
																									"description": `Name of the referent.
This field is effectively required, but due to backwards compatibility is
allowed to be empty. Instances of this type with an empty value here are
almost certainly wrong.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names`,
																									"type": "string",
																								},
																							},
																							"type":                  "object",
																							"x-kubernetes-map-type": "atomic",
																						},
																						"user": map[string]interface{}{
																							"default": "admin",
																							"description": `user is the rados user name.
Default is admin.
More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it`,
																							"type": "string",
																						},
																					},
																					"required": []interface{}{
																						"image",
																						"monitors",
																					},
																					"type": "object",
																				},
																				"scaleIO": map[string]interface{}{
																					"description": `scaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.
Deprecated: ScaleIO is deprecated and the in-tree scaleIO type is no longer supported.`,
																					"properties": map[string]interface{}{
																						"fsType": map[string]interface{}{
																							"default": "xfs",
																							"description": `fsType is the filesystem type to mount.
Must be a filesystem type supported by the host operating system.
Ex. "ext4", "xfs", "ntfs".
Default is "xfs".`,
																							"type": "string",
																						},
																						"gateway": map[string]interface{}{
																							"description": "gateway is the host address of the ScaleIO API Gateway.",
																							"type":        "string",
																						},
																						"protectionDomain": map[string]interface{}{
																							"description": "protectionDomain is the name of the ScaleIO Protection Domain for the configured storage.",
																							"type":        "string",
																						},
																						"readOnly": map[string]interface{}{
																							"description": `readOnly Defaults to false (read/write). ReadOnly here will force
the ReadOnly setting in VolumeMounts.`,
																							"type": "boolean",
																						},
																						"secretRef": map[string]interface{}{
																							"description": `secretRef references to the secret for ScaleIO user and other
sensitive information. If this is not provided, Login operation will fail.`,
																							"properties": map[string]interface{}{
																								"name": map[string]interface{}{
																									"default": "",
																									"description": `Name of the referent.
This field is effectively required, but due to backwards compatibility is
allowed to be empty. Instances of this type with an empty value here are
almost certainly wrong.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names`,
																									"type": "string",
																								},
																							},
																							"type":                  "object",
																							"x-kubernetes-map-type": "atomic",
																						},
																						"sslEnabled": map[string]interface{}{
																							"description": "sslEnabled Flag enable/disable SSL communication with Gateway, default false",
																							"type":        "boolean",
																						},
																						"storageMode": map[string]interface{}{
																							"default": "ThinProvisioned",
																							"description": `storageMode indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned.
Default is ThinProvisioned.`,
																							"type": "string",
																						},
																						"storagePool": map[string]interface{}{
																							"description": "storagePool is the ScaleIO Storage Pool associated with the protection domain.",
																							"type":        "string",
																						},
																						"system": map[string]interface{}{
																							"description": "system is the name of the storage system as configured in ScaleIO.",
																							"type":        "string",
																						},
																						"volumeName": map[string]interface{}{
																							"description": `volumeName is the name of a volume already created in the ScaleIO system
that is associated with this volume source.`,
																							"type": "string",
																						},
																					},
																					"required": []interface{}{
																						"gateway",
																						"secretRef",
																						"system",
																					},
																					"type": "object",
																				},
																				"secret": map[string]interface{}{
																					"description": `secret represents a secret that should populate this volume.
More info: https://kubernetes.io/docs/concepts/storage/volumes#secret`,
																					"properties": map[string]interface{}{
																						"defaultMode": map[string]interface{}{
																							"description": `defaultMode is Optional: mode bits used to set permissions on created files by default.
Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.
YAML accepts both octal and decimal values, JSON requires decimal values
for mode bits. Defaults to 0644.
Directories within the path are not affected by this setting.
This might be in conflict with other options that affect the file
mode, like fsGroup, and the result can be other mode bits set.`,
																							"format": "int32",
																							"type":   "integer",
																						},
																						"items": map[string]interface{}{
																							"description": `items If unspecified, each key-value pair in the Data field of the referenced
Secret will be projected into the volume as a file whose name is the
key and content is the value. If specified, the listed keys will be
projected into the specified paths, and unlisted keys will not be
present. If a key is specified which is not present in the Secret,
the volume setup will error unless it is marked optional. Paths must be
relative and may not contain the '..' path or start with '..'.`,
																							"items": map[string]interface{}{
																								"description": "Maps a string key to a path within a volume.",
																								"properties": map[string]interface{}{
																									"key": map[string]interface{}{
																										"description": "key is the key to project.",
																										"type":        "string",
																									},
																									"mode": map[string]interface{}{
																										"description": `mode is Optional: mode bits used to set permissions on this file.
Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.
YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.
If not specified, the volume defaultMode will be used.
This might be in conflict with other options that affect the file
mode, like fsGroup, and the result can be other mode bits set.`,
																										"format": "int32",
																										"type":   "integer",
																									},
																									"path": map[string]interface{}{
																										"description": `path is the relative path of the file to map the key to.
May not be an absolute path.
May not contain the path element '..'.
May not start with the string '..'.`,
																										"type": "string",
																									},
																								},
																								"required": []interface{}{
																									"key",
																									"path",
																								},
																								"type": "object",
																							},
																							"type":                   "array",
																							"x-kubernetes-list-type": "atomic",
																						},
																						"optional": map[string]interface{}{
																							"description": "optional field specify whether the Secret or its keys must be defined",
																							"type":        "boolean",
																						},
																						"secretName": map[string]interface{}{
																							"description": `secretName is the name of the secret in the pod's namespace to use.
More info: https://kubernetes.io/docs/concepts/storage/volumes#secret`,
																							"type": "string",
																						},
																					},
																					"type": "object",
																				},
																				"storageos": map[string]interface{}{
																					"description": `storageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.
Deprecated: StorageOS is deprecated and the in-tree storageos type is no longer supported.`,
																					"properties": map[string]interface{}{
																						"fsType": map[string]interface{}{
																							"description": `fsType is the filesystem type to mount.
Must be a filesystem type supported by the host operating system.
Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.`,
																							"type": "string",
																						},
																						"readOnly": map[string]interface{}{
																							"description": `readOnly defaults to false (read/write). ReadOnly here will force
the ReadOnly setting in VolumeMounts.`,
																							"type": "boolean",
																						},
																						"secretRef": map[string]interface{}{
																							"description": `secretRef specifies the secret to use for obtaining the StorageOS API
credentials.  If not specified, default values will be attempted.`,
																							"properties": map[string]interface{}{
																								"name": map[string]interface{}{
																									"default": "",
																									"description": `Name of the referent.
This field is effectively required, but due to backwards compatibility is
allowed to be empty. Instances of this type with an empty value here are
almost certainly wrong.
More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names`,
																									"type": "string",
																								},
																							},
																							"type":                  "object",
																							"x-kubernetes-map-type": "atomic",
																						},
																						"volumeName": map[string]interface{}{
																							"description": `volumeName is the human-readable name of the StorageOS volume.  Volume
names are only unique within a namespace.`,
																							"type": "string",
																						},
																						"volumeNamespace": map[string]interface{}{
																							"description": `volumeNamespace specifies the scope of the volume within StorageOS.  If no
namespace is specified then the Pod's namespace will be used.  This allows the
Kubernetes name scoping to be mirrored within StorageOS for tighter integration.
Set VolumeName to any name to override the default behaviour.
Set to "default" if you are not using namespaces within StorageOS.
Namespaces that do not pre-exist within StorageOS will be created.`,
																							"type": "string",
																						},
																					},
																					"type": "object",
																				},
																				"vsphereVolume": map[string]interface{}{
																					"description": `vsphereVolume represents a vSphere volume attached and mounted on kubelets host machine.
Deprecated: VsphereVolume is deprecated. All operations for the in-tree vsphereVolume type
are redirected to the csi.vsphere.vmware.com CSI driver.`,
																					"properties": map[string]interface{}{
																						"fsType": map[string]interface{}{
																							"description": `fsType is filesystem type to mount.
Must be a filesystem type supported by the host operating system.
Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.`,
																							"type": "string",
																						},
																						"storagePolicyID": map[string]interface{}{
																							"description": "storagePolicyID is the storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName.",
																							"type":        "string",
																						},
																						"storagePolicyName": map[string]interface{}{
																							"description": "storagePolicyName is the storage Policy Based Management (SPBM) profile name.",
																							"type":        "string",
																						},
																						"volumePath": map[string]interface{}{
																							"description": "volumePath is the path that identifies vSphere volume vmdk",
																							"type":        "string",
																						},
																					},
																					"required": []interface{}{
																						"volumePath",
																					},
																					"type": "object",
																				},
																			},
																			"required": []interface{}{
																				"name",
																			},
																			"type": "object",
																		},
																		"type": "array",
																		"x-kubernetes-list-map-keys": []interface{}{
																			"name",
																		},
																		"x-kubernetes-list-type": "map",
																	},
																},
																"required": []interface{}{
																	"containers",
																},
																"type": "object",
															},
														},
														"type": "object",
													},
													"ttlSecondsAfterFinished": map[string]interface{}{
														"description": `ttlSecondsAfterFinished limits the lifetime of a Job that has finished
execution (either Complete or Failed). If this field is set,
ttlSecondsAfterFinished after the Job finishes, it is eligible to be
automatically deleted. When the Job is being deleted, its lifecycle
guarantees (e.g. finalizers) will be honored. If this field is unset,
the Job won't be automatically deleted. If this field is set to zero,
the Job becomes eligible to be deleted immediately after it finishes.`,
														"format": "int32",
														"type":   "integer",
													},
												},
												"required": []interface{}{
													"template",
												},
												"type": "object",
											},
											"maxReplicaCount": map[string]interface{}{
												"format":  "int32",
												"minimum": 1,
												"type":    "integer",
											},
											"minReplicaCount": map[string]interface{}{
												"format":  "int32",
												"minimum": 0,
												"type":    "integer",
											},
											"pollingInterval": map[string]interface{}{
												"format":  "int32",
												"minimum": 1,
												"type":    "integer",
											},
											"rollout": map[string]interface{}{
												"description": "Rollout defines the strategy for job rollouts",
												"properties": map[string]interface{}{
													"propagationPolicy": map[string]interface{}{
														"enum": []interface{}{
															"foreground",
															"background",
														},
														"type": "string",
													},
													"strategy": map[string]interface{}{
														"enum": []interface{}{
															"gradual",
															"immediate",
														},
														"type": "string",
													},
												},
												"type": "object",
											},
											"rolloutStrategy": map[string]interface{}{
												"description": "Deprecated: Use Rollout.Strategy instead (see https://github.com/kedacore/keda/issues/3596).",
												"enum": []interface{}{
													"gradual",
													"immediate",
												},
												"type": "string",
											},
											"scalingStrategy": map[string]interface{}{
												"description": "ScalingStrategy defines the strategy of Scaling",
												"properties": map[string]interface{}{
													"customScalingQueueLengthDeduction": map[string]interface{}{
														"format": "int32",
														"type":   "integer",
													},
													"customScalingRunningJobPercentage": map[string]interface{}{
														"type": "string",
													},
													"multipleScalersCalculation": map[string]interface{}{
														"enum": []interface{}{
															"min",
															"avg",
															"sum",
															"max",
														},
														"type": "string",
													},
													"pendingPodConditions": map[string]interface{}{
														"items": map[string]interface{}{
															"type": "string",
														},
														"type": "array",
													},
													"strategy": map[string]interface{}{
														"enum": []interface{}{
															"default",
															"custom",
															"accurate",
															"eager",
														},
														"type": "string",
													},
												},
												"type": "object",
											},
											"successfulJobsHistoryLimit": map[string]interface{}{
												"format":  "int32",
												"minimum": 0,
												"type":    "integer",
											},
											"triggers": map[string]interface{}{
												"items": map[string]interface{}{
													"description": "ScaleTriggers reference the scaler that will be used",
													"properties": map[string]interface{}{
														"authenticationRef": map[string]interface{}{
															"description": `AuthenticationRef points to the TriggerAuthentication or ClusterTriggerAuthentication object that
is used to authenticate the scaler with the environment`,
															"properties": map[string]interface{}{
																"kind": map[string]interface{}{
																	"description": "Kind of the resource being referred to. Defaults to TriggerAuthentication.",
																	"enum": []interface{}{
																		"TriggerAuthentication",
																		"ClusterTriggerAuthentication",
																	},
																	"type": "string",
																},
																"name": map[string]interface{}{
																	"type": "string",
																},
															},
															"required": []interface{}{
																"name",
															},
															"type": "object",
														},
														"metadata": map[string]interface{}{
															"additionalProperties": map[string]interface{}{
																"type": "string",
															},
															"type": "object",
														},
														"name": map[string]interface{}{
															"type": "string",
														},
														"type": map[string]interface{}{
															"minLength": 1,
															"type":      "string",
														},
														"useCachedMetrics": map[string]interface{}{
															"type": "boolean",
														},
													},
													"required": []interface{}{
														"metadata",
														"type",
													},
													"type": "object",
												},
												"minItems": 1,
												"type":     "array",
											},
										},
										"required": []interface{}{
											"jobTargetRef",
											"triggers",
										},
										"type": "object",
									},
									"status": map[string]interface{}{
										"description": "ScaledJobStatus defines the observed state of ScaledJob",
										"properties": map[string]interface{}{
											"Paused": map[string]interface{}{
												"type": "string",
											},
											"authenticationsTypes": map[string]interface{}{
												"type": "string",
											},
											"conditions": map[string]interface{}{
												"description": "Conditions an array representation to store multiple Conditions",
												"items": map[string]interface{}{
													"description": "Condition to store the condition state",
													"properties": map[string]interface{}{
														"message": map[string]interface{}{
															"description": "A human readable message indicating details about the transition.",
															"type":        "string",
														},
														"reason": map[string]interface{}{
															"description": "The reason for the condition's last transition.",
															"type":        "string",
														},
														"status": map[string]interface{}{
															"description": "Status of the condition, one of True, False, Unknown.",
															"type":        "string",
														},
														"type": map[string]interface{}{
															"description": "Type of condition",
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
											"externalMetricNames": map[string]interface{}{
												"items": map[string]interface{}{
													"type": "string",
												},
												"type": "array",
											},
											"lastActiveTime": map[string]interface{}{
												"format": "date-time",
												"type":   "string",
											},
											"triggersActivity": map[string]interface{}{
												"additionalProperties": map[string]interface{}{
													"description": "TriggerActivityStatus represents the activity status of an external trigger",
													"properties": map[string]interface{}{
														"isActive": map[string]interface{}{
															"type": "boolean",
														},
													},
													"type": "object",
												},
												"type": "object",
											},
											"triggersTypes": map[string]interface{}{
												"type": "string",
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
			},
		},
	}

	return mutate.MutateCRDScaledjobsKedaSh(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;create;update;patch;delete

// CreateCRDScaledobjectsKedaSh creates the CustomResourceDefinition resource with name scaledobjects.keda.sh.
func CreateCRDScaledobjectsKedaSh(
	parent *platformv1alpha1.AutoscalingComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					"controller-gen.kubebuilder.io/version": "v0.21.0",
				},
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "keda-operator",
					"app.kubernetes.io/component":   "operator",
					"app.kubernetes.io/part-of":     "keda-operator",
					"app.kubernetes.io/version":     parent.Spec.Keda.Version, //  controlled by field: keda.version
					"platform.nukleros.io/category": "autoscaling",
					"platform.nukleros.io/project":  "keda",
				},
				"name": "scaledobjects.keda.sh",
			},
			"spec": map[string]interface{}{
				"group": "keda.sh",
				"names": map[string]interface{}{
					"kind":     "ScaledObject",
					"listKind": "ScaledObjectList",
					"plural":   "scaledobjects",
					"shortNames": []interface{}{
						"so",
					},
					"singular": "scaledobject",
				},
				"scope": "Namespaced",
				"versions": []interface{}{
					map[string]interface{}{
						"additionalPrinterColumns": []interface{}{
							map[string]interface{}{
								"jsonPath": ".status.scaleTargetKind",
								"name":     "ScaleTargetKind",
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".spec.scaleTargetRef.name",
								"name":     "ScaleTargetName",
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".spec.minReplicaCount",
								"name":     "Min",
								"type":     "integer",
							},
							map[string]interface{}{
								"jsonPath": ".spec.maxReplicaCount",
								"name":     "Max",
								"type":     "integer",
							},
							map[string]interface{}{
								"jsonPath": ".status.conditions[?(@.type==\"Ready\")].status",
								"name":     "Ready",
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".status.conditions[?(@.type==\"Active\")].status",
								"name":     "Active",
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".status.conditions[?(@.type==\"Fallback\")].status",
								"name":     "Fallback",
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".status.conditions[?(@.type==\"Paused\")].status",
								"name":     "Paused",
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".status.triggersTypes",
								"name":     "Triggers",
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".status.authenticationsTypes",
								"name":     "Authentications",
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".metadata.creationTimestamp",
								"name":     "Age",
								"type":     "date",
							},
						},
						"name": "v1alpha1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"description": "ScaledObject is a specification for a ScaledObject resource",
								"properties": map[string]interface{}{
									"apiVersion": map[string]interface{}{
										"description": `APIVersion defines the versioned schema of this representation of an object.
Servers should convert recognized schemas to the latest internal value, and
may reject unrecognized values.
More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources`,
										"type": "string",
									},
									"kind": map[string]interface{}{
										"description": `Kind is a string value representing the REST resource this object represents.
Servers may infer this from the endpoint the client submits requests to.
Cannot be updated.
In CamelCase.
More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds`,
										"type": "string",
									},
									"metadata": map[string]interface{}{
										"type": "object",
									},
									"spec": map[string]interface{}{
										"description": "ScaledObjectSpec is the spec for a ScaledObject resource",
										"properties": map[string]interface{}{
											"advanced": map[string]interface{}{
												"description": "AdvancedConfig specifies advance scaling options",
												"properties": map[string]interface{}{
													"horizontalPodAutoscalerConfig": map[string]interface{}{
														"description": "HorizontalPodAutoscalerConfig specifies horizontal scale config",
														"properties": map[string]interface{}{
															"behavior": map[string]interface{}{
																"description": `HorizontalPodAutoscalerBehavior configures the scaling behavior of the target
in both Up and Down directions (scaleUp and scaleDown fields respectively).`,
																"properties": map[string]interface{}{
																	"scaleDown": map[string]interface{}{
																		"description": `scaleDown is scaling policy for scaling Down.
If not set, the default value is to allow to scale down to minReplicas pods, with a
300 second stabilization window (i.e., the highest recommendation for
the last 300sec is used).`,
																		"properties": map[string]interface{}{
																			"policies": map[string]interface{}{
																				"description": `policies is a list of potential scaling polices which can be used during scaling.
If not set, use the default values:
- For scale up: allow doubling the number of pods, or an absolute change of 4 pods in a 15s window.
- For scale down: allow all pods to be removed in a 15s window.`,
																				"items": map[string]interface{}{
																					"description": "HPAScalingPolicy is a single policy which must hold true for a specified past interval.",
																					"properties": map[string]interface{}{
																						"periodSeconds": map[string]interface{}{
																							"description": `periodSeconds specifies the window of time for which the policy should hold true.
PeriodSeconds must be greater than zero and less than or equal to 1800 (30 min).`,
																							"format": "int32",
																							"type":   "integer",
																						},
																						"type": map[string]interface{}{
																							"description": "type is used to specify the scaling policy.",
																							"type":        "string",
																						},
																						"value": map[string]interface{}{
																							"description": `value contains the amount of change which is permitted by the policy.
It must be greater than zero`,
																							"format": "int32",
																							"type":   "integer",
																						},
																					},
																					"required": []interface{}{
																						"periodSeconds",
																						"type",
																						"value",
																					},
																					"type": "object",
																				},
																				"type":                   "array",
																				"x-kubernetes-list-type": "atomic",
																			},
																			"selectPolicy": map[string]interface{}{
																				"description": `selectPolicy is used to specify which policy should be used.
If not set, the default value Max is used.`,
																				"type": "string",
																			},
																			"stabilizationWindowSeconds": map[string]interface{}{
																				"description": `stabilizationWindowSeconds is the number of seconds for which past recommendations should be
considered while scaling up or scaling down.
StabilizationWindowSeconds must be greater than or equal to zero and less than or equal to 3600 (one hour).
If not set, use the default values:
- For scale up: 0 (i.e. no stabilization is done).
- For scale down: 300 (i.e. the stabilization window is 300 seconds long).`,
																				"format":  "int32",
																				"maximum": 3600,
																				"minimum": 0,
																				"type":    "integer",
																			},
																			"tolerance": map[string]interface{}{
																				"anyOf": []interface{}{
																					map[string]interface{}{
																						"type": "integer",
																					},
																					map[string]interface{}{
																						"type": "string",
																					},
																				},
																				"description": `tolerance is the tolerance on the ratio between the current and desired
metric value under which no updates are made to the desired number of
replicas (e.g. 0.01 for 1%). Must be greater than or equal to zero. If not
set, the default cluster-wide tolerance is applied (by default 10%).

For example, if autoscaling is configured with a memory consumption target of 100Mi,
and scale-down and scale-up tolerances of 5% and 1% respectively, scaling will be
triggered when the actual consumption falls below 95Mi or exceeds 101Mi.

This is an beta field and requires the HPAConfigurableTolerance feature
gate to be enabled.`,
																				"pattern":                    `^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$`,
																				"x-kubernetes-int-or-string": true,
																			},
																		},
																		"type": "object",
																	},
																	"scaleUp": map[string]interface{}{
																		"description": `scaleUp is scaling policy for scaling Up.
If not set, the default value is the higher of:
  * increase no more than 4 pods per 60 seconds
  * double the number of pods per 60 seconds
No stabilization is used.`,
																		"properties": map[string]interface{}{
																			"policies": map[string]interface{}{
																				"description": `policies is a list of potential scaling polices which can be used during scaling.
If not set, use the default values:
- For scale up: allow doubling the number of pods, or an absolute change of 4 pods in a 15s window.
- For scale down: allow all pods to be removed in a 15s window.`,
																				"items": map[string]interface{}{
																					"description": "HPAScalingPolicy is a single policy which must hold true for a specified past interval.",
																					"properties": map[string]interface{}{
																						"periodSeconds": map[string]interface{}{
																							"description": `periodSeconds specifies the window of time for which the policy should hold true.
PeriodSeconds must be greater than zero and less than or equal to 1800 (30 min).`,
																							"format": "int32",
																							"type":   "integer",
																						},
																						"type": map[string]interface{}{
																							"description": "type is used to specify the scaling policy.",
																							"type":        "string",
																						},
																						"value": map[string]interface{}{
																							"description": `value contains the amount of change which is permitted by the policy.
It must be greater than zero`,
																							"format": "int32",
																							"type":   "integer",
																						},
																					},
																					"required": []interface{}{
																						"periodSeconds",
																						"type",
																						"value",
																					},
																					"type": "object",
																				},
																				"type":                   "array",
																				"x-kubernetes-list-type": "atomic",
																			},
																			"selectPolicy": map[string]interface{}{
																				"description": `selectPolicy is used to specify which policy should be used.
If not set, the default value Max is used.`,
																				"type": "string",
																			},
																			"stabilizationWindowSeconds": map[string]interface{}{
																				"description": `stabilizationWindowSeconds is the number of seconds for which past recommendations should be
considered while scaling up or scaling down.
StabilizationWindowSeconds must be greater than or equal to zero and less than or equal to 3600 (one hour).
If not set, use the default values:
- For scale up: 0 (i.e. no stabilization is done).
- For scale down: 300 (i.e. the stabilization window is 300 seconds long).`,
																				"format":  "int32",
																				"maximum": 3600,
																				"minimum": 0,
																				"type":    "integer",
																			},
																			"tolerance": map[string]interface{}{
																				"anyOf": []interface{}{
																					map[string]interface{}{
																						"type": "integer",
																					},
																					map[string]interface{}{
																						"type": "string",
																					},
																				},
																				"description": `tolerance is the tolerance on the ratio between the current and desired
metric value under which no updates are made to the desired number of
replicas (e.g. 0.01 for 1%). Must be greater than or equal to zero. If not
set, the default cluster-wide tolerance is applied (by default 10%).

For example, if autoscaling is configured with a memory consumption target of 100Mi,
and scale-down and scale-up tolerances of 5% and 1% respectively, scaling will be
triggered when the actual consumption falls below 95Mi or exceeds 101Mi.

This is an beta field and requires the HPAConfigurableTolerance feature
gate to be enabled.`,
																				"pattern":                    `^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$`,
																				"x-kubernetes-int-or-string": true,
																			},
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
															"name": map[string]interface{}{
																"type": "string",
															},
														},
														"type": "object",
													},
													"restoreToOriginalReplicaCount": map[string]interface{}{
														"type": "boolean",
													},
													"scalingModifiers": map[string]interface{}{
														"description": "ScalingModifiers describes advanced scaling logic options like formula",
														"properties": map[string]interface{}{
															"activationTarget": map[string]interface{}{
																"type": "string",
															},
															"formula": map[string]interface{}{
																"type": "string",
															},
															"metricType": map[string]interface{}{
																"description": `MetricTargetType specifies the type of metric being targeted, and should be either
"Value", "AverageValue", or "Utilization"`,
																"enum": []interface{}{
																	"AverageValue",
																	"Value",
																},
																"type": "string",
															},
															"target": map[string]interface{}{
																"type": "string",
															},
														},
														"type": "object",
													},
												},
												"type": "object",
											},
											"cooldownPeriod": map[string]interface{}{
												"format":  "int32",
												"minimum": 0,
												"type":    "integer",
											},
											"fallback": map[string]interface{}{
												"description": "Fallback is the spec for fallback options",
												"properties": map[string]interface{}{
													"behavior": map[string]interface{}{
														"default": "static",
														"enum": []interface{}{
															"static",
															"currentReplicas",
															"currentReplicasIfHigher",
															"currentReplicasIfLower",
															"scalingModifiers",
														},
														"type": "string",
													},
													"failureThreshold": map[string]interface{}{
														"format":  "int32",
														"minimum": 0,
														"type":    "integer",
													},
													"replicas": map[string]interface{}{
														"format":  "int32",
														"minimum": 0,
														"type":    "integer",
													},
												},
												"required": []interface{}{
													"failureThreshold",
													"replicas",
												},
												"type": "object",
											},
											"idleReplicaCount": map[string]interface{}{
												"format":  "int32",
												"minimum": 0,
												"type":    "integer",
											},
											"initialCooldownPeriod": map[string]interface{}{
												"format":  "int32",
												"minimum": 0,
												"type":    "integer",
											},
											"maxReplicaCount": map[string]interface{}{
												"format":  "int32",
												"minimum": 1,
												"type":    "integer",
											},
											"minReplicaCount": map[string]interface{}{
												"format":  "int32",
												"minimum": 0,
												"type":    "integer",
											},
											"pollingInterval": map[string]interface{}{
												"format":  "int32",
												"minimum": 1,
												"type":    "integer",
											},
											"scaleTargetRef": map[string]interface{}{
												"description": "ScaleTarget holds the reference to the scale target Object",
												"properties": map[string]interface{}{
													"apiVersion": map[string]interface{}{
														"type": "string",
													},
													"envSourceContainerName": map[string]interface{}{
														"type": "string",
													},
													"kind": map[string]interface{}{
														"type": "string",
													},
													"name": map[string]interface{}{
														"minLength": 1,
														"type":      "string",
													},
												},
												"required": []interface{}{
													"name",
												},
												"type": "object",
											},
											"triggers": map[string]interface{}{
												"items": map[string]interface{}{
													"description": "ScaleTriggers reference the scaler that will be used",
													"properties": map[string]interface{}{
														"authenticationRef": map[string]interface{}{
															"description": `AuthenticationRef points to the TriggerAuthentication or ClusterTriggerAuthentication object that
is used to authenticate the scaler with the environment`,
															"properties": map[string]interface{}{
																"kind": map[string]interface{}{
																	"description": "Kind of the resource being referred to. Defaults to TriggerAuthentication.",
																	"enum": []interface{}{
																		"TriggerAuthentication",
																		"ClusterTriggerAuthentication",
																	},
																	"type": "string",
																},
																"name": map[string]interface{}{
																	"type": "string",
																},
															},
															"required": []interface{}{
																"name",
															},
															"type": "object",
														},
														"metadata": map[string]interface{}{
															"additionalProperties": map[string]interface{}{
																"type": "string",
															},
															"type": "object",
														},
														"metricType": map[string]interface{}{
															"description": `MetricTargetType specifies the type of metric being targeted, and should be either
"Value", "AverageValue", or "Utilization"`,
															"type": "string",
														},
														"name": map[string]interface{}{
															"type": "string",
														},
														"type": map[string]interface{}{
															"minLength": 1,
															"type":      "string",
														},
														"useCachedMetrics": map[string]interface{}{
															"type": "boolean",
														},
													},
													"required": []interface{}{
														"metadata",
														"type",
													},
													"type": "object",
												},
												"minItems": 1,
												"type":     "array",
											},
										},
										"required": []interface{}{
											"scaleTargetRef",
											"triggers",
										},
										"type": "object",
										"x-kubernetes-validations": []interface{}{
											map[string]interface{}{
												"message": "minReplicaCount must be less than or equal to maxReplicaCount",
												"rule":    "!has(self.minReplicaCount) || self.minReplicaCount <= (has(self.maxReplicaCount) ? self.maxReplicaCount : 100)",
											},
										},
									},
									"status": map[string]interface{}{
										"description": "ScaledObjectStatus is the status for a ScaledObject resource",
										"properties": map[string]interface{}{
											"authenticationsTypes": map[string]interface{}{
												"type": "string",
											},
											"compositeScalerName": map[string]interface{}{
												"type": "string",
											},
											"conditions": map[string]interface{}{
												"description": "Conditions an array representation to store multiple Conditions",
												"items": map[string]interface{}{
													"description": "Condition to store the condition state",
													"properties": map[string]interface{}{
														"message": map[string]interface{}{
															"description": "A human readable message indicating details about the transition.",
															"type":        "string",
														},
														"reason": map[string]interface{}{
															"description": "The reason for the condition's last transition.",
															"type":        "string",
														},
														"status": map[string]interface{}{
															"description": "Status of the condition, one of True, False, Unknown.",
															"type":        "string",
														},
														"type": map[string]interface{}{
															"description": "Type of condition",
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
											"externalMetricNames": map[string]interface{}{
												"items": map[string]interface{}{
													"type": "string",
												},
												"type": "array",
											},
											"health": map[string]interface{}{
												"additionalProperties": map[string]interface{}{
													"description": "HealthStatus is the status for a ScaledObject's health",
													"properties": map[string]interface{}{
														"numberOfFailures": map[string]interface{}{
															"format": "int32",
															"type":   "integer",
														},
														"status": map[string]interface{}{
															"description": "HealthStatusType is an indication of whether the health status is happy or failing",
															"type":        "string",
														},
													},
													"type": "object",
												},
												"type": "object",
											},
											"hpaName": map[string]interface{}{
												"type": "string",
											},
											"lastActiveTime": map[string]interface{}{
												"format": "date-time",
												"type":   "string",
											},
											"originalReplicaCount": map[string]interface{}{
												"format": "int32",
												"type":   "integer",
											},
											"pausedReplicaCount": map[string]interface{}{
												"format": "int32",
												"type":   "integer",
											},
											"resourceMetricNames": map[string]interface{}{
												"items": map[string]interface{}{
													"type": "string",
												},
												"type": "array",
											},
											"scaleTargetGVKR": map[string]interface{}{
												"description": "GroupVersionKindResource provides unified structure for schema.GroupVersionKind and Resource",
												"properties": map[string]interface{}{
													"group": map[string]interface{}{
														"type": "string",
													},
													"kind": map[string]interface{}{
														"type": "string",
													},
													"resource": map[string]interface{}{
														"type": "string",
													},
													"version": map[string]interface{}{
														"type": "string",
													},
												},
												"required": []interface{}{
													"group",
													"kind",
													"resource",
													"version",
												},
												"type": "object",
											},
											"scaleTargetKind": map[string]interface{}{
												"type": "string",
											},
											"triggersActivity": map[string]interface{}{
												"additionalProperties": map[string]interface{}{
													"description": "TriggerActivityStatus represents the activity status of an external trigger",
													"properties": map[string]interface{}{
														"isActive": map[string]interface{}{
															"type": "boolean",
														},
													},
													"type": "object",
												},
												"type": "object",
											},
											"triggersTypes": map[string]interface{}{
												"type": "string",
											},
										},
										"type": "object",
									},
								},
								"required": []interface{}{
									"spec",
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
			},
		},
	}

	return mutate.MutateCRDScaledobjectsKedaSh(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;create;update;patch;delete

// CreateCRDTriggerauthenticationsKedaSh creates the CustomResourceDefinition resource with name triggerauthentications.keda.sh.
func CreateCRDTriggerauthenticationsKedaSh(
	parent *platformv1alpha1.AutoscalingComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					"controller-gen.kubebuilder.io/version": "v0.21.0",
				},
				"labels": map[string]interface{}{
					"app.kubernetes.io/name":        "keda-operator",
					"app.kubernetes.io/component":   "operator",
					"app.kubernetes.io/part-of":     "keda-operator",
					"app.kubernetes.io/version":     parent.Spec.Keda.Version, //  controlled by field: keda.version
					"platform.nukleros.io/category": "autoscaling",
					"platform.nukleros.io/project":  "keda",
				},
				"name": "triggerauthentications.keda.sh",
			},
			"spec": map[string]interface{}{
				"group": "keda.sh",
				"names": map[string]interface{}{
					"kind":     "TriggerAuthentication",
					"listKind": "TriggerAuthenticationList",
					"plural":   "triggerauthentications",
					"shortNames": []interface{}{
						"ta",
						"triggerauth",
					},
					"singular": "triggerauthentication",
				},
				"scope": "Namespaced",
				"versions": []interface{}{
					map[string]interface{}{
						"additionalPrinterColumns": []interface{}{
							map[string]interface{}{
								"jsonPath": ".spec.podIdentity.provider",
								"name":     "PodIdentity",
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".spec.secretTargetRef[*].name",
								"name":     "Secret",
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".spec.env[*].name",
								"name":     "Env",
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".spec.hashiCorpVault.address",
								"name":     "VaultAddress",
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".status.scaledobjects",
								"name":     "ScaledObjects",
								"priority": 1,
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".status.scaledjobs",
								"name":     "ScaledJobs",
								"priority": 1,
								"type":     "string",
							},
						},
						"name": "v1alpha1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"description": "TriggerAuthentication defines how a trigger can authenticate",
								"properties": map[string]interface{}{
									"apiVersion": map[string]interface{}{
										"description": `APIVersion defines the versioned schema of this representation of an object.
Servers should convert recognized schemas to the latest internal value, and
may reject unrecognized values.
More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources`,
										"type": "string",
									},
									"kind": map[string]interface{}{
										"description": `Kind is a string value representing the REST resource this object represents.
Servers may infer this from the endpoint the client submits requests to.
Cannot be updated.
In CamelCase.
More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds`,
										"type": "string",
									},
									"metadata": map[string]interface{}{
										"type": "object",
									},
									"spec": map[string]interface{}{
										"description": "TriggerAuthenticationSpec defines the various ways to authenticate",
										"properties": map[string]interface{}{
											"awsSecretManager": map[string]interface{}{
												"description": "AwsSecretManager is used to authenticate using AwsSecretManager",
												"properties": map[string]interface{}{
													"credentials": map[string]interface{}{
														"properties": map[string]interface{}{
															"accessKey": map[string]interface{}{
																"properties": map[string]interface{}{
																	"valueFrom": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"secretKeyRef": map[string]interface{}{
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"type": "string",
																					},
																					"name": map[string]interface{}{
																						"type": "string",
																					},
																				},
																				"required": []interface{}{
																					"key",
																					"name",
																				},
																				"type": "object",
																			},
																		},
																		"required": []interface{}{
																			"secretKeyRef",
																		},
																		"type": "object",
																	},
																},
																"required": []interface{}{
																	"valueFrom",
																},
																"type": "object",
															},
															"accessSecretKey": map[string]interface{}{
																"properties": map[string]interface{}{
																	"valueFrom": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"secretKeyRef": map[string]interface{}{
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"type": "string",
																					},
																					"name": map[string]interface{}{
																						"type": "string",
																					},
																				},
																				"required": []interface{}{
																					"key",
																					"name",
																				},
																				"type": "object",
																			},
																		},
																		"required": []interface{}{
																			"secretKeyRef",
																		},
																		"type": "object",
																	},
																},
																"required": []interface{}{
																	"valueFrom",
																},
																"type": "object",
															},
															"accessToken": map[string]interface{}{
																"properties": map[string]interface{}{
																	"valueFrom": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"secretKeyRef": map[string]interface{}{
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"type": "string",
																					},
																					"name": map[string]interface{}{
																						"type": "string",
																					},
																				},
																				"required": []interface{}{
																					"key",
																					"name",
																				},
																				"type": "object",
																			},
																		},
																		"required": []interface{}{
																			"secretKeyRef",
																		},
																		"type": "object",
																	},
																},
																"required": []interface{}{
																	"valueFrom",
																},
																"type": "object",
															},
														},
														"required": []interface{}{
															"accessKey",
															"accessSecretKey",
														},
														"type": "object",
													},
													"podIdentity": map[string]interface{}{
														"description": `AuthPodIdentity allows users to select the platform native identity
mechanism`,
														"properties": map[string]interface{}{
															"externalID": map[string]interface{}{
																"description": "ExternalID sets the External ID to be used when assuming an identity. This is only applicable when using AWS pod identity with a RoleArn.",
																"type":        "string",
															},
															"identityAuthorityHost": map[string]interface{}{
																"description": "Set identityAuthorityHost to override the default Azure authority host. If this is set, then the IdentityTenantID must also be set",
																"type":        "string",
															},
															"identityId": map[string]interface{}{
																"type": "string",
															},
															"identityOwner": map[string]interface{}{
																"description": "IdentityOwner configures which identity has to be used during auto discovery, keda or the scaled workload. Mutually exclusive with roleArn",
																"enum": []interface{}{
																	"keda",
																	"workload",
																},
																"type": "string",
															},
															"identityTenantId": map[string]interface{}{
																"description": "Set identityTenantId to override the default Azure tenant id. If this is set, then the IdentityID must also be set",
																"type":        "string",
															},
															"provider": map[string]interface{}{
																"description": "PodIdentityProvider contains the list of providers",
																"enum": []interface{}{
																	"azure-workload",
																	"gcp",
																	"aws",
																	"aws-eks",
																	"none",
																},
																"type": "string",
															},
															"roleArn": map[string]interface{}{
																"description": "RoleArn sets the AWS RoleArn to be used. Mutually exclusive with IdentityOwner",
																"type":        "string",
															},
														},
														"required": []interface{}{
															"provider",
														},
														"type": "object",
													},
													"region": map[string]interface{}{
														"type": "string",
													},
													"secrets": map[string]interface{}{
														"items": map[string]interface{}{
															"properties": map[string]interface{}{
																"name": map[string]interface{}{
																	"type": "string",
																},
																"parameter": map[string]interface{}{
																	"type": "string",
																},
																"secretKey": map[string]interface{}{
																	"type": "string",
																},
																"versionId": map[string]interface{}{
																	"type": "string",
																},
																"versionStage": map[string]interface{}{
																	"type": "string",
																},
															},
															"required": []interface{}{
																"name",
																"parameter",
															},
															"type": "object",
														},
														"minItems": 1,
														"type":     "array",
													},
												},
												"required": []interface{}{
													"secrets",
												},
												"type": "object",
											},
											"azureKeyVault": map[string]interface{}{
												"description": "AzureKeyVault is used to authenticate using Azure Key Vault",
												"properties": map[string]interface{}{
													"cloud": map[string]interface{}{
														"properties": map[string]interface{}{
															"activeDirectoryEndpoint": map[string]interface{}{
																"type": "string",
															},
															"keyVaultResourceURL": map[string]interface{}{
																"type": "string",
															},
															"type": map[string]interface{}{
																"type": "string",
															},
														},
														"required": []interface{}{
															"type",
														},
														"type": "object",
													},
													"credentials": map[string]interface{}{
														"properties": map[string]interface{}{
															"clientId": map[string]interface{}{
																"type": "string",
															},
															"clientSecret": map[string]interface{}{
																"properties": map[string]interface{}{
																	"valueFrom": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"secretKeyRef": map[string]interface{}{
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"type": "string",
																					},
																					"name": map[string]interface{}{
																						"type": "string",
																					},
																				},
																				"required": []interface{}{
																					"key",
																					"name",
																				},
																				"type": "object",
																			},
																		},
																		"required": []interface{}{
																			"secretKeyRef",
																		},
																		"type": "object",
																	},
																},
																"required": []interface{}{
																	"valueFrom",
																},
																"type": "object",
															},
															"tenantId": map[string]interface{}{
																"type": "string",
															},
														},
														"required": []interface{}{
															"clientId",
															"clientSecret",
															"tenantId",
														},
														"type": "object",
													},
													"podIdentity": map[string]interface{}{
														"description": `AuthPodIdentity allows users to select the platform native identity
mechanism`,
														"properties": map[string]interface{}{
															"externalID": map[string]interface{}{
																"description": "ExternalID sets the External ID to be used when assuming an identity. This is only applicable when using AWS pod identity with a RoleArn.",
																"type":        "string",
															},
															"identityAuthorityHost": map[string]interface{}{
																"description": "Set identityAuthorityHost to override the default Azure authority host. If this is set, then the IdentityTenantID must also be set",
																"type":        "string",
															},
															"identityId": map[string]interface{}{
																"type": "string",
															},
															"identityOwner": map[string]interface{}{
																"description": "IdentityOwner configures which identity has to be used during auto discovery, keda or the scaled workload. Mutually exclusive with roleArn",
																"enum": []interface{}{
																	"keda",
																	"workload",
																},
																"type": "string",
															},
															"identityTenantId": map[string]interface{}{
																"description": "Set identityTenantId to override the default Azure tenant id. If this is set, then the IdentityID must also be set",
																"type":        "string",
															},
															"provider": map[string]interface{}{
																"description": "PodIdentityProvider contains the list of providers",
																"enum": []interface{}{
																	"azure-workload",
																	"gcp",
																	"aws",
																	"aws-eks",
																	"none",
																},
																"type": "string",
															},
															"roleArn": map[string]interface{}{
																"description": "RoleArn sets the AWS RoleArn to be used. Mutually exclusive with IdentityOwner",
																"type":        "string",
															},
														},
														"required": []interface{}{
															"provider",
														},
														"type": "object",
													},
													"secrets": map[string]interface{}{
														"items": map[string]interface{}{
															"properties": map[string]interface{}{
																"name": map[string]interface{}{
																	"type": "string",
																},
																"parameter": map[string]interface{}{
																	"type": "string",
																},
																"version": map[string]interface{}{
																	"type": "string",
																},
															},
															"required": []interface{}{
																"name",
																"parameter",
															},
															"type": "object",
														},
														"minItems": 1,
														"type":     "array",
													},
													"vaultUri": map[string]interface{}{
														"type": "string",
													},
												},
												"required": []interface{}{
													"secrets",
													"vaultUri",
												},
												"type": "object",
											},
											"azureServicePrincipal": map[string]interface{}{
												"description": `AzureServicePrincipal defines Microsoft Entra service principal credentials
that can be shared by Azure scalers.`,
												"properties": map[string]interface{}{
													"activeDirectoryEndpoint": map[string]interface{}{
														"description": "ActiveDirectoryEndpoint is required when cloud is set to Private.",
														"type":        "string",
													},
													"clientCertificate": map[string]interface{}{
														"description": `AzureServicePrincipalCredential references credential data stored in a
Kubernetes Secret.`,
														"properties": map[string]interface{}{
															"valueFrom": map[string]interface{}{
																"properties": map[string]interface{}{
																	"secretKeyRef": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"type": "string",
																			},
																			"name": map[string]interface{}{
																				"type": "string",
																			},
																		},
																		"required": []interface{}{
																			"key",
																			"name",
																		},
																		"type": "object",
																	},
																},
																"required": []interface{}{
																	"secretKeyRef",
																},
																"type": "object",
															},
														},
														"required": []interface{}{
															"valueFrom",
														},
														"type": "object",
													},
													"clientCertificatePassword": map[string]interface{}{
														"description": `AzureServicePrincipalCredential references credential data stored in a
Kubernetes Secret.`,
														"properties": map[string]interface{}{
															"valueFrom": map[string]interface{}{
																"properties": map[string]interface{}{
																	"secretKeyRef": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"type": "string",
																			},
																			"name": map[string]interface{}{
																				"type": "string",
																			},
																		},
																		"required": []interface{}{
																			"key",
																			"name",
																		},
																		"type": "object",
																	},
																},
																"required": []interface{}{
																	"secretKeyRef",
																},
																"type": "object",
															},
														},
														"required": []interface{}{
															"valueFrom",
														},
														"type": "object",
													},
													"clientId": map[string]interface{}{
														"type": "string",
													},
													"clientSecret": map[string]interface{}{
														"description": `AzureServicePrincipalCredential references credential data stored in a
Kubernetes Secret.`,
														"properties": map[string]interface{}{
															"valueFrom": map[string]interface{}{
																"properties": map[string]interface{}{
																	"secretKeyRef": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"type": "string",
																			},
																			"name": map[string]interface{}{
																				"type": "string",
																			},
																		},
																		"required": []interface{}{
																			"key",
																			"name",
																		},
																		"type": "object",
																	},
																},
																"required": []interface{}{
																	"secretKeyRef",
																},
																"type": "object",
															},
														},
														"required": []interface{}{
															"valueFrom",
														},
														"type": "object",
													},
													"cloud": map[string]interface{}{
														"description": "Cloud specifies the Azure cloud environment. Defaults to AzurePublicCloud.",
														"type":        "string",
													},
													"tenantId": map[string]interface{}{
														"type": "string",
													},
												},
												"required": []interface{}{
													"clientId",
													"tenantId",
												},
												"type": "object",
											},
											"boundServiceAccountToken": map[string]interface{}{
												"items": map[string]interface{}{
													"properties": map[string]interface{}{
														"parameter": map[string]interface{}{
															"type": "string",
														},
														"serviceAccountName": map[string]interface{}{
															"type": "string",
														},
													},
													"required": []interface{}{
														"parameter",
														"serviceAccountName",
													},
													"type": "object",
												},
												"type": "array",
											},
											"configMapTargetRef": map[string]interface{}{
												"items": map[string]interface{}{
													"description": "AuthConfigMapTargetRef is used to authenticate using a reference to a config map",
													"properties": map[string]interface{}{
														"key": map[string]interface{}{
															"type": "string",
														},
														"name": map[string]interface{}{
															"type": "string",
														},
														"parameter": map[string]interface{}{
															"type": "string",
														},
													},
													"required": []interface{}{
														"key",
														"name",
														"parameter",
													},
													"type": "object",
												},
												"type": "array",
											},
											"env": map[string]interface{}{
												"items": map[string]interface{}{
													"description": `AuthEnvironment is used to authenticate using environment variables
in the destination ScaleTarget spec`,
													"properties": map[string]interface{}{
														"containerName": map[string]interface{}{
															"type": "string",
														},
														"name": map[string]interface{}{
															"type": "string",
														},
														"parameter": map[string]interface{}{
															"type": "string",
														},
													},
													"required": []interface{}{
														"name",
														"parameter",
													},
													"type": "object",
												},
												"type": "array",
											},
											"filePath": map[string]interface{}{
												"description": `FilePath specifies a file containing auth parameters as JSON map[string]string.
When set, auth params are read directly from this file instead.`,
												"type": "string",
											},
											"gcpSecretManager": map[string]interface{}{
												"properties": map[string]interface{}{
													"credentials": map[string]interface{}{
														"properties": map[string]interface{}{
															"clientSecret": map[string]interface{}{
																"properties": map[string]interface{}{
																	"valueFrom": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"secretKeyRef": map[string]interface{}{
																				"properties": map[string]interface{}{
																					"key": map[string]interface{}{
																						"type": "string",
																					},
																					"name": map[string]interface{}{
																						"type": "string",
																					},
																				},
																				"required": []interface{}{
																					"key",
																					"name",
																				},
																				"type": "object",
																			},
																		},
																		"required": []interface{}{
																			"secretKeyRef",
																		},
																		"type": "object",
																	},
																},
																"required": []interface{}{
																	"valueFrom",
																},
																"type": "object",
															},
														},
														"required": []interface{}{
															"clientSecret",
														},
														"type": "object",
													},
													"podIdentity": map[string]interface{}{
														"description": `AuthPodIdentity allows users to select the platform native identity
mechanism`,
														"properties": map[string]interface{}{
															"externalID": map[string]interface{}{
																"description": "ExternalID sets the External ID to be used when assuming an identity. This is only applicable when using AWS pod identity with a RoleArn.",
																"type":        "string",
															},
															"identityAuthorityHost": map[string]interface{}{
																"description": "Set identityAuthorityHost to override the default Azure authority host. If this is set, then the IdentityTenantID must also be set",
																"type":        "string",
															},
															"identityId": map[string]interface{}{
																"type": "string",
															},
															"identityOwner": map[string]interface{}{
																"description": "IdentityOwner configures which identity has to be used during auto discovery, keda or the scaled workload. Mutually exclusive with roleArn",
																"enum": []interface{}{
																	"keda",
																	"workload",
																},
																"type": "string",
															},
															"identityTenantId": map[string]interface{}{
																"description": "Set identityTenantId to override the default Azure tenant id. If this is set, then the IdentityID must also be set",
																"type":        "string",
															},
															"provider": map[string]interface{}{
																"description": "PodIdentityProvider contains the list of providers",
																"enum": []interface{}{
																	"azure-workload",
																	"gcp",
																	"aws",
																	"aws-eks",
																	"none",
																},
																"type": "string",
															},
															"roleArn": map[string]interface{}{
																"description": "RoleArn sets the AWS RoleArn to be used. Mutually exclusive with IdentityOwner",
																"type":        "string",
															},
														},
														"required": []interface{}{
															"provider",
														},
														"type": "object",
													},
													"secrets": map[string]interface{}{
														"items": map[string]interface{}{
															"properties": map[string]interface{}{
																"id": map[string]interface{}{
																	"type": "string",
																},
																"parameter": map[string]interface{}{
																	"type": "string",
																},
																"version": map[string]interface{}{
																	"type": "string",
																},
															},
															"required": []interface{}{
																"id",
																"parameter",
															},
															"type": "object",
														},
														"minItems": 1,
														"type":     "array",
													},
												},
												"required": []interface{}{
													"secrets",
												},
												"type": "object",
											},
											"hashiCorpVault": map[string]interface{}{
												"description": "HashiCorpVault is used to authenticate using Hashicorp Vault",
												"properties": map[string]interface{}{
													"address": map[string]interface{}{
														"type": "string",
													},
													"authentication": map[string]interface{}{
														"description": "VaultAuthentication contains the list of Hashicorp Vault authentication methods",
														"enum": []interface{}{
															"token",
															"kubernetes",
														},
														"type": "string",
													},
													"credential": map[string]interface{}{
														"description": "Credential defines the Hashicorp Vault credentials depending on the authentication method",
														"properties": map[string]interface{}{
															"serviceAccount": map[string]interface{}{
																"type": "string",
															},
															"serviceAccountName": map[string]interface{}{
																"type": "string",
															},
															"token": map[string]interface{}{
																"type": "string",
															},
														},
														"type": "object",
													},
													"mount": map[string]interface{}{
														"type": "string",
													},
													"namespace": map[string]interface{}{
														"type": "string",
													},
													"role": map[string]interface{}{
														"type": "string",
													},
													"secrets": map[string]interface{}{
														"items": map[string]interface{}{
															"description": "VaultSecret defines the mapping between the path of the secret in Vault to the parameter",
															"properties": map[string]interface{}{
																"key": map[string]interface{}{
																	"type": "string",
																},
																"parameter": map[string]interface{}{
																	"type": "string",
																},
																"path": map[string]interface{}{
																	"type": "string",
																},
																"pkiData": map[string]interface{}{
																	"properties": map[string]interface{}{
																		"altNames": map[string]interface{}{
																			"type": "string",
																		},
																		"commonName": map[string]interface{}{
																			"type": "string",
																		},
																		"format": map[string]interface{}{
																			"type": "string",
																		},
																		"ipSans": map[string]interface{}{
																			"type": "string",
																		},
																		"otherSans": map[string]interface{}{
																			"type": "string",
																		},
																		"ttl": map[string]interface{}{
																			"type": "string",
																		},
																		"uriSans": map[string]interface{}{
																			"type": "string",
																		},
																	},
																	"type": "object",
																},
																"type": map[string]interface{}{
																	"description": "VaultSecretType defines the type of vault secret",
																	"enum": []interface{}{
																		"",
																		"secretV2",
																		"secret",
																		"pki",
																	},
																	"type": "string",
																},
															},
															"required": []interface{}{
																"key",
																"parameter",
																"path",
															},
															"type": "object",
														},
														"minItems": 1,
														"type":     "array",
													},
												},
												"required": []interface{}{
													"address",
													"authentication",
													"secrets",
												},
												"type": "object",
											},
											"oauth2": map[string]interface{}{
												"properties": map[string]interface{}{
													"clientId": map[string]interface{}{
														"type": "string",
													},
													"clientSecret": map[string]interface{}{
														"properties": map[string]interface{}{
															"valueFrom": map[string]interface{}{
																"properties": map[string]interface{}{
																	"secretKeyRef": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"key": map[string]interface{}{
																				"type": "string",
																			},
																			"name": map[string]interface{}{
																				"type": "string",
																			},
																		},
																		"required": []interface{}{
																			"key",
																			"name",
																		},
																		"type": "object",
																	},
																},
																"required": []interface{}{
																	"secretKeyRef",
																},
																"type": "object",
															},
														},
														"required": []interface{}{
															"valueFrom",
														},
														"type": "object",
													},
													"scopes": map[string]interface{}{
														"items": map[string]interface{}{
															"type": "string",
														},
														"type": "array",
													},
													"tokenUrl": map[string]interface{}{
														"type": "string",
													},
													"tokenUrlParams": map[string]interface{}{
														"additionalProperties": map[string]interface{}{
															"type": "string",
														},
														"type": "object",
													},
													"type": map[string]interface{}{
														"default": "clientCredentials",
														"enum": []interface{}{
															"clientCredentials",
														},
														"type": "string",
													},
												},
												"required": []interface{}{
													"clientId",
													"tokenUrl",
													"type",
												},
												"type": "object",
											},
											"podIdentity": map[string]interface{}{
												"description": `AuthPodIdentity allows users to select the platform native identity
mechanism`,
												"properties": map[string]interface{}{
													"externalID": map[string]interface{}{
														"description": "ExternalID sets the External ID to be used when assuming an identity. This is only applicable when using AWS pod identity with a RoleArn.",
														"type":        "string",
													},
													"identityAuthorityHost": map[string]interface{}{
														"description": "Set identityAuthorityHost to override the default Azure authority host. If this is set, then the IdentityTenantID must also be set",
														"type":        "string",
													},
													"identityId": map[string]interface{}{
														"type": "string",
													},
													"identityOwner": map[string]interface{}{
														"description": "IdentityOwner configures which identity has to be used during auto discovery, keda or the scaled workload. Mutually exclusive with roleArn",
														"enum": []interface{}{
															"keda",
															"workload",
														},
														"type": "string",
													},
													"identityTenantId": map[string]interface{}{
														"description": "Set identityTenantId to override the default Azure tenant id. If this is set, then the IdentityID must also be set",
														"type":        "string",
													},
													"provider": map[string]interface{}{
														"description": "PodIdentityProvider contains the list of providers",
														"enum": []interface{}{
															"azure-workload",
															"gcp",
															"aws",
															"aws-eks",
															"none",
														},
														"type": "string",
													},
													"roleArn": map[string]interface{}{
														"description": "RoleArn sets the AWS RoleArn to be used. Mutually exclusive with IdentityOwner",
														"type":        "string",
													},
												},
												"required": []interface{}{
													"provider",
												},
												"type": "object",
											},
											"secretTargetRef": map[string]interface{}{
												"items": map[string]interface{}{
													"description": "AuthSecretTargetRef is used to authenticate using a reference to a secret",
													"properties": map[string]interface{}{
														"key": map[string]interface{}{
															"type": "string",
														},
														"name": map[string]interface{}{
															"type": "string",
														},
														"parameter": map[string]interface{}{
															"type": "string",
														},
													},
													"required": []interface{}{
														"key",
														"name",
														"parameter",
													},
													"type": "object",
												},
												"type": "array",
											},
										},
										"type": "object",
									},
									"status": map[string]interface{}{
										"description": "TriggerAuthenticationStatus defines the observed state of TriggerAuthentication",
										"properties": map[string]interface{}{
											"scaledjobs": map[string]interface{}{
												"type": "string",
											},
											"scaledobjects": map[string]interface{}{
												"type": "string",
											},
										},
										"type": "object",
									},
								},
								"required": []interface{}{
									"spec",
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
			},
		},
	}

	return mutate.MutateCRDTriggerauthenticationsKedaSh(resourceObj, parent, collection, reconciler, req)
}
