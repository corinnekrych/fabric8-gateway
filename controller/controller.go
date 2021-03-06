package controller

import (
	openshfitclient "github.com/openshift/client-go/apps/clientset/versioned"
	"github.com/openshift/api/apps/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/apimachinery/pkg/util/runtime"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	appslisters "github.com/openshift/client-go/apps/listers/apps/v1"
	appsinformers "github.com/openshift/client-go/apps/informers/externalversions/apps/v1"
	"github.com/golang/glog"
	"k8s.io/apimachinery/pkg/util/wait"
	"time"
	"fmt"
)

// Controller is the controller watching Deployment
type Controller struct {
	// client is a standard kubernetes client
	client openshfitclient.Interface
	// threescaleClient is the client API for 3scale REST API
	//threescaleClient clientset.Interface

	deploymentsLister appslisters.DeploymentConfigLister
	deploymentsSynced cache.InformerSynced

	// workqueue is a rate limited work queue.
	workqueue workqueue.RateLimitingInterface
	//recorder record.EventRecorder
}

// GatewayController returns a new controller that watch deployment and genrate Gateway API.
func GatewayController(
	client openshfitclient.Interface,
	//threescaleClient clientset.Interface,

	deploymentInformer appsinformers.DeploymentConfigInformer) *Controller {

	controller := &Controller{
		client:     client,
		//threescaleClient:   threescaleClient,
		deploymentsLister: deploymentInformer.Lister(),
		deploymentsSynced: deploymentInformer.Informer().HasSynced,
		workqueue:         workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "GatewayQueue"),
	}

	glog.Info("--> Instantiate Setting up event handlers")
	// Set up an event handler for when Deployment resources change.
	deploymentInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: controller.handleObject,
		// UpdateFunc is called when an existing resource is modified.
		// UpdateFunc is also called when a re-synchronization happens, and it gets called even if nothing changes.
		UpdateFunc: func(old, new interface{}) {
			//glog.Info("--> GatewayController::deploymentInformer::UpdateFunc\n")
			newDepl := new.(*v1.DeploymentConfig)
			oldDepl := old.(*v1.DeploymentConfig)
			var foundDeploymentFinishedCondition *v1.DeploymentCondition
			if newDepl.ResourceVersion == oldDepl.ResourceVersion { // no changes, do nothing
				glog.Infof("--> GatewayController::deploymentInformer::Same ResourceRevision %s of DeploymentConfig of %s with version %d ...\n", oldDepl.ResourceVersion, oldDepl.Name, newDepl.Status.LatestVersion)
				return
			} else { // there are some changes in DC
				deploymentConditions := newDepl.Status.Conditions
				for _, v := range deploymentConditions {
					if v.Reason == "NewReplicationControllerAvailable" {
						foundDeploymentFinishedCondition = &v
					}
				}
				// if the message is sent that "replication controller successfully rolled out" and this is really about the latest generation.
				if foundDeploymentFinishedCondition != nil && newDepl.ObjectMeta.Generation == newDepl.Status.ObservedGeneration {
					glog.Infof("--> GatewayController::deploymentInformer::Going to version %d with generation %d of %s\n", newDepl.Status.LatestVersion, newDepl.ObjectMeta.Generation, newDepl.ObjectMeta.Name)
					glog.Infof("--> NEW version of DC is %#v \n", newDepl)
					controller.handleObject(new)
				}
			}


		},
		DeleteFunc: controller.handleObject,
	})

	return controller
}

// Run will set up the event handlers for types we are interested in, as well
// as syncing informer caches and starting workers. It will block until stopCh
// is closed, at which point it will shutdown the workqueue and wait for
// workers to finish processing their current work items.
func (c *Controller) Run(threadiness int, stopCh <-chan struct{}) error {
	glog.Info("--> Setting up event handlers")
	defer runtime.HandleCrash()
	defer c.workqueue.ShutDown()

	// Start the informer factories to begin populating the informer caches
	glog.Info("--> Starting Gateway Controller")

	// Wait for the caches to be synced before starting workers
	glog.Info("Waiting for informer caches to sync")
	if ok := cache.WaitForCacheSync(stopCh, c.deploymentsSynced); !ok {
		return fmt.Errorf("failed to wait for caches to sync")
	}

	glog.Info("--> Starting workers")

	// Launch workers to process Foo resources
	for i := 0; i < threadiness; i++ {
		go wait.Until(c.runWorker, time.Second, stopCh)
	}

	glog.Info("--> Started workers")
	<-stopCh
	glog.Info("---> Shutting down workers")

	return nil
}

// runWorker is a long-running function that will continually call the
// processNextWorkItem function in order to read and process a message on the
// workqueue.
func (c *Controller) runWorker() {
	for c.processNextWorkItem() {
	}
}

// processNextWorkItem will read a single work item off the workqueue and
// attempt to process it, by calling the syncHandler.
func (c *Controller) processNextWorkItem() bool {
	glog.Info("--> GatewayController::processNextWorkItem\n")
	obj, shutdown := c.workqueue.Get()
	glog.Infof("--> GatewayController::processNextWorkItem found it %s!!!!!!\n", shutdown)

	if shutdown {
		return false
	}

	// We wrap this block in a func so we can defer c.workqueue.Done.
	err := func(obj interface{}) error {
		defer c.workqueue.Done(obj)
		var key string
		var ok bool
		if key, ok = obj.(string); !ok {
			c.workqueue.Forget(obj)
			runtime.HandleError(fmt.Errorf("expected string in workqueue but got %#v", obj))
			return nil
		}
		// Run the syncHandler, passing it the namespace/name string of the
		if err := c.syncHandler(key); err != nil {
			return fmt.Errorf("error syncing '%s': %s", key, err.Error())
		}
		c.workqueue.Forget(obj)
		glog.Infof("Successfully synced '%s'", key)
		return nil
	}(obj)

	if err != nil {
		runtime.HandleError(err)
		return true
	}

	return true
}

// handleObject is triggered either by an Add/Update/Delete
// on Deployment. If the Service/s deployment got the label/annotation threeScale// we enqueue the
func (c *Controller) handleObject(obj interface{}) {
	glog.Info("--> In Controller::handleObject\n")
	var object metav1.Object
	var ok bool
	if object, ok = obj.(metav1.Object); !ok {
		tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
		if !ok {
			runtime.HandleError(fmt.Errorf("error decoding object, invalid type"))
			return
		}
		object, ok = tombstone.Obj.(metav1.Object)
		if !ok {
			runtime.HandleError(fmt.Errorf("error decoding object tombstone, invalid type"))
			return
		}
		glog.Infof("Recovered deleted object '%s' from tombstone", object.GetName())
	}
	glog.Infof("--------> Processing object: %s", object.GetName())

	// TODO
	//c.workqueue.AddRateLimited(obj)

}

// syncHandler hold the boen of the work to be doen as part of a sync loop
func (c *Controller) syncHandler(key string) error {
	glog.Info("--> In Controller::handleObject\n")
	return nil
}