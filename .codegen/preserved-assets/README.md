# Support Services Operator

A Kubernetes operator to manage cluster support services.

Support services operator has been built with
[operator-builder](https://github.com/nukleros/operator-builder).  This allows
us to generate all the source code for the project from a set of Kubernetes yaml
manifest for the managed Kubernetes resources with commented markers.  The following quickstart walks through the
process of re-generating the source code and running it to test changes.  See
the [operator-builder
docs](https://github.com/nukleros/operator-builder/tree/main/docs) for more
info.

Make sure you have [operator-builder
installed](https://github.com/nukleros/operator-builder/blob/main/docs/installation.md)
before proceeding.

Use [kind](https://kind.sigs.k8s.io/) to spin up a local Kubernetes cluster for
testing.

## Quickstart

### Generate Source

The manifests that define the managed resources are in the `.codegen` directory.

```bash
cd .codegen
```

There is a Makefile that facilitates common operations.  If you have
operator-builder installed somewhere besides `/usr/local/bin/` set the following
env var:

```bash
export OPERATOR_BUILDER=/path/to/operator-builder
```

Remove the existing source code:

```bash
make operator-clean
```

Initialize a new codebase:

```bash
make operator-init
```

Build the APIs and controller code:

```bash
make operator-create
```

Ensure go dependencies are tidied:

```bash
go mod tidy
```

### Test Operator

Install CRDs:

```bash
make install
```

Run the controller for the support services operator locally.  It will use your
kubeconfig to connect to the Kubernetes API.

```bash
make run
```

There are sample manifests for each custom resource in the `config/samples`
directory.  Create all the support services:

```bash
kubectl apply -f config/samples
```

Check the outcome.  One of the custom resources represents a cert-mangaer
installation.  You can view the spec:

```bash
kubectl get certmanager certmanager-sample -o=jsonpath='{.spec}'
```

You can see the pods that were created as a part of the cert-manager
installation.  Note there are two replicas for each deployment.

```bash
kubectl get po -n nukleros-certs-system
```

Update the certmanager resource.  Set `spec.cainjector.replicas` to 1:

```bash
kubectl edit certmanager certmanager-sample
```

Check the pods again to ensure there is now just one cainjector pod.

```bash
kubectl get po -n nukleros-certs-system
```

### Clean Up

Now let's delete the support services components.  This will remove the various
support services installations.

```bash
kubectl delete externaldns externaldns-sample
kubectl delete externalsecrets externalsecrets-sample
kubectl delete reloader reloader-sample
kubectl delete certmanager certmanager-sample
```

The supportservices resource orchestrates values that need to be shared by
different components.  You can now delete that as well.

```bash
kubectl delete supportservices supportservices-sample
```

You can now stop the controller that you ran with `make run` by hitting Ctrl-C
in that window.

Finally, remove the CRDs:

```bash
make uninstall
```

### Preserve Manually Managed Assets

If you make any changes to files in the codebase, and you want to preserve those
outside of the code generation lifecycle, add that file to the `preserve` and
`restore` make targets defined in `.codegen/Makefile`.  When you delete the
codebase with `make operator-clean` they will automatically be saved.  After
code is generated you can restore them with:

```
make restore
```

## Deploy the Controller Manager

First, set the image:

    export IMG=myrepo/myproject:v0.1.0

Now you can build and push the image:

    make docker-build
    make docker-push

Then deploy:

    make deploy

To clean up:

    make undeploy

## Companion CLI

See the [operator-builder
docs](https://github.com/nukleros/operator-builder/blob/main/docs/companion-cli.md)
for more info on the companion CLI.

To build the companion CLI:

    make build-cli

The CLI binary will get saved to the bin directory.  You can see the help
message with:

    ./bin/ssctl help

