package templaterouter

import (
	"fmt"
	"log"

	"github.com/kubernetes-sigs/kubebuilder/pkg/controller"
	"github.com/kubernetes-sigs/kubebuilder/pkg/controller/types"

	routerv1alpha1 "github.com/knobunc/router-operator/pkg/apis/router/v1alpha1"
	routerv1alpha1client "github.com/knobunc/router-operator/pkg/client/clientset/versioned/typed/router/v1alpha1"
	routerv1alpha1informer "github.com/knobunc/router-operator/pkg/client/informers/externalversions/router/v1alpha1"
	routerv1alpha1lister "github.com/knobunc/router-operator/pkg/client/listers/router/v1alpha1"
	"github.com/knobunc/router-operator/pkg/inject/args"

	originappsv1 "github.com/openshift/client-go/apps/clientset/versioned/typed/apps/v1"

	appsapi "github.com/openshift/api/apps/v1"

	kapi "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// EDIT THIS FILE
// This files was created by "kubebuilder create resource" for you to edit.
// Controller implementation logic for TemplateRouter resources goes here.

func (bc *TemplateRouterController) Reconcile(k types.ReconcileKey) error {
	// INSERT YOUR CODE HERE
	log.Printf("Reconciling templaterouter.TemplateRouterController %s.%s", k.Namespace, k.Name)

	// Pull the template router CRD object
	tr, err := bc.templaterouterLister.TemplateRouters(k.Namespace).Get(k.Name)
	if err != nil {
		// XXX Had an error, assume it wasn't there and delete the dc
		log.Printf("No tr, Deleting")

		deleteOptions := metav1.DeleteOptions{}
		if err := bc.dcClient.DeploymentConfigs(k.Namespace).Delete(k.Name, &deleteOptions); err != nil {
			log.Panicf("Couldn't delete dc %v", err)
		}
	} else {
		// Create or update
		log.Printf("Got tr, Creating or Updating")
		tr.SetDefaults()

		// Pull the DC
		dc, err := bc.dcClient.DeploymentConfigs(k.Namespace).Get(k.Name, metav1.GetOptions{})
		createDC := false
		if err != nil {
			// Assume it doesn't exist and just make it
			// TODO the right thing :-)
			log.Printf("Need to make a new dc")
			dc = newDeploymentConfig()
			createDC = true
		} else {
			log.Printf("Read dc")
		}

		if updateDC(tr, dc, k) {
			// log.Printf("New dc: %#v\n\n", dc)
			// log.Printf("New template: %#v\n\n", dc.Spec.Template)
			if createDC {
				if _, err := bc.dcClient.DeploymentConfigs(k.Namespace).Create(dc); err != nil {
					log.Panicf("Couldn't create dc %v", err)
				}
				log.Printf("Created new dc")
			} else {
				if _, err := bc.dcClient.DeploymentConfigs(k.Namespace).Update(dc); err != nil {
					log.Panicf("Couldn't update dc %v", err)
				}
				log.Printf("Updated dc")
			}
		}
	}

	return nil
}

// +controller:group=router,version=v1alpha1,kind=TemplateRouter,resource=templaterouters
type TemplateRouterController struct {
	// INSERT ADDITIONAL FIELDS HERE
	templaterouterLister routerv1alpha1lister.TemplateRouterLister
	templaterouterclient routerv1alpha1client.RouterV1alpha1Interface
	dcClient             originappsv1.DeploymentConfigsGetter
}

// ProvideController provides a controller that will be run at startup.  Kubebuilder will use codegeneration
// to automatically register this controller in the inject package
func ProvideController(arguments args.InjectArgs) (*controller.GenericController, error) {
	// INSERT INITIALIZATIONS FOR ADDITIONAL FIELDS HERE
	bc := &TemplateRouterController{
		templaterouterLister: arguments.ControllerManager.GetInformerProvider(&routerv1alpha1.TemplateRouter{}).(routerv1alpha1informer.TemplateRouterInformer).Lister(),
		templaterouterclient: arguments.Clientset.RouterV1alpha1(),
		dcClient:             arguments.OriginAppsClientset.AppsV1(),
	}

	// Create a new controller that will call TemplateRouterController.Reconcile on changes to TemplateRouters
	gc := &controller.GenericController{
		Name:             "TemplateRouterController",
		Reconcile:        bc.Reconcile,
		InformerRegistry: arguments.ControllerManager,
	}
	if err := gc.Watch(&routerv1alpha1.TemplateRouter{}); err != nil {
		return gc, err
	}

	// INSERT ADDITIONAL WATCHES HERE BY CALLING gc.Watch.*() FUNCTIONS
	// NOTE: Informers for Kubernetes resources *MUST* be registered in the pkg/inject package so that they are started.
	return gc, nil
}

func newDeploymentConfig() *appsapi.DeploymentConfig {
	maxUnavailable := intstr.FromString("25%")
	return &appsapi.DeploymentConfig{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "",
			Namespace: "",
		},
		Spec: appsapi.DeploymentConfigSpec{
			Replicas: 1,
			Selector: map[string]string{"a": "XXX"},
			Strategy: appsapi.DeploymentStrategy{
				Type:      appsapi.DeploymentStrategyTypeRolling,
				Resources: kapi.ResourceRequirements{},
				RollingParams: &appsapi.RollingDeploymentStrategyParams{
					MaxUnavailable: &maxUnavailable,
				},
			},
			Template: &kapi.PodTemplateSpec{
				Spec: kapi.PodSpec{
					Containers: []kapi.Container{
						newContainerSpec(),
					},
					RestartPolicy:   kapi.RestartPolicyAlways,
					DNSPolicy:       kapi.DNSClusterFirst,
					SchedulerName:   kapi.DefaultSchedulerName,
					HostNetwork:     true, // XXX
					SecurityContext: &kapi.PodSecurityContext{},
				},
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{"a": "XXX"},
				},
			},
			Triggers: []appsapi.DeploymentTriggerPolicy{
				appsapi.DeploymentTriggerPolicy{
					Type: appsapi.DeploymentTriggerOnConfigChange,
				},
			},
		},
	}
}

func newContainerSpec() kapi.Container {
	return kapi.Container{
		Name:  "router",
		Image: "openshift/origin-haproxy-router:v3.9.0-alpha.4",
		Env: []kapi.EnvVar{
			{
				Name:  "ENV1",
				Value: "VAL1",
			},
		},
		ImagePullPolicy:          kapi.PullIfNotPresent,
		TerminationMessagePath:   "/dev/termination-log",
		TerminationMessagePolicy: kapi.TerminationMessageReadFile,
		Resources: kapi.ResourceRequirements{
			Requests: kapi.ResourceList{
				kapi.ResourceCPU:    resource.MustParse("100m"),
				kapi.ResourceMemory: resource.MustParse("256Mi"),
			},
		},
	}
}

func mkintp(i int) *int64 {
	v := int64(i)
	return &v
}

// SetDefaults sets the default vaules for the TemplateRouter spec and returns true if the spec was changed
func updateDC(t *routerv1alpha1.TemplateRouter, dc *appsapi.DeploymentConfig, k types.ReconcileKey) bool {
	changed := false

	dcm := &dc.ObjectMeta
	if dcm.Name != k.Name {
		log.Printf("Name changed: '%s' to '%s'", dcm.Name, k.Name)
		dcm.Name = k.Name
		changed = true
	}
	if dcm.Namespace != k.Namespace {
		log.Printf("Namespace changed: '%s' to '%s'", dcm.Namespace, k.Namespace)
		dcm.Namespace = k.Namespace
		changed = true
	}

	ts := &t.Spec
	dcs := &dc.Spec
	if dcs.Replicas != ts.Replicas {
		log.Printf("Replicas changed: '%d' to '%d'", dcs.Replicas, ts.Replicas)
		dcs.Replicas = ts.Replicas
		changed = true
	}

	if len(dcs.Template.Spec.Containers) != 1 {
		log.Printf("Container count changed: %d to 1", len(dcs.Template.Spec.Containers))
		dcs.Template.Spec.Containers = []kapi.Container{
			newContainerSpec(),
		}
		changed = true
	}
	c := dcs.Template.Spec.Containers[0]

	image := fmt.Sprintf("%s:%s", ts.BaseImage, ts.Version)
	if c.Image != image {
		log.Printf("Image changed: '%s' to '%s'", c.Image, image)
		c.Image = image
		changed = true
	}

	dcs.Template.Spec.Containers[0] = c

	return changed
}
