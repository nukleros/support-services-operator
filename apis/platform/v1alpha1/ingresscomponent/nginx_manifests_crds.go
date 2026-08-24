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

package ingresscomponent

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	platformv1alpha1 "github.com/nukleros/support-services-operator/apis/platform/v1alpha1"
	"github.com/nukleros/support-services-operator/apis/platform/v1alpha1/ingresscomponent/mutate"
	setupv1alpha1 "github.com/nukleros/support-services-operator/apis/setup/v1alpha1"
)

// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;create;update;patch;delete

// CreateCRDDnsendpointsExternaldnsNginxOrg creates the CustomResourceDefinition resource with name dnsendpoints.externaldns.nginx.org.
func CreateCRDDnsendpointsExternaldnsNginxOrg(
	parent *platformv1alpha1.IngressComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if parent.Spec.Nginx.Include != true {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:field=nginx.include,value=true,include
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					"controller-gen.kubebuilder.io/version": "v0.21.0",
				},
				"name": "dnsendpoints.externaldns.nginx.org",
				"labels": map[string]interface{}{
					"platform.nukleros.io/category": "ingress",
					"platform.nukleros.io/project":  "nginx-ingress-controller",
				},
			},
			"spec": map[string]interface{}{
				"group": "externaldns.nginx.org",
				"names": map[string]interface{}{
					"kind":     "DNSEndpoint",
					"listKind": "DNSEndpointList",
					"plural":   "dnsendpoints",
					"singular": "dnsendpoint",
				},
				"scope": "Namespaced",
				"versions": []interface{}{
					map[string]interface{}{
						"name": "v1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"description": "DNSEndpoint is the CRD wrapper for Endpoint",
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
										"description": "DNSEndpointSpec holds information about endpoints.",
										"properties": map[string]interface{}{
											"endpoints": map[string]interface{}{
												"items": map[string]interface{}{
													"description": "Endpoint describes DNS Endpoint.",
													"properties": map[string]interface{}{
														"dnsName": map[string]interface{}{
															"description": "The hostname for the DNS record",
															"type":        "string",
														},
														"labels": map[string]interface{}{
															"additionalProperties": map[string]interface{}{
																"type": "string",
															},
															"description": "Labels stores labels defined for the Endpoint",
															"type":        "object",
														},
														"providerSpecific": map[string]interface{}{
															"description": "ProviderSpecific stores provider specific config",
															"items": map[string]interface{}{
																"description": "ProviderSpecificProperty represents provider specific config property.",
																"properties": map[string]interface{}{
																	"name": map[string]interface{}{
																		"description": "Name of the property",
																		"type":        "string",
																	},
																	"value": map[string]interface{}{
																		"description": "Value of the property",
																		"type":        "string",
																	},
																},
																"type": "object",
															},
															"type": "array",
														},
														"recordTTL": map[string]interface{}{
															"description": "TTL for the record",
															"format":      "int64",
															"type":        "integer",
														},
														"recordType": map[string]interface{}{
															"description": "RecordType type of record, e.g. CNAME, A, SRV, TXT, MX",
															"type":        "string",
														},
														"targets": map[string]interface{}{
															"description": "The targets the DNS service points to",
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
										},
										"type": "object",
									},
									"status": map[string]interface{}{
										"description": "DNSEndpointStatus represents generation observed by the external dns controller.",
										"properties": map[string]interface{}{
											"observedGeneration": map[string]interface{}{
												"description": "The generation observed by by the external-dns controller.",
												"format":      "int64",
												"type":        "integer",
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

	return mutate.MutateCRDDnsendpointsExternaldnsNginxOrg(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;create;update;patch;delete

// CreateCRDApdoslogconfsAppprotectdosF5Com creates the CustomResourceDefinition resource with name apdoslogconfs.appprotectdos.f5.com.
func CreateCRDApdoslogconfsAppprotectdosF5Com(
	parent *platformv1alpha1.IngressComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if parent.Spec.Nginx.Include != true {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:field=nginx.include,value=true,include
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					"controller-gen.kubebuilder.io/version": "v0.9.2",
				},
				"creationTimestamp": nil,
				"name":              "apdoslogconfs.appprotectdos.f5.com",
				"labels": map[string]interface{}{
					"platform.nukleros.io/category": "ingress",
					"platform.nukleros.io/project":  "nginx-ingress-controller",
				},
			},
			"spec": map[string]interface{}{
				"group": "appprotectdos.f5.com",
				"names": map[string]interface{}{
					"kind":     "APDosLogConf",
					"listKind": "APDosLogConfList",
					"plural":   "apdoslogconfs",
					"singular": "apdoslogconf",
				},
				"scope": "Namespaced",
				"versions": []interface{}{
					map[string]interface{}{
						"name": "v1beta1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"description": "APDosLogConf is the Schema for the APDosLogConfs API",
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
										"description": "APDosLogConfSpec defines the desired state of APDosLogConf",
										"properties": map[string]interface{}{
											"content": map[string]interface{}{
												"properties": map[string]interface{}{
													"format": map[string]interface{}{
														"enum": []interface{}{
															"splunk",
															"arcsight",
															"user-defined",
														},
														"type": "string",
													},
													"format_string": map[string]interface{}{
														"type": "string",
													},
													"max_message_size": map[string]interface{}{
														"pattern": "^([1-9]|[1-5][0-9]|6[0-4])k$",
														"type":    "string",
													},
												},
												"type": "object",
											},
											"filter": map[string]interface{}{
												"properties": map[string]interface{}{
													"traffic-mitigation-stats": map[string]interface{}{
														"enum": []interface{}{
															"none",
															"all",
														},
														"default": "all",
														"type":    "string",
													},
													"bad-actors": map[string]interface{}{
														"pattern": "^(none|all|top ([1-9]|[1-9][0-9]|[1-9][0-9]{2,4}|100000))$",
														"default": "top 10",
														"type":    "string",
													},
													"attack-signatures": map[string]interface{}{
														"pattern": "^(none|all|top ([1-9]|[1-9][0-9]|[1-9][0-9]{2,4}|100000))$",
														"default": "top 10",
														"type":    "string",
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
						},
						"served":  true,
						"storage": true,
					},
				},
			},
		},
	}

	return mutate.MutateCRDApdoslogconfsAppprotectdosF5Com(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;create;update;patch;delete

// CreateCRDAplogconfsAppprotectF5Com creates the CustomResourceDefinition resource with name aplogconfs.appprotect.f5.com.
func CreateCRDAplogconfsAppprotectF5Com(
	parent *platformv1alpha1.IngressComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if parent.Spec.Nginx.Include != true {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:field=nginx.include,value=true,include
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					"controller-gen.kubebuilder.io/version": "v0.13.0",
				},
				"name": "aplogconfs.appprotect.f5.com",
				"labels": map[string]interface{}{
					"platform.nukleros.io/category": "ingress",
					"platform.nukleros.io/project":  "nginx-ingress-controller",
				},
			},
			"spec": map[string]interface{}{
				"group": "appprotect.f5.com",
				"names": map[string]interface{}{
					"kind":     "APLogConf",
					"listKind": "APLogConfList",
					"plural":   "aplogconfs",
					"singular": "aplogconf",
				},
				"scope": "Namespaced",
				"versions": []interface{}{
					map[string]interface{}{
						"name": "v1beta1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"description": "APLogConf is the Schema for the APLogConfs API",
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
										"description": "APLogConfSpec defines the desired state of APLogConf",
										"properties": map[string]interface{}{
											"content": map[string]interface{}{
												"properties": map[string]interface{}{
													"escaping_characters": map[string]interface{}{
														"items": map[string]interface{}{
															"properties": map[string]interface{}{
																"from": map[string]interface{}{
																	"type": "string",
																},
																"to": map[string]interface{}{
																	"type": "string",
																},
															},
															"type": "object",
														},
														"type": "array",
													},
													"format": map[string]interface{}{
														"enum": []interface{}{
															"splunk",
															"arcsight",
															"default",
															"user-defined",
															"grpc",
														},
														"type": "string",
													},
													"format_string": map[string]interface{}{
														"type": "string",
													},
													"list_delimiter": map[string]interface{}{
														"type": "string",
													},
													"list_prefix": map[string]interface{}{
														"type": "string",
													},
													"list_suffix": map[string]interface{}{
														"type": "string",
													},
													"max_message_size": map[string]interface{}{
														"pattern": "^([1-9]|[1-5][0-9]|6[0-4])k$",
														"type":    "string",
													},
													"max_request_size": map[string]interface{}{
														"pattern": "^([1-9]|[1-9][0-9]|[1-9][0-9]{2}|[1-9][0-9]{3}|10[0-2][0-9][0-9]|[1-9]k|10k|any)$",
														"type":    "string",
													},
												},
												"type": "object",
											},
											"filter": map[string]interface{}{
												"properties": map[string]interface{}{
													"request_type": map[string]interface{}{
														"enum": []interface{}{
															"all",
															"illegal",
															"blocked",
														},
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
						},
						"served":  true,
						"storage": true,
					},
				},
			},
		},
	}

	return mutate.MutateCRDAplogconfsAppprotectF5Com(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;create;update;patch;delete

// CreateCRDTransportserversK8sNginxOrg creates the CustomResourceDefinition resource with name transportservers.k8s.nginx.org.
func CreateCRDTransportserversK8sNginxOrg(
	parent *platformv1alpha1.IngressComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if parent.Spec.Nginx.Include != true {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:field=nginx.include,value=true,include
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					"controller-gen.kubebuilder.io/version": "v0.21.0",
				},
				"name": "transportservers.k8s.nginx.org",
				"labels": map[string]interface{}{
					"platform.nukleros.io/category": "ingress",
					"platform.nukleros.io/project":  "nginx-ingress-controller",
				},
			},
			"spec": map[string]interface{}{
				"group": "k8s.nginx.org",
				"names": map[string]interface{}{
					"kind":     "TransportServer",
					"listKind": "TransportServerList",
					"plural":   "transportservers",
					"shortNames": []interface{}{
						"ts",
					},
					"singular": "transportserver",
				},
				"scope": "Namespaced",
				"versions": []interface{}{
					map[string]interface{}{
						"additionalPrinterColumns": []interface{}{
							map[string]interface{}{
								"description": "Current state of the TransportServer. If the resource has a valid status, it means it has been validated and accepted by the Ingress Controller.",
								"jsonPath":    ".status.state",
								"name":        "State",
								"type":        "string",
							},
							map[string]interface{}{
								"jsonPath": ".status.reason",
								"name":     "Reason",
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".metadata.creationTimestamp",
								"name":     "Age",
								"type":     "date",
							},
						},
						"name": "v1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"description": "TransportServer defines the TransportServer resource.",
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
										"description": "TransportServerSpec is the spec of the TransportServer resource.",
										"properties": map[string]interface{}{
											"action": map[string]interface{}{
												"description": "The action to perform for a request.",
												"properties": map[string]interface{}{
													"pass": map[string]interface{}{
														"description": "Passes connections/datagrams to an upstream. The upstream with that name must be defined in the resource.",
														"type":        "string",
													},
												},
												"type": "object",
											},
											"host": map[string]interface{}{
												"description": "The host (domain name) of the server. Must be a valid subdomain as defined in RFC 1123, such as my-app or hello.example.com. When using a wildcard domain like *.example.com the domain must be contained in double quotes. The host value needs to be unique among all Ingress and VirtualServer resources.",
												"type":        "string",
											},
											"ingressClassName": map[string]interface{}{
												"description": "Specifies which Ingress Controller must handle the VirtualServer resource.",
												"type":        "string",
											},
											"listener": map[string]interface{}{
												"description": "Sets a custom HTTP and/or HTTPS listener. Valid fields are listener.http and listener.https. Each field must reference the name of a valid listener defined in a GlobalConfiguration resource",
												"properties": map[string]interface{}{
													"name": map[string]interface{}{
														"description": "The name of a listener defined in a GlobalConfiguration resource.",
														"type":        "string",
													},
													"protocol": map[string]interface{}{
														"description": "The protocol of the listener.",
														"type":        "string",
													},
												},
												"type": "object",
											},
											"serverSnippets": map[string]interface{}{
												"description": "Sets a custom snippet in server context. Overrides the server-snippets ConfigMap key.",
												"type":        "string",
											},
											"sessionParameters": map[string]interface{}{
												"description": "The parameters of the session to be used for the Server context",
												"properties": map[string]interface{}{
													"timeout": map[string]interface{}{
														"description": "The timeout between two successive read or write operations on client or proxied server connections. The default is 10m.",
														"type":        "string",
													},
												},
												"type": "object",
											},
											"streamSnippets": map[string]interface{}{
												"description": "Sets a custom snippet in the stream context. Overrides the stream-snippets ConfigMap key.",
												"type":        "string",
											},
											"tls": map[string]interface{}{
												"description": "The TLS termination configuration.",
												"properties": map[string]interface{}{
													"secret": map[string]interface{}{
														"type": "string",
													},
												},
												"type": "object",
											},
											"upstreamParameters": map[string]interface{}{
												"description": "UpstreamParameters defines parameters for an upstream.",
												"properties": map[string]interface{}{
													"connectTimeout": map[string]interface{}{
														"description": "The timeout for establishing a connection with a proxied server.  The default is 60s.",
														"type":        "string",
													},
													"nextUpstream": map[string]interface{}{
														"description": "If a connection to the proxied server cannot be established, determines whether a client connection will be passed to the next server. The default is true.",
														"type":        "boolean",
													},
													"nextUpstreamTimeout": map[string]interface{}{
														"description": "The time allowed to pass a connection to the next server. The default is 0.",
														"type":        "string",
													},
													"nextUpstreamTries": map[string]interface{}{
														"description": "The number of tries for passing a connection to the next server. The default is 0.",
														"type":        "integer",
													},
													"udpRequests": map[string]interface{}{
														"description": "The number of datagrams, after receiving which, the next datagram from the same client starts a new session. The default is 0.",
														"type":        "integer",
													},
													"udpResponses": map[string]interface{}{
														"description": "The number of datagrams expected from the proxied server in response to a client datagram.  By default, the number of datagrams is not limited.",
														"type":        "integer",
													},
												},
												"type": "object",
											},
											"upstreams": map[string]interface{}{
												"description": "A list of upstreams.",
												"items": map[string]interface{}{
													"description": "TransportServerUpstream defines an upstream.",
													"properties": map[string]interface{}{
														"backup": map[string]interface{}{
															"description": "The name of the backup service of type ExternalName. This will be used when the primary servers are unavailable. Note: The parameter cannot be used along with the random, hash or ip_hash load balancing methods.",
															"type":        "string",
														},
														"backupPort": map[string]interface{}{
															"description": "The port of the backup service. The backup port is required if the backup service name is provided. The port must fall into the range 1..65535.",
															"type":        "integer",
														},
														"failTimeout": map[string]interface{}{
															"description": "Sets the number of unsuccessful attempts to communicate with the server that should happen in the duration set by the failTimeout parameter to consider the server unavailable. The default is 1.",
															"type":        "string",
														},
														"healthCheck": map[string]interface{}{
															"description": "The health check configuration for the Upstream. Note: this feature is supported only in NGINX Plus.",
															"properties": map[string]interface{}{
																"enable": map[string]interface{}{
																	"description": "Enables a health check for an upstream server. The default is false.",
																	"type":        "boolean",
																},
																"fails": map[string]interface{}{
																	"description": "The number of consecutive failed health checks of a particular upstream server after which this server will be considered unhealthy. The default is 1.",
																	"type":        "integer",
																},
																"interval": map[string]interface{}{
																	"description": "The interval between two consecutive health checks. The default is 5s.",
																	"type":        "string",
																},
																"jitter": map[string]interface{}{
																	"description": "The time within which each health check will be randomly delayed. By default, there is no delay.",
																	"type":        "string",
																},
																"match": map[string]interface{}{
																	"description": "Controls the data to send and the response to expect for the healthcheck.",
																	"properties": map[string]interface{}{
																		"expect": map[string]interface{}{
																			"description": "A literal string or a regular expression that the data obtained from the server should match. The regular expression is specified with the preceding ~* modifier (for case-insensitive matching), or the ~ modifier (for case-sensitive matching). NGINX Ingress Controller validates a regular expression using the RE2 syntax.",
																			"type":        "string",
																		},
																		"send": map[string]interface{}{
																			"description": "A string to send to an upstream server.",
																			"type":        "string",
																		},
																	},
																	"type": "object",
																},
																"passes": map[string]interface{}{
																	"description": "The number of consecutive passed health checks of a particular upstream server after which the server will be considered healthy. The default is 1.",
																	"type":        "integer",
																},
																"port": map[string]interface{}{
																	"description": "The port used for health check requests. By default, the server port is used. Note: in contrast with the port of the upstream, this port is not a service port, but a port of a pod.",
																	"type":        "integer",
																},
																"timeout": map[string]interface{}{
																	"description": "This overrides the timeout set by proxy_timeout which is set in SessionParameters for health checks. The default value is 5s.",
																	"type":        "string",
																},
															},
															"type": "object",
														},
														"loadBalancingMethod": map[string]interface{}{
															"description": "The method used to load balance the upstream servers. By default, connections are distributed between the servers using a weighted round-robin balancing method.",
															"type":        "string",
														},
														"maxConns": map[string]interface{}{
															"description": "Sets the time during which the specified number of unsuccessful attempts to communicate with the server should happen to consider the server unavailable and the period of time the server will be considered unavailable. The default is 10s.",
															"type":        "integer",
														},
														"maxFails": map[string]interface{}{
															"description": "Sets the number of maximum connections to the proxied server. Default value is zero, meaning there is no limit. The default is 0.",
															"type":        "integer",
														},
														"name": map[string]interface{}{
															"description": "The name of the upstream. Must be a valid DNS label as defined in RFC 1035. For example, hello and upstream-123 are valid. The name must be unique among all upstreams of the resource.",
															"type":        "string",
														},
														"port": map[string]interface{}{
															"description": "The port of the service. If the service doesn’t define that port, NGINX will assume the service has zero endpoints and close client connections/ignore datagrams. The port must fall into the range 1..65535.",
															"type":        "integer",
														},
														"service": map[string]interface{}{
															"description": "The name of a service. The service must belong to the same namespace as the resource. If the service doesn’t exist, NGINX will assume the service has zero endpoints and close client connections/ignore datagrams.",
															"type":        "string",
														},
													},
													"type": "object",
												},
												"type": "array",
											},
										},
										"type": "object",
									},
									"status": map[string]interface{}{
										"description": "The status of the TransportServer resource",
										"properties": map[string]interface{}{
											"message": map[string]interface{}{
												"description": "The message of the current state of the resource. It can contain more detailed information about the reason.",
												"type":        "string",
											},
											"reason": map[string]interface{}{
												"description": "The reason of the current state of the resource.",
												"type":        "string",
											},
											"state": map[string]interface{}{
												"description": "Represents the current state of the resource. Possible values: Valid (resource validated and accepted), Invalid (validation failed or config reload failed), or Warning (validated but may work in degraded state).",
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
			},
		},
	}

	return mutate.MutateCRDTransportserversK8sNginxOrg(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;create;update;patch;delete

// CreateCRDVirtualserverroutesK8sNginxOrg creates the CustomResourceDefinition resource with name virtualserverroutes.k8s.nginx.org.
func CreateCRDVirtualserverroutesK8sNginxOrg(
	parent *platformv1alpha1.IngressComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if parent.Spec.Nginx.Include != true {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:field=nginx.include,value=true,include
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					"controller-gen.kubebuilder.io/version": "v0.21.0",
				},
				"name": "virtualserverroutes.k8s.nginx.org",
				"labels": map[string]interface{}{
					"platform.nukleros.io/category": "ingress",
					"platform.nukleros.io/project":  "nginx-ingress-controller",
				},
			},
			"spec": map[string]interface{}{
				"group": "k8s.nginx.org",
				"names": map[string]interface{}{
					"kind":     "VirtualServerRoute",
					"listKind": "VirtualServerRouteList",
					"plural":   "virtualserverroutes",
					"shortNames": []interface{}{
						"vsr",
					},
					"singular": "virtualserverroute",
				},
				"scope": "Namespaced",
				"versions": []interface{}{
					map[string]interface{}{
						"additionalPrinterColumns": []interface{}{
							map[string]interface{}{
								"description": "Current state of the VirtualServerRoute. If the resource has a valid status, it means it has been validated and accepted by the Ingress Controller.",
								"jsonPath":    ".status.state",
								"name":        "State",
								"type":        "string",
							},
							map[string]interface{}{
								"jsonPath": ".spec.host",
								"name":     "Host",
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".status.externalEndpoints[*].ip",
								"name":     "IP",
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".status.externalEndpoints[*].hostname",
								"name":     "ExternalHostname",
								"priority": 1,
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".status.externalEndpoints[*].ports",
								"name":     "Ports",
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".metadata.creationTimestamp",
								"name":     "Age",
								"type":     "date",
							},
						},
						"name": "v1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"description": "VirtualServerRoute defines the VirtualServerRoute resource.",
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
										"description": "VirtualServerRouteSpec is the spec of the VirtualServerRoute resource.",
										"properties": map[string]interface{}{
											"host": map[string]interface{}{
												"description": "The host (domain name) of the server. Must be a valid subdomain as defined in RFC 1123, such as my-app or hello.example.com. When using a wildcard domain like *.example.com the domain must be contained in double quotes. Must be the same as the host of the VirtualServer that references this resource.",
												"type":        "string",
											},
											"ingressClassName": map[string]interface{}{
												"description": "Specifies which Ingress Controller must handle the VirtualServerRoute resource. Must be the same as the ingressClassName of the VirtualServer that references this resource.",
												"type":        "string",
											},
											"subroutes": map[string]interface{}{
												"description": "A list of subroutes.",
												"items": map[string]interface{}{
													"description": "Route defines a route.",
													"properties": map[string]interface{}{
														"action": map[string]interface{}{
															"description": "The default action to perform for a request.",
															"properties": map[string]interface{}{
																"pass": map[string]interface{}{
																	"description": "Passes requests to an upstream. The upstream with that name must be defined in the resource.",
																	"type":        "string",
																},
																"proxy": map[string]interface{}{
																	"description": "Passes requests to an upstream with the ability to modify the request/response (for example, rewrite the URI or modify the headers).",
																	"properties": map[string]interface{}{
																		"requestHeaders": map[string]interface{}{
																			"description": "The request headers modifications.",
																			"properties": map[string]interface{}{
																				"pass": map[string]interface{}{
																					"description": "Passes the original request headers to the proxied upstream server.  Default is true.",
																					"type":        "boolean",
																				},
																				"set": map[string]interface{}{
																					"description": "Allows redefining or appending fields to present request headers passed to the proxied upstream servers.",
																					"items": map[string]interface{}{
																						"description": "Header defines an HTTP Header.",
																						"properties": map[string]interface{}{
																							"name": map[string]interface{}{
																								"description": "The name of the header.",
																								"type":        "string",
																							},
																							"value": map[string]interface{}{
																								"description": "The value of the header.",
																								"type":        "string",
																							},
																						},
																						"type": "object",
																					},
																					"type": "array",
																				},
																			},
																			"type": "object",
																		},
																		"responseHeaders": map[string]interface{}{
																			"description": "The response headers modifications.",
																			"properties": map[string]interface{}{
																				"add": map[string]interface{}{
																					"description": "Adds headers to the response to the client.",
																					"items": map[string]interface{}{
																						"description": "AddHeader defines an HTTP Header with an optional Always field to use with the add_header NGINX directive.",
																						"properties": map[string]interface{}{
																							"always": map[string]interface{}{
																								"description": "If set to true, add the header regardless of the response status code**. Default is false.",
																								"type":        "boolean",
																							},
																							"name": map[string]interface{}{
																								"description": "The name of the header.",
																								"type":        "string",
																							},
																							"value": map[string]interface{}{
																								"description": "The value of the header.",
																								"type":        "string",
																							},
																						},
																						"type": "object",
																					},
																					"type": "array",
																				},
																				"hide": map[string]interface{}{
																					"description": "The headers that will not be passed* in the response to the client from a proxied upstream server.",
																					"items": map[string]interface{}{
																						"type": "string",
																					},
																					"type": "array",
																				},
																				"ignore": map[string]interface{}{
																					"description": "Disables processing of certain headers** to the client from a proxied upstream server.",
																					"items": map[string]interface{}{
																						"type": "string",
																					},
																					"type": "array",
																				},
																				"pass": map[string]interface{}{
																					"description": "Allows passing the hidden header fields* to the client from a proxied upstream server.",
																					"items": map[string]interface{}{
																						"type": "string",
																					},
																					"type": "array",
																				},
																			},
																			"type": "object",
																		},
																		"rewritePath": map[string]interface{}{
																			"description": "The rewritten URI. If the route path is a regular expression – starts with ~ – the rewritePath can include capture groups with $1-9. For example $1 for the first group, and so on. For more information, check the rewrite example.",
																			"type":        "string",
																		},
																		"upstream": map[string]interface{}{
																			"description": "The name of the upstream which the requests will be proxied to. The upstream with that name must be defined in the resource.",
																			"type":        "string",
																		},
																	},
																	"type": "object",
																},
																"redirect": map[string]interface{}{
																	"description": "Redirects requests to a provided URL.",
																	"properties": map[string]interface{}{
																		"code": map[string]interface{}{
																			"description": "The status code of a redirect. The allowed values are: 301, 302, 307 or 308. The default is 301.",
																			"type":        "integer",
																		},
																		"url": map[string]interface{}{
																			"description": "The URL to redirect the request to. Supported NGINX variables: $scheme, $http_x_forwarded_proto, $request_uri or $host. Variables must be enclosed in curly braces. For example: ${host}${request_uri}.",
																			"type":        "string",
																		},
																	},
																	"type": "object",
																},
																"return": map[string]interface{}{
																	"description": "Returns a preconfigured response.",
																	"properties": map[string]interface{}{
																		"body": map[string]interface{}{
																			"description": `The body of the response. Supports NGINX variables*. Variables must be enclosed in curly brackets. For example: Request is ${request_uri}\n.`,
																			"type":        "string",
																		},
																		"code": map[string]interface{}{
																			"description": "The status code of the response. The allowed values are: 2XX, 4XX or 5XX. The default is 200.",
																			"type":        "integer",
																		},
																		"headers": map[string]interface{}{
																			"description": "The custom headers of the response.",
																			"items": map[string]interface{}{
																				"description": "Header defines an HTTP Header.",
																				"properties": map[string]interface{}{
																					"name": map[string]interface{}{
																						"description": "The name of the header.",
																						"type":        "string",
																					},
																					"value": map[string]interface{}{
																						"description": "The value of the header.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"type": "array",
																		},
																		"type": map[string]interface{}{
																			"description": "The MIME type of the response. The default is text/plain.",
																			"type":        "string",
																		},
																	},
																	"type": "object",
																},
															},
															"type": "object",
														},
														"add-header-inherit": map[string]interface{}{
															"description": "Controls header inheritance behavior at the location level. Allowed values are: on, off, merge. When set to \"merge\", headers from this context are merged with headers in child contexts. When set to \"on\", standard NGINX inheritance applies. When set to \"off\", no headers are inherited from parent contexts.",
															"enum": []interface{}{
																"on",
																"off",
																"merge",
															},
															"type": "string",
														},
														"dos": map[string]interface{}{
															"description": "A reference to a DosProtectedResource, setting this enables DOS protection of the VirtualServer route.",
															"type":        "string",
														},
														"errorPages": map[string]interface{}{
															"description": "The custom responses for error codes. NGINX will use those responses instead of returning the error responses from the upstream servers or the default responses generated by NGINX. A custom response can be a redirect or a canned response. For example, a redirect to another URL if an upstream server responded with a 404 status code.",
															"items": map[string]interface{}{
																"description": "ErrorPage defines an ErrorPage in a Route.",
																"properties": map[string]interface{}{
																	"codes": map[string]interface{}{
																		"description": "A list of error status codes.",
																		"items": map[string]interface{}{
																			"type": "integer",
																		},
																		"type": "array",
																	},
																	"redirect": map[string]interface{}{
																		"description": "The canned response action for the given status codes.",
																		"properties": map[string]interface{}{
																			"code": map[string]interface{}{
																				"description": "The status code of a redirect. The allowed values are: 301, 302, 307 or 308. The default is 301.",
																				"type":        "integer",
																			},
																			"url": map[string]interface{}{
																				"description": "The URL to redirect the request to. Supported NGINX variables: $scheme, $http_x_forwarded_proto, $request_uri or $host. Variables must be enclosed in curly braces. For example: ${host}${request_uri}.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																	"return": map[string]interface{}{
																		"description": "The redirect action for the given status codes.",
																		"properties": map[string]interface{}{
																			"body": map[string]interface{}{
																				"description": `The body of the response. Supports NGINX variables*. Variables must be enclosed in curly brackets. For example: Request is ${request_uri}\n.`,
																				"type":        "string",
																			},
																			"code": map[string]interface{}{
																				"description": "The status code of the response. The allowed values are: 2XX, 4XX or 5XX. The default is 200.",
																				"type":        "integer",
																			},
																			"headers": map[string]interface{}{
																				"description": "The custom headers of the response.",
																				"items": map[string]interface{}{
																					"description": "Header defines an HTTP Header.",
																					"properties": map[string]interface{}{
																						"name": map[string]interface{}{
																							"description": "The name of the header.",
																							"type":        "string",
																						},
																						"value": map[string]interface{}{
																							"description": "The value of the header.",
																							"type":        "string",
																						},
																					},
																					"type": "object",
																				},
																				"type": "array",
																			},
																			"type": map[string]interface{}{
																				"description": "The MIME type of the response. The default is text/plain.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
															"type": "array",
														},
														"location-snippets": map[string]interface{}{
															"description": "Sets a custom snippet in the location context. Overrides the location-snippets ConfigMap key.",
															"type":        "string",
														},
														"matches": map[string]interface{}{
															"description": "The matching rules for advanced content-based routing. Requires the default Action or Splits. Unmatched requests will be handled by the default Action or Splits.",
															"items": map[string]interface{}{
																"description": "Match defines a match.",
																"properties": map[string]interface{}{
																	"action": map[string]interface{}{
																		"description": "The action to perform for a request.",
																		"properties": map[string]interface{}{
																			"pass": map[string]interface{}{
																				"description": "Passes requests to an upstream. The upstream with that name must be defined in the resource.",
																				"type":        "string",
																			},
																			"proxy": map[string]interface{}{
																				"description": "Passes requests to an upstream with the ability to modify the request/response (for example, rewrite the URI or modify the headers).",
																				"properties": map[string]interface{}{
																					"requestHeaders": map[string]interface{}{
																						"description": "The request headers modifications.",
																						"properties": map[string]interface{}{
																							"pass": map[string]interface{}{
																								"description": "Passes the original request headers to the proxied upstream server.  Default is true.",
																								"type":        "boolean",
																							},
																							"set": map[string]interface{}{
																								"description": "Allows redefining or appending fields to present request headers passed to the proxied upstream servers.",
																								"items": map[string]interface{}{
																									"description": "Header defines an HTTP Header.",
																									"properties": map[string]interface{}{
																										"name": map[string]interface{}{
																											"description": "The name of the header.",
																											"type":        "string",
																										},
																										"value": map[string]interface{}{
																											"description": "The value of the header.",
																											"type":        "string",
																										},
																									},
																									"type": "object",
																								},
																								"type": "array",
																							},
																						},
																						"type": "object",
																					},
																					"responseHeaders": map[string]interface{}{
																						"description": "The response headers modifications.",
																						"properties": map[string]interface{}{
																							"add": map[string]interface{}{
																								"description": "Adds headers to the response to the client.",
																								"items": map[string]interface{}{
																									"description": "AddHeader defines an HTTP Header with an optional Always field to use with the add_header NGINX directive.",
																									"properties": map[string]interface{}{
																										"always": map[string]interface{}{
																											"description": "If set to true, add the header regardless of the response status code**. Default is false.",
																											"type":        "boolean",
																										},
																										"name": map[string]interface{}{
																											"description": "The name of the header.",
																											"type":        "string",
																										},
																										"value": map[string]interface{}{
																											"description": "The value of the header.",
																											"type":        "string",
																										},
																									},
																									"type": "object",
																								},
																								"type": "array",
																							},
																							"hide": map[string]interface{}{
																								"description": "The headers that will not be passed* in the response to the client from a proxied upstream server.",
																								"items": map[string]interface{}{
																									"type": "string",
																								},
																								"type": "array",
																							},
																							"ignore": map[string]interface{}{
																								"description": "Disables processing of certain headers** to the client from a proxied upstream server.",
																								"items": map[string]interface{}{
																									"type": "string",
																								},
																								"type": "array",
																							},
																							"pass": map[string]interface{}{
																								"description": "Allows passing the hidden header fields* to the client from a proxied upstream server.",
																								"items": map[string]interface{}{
																									"type": "string",
																								},
																								"type": "array",
																							},
																						},
																						"type": "object",
																					},
																					"rewritePath": map[string]interface{}{
																						"description": "The rewritten URI. If the route path is a regular expression – starts with ~ – the rewritePath can include capture groups with $1-9. For example $1 for the first group, and so on. For more information, check the rewrite example.",
																						"type":        "string",
																					},
																					"upstream": map[string]interface{}{
																						"description": "The name of the upstream which the requests will be proxied to. The upstream with that name must be defined in the resource.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"redirect": map[string]interface{}{
																				"description": "Redirects requests to a provided URL.",
																				"properties": map[string]interface{}{
																					"code": map[string]interface{}{
																						"description": "The status code of a redirect. The allowed values are: 301, 302, 307 or 308. The default is 301.",
																						"type":        "integer",
																					},
																					"url": map[string]interface{}{
																						"description": "The URL to redirect the request to. Supported NGINX variables: $scheme, $http_x_forwarded_proto, $request_uri or $host. Variables must be enclosed in curly braces. For example: ${host}${request_uri}.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"return": map[string]interface{}{
																				"description": "Returns a preconfigured response.",
																				"properties": map[string]interface{}{
																					"body": map[string]interface{}{
																						"description": `The body of the response. Supports NGINX variables*. Variables must be enclosed in curly brackets. For example: Request is ${request_uri}\n.`,
																						"type":        "string",
																					},
																					"code": map[string]interface{}{
																						"description": "The status code of the response. The allowed values are: 2XX, 4XX or 5XX. The default is 200.",
																						"type":        "integer",
																					},
																					"headers": map[string]interface{}{
																						"description": "The custom headers of the response.",
																						"items": map[string]interface{}{
																							"description": "Header defines an HTTP Header.",
																							"properties": map[string]interface{}{
																								"name": map[string]interface{}{
																									"description": "The name of the header.",
																									"type":        "string",
																								},
																								"value": map[string]interface{}{
																									"description": "The value of the header.",
																									"type":        "string",
																								},
																							},
																							"type": "object",
																						},
																						"type": "array",
																					},
																					"type": map[string]interface{}{
																						"description": "The MIME type of the response. The default is text/plain.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																	"conditions": map[string]interface{}{
																		"description": "A list of conditions. Must include at least 1 condition.",
																		"items": map[string]interface{}{
																			"description": "Condition defines a condition in a MatchRule.",
																			"properties": map[string]interface{}{
																				"argument": map[string]interface{}{
																					"description": "The name of an argument. Must consist of alphanumeric characters or _.",
																					"type":        "string",
																				},
																				"cookie": map[string]interface{}{
																					"description": "The name of a cookie. Must consist of alphanumeric characters or _.",
																					"type":        "string",
																				},
																				"header": map[string]interface{}{
																					"description": "The name of a header. Must consist of alphanumeric characters or -.",
																					"type":        "string",
																				},
																				"value": map[string]interface{}{
																					"description": "The value to match the condition against.",
																					"type":        "string",
																				},
																				"variable": map[string]interface{}{
																					"description": "The name of an NGINX variable. Must start with $.",
																					"type":        "string",
																				},
																			},
																			"type": "object",
																		},
																		"type": "array",
																	},
																	"splits": map[string]interface{}{
																		"description": "The splits configuration for traffic splitting. Must include at least 2 splits.",
																		"items": map[string]interface{}{
																			"description": "Split defines a split.",
																			"properties": map[string]interface{}{
																				"action": map[string]interface{}{
																					"description": "The action to perform for a request.",
																					"properties": map[string]interface{}{
																						"pass": map[string]interface{}{
																							"description": "Passes requests to an upstream. The upstream with that name must be defined in the resource.",
																							"type":        "string",
																						},
																						"proxy": map[string]interface{}{
																							"description": "Passes requests to an upstream with the ability to modify the request/response (for example, rewrite the URI or modify the headers).",
																							"properties": map[string]interface{}{
																								"requestHeaders": map[string]interface{}{
																									"description": "The request headers modifications.",
																									"properties": map[string]interface{}{
																										"pass": map[string]interface{}{
																											"description": "Passes the original request headers to the proxied upstream server.  Default is true.",
																											"type":        "boolean",
																										},
																										"set": map[string]interface{}{
																											"description": "Allows redefining or appending fields to present request headers passed to the proxied upstream servers.",
																											"items": map[string]interface{}{
																												"description": "Header defines an HTTP Header.",
																												"properties": map[string]interface{}{
																													"name": map[string]interface{}{
																														"description": "The name of the header.",
																														"type":        "string",
																													},
																													"value": map[string]interface{}{
																														"description": "The value of the header.",
																														"type":        "string",
																													},
																												},
																												"type": "object",
																											},
																											"type": "array",
																										},
																									},
																									"type": "object",
																								},
																								"responseHeaders": map[string]interface{}{
																									"description": "The response headers modifications.",
																									"properties": map[string]interface{}{
																										"add": map[string]interface{}{
																											"description": "Adds headers to the response to the client.",
																											"items": map[string]interface{}{
																												"description": "AddHeader defines an HTTP Header with an optional Always field to use with the add_header NGINX directive.",
																												"properties": map[string]interface{}{
																													"always": map[string]interface{}{
																														"description": "If set to true, add the header regardless of the response status code**. Default is false.",
																														"type":        "boolean",
																													},
																													"name": map[string]interface{}{
																														"description": "The name of the header.",
																														"type":        "string",
																													},
																													"value": map[string]interface{}{
																														"description": "The value of the header.",
																														"type":        "string",
																													},
																												},
																												"type": "object",
																											},
																											"type": "array",
																										},
																										"hide": map[string]interface{}{
																											"description": "The headers that will not be passed* in the response to the client from a proxied upstream server.",
																											"items": map[string]interface{}{
																												"type": "string",
																											},
																											"type": "array",
																										},
																										"ignore": map[string]interface{}{
																											"description": "Disables processing of certain headers** to the client from a proxied upstream server.",
																											"items": map[string]interface{}{
																												"type": "string",
																											},
																											"type": "array",
																										},
																										"pass": map[string]interface{}{
																											"description": "Allows passing the hidden header fields* to the client from a proxied upstream server.",
																											"items": map[string]interface{}{
																												"type": "string",
																											},
																											"type": "array",
																										},
																									},
																									"type": "object",
																								},
																								"rewritePath": map[string]interface{}{
																									"description": "The rewritten URI. If the route path is a regular expression – starts with ~ – the rewritePath can include capture groups with $1-9. For example $1 for the first group, and so on. For more information, check the rewrite example.",
																									"type":        "string",
																								},
																								"upstream": map[string]interface{}{
																									"description": "The name of the upstream which the requests will be proxied to. The upstream with that name must be defined in the resource.",
																									"type":        "string",
																								},
																							},
																							"type": "object",
																						},
																						"redirect": map[string]interface{}{
																							"description": "Redirects requests to a provided URL.",
																							"properties": map[string]interface{}{
																								"code": map[string]interface{}{
																									"description": "The status code of a redirect. The allowed values are: 301, 302, 307 or 308. The default is 301.",
																									"type":        "integer",
																								},
																								"url": map[string]interface{}{
																									"description": "The URL to redirect the request to. Supported NGINX variables: $scheme, $http_x_forwarded_proto, $request_uri or $host. Variables must be enclosed in curly braces. For example: ${host}${request_uri}.",
																									"type":        "string",
																								},
																							},
																							"type": "object",
																						},
																						"return": map[string]interface{}{
																							"description": "Returns a preconfigured response.",
																							"properties": map[string]interface{}{
																								"body": map[string]interface{}{
																									"description": `The body of the response. Supports NGINX variables*. Variables must be enclosed in curly brackets. For example: Request is ${request_uri}\n.`,
																									"type":        "string",
																								},
																								"code": map[string]interface{}{
																									"description": "The status code of the response. The allowed values are: 2XX, 4XX or 5XX. The default is 200.",
																									"type":        "integer",
																								},
																								"headers": map[string]interface{}{
																									"description": "The custom headers of the response.",
																									"items": map[string]interface{}{
																										"description": "Header defines an HTTP Header.",
																										"properties": map[string]interface{}{
																											"name": map[string]interface{}{
																												"description": "The name of the header.",
																												"type":        "string",
																											},
																											"value": map[string]interface{}{
																												"description": "The value of the header.",
																												"type":        "string",
																											},
																										},
																										"type": "object",
																									},
																									"type": "array",
																								},
																								"type": map[string]interface{}{
																									"description": "The MIME type of the response. The default is text/plain.",
																									"type":        "string",
																								},
																							},
																							"type": "object",
																						},
																					},
																					"type": "object",
																				},
																				"weight": map[string]interface{}{
																					"description": "The weight of an action. Must fall into the range 0..100. The sum of the weights of all splits must be equal to 100.",
																					"type":        "integer",
																				},
																			},
																			"type": "object",
																		},
																		"type": "array",
																	},
																},
																"type": "object",
															},
															"type": "array",
														},
														"path": map[string]interface{}{
															"description": `The path of the route. NGINX will match it against the URI of a request. Possible values are: a prefix ( / , /path ), a longest prefix match ( ^~/images/ ), an exact match ( =/exact/match ), a case-insensitive regular expression ( ~*^/Bar.*\.jpg ) or a case-sensitive regular expression ( ~^/foo.*\.jpg ). In the case of a prefix match (must start with / ), a longest prefix match (must start with ^~ ) or an exact match (must start with = ), the path must not include any whitespace characters, { , } or ;. In the case of the regex matches, all double quotes " must be escaped and the match can’t end in an unescaped backslash \. The path must be unique among the paths of all routes of the VirtualServer. Check the location directive for more information.`,
															"type":        "string",
														},
														"policies": map[string]interface{}{
															"description": "A list of policies. The policies override the policies of the same type defined in the spec of the VirtualServer.",
															"items": map[string]interface{}{
																"description": "PolicyReference references a policy by name and an optional namespace.",
																"properties": map[string]interface{}{
																	"name": map[string]interface{}{
																		"description": "The name of a policy. If the policy doesn’t exist or invalid, NGINX will respond with an error response with the 500 status code.",
																		"type":        "string",
																	},
																	"namespace": map[string]interface{}{
																		"description": "The namespace of a policy. If not specified, the namespace of the VirtualServer resource is used.",
																		"type":        "string",
																	},
																},
																"type": "object",
															},
															"type": "array",
														},
														"route": map[string]interface{}{
															"description": "The name of a VirtualServerRoute resource that defines this route. If the VirtualServerRoute belongs to a different namespace than the VirtualServer, you need to include the namespace. For example, tea-namespace/tea.",
															"type":        "string",
														},
														"routeSelector": map[string]interface{}{
															"description": "The RouteSelector allows selecting VirtualServerRoute resources using label selectors.",
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
														"splits": map[string]interface{}{
															"description": "The default splits configuration for traffic splitting. Must include at least 2 splits.",
															"items": map[string]interface{}{
																"description": "Split defines a split.",
																"properties": map[string]interface{}{
																	"action": map[string]interface{}{
																		"description": "The action to perform for a request.",
																		"properties": map[string]interface{}{
																			"pass": map[string]interface{}{
																				"description": "Passes requests to an upstream. The upstream with that name must be defined in the resource.",
																				"type":        "string",
																			},
																			"proxy": map[string]interface{}{
																				"description": "Passes requests to an upstream with the ability to modify the request/response (for example, rewrite the URI or modify the headers).",
																				"properties": map[string]interface{}{
																					"requestHeaders": map[string]interface{}{
																						"description": "The request headers modifications.",
																						"properties": map[string]interface{}{
																							"pass": map[string]interface{}{
																								"description": "Passes the original request headers to the proxied upstream server.  Default is true.",
																								"type":        "boolean",
																							},
																							"set": map[string]interface{}{
																								"description": "Allows redefining or appending fields to present request headers passed to the proxied upstream servers.",
																								"items": map[string]interface{}{
																									"description": "Header defines an HTTP Header.",
																									"properties": map[string]interface{}{
																										"name": map[string]interface{}{
																											"description": "The name of the header.",
																											"type":        "string",
																										},
																										"value": map[string]interface{}{
																											"description": "The value of the header.",
																											"type":        "string",
																										},
																									},
																									"type": "object",
																								},
																								"type": "array",
																							},
																						},
																						"type": "object",
																					},
																					"responseHeaders": map[string]interface{}{
																						"description": "The response headers modifications.",
																						"properties": map[string]interface{}{
																							"add": map[string]interface{}{
																								"description": "Adds headers to the response to the client.",
																								"items": map[string]interface{}{
																									"description": "AddHeader defines an HTTP Header with an optional Always field to use with the add_header NGINX directive.",
																									"properties": map[string]interface{}{
																										"always": map[string]interface{}{
																											"description": "If set to true, add the header regardless of the response status code**. Default is false.",
																											"type":        "boolean",
																										},
																										"name": map[string]interface{}{
																											"description": "The name of the header.",
																											"type":        "string",
																										},
																										"value": map[string]interface{}{
																											"description": "The value of the header.",
																											"type":        "string",
																										},
																									},
																									"type": "object",
																								},
																								"type": "array",
																							},
																							"hide": map[string]interface{}{
																								"description": "The headers that will not be passed* in the response to the client from a proxied upstream server.",
																								"items": map[string]interface{}{
																									"type": "string",
																								},
																								"type": "array",
																							},
																							"ignore": map[string]interface{}{
																								"description": "Disables processing of certain headers** to the client from a proxied upstream server.",
																								"items": map[string]interface{}{
																									"type": "string",
																								},
																								"type": "array",
																							},
																							"pass": map[string]interface{}{
																								"description": "Allows passing the hidden header fields* to the client from a proxied upstream server.",
																								"items": map[string]interface{}{
																									"type": "string",
																								},
																								"type": "array",
																							},
																						},
																						"type": "object",
																					},
																					"rewritePath": map[string]interface{}{
																						"description": "The rewritten URI. If the route path is a regular expression – starts with ~ – the rewritePath can include capture groups with $1-9. For example $1 for the first group, and so on. For more information, check the rewrite example.",
																						"type":        "string",
																					},
																					"upstream": map[string]interface{}{
																						"description": "The name of the upstream which the requests will be proxied to. The upstream with that name must be defined in the resource.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"redirect": map[string]interface{}{
																				"description": "Redirects requests to a provided URL.",
																				"properties": map[string]interface{}{
																					"code": map[string]interface{}{
																						"description": "The status code of a redirect. The allowed values are: 301, 302, 307 or 308. The default is 301.",
																						"type":        "integer",
																					},
																					"url": map[string]interface{}{
																						"description": "The URL to redirect the request to. Supported NGINX variables: $scheme, $http_x_forwarded_proto, $request_uri or $host. Variables must be enclosed in curly braces. For example: ${host}${request_uri}.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"return": map[string]interface{}{
																				"description": "Returns a preconfigured response.",
																				"properties": map[string]interface{}{
																					"body": map[string]interface{}{
																						"description": `The body of the response. Supports NGINX variables*. Variables must be enclosed in curly brackets. For example: Request is ${request_uri}\n.`,
																						"type":        "string",
																					},
																					"code": map[string]interface{}{
																						"description": "The status code of the response. The allowed values are: 2XX, 4XX or 5XX. The default is 200.",
																						"type":        "integer",
																					},
																					"headers": map[string]interface{}{
																						"description": "The custom headers of the response.",
																						"items": map[string]interface{}{
																							"description": "Header defines an HTTP Header.",
																							"properties": map[string]interface{}{
																								"name": map[string]interface{}{
																									"description": "The name of the header.",
																									"type":        "string",
																								},
																								"value": map[string]interface{}{
																									"description": "The value of the header.",
																									"type":        "string",
																								},
																							},
																							"type": "object",
																						},
																						"type": "array",
																					},
																					"type": map[string]interface{}{
																						"description": "The MIME type of the response. The default is text/plain.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																	"weight": map[string]interface{}{
																		"description": "The weight of an action. Must fall into the range 0..100. The sum of the weights of all splits must be equal to 100.",
																		"type":        "integer",
																	},
																},
																"type": "object",
															},
															"type": "array",
														},
													},
													"type": "object",
												},
												"type": "array",
											},
											"upstreams": map[string]interface{}{
												"description": "A list of upstreams.",
												"items": map[string]interface{}{
													"description": "Upstream defines an upstream.",
													"properties": map[string]interface{}{
														"backup": map[string]interface{}{
															"description": "The name of the backup service of type ExternalName. This will be used when the primary servers are unavailable. Note: The parameter cannot be used along with the random, hash or ip_hash load balancing methods.",
															"type":        "string",
														},
														"backupPort": map[string]interface{}{
															"description": "The port of the backup service. The backup port is required if the backup service name is provided. The port must fall into the range 1..65535.",
															"type":        "integer",
														},
														"buffer-size": map[string]interface{}{
															"description": "Sets the size of the buffer used for reading the first part of a response received from the upstream server. The default is set in the proxy-buffer-size ConfigMap key.",
															"type":        "string",
														},
														"buffering": map[string]interface{}{
															"description": "Enables buffering of responses from the upstream server.  The default is set in the proxy-buffering ConfigMap key.",
															"type":        "boolean",
														},
														"buffers": map[string]interface{}{
															"description": "Configures the buffers used for reading a response from the upstream server for a single connection.",
															"properties": map[string]interface{}{
																"number": map[string]interface{}{
																	"description": "Configures the number of buffers. The default is set in the proxy-buffers ConfigMap key.",
																	"type":        "integer",
																},
																"size": map[string]interface{}{
																	"description": "Configures the size of a buffer. The default is set in the proxy-buffers ConfigMap key.",
																	"type":        "string",
																},
															},
															"type": "object",
														},
														"busy-buffers-size": map[string]interface{}{
															"description": "Sets the size of the buffers used for reading a response from the upstream server when the proxy_buffering is enabled. The default is set in the proxy-busy-buffers-size ConfigMap key.'",
															"type":        "string",
														},
														"client-body-buffer-size": map[string]interface{}{
															"description": `ClientBodyBufferSize sets the size of the buffer used for reading the client request body. Must be specified as a number followed by:
'k' for kilobytes or 'm' for megabytes.
Examples: "10m" or "512k".`,
															"pattern": `^\d+[kKmM]?$`,
															"type":    "string",
														},
														"client-max-body-size": map[string]interface{}{
															"description": "Sets the maximum allowed size of the client request body. The default is set in the client-max-body-size ConfigMap key.",
															"type":        "string",
														},
														"connect-timeout": map[string]interface{}{
															"description": "The timeout for establishing a connection with an upstream server. The default is specified in the proxy-connect-timeout ConfigMap key.",
															"type":        "string",
														},
														"fail-timeout": map[string]interface{}{
															"description": "The time during which the specified number of unsuccessful attempts to communicate with an upstream server should happen to consider the server unavailable. The default is set in the fail-timeout ConfigMap key.",
															"type":        "string",
														},
														"healthCheck": map[string]interface{}{
															"description": "The health check configuration for the Upstream. Note: this feature is supported only in NGINX Plus.",
															"properties": map[string]interface{}{
																"connect-timeout": map[string]interface{}{
																	"description": "The timeout for establishing a connection with an upstream server. By default, the connect-timeout of the upstream is used.",
																	"type":        "string",
																},
																"enable": map[string]interface{}{
																	"description": "Enables a health check for an upstream server. The default is false.",
																	"type":        "boolean",
																},
																"fails": map[string]interface{}{
																	"description": "The number of consecutive failed health checks of a particular upstream server after which this server will be considered unhealthy. The default is 1.",
																	"type":        "integer",
																},
																"grpcService": map[string]interface{}{
																	"description": "The gRPC service to be monitored on the upstream server. Only valid on gRPC type upstreams.",
																	"type":        "string",
																},
																"grpcStatus": map[string]interface{}{
																	"description": "The expected gRPC status code of the upstream server response to the Check method. Configure this field only if your gRPC services do not implement the gRPC health checking protocol. For example, configure 12 if the upstream server responds with 12 (UNIMPLEMENTED) status code. Only valid on gRPC type upstreams.",
																	"type":        "integer",
																},
																"headers": map[string]interface{}{
																	"description": "The request headers used for health check requests. NGINX Plus always sets the Host, User-Agent and Connection headers for health check requests.",
																	"items": map[string]interface{}{
																		"description": "Header defines an HTTP Header.",
																		"properties": map[string]interface{}{
																			"name": map[string]interface{}{
																				"description": "The name of the header.",
																				"type":        "string",
																			},
																			"value": map[string]interface{}{
																				"description": "The value of the header.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																	"type": "array",
																},
																"interval": map[string]interface{}{
																	"description": "The interval between two consecutive health checks. The default is 5s.",
																	"type":        "string",
																},
																"jitter": map[string]interface{}{
																	"description": "The time within which each health check will be randomly delayed. By default, there is no delay.",
																	"type":        "string",
																},
																"keepalive-time": map[string]interface{}{
																	"description": "Enables keepalive connections for health checks and specifies the time during which requests can be processed through one keepalive connection. The default is 60s.",
																	"type":        "string",
																},
																"mandatory": map[string]interface{}{
																	"description": "Require every newly added server to pass all configured health checks before NGINX Plus sends traffic to it. If this is not specified, or is set to false, the server will be initially considered healthy. When combined with slow-start, it gives a new server more time to connect to databases and “warm up” before being asked to handle their full share of traffic.",
																	"type":        "boolean",
																},
																"passes": map[string]interface{}{
																	"description": "The number of consecutive passed health checks of a particular upstream server after which the server will be considered healthy. The default is 1.",
																	"type":        "integer",
																},
																"path": map[string]interface{}{
																	"description": "The path used for health check requests. The default is /. This is not configurable for gRPC type upstreams.",
																	"type":        "string",
																},
																"persistent": map[string]interface{}{
																	"description": "Set the initial “up” state for a server after reload if the server was considered healthy before reload. Enabling persistent requires that the mandatory parameter is also set to true.",
																	"type":        "boolean",
																},
																"port": map[string]interface{}{
																	"description": "The port used for health check requests. By default, the server port is used. Note: in contrast with the port of the upstream, this port is not a service port, but a port of a pod.",
																	"type":        "integer",
																},
																"read-timeout": map[string]interface{}{
																	"description": "The timeout for reading a response from an upstream server. By default, the read-timeout of the upstream is used.",
																	"type":        "string",
																},
																"send-timeout": map[string]interface{}{
																	"description": "The timeout for transmitting a request to an upstream server. By default, the send-timeout of the upstream is used.",
																	"type":        "string",
																},
																"statusMatch": map[string]interface{}{
																	"description": "The expected response status codes of a health check. By default, the response should have status code 2xx or 3xx. Examples: \"200\", \"! 500\", \"301-303 307\". This not supported for gRPC type upstreams.",
																	"type":        "string",
																},
																"tls": map[string]interface{}{
																	"description": "The TLS configuration used for health check requests. By default, the tls field of the upstream is used.",
																	"properties": map[string]interface{}{
																		"enable": map[string]interface{}{
																			"description": "Enables HTTPS for requests to upstream servers. The default is False , meaning that HTTP will be used. Note: by default, NGINX will not verify the upstream server certificate. To enable the verification, configure an EgressMTLS Policy.",
																			"type":        "boolean",
																		},
																	},
																	"type": "object",
																},
															},
															"type": "object",
														},
														"keepalive": map[string]interface{}{
															"description": "Configures the cache for connections to upstream servers. The value 0 disables the cache. The default is set in the keepalive ConfigMap key.",
															"type":        "integer",
														},
														"lb-method": map[string]interface{}{
															"description": "The load balancing method. To use the round-robin method, specify round_robin. The default is specified in the lb-method ConfigMap key.",
															"type":        "string",
														},
														"max-conns": map[string]interface{}{
															"description": "The maximum number of simultaneous active connections to an upstream server. By default there is no limit. Note: if keepalive connections are enabled, the total number of active and idle keepalive connections to an upstream server may exceed the max_conns value.",
															"type":        "integer",
														},
														"max-fails": map[string]interface{}{
															"description": "The number of unsuccessful attempts to communicate with an upstream server that should happen in the duration set by the fail-timeout to consider the server unavailable. The default is set in the max-fails ConfigMap key.",
															"type":        "integer",
														},
														"name": map[string]interface{}{
															"description": "The name of the upstream. Must be a valid DNS label as defined in RFC 1035. For example, hello and upstream-123 are valid. The name must be unique among all upstreams of the resource.",
															"type":        "string",
														},
														"next-upstream": map[string]interface{}{
															"description": "Specifies in which cases a request should be passed to the next upstream server. The default is error timeout.",
															"type":        "string",
														},
														"next-upstream-timeout": map[string]interface{}{
															"description": "The time during which a request can be passed to the next upstream server. The 0 value turns off the time limit. The default is 0.",
															"type":        "string",
														},
														"next-upstream-tries": map[string]interface{}{
															"description": "The number of possible tries for passing a request to the next upstream server. The 0 value turns off this limit. The default is 0.",
															"type":        "integer",
														},
														"ntlm": map[string]interface{}{
															"description": "Allows proxying requests with NTLM Authentication. In order for NTLM authentication to work, it is necessary to enable keepalive connections to upstream servers using the keepalive field. Note: this feature is supported only in NGINX Plus.",
															"type":        "boolean",
														},
														"port": map[string]interface{}{
															"description": "The port of the service. If the service doesn’t define that port, NGINX will assume the service has zero endpoints and return a 502 response for requests for this upstream. The port must fall into the range 1..65535.",
															"type":        "integer",
														},
														"queue": map[string]interface{}{
															"description": "Configures a queue for an upstream. A client request will be placed into the queue if an upstream server cannot be selected immediately while processing the request. By default, no queue is configured. Note: this feature is supported only in NGINX Plus.",
															"properties": map[string]interface{}{
																"size": map[string]interface{}{
																	"description": "The size of the queue.",
																	"type":        "integer",
																},
																"timeout": map[string]interface{}{
																	"description": "The timeout of the queue. A request cannot be queued for a period longer than the timeout. The default is 60s.",
																	"type":        "string",
																},
															},
															"type": "object",
														},
														"read-timeout": map[string]interface{}{
															"description": "The timeout for reading a response from an upstream server. The default is specified in the proxy-read-timeout ConfigMap key.",
															"type":        "string",
														},
														"send-timeout": map[string]interface{}{
															"description": "The timeout for transmitting a request to an upstream server. The default is specified in the proxy-send-timeout ConfigMap key.",
															"type":        "string",
														},
														"service": map[string]interface{}{
															"description": "The name of a service. If the Service belongs to a different namespace than the VirtualServer or VirtualServerRoute, you need to include the namespace. For example, tea-namespace/tea. If the service doesn’t exist, NGINX will assume the service has zero endpoints and return a 502 response for requests for this upstream. For NGINX Plus only, services of type ExternalName are also supported in the same namespace.",
															"type":        "string",
														},
														"sessionCookie": map[string]interface{}{
															"description": "The SessionCookie field configures session persistence which allows requests from the same client to be passed to the same upstream server. The information about the designated upstream server is passed in a session cookie generated by NGINX.",
															"properties": map[string]interface{}{
																"domain": map[string]interface{}{
																	"description": "The domain for which the cookie is set.",
																	"type":        "string",
																},
																"enable": map[string]interface{}{
																	"description": "Enables session persistence with a session cookie for an upstream server. The default is false.",
																	"type":        "boolean",
																},
																"expires": map[string]interface{}{
																	"description": "The time for which a browser should keep the cookie. Can be set to the special value max, which will cause the cookie to expire on 31 Dec 2037 23:55:55 GMT.",
																	"type":        "string",
																},
																"httpOnly": map[string]interface{}{
																	"description": "Adds the HttpOnly attribute to the cookie.",
																	"type":        "boolean",
																},
																"name": map[string]interface{}{
																	"description": "The name of the cookie.",
																	"type":        "string",
																},
																"path": map[string]interface{}{
																	"description": "The path for which the cookie is set.",
																	"type":        "string",
																},
																"samesite": map[string]interface{}{
																	"description": "Adds the SameSite attribute to the cookie. The allowed values are: strict, lax, none",
																	"type":        "string",
																},
																"secure": map[string]interface{}{
																	"description": "Adds the Secure attribute to the cookie.",
																	"type":        "boolean",
																},
															},
															"type": "object",
														},
														"slow-start": map[string]interface{}{
															"description": "The slow start allows an upstream server to gradually recover its weight from 0 to its nominal value after it has been recovered or became available or when the server becomes available after a period of time it was considered unavailable. By default, the slow start is disabled. Note: The parameter cannot be used along with the random, hash or ip_hash load balancing methods and will be ignored.",
															"type":        "string",
														},
														"subselector": map[string]interface{}{
															"additionalProperties": map[string]interface{}{
																"type": "string",
															},
															"description": "Selects the pods within the service using label keys and values. By default, all pods of the service are selected. Note: the specified labels are expected to be present in the pods when they are created. If the pod labels are updated, NGINX Ingress Controller will not see that change until the number of the pods is changed.",
															"type":        "object",
														},
														"tls": map[string]interface{}{
															"description": "The TLS configuration for the Upstream.",
															"properties": map[string]interface{}{
																"enable": map[string]interface{}{
																	"description": "Enables HTTPS for requests to upstream servers. The default is False , meaning that HTTP will be used. Note: by default, NGINX will not verify the upstream server certificate. To enable the verification, configure an EgressMTLS Policy.",
																	"type":        "boolean",
																},
															},
															"type": "object",
														},
														"type": map[string]interface{}{
															"description": "The type of the upstream. Supported values are http and grpc. The default is http. For gRPC, it is necessary to enable HTTP/2 in the ConfigMap and configure TLS termination in the VirtualServer.",
															"type":        "string",
														},
														"use-cluster-ip": map[string]interface{}{
															"description": "Enables using the Cluster IP and port of the service instead of the default behavior of using the IP and port of the pods. When this field is enabled, the fields that configure NGINX behavior related to multiple upstream servers (like lb-method and next-upstream) will have no effect, as NGINX Ingress Controller will configure NGINX with only one upstream server that will match the service Cluster IP.",
															"type":        "boolean",
														},
													},
													"type": "object",
												},
												"type": "array",
											},
										},
										"type": "object",
									},
									"status": map[string]interface{}{
										"description": "VirtualServerRouteStatus defines the status for the VirtualServerRoute resource.",
										"properties": map[string]interface{}{
											"externalEndpoints": map[string]interface{}{
												"description": "Defines the IPs, hostnames and ports used to connect to this resource.",
												"items": map[string]interface{}{
													"description": "ExternalEndpoint defines the IP/ Hostname and ports used to connect to this resource.",
													"properties": map[string]interface{}{
														"hostname": map[string]interface{}{
															"type": "string",
														},
														"ip": map[string]interface{}{
															"type": "string",
														},
														"ports": map[string]interface{}{
															"type": "string",
														},
													},
													"type": "object",
												},
												"type": "array",
											},
											"message": map[string]interface{}{
												"description": "The message of the current state of the resource. It can contain more detailed information about the reason.",
												"type":        "string",
											},
											"reason": map[string]interface{}{
												"description": "The reason of the current state of the resource.",
												"type":        "string",
											},
											"referencedBy": map[string]interface{}{
												"description": "Defines how other resources reference this resource.",
												"type":        "string",
											},
											"state": map[string]interface{}{
												"description": "Represents the current state of the resource. There are three possible values: Valid, Invalid and Warning. Valid indicates that the resource has been validated and accepted by the Ingress Controller. Invalid means the resource failed validation or NGINX",
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
			},
		},
	}

	return mutate.MutateCRDVirtualserverroutesK8sNginxOrg(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;create;update;patch;delete

// CreateCRDApdospoliciesAppprotectdosF5Com creates the CustomResourceDefinition resource with name apdospolicies.appprotectdos.f5.com.
func CreateCRDApdospoliciesAppprotectdosF5Com(
	parent *platformv1alpha1.IngressComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if parent.Spec.Nginx.Include != true {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:field=nginx.include,value=true,include
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					"controller-gen.kubebuilder.io/version": "v0.9.2",
				},
				"creationTimestamp": nil,
				"name":              "apdospolicies.appprotectdos.f5.com",
				"labels": map[string]interface{}{
					"platform.nukleros.io/category": "ingress",
					"platform.nukleros.io/project":  "nginx-ingress-controller",
				},
			},
			"spec": map[string]interface{}{
				"group": "appprotectdos.f5.com",
				"names": map[string]interface{}{
					"kind":     "APDosPolicy",
					"listKind": "APDosPoliciesList",
					"plural":   "apdospolicies",
					"singular": "apdospolicy",
				},
				"scope": "Namespaced",
				"versions": []interface{}{
					map[string]interface{}{
						"name": "v1beta1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"type":        "object",
								"description": "APDosPolicy is the Schema for the APDosPolicy API",
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
										"type":        "object",
										"description": "APDosPolicySpec defines the desired state of APDosPolicy",
										"properties": map[string]interface{}{
											"mitigation_mode": map[string]interface{}{
												"enum": []interface{}{
													"standard",
													"conservative",
													"none",
												},
												"default": "standard",
												"type":    "string",
											},
											"signatures": map[string]interface{}{
												"enum": []interface{}{
													"on",
													"off",
												},
												"default": "on",
												"type":    "string",
											},
											"bad_actors": map[string]interface{}{
												"enum": []interface{}{
													"on",
													"off",
												},
												"default": "on",
												"type":    "string",
											},
											"automation_tools_detection": map[string]interface{}{
												"enum": []interface{}{
													"on",
													"off",
												},
												"default": "on",
												"type":    "string",
											},
											"tls_fingerprint": map[string]interface{}{
												"enum": []interface{}{
													"on",
													"off",
												},
												"default": "on",
												"type":    "string",
											},
										},
									},
								},
							},
						},
						"served":  true,
						"storage": true,
					},
				},
			},
		},
	}

	return mutate.MutateCRDApdospoliciesAppprotectdosF5Com(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;create;update;patch;delete

// CreateCRDDosprotectedresourcesAppprotectdosF5Com creates the CustomResourceDefinition resource with name dosprotectedresources.appprotectdos.f5.com.
func CreateCRDDosprotectedresourcesAppprotectdosF5Com(
	parent *platformv1alpha1.IngressComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if parent.Spec.Nginx.Include != true {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:field=nginx.include,value=true,include
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					"controller-gen.kubebuilder.io/version": "v0.21.0",
				},
				"name": "dosprotectedresources.appprotectdos.f5.com",
				"labels": map[string]interface{}{
					"platform.nukleros.io/category": "ingress",
					"platform.nukleros.io/project":  "nginx-ingress-controller",
				},
			},
			"spec": map[string]interface{}{
				"group": "appprotectdos.f5.com",
				"names": map[string]interface{}{
					"kind":     "DosProtectedResource",
					"listKind": "DosProtectedResourceList",
					"plural":   "dosprotectedresources",
					"shortNames": []interface{}{
						"pr",
					},
					"singular": "dosprotectedresource",
				},
				"scope": "Namespaced",
				"versions": []interface{}{
					map[string]interface{}{
						"name": "v1beta1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"description": "DosProtectedResource defines a Dos protected resource.",
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
										"description": "DosProtectedResourceSpec defines the properties and values a DosProtectedResource can have.",
										"properties": map[string]interface{}{
											"allowList": map[string]interface{}{
												"description": "AllowList is a list of allowed IPs and subnet masks",
												"items": map[string]interface{}{
													"description": "AllowListEntry represents an IP address and a subnet mask.",
													"properties": map[string]interface{}{
														"ipWithMask": map[string]interface{}{
															"type": "string",
														},
													},
													"type": "object",
												},
												"type": "array",
											},
											"apDosMonitor": map[string]interface{}{
												"description": "ApDosMonitor is how NGINX App Protect DoS monitors the stress level of the protected object. The monitor requests are sent from localhost (127.0.0.1). Default value: URI - None, protocol - http1, timeout - NGINX App Protect DoS default.",
												"properties": map[string]interface{}{
													"protocol": map[string]interface{}{
														"description": "Protocol determines if the server listens on http1 / http2 / grpc / websocket. The default is http1.",
														"enum": []interface{}{
															"http1",
															"http2",
															"grpc",
															"websocket",
														},
														"type": "string",
													},
													"timeout": map[string]interface{}{
														"description": "Timeout determines how long (in seconds) should NGINX App Protect DoS wait for a response. Default is 10 seconds for http1/http2 and 5 seconds for grpc.",
														"format":      "int64",
														"type":        "integer",
													},
													"uri": map[string]interface{}{
														"description": "URI is the destination to the desired protected object in the nginx.conf:",
														"type":        "string",
													},
												},
												"type": "object",
											},
											"apDosPolicy": map[string]interface{}{
												"description": "ApDosPolicy is the namespace/name of a ApDosPolicy resource",
												"type":        "string",
											},
											"dosAccessLogDest": map[string]interface{}{
												"description": "DosAccessLogDest is the network address for the access logs",
												"type":        "string",
											},
											"dosSecurityLog": map[string]interface{}{
												"description": "DosSecurityLog defines the security log of the DosProtectedResource.",
												"properties": map[string]interface{}{
													"apDosLogConf": map[string]interface{}{
														"description": "ApDosLogConf is the namespace/name of a APDosLogConf resource",
														"type":        "string",
													},
													"dosLogDest": map[string]interface{}{
														"description": "DosLogDest is the network address of a logging service, can be either IP or DNS name.",
														"type":        "string",
													},
													"enable": map[string]interface{}{
														"description": "Enable enables the security logging feature if set to true",
														"type":        "boolean",
													},
												},
												"type": "object",
											},
											"enable": map[string]interface{}{
												"description": "Enable enables the DOS feature if set to true",
												"type":        "boolean",
											},
											"name": map[string]interface{}{
												"description": "Name is the name of protected object, max of 63 characters.",
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
					},
				},
			},
		},
	}

	return mutate.MutateCRDDosprotectedresourcesAppprotectdosF5Com(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;create;update;patch;delete

// CreateCRDGlobalconfigurationsK8sNginxOrg creates the CustomResourceDefinition resource with name globalconfigurations.k8s.nginx.org.
func CreateCRDGlobalconfigurationsK8sNginxOrg(
	parent *platformv1alpha1.IngressComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if parent.Spec.Nginx.Include != true {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:field=nginx.include,value=true,include
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					"controller-gen.kubebuilder.io/version": "v0.21.0",
				},
				"name": "globalconfigurations.k8s.nginx.org",
				"labels": map[string]interface{}{
					"platform.nukleros.io/category": "ingress",
					"platform.nukleros.io/project":  "nginx-ingress-controller",
				},
			},
			"spec": map[string]interface{}{
				"group": "k8s.nginx.org",
				"names": map[string]interface{}{
					"kind":     "GlobalConfiguration",
					"listKind": "GlobalConfigurationList",
					"plural":   "globalconfigurations",
					"shortNames": []interface{}{
						"gc",
					},
					"singular": "globalconfiguration",
				},
				"scope": "Namespaced",
				"versions": []interface{}{
					map[string]interface{}{
						"name": "v1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"description": "GlobalConfiguration defines the GlobalConfiguration resource.",
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
										"description": "GlobalConfigurationSpec resource defines the global configuration parameters of the Ingress Controller.",
										"properties": map[string]interface{}{
											"listeners": map[string]interface{}{
												"description": "Listeners field of the GlobalConfigurationSpec resource",
												"items": map[string]interface{}{
													"description": "Listener defines a listener.",
													"properties": map[string]interface{}{
														"ipv4": map[string]interface{}{
															"description": "Specifies the IPv4 address to listen on.",
															"type":        "string",
														},
														"ipv6": map[string]interface{}{
															"description": "ipv6 addresse that NGINX will listen on.",
															"type":        "string",
														},
														"name": map[string]interface{}{
															"description": "The name of the listener. The name must be unique across all listeners.",
															"type":        "string",
														},
														"port": map[string]interface{}{
															"description": "The port on which the listener will accept connections.",
															"type":        "integer",
														},
														"protocol": map[string]interface{}{
															"description": "The protocol of the listener. For example, HTTP.",
															"type":        "string",
														},
														"ssl": map[string]interface{}{
															"description": "Whether the listener will be listening for SSL connections",
															"type":        "boolean",
														},
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
					},
				},
			},
		},
	}

	return mutate.MutateCRDGlobalconfigurationsK8sNginxOrg(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;create;update;patch;delete

// CreateCRDAppoliciesAppprotectF5Com creates the CustomResourceDefinition resource with name appolicies.appprotect.f5.com.
func CreateCRDAppoliciesAppprotectF5Com(
	parent *platformv1alpha1.IngressComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if parent.Spec.Nginx.Include != true {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:field=nginx.include,value=true,include
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					"controller-gen.kubebuilder.io/version": "v0.13.0",
				},
				"name": "appolicies.appprotect.f5.com",
				"labels": map[string]interface{}{
					"platform.nukleros.io/category": "ingress",
					"platform.nukleros.io/project":  "nginx-ingress-controller",
				},
			},
			"spec": map[string]interface{}{
				"group": "appprotect.f5.com",
				"names": map[string]interface{}{
					"kind":     "APPolicy",
					"listKind": "APPolicyList",
					"plural":   "appolicies",
					"singular": "appolicy",
				},
				"scope": "Namespaced",
				"versions": []interface{}{
					map[string]interface{}{
						"name": "v1beta1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"description": "APPolicyConfig is the Schema for the APPolicyconfigs API",
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
										"description": "APPolicySpec defines the desired state of APPolicy",
										"properties": map[string]interface{}{
											"modifications": map[string]interface{}{
												"items": map[string]interface{}{
													"properties": map[string]interface{}{
														"action": map[string]interface{}{
															"type": "string",
														},
														"description": map[string]interface{}{
															"type": "string",
														},
														"entity": map[string]interface{}{
															"properties": map[string]interface{}{
																"name": map[string]interface{}{
																	"type": "string",
																},
															},
															"type": "object",
														},
														"entityChanges": map[string]interface{}{
															"properties": map[string]interface{}{
																"type": map[string]interface{}{
																	"type": "string",
																},
															},
															"type": "object",
														},
													},
													"type":                                 "object",
													"x-kubernetes-preserve-unknown-fields": true,
												},
												"type": "array",
											},
											"modificationsReference": map[string]interface{}{
												"properties": map[string]interface{}{
													"link": map[string]interface{}{
														"pattern": "^http",
														"type":    "string",
													},
												},
												"type": "object",
											},
											"policy": map[string]interface{}{
												"description": "Defines the App Protect policy",
												"properties": map[string]interface{}{
													"applicationLanguage": map[string]interface{}{
														"enum": []interface{}{
															"iso-8859-10",
															"iso-8859-6",
															"windows-1255",
															"auto-detect",
															"koi8-r",
															"gb18030",
															"iso-8859-8",
															"windows-1250",
															"iso-8859-9",
															"windows-1252",
															"iso-8859-16",
															"gb2312",
															"iso-8859-2",
															"iso-8859-5",
															"windows-1257",
															"windows-1256",
															"iso-8859-13",
															"windows-874",
															"windows-1253",
															"iso-8859-3",
															"euc-jp",
															"utf-8",
															"gbk",
															"windows-1251",
															"big5",
															"iso-8859-1",
															"shift_jis",
															"euc-kr",
															"iso-8859-4",
															"iso-8859-7",
															"iso-8859-15",
														},
														"type": "string",
													},
													"blocking-settings": map[string]interface{}{
														"properties": map[string]interface{}{
															"evasions": map[string]interface{}{
																"items": map[string]interface{}{
																	"properties": map[string]interface{}{
																		"description": map[string]interface{}{
																			"enum": []interface{}{
																				"%u decoding",
																				"Apache whitespace",
																				"Bad unescape",
																				"Bare byte decoding",
																				"Directory traversals",
																				"IIS backslashes",
																				"IIS Unicode codepoints",
																				"Multiple decoding",
																				"Multiple slashes",
																				"Semicolon path parameters",
																				"Trailing dot",
																				"Trailing slash",
																			},
																			"type": "string",
																		},
																		"enabled": map[string]interface{}{
																			"type": "boolean",
																		},
																		"maxDecodingPasses": map[string]interface{}{
																			"type": "integer",
																		},
																	},
																	"type": "object",
																},
																"type": "array",
															},
															"http-protocols": map[string]interface{}{
																"items": map[string]interface{}{
																	"properties": map[string]interface{}{
																		"description": map[string]interface{}{
																			"enum": []interface{}{
																				"Unescaped space in URL",
																				"Unparsable request content",
																				"Several Content-Length headers",
																				"POST request with Content-Length: 0",
																				"Null in request",
																				"No Host header in HTTP/1.1 request",
																				"Multiple host headers",
																				"Host header contains IP address",
																				"High ASCII characters in headers",
																				"Header name with no header value",
																				"CRLF characters before request start",
																				"Content length should be a positive number",
																				"Chunked request with Content-Length header",
																				"Check maximum number of cookies",
																				"Check maximum number of parameters",
																				"Check maximum number of headers",
																				"Body in GET or HEAD requests",
																				"Bad multipart/form-data request parsing",
																				"Bad multipart parameters parsing",
																				"Bad HTTP version",
																				"Bad host header value",
																			},
																			"type": "string",
																		},
																		"enabled": map[string]interface{}{
																			"type": "boolean",
																		},
																		"maxCookies": map[string]interface{}{
																			"maximum": 100,
																			"minimum": 1,
																			"type":    "integer",
																		},
																		"maxHeaders": map[string]interface{}{
																			"maximum": 150,
																			"minimum": 1,
																			"type":    "integer",
																		},
																		"maxParams": map[string]interface{}{
																			"maximum": 5000,
																			"minimum": 1,
																			"type":    "integer",
																		},
																	},
																	"type": "object",
																},
																"type": "array",
															},
															"violations": map[string]interface{}{
																"items": map[string]interface{}{
																	"properties": map[string]interface{}{
																		"alarm": map[string]interface{}{
																			"type": "boolean",
																		},
																		"block": map[string]interface{}{
																			"type": "boolean",
																		},
																		"description": map[string]interface{}{
																			"type": "string",
																		},
																		"name": map[string]interface{}{
																			"enum": []interface{}{
																				"VIOL_ACCESS_INVALID",
																				"VIOL_ACCESS_MALFORMED",
																				"VIOL_ACCESS_MISSING",
																				"VIOL_ACCESS_UNAUTHORIZED",
																				"VIOL_ASM_COOKIE_HIJACKING",
																				"VIOL_ASM_COOKIE_MODIFIED",
																				"VIOL_BLACKLISTED_IP",
																				"VIOL_BOT_CLIENT",
																				"VIOL_BRUTE_FORCE",
																				"VIOL_COOKIE_EXPIRED",
																				"VIOL_COOKIE_LENGTH",
																				"VIOL_COOKIE_MALFORMED",
																				"VIOL_COOKIE_MODIFIED",
																				"VIOL_CSRF",
																				"VIOL_DATA_GUARD",
																				"VIOL_ENCODING",
																				"VIOL_EVASION",
																				"VIOL_FILE_UPLOAD",
																				"VIOL_FILE_UPLOAD_IN_BODY",
																				"VIOL_FILETYPE",
																				"VIOL_GEOLOCATION",
																				"VIOL_GRAPHQL_ERROR_RESPONSE",
																				"VIOL_GRAPHQL_FORMAT",
																				"VIOL_GRAPHQL_INTROSPECTION_QUERY",
																				"VIOL_GRAPHQL_MALFORMED",
																				"VIOL_GRPC_FORMAT",
																				"VIOL_GRPC_MALFORMED",
																				"VIOL_GRPC_METHOD",
																				"VIOL_HEADER_LENGTH",
																				"VIOL_HEADER_METACHAR",
																				"VIOL_HEADER_REPEATED",
																				"VIOL_HTTP_PROTOCOL",
																				"VIOL_HTTP_RESPONSE_STATUS",
																				"VIOL_JSON_FORMAT",
																				"VIOL_JSON_MALFORMED",
																				"VIOL_JSON_SCHEMA",
																				"VIOL_LOGIN",
																				"VIOL_LOGIN_URL_BYPASSED",
																				"VIOL_LOGIN_URL_EXPIRED",
																				"VIOL_MANDATORY_HEADER",
																				"VIOL_MANDATORY_PARAMETER",
																				"VIOL_MANDATORY_REQUEST_BODY",
																				"VIOL_MALICIOUS_IP",
																				"VIOL_METHOD",
																				"VIOL_PARAMETER",
																				"VIOL_PARAMETER_ARRAY_VALUE",
																				"VIOL_PARAMETER_DATA_TYPE",
																				"VIOL_PARAMETER_EMPTY_VALUE",
																				"VIOL_PARAMETER_LOCATION",
																				"VIOL_PARAMETER_MULTIPART_NULL_VALUE",
																				"VIOL_PARAMETER_NAME_METACHAR",
																				"VIOL_PARAMETER_NUMERIC_VALUE",
																				"VIOL_PARAMETER_REPEATED",
																				"VIOL_PARAMETER_STATIC_VALUE",
																				"VIOL_PARAMETER_VALUE_BASE64",
																				"VIOL_PARAMETER_VALUE_LENGTH",
																				"VIOL_PARAMETER_VALUE_METACHAR",
																				"VIOL_PARAMETER_VALUE_REGEXP",
																				"VIOL_POST_DATA_LENGTH",
																				"VIOL_QUERY_STRING_LENGTH",
																				"VIOL_RATING_NEED_EXAMINATION",
																				"VIOL_RATING_THREAT",
																				"VIOL_REQUEST_LENGTH",
																				"VIOL_REQUEST_MAX_LENGTH",
																				"VIOL_THREAT_CAMPAIGN",
																				"VIOL_URL",
																				"VIOL_URL_CONTENT_TYPE",
																				"VIOL_URL_LENGTH",
																				"VIOL_URL_METACHAR",
																				"VIOL_WEBSOCKET_BAD_REQUEST",
																				"VIOL_XML_FORMAT",
																				"VIOL_XML_MALFORMED",
																			},
																			"type": "string",
																		},
																	},
																	"type": "object",
																},
																"type": "array",
															},
														},
														"type": "object",
													},
													"blockingSettingReference": map[string]interface{}{
														"properties": map[string]interface{}{
															"link": map[string]interface{}{
																"pattern": "^http",
																"type":    "string",
															},
														},
														"type": "object",
													},
													"bot-defense": map[string]interface{}{
														"properties": map[string]interface{}{
															"mitigations": map[string]interface{}{
																"properties": map[string]interface{}{
																	"anomalies": map[string]interface{}{
																		"items": map[string]interface{}{
																			"properties": map[string]interface{}{
																				"$action": map[string]interface{}{
																					"enum": []interface{}{
																						"delete",
																					},
																					"type": "string",
																				},
																				"action": map[string]interface{}{
																					"enum": []interface{}{
																						"alarm",
																						"block",
																						"default",
																						"detect",
																						"ignore",
																					},
																					"type": "string",
																				},
																				"name": map[string]interface{}{
																					"type": "string",
																				},
																				"scoreThreshold": map[string]interface{}{
																					"anyOf": []interface{}{
																						map[string]interface{}{
																							"type": "integer",
																						},
																						map[string]interface{}{
																							"type": "string",
																						},
																					},
																					"x-kubernetes-int-or-string": true,
																				},
																			},
																			"type": "object",
																		},
																		"type": "array",
																	},
																	"browsers": map[string]interface{}{
																		"items": map[string]interface{}{
																			"properties": map[string]interface{}{
																				"$action": map[string]interface{}{
																					"enum": []interface{}{
																						"delete",
																					},
																					"type": "string",
																				},
																				"action": map[string]interface{}{
																					"enum": []interface{}{
																						"alarm",
																						"block",
																						"detect",
																					},
																					"type": "string",
																				},
																				"maxVersion": map[string]interface{}{
																					"maximum": 2147483647,
																					"minimum": 0,
																					"type":    "integer",
																				},
																				"minVersion": map[string]interface{}{
																					"maximum": 2147483647,
																					"minimum": 0,
																					"type":    "integer",
																				},
																				"name": map[string]interface{}{
																					"type": "string",
																				},
																			},
																			"type": "object",
																		},
																		"type": "array",
																	},
																	"classes": map[string]interface{}{
																		"items": map[string]interface{}{
																			"properties": map[string]interface{}{
																				"action": map[string]interface{}{
																					"enum": []interface{}{
																						"alarm",
																						"block",
																						"detect",
																						"ignore",
																					},
																					"type": "string",
																				},
																				"name": map[string]interface{}{
																					"enum": []interface{}{
																						"browser",
																						"malicious-bot",
																						"suspicious-browser",
																						"trusted-bot",
																						"unknown",
																						"untrusted-bot",
																					},
																					"type": "string",
																				},
																			},
																			"type": "object",
																		},
																		"type": "array",
																	},
																	"signatures": map[string]interface{}{
																		"items": map[string]interface{}{
																			"properties": map[string]interface{}{
																				"$action": map[string]interface{}{
																					"enum": []interface{}{
																						"delete",
																					},
																					"type": "string",
																				},
																				"action": map[string]interface{}{
																					"enum": []interface{}{
																						"alarm",
																						"block",
																						"detect",
																						"ignore",
																					},
																					"type": "string",
																				},
																				"name": map[string]interface{}{
																					"type": "string",
																				},
																			},
																			"type": "object",
																		},
																		"type": "array",
																	},
																},
																"type": "object",
															},
															"settings": map[string]interface{}{
																"properties": map[string]interface{}{
																	"caseSensitiveHttpHeaders": map[string]interface{}{
																		"type": "boolean",
																	},
																	"isEnabled": map[string]interface{}{
																		"type": "boolean",
																	},
																},
																"type": "object",
															},
														},
														"type": "object",
													},
													"browser-definitions": map[string]interface{}{
														"items": map[string]interface{}{
															"properties": map[string]interface{}{
																"$action": map[string]interface{}{
																	"enum": []interface{}{
																		"delete",
																	},
																	"type": "string",
																},
																"isUserDefined": map[string]interface{}{
																	"type": "boolean",
																},
																"matchRegex": map[string]interface{}{
																	"type": "string",
																},
																"matchString": map[string]interface{}{
																	"type": "string",
																},
																"name": map[string]interface{}{
																	"type": "string",
																},
															},
															"type": "object",
														},
														"type": "array",
													},
													"caseInsensitive": map[string]interface{}{
														"type": "boolean",
													},
													"character-sets": map[string]interface{}{
														"items": map[string]interface{}{
															"properties": map[string]interface{}{
																"characterSet": map[string]interface{}{
																	"items": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"isAllowed": map[string]interface{}{
																				"type": "boolean",
																			},
																			"metachar": map[string]interface{}{
																				"type": "string",
																			},
																		},
																		"type": "object",
																	},
																	"type": "array",
																},
																"characterSetType": map[string]interface{}{
																	"enum": []interface{}{
																		"gwt-content",
																		"header",
																		"json-content",
																		"parameter-name",
																		"parameter-value",
																		"plain-text-content",
																		"url",
																		"xml-content",
																	},
																	"type": "string",
																},
															},
															"type": "object",
														},
														"type": "array",
													},
													"characterSetReference": map[string]interface{}{
														"properties": map[string]interface{}{
															"link": map[string]interface{}{
																"pattern": "^http",
																"type":    "string",
															},
														},
														"type": "object",
													},
													"cookie-settings": map[string]interface{}{
														"properties": map[string]interface{}{
															"maximumCookieHeaderLength": map[string]interface{}{
																"anyOf": []interface{}{
																	map[string]interface{}{
																		"type": "integer",
																	},
																	map[string]interface{}{
																		"type": "string",
																	},
																},
																"x-kubernetes-int-or-string": true,
															},
														},
														"type": "object",
													},
													"cookieReference": map[string]interface{}{
														"properties": map[string]interface{}{
															"link": map[string]interface{}{
																"pattern": "^http",
																"type":    "string",
															},
														},
														"type": "object",
													},
													"cookieSettingsReference": map[string]interface{}{
														"properties": map[string]interface{}{
															"link": map[string]interface{}{
																"pattern": "^http",
																"type":    "string",
															},
														},
														"type": "object",
													},
													"cookies": map[string]interface{}{
														"items": map[string]interface{}{
															"properties": map[string]interface{}{
																"$action": map[string]interface{}{
																	"enum": []interface{}{
																		"delete",
																	},
																	"type": "string",
																},
																"accessibleOnlyThroughTheHttpProtocol": map[string]interface{}{
																	"type": "boolean",
																},
																"attackSignaturesCheck": map[string]interface{}{
																	"type": "boolean",
																},
																"decodeValueAsBase64": map[string]interface{}{
																	"enum": []interface{}{
																		"enabled",
																		"disabled",
																		"required",
																	},
																	"type": "string",
																},
																"enforcementType": map[string]interface{}{
																	"type": "string",
																},
																"insertSameSiteAttribute": map[string]interface{}{
																	"enum": []interface{}{
																		"lax",
																		"none",
																		"none-value",
																		"strict",
																	},
																	"type": "string",
																},
																"maskValueInLogs": map[string]interface{}{
																	"type": "boolean",
																},
																"name": map[string]interface{}{
																	"type": "string",
																},
																"securedOverHttpsConnection": map[string]interface{}{
																	"type": "boolean",
																},
																"signatureOverrides": map[string]interface{}{
																	"items": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"enabled": map[string]interface{}{
																				"type": "boolean",
																			},
																			"name": map[string]interface{}{
																				"type": "string",
																			},
																			"signatureId": map[string]interface{}{
																				"type": "integer",
																			},
																			"tag": map[string]interface{}{
																				"type": "string",
																			},
																		},
																		"type": "object",
																	},
																	"type": "array",
																},
																"type": map[string]interface{}{
																	"enum": []interface{}{
																		"explicit",
																		"wildcard",
																	},
																	"type": "string",
																},
																"wildcardOrder": map[string]interface{}{
																	"type": "integer",
																},
															},
															"type": "object",
														},
														"type": "array",
													},
													"csrf-protection": map[string]interface{}{
														"properties": map[string]interface{}{
															"enabled": map[string]interface{}{
																"type": "boolean",
															},
															"expirationTimeInSeconds": map[string]interface{}{
																"pattern": `disabled|\d+`,
																"type":    "string",
															},
															"sslOnly": map[string]interface{}{
																"type": "boolean",
															},
														},
														"type": "object",
													},
													"csrf-urls": map[string]interface{}{
														"items": map[string]interface{}{
															"properties": map[string]interface{}{
																"$action": map[string]interface{}{
																	"enum": []interface{}{
																		"delete",
																	},
																	"type": "string",
																},
																"enforcementAction": map[string]interface{}{
																	"enum": []interface{}{
																		"verify-origin",
																		"none",
																	},
																	"type": "string",
																},
																"method": map[string]interface{}{
																	"enum": []interface{}{
																		"GET",
																		"POST",
																		"any",
																	},
																	"type": "string",
																},
																"url": map[string]interface{}{
																	"type": "string",
																},
																"wildcardOrder": map[string]interface{}{
																	"type": "integer",
																},
															},
															"type": "object",
														},
														"type": "array",
													},
													"data-guard": map[string]interface{}{
														"properties": map[string]interface{}{
															"creditCardNumbers": map[string]interface{}{
																"type": "boolean",
															},
															"customPatterns": map[string]interface{}{
																"type": "boolean",
															},
															"customPatternsList": map[string]interface{}{
																"items": map[string]interface{}{
																	"type": "string",
																},
																"type": "array",
															},
															"enabled": map[string]interface{}{
																"type": "boolean",
															},
															"enforcementMode": map[string]interface{}{
																"enum": []interface{}{
																	"ignore-urls-in-list",
																	"enforce-urls-in-list",
																},
																"type": "string",
															},
															"enforcementUrls": map[string]interface{}{
																"items": map[string]interface{}{
																	"type": "string",
																},
																"type": "array",
															},
															"firstCustomCharactersToExpose": map[string]interface{}{
																"type": "integer",
															},
															"lastCcnDigitsToExpose": map[string]interface{}{
																"type": "integer",
															},
															"lastCustomCharactersToExpose": map[string]interface{}{
																"type": "integer",
															},
															"lastSsnDigitsToExpose": map[string]interface{}{
																"type": "integer",
															},
															"maskData": map[string]interface{}{
																"type": "boolean",
															},
															"usSocialSecurityNumbers": map[string]interface{}{
																"type": "boolean",
															},
														},
														"type": "object",
													},
													"dataGuardReference": map[string]interface{}{
														"properties": map[string]interface{}{
															"link": map[string]interface{}{
																"pattern": "^http",
																"type":    "string",
															},
														},
														"type": "object",
													},
													"description": map[string]interface{}{
														"type": "string",
													},
													"disallowed-geolocations": map[string]interface{}{
														"items": map[string]interface{}{
															"properties": map[string]interface{}{
																"$action": map[string]interface{}{
																	"enum": []interface{}{
																		"delete",
																	},
																	"type": "string",
																},
																"countryCode": map[string]interface{}{
																	"enum": []interface{}{
																		"AF",
																		"AX",
																		"AL",
																		"DZ",
																		"AS",
																		"AD",
																		"AO",
																		"AI",
																		"A1",
																		"AQ",
																		"AG",
																		"AR",
																		"AM",
																		"AW",
																		"AU",
																		"AT",
																		"AZ",
																		"BS",
																		"BH",
																		"BD",
																		"BB",
																		"BY",
																		"BE",
																		"BZ",
																		"BJ",
																		"BM",
																		"BT",
																		"BO",
																		"BA",
																		"BW",
																		"BV",
																		"BR",
																		"IO",
																		"BN",
																		"BG",
																		"BF",
																		"BI",
																		"KH",
																		"CM",
																		"CA",
																		"CV",
																		"KY",
																		"CF",
																		"TD",
																		"CL",
																		"CN",
																		"CX",
																		"CC",
																		"CO",
																		"KM",
																		"CG",
																		"CD",
																		"CK",
																		"CR",
																		"CI",
																		"HR",
																		"CU",
																		"CY",
																		"CZ",
																		"DK",
																		"DJ",
																		"DM",
																		"DO",
																		"EC",
																		"EG",
																		"SV",
																		"GQ",
																		"ER",
																		"EE",
																		"ET",
																		"FK",
																		"FO",
																		"FJ",
																		"FI",
																		"FR",
																		"FX",
																		"GF",
																		"PF",
																		"TF",
																		"GA",
																		"GM",
																		"GE",
																		"DE",
																		"GH",
																		"GI",
																		"GR",
																		"GL",
																		"GD",
																		"GP",
																		"GU",
																		"GT",
																		"GG",
																		"GN",
																		"GW",
																		"GY",
																		"HT",
																		"HM",
																		"VA",
																		"HN",
																		"HK",
																		"HU",
																		"IS",
																		"IN",
																		"ID",
																		"IR",
																		"IQ",
																		"IE",
																		"IM",
																		"IL",
																		"IT",
																		"JM",
																		"JP",
																		"JE",
																		"JO",
																		"KZ",
																		"KE",
																		"KI",
																		"KP",
																		"KR",
																		"KW",
																		"KG",
																		"LA",
																		"LV",
																		"LB",
																		"LS",
																		"LR",
																		"LY",
																		"LI",
																		"LT",
																		"LU",
																		"MO",
																		"MK",
																		"MG",
																		"MW",
																		"MY",
																		"MV",
																		"ML",
																		"MT",
																		"MH",
																		"MQ",
																		"MR",
																		"MU",
																		"YT",
																		"MX",
																		"FM",
																		"MD",
																		"MC",
																		"MN",
																		"ME",
																		"MS",
																		"MA",
																		"MZ",
																		"MM",
																		"ZZ",
																		"NA",
																		"NR",
																		"NP",
																		"NL",
																		"AN",
																		"NC",
																		"NZ",
																		"NI",
																		"NE",
																		"NG",
																		"NU",
																		"NF",
																		"MP",
																		"NO",
																		"OM",
																		"PK",
																		"PW",
																		"PS",
																		"PA",
																		"PG",
																		"PY",
																		"PE",
																		"PH",
																		"PN",
																		"PL",
																		"PT",
																		"PR",
																		"QA",
																		"RE",
																		"RO",
																		"RU",
																		"RW",
																		"BL",
																		"SH",
																		"KN",
																		"LC",
																		"MF",
																		"PM",
																		"VC",
																		"WS",
																		"SM",
																		"ST",
																		"A2",
																		"SA",
																		"SN",
																		"RS",
																		"SC",
																		"SL",
																		"SG",
																		"SK",
																		"SI",
																		"SB",
																		"SO",
																		"ZA",
																		"GS",
																		"ES",
																		"LK",
																		"SD",
																		"SR",
																		"SJ",
																		"SZ",
																		"SE",
																		"CH",
																		"SY",
																		"TW",
																		"TJ",
																		"TZ",
																		"TH",
																		"TL",
																		"TG",
																		"TK",
																		"TO",
																		"TT",
																		"TN",
																		"TR",
																		"TM",
																		"TC",
																		"TV",
																		"UG",
																		"UA",
																		"AE",
																		"GB",
																		"US",
																		"UM",
																		"UY",
																		"UZ",
																		"VU",
																		"VE",
																		"VN",
																		"VG",
																		"VI",
																		"WF",
																		"EH",
																		"YE",
																		"ZM",
																		"ZW",
																	},
																	"type": "string",
																},
																"countryName": map[string]interface{}{
																	"enum": []interface{}{
																		"Afghanistan",
																		"Aland Islands",
																		"Albania",
																		"Algeria",
																		"American Samoa",
																		"Andorra",
																		"Angola",
																		"Anguilla",
																		"Anonymous Proxy",
																		"Antarctica",
																		"Antigua and Barbuda",
																		"Argentina",
																		"Armenia",
																		"Aruba",
																		"Australia",
																		"Austria",
																		"Azerbaijan",
																		"Bahamas",
																		"Bahrain",
																		"Bangladesh",
																		"Barbados",
																		"Belarus",
																		"Belgium",
																		"Belize",
																		"Benin",
																		"Bermuda",
																		"Bhutan",
																		"Bolivia",
																		"Bosnia and Herzegovina",
																		"Botswana",
																		"Bouvet Island",
																		"Brazil",
																		"British Indian Ocean Territory",
																		"Brunei Darussalam",
																		"Bulgaria",
																		"Burkina Faso",
																		"Burundi",
																		"Cambodia",
																		"Cameroon",
																		"Canada",
																		"Cape Verde",
																		"Cayman Islands",
																		"Central African Republic",
																		"Chad",
																		"Chile",
																		"China",
																		"Christmas Island",
																		"Cocos (Keeling) Islands",
																		"Colombia",
																		"Comoros",
																		"Congo",
																		"Congo, The Democratic Republic of the",
																		"Cook Islands",
																		"Costa Rica",
																		"Cote D'Ivoire",
																		"Croatia",
																		"Cuba",
																		"Cyprus",
																		"Czech Republic",
																		"Denmark",
																		"Djibouti",
																		"Dominica",
																		"Dominican Republic",
																		"Ecuador",
																		"Egypt",
																		"El Salvador",
																		"Equatorial Guinea",
																		"Eritrea",
																		"Estonia",
																		"Ethiopia",
																		"Falkland Islands (Malvinas)",
																		"Faroe Islands",
																		"Fiji",
																		"Finland",
																		"France",
																		"France, Metropolitan",
																		"French Guiana",
																		"French Polynesia",
																		"French Southern Territories",
																		"Gabon",
																		"Gambia",
																		"Georgia",
																		"Germany",
																		"Ghana",
																		"Gibraltar",
																		"Greece",
																		"Greenland",
																		"Grenada",
																		"Guadeloupe",
																		"Guam",
																		"Guatemala",
																		"Guernsey",
																		"Guinea",
																		"Guinea-Bissau",
																		"Guyana",
																		"Haiti",
																		"Heard Island and McDonald Islands",
																		"Holy See (Vatican City State)",
																		"Honduras",
																		"Hong Kong",
																		"Hungary",
																		"Iceland",
																		"India",
																		"Indonesia",
																		"Iran, Islamic Republic of",
																		"Iraq",
																		"Ireland",
																		"Isle of Man",
																		"Israel",
																		"Italy",
																		"Jamaica",
																		"Japan",
																		"Jersey",
																		"Jordan",
																		"Kazakhstan",
																		"Kenya",
																		"Kiribati",
																		"Korea, Democratic People's Republic of",
																		"Korea, Republic of",
																		"Kuwait",
																		"Kyrgyzstan",
																		"Lao People's Democratic Republic",
																		"Latvia",
																		"Lebanon",
																		"Lesotho",
																		"Liberia",
																		"Libyan Arab Jamahiriya",
																		"Liechtenstein",
																		"Lithuania",
																		"Luxembourg",
																		"Macau",
																		"Macedonia",
																		"Madagascar",
																		"Malawi",
																		"Malaysia",
																		"Maldives",
																		"Mali",
																		"Malta",
																		"Marshall Islands",
																		"Martinique",
																		"Mauritania",
																		"Mauritius",
																		"Mayotte",
																		"Mexico",
																		"Micronesia, Federated States of",
																		"Moldova, Republic of",
																		"Monaco",
																		"Mongolia",
																		"Montenegro",
																		"Montserrat",
																		"Morocco",
																		"Mozambique",
																		"Myanmar",
																		"N/A",
																		"Namibia",
																		"Nauru",
																		"Nepal",
																		"Netherlands",
																		"Netherlands Antilles",
																		"New Caledonia",
																		"New Zealand",
																		"Nicaragua",
																		"Niger",
																		"Nigeria",
																		"Niue",
																		"Norfolk Island",
																		"Northern Mariana Islands",
																		"Norway",
																		"Oman",
																		"Other",
																		"Pakistan",
																		"Palau",
																		"Palestinian Territory",
																		"Panama",
																		"Papua New Guinea",
																		"Paraguay",
																		"Peru",
																		"Philippines",
																		"Pitcairn Islands",
																		"Poland",
																		"Portugal",
																		"Puerto Rico",
																		"Qatar",
																		"Reunion",
																		"Romania",
																		"Russian Federation",
																		"Rwanda",
																		"Saint Barthelemy",
																		"Saint Helena",
																		"Saint Kitts and Nevis",
																		"Saint Lucia",
																		"Saint Martin",
																		"Saint Pierre and Miquelon",
																		"Saint Vincent and the Grenadines",
																		"Samoa",
																		"San Marino",
																		"Sao Tome and Principe",
																		"Satellite Provider",
																		"Saudi Arabia",
																		"Senegal",
																		"Serbia",
																		"Seychelles",
																		"Sierra Leone",
																		"Singapore",
																		"Slovakia",
																		"Slovenia",
																		"Solomon Islands",
																		"Somalia",
																		"South Africa",
																		"South Georgia and the South Sandwich Islands",
																		"Spain",
																		"Sri Lanka",
																		"Sudan",
																		"Suriname",
																		"Svalbard and Jan Mayen",
																		"Swaziland",
																		"Sweden",
																		"Switzerland",
																		"Syrian Arab Republic",
																		"Taiwan",
																		"Tajikistan",
																		"Tanzania, United Republic of",
																		"Thailand",
																		"Timor-Leste",
																		"Togo",
																		"Tokelau",
																		"Tonga",
																		"Trinidad and Tobago",
																		"Tunisia",
																		"Turkey",
																		"Turkmenistan",
																		"Turks and Caicos Islands",
																		"Tuvalu",
																		"Uganda",
																		"Ukraine",
																		"United Arab Emirates",
																		"United Kingdom",
																		"United States",
																		"United States Minor Outlying Islands",
																		"Uruguay",
																		"Uzbekistan",
																		"Vanuatu",
																		"Venezuela",
																		"Vietnam",
																		"Virgin Islands, British",
																		"Virgin Islands, U.S.",
																		"Wallis and Futuna",
																		"Western Sahara",
																		"Yemen",
																		"Zambia",
																		"Zimbabwe",
																	},
																	"type": "string",
																},
															},
															"type": "object",
														},
														"type": "array",
													},
													"disallowedGeolocationReference": map[string]interface{}{
														"properties": map[string]interface{}{
															"link": map[string]interface{}{
																"pattern": "^http",
																"type":    "string",
															},
														},
														"type": "object",
													},
													"enablePassiveMode": map[string]interface{}{
														"type": "boolean",
													},
													"enforcementMode": map[string]interface{}{
														"enum": []interface{}{
															"transparent",
															"blocking",
														},
														"type": "string",
													},
													"enforcer-settings": map[string]interface{}{
														"properties": map[string]interface{}{
															"enforcerStateCookies": map[string]interface{}{
																"properties": map[string]interface{}{
																	"httpOnlyAttribute": map[string]interface{}{
																		"type": "boolean",
																	},
																	"sameSiteAttribute": map[string]interface{}{
																		"enum": []interface{}{
																			"lax",
																			"none",
																			"none-value",
																			"strict",
																		},
																		"type": "string",
																	},
																	"secureAttribute": map[string]interface{}{
																		"enum": []interface{}{
																			"always",
																			"never",
																		},
																		"type": "string",
																	},
																},
																"type": "object",
															},
														},
														"type": "object",
													},
													"filetypeReference": map[string]interface{}{
														"properties": map[string]interface{}{
															"link": map[string]interface{}{
																"pattern": "^http",
																"type":    "string",
															},
														},
														"type": "object",
													},
													"filetypes": map[string]interface{}{
														"items": map[string]interface{}{
															"properties": map[string]interface{}{
																"$action": map[string]interface{}{
																	"enum": []interface{}{
																		"delete",
																	},
																	"type": "string",
																},
																"allowed": map[string]interface{}{
																	"type": "boolean",
																},
																"checkPostDataLength": map[string]interface{}{
																	"type": "boolean",
																},
																"checkQueryStringLength": map[string]interface{}{
																	"type": "boolean",
																},
																"checkRequestLength": map[string]interface{}{
																	"type": "boolean",
																},
																"checkUrlLength": map[string]interface{}{
																	"type": "boolean",
																},
																"name": map[string]interface{}{
																	"type": "string",
																},
																"postDataLength": map[string]interface{}{
																	"type": "integer",
																},
																"queryStringLength": map[string]interface{}{
																	"type": "integer",
																},
																"requestLength": map[string]interface{}{
																	"type": "integer",
																},
																"responseCheck": map[string]interface{}{
																	"type": "boolean",
																},
																"type": map[string]interface{}{
																	"enum": []interface{}{
																		"explicit",
																		"wildcard",
																	},
																	"type": "string",
																},
																"urlLength": map[string]interface{}{
																	"type": "integer",
																},
																"wildcardOrder": map[string]interface{}{
																	"type": "integer",
																},
															},
															"type": "object",
														},
														"type": "array",
													},
													"fullPath": map[string]interface{}{
														"type": "string",
													},
													"general": map[string]interface{}{
														"properties": map[string]interface{}{
															"allowedResponseCodes": map[string]interface{}{
																"items": map[string]interface{}{
																	"format":  "int32",
																	"maximum": 999,
																	"minimum": 100,
																	"type":    "integer",
																},
																"type": "array",
															},
															"customXffHeaders": map[string]interface{}{
																"items": map[string]interface{}{
																	"type": "string",
																},
																"type": "array",
															},
															"maskCreditCardNumbersInRequest": map[string]interface{}{
																"type": "boolean",
															},
															"trustXff": map[string]interface{}{
																"type": "boolean",
															},
														},
														"type": "object",
													},
													"generalReference": map[string]interface{}{
														"properties": map[string]interface{}{
															"link": map[string]interface{}{
																"pattern": "^http",
																"type":    "string",
															},
														},
														"type": "object",
													},
													"graphql-profiles": map[string]interface{}{
														"items": map[string]interface{}{
															"properties": map[string]interface{}{
																"$action": map[string]interface{}{
																	"enum": []interface{}{
																		"delete",
																	},
																	"type": "string",
																},
																"attackSignaturesCheck": map[string]interface{}{
																	"type": "boolean",
																},
																"defenseAttributes": map[string]interface{}{
																	"properties": map[string]interface{}{
																		"allowIntrospectionQueries": map[string]interface{}{
																			"type": "boolean",
																		},
																		"maximumBatchedQueries": map[string]interface{}{
																			"anyOf": []interface{}{
																				map[string]interface{}{
																					"type": "integer",
																				},
																				map[string]interface{}{
																					"type": "string",
																				},
																			},
																			"x-kubernetes-int-or-string": true,
																		},
																		"maximumQueryCost": map[string]interface{}{
																			"anyOf": []interface{}{
																				map[string]interface{}{
																					"type": "integer",
																				},
																				map[string]interface{}{
																					"type": "string",
																				},
																			},
																			"x-kubernetes-int-or-string": true,
																		},
																		"maximumStructureDepth": map[string]interface{}{
																			"anyOf": []interface{}{
																				map[string]interface{}{
																					"type": "integer",
																				},
																				map[string]interface{}{
																					"type": "string",
																				},
																			},
																			"x-kubernetes-int-or-string": true,
																		},
																		"maximumTotalLength": map[string]interface{}{
																			"anyOf": []interface{}{
																				map[string]interface{}{
																					"type": "integer",
																				},
																				map[string]interface{}{
																					"type": "string",
																				},
																			},
																			"x-kubernetes-int-or-string": true,
																		},
																		"maximumValueLength": map[string]interface{}{
																			"anyOf": []interface{}{
																				map[string]interface{}{
																					"type": "integer",
																				},
																				map[string]interface{}{
																					"type": "string",
																				},
																			},
																			"x-kubernetes-int-or-string": true,
																		},
																		"tolerateParsingWarnings": map[string]interface{}{
																			"type": "boolean",
																		},
																	},
																	"type": "object",
																},
																"description": map[string]interface{}{
																	"type": "string",
																},
																"metacharElementCheck": map[string]interface{}{
																	"type": "boolean",
																},
																"metacharOverrides": map[string]interface{}{
																	"items": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"isAllowed": map[string]interface{}{
																				"type": "boolean",
																			},
																			"metachar": map[string]interface{}{
																				"type": "string",
																			},
																		},
																		"type": "object",
																	},
																	"type": "array",
																},
																"name": map[string]interface{}{
																	"type": "string",
																},
																"responseEnforcement": map[string]interface{}{
																	"properties": map[string]interface{}{
																		"blockDisallowedPatterns": map[string]interface{}{
																			"type": "boolean",
																		},
																		"disallowedPatterns": map[string]interface{}{
																			"items": map[string]interface{}{
																				"type": "string",
																			},
																			"type": "array",
																		},
																	},
																	"type": "object",
																},
																"sensitiveData": map[string]interface{}{
																	"items": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"parameterName": map[string]interface{}{
																				"type": "string",
																			},
																		},
																		"type": "object",
																	},
																	"type": "array",
																},
																"signatureOverrides": map[string]interface{}{
																	"items": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"enabled": map[string]interface{}{
																				"type": "boolean",
																			},
																			"name": map[string]interface{}{
																				"type": "string",
																			},
																			"signatureId": map[string]interface{}{
																				"type": "integer",
																			},
																			"tag": map[string]interface{}{
																				"type": "string",
																			},
																		},
																		"type": "object",
																	},
																	"type": "array",
																},
															},
															"type": "object",
														},
														"type": "array",
													},
													"grpc-profiles": map[string]interface{}{
														"items": map[string]interface{}{
															"properties": map[string]interface{}{
																"$action": map[string]interface{}{
																	"enum": []interface{}{
																		"delete",
																	},
																	"type": "string",
																},
																"associateUrls": map[string]interface{}{
																	"type": "boolean",
																},
																"attackSignaturesCheck": map[string]interface{}{
																	"type": "boolean",
																},
																"decodeStringValuesAsBase64": map[string]interface{}{
																	"enum": []interface{}{
																		"disabled",
																		"enabled",
																	},
																	"type": "string",
																},
																"defenseAttributes": map[string]interface{}{
																	"properties": map[string]interface{}{
																		"allowUnknownFields": map[string]interface{}{
																			"type": "boolean",
																		},
																		"maximumDataLength": map[string]interface{}{
																			"anyOf": []interface{}{
																				map[string]interface{}{
																					"type": "integer",
																				},
																				map[string]interface{}{
																					"type": "string",
																				},
																			},
																			"x-kubernetes-int-or-string": true,
																		},
																	},
																	"type": "object",
																},
																"description": map[string]interface{}{
																	"type": "string",
																},
																"hasIdlFiles": map[string]interface{}{
																	"type": "boolean",
																},
																"idlFiles": map[string]interface{}{
																	"items": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"idlFile": map[string]interface{}{
																				"properties": map[string]interface{}{
																					"contents": map[string]interface{}{
																						"type": "string",
																					},
																					"fileName": map[string]interface{}{
																						"type": "string",
																					},
																					"isBase64": map[string]interface{}{
																						"type": "boolean",
																					},
																				},
																				"type": "object",
																			},
																			"importUrl": map[string]interface{}{
																				"type": "string",
																			},
																			"isPrimary": map[string]interface{}{
																				"type": "boolean",
																			},
																			"primaryIdlFileName": map[string]interface{}{
																				"type": "string",
																			},
																		},
																		"type": "object",
																	},
																	"type": "array",
																},
																"metacharCheck": map[string]interface{}{
																	"type": "boolean",
																},
																"metacharElementCheck": map[string]interface{}{
																	"type": "boolean",
																},
																"name": map[string]interface{}{
																	"type": "string",
																},
																"signatureOverrides": map[string]interface{}{
																	"items": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"enabled": map[string]interface{}{
																				"type": "boolean",
																			},
																			"name": map[string]interface{}{
																				"type": "string",
																			},
																			"signatureId": map[string]interface{}{
																				"type": "integer",
																			},
																			"tag": map[string]interface{}{
																				"type": "string",
																			},
																		},
																		"type": "object",
																	},
																	"type": "array",
																},
															},
															"type": "object",
														},
														"type": "array",
													},
													"header-settings": map[string]interface{}{
														"properties": map[string]interface{}{
															"maximumHttpHeaderLength": map[string]interface{}{
																"anyOf": []interface{}{
																	map[string]interface{}{
																		"type": "integer",
																	},
																	map[string]interface{}{
																		"type": "string",
																	},
																},
																"x-kubernetes-int-or-string": true,
															},
														},
														"type": "object",
													},
													"headerReference": map[string]interface{}{
														"properties": map[string]interface{}{
															"link": map[string]interface{}{
																"pattern": "^http",
																"type":    "string",
															},
														},
														"type": "object",
													},
													"headerSettingsReference": map[string]interface{}{
														"properties": map[string]interface{}{
															"link": map[string]interface{}{
																"pattern": "^http",
																"type":    "string",
															},
														},
														"type": "object",
													},
													"headers": map[string]interface{}{
														"items": map[string]interface{}{
															"properties": map[string]interface{}{
																"$action": map[string]interface{}{
																	"enum": []interface{}{
																		"delete",
																	},
																	"type": "string",
																},
																"allowRepeatedOccurrences": map[string]interface{}{
																	"type": "boolean",
																},
																"base64Decoding": map[string]interface{}{
																	"type": "boolean",
																},
																"checkSignatures": map[string]interface{}{
																	"type": "boolean",
																},
																"decodeValueAsBase64": map[string]interface{}{
																	"enum": []interface{}{
																		"enabled",
																		"disabled",
																		"required",
																	},
																	"type": "string",
																},
																"htmlNormalization": map[string]interface{}{
																	"type": "boolean",
																},
																"mandatory": map[string]interface{}{
																	"type": "boolean",
																},
																"maskValueInLogs": map[string]interface{}{
																	"type": "boolean",
																},
																"name": map[string]interface{}{
																	"type": "string",
																},
																"normalizationViolations": map[string]interface{}{
																	"type": "boolean",
																},
																"percentDecoding": map[string]interface{}{
																	"type": "boolean",
																},
																"signatureOverrides": map[string]interface{}{
																	"items": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"enabled": map[string]interface{}{
																				"type": "boolean",
																			},
																			"name": map[string]interface{}{
																				"type": "string",
																			},
																			"signatureId": map[string]interface{}{
																				"type": "integer",
																			},
																			"tag": map[string]interface{}{
																				"type": "string",
																			},
																		},
																		"type": "object",
																	},
																	"type": "array",
																},
																"type": map[string]interface{}{
																	"enum": []interface{}{
																		"explicit",
																		"wildcard",
																	},
																	"type": "string",
																},
																"urlNormalization": map[string]interface{}{
																	"type": "boolean",
																},
																"wildcardOrder": map[string]interface{}{
																	"type": "integer",
																},
															},
															"type": "object",
														},
														"type": "array",
													},
													"host-names": map[string]interface{}{
														"items": map[string]interface{}{
															"properties": map[string]interface{}{
																"$action": map[string]interface{}{
																	"enum": []interface{}{
																		"delete",
																	},
																	"type": "string",
																},
																"includeSubdomains": map[string]interface{}{
																	"type": "boolean",
																},
																"name": map[string]interface{}{
																	"type": "string",
																},
															},
															"type": "object",
														},
														"type": "array",
													},
													"idl-files": map[string]interface{}{
														"items": map[string]interface{}{
															"properties": map[string]interface{}{
																"contents": map[string]interface{}{
																	"type": "string",
																},
																"fileName": map[string]interface{}{
																	"type": "string",
																},
																"isBase64": map[string]interface{}{
																	"type": "boolean",
																},
															},
															"type": "object",
														},
														"type": "array",
													},
													"ip-intelligence": map[string]interface{}{
														"type":                                 "object",
														"x-kubernetes-preserve-unknown-fields": true,
													},
													"json-profiles": map[string]interface{}{
														"items": map[string]interface{}{
															"properties": map[string]interface{}{
																"$action": map[string]interface{}{
																	"enum": []interface{}{
																		"delete",
																	},
																	"type": "string",
																},
																"attackSignaturesCheck": map[string]interface{}{
																	"type": "boolean",
																},
																"defenseAttributes": map[string]interface{}{
																	"properties": map[string]interface{}{
																		"maximumArrayLength": map[string]interface{}{
																			"anyOf": []interface{}{
																				map[string]interface{}{
																					"type": "integer",
																				},
																				map[string]interface{}{
																					"type": "string",
																				},
																			},
																			"x-kubernetes-int-or-string": true,
																		},
																		"maximumStructureDepth": map[string]interface{}{
																			"anyOf": []interface{}{
																				map[string]interface{}{
																					"type": "integer",
																				},
																				map[string]interface{}{
																					"type": "string",
																				},
																			},
																			"x-kubernetes-int-or-string": true,
																		},
																		"maximumTotalLengthOfJSONData": map[string]interface{}{
																			"anyOf": []interface{}{
																				map[string]interface{}{
																					"type": "integer",
																				},
																				map[string]interface{}{
																					"type": "string",
																				},
																			},
																			"x-kubernetes-int-or-string": true,
																		},
																		"maximumValueLength": map[string]interface{}{
																			"anyOf": []interface{}{
																				map[string]interface{}{
																					"type": "integer",
																				},
																				map[string]interface{}{
																					"type": "string",
																				},
																			},
																			"x-kubernetes-int-or-string": true,
																		},
																		"tolerateJSONParsingWarnings": map[string]interface{}{
																			"type": "boolean",
																		},
																	},
																	"type": "object",
																},
																"description": map[string]interface{}{
																	"type": "string",
																},
																"handleJsonValuesAsParameters": map[string]interface{}{
																	"type": "boolean",
																},
																"hasValidationFiles": map[string]interface{}{
																	"type": "boolean",
																},
																"metacharOverrides": map[string]interface{}{
																	"items": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"isAllowed": map[string]interface{}{
																				"type": "boolean",
																			},
																			"metachar": map[string]interface{}{
																				"type": "string",
																			},
																		},
																		"type": "object",
																	},
																	"type": "array",
																},
																"name": map[string]interface{}{
																	"type": "string",
																},
																"signatureOverrides": map[string]interface{}{
																	"items": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"enabled": map[string]interface{}{
																				"type": "boolean",
																			},
																			"name": map[string]interface{}{
																				"type": "string",
																			},
																			"signatureId": map[string]interface{}{
																				"type": "integer",
																			},
																			"tag": map[string]interface{}{
																				"type": "string",
																			},
																		},
																		"type": "object",
																	},
																	"type": "array",
																},
																"validationFiles": map[string]interface{}{
																	"items": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"importUrl": map[string]interface{}{
																				"type": "string",
																			},
																			"isPrimary": map[string]interface{}{
																				"type": "boolean",
																			},
																			"jsonValidationFile": map[string]interface{}{
																				"properties": map[string]interface{}{
																					"$action": map[string]interface{}{
																						"enum": []interface{}{
																							"delete",
																						},
																						"type": "string",
																					},
																					"contents": map[string]interface{}{
																						"type": "string",
																					},
																					"fileName": map[string]interface{}{
																						"type": "string",
																					},
																					"isBase64": map[string]interface{}{
																						"type": "boolean",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																	"type": "array",
																},
															},
															"type": "object",
														},
														"type": "array",
													},
													"json-validation-files": map[string]interface{}{
														"items": map[string]interface{}{
															"properties": map[string]interface{}{
																"$action": map[string]interface{}{
																	"enum": []interface{}{
																		"delete",
																	},
																	"type": "string",
																},
																"contents": map[string]interface{}{
																	"type": "string",
																},
																"fileName": map[string]interface{}{
																	"type": "string",
																},
																"isBase64": map[string]interface{}{
																	"type": "boolean",
																},
															},
															"type": "object",
														},
														"type": "array",
													},
													"jsonProfileReference": map[string]interface{}{
														"properties": map[string]interface{}{
															"link": map[string]interface{}{
																"pattern": "^http",
																"type":    "string",
															},
														},
														"type": "object",
													},
													"jsonValidationFileReference": map[string]interface{}{
														"properties": map[string]interface{}{
															"link": map[string]interface{}{
																"pattern": "^http",
																"type":    "string",
															},
														},
														"type": "object",
													},
													"methodReference": map[string]interface{}{
														"properties": map[string]interface{}{
															"link": map[string]interface{}{
																"pattern": "^http",
																"type":    "string",
															},
														},
														"type": "object",
													},
													"methods": map[string]interface{}{
														"items": map[string]interface{}{
															"properties": map[string]interface{}{
																"$action": map[string]interface{}{
																	"enum": []interface{}{
																		"delete",
																	},
																	"type": "string",
																},
																"name": map[string]interface{}{
																	"type": "string",
																},
															},
															"type": "object",
														},
														"type": "array",
													},
													"name": map[string]interface{}{
														"type": "string",
													},
													"open-api-files": map[string]interface{}{
														"items": map[string]interface{}{
															"properties": map[string]interface{}{
																"link": map[string]interface{}{
																	"pattern": "^http",
																	"type":    "string",
																},
															},
															"type": "object",
														},
														"type": "array",
													},
													"parameterReference": map[string]interface{}{
														"properties": map[string]interface{}{
															"link": map[string]interface{}{
																"pattern": "^http",
																"type":    "string",
															},
														},
														"type": "object",
													},
													"parameters": map[string]interface{}{
														"items": map[string]interface{}{
															"properties": map[string]interface{}{
																"$action": map[string]interface{}{
																	"enum": []interface{}{
																		"delete",
																	},
																	"type": "string",
																},
																"allowEmptyValue": map[string]interface{}{
																	"type": "boolean",
																},
																"allowRepeatedParameterName": map[string]interface{}{
																	"type": "boolean",
																},
																"arraySerializationFormat": map[string]interface{}{
																	"enum": []interface{}{
																		"csv",
																		"form",
																		"label",
																		"matrix",
																		"multi",
																		"multipart",
																		"pipe",
																		"ssv",
																		"tsv",
																	},
																	"type": "string",
																},
																"attackSignaturesCheck": map[string]interface{}{
																	"type": "boolean",
																},
																"checkMaxValue": map[string]interface{}{
																	"type": "boolean",
																},
																"checkMaxValueLength": map[string]interface{}{
																	"type": "boolean",
																},
																"checkMetachars": map[string]interface{}{
																	"type": "boolean",
																},
																"checkMinValue": map[string]interface{}{
																	"type": "boolean",
																},
																"checkMinValueLength": map[string]interface{}{
																	"type": "boolean",
																},
																"checkMultipleOfValue": map[string]interface{}{
																	"type": "boolean",
																},
																"contentProfile": map[string]interface{}{
																	"properties": map[string]interface{}{
																		"name": map[string]interface{}{
																			"type": "string",
																		},
																	},
																	"type": "object",
																},
																"dataType": map[string]interface{}{
																	"enum": []interface{}{
																		"alpha-numeric",
																		"binary",
																		"boolean",
																		"decimal",
																		"email",
																		"integer",
																		"none",
																		"phone",
																	},
																	"type": "string",
																},
																"decodeValueAsBase64": map[string]interface{}{
																	"enum": []interface{}{
																		"enabled",
																		"disabled",
																		"required",
																	},
																	"type": "string",
																},
																"disallowFileUploadOfExecutables": map[string]interface{}{
																	"type": "boolean",
																},
																"enableRegularExpression": map[string]interface{}{
																	"type": "boolean",
																},
																"exclusiveMax": map[string]interface{}{
																	"type": "boolean",
																},
																"exclusiveMin": map[string]interface{}{
																	"type": "boolean",
																},
																"isBase64": map[string]interface{}{
																	"type": "boolean",
																},
																"isCookie": map[string]interface{}{
																	"type": "boolean",
																},
																"isHeader": map[string]interface{}{
																	"type": "boolean",
																},
																"level": map[string]interface{}{
																	"enum": []interface{}{
																		"global",
																		"url",
																	},
																	"type": "string",
																},
																"mandatory": map[string]interface{}{
																	"type": "boolean",
																},
																"maximumLength": map[string]interface{}{
																	"type": "integer",
																},
																"maximumValue": map[string]interface{}{
																	"type": "integer",
																},
																"metacharsOnParameterValueCheck": map[string]interface{}{
																	"type": "boolean",
																},
																"minimumLength": map[string]interface{}{
																	"type": "integer",
																},
																"minimumValue": map[string]interface{}{
																	"type": "integer",
																},
																"multipleOf": map[string]interface{}{
																	"type": "integer",
																},
																"name": map[string]interface{}{
																	"type": "string",
																},
																"nameMetacharOverrides": map[string]interface{}{
																	"items": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"isAllowed": map[string]interface{}{
																				"type": "boolean",
																			},
																			"metachar": map[string]interface{}{
																				"type": "string",
																			},
																		},
																		"type": "object",
																	},
																	"type": "array",
																},
																"objectSerializationStyle": map[string]interface{}{
																	"type": "string",
																},
																"parameterEnumValues": map[string]interface{}{
																	"items": map[string]interface{}{
																		"type": "string",
																	},
																	"type": "array",
																},
																"parameterLocation": map[string]interface{}{
																	"enum": []interface{}{
																		"any",
																		"cookie",
																		"form-data",
																		"header",
																		"path",
																		"query",
																	},
																	"type": "string",
																},
																"regularExpression": map[string]interface{}{
																	"type": "string",
																},
																"sensitiveParameter": map[string]interface{}{
																	"type": "boolean",
																},
																"signatureOverrides": map[string]interface{}{
																	"items": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"enabled": map[string]interface{}{
																				"type": "boolean",
																			},
																			"name": map[string]interface{}{
																				"type": "string",
																			},
																			"signatureId": map[string]interface{}{
																				"type": "integer",
																			},
																			"tag": map[string]interface{}{
																				"type": "string",
																			},
																		},
																		"type": "object",
																	},
																	"type": "array",
																},
																"staticValues": map[string]interface{}{
																	"type": "string",
																},
																"type": map[string]interface{}{
																	"enum": []interface{}{
																		"explicit",
																		"wildcard",
																	},
																	"type": "string",
																},
																"url": map[string]interface{}{
																	"properties": map[string]interface{}{
																		"method": map[string]interface{}{
																			"enum": []interface{}{
																				"ACL",
																				"BCOPY",
																				"BDELETE",
																				"BMOVE",
																				"BPROPFIND",
																				"BPROPPATCH",
																				"CHECKIN",
																				"CHECKOUT",
																				"CONNECT",
																				"COPY",
																				"DELETE",
																				"GET",
																				"HEAD",
																				"LINK",
																				"LOCK",
																				"MERGE",
																				"MKCOL",
																				"MKWORKSPACE",
																				"MOVE",
																				"NOTIFY",
																				"OPTIONS",
																				"PATCH",
																				"POLL",
																				"POST",
																				"PROPFIND",
																				"PROPPATCH",
																				"PUT",
																				"REPORT",
																				"RPC_IN_DATA",
																				"RPC_OUT_DATA",
																				"SEARCH",
																				"SUBSCRIBE",
																				"TRACE",
																				"TRACK",
																				"UNLINK",
																				"UNLOCK",
																				"UNSUBSCRIBE",
																				"VERSION_CONTROL",
																				"X-MS-ENUMATTS",
																				"*",
																			},
																			"type": "string",
																		},
																		"name": map[string]interface{}{
																			"type": "string",
																		},
																		"protocol": map[string]interface{}{
																			"enum": []interface{}{
																				"http",
																				"https",
																			},
																			"type": "string",
																		},
																		"type": map[string]interface{}{
																			"enum": []interface{}{
																				"explicit",
																				"wildcard",
																			},
																			"type": "string",
																		},
																	},
																	"type": "object",
																},
																"valueMetacharOverrides": map[string]interface{}{
																	"items": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"isAllowed": map[string]interface{}{
																				"type": "boolean",
																			},
																			"metachar": map[string]interface{}{
																				"type": "string",
																			},
																		},
																		"type": "object",
																	},
																	"type": "array",
																},
																"valueType": map[string]interface{}{
																	"enum": []interface{}{
																		"array",
																		"auto-detect",
																		"dynamic-content",
																		"dynamic-parameter-name",
																		"ignore",
																		"json",
																		"object",
																		"openapi-array",
																		"static-content",
																		"user-input",
																		"xml",
																	},
																	"type": "string",
																},
																"wildcardOrder": map[string]interface{}{
																	"type": "integer",
																},
															},
															"type": "object",
														},
														"type": "array",
													},
													"response-pages": map[string]interface{}{
														"items": map[string]interface{}{
															"properties": map[string]interface{}{
																"ajaxActionType": map[string]interface{}{
																	"enum": []interface{}{
																		"alert-popup",
																		"custom",
																		"redirect",
																	},
																	"type": "string",
																},
																"ajaxCustomContent": map[string]interface{}{
																	"type": "string",
																},
																"ajaxEnabled": map[string]interface{}{
																	"type": "boolean",
																},
																"ajaxPopupMessage": map[string]interface{}{
																	"type": "string",
																},
																"ajaxRedirectUrl": map[string]interface{}{
																	"type": "string",
																},
																"grpcStatusCode": map[string]interface{}{
																	"pattern": "ABORTED|ALREADY_EXISTS|CANCELLED|DATA_LOSS|DEADLINE_EXCEEDED|FAILED_PRECONDITION|INTERNAL|INVALID_ARGUMENT|NOT_FOUND|OK|OUT_OF_RANGE|PERMISSION_DENIED|RESOURCE_EXHAUSTED|UNAUTHENTICATED|UNAVAILABLE|UNIMPLEMENTED|UNKNOWN|d+",
																	"type":    "string",
																},
																"grpcStatusMessage": map[string]interface{}{
																	"type": "string",
																},
																"responseActionType": map[string]interface{}{
																	"enum": []interface{}{
																		"custom",
																		"default",
																		"erase-cookies",
																		"redirect",
																		"soap-fault",
																	},
																	"type": "string",
																},
																"responseContent": map[string]interface{}{
																	"type": "string",
																},
																"responseHeader": map[string]interface{}{
																	"type": "string",
																},
																"responsePageType": map[string]interface{}{
																	"enum": []interface{}{
																		"ajax",
																		"ajax-login",
																		"captcha",
																		"captcha-fail",
																		"default",
																		"failed-login-honeypot",
																		"failed-login-honeypot-ajax",
																		"hijack",
																		"leaked-credentials",
																		"leaked-credentials-ajax",
																		"mobile",
																		"persistent-flow",
																		"xml",
																		"grpc",
																	},
																	"type": "string",
																},
																"responseRedirectUrl": map[string]interface{}{
																	"type": "string",
																},
															},
															"type": "object",
														},
														"type": "array",
													},
													"responsePageReference": map[string]interface{}{
														"properties": map[string]interface{}{
															"link": map[string]interface{}{
																"pattern": "^http",
																"type":    "string",
															},
														},
														"type": "object",
													},
													"sensitive-parameters": map[string]interface{}{
														"items": map[string]interface{}{
															"properties": map[string]interface{}{
																"$action": map[string]interface{}{
																	"enum": []interface{}{
																		"delete",
																	},
																	"type": "string",
																},
																"name": map[string]interface{}{
																	"type": "string",
																},
															},
															"type": "object",
														},
														"type": "array",
													},
													"sensitiveParameterReference": map[string]interface{}{
														"properties": map[string]interface{}{
															"link": map[string]interface{}{
																"pattern": "^http",
																"type":    "string",
															},
														},
														"type": "object",
													},
													"server-technologies": map[string]interface{}{
														"items": map[string]interface{}{
															"properties": map[string]interface{}{
																"$action": map[string]interface{}{
																	"enum": []interface{}{
																		"delete",
																	},
																	"type": "string",
																},
																"serverTechnologyName": map[string]interface{}{
																	"enum": []interface{}{
																		"Jenkins",
																		"SharePoint",
																		"Oracle Application Server",
																		"Python",
																		"Oracle Identity Manager",
																		"Spring Boot",
																		"CouchDB",
																		"SQLite",
																		"Handlebars",
																		"Mustache",
																		"Prototype",
																		"Zend",
																		"Redis",
																		"Underscore.js",
																		"Ember.js",
																		"ZURB Foundation",
																		"ef.js",
																		"Vue.js",
																		"UIKit",
																		"TYPO3 CMS",
																		"RequireJS",
																		"React",
																		"MooTools",
																		"Laravel",
																		"GraphQL",
																		"Google Web Toolkit",
																		"Express.js",
																		"CodeIgniter",
																		"Backbone.js",
																		"AngularJS",
																		"JavaScript",
																		"Nginx",
																		"Jetty",
																		"Joomla",
																		"JavaServer Faces (JSF)",
																		"Ruby",
																		"MongoDB",
																		"Django",
																		"Node.js",
																		"Citrix",
																		"JBoss",
																		"Elasticsearch",
																		"Apache Struts",
																		"XML",
																		"PostgreSQL",
																		"IBM DB2",
																		"Sybase/ASE",
																		"CGI",
																		"Proxy Servers",
																		"SSI (Server Side Includes)",
																		"Cisco",
																		"Novell",
																		"Macromedia JRun",
																		"BEA Systems WebLogic Server",
																		"Lotus Domino",
																		"MySQL",
																		"Oracle",
																		"Microsoft SQL Server",
																		"PHP",
																		"Outlook Web Access",
																		"Apache/NCSA HTTP Server",
																		"Apache Tomcat",
																		"WordPress",
																		"Macromedia ColdFusion",
																		"Unix/Linux",
																		"Microsoft Windows",
																		"ASP.NET",
																		"Front Page Server Extensions (FPSE)",
																		"IIS",
																		"WebDAV",
																		"ASP",
																		"Java Servlets/JSP",
																		"jQuery",
																	},
																	"type": "string",
																},
															},
															"type": "object",
														},
														"type": "array",
													},
													"serverTechnologyReference": map[string]interface{}{
														"properties": map[string]interface{}{
															"link": map[string]interface{}{
																"pattern": "^http",
																"type":    "string",
															},
														},
														"type": "object",
													},
													"signature-requirements": map[string]interface{}{
														"items": map[string]interface{}{
															"properties": map[string]interface{}{
																"$action": map[string]interface{}{
																	"enum": []interface{}{
																		"delete",
																	},
																	"type": "string",
																},
																"tag": map[string]interface{}{
																	"type": "string",
																},
															},
															"type": "object",
														},
														"type": "array",
													},
													"signature-sets": map[string]interface{}{
														"items": map[string]interface{}{
															"properties": map[string]interface{}{
																"$action": map[string]interface{}{
																	"enum": []interface{}{
																		"delete",
																	},
																	"type": "string",
																},
																"alarm": map[string]interface{}{
																	"type": "boolean",
																},
																"block": map[string]interface{}{
																	"type": "boolean",
																},
																"name": map[string]interface{}{
																	"type": "string",
																},
															},
															"type":                                 "object",
															"x-kubernetes-preserve-unknown-fields": true,
														},
														"type": "array",
													},
													"signature-settings": map[string]interface{}{
														"properties": map[string]interface{}{
															"attackSignatureFalsePositiveMode": map[string]interface{}{
																"enum": []interface{}{
																	"detect",
																	"detect-and-allow",
																	"disabled",
																},
																"type": "string",
															},
															"minimumAccuracyForAutoAddedSignatures": map[string]interface{}{
																"enum": []interface{}{
																	"high",
																	"low",
																	"medium",
																},
																"type": "string",
															},
														},
														"type": "object",
													},
													"signatureReference": map[string]interface{}{
														"properties": map[string]interface{}{
															"link": map[string]interface{}{
																"pattern": "^http",
																"type":    "string",
															},
														},
														"type": "object",
													},
													"signatureSetReference": map[string]interface{}{
														"properties": map[string]interface{}{
															"link": map[string]interface{}{
																"pattern": "^http",
																"type":    "string",
															},
														},
														"type": "object",
													},
													"signatureSettingReference": map[string]interface{}{
														"properties": map[string]interface{}{
															"link": map[string]interface{}{
																"pattern": "^http",
																"type":    "string",
															},
														},
														"type": "object",
													},
													"signatures": map[string]interface{}{
														"items": map[string]interface{}{
															"properties": map[string]interface{}{
																"enabled": map[string]interface{}{
																	"type": "boolean",
																},
																"name": map[string]interface{}{
																	"type": "string",
																},
																"signatureId": map[string]interface{}{
																	"type": "integer",
																},
																"tag": map[string]interface{}{
																	"type": "string",
																},
															},
															"type": "object",
														},
														"type": "array",
													},
													"softwareVersion": map[string]interface{}{
														"type": "string",
													},
													"template": map[string]interface{}{
														"properties": map[string]interface{}{
															"name": map[string]interface{}{
																"type": "string",
															},
														},
														"type": "object",
													},
													"threat-campaigns": map[string]interface{}{
														"items": map[string]interface{}{
															"properties": map[string]interface{}{
																"isEnabled": map[string]interface{}{
																	"type": "boolean",
																},
																"name": map[string]interface{}{
																	"type": "string",
																},
															},
															"type": "object",
														},
														"type": "array",
													},
													"threatCampaignReference": map[string]interface{}{
														"properties": map[string]interface{}{
															"link": map[string]interface{}{
																"pattern": "^http",
																"type":    "string",
															},
														},
														"type": "object",
													},
													"urlReference": map[string]interface{}{
														"properties": map[string]interface{}{
															"link": map[string]interface{}{
																"pattern": "^http",
																"type":    "string",
															},
														},
														"type": "object",
													},
													"urls": map[string]interface{}{
														"items": map[string]interface{}{
															"properties": map[string]interface{}{
																"$action": map[string]interface{}{
																	"enum": []interface{}{
																		"delete",
																	},
																	"type": "string",
																},
																"allowRenderingInFrames": map[string]interface{}{
																	"enum": []interface{}{
																		"never",
																		"only-same",
																	},
																	"type": "string",
																},
																"allowRenderingInFramesOnlyFrom": map[string]interface{}{
																	"type": "string",
																},
																"attackSignaturesCheck": map[string]interface{}{
																	"type": "boolean",
																},
																"clickjackingProtection": map[string]interface{}{
																	"type": "boolean",
																},
																"description": map[string]interface{}{
																	"type": "string",
																},
																"disallowFileUploadOfExecutables": map[string]interface{}{
																	"type": "boolean",
																},
																"html5CrossOriginRequestsEnforcement": map[string]interface{}{
																	"properties": map[string]interface{}{
																		"allowOriginsEnforcementMode": map[string]interface{}{
																			"enum": []interface{}{
																				"replace-with",
																				"unmodified",
																			},
																			"type": "string",
																		},
																		"checkAllowedMethods": map[string]interface{}{
																			"type": "boolean",
																		},
																		"crossDomainAllowedOrigin": map[string]interface{}{
																			"items": map[string]interface{}{
																				"properties": map[string]interface{}{
																					"includeSubDomains": map[string]interface{}{
																						"type": "boolean",
																					},
																					"originName": map[string]interface{}{
																						"type": "string",
																					},
																					"originPort": map[string]interface{}{
																						"anyOf": []interface{}{
																							map[string]interface{}{
																								"type": "integer",
																							},
																							map[string]interface{}{
																								"type": "string",
																							},
																						},
																						"x-kubernetes-int-or-string": true,
																					},
																					"originProtocol": map[string]interface{}{
																						"enum": []interface{}{
																							"http",
																							"http/https",
																							"https",
																						},
																						"type": "string",
																					},
																				},
																				"type": "object",
																			},
																			"type": "array",
																		},
																		"enforcementMode": map[string]interface{}{
																			"enum": []interface{}{
																				"disabled",
																				"enforce",
																			},
																			"type": "string",
																		},
																	},
																	"type": "object",
																},
																"isAllowed": map[string]interface{}{
																	"type": "boolean",
																},
																"mandatoryBody": map[string]interface{}{
																	"type": "boolean",
																},
																"metacharOverrides": map[string]interface{}{
																	"items": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"isAllowed": map[string]interface{}{
																				"type": "boolean",
																			},
																			"metachar": map[string]interface{}{
																				"type": "string",
																			},
																		},
																		"type": "object",
																	},
																	"type": "array",
																},
																"metacharsOnUrlCheck": map[string]interface{}{
																	"type": "boolean",
																},
																"method": map[string]interface{}{
																	"enum": []interface{}{
																		"ACL",
																		"BCOPY",
																		"BDELETE",
																		"BMOVE",
																		"BPROPFIND",
																		"BPROPPATCH",
																		"CHECKIN",
																		"CHECKOUT",
																		"CONNECT",
																		"COPY",
																		"DELETE",
																		"GET",
																		"HEAD",
																		"LINK",
																		"LOCK",
																		"MERGE",
																		"MKCOL",
																		"MKWORKSPACE",
																		"MOVE",
																		"NOTIFY",
																		"OPTIONS",
																		"PATCH",
																		"POLL",
																		"POST",
																		"PROPFIND",
																		"PROPPATCH",
																		"PUT",
																		"REPORT",
																		"RPC_IN_DATA",
																		"RPC_OUT_DATA",
																		"SEARCH",
																		"SUBSCRIBE",
																		"TRACE",
																		"TRACK",
																		"UNLINK",
																		"UNLOCK",
																		"UNSUBSCRIBE",
																		"VERSION_CONTROL",
																		"X-MS-ENUMATTS",
																		"*",
																	},
																	"type": "string",
																},
																"methodOverrides": map[string]interface{}{
																	"items": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"allowed": map[string]interface{}{
																				"type": "boolean",
																			},
																			"method": map[string]interface{}{
																				"enum": []interface{}{
																					"ACL",
																					"BCOPY",
																					"BDELETE",
																					"BMOVE",
																					"BPROPFIND",
																					"BPROPPATCH",
																					"CHECKIN",
																					"CHECKOUT",
																					"CONNECT",
																					"COPY",
																					"DELETE",
																					"GET",
																					"HEAD",
																					"LINK",
																					"LOCK",
																					"MERGE",
																					"MKCOL",
																					"MKWORKSPACE",
																					"MOVE",
																					"NOTIFY",
																					"OPTIONS",
																					"PATCH",
																					"POLL",
																					"POST",
																					"PROPFIND",
																					"PROPPATCH",
																					"PUT",
																					"REPORT",
																					"RPC_IN_DATA",
																					"RPC_OUT_DATA",
																					"SEARCH",
																					"SUBSCRIBE",
																					"TRACE",
																					"TRACK",
																					"UNLINK",
																					"UNLOCK",
																					"UNSUBSCRIBE",
																					"VERSION_CONTROL",
																					"X-MS-ENUMATTS",
																				},
																				"type": "string",
																			},
																		},
																		"type": "object",
																	},
																	"type": "array",
																},
																"methodsOverrideOnUrlCheck": map[string]interface{}{
																	"type": "boolean",
																},
																"name": map[string]interface{}{
																	"type": "string",
																},
																"operationId": map[string]interface{}{
																	"type": "string",
																},
																"positionalParameters": map[string]interface{}{
																	"items": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"parameter": map[string]interface{}{
																				"properties": map[string]interface{}{
																					"$action": map[string]interface{}{
																						"enum": []interface{}{
																							"delete",
																						},
																						"type": "string",
																					},
																					"allowEmptyValue": map[string]interface{}{
																						"type": "boolean",
																					},
																					"allowRepeatedParameterName": map[string]interface{}{
																						"type": "boolean",
																					},
																					"arraySerializationFormat": map[string]interface{}{
																						"enum": []interface{}{
																							"csv",
																							"form",
																							"label",
																							"matrix",
																							"multi",
																							"multipart",
																							"pipe",
																							"ssv",
																							"tsv",
																						},
																						"type": "string",
																					},
																					"attackSignaturesCheck": map[string]interface{}{
																						"type": "boolean",
																					},
																					"checkMaxValue": map[string]interface{}{
																						"type": "boolean",
																					},
																					"checkMaxValueLength": map[string]interface{}{
																						"type": "boolean",
																					},
																					"checkMetachars": map[string]interface{}{
																						"type": "boolean",
																					},
																					"checkMinValue": map[string]interface{}{
																						"type": "boolean",
																					},
																					"checkMinValueLength": map[string]interface{}{
																						"type": "boolean",
																					},
																					"checkMultipleOfValue": map[string]interface{}{
																						"type": "boolean",
																					},
																					"contentProfile": map[string]interface{}{
																						"properties": map[string]interface{}{
																							"name": map[string]interface{}{
																								"type": "string",
																							},
																						},
																						"type": "object",
																					},
																					"dataType": map[string]interface{}{
																						"enum": []interface{}{
																							"alpha-numeric",
																							"binary",
																							"boolean",
																							"decimal",
																							"email",
																							"integer",
																							"none",
																							"phone",
																						},
																						"type": "string",
																					},
																					"decodeValueAsBase64": map[string]interface{}{
																						"enum": []interface{}{
																							"enabled",
																							"disabled",
																							"required",
																						},
																						"type": "string",
																					},
																					"disallowFileUploadOfExecutables": map[string]interface{}{
																						"type": "boolean",
																					},
																					"enableRegularExpression": map[string]interface{}{
																						"type": "boolean",
																					},
																					"exclusiveMax": map[string]interface{}{
																						"type": "boolean",
																					},
																					"exclusiveMin": map[string]interface{}{
																						"type": "boolean",
																					},
																					"isBase64": map[string]interface{}{
																						"type": "boolean",
																					},
																					"isCookie": map[string]interface{}{
																						"type": "boolean",
																					},
																					"isHeader": map[string]interface{}{
																						"type": "boolean",
																					},
																					"level": map[string]interface{}{
																						"enum": []interface{}{
																							"global",
																							"url",
																						},
																						"type": "string",
																					},
																					"mandatory": map[string]interface{}{
																						"type": "boolean",
																					},
																					"maximumLength": map[string]interface{}{
																						"type": "integer",
																					},
																					"maximumValue": map[string]interface{}{
																						"type": "integer",
																					},
																					"metacharsOnParameterValueCheck": map[string]interface{}{
																						"type": "boolean",
																					},
																					"minimumLength": map[string]interface{}{
																						"type": "integer",
																					},
																					"minimumValue": map[string]interface{}{
																						"type": "integer",
																					},
																					"multipleOf": map[string]interface{}{
																						"type": "integer",
																					},
																					"name": map[string]interface{}{
																						"type": "string",
																					},
																					"nameMetacharOverrides": map[string]interface{}{
																						"items": map[string]interface{}{
																							"properties": map[string]interface{}{
																								"isAllowed": map[string]interface{}{
																									"type": "boolean",
																								},
																								"metachar": map[string]interface{}{
																									"type": "string",
																								},
																							},
																							"type": "object",
																						},
																						"type": "array",
																					},
																					"objectSerializationStyle": map[string]interface{}{
																						"type": "string",
																					},
																					"parameterEnumValues": map[string]interface{}{
																						"items": map[string]interface{}{
																							"type": "string",
																						},
																						"type": "array",
																					},
																					"parameterLocation": map[string]interface{}{
																						"enum": []interface{}{
																							"any",
																							"cookie",
																							"form-data",
																							"header",
																							"path",
																							"query",
																						},
																						"type": "string",
																					},
																					"regularExpression": map[string]interface{}{
																						"type": "string",
																					},
																					"sensitiveParameter": map[string]interface{}{
																						"type": "boolean",
																					},
																					"signatureOverrides": map[string]interface{}{
																						"items": map[string]interface{}{
																							"properties": map[string]interface{}{
																								"enabled": map[string]interface{}{
																									"type": "boolean",
																								},
																								"name": map[string]interface{}{
																									"type": "string",
																								},
																								"signatureId": map[string]interface{}{
																									"type": "integer",
																								},
																								"tag": map[string]interface{}{
																									"type": "string",
																								},
																							},
																							"type": "object",
																						},
																						"type": "array",
																					},
																					"staticValues": map[string]interface{}{
																						"type": "string",
																					},
																					"type": map[string]interface{}{
																						"enum": []interface{}{
																							"explicit",
																							"wildcard",
																						},
																						"type": "string",
																					},
																					"url": map[string]interface{}{
																						"properties": map[string]interface{}{
																							"method": map[string]interface{}{
																								"enum": []interface{}{
																									"ACL",
																									"BCOPY",
																									"BDELETE",
																									"BMOVE",
																									"BPROPFIND",
																									"BPROPPATCH",
																									"CHECKIN",
																									"CHECKOUT",
																									"CONNECT",
																									"COPY",
																									"DELETE",
																									"GET",
																									"HEAD",
																									"LINK",
																									"LOCK",
																									"MERGE",
																									"MKCOL",
																									"MKWORKSPACE",
																									"MOVE",
																									"NOTIFY",
																									"OPTIONS",
																									"PATCH",
																									"POLL",
																									"POST",
																									"PROPFIND",
																									"PROPPATCH",
																									"PUT",
																									"REPORT",
																									"RPC_IN_DATA",
																									"RPC_OUT_DATA",
																									"SEARCH",
																									"SUBSCRIBE",
																									"TRACE",
																									"TRACK",
																									"UNLINK",
																									"UNLOCK",
																									"UNSUBSCRIBE",
																									"VERSION_CONTROL",
																									"X-MS-ENUMATTS",
																									"*",
																								},
																								"type": "string",
																							},
																							"name": map[string]interface{}{
																								"type": "string",
																							},
																							"protocol": map[string]interface{}{
																								"enum": []interface{}{
																									"http",
																									"https",
																								},
																								"type": "string",
																							},
																							"type": map[string]interface{}{
																								"enum": []interface{}{
																									"explicit",
																									"wildcard",
																								},
																								"type": "string",
																							},
																						},
																						"type": "object",
																					},
																					"valueMetacharOverrides": map[string]interface{}{
																						"items": map[string]interface{}{
																							"properties": map[string]interface{}{
																								"isAllowed": map[string]interface{}{
																									"type": "boolean",
																								},
																								"metachar": map[string]interface{}{
																									"type": "string",
																								},
																							},
																							"type": "object",
																						},
																						"type": "array",
																					},
																					"valueType": map[string]interface{}{
																						"enum": []interface{}{
																							"array",
																							"auto-detect",
																							"dynamic-content",
																							"dynamic-parameter-name",
																							"ignore",
																							"json",
																							"object",
																							"openapi-array",
																							"static-content",
																							"user-input",
																							"xml",
																						},
																						"type": "string",
																					},
																					"wildcardOrder": map[string]interface{}{
																						"type": "integer",
																					},
																				},
																				"type": "object",
																			},
																			"urlSegmentIndex": map[string]interface{}{
																				"type": "integer",
																			},
																		},
																		"type": "object",
																	},
																	"type": "array",
																},
																"protocol": map[string]interface{}{
																	"enum": []interface{}{
																		"http",
																		"https",
																	},
																	"type": "string",
																},
																"signatureOverrides": map[string]interface{}{
																	"items": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"enabled": map[string]interface{}{
																				"type": "boolean",
																			},
																			"name": map[string]interface{}{
																				"type": "string",
																			},
																			"signatureId": map[string]interface{}{
																				"type": "integer",
																			},
																			"tag": map[string]interface{}{
																				"type": "string",
																			},
																		},
																		"type": "object",
																	},
																	"type": "array",
																},
																"type": map[string]interface{}{
																	"enum": []interface{}{
																		"explicit",
																		"wildcard",
																	},
																	"type": "string",
																},
																"urlContentProfiles": map[string]interface{}{
																	"items": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"contentProfile": map[string]interface{}{
																				"properties": map[string]interface{}{
																					"name": map[string]interface{}{
																						"type": "string",
																					},
																				},
																				"type": "object",
																			},
																			"headerName": map[string]interface{}{
																				"type": "string",
																			},
																			"headerOrder": map[string]interface{}{
																				"anyOf": []interface{}{
																					map[string]interface{}{
																						"type": "integer",
																					},
																					map[string]interface{}{
																						"type": "string",
																					},
																				},
																				"x-kubernetes-int-or-string": true,
																			},
																			"headerValue": map[string]interface{}{
																				"type": "string",
																			},
																			"name": map[string]interface{}{
																				"type": "string",
																			},
																			"type": map[string]interface{}{
																				"enum": []interface{}{
																					"apply-content-signatures",
																					"apply-value-and-content-signatures",
																					"disallow",
																					"do-nothing",
																					"form-data",
																					"gwt",
																					"json",
																					"xml",
																					"grpc",
																				},
																				"type": "string",
																			},
																		},
																		"type": "object",
																	},
																	"type": "array",
																},
																"wildcardOrder": map[string]interface{}{
																	"type": "integer",
																},
															},
															"type": "object",
														},
														"type": "array",
													},
													"whitelist-ips": map[string]interface{}{
														"items": map[string]interface{}{
															"properties": map[string]interface{}{
																"$action": map[string]interface{}{
																	"enum": []interface{}{
																		"delete",
																	},
																	"type": "string",
																},
																"blockRequests": map[string]interface{}{
																	"enum": []interface{}{
																		"always",
																		"never",
																		"policy-default",
																	},
																	"type": "string",
																},
																"ipAddress": map[string]interface{}{
																	"pattern": `[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}`,
																	"type":    "string",
																},
																"ipMask": map[string]interface{}{
																	"pattern": `[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}`,
																	"type":    "string",
																},
																"neverLogRequests": map[string]interface{}{
																	"type": "boolean",
																},
															},
															"type": "object",
														},
														"type": "array",
													},
													"whitelistIpReference": map[string]interface{}{
														"properties": map[string]interface{}{
															"link": map[string]interface{}{
																"pattern": "^http",
																"type":    "string",
															},
														},
														"type": "object",
													},
													"xml-profiles": map[string]interface{}{
														"items": map[string]interface{}{
															"properties": map[string]interface{}{
																"$action": map[string]interface{}{
																	"enum": []interface{}{
																		"delete",
																	},
																	"type": "string",
																},
																"attackSignaturesCheck": map[string]interface{}{
																	"type": "boolean",
																},
																"defenseAttributes": map[string]interface{}{
																	"properties": map[string]interface{}{
																		"allowCDATA": map[string]interface{}{
																			"type": "boolean",
																		},
																		"allowDTDs": map[string]interface{}{
																			"type": "boolean",
																		},
																		"allowExternalReferences": map[string]interface{}{
																			"type": "boolean",
																		},
																		"allowProcessingInstructions": map[string]interface{}{
																			"type": "boolean",
																		},
																		"maximumAttributeValueLength": map[string]interface{}{
																			"anyOf": []interface{}{
																				map[string]interface{}{
																					"type": "integer",
																				},
																				map[string]interface{}{
																					"type": "string",
																				},
																			},
																			"x-kubernetes-int-or-string": true,
																		},
																		"maximumAttributesPerElement": map[string]interface{}{
																			"anyOf": []interface{}{
																				map[string]interface{}{
																					"type": "integer",
																				},
																				map[string]interface{}{
																					"type": "string",
																				},
																			},
																			"x-kubernetes-int-or-string": true,
																		},
																		"maximumChildrenPerElement": map[string]interface{}{
																			"anyOf": []interface{}{
																				map[string]interface{}{
																					"type": "integer",
																				},
																				map[string]interface{}{
																					"type": "string",
																				},
																			},
																			"x-kubernetes-int-or-string": true,
																		},
																		"maximumDocumentDepth": map[string]interface{}{
																			"anyOf": []interface{}{
																				map[string]interface{}{
																					"type": "integer",
																				},
																				map[string]interface{}{
																					"type": "string",
																				},
																			},
																			"x-kubernetes-int-or-string": true,
																		},
																		"maximumDocumentSize": map[string]interface{}{
																			"anyOf": []interface{}{
																				map[string]interface{}{
																					"type": "integer",
																				},
																				map[string]interface{}{
																					"type": "string",
																				},
																			},
																			"x-kubernetes-int-or-string": true,
																		},
																		"maximumElements": map[string]interface{}{
																			"anyOf": []interface{}{
																				map[string]interface{}{
																					"type": "integer",
																				},
																				map[string]interface{}{
																					"type": "string",
																				},
																			},
																			"x-kubernetes-int-or-string": true,
																		},
																		"maximumNSDeclarations": map[string]interface{}{
																			"anyOf": []interface{}{
																				map[string]interface{}{
																					"type": "integer",
																				},
																				map[string]interface{}{
																					"type": "string",
																				},
																			},
																			"x-kubernetes-int-or-string": true,
																		},
																		"maximumNameLength": map[string]interface{}{
																			"anyOf": []interface{}{
																				map[string]interface{}{
																					"type": "integer",
																				},
																				map[string]interface{}{
																					"type": "string",
																				},
																			},
																			"x-kubernetes-int-or-string": true,
																		},
																		"maximumNamespaceLength": map[string]interface{}{
																			"anyOf": []interface{}{
																				map[string]interface{}{
																					"type": "integer",
																				},
																				map[string]interface{}{
																					"type": "string",
																				},
																			},
																			"x-kubernetes-int-or-string": true,
																		},
																		"tolerateCloseTagShorthand": map[string]interface{}{
																			"type": "boolean",
																		},
																		"tolerateLeadingWhiteSpace": map[string]interface{}{
																			"type": "boolean",
																		},
																		"tolerateNumericNames": map[string]interface{}{
																			"type": "boolean",
																		},
																	},
																	"type": "object",
																},
																"description": map[string]interface{}{
																	"type": "string",
																},
																"enableWss": map[string]interface{}{
																	"type": "boolean",
																},
																"followSchemaLinks": map[string]interface{}{
																	"type": "boolean",
																},
																"name": map[string]interface{}{
																	"type": "string",
																},
																"signatureOverrides": map[string]interface{}{
																	"items": map[string]interface{}{
																		"properties": map[string]interface{}{
																			"enabled": map[string]interface{}{
																				"type": "boolean",
																			},
																			"name": map[string]interface{}{
																				"type": "string",
																			},
																			"signatureId": map[string]interface{}{
																				"type": "integer",
																			},
																			"tag": map[string]interface{}{
																				"type": "string",
																			},
																		},
																		"type": "object",
																	},
																	"type": "array",
																},
																"useXmlResponsePage": map[string]interface{}{
																	"type": "boolean",
																},
															},
															"type": "object",
														},
														"type": "array",
													},
													"xml-validation-files": map[string]interface{}{
														"items": map[string]interface{}{
															"properties": map[string]interface{}{
																"$action": map[string]interface{}{
																	"enum": []interface{}{
																		"delete",
																	},
																	"type": "string",
																},
																"contents": map[string]interface{}{
																	"type": "string",
																},
																"fileName": map[string]interface{}{
																	"type": "string",
																},
																"isBase64": map[string]interface{}{
																	"type": "boolean",
																},
															},
															"type": "object",
														},
														"type": "array",
													},
													"xmlProfileReference": map[string]interface{}{
														"properties": map[string]interface{}{
															"link": map[string]interface{}{
																"pattern": "^http",
																"type":    "string",
															},
														},
														"type": "object",
													},
													"xmlValidationFileReference": map[string]interface{}{
														"properties": map[string]interface{}{
															"link": map[string]interface{}{
																"pattern": "^http",
																"type":    "string",
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
								},
								"type": "object",
							},
						},
						"served":  true,
						"storage": true,
					},
				},
			},
		},
	}

	return mutate.MutateCRDAppoliciesAppprotectF5Com(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;create;update;patch;delete

// CreateCRDPoliciesK8sNginxOrg creates the CustomResourceDefinition resource with name policies.k8s.nginx.org.
func CreateCRDPoliciesK8sNginxOrg(
	parent *platformv1alpha1.IngressComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if parent.Spec.Nginx.Include != true {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:field=nginx.include,value=true,include
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					"controller-gen.kubebuilder.io/version": "v0.21.0",
				},
				"name": "policies.k8s.nginx.org",
				"labels": map[string]interface{}{
					"platform.nukleros.io/category": "ingress",
					"platform.nukleros.io/project":  "nginx-ingress-controller",
				},
			},
			"spec": map[string]interface{}{
				"group": "k8s.nginx.org",
				"names": map[string]interface{}{
					"kind":     "Policy",
					"listKind": "PolicyList",
					"plural":   "policies",
					"shortNames": []interface{}{
						"pol",
					},
					"singular": "policy",
				},
				"scope": "Namespaced",
				"versions": []interface{}{
					map[string]interface{}{
						"additionalPrinterColumns": []interface{}{
							map[string]interface{}{
								"description": "Current state of the Policy. If the resource has a valid status, it means it has been validated and accepted by the Ingress Controller.",
								"jsonPath":    ".status.state",
								"name":        "State",
								"type":        "string",
							},
							map[string]interface{}{
								"jsonPath": ".metadata.creationTimestamp",
								"name":     "Age",
								"type":     "date",
							},
						},
						"name": "v1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"description": "Policy defines a Policy for VirtualServer and VirtualServerRoute resources.",
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
										"description": `PolicySpec is the spec of the Policy resource.
The spec includes multiple fields, where each field represents a different policy.
Only one policy (field) is allowed.`,
										"properties": map[string]interface{}{
											"accessControl": map[string]interface{}{
												"description": "The access control policy based on the client IP address.",
												"properties": map[string]interface{}{
													"allow": map[string]interface{}{
														"items": map[string]interface{}{
															"type": "string",
														},
														"type": "array",
													},
													"deny": map[string]interface{}{
														"items": map[string]interface{}{
															"type": "string",
														},
														"type": "array",
													},
												},
												"type": "object",
											},
											"apiKey": map[string]interface{}{
												"description": "The API Key policy configures NGINX to authorize requests which provide a valid API Key in a specified header or query param.",
												"properties": map[string]interface{}{
													"clientSecret": map[string]interface{}{
														"description": "The key to which the API key is applied. Can contain text, variables, or a combination of them. Accepted variables are $http_, $arg_, $cookie_.",
														"type":        "string",
													},
													"suppliedIn": map[string]interface{}{
														"description": "The location of the API Key. For example, $http_auth, $arg_apikey, $cookie_auth. Accepted variables are $http_, $arg_, $cookie_.",
														"properties": map[string]interface{}{
															"header": map[string]interface{}{
																"description": "The location of the API Key as a request header. For example, $http_auth. Accepted variables are $http_.",
																"items": map[string]interface{}{
																	"type": "string",
																},
																"type": "array",
															},
															"query": map[string]interface{}{
																"description": "The location of the API Key as a query param. For example, $arg_apikey. Accepted variables are $arg_.",
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
											"basicAuth": map[string]interface{}{
												"description": "The basic auth policy configures NGINX to authenticate client requests using HTTP Basic authentication credentials.",
												"properties": map[string]interface{}{
													"realm": map[string]interface{}{
														"description": "The realm for the basic authentication.",
														"type":        "string",
													},
													"secret": map[string]interface{}{
														"description": "The name of the Kubernetes secret that stores the Htpasswd configuration. It must be in the same namespace as the Policy resource. The secret must be of the type nginx.org/htpasswd, and the config must be stored in the secret under the key htpasswd, otherwise the secret will be rejected as invalid.",
														"type":        "string",
													},
												},
												"type": "object",
											},
											"cache": map[string]interface{}{
												"description": "The Cache Key defines a cache policy for proxy caching",
												"properties": map[string]interface{}{
													"allowedCodes": map[string]interface{}{
														"description": `AllowedCodes defines which HTTP response codes should be cached.
Accepts either:
- The string "any" to cache all response codes (must be the only element)
- A list of HTTP status codes as integers (100-599)
Examples: ["any"], [200, 301, 404], [200].
Invalid: ["any", 200] (cannot mix "any" with specific codes).`,
														"items": map[string]interface{}{
															"anyOf": []interface{}{
																map[string]interface{}{
																	"type": "integer",
																},
																map[string]interface{}{
																	"type": "string",
																},
															},
															"x-kubernetes-int-or-string": true,
														},
														"type": "array",
													},
													"allowedMethods": map[string]interface{}{
														"description": `AllowedMethods defines which HTTP methods should be cached.
Only "GET", "HEAD", and "POST" are supported by NGINX proxy_cache_methods directive.
GET and HEAD are always cached by default even if not specified.
Maximum of 3 items allowed. Examples: ["GET"], ["GET", "HEAD", "POST"].
Invalid methods: PUT, DELETE, PATCH, etc.`,
														"items": map[string]interface{}{
															"type": "string",
														},
														"maxItems": 3,
														"type":     "array",
														"x-kubernetes-validations": []interface{}{
															map[string]interface{}{
																"message": "allowed methods must be one of: GET, HEAD, POST",
																"rule":    "self.all(method, method in ['GET', 'HEAD', 'POST'])",
															},
														},
													},
													"cacheBackgroundUpdate": map[string]interface{}{
														"default": false,
														"description": `CacheBackgroundUpdate allows starting a background subrequest to update an expired cache item (proxy_cache_background_update).
A stale cached response is returned to the client while the cache is being updated.`,
														"type": "boolean",
													},
													"cacheKey": map[string]interface{}{
														"description": `CacheKey defines a key for caching (proxy_cache_key).
By default, close to "$scheme$proxy_host$uri$is_args$args".
Must not contain command execution patterns: $(, ` + "`" + `, ;, &&, ||`,
														"maxLength": 1024,
														"type":      "string",
														"x-kubernetes-validations": []interface{}{
															map[string]interface{}{
																"message": "cache key must not contain command execution patterns: $(, `, ;, &&, ||",
																"rule":    "!self.contains('$(') && !self.contains('`') && !self.contains(';') && !self.contains('&&') && !self.contains('||')",
															},
														},
													},
													"cacheMinUses": map[string]interface{}{
														"description": "CacheMinUses sets the number of requests after which the response will be cached (proxy_cache_min_uses).",
														"minimum":     1,
														"type":        "integer",
													},
													"cachePurgeAllow": map[string]interface{}{
														"description": `CachePurgeAllow defines IP addresses or CIDR blocks allowed to purge cache.
This feature is only available in NGINX Plus.
Examples: ["192.168.1.100", "10.0.0.0/8", "::1"].
Invalid in NGINX OSS (will be ignored).`,
														"items": map[string]interface{}{
															"type": "string",
														},
														"type": "array",
													},
													"cacheRevalidate": map[string]interface{}{
														"default": false,
														"description": `CacheRevalidate enables revalidation of expired cache items using conditional requests (proxy_cache_revalidate).
Uses "If-Modified-Since" and "If-None-Match" header fields.`,
														"type": "boolean",
													},
													"cacheUseStale": map[string]interface{}{
														"description": `CacheUseStale determines in which cases a stale cached response can be used (proxy_cache_use_stale).
Valid parameters: error, timeout, invalid_header, updating, http_500, http_502, http_503, http_504, http_403, http_404, http_429, off.`,
														"items": map[string]interface{}{
															"type": "string",
														},
														"maxItems": 11,
														"type":     "array",
													},
													"cacheZoneName": map[string]interface{}{
														"description": `CacheZoneName defines the name of the cache zone. Must start with a lowercase letter,
followed by alphanumeric characters or underscores, and end with an alphanumeric character.
Single lowercase letters are also allowed. Examples: "cache", "my_cache", "cache1".`,
														"pattern": "^[a-z][a-zA-Z0-9_]*[a-zA-Z0-9]$|^[a-z]$",
														"type":    "string",
													},
													"cacheZoneSize": map[string]interface{}{
														"description": `CacheZoneSize defines the size of the cache zone. Must be a number followed by a size unit:
'k' or 'K' for kilobytes, 'm' or 'M' for megabytes, or 'g' or 'G' for gigabytes.
Examples: "10m", "1g", "512k".`,
														"pattern": "^[0-9]+[kmgKMG]$",
														"type":    "string",
													},
													"conditions": map[string]interface{}{
														"description": "Conditions defines when responses should not be cached or taken from cache.",
														"properties": map[string]interface{}{
															"bypass": map[string]interface{}{
																"description": `Bypass defines conditions under which the response will not be taken from a cache (proxy_cache_bypass).
If at least one value of the string parameters is not empty and is not equal to "0" then the response will not be taken from the cache.`,
																"items": map[string]interface{}{
																	"type": "string",
																},
																"type": "array",
															},
															"noCache": map[string]interface{}{
																"description": `NoCache defines conditions under which the response will not be saved to a cache (proxy_no_cache).
If at least one value of the string parameters is not empty and is not equal to "0" then the response will not be saved.`,
																"items": map[string]interface{}{
																	"type": "string",
																},
																"type": "array",
															},
														},
														"type": "object",
													},
													"inactive": map[string]interface{}{
														"description": `Inactive sets the time after which cached data that are not accessed get removed from the cache (inactive parameter).
By default, inactive is set to 10 minutes.`,
														"pattern": "^[0-9]+[smhd]$",
														"type":    "string",
													},
													"levels": map[string]interface{}{
														"description": `Levels defines the cache directory hierarchy levels for storing cached files.
Must be in format "X:Y" or "X:Y:Z" where X, Y, Z are either 1 or 2.
This controls the number of subdirectory levels and their name lengths.
Examples: "1:2", "2:2", "1:2:2".
Invalid: "3:1", "1:3", "1:2:3".`,
														"pattern": "^[12](?::[12]){0,2}$",
														"type":    "string",
													},
													"lock": map[string]interface{}{
														"description": "Lock configures cache locking to prevent multiple identical requests from populating the same cache element simultaneously.",
														"properties": map[string]interface{}{
															"age": map[string]interface{}{
																"description": `Age sets the maximum time a cache lock can be held (proxy_cache_lock_age).
If the last request passed to the proxied server for populating a new cache element has not completed for the specified time, one more request may be passed.`,
																"pattern": "^[0-9]+[smhd]$",
																"type":    "string",
															},
															"enable": map[string]interface{}{
																"default": false,
																"description": `Enable sets whether cache locking is enabled (proxy_cache_lock).
When enabled, only one request at a time will be allowed to populate a new cache element according to the proxy_cache_key.`,
																"type": "boolean",
															},
															"timeout": map[string]interface{}{
																"description": `Timeout sets a timeout for proxy_cache_lock.
When the time expires, the request will be passed to the proxied server, however, the response will not be cached.`,
																"pattern": "^[0-9]+[smhd]$",
																"type":    "string",
															},
														},
														"type": "object",
														"x-kubernetes-validations": []interface{}{
															map[string]interface{}{
																"message": "timeout or age require enable=true",
																"rule":    "(!has(self.timeout) && !has(self.age)) || self.enable",
															},
														},
													},
													"manager": map[string]interface{}{
														"description": "Manager configures the cache manager process parameters (manager_files, manager_sleep, manager_threshold).",
														"properties": map[string]interface{}{
															"files": map[string]interface{}{
																"description": `Files sets the maximum number of files that will be deleted in one iteration by the cache manager.
During one iteration no more than manager_files items are deleted (by default, 100).`,
																"minimum": 1,
																"type":    "integer",
															},
															"sleep": map[string]interface{}{
																"description": `Sleep sets the pause between cache manager iterations.
Between iterations, a pause configured by manager_sleep (by default, 50 milliseconds) is made.`,
																"pattern": "^[0-9]+[mu]?s$",
																"type":    "string",
															},
															"threshold": map[string]interface{}{
																"description": `Threshold sets the maximum duration of one cache manager iteration.
The duration of one iteration is limited by manager_threshold (by default, 200 milliseconds).`,
																"pattern": "^[0-9]+[mu]?s$",
																"type":    "string",
															},
														},
														"type": "object",
													},
													"maxSize": map[string]interface{}{
														"description": `MaxSize sets the maximum cache size (max_size parameter).
When the size is exceeded, the cache manager removes the least recently used data.`,
														"pattern": "^[0-9]+[kmgKMG]$",
														"type":    "string",
													},
													"minFree": map[string]interface{}{
														"description": `MinFree sets the minimum amount of free space required on the file system with cache (min_free parameter).
When there is not enough free space, the cache manager removes the least recently used data.`,
														"pattern": "^[0-9]+[kmgKMG]$",
														"type":    "string",
													},
													"overrideUpstreamCache": map[string]interface{}{
														"default": false,
														"description": `OverrideUpstreamCache controls whether to override upstream cache headers
(using proxy_ignore_headers directive). When true, NGINX will ignore
cache-related headers from upstream servers like Cache-Control, Expires, etc.
Default: false.`,
														"type": "boolean",
													},
													"time": map[string]interface{}{
														"description": `Time defines the default cache time. Required when allowedCodes is specified.
Must be a number followed by a time unit:
's' for seconds, 'm' for minutes, 'h' for hours, 'd' for days.
Examples: "30s", "5m", "1h", "2d".`,
														"pattern": "^[0-9]+[smhd]$",
														"type":    "string",
													},
													"useTempPath": map[string]interface{}{
														"default": false,
														"description": `UseTempPath controls whether temporary files and the cache are put on different file systems (use_temp_path parameter).
If set to false, temporary files will be put directly in the cache directory (use_temp_path=off).
Default: false (use_temp_path=off, which puts temp files directly in cache directory for better performance).`,
														"type": "boolean",
													},
												},
												"required": []interface{}{
													"cacheZoneName",
													"cacheZoneSize",
												},
												"type": "object",
												"x-kubernetes-validations": []interface{}{
													map[string]interface{}{
														"message": "time is required when allowedCodes is specified",
														"rule":    "!has(self.allowedCodes) || (has(self.allowedCodes) && has(self.time))",
													},
												},
											},
											"cors": map[string]interface{}{
												"description": "The CORS policy configures Cross-Origin Resource Sharing headers",
												"properties": map[string]interface{}{
													"allowCredentials": map[string]interface{}{
														"default": false,
														"description": `AllowCredentials indicates whether the response to the request can be exposed when the credentials flag is true.
When used as part of a response to a preflight request, this indicates whether the actual request can be made using credentials.`,
														"type": "boolean",
													},
													"allowHeaders": map[string]interface{}{
														"description": `AllowHeaders defines the headers that are allowed in cross-origin requests.
Common safe headers: ["Accept", "Accept-Language", "Content-Language", "Content-Type"]
Custom headers: ["Authorization", "X-Requested-With", "X-Custom-Header"]`,
														"items": map[string]interface{}{
															"type": "string",
														},
														"type": "array",
														"x-kubernetes-validations": []interface{}{
															map[string]interface{}{
																"message": "header name cannot be empty",
																"rule":    "self.all(header, header != '')",
															},
														},
													},
													"allowMethods": map[string]interface{}{
														"description": "AllowMethods defines the HTTP methods that are allowed for cross-origin requests.",
														"items": map[string]interface{}{
															"type": "string",
														},
														"type": "array",
														"x-kubernetes-validations": []interface{}{
															map[string]interface{}{
																"message": "method name cannot be empty",
																"rule":    "self.all(method, method != '')",
															},
														},
													},
													"allowOrigin": map[string]interface{}{
														"description": `AllowOrigin defines the origins that are allowed to make cross-origin requests.
Can be exact domains, single wildcards, or "*" for all origins.
Examples: ["https://example.com", "https://*.mydomain.com", "*"]
Security: When allowCredentials is true, wildcard "*" is not allowed per CORS specification.
The server must specify explicit origins for credentialed requests.`,
														"items": map[string]interface{}{
															"type": "string",
														},
														"minItems": 1,
														"type":     "array",
														"x-kubernetes-validations": []interface{}{
															map[string]interface{}{
																"message": "origin cannot be empty",
																"rule":    "self.all(origin, origin != '')",
															},
														},
													},
													"exposeHeaders": map[string]interface{}{
														"description": `ExposeHeaders defines the headers that browsers are allowed to access.
Use this field to expose additional custom headers to the browser.
Example: ["X-Total-Count", "X-Page-Size", "X-RateLimit-Remaining"]
Note: Set-Cookie headers cannot be exposed via CORS per official MDN specification.`,
														"items": map[string]interface{}{
															"type": "string",
														},
														"type": "array",
														"x-kubernetes-validations": []interface{}{
															map[string]interface{}{
																"message": "header name cannot be empty",
																"rule":    "self.all(header, header != '')",
															},
														},
													},
													"maxAge": map[string]interface{}{
														"default": 86400,
														"description": `MaxAge defines how long (in seconds) the results of a preflight request can be cached.
Default: 86400 (24 hours). Maximum recommended value is 86400 (24 hours).`,
														"minimum": 0,
														"type":    "integer",
													},
												},
												"required": []interface{}{
													"allowOrigin",
												},
												"type": "object",
												"x-kubernetes-validations": []interface{}{
													map[string]interface{}{
														"message": "cannot use wildcard '*' for allowOrigin when allowCredentials is true for security reasons",
														"rule":    "!(self.allowOrigin.exists(origin, origin == '*') && has(self.allowCredentials) && self.allowCredentials == true)",
													},
												},
											},
											"egressMTLS": map[string]interface{}{
												"description": "The EgressMTLS policy configures upstreams authentication and certificate verification.",
												"properties": map[string]interface{}{
													"ciphers": map[string]interface{}{
														"description": "Specifies the enabled ciphers for requests to an upstream HTTPS server. The default is DEFAULT.",
														"type":        "string",
													},
													"protocols": map[string]interface{}{
														"description": "Specifies the protocols for requests to an upstream HTTPS server. The default is TLSv1 TLSv1.1 TLSv1.2.",
														"type":        "string",
													},
													"serverName": map[string]interface{}{
														"description": "Enables passing of the server name through Server Name Indication extension.",
														"type":        "boolean",
													},
													"sessionReuse": map[string]interface{}{
														"description": "Enables reuse of SSL sessions to the upstreams. The default is true.",
														"type":        "boolean",
													},
													"sslName": map[string]interface{}{
														"description": "Allows overriding the server name used to verify the certificate of the upstream HTTPS server.",
														"type":        "string",
													},
													"tlsSecret": map[string]interface{}{
														"description": "The name of the Kubernetes secret that stores the TLS certificate and key. It must be in the same namespace as the Policy resource. The secret must be of the type kubernetes.io/tls, the certificate must be stored in the secret under the key tls.crt, and the key must be stored under the key tls.key, otherwise the secret will be rejected as invalid.",
														"type":        "string",
													},
													"trustedCertSecret": map[string]interface{}{
														"description": "The name of the Kubernetes secret that stores the CA certificate. It must be in the same namespace as the Policy resource. The secret must be of the type nginx.org/ca, and the certificate must be stored in the secret under the key ca.crt, otherwise the secret will be rejected as invalid.",
														"type":        "string",
													},
													"verifyDepth": map[string]interface{}{
														"description": "Sets the verification depth in the proxied HTTPS server certificates chain. The default is 1.",
														"type":        "integer",
													},
													"verifyServer": map[string]interface{}{
														"description": "Enables verification of the upstream HTTPS server certificate.",
														"type":        "boolean",
													},
												},
												"type": "object",
											},
											"externalAuth": map[string]interface{}{
												"description": "The ExternalAuth policy configures NGINX to authenticate client requests using an external authentication server, which can be used for example with the oauth2-proxy or any custom authentication server.",
												"properties": map[string]interface{}{
													"authServiceName": map[string]interface{}{
														"description": "AuthServiceName is the name of the Kubernetes service to which the request will be sent for authentication.  It can be in the same namespace as the Policy resource or in a different namespace. If the service is in a different namespace, it should be specified in the format <namespace>/<service>. For example, auth-service or auth-namespace/auth-service.",
														"pattern":     `^([a-z0-9]([-a-z0-9]*[a-z0-9])?\/)?[a-z0-9]([-a-z0-9]*[a-z0-9])?$`,
														"type":        "string",
													},
													"authServicePorts": map[string]interface{}{
														"description": "AuthServicePorts are the ports of the Kubernetes service to which requests will be sent for authentication. If not specified, the ports will be looked up from the service definition. This field is only required if the user wants to choose a specific port from the service definition, otherwise the first port will be used by default.",
														"items": map[string]interface{}{
															"type": "integer",
														},
														"type": "array",
													},
													"authSigninRedirectBasePath": map[string]interface{}{
														"description": "AuthSigninRedirectBasePath is the base path for the NGINX location block that handles sign-in redirect requests from the external authentication server. For example, oauth2-proxy expects /oauth2. If not specified, defaults to /oauth2.",
														"pattern":     `^/[a-zA-Z0-9._~:/?#\[\]@!$&'()*+,;=-]*$`,
														"type":        "string",
													},
													"authSigninURI": map[string]interface{}{
														"description": "AuthSigninURI is the URI which requests will be redirected to if the external authentication server determines that the client needs to be authenticated. This is typically used when the external authentication server is an oauth2-proxy or any custom authentication server that requires redirection for authentication. The URI is a relative URI, for example /signin.",
														"pattern":     "^/.*$",
														"type":        "string",
													},
													"authSnippets": map[string]interface{}{
														"description": "AuthSnippets can be used to add custom configuration snippets to the location block of the external authentication configuration. This can be used for example to add additional headers to the request sent to the external authentication server, or to configure additional parameters for the auth_request module. The content of this field will be added as-is to the location block, so it must be a valid NGINX configuration snippet.",
														"type":        "string",
													},
													"authURI": map[string]interface{}{
														"default":     "/",
														"description": "AuthURI is the URI of the external authentication server to which the request will be sent for authentication. The URI is a relative URI, for example /auth.",
														"pattern":     "^/.*$",
														"type":        "string",
													},
													"sniName": map[string]interface{}{
														"description": "SNIName sets the server name used for SNI and certificate verification when connecting to the external authentication server over TLS. If not specified, defaults to <service-name>.<namespace>.svc derived from authServiceName.",
														"pattern":     `^[a-zA-Z0-9]([-a-zA-Z0-9]*[a-zA-Z0-9])?(\.[a-zA-Z0-9]([-a-zA-Z0-9]*[a-zA-Z0-9])?)*$`,
														"type":        "string",
													},
													"sslEnabled": map[string]interface{}{
														"default":     false,
														"description": "SSLEnabled enables HTTPS when proxying requests to the external authentication server. Default is false.",
														"type":        "boolean",
													},
													"sslVerify": map[string]interface{}{
														"default":     false,
														"description": "SSLVerify enables verification of the external authentication server's SSL certificate. Default is false.",
														"type":        "boolean",
													},
													"sslVerifyDepth": map[string]interface{}{
														"default":     1,
														"description": "SSLVerifyDepth sets the verification depth in the external authentication server certificates chain. Default is 1.",
														"minimum":     0,
														"type":        "integer",
													},
													"trustedCertSecret": map[string]interface{}{
														"description": "TrustedCertSecret is the name of the Kubernetes secret that stores the CA certificate for external authentication server certificate verification. It can be in the same namespace as the Policy resource or in a different namespace specified as <namespace>/<secret>. The secret must be of the type nginx.org/ca, and the certificate must be stored under the key ca.crt.",
														"pattern":     `^([a-z0-9]([-a-z0-9]*[a-z0-9])?\/)?[a-z0-9]([-a-z0-9]*[a-z0-9])?$`,
														"type":        "string",
													},
												},
												"required": []interface{}{
													"authServiceName",
													"authURI",
												},
												"type": "object",
											},
											"ingressClassName": map[string]interface{}{
												"description": "Specifies which instance of NGINX Ingress Controller must handle the Policy resource.",
												"type":        "string",
											},
											"ingressMTLS": map[string]interface{}{
												"description": "The IngressMTLS policy configures client certificate verification.",
												"properties": map[string]interface{}{
													"clientCertSecret": map[string]interface{}{
														"description": "The name of the Kubernetes secret that stores the CA certificate. It must be in the same namespace as the Policy resource. The secret must be of the type nginx.org/ca, and the certificate must be stored in the secret under the key ca.crt, otherwise the secret will be rejected as invalid.",
														"type":        "string",
													},
													"crlFileName": map[string]interface{}{
														"description": "The file name of the Certificate Revocation List. NGINX Ingress Controller will look for this file in /etc/nginx/secrets",
														"type":        "string",
													},
													"verifyClient": map[string]interface{}{
														"description": "Verification for the client. Possible values are \"on\", \"off\", \"optional\", \"optional_no_ca\". The default is \"on\".",
														"type":        "string",
													},
													"verifyDepth": map[string]interface{}{
														"description": "Sets the verification depth in the client certificates chain. The default is 1.",
														"type":        "integer",
													},
												},
												"type": "object",
											},
											"jwt": map[string]interface{}{
												"description": "The JWT policy configures NGINX Plus to authenticate client requests using JSON Web Tokens.",
												"properties": map[string]interface{}{
													"jwksURI": map[string]interface{}{
														"description": "The remote URI where the request will be sent to retrieve JSON Web Key set",
														"type":        "string",
													},
													"keyCache": map[string]interface{}{
														"description": "Enables in-memory caching of JWKS (JSON Web Key Sets) that are obtained from the jwksURI and sets a valid time for expiration.",
														"type":        "string",
													},
													"realm": map[string]interface{}{
														"description": "The realm of the JWT.",
														"type":        "string",
													},
													"secret": map[string]interface{}{
														"description": "The name of the Kubernetes secret that stores the Htpasswd configuration. It must be in the same namespace as the Policy resource. The secret must be of the type nginx.org/htpasswd, and the config must be stored in the secret under the key htpasswd, otherwise the secret will be rejected as invalid.",
														"type":        "string",
													},
													"sniEnabled": map[string]interface{}{
														"description": "Enables SNI (Server Name Indication) for the JWT policy. This is useful when the remote server requires SNI to serve the correct certificate.",
														"type":        "boolean",
													},
													"sniName": map[string]interface{}{
														"description": "The SNI name to use when connecting to the remote server. If not set, the hostname from the ``jwksURI`` will be used.",
														"type":        "string",
													},
													"sslVerify": map[string]interface{}{
														"default":     false,
														"description": "Enables verification of the JWKS server SSL certificate. Default is false.",
														"type":        "boolean",
													},
													"sslVerifyDepth": map[string]interface{}{
														"default":     1,
														"description": "Sets the verification depth in the JWKS server certificates chain. The default is 1.",
														"minimum":     0,
														"type":        "integer",
													},
													"token": map[string]interface{}{
														"description": "The token specifies a variable that contains the JSON Web Token. By default the JWT is passed in the Authorization header as a Bearer Token. JWT may be also passed as a cookie or a part of a query string, for example: $cookie_auth_token. Accepted variables are $http_, $arg_, $cookie_.",
														"type":        "string",
													},
													"trustedCertSecret": map[string]interface{}{
														"description": "The name of the Kubernetes secret that stores the CA certificate for JWKS server verification. It must be in the same namespace as the Policy resource. The secret must be of the type nginx.org/ca, and the certificate must be stored in the secret under the key ca.crt.",
														"pattern":     "^[a-z0-9]([-a-z0-9]*[a-z0-9])?$",
														"type":        "string",
													},
												},
												"type": "object",
											},
											"oidc": map[string]interface{}{
												"description": "The OpenID Connect policy configures NGINX to authenticate client requests by validating a JWT token against an OAuth2/OIDC token provider, such as Auth0 or Keycloak.",
												"properties": map[string]interface{}{
													"accessTokenEnable": map[string]interface{}{
														"description": "Option of whether Bearer token is used to authorize NGINX to access protected backend.",
														"type":        "boolean",
													},
													"authEndpoint": map[string]interface{}{
														"description": "URL for the authorization endpoint provided by your OpenID Connect provider.",
														"type":        "string",
													},
													"authExtraArgs": map[string]interface{}{
														"description": "A list of extra URL arguments to pass to the authorization endpoint provided by your OpenID Connect provider. Arguments must be URL encoded, multiple arguments may be included in the list, for example [ arg1=value1, arg2=value2 ]",
														"items": map[string]interface{}{
															"type": "string",
														},
														"type": "array",
													},
													"clientID": map[string]interface{}{
														"description": "The client ID provided by your OpenID Connect provider.",
														"type":        "string",
													},
													"clientSecret": map[string]interface{}{
														"description": "The name of the Kubernetes secret that stores the client secret provided by your OpenID Connect provider. It must be in the same namespace as the Policy resource. The secret must be of the type nginx.org/oidc, and the secret under the key client-secret, otherwise the secret will be rejected as invalid. If PKCE is enabled, this should be not configured.",
														"type":        "string",
													},
													"endSessionEndpoint": map[string]interface{}{
														"description": "URL provided by your OpenID Connect provider to request the end user be logged out.",
														"type":        "string",
													},
													"jwksURI": map[string]interface{}{
														"description": "URL for the JSON Web Key Set (JWK) document provided by your OpenID Connect provider.",
														"type":        "string",
													},
													"pkceEnable": map[string]interface{}{
														"description": "Switches Proof Key for Code Exchange on. The OpenID client needs to be in public mode. clientSecret is not used in this mode.",
														"type":        "boolean",
													},
													"postLogoutRedirectURI": map[string]interface{}{
														"description": "URI to redirect to after the logout has been performed. Requires endSessionEndpoint. The default is /_logout.",
														"type":        "string",
													},
													"redirectURI": map[string]interface{}{
														"description": "Allows overriding the default redirect URI. The default is /_codexch.",
														"type":        "string",
													},
													"scope": map[string]interface{}{
														"description": "List of OpenID Connect scopes. The scope openid always needs to be present and others can be added concatenating them with a + sign, for example openid+profile+email, openid+email+userDefinedScope. The default is openid.",
														"type":        "string",
													},
													"sslVerify": map[string]interface{}{
														"default":     false,
														"description": "Enables verification of the IDP server SSL certificate. Default is false.",
														"type":        "boolean",
													},
													"sslVerifyDepth": map[string]interface{}{
														"default":     1,
														"description": "Sets the verification depth in the IDP server certificates chain. The default is 1.",
														"minimum":     0,
														"type":        "integer",
													},
													"tokenEndpoint": map[string]interface{}{
														"description": "URL for the token endpoint provided by your OpenID Connect provider.",
														"type":        "string",
													},
													"trustedCertSecret": map[string]interface{}{
														"description": "The name of the Kubernetes secret that stores the CA certificate for IDP server verification. It must be in the same namespace as the Policy resource. The secret must be of the type nginx.org/ca, and the certificate must be stored in the secret under the key ca.crt.",
														"pattern":     "^[a-z0-9]([-a-z0-9]*[a-z0-9])?$",
														"type":        "string",
													},
													"zoneSyncLeeway": map[string]interface{}{
														"description": "Specifies the maximum timeout in milliseconds for synchronizing ID/access tokens and shared values between Ingress Controller pods. The default is 200.",
														"type":        "integer",
													},
												},
												"type": "object",
												"x-kubernetes-validations": []interface{}{
													map[string]interface{}{
														"message": "trustedCertSecret can be set only if sslVerify is true",
														"rule":    "(self.sslVerify == true) || (self.sslVerify == false && !has(self.trustedCertSecret))",
													},
												},
											},
											"rateLimit": map[string]interface{}{
												"description": "The rate limit policy controls the rate of processing requests per a defined key.",
												"properties": map[string]interface{}{
													"burst": map[string]interface{}{
														"description": "Excessive requests are delayed until their number exceeds the burst size, in which case the request is terminated with an error.",
														"type":        "integer",
													},
													"condition": map[string]interface{}{
														"description": "Add a condition to a rate-limit policy.",
														"properties": map[string]interface{}{
															"default": map[string]interface{}{
																"description": "sets the rate limit in this policy to be the default if no conditions are met. In a group of policies with the same condition, only one policy can be the default.",
																"type":        "boolean",
															},
															"jwt": map[string]interface{}{
																"description": "defines a JWT condition to rate limit against.",
																"properties": map[string]interface{}{
																	"claim": map[string]interface{}{
																		"description": "the JWT claim to be rate limit by. Nested claims should be separated by \".\"",
																		"pattern":     `^([^$\s"'])*$`,
																		"type":        "string",
																	},
																	"match": map[string]interface{}{
																		"description": "the value of the claim to match against.",
																		"pattern":     `^([^$\s."'])*$`,
																		"type":        "string",
																	},
																},
																"required": []interface{}{
																	"claim",
																	"match",
																},
																"type": "object",
															},
															"variables": map[string]interface{}{
																"description": "defines a Variables condition to rate limit against.",
																"items": map[string]interface{}{
																	"description": "VariableCondition defines a condition to rate limit by a variable.",
																	"properties": map[string]interface{}{
																		"match": map[string]interface{}{
																			"description": "the value of the variable to match against.",
																			"pattern":     `^([^\s"'])*$`,
																			"type":        "string",
																		},
																		"name": map[string]interface{}{
																			"description": "the name of the variable to match against.",
																			"pattern":     `^([^\s"'])*$`,
																			"type":        "string",
																		},
																	},
																	"required": []interface{}{
																		"match",
																		"name",
																	},
																	"type": "object",
																},
																"maxItems": 1,
																"type":     "array",
															},
														},
														"type": "object",
													},
													"delay": map[string]interface{}{
														"description": "The delay parameter specifies a limit at which excessive requests become delayed. If not set all excessive requests are delayed.",
														"type":        "integer",
													},
													"dryRun": map[string]interface{}{
														"description": "Enables the dry run mode. In this mode, the rate limit is not actually applied, but the number of excessive requests is accounted as usual in the shared memory zone.",
														"type":        "boolean",
													},
													"key": map[string]interface{}{
														"description": `The key to which the rate limit is applied. Can contain text, variables, or a combination of them.
Variables must be surrounded by ${}. For example: ${binary_remote_addr}. Accepted variables are
$binary_remote_addr, $request_uri, $request_method, $url, $http_, $args, $arg_, $cookie_,$jwt_claim_ .`,
														"type": "string",
													},
													"logLevel": map[string]interface{}{
														"description": "Sets the desired logging level for cases when the server refuses to process requests due to rate exceeding, or delays request processing. Allowed values are info, notice, warn or error. Default is error.",
														"type":        "string",
													},
													"noDelay": map[string]interface{}{
														"description": "Disables the delaying of excessive requests while requests are being limited. Overrides delay if both are set.",
														"type":        "boolean",
													},
													"rate": map[string]interface{}{
														"description": "The rate of requests permitted. The rate is specified in requests per second (r/s) or requests per minute (r/m).",
														"type":        "string",
													},
													"rejectCode": map[string]interface{}{
														"description": "Sets the status code to return in response to rejected requests. Must fall into the range 400..599. Default is 503.",
														"type":        "integer",
													},
													"scale": map[string]interface{}{
														"description": "Enables a constant rate-limit by dividing the configured rate by the number of nginx-ingress pods currently serving traffic. This adjustment ensures that the rate-limit remains consistent, even as the number of nginx-pods fluctuates due to autoscaling. This will not work properly if requests from a client are not evenly distributed across all ingress pods (Such as with sticky sessions, long lived TCP Connections with many requests, and so forth). In such cases using zone-sync instead would give better results. Enabling zone-sync will suppress this setting.",
														"type":        "boolean",
													},
													"zoneSize": map[string]interface{}{
														"description": "Size of the shared memory zone. Only positive values are allowed. Allowed suffixes are k or m, if none are present k is assumed.",
														"type":        "string",
													},
												},
												"type": "object",
											},
											"waf": map[string]interface{}{
												"description": "The WAF policy configures WAF and log configuration policies for NGINX AppProtect",
												"properties": map[string]interface{}{
													"apBundle": map[string]interface{}{
														"description": "The App Protect WAF policy bundle. Mutually exclusive with apPolicy.",
														"type":        "string",
													},
													"apPolicy": map[string]interface{}{
														"description": "The App Protect WAF policy of the WAF. Accepts an optional namespace. Mutually exclusive with apBundle.",
														"type":        "string",
													},
													"enable": map[string]interface{}{
														"description": "Enables NGINX App Protect WAF.",
														"type":        "boolean",
													},
													"securityLog": map[string]interface{}{
														"description": "SecurityLog defines the security log of a WAF policy.",
														"properties": map[string]interface{}{
															"apLogBundle": map[string]interface{}{
																"description": "The App Protect WAF log bundle resource. Only works with apBundle.",
																"type":        "string",
															},
															"apLogConf": map[string]interface{}{
																"description": "The App Protect WAF log conf resource. Accepts an optional namespace. Only works with apPolicy.",
																"type":        "string",
															},
															"enable": map[string]interface{}{
																"description": "Enables security log.",
																"type":        "boolean",
															},
															"logDest": map[string]interface{}{
																"description": "The log destination for the security log. Only accepted variables are syslog:server=<ip-address>; localhost; fqdn>:<port>, stderr, <absolute path to file>.",
																"type":        "string",
															},
														},
														"type": "object",
													},
													"securityLogs": map[string]interface{}{
														"items": map[string]interface{}{
															"description": "SecurityLog defines the security log of a WAF policy.",
															"properties": map[string]interface{}{
																"apLogBundle": map[string]interface{}{
																	"description": "The App Protect WAF log bundle resource. Only works with apBundle.",
																	"type":        "string",
																},
																"apLogConf": map[string]interface{}{
																	"description": "The App Protect WAF log conf resource. Accepts an optional namespace. Only works with apPolicy.",
																	"type":        "string",
																},
																"enable": map[string]interface{}{
																	"description": "Enables security log.",
																	"type":        "boolean",
																},
																"logDest": map[string]interface{}{
																	"description": "The log destination for the security log. Only accepted variables are syslog:server=<ip-address>; localhost; fqdn>:<port>, stderr, <absolute path to file>.",
																	"type":        "string",
																},
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
									"status": map[string]interface{}{
										"description": "the status of the Policy resource",
										"properties": map[string]interface{}{
											"message": map[string]interface{}{
												"description": "The message of the current state of the resource. It can contain more detailed information about the reason.",
												"type":        "string",
											},
											"reason": map[string]interface{}{
												"description": "The reason of the current state of the resource.",
												"type":        "string",
											},
											"state": map[string]interface{}{
												"description": "Represents the current state of the resource. There are three possible values: Valid, Invalid and Warning. Valid indicates that the resource has been validated and accepted by the Ingress Controller. Invalid means the resource failed validation or",
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
			},
		},
	}

	return mutate.MutateCRDPoliciesK8sNginxOrg(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;create;update;patch;delete

// CreateCRDApusersigsAppprotectF5Com creates the CustomResourceDefinition resource with name apusersigs.appprotect.f5.com.
func CreateCRDApusersigsAppprotectF5Com(
	parent *platformv1alpha1.IngressComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if parent.Spec.Nginx.Include != true {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:field=nginx.include,value=true,include
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					"controller-gen.kubebuilder.io/version": "v0.13.0",
				},
				"name": "apusersigs.appprotect.f5.com",
				"labels": map[string]interface{}{
					"platform.nukleros.io/category": "ingress",
					"platform.nukleros.io/project":  "nginx-ingress-controller",
				},
			},
			"spec": map[string]interface{}{
				"group": "appprotect.f5.com",
				"names": map[string]interface{}{
					"kind":     "APUserSig",
					"listKind": "APUserSigList",
					"plural":   "apusersigs",
					"singular": "apusersig",
				},
				"scope": "Namespaced",
				"versions": []interface{}{
					map[string]interface{}{
						"name": "v1beta1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"description": "APUserSig is the Schema for the apusersigs API",
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
										"description": "APUserSigSpec defines the desired state of APUserSig",
										"properties": map[string]interface{}{
											"properties": map[string]interface{}{
												"type": "string",
											},
											"signatures": map[string]interface{}{
												"items": map[string]interface{}{
													"properties": map[string]interface{}{
														"accuracy": map[string]interface{}{
															"enum": []interface{}{
																"high",
																"medium",
																"low",
															},
															"type": "string",
														},
														"attackType": map[string]interface{}{
															"properties": map[string]interface{}{
																"name": map[string]interface{}{
																	"type": "string",
																},
															},
															"type": "object",
														},
														"description": map[string]interface{}{
															"type": "string",
														},
														"name": map[string]interface{}{
															"type": "string",
														},
														"references": map[string]interface{}{
															"properties": map[string]interface{}{
																"type": map[string]interface{}{
																	"enum": []interface{}{
																		"bugtraq",
																		"cve",
																		"nessus",
																		"url",
																	},
																	"type": "string",
																},
																"value": map[string]interface{}{
																	"type": "string",
																},
															},
															"type": "object",
														},
														"risk": map[string]interface{}{
															"enum": []interface{}{
																"high",
																"medium",
																"low",
															},
															"type": "string",
														},
														"rule": map[string]interface{}{
															"type": "string",
														},
														"signatureType": map[string]interface{}{
															"enum": []interface{}{
																"request",
																"response",
															},
															"type": "string",
														},
														"systems": map[string]interface{}{
															"items": map[string]interface{}{
																"properties": map[string]interface{}{
																	"name": map[string]interface{}{
																		"type": "string",
																	},
																},
																"type": "object",
															},
															"type": "array",
														},
													},
													"type": "object",
												},
												"type": "array",
											},
											"softwareVersion": map[string]interface{}{
												"type": "string",
											},
											"tag": map[string]interface{}{
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
					},
				},
			},
		},
	}

	return mutate.MutateCRDApusersigsAppprotectF5Com(resourceObj, parent, collection, reconciler, req)
}

// +kubebuilder:rbac:groups=apiextensions.k8s.io,resources=customresourcedefinitions,verbs=get;list;watch;create;update;patch;delete

// CreateCRDVirtualserversK8sNginxOrg creates the CustomResourceDefinition resource with name virtualservers.k8s.nginx.org.
func CreateCRDVirtualserversK8sNginxOrg(
	parent *platformv1alpha1.IngressComponent,
	collection *setupv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	if parent.Spec.Nginx.Include != true {
		return []client.Object{}, nil
	}

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// +operator-builder:resource:field=nginx.include,value=true,include
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{
					"controller-gen.kubebuilder.io/version": "v0.21.0",
				},
				"name": "virtualservers.k8s.nginx.org",
				"labels": map[string]interface{}{
					"platform.nukleros.io/category": "ingress",
					"platform.nukleros.io/project":  "nginx-ingress-controller",
				},
			},
			"spec": map[string]interface{}{
				"group": "k8s.nginx.org",
				"names": map[string]interface{}{
					"kind":     "VirtualServer",
					"listKind": "VirtualServerList",
					"plural":   "virtualservers",
					"shortNames": []interface{}{
						"vs",
					},
					"singular": "virtualserver",
				},
				"scope": "Namespaced",
				"versions": []interface{}{
					map[string]interface{}{
						"additionalPrinterColumns": []interface{}{
							map[string]interface{}{
								"description": "Current state of the VirtualServer. If the resource has a valid status, it means it has been validated and accepted by the Ingress Controller.",
								"jsonPath":    ".status.state",
								"name":        "State",
								"type":        "string",
							},
							map[string]interface{}{
								"jsonPath": ".spec.host",
								"name":     "Host",
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".status.externalEndpoints[*].ip",
								"name":     "IP",
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".status.externalEndpoints[*].hostname",
								"name":     "ExternalHostname",
								"priority": 1,
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".status.externalEndpoints[*].ports",
								"name":     "Ports",
								"type":     "string",
							},
							map[string]interface{}{
								"jsonPath": ".metadata.creationTimestamp",
								"name":     "Age",
								"type":     "date",
							},
						},
						"name": "v1",
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"description": "VirtualServer defines the VirtualServer resource.",
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
										"description": "VirtualServerSpec is the spec of the VirtualServer resource.",
										"properties": map[string]interface{}{
											"add-header-inherit": map[string]interface{}{
												"description": "Controls header inheritance behavior at the server level. Allowed values are: on, off, merge. When set to \"merge\", headers from this context are merged with headers in child contexts. When set to \"on\", standard NGINX inheritance applies. When set to \"off\", no headers are inherited from parent contexts.",
												"enum": []interface{}{
													"on",
													"off",
													"merge",
												},
												"type": "string",
											},
											"dos": map[string]interface{}{
												"description": "A reference to a DosProtectedResource, setting this enables DOS protection of the VirtualServer route.",
												"type":        "string",
											},
											"externalDNS": map[string]interface{}{
												"description": "The externalDNS configuration for a VirtualServer.",
												"properties": map[string]interface{}{
													"enable": map[string]interface{}{
														"description": "Enables ExternalDNS integration for a VirtualServer resource. The default is false.",
														"type":        "boolean",
													},
													"labels": map[string]interface{}{
														"additionalProperties": map[string]interface{}{
															"type": "string",
														},
														"description": "Configure labels to be applied to the Endpoint resources that will be consumed by ExternalDNS.",
														"type":        "object",
													},
													"providerSpecific": map[string]interface{}{
														"description": "Configure provider specific properties which holds the name and value of a configuration which is specific to individual DNS providers.",
														"items": map[string]interface{}{
															"description": `ProviderSpecificProperty defines specific property
for using with ExternalDNS sub-resource.`,
															"properties": map[string]interface{}{
																"name": map[string]interface{}{
																	"description": "Name of the property",
																	"type":        "string",
																},
																"value": map[string]interface{}{
																	"description": "Value of the property",
																	"type":        "string",
																},
															},
															"type": "object",
														},
														"type": "array",
													},
													"recordTTL": map[string]interface{}{
														"description": "TTL for the DNS record. This defaults to 0 if not defined.",
														"format":      "int64",
														"type":        "integer",
													},
													"recordType": map[string]interface{}{
														"description": "The record Type that should be created, e.g. “A”, “AAAA”, “CNAME”. This is automatically computed based on the external endpoints if not defined.",
														"type":        "string",
													},
												},
												"type": "object",
											},
											"gunzip": map[string]interface{}{
												"description": "Enables or disables decompression of gzipped responses for clients. Allowed values “on”/“off”, “true”/“false” or “yes”/“no”. If the gunzip value is not set, it defaults to off.",
												"type":        "boolean",
											},
											"host": map[string]interface{}{
												"description": "The host (domain name) of the server. Must be a valid subdomain as defined in RFC 1123, such as my-app or hello.example.com. When using a wildcard domain like *.example.com the domain must be contained in double quotes. The host value needs to be unique among all Ingress and VirtualServer resources.",
												"type":        "string",
											},
											"http-snippets": map[string]interface{}{
												"description": "Sets a custom snippet in the http context.",
												"type":        "string",
											},
											"ingressClassName": map[string]interface{}{
												"description": "Specifies which Ingress Controller must handle the VirtualServerRoute resource. Must be the same as the ingressClassName of the VirtualServer that references this resource.",
												"type":        "string",
											},
											"internalRoute": map[string]interface{}{
												"description": "InternalRoute allows for the configuration of internal routing.",
												"type":        "boolean",
											},
											"listener": map[string]interface{}{
												"description": "Sets a custom HTTP and/or HTTPS listener. Valid fields are listener.http and listener.https. Each field must reference the name of a valid listener defined in a GlobalConfiguration resource",
												"properties": map[string]interface{}{
													"http": map[string]interface{}{
														"description": "The name of an HTTP listener defined in a GlobalConfiguration resource.",
														"type":        "string",
													},
													"https": map[string]interface{}{
														"description": "The name of an HTTPS listener defined in a GlobalConfiguration resource.",
														"type":        "string",
													},
												},
												"type": "object",
											},
											"policies": map[string]interface{}{
												"description": "A list of policies.",
												"items": map[string]interface{}{
													"description": "PolicyReference references a policy by name and an optional namespace.",
													"properties": map[string]interface{}{
														"name": map[string]interface{}{
															"description": "The name of a policy. If the policy doesn’t exist or invalid, NGINX will respond with an error response with the 500 status code.",
															"type":        "string",
														},
														"namespace": map[string]interface{}{
															"description": "The namespace of a policy. If not specified, the namespace of the VirtualServer resource is used.",
															"type":        "string",
														},
													},
													"type": "object",
												},
												"type": "array",
											},
											"routes": map[string]interface{}{
												"description": "A list of routes.",
												"items": map[string]interface{}{
													"description": "Route defines a route.",
													"properties": map[string]interface{}{
														"action": map[string]interface{}{
															"description": "The default action to perform for a request.",
															"properties": map[string]interface{}{
																"pass": map[string]interface{}{
																	"description": "Passes requests to an upstream. The upstream with that name must be defined in the resource.",
																	"type":        "string",
																},
																"proxy": map[string]interface{}{
																	"description": "Passes requests to an upstream with the ability to modify the request/response (for example, rewrite the URI or modify the headers).",
																	"properties": map[string]interface{}{
																		"requestHeaders": map[string]interface{}{
																			"description": "The request headers modifications.",
																			"properties": map[string]interface{}{
																				"pass": map[string]interface{}{
																					"description": "Passes the original request headers to the proxied upstream server.  Default is true.",
																					"type":        "boolean",
																				},
																				"set": map[string]interface{}{
																					"description": "Allows redefining or appending fields to present request headers passed to the proxied upstream servers.",
																					"items": map[string]interface{}{
																						"description": "Header defines an HTTP Header.",
																						"properties": map[string]interface{}{
																							"name": map[string]interface{}{
																								"description": "The name of the header.",
																								"type":        "string",
																							},
																							"value": map[string]interface{}{
																								"description": "The value of the header.",
																								"type":        "string",
																							},
																						},
																						"type": "object",
																					},
																					"type": "array",
																				},
																			},
																			"type": "object",
																		},
																		"responseHeaders": map[string]interface{}{
																			"description": "The response headers modifications.",
																			"properties": map[string]interface{}{
																				"add": map[string]interface{}{
																					"description": "Adds headers to the response to the client.",
																					"items": map[string]interface{}{
																						"description": "AddHeader defines an HTTP Header with an optional Always field to use with the add_header NGINX directive.",
																						"properties": map[string]interface{}{
																							"always": map[string]interface{}{
																								"description": "If set to true, add the header regardless of the response status code**. Default is false.",
																								"type":        "boolean",
																							},
																							"name": map[string]interface{}{
																								"description": "The name of the header.",
																								"type":        "string",
																							},
																							"value": map[string]interface{}{
																								"description": "The value of the header.",
																								"type":        "string",
																							},
																						},
																						"type": "object",
																					},
																					"type": "array",
																				},
																				"hide": map[string]interface{}{
																					"description": "The headers that will not be passed* in the response to the client from a proxied upstream server.",
																					"items": map[string]interface{}{
																						"type": "string",
																					},
																					"type": "array",
																				},
																				"ignore": map[string]interface{}{
																					"description": "Disables processing of certain headers** to the client from a proxied upstream server.",
																					"items": map[string]interface{}{
																						"type": "string",
																					},
																					"type": "array",
																				},
																				"pass": map[string]interface{}{
																					"description": "Allows passing the hidden header fields* to the client from a proxied upstream server.",
																					"items": map[string]interface{}{
																						"type": "string",
																					},
																					"type": "array",
																				},
																			},
																			"type": "object",
																		},
																		"rewritePath": map[string]interface{}{
																			"description": "The rewritten URI. If the route path is a regular expression – starts with ~ – the rewritePath can include capture groups with $1-9. For example $1 for the first group, and so on. For more information, check the rewrite example.",
																			"type":        "string",
																		},
																		"upstream": map[string]interface{}{
																			"description": "The name of the upstream which the requests will be proxied to. The upstream with that name must be defined in the resource.",
																			"type":        "string",
																		},
																	},
																	"type": "object",
																},
																"redirect": map[string]interface{}{
																	"description": "Redirects requests to a provided URL.",
																	"properties": map[string]interface{}{
																		"code": map[string]interface{}{
																			"description": "The status code of a redirect. The allowed values are: 301, 302, 307 or 308. The default is 301.",
																			"type":        "integer",
																		},
																		"url": map[string]interface{}{
																			"description": "The URL to redirect the request to. Supported NGINX variables: $scheme, $http_x_forwarded_proto, $request_uri or $host. Variables must be enclosed in curly braces. For example: ${host}${request_uri}.",
																			"type":        "string",
																		},
																	},
																	"type": "object",
																},
																"return": map[string]interface{}{
																	"description": "Returns a preconfigured response.",
																	"properties": map[string]interface{}{
																		"body": map[string]interface{}{
																			"description": `The body of the response. Supports NGINX variables*. Variables must be enclosed in curly brackets. For example: Request is ${request_uri}\n.`,
																			"type":        "string",
																		},
																		"code": map[string]interface{}{
																			"description": "The status code of the response. The allowed values are: 2XX, 4XX or 5XX. The default is 200.",
																			"type":        "integer",
																		},
																		"headers": map[string]interface{}{
																			"description": "The custom headers of the response.",
																			"items": map[string]interface{}{
																				"description": "Header defines an HTTP Header.",
																				"properties": map[string]interface{}{
																					"name": map[string]interface{}{
																						"description": "The name of the header.",
																						"type":        "string",
																					},
																					"value": map[string]interface{}{
																						"description": "The value of the header.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"type": "array",
																		},
																		"type": map[string]interface{}{
																			"description": "The MIME type of the response. The default is text/plain.",
																			"type":        "string",
																		},
																	},
																	"type": "object",
																},
															},
															"type": "object",
														},
														"add-header-inherit": map[string]interface{}{
															"description": "Controls header inheritance behavior at the location level. Allowed values are: on, off, merge. When set to \"merge\", headers from this context are merged with headers in child contexts. When set to \"on\", standard NGINX inheritance applies. When set to \"off\", no headers are inherited from parent contexts.",
															"enum": []interface{}{
																"on",
																"off",
																"merge",
															},
															"type": "string",
														},
														"dos": map[string]interface{}{
															"description": "A reference to a DosProtectedResource, setting this enables DOS protection of the VirtualServer route.",
															"type":        "string",
														},
														"errorPages": map[string]interface{}{
															"description": "The custom responses for error codes. NGINX will use those responses instead of returning the error responses from the upstream servers or the default responses generated by NGINX. A custom response can be a redirect or a canned response. For example, a redirect to another URL if an upstream server responded with a 404 status code.",
															"items": map[string]interface{}{
																"description": "ErrorPage defines an ErrorPage in a Route.",
																"properties": map[string]interface{}{
																	"codes": map[string]interface{}{
																		"description": "A list of error status codes.",
																		"items": map[string]interface{}{
																			"type": "integer",
																		},
																		"type": "array",
																	},
																	"redirect": map[string]interface{}{
																		"description": "The canned response action for the given status codes.",
																		"properties": map[string]interface{}{
																			"code": map[string]interface{}{
																				"description": "The status code of a redirect. The allowed values are: 301, 302, 307 or 308. The default is 301.",
																				"type":        "integer",
																			},
																			"url": map[string]interface{}{
																				"description": "The URL to redirect the request to. Supported NGINX variables: $scheme, $http_x_forwarded_proto, $request_uri or $host. Variables must be enclosed in curly braces. For example: ${host}${request_uri}.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																	"return": map[string]interface{}{
																		"description": "The redirect action for the given status codes.",
																		"properties": map[string]interface{}{
																			"body": map[string]interface{}{
																				"description": `The body of the response. Supports NGINX variables*. Variables must be enclosed in curly brackets. For example: Request is ${request_uri}\n.`,
																				"type":        "string",
																			},
																			"code": map[string]interface{}{
																				"description": "The status code of the response. The allowed values are: 2XX, 4XX or 5XX. The default is 200.",
																				"type":        "integer",
																			},
																			"headers": map[string]interface{}{
																				"description": "The custom headers of the response.",
																				"items": map[string]interface{}{
																					"description": "Header defines an HTTP Header.",
																					"properties": map[string]interface{}{
																						"name": map[string]interface{}{
																							"description": "The name of the header.",
																							"type":        "string",
																						},
																						"value": map[string]interface{}{
																							"description": "The value of the header.",
																							"type":        "string",
																						},
																					},
																					"type": "object",
																				},
																				"type": "array",
																			},
																			"type": map[string]interface{}{
																				"description": "The MIME type of the response. The default is text/plain.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																},
																"type": "object",
															},
															"type": "array",
														},
														"location-snippets": map[string]interface{}{
															"description": "Sets a custom snippet in the location context. Overrides the location-snippets ConfigMap key.",
															"type":        "string",
														},
														"matches": map[string]interface{}{
															"description": "The matching rules for advanced content-based routing. Requires the default Action or Splits. Unmatched requests will be handled by the default Action or Splits.",
															"items": map[string]interface{}{
																"description": "Match defines a match.",
																"properties": map[string]interface{}{
																	"action": map[string]interface{}{
																		"description": "The action to perform for a request.",
																		"properties": map[string]interface{}{
																			"pass": map[string]interface{}{
																				"description": "Passes requests to an upstream. The upstream with that name must be defined in the resource.",
																				"type":        "string",
																			},
																			"proxy": map[string]interface{}{
																				"description": "Passes requests to an upstream with the ability to modify the request/response (for example, rewrite the URI or modify the headers).",
																				"properties": map[string]interface{}{
																					"requestHeaders": map[string]interface{}{
																						"description": "The request headers modifications.",
																						"properties": map[string]interface{}{
																							"pass": map[string]interface{}{
																								"description": "Passes the original request headers to the proxied upstream server.  Default is true.",
																								"type":        "boolean",
																							},
																							"set": map[string]interface{}{
																								"description": "Allows redefining or appending fields to present request headers passed to the proxied upstream servers.",
																								"items": map[string]interface{}{
																									"description": "Header defines an HTTP Header.",
																									"properties": map[string]interface{}{
																										"name": map[string]interface{}{
																											"description": "The name of the header.",
																											"type":        "string",
																										},
																										"value": map[string]interface{}{
																											"description": "The value of the header.",
																											"type":        "string",
																										},
																									},
																									"type": "object",
																								},
																								"type": "array",
																							},
																						},
																						"type": "object",
																					},
																					"responseHeaders": map[string]interface{}{
																						"description": "The response headers modifications.",
																						"properties": map[string]interface{}{
																							"add": map[string]interface{}{
																								"description": "Adds headers to the response to the client.",
																								"items": map[string]interface{}{
																									"description": "AddHeader defines an HTTP Header with an optional Always field to use with the add_header NGINX directive.",
																									"properties": map[string]interface{}{
																										"always": map[string]interface{}{
																											"description": "If set to true, add the header regardless of the response status code**. Default is false.",
																											"type":        "boolean",
																										},
																										"name": map[string]interface{}{
																											"description": "The name of the header.",
																											"type":        "string",
																										},
																										"value": map[string]interface{}{
																											"description": "The value of the header.",
																											"type":        "string",
																										},
																									},
																									"type": "object",
																								},
																								"type": "array",
																							},
																							"hide": map[string]interface{}{
																								"description": "The headers that will not be passed* in the response to the client from a proxied upstream server.",
																								"items": map[string]interface{}{
																									"type": "string",
																								},
																								"type": "array",
																							},
																							"ignore": map[string]interface{}{
																								"description": "Disables processing of certain headers** to the client from a proxied upstream server.",
																								"items": map[string]interface{}{
																									"type": "string",
																								},
																								"type": "array",
																							},
																							"pass": map[string]interface{}{
																								"description": "Allows passing the hidden header fields* to the client from a proxied upstream server.",
																								"items": map[string]interface{}{
																									"type": "string",
																								},
																								"type": "array",
																							},
																						},
																						"type": "object",
																					},
																					"rewritePath": map[string]interface{}{
																						"description": "The rewritten URI. If the route path is a regular expression – starts with ~ – the rewritePath can include capture groups with $1-9. For example $1 for the first group, and so on. For more information, check the rewrite example.",
																						"type":        "string",
																					},
																					"upstream": map[string]interface{}{
																						"description": "The name of the upstream which the requests will be proxied to. The upstream with that name must be defined in the resource.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"redirect": map[string]interface{}{
																				"description": "Redirects requests to a provided URL.",
																				"properties": map[string]interface{}{
																					"code": map[string]interface{}{
																						"description": "The status code of a redirect. The allowed values are: 301, 302, 307 or 308. The default is 301.",
																						"type":        "integer",
																					},
																					"url": map[string]interface{}{
																						"description": "The URL to redirect the request to. Supported NGINX variables: $scheme, $http_x_forwarded_proto, $request_uri or $host. Variables must be enclosed in curly braces. For example: ${host}${request_uri}.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"return": map[string]interface{}{
																				"description": "Returns a preconfigured response.",
																				"properties": map[string]interface{}{
																					"body": map[string]interface{}{
																						"description": `The body of the response. Supports NGINX variables*. Variables must be enclosed in curly brackets. For example: Request is ${request_uri}\n.`,
																						"type":        "string",
																					},
																					"code": map[string]interface{}{
																						"description": "The status code of the response. The allowed values are: 2XX, 4XX or 5XX. The default is 200.",
																						"type":        "integer",
																					},
																					"headers": map[string]interface{}{
																						"description": "The custom headers of the response.",
																						"items": map[string]interface{}{
																							"description": "Header defines an HTTP Header.",
																							"properties": map[string]interface{}{
																								"name": map[string]interface{}{
																									"description": "The name of the header.",
																									"type":        "string",
																								},
																								"value": map[string]interface{}{
																									"description": "The value of the header.",
																									"type":        "string",
																								},
																							},
																							"type": "object",
																						},
																						"type": "array",
																					},
																					"type": map[string]interface{}{
																						"description": "The MIME type of the response. The default is text/plain.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																	"conditions": map[string]interface{}{
																		"description": "A list of conditions. Must include at least 1 condition.",
																		"items": map[string]interface{}{
																			"description": "Condition defines a condition in a MatchRule.",
																			"properties": map[string]interface{}{
																				"argument": map[string]interface{}{
																					"description": "The name of an argument. Must consist of alphanumeric characters or _.",
																					"type":        "string",
																				},
																				"cookie": map[string]interface{}{
																					"description": "The name of a cookie. Must consist of alphanumeric characters or _.",
																					"type":        "string",
																				},
																				"header": map[string]interface{}{
																					"description": "The name of a header. Must consist of alphanumeric characters or -.",
																					"type":        "string",
																				},
																				"value": map[string]interface{}{
																					"description": "The value to match the condition against.",
																					"type":        "string",
																				},
																				"variable": map[string]interface{}{
																					"description": "The name of an NGINX variable. Must start with $.",
																					"type":        "string",
																				},
																			},
																			"type": "object",
																		},
																		"type": "array",
																	},
																	"splits": map[string]interface{}{
																		"description": "The splits configuration for traffic splitting. Must include at least 2 splits.",
																		"items": map[string]interface{}{
																			"description": "Split defines a split.",
																			"properties": map[string]interface{}{
																				"action": map[string]interface{}{
																					"description": "The action to perform for a request.",
																					"properties": map[string]interface{}{
																						"pass": map[string]interface{}{
																							"description": "Passes requests to an upstream. The upstream with that name must be defined in the resource.",
																							"type":        "string",
																						},
																						"proxy": map[string]interface{}{
																							"description": "Passes requests to an upstream with the ability to modify the request/response (for example, rewrite the URI or modify the headers).",
																							"properties": map[string]interface{}{
																								"requestHeaders": map[string]interface{}{
																									"description": "The request headers modifications.",
																									"properties": map[string]interface{}{
																										"pass": map[string]interface{}{
																											"description": "Passes the original request headers to the proxied upstream server.  Default is true.",
																											"type":        "boolean",
																										},
																										"set": map[string]interface{}{
																											"description": "Allows redefining or appending fields to present request headers passed to the proxied upstream servers.",
																											"items": map[string]interface{}{
																												"description": "Header defines an HTTP Header.",
																												"properties": map[string]interface{}{
																													"name": map[string]interface{}{
																														"description": "The name of the header.",
																														"type":        "string",
																													},
																													"value": map[string]interface{}{
																														"description": "The value of the header.",
																														"type":        "string",
																													},
																												},
																												"type": "object",
																											},
																											"type": "array",
																										},
																									},
																									"type": "object",
																								},
																								"responseHeaders": map[string]interface{}{
																									"description": "The response headers modifications.",
																									"properties": map[string]interface{}{
																										"add": map[string]interface{}{
																											"description": "Adds headers to the response to the client.",
																											"items": map[string]interface{}{
																												"description": "AddHeader defines an HTTP Header with an optional Always field to use with the add_header NGINX directive.",
																												"properties": map[string]interface{}{
																													"always": map[string]interface{}{
																														"description": "If set to true, add the header regardless of the response status code**. Default is false.",
																														"type":        "boolean",
																													},
																													"name": map[string]interface{}{
																														"description": "The name of the header.",
																														"type":        "string",
																													},
																													"value": map[string]interface{}{
																														"description": "The value of the header.",
																														"type":        "string",
																													},
																												},
																												"type": "object",
																											},
																											"type": "array",
																										},
																										"hide": map[string]interface{}{
																											"description": "The headers that will not be passed* in the response to the client from a proxied upstream server.",
																											"items": map[string]interface{}{
																												"type": "string",
																											},
																											"type": "array",
																										},
																										"ignore": map[string]interface{}{
																											"description": "Disables processing of certain headers** to the client from a proxied upstream server.",
																											"items": map[string]interface{}{
																												"type": "string",
																											},
																											"type": "array",
																										},
																										"pass": map[string]interface{}{
																											"description": "Allows passing the hidden header fields* to the client from a proxied upstream server.",
																											"items": map[string]interface{}{
																												"type": "string",
																											},
																											"type": "array",
																										},
																									},
																									"type": "object",
																								},
																								"rewritePath": map[string]interface{}{
																									"description": "The rewritten URI. If the route path is a regular expression – starts with ~ – the rewritePath can include capture groups with $1-9. For example $1 for the first group, and so on. For more information, check the rewrite example.",
																									"type":        "string",
																								},
																								"upstream": map[string]interface{}{
																									"description": "The name of the upstream which the requests will be proxied to. The upstream with that name must be defined in the resource.",
																									"type":        "string",
																								},
																							},
																							"type": "object",
																						},
																						"redirect": map[string]interface{}{
																							"description": "Redirects requests to a provided URL.",
																							"properties": map[string]interface{}{
																								"code": map[string]interface{}{
																									"description": "The status code of a redirect. The allowed values are: 301, 302, 307 or 308. The default is 301.",
																									"type":        "integer",
																								},
																								"url": map[string]interface{}{
																									"description": "The URL to redirect the request to. Supported NGINX variables: $scheme, $http_x_forwarded_proto, $request_uri or $host. Variables must be enclosed in curly braces. For example: ${host}${request_uri}.",
																									"type":        "string",
																								},
																							},
																							"type": "object",
																						},
																						"return": map[string]interface{}{
																							"description": "Returns a preconfigured response.",
																							"properties": map[string]interface{}{
																								"body": map[string]interface{}{
																									"description": `The body of the response. Supports NGINX variables*. Variables must be enclosed in curly brackets. For example: Request is ${request_uri}\n.`,
																									"type":        "string",
																								},
																								"code": map[string]interface{}{
																									"description": "The status code of the response. The allowed values are: 2XX, 4XX or 5XX. The default is 200.",
																									"type":        "integer",
																								},
																								"headers": map[string]interface{}{
																									"description": "The custom headers of the response.",
																									"items": map[string]interface{}{
																										"description": "Header defines an HTTP Header.",
																										"properties": map[string]interface{}{
																											"name": map[string]interface{}{
																												"description": "The name of the header.",
																												"type":        "string",
																											},
																											"value": map[string]interface{}{
																												"description": "The value of the header.",
																												"type":        "string",
																											},
																										},
																										"type": "object",
																									},
																									"type": "array",
																								},
																								"type": map[string]interface{}{
																									"description": "The MIME type of the response. The default is text/plain.",
																									"type":        "string",
																								},
																							},
																							"type": "object",
																						},
																					},
																					"type": "object",
																				},
																				"weight": map[string]interface{}{
																					"description": "The weight of an action. Must fall into the range 0..100. The sum of the weights of all splits must be equal to 100.",
																					"type":        "integer",
																				},
																			},
																			"type": "object",
																		},
																		"type": "array",
																	},
																},
																"type": "object",
															},
															"type": "array",
														},
														"path": map[string]interface{}{
															"description": `The path of the route. NGINX will match it against the URI of a request. Possible values are: a prefix ( / , /path ), a longest prefix match ( ^~/images/ ), an exact match ( =/exact/match ), a case-insensitive regular expression ( ~*^/Bar.*\.jpg ) or a case-sensitive regular expression ( ~^/foo.*\.jpg ). In the case of a prefix match (must start with / ), a longest prefix match (must start with ^~ ) or an exact match (must start with = ), the path must not include any whitespace characters, { , } or ;. In the case of the regex matches, all double quotes " must be escaped and the match can’t end in an unescaped backslash \. The path must be unique among the paths of all routes of the VirtualServer. Check the location directive for more information.`,
															"type":        "string",
														},
														"policies": map[string]interface{}{
															"description": "A list of policies. The policies override the policies of the same type defined in the spec of the VirtualServer.",
															"items": map[string]interface{}{
																"description": "PolicyReference references a policy by name and an optional namespace.",
																"properties": map[string]interface{}{
																	"name": map[string]interface{}{
																		"description": "The name of a policy. If the policy doesn’t exist or invalid, NGINX will respond with an error response with the 500 status code.",
																		"type":        "string",
																	},
																	"namespace": map[string]interface{}{
																		"description": "The namespace of a policy. If not specified, the namespace of the VirtualServer resource is used.",
																		"type":        "string",
																	},
																},
																"type": "object",
															},
															"type": "array",
														},
														"route": map[string]interface{}{
															"description": "The name of a VirtualServerRoute resource that defines this route. If the VirtualServerRoute belongs to a different namespace than the VirtualServer, you need to include the namespace. For example, tea-namespace/tea.",
															"type":        "string",
														},
														"routeSelector": map[string]interface{}{
															"description": "The RouteSelector allows selecting VirtualServerRoute resources using label selectors.",
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
														"splits": map[string]interface{}{
															"description": "The default splits configuration for traffic splitting. Must include at least 2 splits.",
															"items": map[string]interface{}{
																"description": "Split defines a split.",
																"properties": map[string]interface{}{
																	"action": map[string]interface{}{
																		"description": "The action to perform for a request.",
																		"properties": map[string]interface{}{
																			"pass": map[string]interface{}{
																				"description": "Passes requests to an upstream. The upstream with that name must be defined in the resource.",
																				"type":        "string",
																			},
																			"proxy": map[string]interface{}{
																				"description": "Passes requests to an upstream with the ability to modify the request/response (for example, rewrite the URI or modify the headers).",
																				"properties": map[string]interface{}{
																					"requestHeaders": map[string]interface{}{
																						"description": "The request headers modifications.",
																						"properties": map[string]interface{}{
																							"pass": map[string]interface{}{
																								"description": "Passes the original request headers to the proxied upstream server.  Default is true.",
																								"type":        "boolean",
																							},
																							"set": map[string]interface{}{
																								"description": "Allows redefining or appending fields to present request headers passed to the proxied upstream servers.",
																								"items": map[string]interface{}{
																									"description": "Header defines an HTTP Header.",
																									"properties": map[string]interface{}{
																										"name": map[string]interface{}{
																											"description": "The name of the header.",
																											"type":        "string",
																										},
																										"value": map[string]interface{}{
																											"description": "The value of the header.",
																											"type":        "string",
																										},
																									},
																									"type": "object",
																								},
																								"type": "array",
																							},
																						},
																						"type": "object",
																					},
																					"responseHeaders": map[string]interface{}{
																						"description": "The response headers modifications.",
																						"properties": map[string]interface{}{
																							"add": map[string]interface{}{
																								"description": "Adds headers to the response to the client.",
																								"items": map[string]interface{}{
																									"description": "AddHeader defines an HTTP Header with an optional Always field to use with the add_header NGINX directive.",
																									"properties": map[string]interface{}{
																										"always": map[string]interface{}{
																											"description": "If set to true, add the header regardless of the response status code**. Default is false.",
																											"type":        "boolean",
																										},
																										"name": map[string]interface{}{
																											"description": "The name of the header.",
																											"type":        "string",
																										},
																										"value": map[string]interface{}{
																											"description": "The value of the header.",
																											"type":        "string",
																										},
																									},
																									"type": "object",
																								},
																								"type": "array",
																							},
																							"hide": map[string]interface{}{
																								"description": "The headers that will not be passed* in the response to the client from a proxied upstream server.",
																								"items": map[string]interface{}{
																									"type": "string",
																								},
																								"type": "array",
																							},
																							"ignore": map[string]interface{}{
																								"description": "Disables processing of certain headers** to the client from a proxied upstream server.",
																								"items": map[string]interface{}{
																									"type": "string",
																								},
																								"type": "array",
																							},
																							"pass": map[string]interface{}{
																								"description": "Allows passing the hidden header fields* to the client from a proxied upstream server.",
																								"items": map[string]interface{}{
																									"type": "string",
																								},
																								"type": "array",
																							},
																						},
																						"type": "object",
																					},
																					"rewritePath": map[string]interface{}{
																						"description": "The rewritten URI. If the route path is a regular expression – starts with ~ – the rewritePath can include capture groups with $1-9. For example $1 for the first group, and so on. For more information, check the rewrite example.",
																						"type":        "string",
																					},
																					"upstream": map[string]interface{}{
																						"description": "The name of the upstream which the requests will be proxied to. The upstream with that name must be defined in the resource.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"redirect": map[string]interface{}{
																				"description": "Redirects requests to a provided URL.",
																				"properties": map[string]interface{}{
																					"code": map[string]interface{}{
																						"description": "The status code of a redirect. The allowed values are: 301, 302, 307 or 308. The default is 301.",
																						"type":        "integer",
																					},
																					"url": map[string]interface{}{
																						"description": "The URL to redirect the request to. Supported NGINX variables: $scheme, $http_x_forwarded_proto, $request_uri or $host. Variables must be enclosed in curly braces. For example: ${host}${request_uri}.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																			"return": map[string]interface{}{
																				"description": "Returns a preconfigured response.",
																				"properties": map[string]interface{}{
																					"body": map[string]interface{}{
																						"description": `The body of the response. Supports NGINX variables*. Variables must be enclosed in curly brackets. For example: Request is ${request_uri}\n.`,
																						"type":        "string",
																					},
																					"code": map[string]interface{}{
																						"description": "The status code of the response. The allowed values are: 2XX, 4XX or 5XX. The default is 200.",
																						"type":        "integer",
																					},
																					"headers": map[string]interface{}{
																						"description": "The custom headers of the response.",
																						"items": map[string]interface{}{
																							"description": "Header defines an HTTP Header.",
																							"properties": map[string]interface{}{
																								"name": map[string]interface{}{
																									"description": "The name of the header.",
																									"type":        "string",
																								},
																								"value": map[string]interface{}{
																									"description": "The value of the header.",
																									"type":        "string",
																								},
																							},
																							"type": "object",
																						},
																						"type": "array",
																					},
																					"type": map[string]interface{}{
																						"description": "The MIME type of the response. The default is text/plain.",
																						"type":        "string",
																					},
																				},
																				"type": "object",
																			},
																		},
																		"type": "object",
																	},
																	"weight": map[string]interface{}{
																		"description": "The weight of an action. Must fall into the range 0..100. The sum of the weights of all splits must be equal to 100.",
																		"type":        "integer",
																	},
																},
																"type": "object",
															},
															"type": "array",
														},
													},
													"type": "object",
												},
												"type": "array",
											},
											"server-snippets": map[string]interface{}{
												"description": "Sets a custom snippet in server context. Overrides the server-snippets ConfigMap key.",
												"type":        "string",
											},
											"tls": map[string]interface{}{
												"description": "The TLS termination configuration.",
												"properties": map[string]interface{}{
													"cert-manager": map[string]interface{}{
														"description": "The cert-manager configuration of the TLS for a VirtualServer.",
														"properties": map[string]interface{}{
															"cluster-issuer": map[string]interface{}{
																"description": "the name of a ClusterIssuer. A ClusterIssuer is a cert-manager resource which describes the certificate authority capable of signing certificates. It does not matter which namespace your VirtualServer resides, as ClusterIssuers are non-namespaced resources. Please note that one of issuer and cluster-issuer are required, but they are mutually exclusive - one and only one must be defined.",
																"type":        "string",
															},
															"common-name": map[string]interface{}{
																"description": "This field allows you to configure spec.commonName for the Certificate to be generated. This configuration adds a CN to the x509 certificate.",
																"type":        "string",
															},
															"duration": map[string]interface{}{
																"description": "This field allows you to configure spec.duration field for the Certificate to be generated. Must be specified using a Go time.Duration string format, which does not allow the d (days) suffix. You must specify these values using s, m, and h suffixes instead.",
																"type":        "string",
															},
															"issue-temp-cert": map[string]interface{}{
																"description": "When true, ask cert-manager for a temporary self-signed certificate pending the issuance of the Certificate. This allows HTTPS-only servers to use ACME HTTP01 challenges when the TLS secret does not exist yet.",
																"type":        "boolean",
															},
															"issuer": map[string]interface{}{
																"description": "the name of an Issuer. An Issuer is a cert-manager resource which describes the certificate authority capable of signing certificates. The Issuer must be in the same namespace as the VirtualServer resource. Please note that one of issuer and cluster-issuer are required, but they are mutually exclusive - one and only one must be defined.",
																"type":        "string",
															},
															"issuer-group": map[string]interface{}{
																"description": "The API group of the external issuer controller, for example awspca.cert-manager.io. This is only necessary for out-of-tree issuers. This cannot be defined if cluster-issuer is also defined.",
																"type":        "string",
															},
															"issuer-kind": map[string]interface{}{
																"description": "The kind of the external issuer resource, for example AWSPCAIssuer. This is only necessary for out-of-tree issuers. This cannot be defined if cluster-issuer is also defined.",
																"type":        "string",
															},
															"renew-before": map[string]interface{}{
																"description": "this annotation allows you to configure spec.renewBefore field for the Certificate to be generated. Must be specified using a Go time.Duration string format, which does not allow the d (days) suffix. You must specify these values using s, m, and h suffixes instead.",
																"type":        "string",
															},
															"usages": map[string]interface{}{
																"description": "This field allows you to configure spec.usages field for the Certificate to be generated. Pass a string with comma-separated values i.e. key agreement,digital signature, server auth. An exhaustive list of supported key usages can be found in the the cert-manager api documentation.",
																"type":        "string",
															},
														},
														"type": "object",
													},
													"redirect": map[string]interface{}{
														"description": "The redirect configuration of the TLS for a VirtualServer.",
														"properties": map[string]interface{}{
															"basedOn": map[string]interface{}{
																"description": "The attribute of a request that NGINX will evaluate to send a redirect. The allowed values are scheme (the scheme of the request) or x-forwarded-proto (the X-Forwarded-Proto header of the request). The default is scheme.",
																"type":        "string",
															},
															"code": map[string]interface{}{
																"description": "The status code of a redirect. The allowed values are: 301, 302, 307 or 308. The default is 301.",
																"type":        "integer",
															},
															"enable": map[string]interface{}{
																"description": "Enables a TLS redirect for a VirtualServer. The default is False.",
																"type":        "boolean",
															},
														},
														"type": "object",
													},
													"secret": map[string]interface{}{
														"description": "The name of a secret with a TLS certificate and key. The secret must belong to the same namespace as the VirtualServer. The secret must be of the type kubernetes.io/tls and contain keys named tls.crt and tls.key that contain the certificate and private key as described here. If the secret doesn’t exist or is invalid, NGINX will break any attempt to establish a TLS connection to the host of the VirtualServer. If the secret is not specified but wildcard TLS secret is configured, NGINX will use the wildcard secret for TLS termination.",
														"type":        "string",
													},
												},
												"type": "object",
											},
											"upstreams": map[string]interface{}{
												"description": "A list of upstreams.",
												"items": map[string]interface{}{
													"description": "Upstream defines an upstream.",
													"properties": map[string]interface{}{
														"backup": map[string]interface{}{
															"description": "The name of the backup service of type ExternalName. This will be used when the primary servers are unavailable. Note: The parameter cannot be used along with the random, hash or ip_hash load balancing methods.",
															"type":        "string",
														},
														"backupPort": map[string]interface{}{
															"description": "The port of the backup service. The backup port is required if the backup service name is provided. The port must fall into the range 1..65535.",
															"type":        "integer",
														},
														"buffer-size": map[string]interface{}{
															"description": "Sets the size of the buffer used for reading the first part of a response received from the upstream server. The default is set in the proxy-buffer-size ConfigMap key.",
															"type":        "string",
														},
														"buffering": map[string]interface{}{
															"description": "Enables buffering of responses from the upstream server.  The default is set in the proxy-buffering ConfigMap key.",
															"type":        "boolean",
														},
														"buffers": map[string]interface{}{
															"description": "Configures the buffers used for reading a response from the upstream server for a single connection.",
															"properties": map[string]interface{}{
																"number": map[string]interface{}{
																	"description": "Configures the number of buffers. The default is set in the proxy-buffers ConfigMap key.",
																	"type":        "integer",
																},
																"size": map[string]interface{}{
																	"description": "Configures the size of a buffer. The default is set in the proxy-buffers ConfigMap key.",
																	"type":        "string",
																},
															},
															"type": "object",
														},
														"busy-buffers-size": map[string]interface{}{
															"description": "Sets the size of the buffers used for reading a response from the upstream server when the proxy_buffering is enabled. The default is set in the proxy-busy-buffers-size ConfigMap key.'",
															"type":        "string",
														},
														"client-body-buffer-size": map[string]interface{}{
															"description": `ClientBodyBufferSize sets the size of the buffer used for reading the client request body. Must be specified as a number followed by:
'k' for kilobytes or 'm' for megabytes.
Examples: "10m" or "512k".`,
															"pattern": `^\d+[kKmM]?$`,
															"type":    "string",
														},
														"client-max-body-size": map[string]interface{}{
															"description": "Sets the maximum allowed size of the client request body. The default is set in the client-max-body-size ConfigMap key.",
															"type":        "string",
														},
														"connect-timeout": map[string]interface{}{
															"description": "The timeout for establishing a connection with an upstream server. The default is specified in the proxy-connect-timeout ConfigMap key.",
															"type":        "string",
														},
														"fail-timeout": map[string]interface{}{
															"description": "The time during which the specified number of unsuccessful attempts to communicate with an upstream server should happen to consider the server unavailable. The default is set in the fail-timeout ConfigMap key.",
															"type":        "string",
														},
														"healthCheck": map[string]interface{}{
															"description": "The health check configuration for the Upstream. Note: this feature is supported only in NGINX Plus.",
															"properties": map[string]interface{}{
																"connect-timeout": map[string]interface{}{
																	"description": "The timeout for establishing a connection with an upstream server. By default, the connect-timeout of the upstream is used.",
																	"type":        "string",
																},
																"enable": map[string]interface{}{
																	"description": "Enables a health check for an upstream server. The default is false.",
																	"type":        "boolean",
																},
																"fails": map[string]interface{}{
																	"description": "The number of consecutive failed health checks of a particular upstream server after which this server will be considered unhealthy. The default is 1.",
																	"type":        "integer",
																},
																"grpcService": map[string]interface{}{
																	"description": "The gRPC service to be monitored on the upstream server. Only valid on gRPC type upstreams.",
																	"type":        "string",
																},
																"grpcStatus": map[string]interface{}{
																	"description": "The expected gRPC status code of the upstream server response to the Check method. Configure this field only if your gRPC services do not implement the gRPC health checking protocol. For example, configure 12 if the upstream server responds with 12 (UNIMPLEMENTED) status code. Only valid on gRPC type upstreams.",
																	"type":        "integer",
																},
																"headers": map[string]interface{}{
																	"description": "The request headers used for health check requests. NGINX Plus always sets the Host, User-Agent and Connection headers for health check requests.",
																	"items": map[string]interface{}{
																		"description": "Header defines an HTTP Header.",
																		"properties": map[string]interface{}{
																			"name": map[string]interface{}{
																				"description": "The name of the header.",
																				"type":        "string",
																			},
																			"value": map[string]interface{}{
																				"description": "The value of the header.",
																				"type":        "string",
																			},
																		},
																		"type": "object",
																	},
																	"type": "array",
																},
																"interval": map[string]interface{}{
																	"description": "The interval between two consecutive health checks. The default is 5s.",
																	"type":        "string",
																},
																"jitter": map[string]interface{}{
																	"description": "The time within which each health check will be randomly delayed. By default, there is no delay.",
																	"type":        "string",
																},
																"keepalive-time": map[string]interface{}{
																	"description": "Enables keepalive connections for health checks and specifies the time during which requests can be processed through one keepalive connection. The default is 60s.",
																	"type":        "string",
																},
																"mandatory": map[string]interface{}{
																	"description": "Require every newly added server to pass all configured health checks before NGINX Plus sends traffic to it. If this is not specified, or is set to false, the server will be initially considered healthy. When combined with slow-start, it gives a new server more time to connect to databases and “warm up” before being asked to handle their full share of traffic.",
																	"type":        "boolean",
																},
																"passes": map[string]interface{}{
																	"description": "The number of consecutive passed health checks of a particular upstream server after which the server will be considered healthy. The default is 1.",
																	"type":        "integer",
																},
																"path": map[string]interface{}{
																	"description": "The path used for health check requests. The default is /. This is not configurable for gRPC type upstreams.",
																	"type":        "string",
																},
																"persistent": map[string]interface{}{
																	"description": "Set the initial “up” state for a server after reload if the server was considered healthy before reload. Enabling persistent requires that the mandatory parameter is also set to true.",
																	"type":        "boolean",
																},
																"port": map[string]interface{}{
																	"description": "The port used for health check requests. By default, the server port is used. Note: in contrast with the port of the upstream, this port is not a service port, but a port of a pod.",
																	"type":        "integer",
																},
																"read-timeout": map[string]interface{}{
																	"description": "The timeout for reading a response from an upstream server. By default, the read-timeout of the upstream is used.",
																	"type":        "string",
																},
																"send-timeout": map[string]interface{}{
																	"description": "The timeout for transmitting a request to an upstream server. By default, the send-timeout of the upstream is used.",
																	"type":        "string",
																},
																"statusMatch": map[string]interface{}{
																	"description": "The expected response status codes of a health check. By default, the response should have status code 2xx or 3xx. Examples: \"200\", \"! 500\", \"301-303 307\". This not supported for gRPC type upstreams.",
																	"type":        "string",
																},
																"tls": map[string]interface{}{
																	"description": "The TLS configuration used for health check requests. By default, the tls field of the upstream is used.",
																	"properties": map[string]interface{}{
																		"enable": map[string]interface{}{
																			"description": "Enables HTTPS for requests to upstream servers. The default is False , meaning that HTTP will be used. Note: by default, NGINX will not verify the upstream server certificate. To enable the verification, configure an EgressMTLS Policy.",
																			"type":        "boolean",
																		},
																	},
																	"type": "object",
																},
															},
															"type": "object",
														},
														"keepalive": map[string]interface{}{
															"description": "Configures the cache for connections to upstream servers. The value 0 disables the cache. The default is set in the keepalive ConfigMap key.",
															"type":        "integer",
														},
														"lb-method": map[string]interface{}{
															"description": "The load balancing method. To use the round-robin method, specify round_robin. The default is specified in the lb-method ConfigMap key.",
															"type":        "string",
														},
														"max-conns": map[string]interface{}{
															"description": "The maximum number of simultaneous active connections to an upstream server. By default there is no limit. Note: if keepalive connections are enabled, the total number of active and idle keepalive connections to an upstream server may exceed the max_conns value.",
															"type":        "integer",
														},
														"max-fails": map[string]interface{}{
															"description": "The number of unsuccessful attempts to communicate with an upstream server that should happen in the duration set by the fail-timeout to consider the server unavailable. The default is set in the max-fails ConfigMap key.",
															"type":        "integer",
														},
														"name": map[string]interface{}{
															"description": "The name of the upstream. Must be a valid DNS label as defined in RFC 1035. For example, hello and upstream-123 are valid. The name must be unique among all upstreams of the resource.",
															"type":        "string",
														},
														"next-upstream": map[string]interface{}{
															"description": "Specifies in which cases a request should be passed to the next upstream server. The default is error timeout.",
															"type":        "string",
														},
														"next-upstream-timeout": map[string]interface{}{
															"description": "The time during which a request can be passed to the next upstream server. The 0 value turns off the time limit. The default is 0.",
															"type":        "string",
														},
														"next-upstream-tries": map[string]interface{}{
															"description": "The number of possible tries for passing a request to the next upstream server. The 0 value turns off this limit. The default is 0.",
															"type":        "integer",
														},
														"ntlm": map[string]interface{}{
															"description": "Allows proxying requests with NTLM Authentication. In order for NTLM authentication to work, it is necessary to enable keepalive connections to upstream servers using the keepalive field. Note: this feature is supported only in NGINX Plus.",
															"type":        "boolean",
														},
														"port": map[string]interface{}{
															"description": "The port of the service. If the service doesn’t define that port, NGINX will assume the service has zero endpoints and return a 502 response for requests for this upstream. The port must fall into the range 1..65535.",
															"type":        "integer",
														},
														"queue": map[string]interface{}{
															"description": "Configures a queue for an upstream. A client request will be placed into the queue if an upstream server cannot be selected immediately while processing the request. By default, no queue is configured. Note: this feature is supported only in NGINX Plus.",
															"properties": map[string]interface{}{
																"size": map[string]interface{}{
																	"description": "The size of the queue.",
																	"type":        "integer",
																},
																"timeout": map[string]interface{}{
																	"description": "The timeout of the queue. A request cannot be queued for a period longer than the timeout. The default is 60s.",
																	"type":        "string",
																},
															},
															"type": "object",
														},
														"read-timeout": map[string]interface{}{
															"description": "The timeout for reading a response from an upstream server. The default is specified in the proxy-read-timeout ConfigMap key.",
															"type":        "string",
														},
														"send-timeout": map[string]interface{}{
															"description": "The timeout for transmitting a request to an upstream server. The default is specified in the proxy-send-timeout ConfigMap key.",
															"type":        "string",
														},
														"service": map[string]interface{}{
															"description": "The name of a service. If the Service belongs to a different namespace than the VirtualServer or VirtualServerRoute, you need to include the namespace. For example, tea-namespace/tea. If the service doesn’t exist, NGINX will assume the service has zero endpoints and return a 502 response for requests for this upstream. For NGINX Plus only, services of type ExternalName are also supported in the same namespace.",
															"type":        "string",
														},
														"sessionCookie": map[string]interface{}{
															"description": "The SessionCookie field configures session persistence which allows requests from the same client to be passed to the same upstream server. The information about the designated upstream server is passed in a session cookie generated by NGINX.",
															"properties": map[string]interface{}{
																"domain": map[string]interface{}{
																	"description": "The domain for which the cookie is set.",
																	"type":        "string",
																},
																"enable": map[string]interface{}{
																	"description": "Enables session persistence with a session cookie for an upstream server. The default is false.",
																	"type":        "boolean",
																},
																"expires": map[string]interface{}{
																	"description": "The time for which a browser should keep the cookie. Can be set to the special value max, which will cause the cookie to expire on 31 Dec 2037 23:55:55 GMT.",
																	"type":        "string",
																},
																"httpOnly": map[string]interface{}{
																	"description": "Adds the HttpOnly attribute to the cookie.",
																	"type":        "boolean",
																},
																"name": map[string]interface{}{
																	"description": "The name of the cookie.",
																	"type":        "string",
																},
																"path": map[string]interface{}{
																	"description": "The path for which the cookie is set.",
																	"type":        "string",
																},
																"samesite": map[string]interface{}{
																	"description": "Adds the SameSite attribute to the cookie. The allowed values are: strict, lax, none",
																	"type":        "string",
																},
																"secure": map[string]interface{}{
																	"description": "Adds the Secure attribute to the cookie.",
																	"type":        "boolean",
																},
															},
															"type": "object",
														},
														"slow-start": map[string]interface{}{
															"description": "The slow start allows an upstream server to gradually recover its weight from 0 to its nominal value after it has been recovered or became available or when the server becomes available after a period of time it was considered unavailable. By default, the slow start is disabled. Note: The parameter cannot be used along with the random, hash or ip_hash load balancing methods and will be ignored.",
															"type":        "string",
														},
														"subselector": map[string]interface{}{
															"additionalProperties": map[string]interface{}{
																"type": "string",
															},
															"description": "Selects the pods within the service using label keys and values. By default, all pods of the service are selected. Note: the specified labels are expected to be present in the pods when they are created. If the pod labels are updated, NGINX Ingress Controller will not see that change until the number of the pods is changed.",
															"type":        "object",
														},
														"tls": map[string]interface{}{
															"description": "The TLS configuration for the Upstream.",
															"properties": map[string]interface{}{
																"enable": map[string]interface{}{
																	"description": "Enables HTTPS for requests to upstream servers. The default is False , meaning that HTTP will be used. Note: by default, NGINX will not verify the upstream server certificate. To enable the verification, configure an EgressMTLS Policy.",
																	"type":        "boolean",
																},
															},
															"type": "object",
														},
														"type": map[string]interface{}{
															"description": "The type of the upstream. Supported values are http and grpc. The default is http. For gRPC, it is necessary to enable HTTP/2 in the ConfigMap and configure TLS termination in the VirtualServer.",
															"type":        "string",
														},
														"use-cluster-ip": map[string]interface{}{
															"description": "Enables using the Cluster IP and port of the service instead of the default behavior of using the IP and port of the pods. When this field is enabled, the fields that configure NGINX behavior related to multiple upstream servers (like lb-method and next-upstream) will have no effect, as NGINX Ingress Controller will configure NGINX with only one upstream server that will match the service Cluster IP.",
															"type":        "boolean",
														},
													},
													"type": "object",
												},
												"type": "array",
											},
										},
										"type": "object",
									},
									"status": map[string]interface{}{
										"description": "Status contains the current status of the VirtualServer.",
										"properties": map[string]interface{}{
											"externalEndpoints": map[string]interface{}{
												"items": map[string]interface{}{
													"description": "ExternalEndpoint defines the IP/ Hostname and ports used to connect to this resource.",
													"properties": map[string]interface{}{
														"hostname": map[string]interface{}{
															"type": "string",
														},
														"ip": map[string]interface{}{
															"type": "string",
														},
														"ports": map[string]interface{}{
															"type": "string",
														},
													},
													"type": "object",
												},
												"type": "array",
											},
											"message": map[string]interface{}{
												"type": "string",
											},
											"reason": map[string]interface{}{
												"type": "string",
											},
											"state": map[string]interface{}{
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

	return mutate.MutateCRDVirtualserversK8sNginxOrg(resourceObj, parent, collection, reconciler, req)
}
