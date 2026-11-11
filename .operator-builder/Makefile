#
# code generation
#
INIT_OPTS=init \
	--workload-config .operator-builder/setup.addons.nukleros.io/workload.yaml \
   	--repo github.com/nukleros/support-services-operator \
	--controller-image nukleros/support-services-operator:latest \
    --skip-go-version-check

CREATE_OPTS=create api \
	--workload-config .operator-builder/setup.addons.nukleros.io/workload.yaml \
	--controller \
	--resource

operator-clean: preserve
	rm -rf ../*

# WARN: init/create only need to be run once.  use update to update the operator.
operator-init:
	cd .. && operator-builder $(INIT_OPTS)

operator-create:
	cd .. && operator-builder $(CREATE_OPTS)

operator-update:
	cd .. && operator-builder $(CREATE_OPTS) --force

#
# project management
#
# projects belong to a component when KIND is set, or to the collection when it is not:
#
#   make overlay GROUP=platform KIND=ingress-component PROJECT=kong  # platform.addons.nukleros.io/ingress-component/kong
#   make overlay GROUP=setup PROJECT=gateway-api                     # setup.addons.nukleros.io/gateway-api
#
GROUP ?=
DOMAIN ?= addons.nukleros.io
KIND ?=
PROJECT ?=

WORKLOAD_DIR = $(GROUP).$(DOMAIN)$(if $(KIND),/$(KIND))
PROJECT_DIR = $(WORKLOAD_DIR)/$(PROJECT)
CONFIG_DIR = $(PROJECT_DIR)/config

# fail fast when a required variable is missing (e.g. require-GROUP)
require-%:
	@[ -n "$($*)" ] || (echo "ERROR: $* is required" && exit 1)

PROJECT_REQUIREMENTS = require-GROUP require-PROJECT

# NOTE: the collection namespace already exists, so a namespace is only generated for new components
project: $(PROJECT_REQUIREMENTS)
	@mkdir -p $(CONFIG_DIR) $(PROJECT_DIR)/vendor $(PROJECT_DIR)/manifests
	@touch $(CONFIG_DIR)/values.yaml $(CONFIG_DIR)/vendor.yaml $(CONFIG_DIR)/overlays.yaml
	@$(if $(KIND),[ -f $(WORKLOAD_DIR)/namespace.yaml ] || kubectl create ns nukleros-$(KIND)-system --dry-run=client -o yaml > $(WORKLOAD_DIR)/namespace.yaml)

download: $(PROJECT_REQUIREMENTS)
	vendir sync \
		--file $(CONFIG_DIR)/vendor.yaml \
		--lock-file $(CONFIG_DIR)/vendor.yaml.lock

overlay: $(PROJECT_REQUIREMENTS)
	yot \
		--indent-level 2 \
		--instructions $(CONFIG_DIR)/overlays.yaml \
		--output-directory $(PROJECT_DIR) \
		--values-file $(CONFIG_DIR)/values.yaml
