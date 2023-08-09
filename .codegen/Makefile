OPERATOR_BUILDER ?= /usr/local/bin/operator-builder

#preserve: @ Preserve manually mananged assets
preserve:
	cp -R ../config/install preserved-assets/
	cp ../README.md preserved-assets/
	cp ../Makefile preserved-assets/
	cp ../Dockerfile preserved-assets/
	cp ../Dockerfile.old preserved-assets/
	cp ../go.mod preserved-assets/
	cp ../config/samples/gateway_v1alpha1_glooedge.yaml preserved-assets/
	cp ../apis/gateway/v1alpha1/glooedge_types.go preserved-assets/glooedge_types.go.bak
	cp ../apis/gateway/v1alpha1/externaldns_types.go preserved-assets/externaldns_types.go.bak
	cp ../apis/gateway/v1alpha1/glooedge/mutate/gateway_namespace_gateway_proxy.go preserved-assets/gateway_namespace_gateway_proxy.go.bak
	cp ../apis/gateway/v1alpha1/externaldns/mutate/deployment.go preserved-assets/deployment.go.bak
	cp ../apis/gateway/v1alpha1/externaldns/mutate/deployment_namespace_external_dns_route_US.go preserved-assets/deployment_namespace_external_dns_route_US.go.bak
	cp ../apis/gateway/v1alpha1/externaldns/mutate/deployment_namespace_external_dns_google.go preserved-assets/deployment_namespace_external_dns_google.go.bak

#operator-clean: @ Preserve manually managed assets and then remove all operator source code to prep for re-generation
operator-clean: preserve
	rm -rf ../*

#operator-init: @ Initialize codebase
operator-init:
	cd .. && $(OPERATOR_BUILDER) init \
		--workload-config .codegen/workload.yaml \
		--repo github.com/nukleros/support-services-operator \
		--controller-image nukleros/support-services-operator:latest \
		--skip-go-version-check

#operator-create: @ Create APIs and controller source code and restore manually managed assets
operator-create:
	cd .. && $(OPERATOR_BUILDER) create api \
		--workload-config .codegen/workload.yaml \
		--controller \
		--resource
	cp -R preserved-assets/install ../config
	cp preserved-assets/README.md ../
	cp preserved-assets/Makefile ../
	cp preserved-assets/Dockerfile ../
	cp preserved-assets/Dockerfile.old ../
	cp preserved-assets/go.mod ../
	cp preserved-assets/gateway_v1alpha1_glooedge.yaml ../config/samples/gateway_v1alpha1_glooedge.yaml
	cp preserved-assets/glooedge_types.go.bak ../apis/gateway/v1alpha1/glooedge_types.go
	cp preserved-assets/externaldns_types.go.bak ../apis/gateway/v1alpha1/externaldns_types.go
	cp preserved-assets/gateway_namespace_gateway_proxy.go.bak ../apis/gateway/v1alpha1/glooedge/mutate/gateway_namespace_gateway_proxy.go
	cp preserved-assets/deployment.go.bak ../apis/gateway/v1alpha1/externaldns/mutate/deployment.go
	cp preserved-assets/deployment_namespace_external_dns_route_US.go.bak ../apis/gateway/v1alpha1/externaldns/mutate/deployment_namespace_external_dns_route_US.go
	cp preserved-assets/deployment_namespace_external_dns_google.go.bak ../apis/gateway/v1alpha1/externaldns/mutate/deployment_namespace_external_dns_google.go
	go mod tidy

