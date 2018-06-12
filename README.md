= Fabric8 Gateway

== Overview

The fabric8-gateway is a PoC to see how we can automate Gateway API publishing for a bunch of API in services marked with annotation/labels.

- inspired by [K8s Sample Controller](https://github.com/kubernetes/sample-controller)
- following [Best Pactices](https://github.com/kubernetes/community/blob/master/contributors/devel/controllers.md)

For the Gateway API, the integration is done with link:https://github.com/3scale[3scale].

This is very much a work in progress. Therefore a [TODO list](TODO.md).

== Build & Run in dev mode

==== Pre-requisites

Have the following installed on your machine:

* `go` Install link:https://golang.org/dl/[GoLang from 1.9+], set the environment variable `GOPATH`.
* link:https://github.com/golang/dep[Dep the go dependencies mgt], see link:https://golang.github.io/dep/docs/installation.html[isntall guide]
* minishift v1.18+ running with OpenShift v3.9.0+

==== Get the code

```sh
$ git clone https://github.com/corinnekrych/fabric8-gateway $GOPATH/src/github.com/corinnekrych/fabric8-gateway
```

==== Build

```sh
$ rm -rf vendor; dep ensure -v
$ go build ./...
$ go run main.go -kubeconfig=$HOME/.kube/config -project=fabric8
```

== e2e testing

1. start minishift
on macOS
```
 minishift start --vm-driver=xhyve --memory=7000 --cpus=4 --disk-size=50g
```
2. start DeploymentConfig's controller
```
$ go run main.go -kubeconfig=$HOME/.kube/config -project=fabric8
```
3. deploy the REST-service project you used as sample in `fabric8` project.
The project should use `DeploymentConfig`.
See [fabric8-toggle template](https://github.com/corinnekrych/fabric8-toggles-service/blob/d3c2a4843b154ba9b9342dcb41bef8028fc1d10c/toggles-service-os.yml)
Follow [deployment instruction](https://github.com/corinnekrych/fabric8-toggles-service/tree/threescale.trial#docker-build).