#
# code generation
#
INIT_OPTS=init \
	--workload-config .operator-builder/workload.yaml \
   	--repo github.com/nukleros/support-services-operator \
	--controller-image nukleros/support-services-operator:latest \
    --skip-go-version-check

CREATE_OPTS=create api \
	--workload-config .operator-builder/workload.yaml \
	--controller \
	--resource

operator-clean:
	rm -rf ../*

operator-init:
	cd .. && operator-builder $(INIT_OPTS)

operator-create:
	cd .. && operator-builder $(CREATE_OPTS)

#
# project management
#
DOMAIN ?=
GROUP ?=
KIND ?=
PROJECT ?=
project:
	@mkdir -p $(GROUP).$(DOMAIN)/$(KIND)/$(PROJECT)/config $(GROUP).$(DOMAIN)/$(KIND)/$(PROJECT)/vendor $(GROUP).$(DOMAIN)/$(KIND)/$(PROJECT)/manifests $(GROUP).$(DOMAIN)/$(KIND)/$(PROJECT)/static
	@touch $(GROUP).$(DOMAIN)/$(KIND)/$(PROJECT)/config/values.yaml $(GROUP).$(DOMAIN)/$(KIND)/$(PROJECT)/config/vendor.yaml $(GROUP).$(DOMAIN)/$(KIND)/$(PROJECT)/config/overlays.yaml

download:
	vendir sync --file $(GROUP).$(DOMAIN)/$(KIND)/$(PROJECT)/config/vendor.yaml --lock-file $(GROUP).$(DOMAIN)/$(KIND)/$(PROJECT)/config/vendor.yaml.lock

overlay:
	#yot -I 2 -i $(GROUP).$(DOMAIN)/$(KIND)/$(PROJECT)/config/overlays.yaml -o $(GROUP).$(DOMAIN)/$(KIND)/$(PROJECT)/manifests -f $(GROUP).$(DOMAIN)/$(KIND)/$(PROJECT)/config/values.yaml
	yot -I 2 -i $(GROUP).$(DOMAIN)/$(KIND)/$(PROJECT)/config/overlays.yaml -o $(GROUP).$(DOMAIN)/$(KIND)/$(PROJECT)/manifests -f $(GROUP).$(DOMAIN)/$(KIND)/$(PROJECT)/config/values.yaml

