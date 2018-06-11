
package main

import (
	"flag"
	"github.com/golang/glog"
	"k8s.io/client-go/kubernetes"
	"github.com/corinnekrych/fabric8-gateway/common"
	"github.com/corinnekrych/fabric8-gateway/controller"
	"time"
	kubeinformers "k8s.io/client-go/informers"
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

	// Construct the Kubernetes client
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		glog.Fatalf(">> Failed to create kubernetes client: %v", err)
	}

	stopCh := make(chan struct{})
	defer close(stopCh)

	sharedInformerFactory := kubeinformers.WithNamespace(*project)
	kubeInformerFactory := kubeinformers.NewSharedInformerFactoryWithOptions(client, time.Second*30, sharedInformerFactory)
	gc := controller.GatewayController(client, kubeInformerFactory.Apps().V1().Deployments())
	go kubeInformerFactory.Start(stopCh)

	if err = gc.Run(1, stopCh); err != nil {
		glog.Fatalf("Error running controller: %s", err.Error())
	}
}

