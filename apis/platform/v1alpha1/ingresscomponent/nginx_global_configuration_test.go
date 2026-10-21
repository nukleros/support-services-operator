package ingresscomponent

import (
	"testing"

	"github.com/go-logr/logr"
	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"
	"github.com/nukleros/operator-builder-tools/pkg/resources"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/tools/events"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	platformv1alpha1 "github.com/nukleros/support-services-operator/apis/platform/v1alpha1"
	setupv1alpha1 "github.com/nukleros/support-services-operator/apis/setup/v1alpha1"
)

// reconcileStub stands in for a controller so manifest creation runs the mutate hook.
type reconcileStub struct {
	client.Client
}

// GetController returns no controller.
func (reconcileStub) GetController() controller.Controller { return nil }

// GetManager returns no manager.
func (reconcileStub) GetManager() manager.Manager { return nil }

// GetLogger returns a logger that discards every line.
func (reconcileStub) GetLogger() logr.Logger { return logr.Discard() }

// GetResources returns no child resources.
func (reconcileStub) GetResources(*workload.Request) ([]client.Object, error) {
	return nil, nil
}

// GetEventRecorder returns no recorder.
func (reconcileStub) GetEventRecorder() events.EventRecorder { return nil }

// GetFieldManager returns an empty field manager.
func (reconcileStub) GetFieldManager() string { return "" }

// GetWatches returns no watches.
func (reconcileStub) GetWatches() []client.Object { return nil }

// SetWatch records nothing.
func (reconcileStub) SetWatch(client.Object) {}

// CheckReady reports the workload ready.
func (reconcileStub) CheckReady(*workload.Request) (bool, error) {
	return true, nil
}

// Mutate returns the object unchanged.
func (reconcileStub) Mutate(*workload.Request, client.Object) ([]client.Object, bool, error) {
	// return the object unchanged
	return nil, false, nil
}

// TestCreateDeploymentGlobalConfigurationArg puts the global configuration arg on the nginx container.
func TestCreateDeploymentGlobalConfigurationArg(t *testing.T) {
	// build a deployment parent with the name set
	parent := &platformv1alpha1.IngressComponent{}
	parent.Spec.Namespace = "nukleros-ingress-system"
	parent.Spec.Nginx.Include = true
	parent.Spec.Nginx.InstallType = "deployment"
	parent.Spec.Nginx.Image = "nginx/nginx-ingress"
	parent.Spec.Nginx.Version = "5.5.4"
	parent.Spec.Nginx.Replicas = 1
	parent.Spec.Nginx.GlobalConfigurationName = "nginx-ingress"

	// create the deployment the controller would reconcile
	objs, err := CreateDeploymentNamespaceNginxIngress(
		parent,
		&setupv1alpha1.SupportServices{},
		reconcileStub{},
		&workload.Request{},
	)
	if err != nil || len(objs) != 1 {
		t.Fatalf("deployment: %v %v", len(objs), err)
	}
	obj := objs[0]
	workloadObject, ok := obj.(*unstructured.Unstructured)
	if !ok {
		t.Fatalf("deployment must be *unstructured.Unstructured, got %T", obj)
	}
	raw, found, err := unstructured.NestedFieldNoCopy(
		workloadObject.Object, "spec", "template", "spec", "containers",
	)
	containers, containersOK := raw.([]interface{})
	if err != nil || !found || !containersOK || len(containers) == 0 {
		t.Fatalf("containers: found=%v err=%v", found, err)
	}
	container, ok := containers[0].(map[string]interface{})
	if !ok {
		t.Fatalf("container must be map[string]interface{}, got %T", containers[0])
	}
	rawArgs, _, _ := unstructured.NestedFieldNoCopy(container, "args")
	args, _ := rawArgs.([]interface{})
	var foundArg bool
	for _, arg := range args {
		if arg == "-global-configuration=$(POD_NAMESPACE)/nginx-ingress" {
			foundArg = true
		}
	}
	if !foundArg {
		t.Fatalf("global configuration arg: %v", args)
	}

	// an empty name omits the arg, so the update path must see the difference
	parent.Spec.Nginx.GlobalConfigurationName = ""
	without, err := CreateDeploymentNamespaceNginxIngress(
		parent,
		&setupv1alpha1.SupportServices{},
		reconcileStub{},
		&workload.Request{},
	)
	if err != nil || len(without) != 1 {
		t.Fatalf("deployment without name: %v %v", len(without), err)
	}
	equal, err := resources.AreEqual(workloadObject, without[0])
	if err != nil {
		t.Fatal(err)
	}
	if equal {
		t.Fatal("deployment with the arg compares equal to one without it")
	}
}
