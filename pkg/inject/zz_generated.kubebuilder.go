package inject

import (
	routerv1alpha1 "github.com/knobunc/router-operator/pkg/apis/router/v1alpha1"
	rscheme "github.com/knobunc/router-operator/pkg/client/clientset/versioned/scheme"
	"github.com/knobunc/router-operator/pkg/controller/templaterouter"
	"github.com/knobunc/router-operator/pkg/inject/args"
	"github.com/kubernetes-sigs/kubebuilder/pkg/inject/run"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/scheme"
)

func init() {
	rscheme.AddToScheme(scheme.Scheme)

	// Inject Informers
	Inject = append(Inject, func(arguments args.InjectArgs) error {
		Injector.ControllerManager = arguments.ControllerManager

		if err := arguments.ControllerManager.AddInformerProvider(&routerv1alpha1.TemplateRouter{}, arguments.Informers.Router().V1alpha1().TemplateRouters()); err != nil {
			return err
		}

		// Add Kubernetes informers

		if c, err := templaterouter.ProvideController(arguments); err != nil {
			return err
		} else {
			arguments.ControllerManager.AddController(c)
		}
		return nil
	})

	// Inject CRDs
	Injector.CRDs = append(Injector.CRDs, &routerv1alpha1.TemplateRouterCRD)
	// Inject PolicyRules
	Injector.PolicyRules = append(Injector.PolicyRules, rbacv1.PolicyRule{
		APIGroups: []string{"router.operations.openshift.io"},
		Resources: []string{"*"},
		Verbs:     []string{"*"},
	})
	// Inject GroupVersions
	Injector.GroupVersions = append(Injector.GroupVersions, schema.GroupVersion{
		Group:   "router.operations.openshift.io",
		Version: "v1alpha1",
	})
	Injector.RunFns = append(Injector.RunFns, func(arguments run.RunArguments) error {
		Injector.ControllerManager.RunInformersAndControllers(arguments)
		return nil
	})
}
