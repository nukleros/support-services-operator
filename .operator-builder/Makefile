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
DOMAIN ?=
GROUP ?=
KIND ?=
PROJECT ?=
project:
	@mkdir -p $(GROUP).$(DOMAIN)/$(KIND)/$(PROJECT)/config $(GROUP).$(DOMAIN)/$(KIND)/$(PROJECT)/vendor $(GROUP).$(DOMAIN)/$(KIND)/$(PROJECT)/manifests
	@touch $(GROUP).$(DOMAIN)/$(KIND)/$(PROJECT)/config/values.yaml $(GROUP).$(DOMAIN)/$(KIND)/$(PROJECT)/config/vendor.yaml $(GROUP).$(DOMAIN)/$(KIND)/$(PROJECT)/config/overlays.yaml

download:
	vendir sync \
		--file $(GROUP).$(DOMAIN)/$(KIND)/$(PROJECT)/config/vendor.yaml \
		--lock-file $(GROUP).$(DOMAIN)/$(KIND)/$(PROJECT)/config/vendor.yaml.lock

overlay:
	yot \
		--indent-level 2 \
		--instructions $(GROUP).$(DOMAIN)/$(KIND)/$(PROJECT)/config/overlays.yaml \
		--output-directory $(GROUP).$(DOMAIN)/$(KIND)/$(PROJECT) \
		--values-file $(GROUP).$(DOMAIN)/$(KIND)/$(PROJECT)/config/values.yaml
