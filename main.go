
package main

import (
	"flag"
	"github.com/golang/glog"
	openshiftclient "github.com/openshift/client-go/apps/clientset/versioned"
	openshiftinformers "github.com/openshift/client-go/apps/informers/externalversions"
	"github.com/corinnekrych/fabric8-gateway/controller"
	"time"
	"github.com/corinnekrych/fabric8-gateway/common"
)

var (
	masterURL  string
	kubeconfig string
)

func main() {
	// When running as a pod in-cluster, a kubeconfig is not needed. Instead this will make use of the service
	// account injected into the pod. However, allow the use of a local kubeconfig as this can make local
	// development & testing easier.
	kubeconfig := flag.String("kubeconfig", "", "Path to a kubeconfig file")
	project := flag.String("project", "", "K8s namespace / OS project  to watch")

	// We log to stderr because glog will default to logging to a file.
	// By setting this debugging is easier via `kubectl logs`
	flag.Set("logtostderr", "true")
	flag.Parse()
	glog.Infof("--> Starting gateway-controller with config %s", *kubeconfig)

	// Build the client config - optionally using a provided kubeconfig file.
	config, err := common.GetClientConfig(*kubeconfig)
	if err != nil {
		glog.Fatalf("Failed to load client config: %v", err)
	}
	glog.Infof("--> Running on host %s", config.Host)

	// Construct the openshift client
	client, err := openshiftclient.NewForConfig(config)
	if err != nil {
		glog.Fatalf(">> Failed to create kubernetes client: %v", err)
	}

	stopCh := make(chan struct{})
	defer close(stopCh)

	sharedInformerFactory := openshiftinformers.NewFilteredSharedInformerFactory(client, time.Second*30, *project, nil)
	gc := controller.GatewayController(client, sharedInformerFactory.Apps().V1().DeploymentConfigs())
	go sharedInformerFactory.Start(stopCh)

	if err = gc.Run(1, stopCh); err != nil {
		glog.Fatalf("Error running controller: %s", err.Error())
	}
}

