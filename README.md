# Pocket Operator

Deploy and manage pocket nodes on Kubernetes clusters.

This operator was built with
[operator-builder](https://github.com/nukleros/operator-builder).

## Local Development & Testing

To install the custom resource/s for this operator, make sure you have a
kubeconfig set up for a test cluster, then install the CRDs.

```bash
make install
```

To run the controller locally against your test cluster.

```bash
make run
```

You can then test the operator by creating the sample manifests in another
terminal.

```bash
kubectl apply -f config/samples
```

Give the validators some time to come up.  The validator nodes will go into a
crashloop.  First because the database pod is not yet up, and then because it is
failing password authentication when connecting to the DB.  This is a issue that
will be fixed soon.  In the meantime, we have a workaround.

Stop the controller by hitting Ctrl-C in the terminal where you ran `make run`.
Then run the following commands.

```bash
cd .operator-builder/hack
for v in $(cat validators); do ./update-db-pass.sh $v; done
for v in $(cat validators); do kubectl delete po $v-0; done
```

Run the dev client in the pocket client container.  See the [pocket development
docs](https://github.com/pokt-network/pocket/tree/main/docs/development#running-localnet)
for more info on using the dev client.

```bash
kubectl exec -it pocket-v1-client -- go run app/client/main.go
```

Finally, once testing is complete you can clean up.  Note: you will need to run
the controller to finalize resource deletion with `make run` if you stopped it.

```bash
make uninstall
```

## Code Generation

The following steps will re-generate the codebase from scratch after making
changes to the configurations, source manifests and/or markers.

*Caution*: These steps will permanently delete any changes you have made
directly to the codebase.

Delete the existing codebase.

```bash
cd .operator-builder
make operator-clean
```

Re-build the codebase from the existing configurations and source manifests.

```bash
make operator-init
make operator-create
```

Install the dependencies (postgres operator) in your test cluster.

```bash
make operator-dependencies
```

Copy the modified sample manifests into the `config/samples` directory of the
codebase.

```bash
make operator-samples
cd ../
```

You can now re-test the operator using the [Local Development & Testing
instructions](#local-development-&-testing) above.

## Deploy the Controller Manager

First, set the image name.

```bash
export IMG=myrepo/myproject:v0.1.0
```

Now you can build and push the image.

```bash
make docker-build
make docker-push
```

Then deploy.

```bash
make deploy
```

To clean up.

```bash
make undeploy
```

## Companion CLI

To build the companion CLI.

```bash
make build-cli
```

The CLI binary will get saved to the bin directory.  You can see the help
message with the following.

```bash
./bin/pocketctl help
```
