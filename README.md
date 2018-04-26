This project creates a simple router operator for OpenShift to read a
custom resource definition that defines a desired router state and
then creates the objects needed to make that state exist.

Building
--------

`go build cmd/controller-manager/main.go`


Installing
----------

The operator creates its own CRD if there is not one already there.
Eventually we'll need to make one manually since the operator should
not have privilege to do that... but that can wait.


Running
-------

`./main -kubeconfig $KUBECONFIG`

(Eventually this needs to be made into an image, but I haven't got there yet)


Using
-----

Create a `templaterouter` object:

```yaml
apiVersion: "router.operations.openshift.io/v1alpha1"
kind: "TemplateRouter"
metadata:
  name: "whee"
spec:
  replicas: 1
```

Or just use the example one: `oc create -f example/template_router.yaml`

Then you should see a deployment config get created:

```bash
$ oc get dc whee
NAME      REVISION   DESIRED   CURRENT   TRIGGERED BY
whee      1          1         0         config
```

Then if you edit the `templaterouter` object and set the `replicas` to 2 you will see that get reflected in the deployment congfig:

```bash
$ oc patch templaterouter whee -p '{"spec":{"replicas":2}}' --type=merge
templaterouter.router.operations.openshift.io "whee" patched

$ oc get dc whee
NAME      REVISION   DESIRED   CURRENT   TRIGGERED BY
whee      1          2         2         config
```

Then, we can scale down to 0 the same way.  Or if we delete the
`templaterouter` object then it will clean up the items that were
created:

```bash
$ oc delete templaterouter whee
templaterouter.router.operations.openshift.io "whee" deleted

$ oc get dc whee
Error from server (NotFound): deploymentconfigs.apps.openshift.io "whee" not found
```
