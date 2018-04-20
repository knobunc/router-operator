

package templaterouter

import (
    "log"

    "github.com/kubernetes-sigs/kubebuilder/pkg/controller"
    "github.com/kubernetes-sigs/kubebuilder/pkg/controller/types"

    routerv1alpha1client "github.com/knobunc/router-operator/pkg/client/clientset/versioned/typed/router/v1alpha1"
    routerv1alpha1lister "github.com/knobunc/router-operator/pkg/client/listers/router/v1alpha1"
    routerv1alpha1 "github.com/knobunc/router-operator/pkg/apis/router/v1alpha1"
    routerv1alpha1informer "github.com/knobunc/router-operator/pkg/client/informers/externalversions/router/v1alpha1"
    "github.com/knobunc/router-operator/pkg/inject/args"
)

// EDIT THIS FILE
// This files was created by "kubebuilder create resource" for you to edit.
// Controller implementation logic for TemplateRouter resources goes here.

func (bc *TemplateRouterController) Reconcile(k types.ReconcileKey) error {
    // INSERT YOUR CODE HERE
    log.Printf("Implement the Reconcile function on templaterouter.TemplateRouterController to reconcile %s\n", k.Name)
    return nil
}

// +controller:group=router,version=v1alpha1,kind=TemplateRouter,resource=templaterouters
type TemplateRouterController struct {
    // INSERT ADDITIONAL FIELDS HERE
    templaterouterLister routerv1alpha1lister.TemplateRouterLister
    templaterouterclient routerv1alpha1client.RouterV1alpha1Interface
}

// ProvideController provides a controller that will be run at startup.  Kubebuilder will use codegeneration
// to automatically register this controller in the inject package
func ProvideController(arguments args.InjectArgs) (*controller.GenericController, error) {
    // INSERT INITIALIZATIONS FOR ADDITIONAL FIELDS HERE
    bc := &TemplateRouterController{
        templaterouterLister: arguments.ControllerManager.GetInformerProvider(&routerv1alpha1.TemplateRouter{}).(routerv1alpha1informer.TemplateRouterInformer).Lister(),
        templaterouterclient: arguments.Clientset.RouterV1alpha1(),
    }

    // Create a new controller that will call TemplateRouterController.Reconcile on changes to TemplateRouters
    gc := &controller.GenericController{
        Name: "TemplateRouterController",
        Reconcile: bc.Reconcile,
        InformerRegistry: arguments.ControllerManager,
    }
    if err := gc.Watch(&routerv1alpha1.TemplateRouter{}); err != nil {
        return gc, err
    }

    // INSERT ADDITIONAL WATCHES HERE BY CALLING gc.Watch.*() FUNCTIONS
    // NOTE: Informers for Kubernetes resources *MUST* be registered in the pkg/inject package so that they are started.
    return gc, nil
}
