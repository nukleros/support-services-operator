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

// +kubebuilder:rbac:groups=core,resources=configmaps,verbs=get;list;watch;create;update;patch;delete

// CreateConfigMapNamespaceGatewayProxyEnvoyConfig creates the ConfigMap resource with name gateway-proxy-envoy-config.
func CreateConfigMapNamespaceGatewayProxyEnvoyConfig(
	parent *gatewayv1alpha1.GlooEdge,
	collection *orchestrationv1alpha1.SupportServices,
	reconciler workload.Reconciler,
	req *workload.Request,
) ([]client.Object, error) {

	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "ConfigMap",
			"metadata": map[string]interface{}{
				"name":      "gateway-proxy-envoy-config",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
				"labels": map[string]interface{}{
					"app":              "gloo",
					"gloo":             "gateway-proxy",
					"gateway-proxy-id": "gateway-proxy",
				},
			},
			"data": map[string]interface{}{
				// controlled by field: namespace
				"envoy.yaml": `layered_runtime:
  layers:
  - name: static_layer
    static_layer:
      overload:
        global_downstream_max_connections: 250000
      upstream:
        healthy_panic_threshold:
          value: 50
  - name: admin_layer
    admin_layer: {}
node:
  cluster: gateway
  id: "{{.PodName}}.{{.PodNamespace}}"
  metadata:
    # Specifies the proxy's in-memory xds cache key (see projects/gloo/pkg/xds/envoy.go)
    # This value needs to match discoveryNamespace (or "writeNamespace") in the settings template
    role: ` + parent.Spec.Namespace + `~gateway-proxy
static_resources:
  listeners:
    - name: prometheus_listener
      address:
        socket_address:
          address: 0.0.0.0
          port_value: 8081
      filter_chains:
        - filters:
            - name: envoy.filters.network.http_connection_manager
              typed_config:
                "@type": type.googleapis.com/envoy.extensions.filters.network.http_connection_manager.v3.HttpConnectionManager
                codec_type: AUTO
                stat_prefix: prometheus
                route_config:
                  name: prometheus_route
                  virtual_hosts:
                    - name: prometheus_host
                      domains:
                        - "*"
                      routes:
                        - match:
                            path: "/ready"
                            headers:
                            - name: ":method"
                              exact_match: GET
                          route:
                            cluster: admin_port_cluster
                        - match:
                            prefix: "/metrics"
                            headers:
                            - name: ":method"
                              exact_match: GET
                          route:
                            prefix_rewrite: /stats/prometheus
                            cluster: admin_port_cluster
                http_filters:
                  - name: envoy.filters.http.router
                    typed_config:
                      "@type": type.googleapis.com/envoy.extensions.filters.http.router.v3.Router
  clusters:
  - name: gloo.` + parent.Spec.Namespace + `.svc.cluster.local:9977
    alt_stat_name: xds_cluster
    connect_timeout: 5.000s
    load_assignment:
      cluster_name: gloo.` + parent.Spec.Namespace + `.svc.cluster.local:9977
      endpoints:
      - lb_endpoints:
        - endpoint:
            address:
              socket_address:
                address: gloo.` + parent.Spec.Namespace + `.svc.cluster.local
                port_value: 9977
    http2_protocol_options: {}
    upstream_connection_options:
      tcp_keepalive:
        keepalive_time: 60
    type: STRICT_DNS
    respect_dns_ttl: true
  - name: rest_xds_cluster
    alt_stat_name: rest_xds_cluster
    connect_timeout: 5.000s
    load_assignment:
      cluster_name: rest_xds_cluster
      endpoints:
      - lb_endpoints:
        - endpoint:
            address:
              socket_address:
                address: gloo.` + parent.Spec.Namespace + `.svc.cluster.local
                port_value: 9976
    upstream_connection_options:
      tcp_keepalive:
        keepalive_time: 60
    type: STRICT_DNS
    respect_dns_ttl: true
  - name: wasm-cache
    connect_timeout: 5.000s
    load_assignment:
      cluster_name: wasm-cache
      endpoints:
      - lb_endpoints:
        - endpoint:
            address:
              socket_address:
                address: gloo.` + parent.Spec.Namespace + `.svc.cluster.local
                port_value: 9979
    upstream_connection_options:
      tcp_keepalive:
        keepalive_time: 60
    type: STRICT_DNS
    respect_dns_ttl: true
  - name: admin_port_cluster
    connect_timeout: 5.000s
    type: STATIC
    lb_policy: ROUND_ROBIN
    load_assignment:
      cluster_name: admin_port_cluster
      endpoints:
      - lb_endpoints:
        - endpoint:
            address:
              socket_address:
                address: 127.0.0.1
                port_value: 19000

dynamic_resources:
  ads_config:
    transport_api_version: V3
    api_type: GRPC
    rate_limit_settings: {}
    grpc_services:
    - envoy_grpc: {cluster_name: gloo.` + parent.Spec.Namespace + `.svc.cluster.local:9977}
  cds_config:
    resource_api_version: V3
    ads: {}
  lds_config:
    resource_api_version: V3
    ads: {}
admin:
  access_log_path: /dev/null
  address:
    socket_address:
      address: 127.0.0.1
      port_value: 19000
`,
			},
		},
	}

	return mutate.MutateConfigMapNamespaceGatewayProxyEnvoyConfig(resourceObj, parent, collection, reconciler, req)
}
