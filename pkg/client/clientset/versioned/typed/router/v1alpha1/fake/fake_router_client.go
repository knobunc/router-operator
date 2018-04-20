package fake

import (
	v1alpha1 "github.com/knobunc/router-operator/pkg/client/clientset/versioned/typed/router/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeRouterV1alpha1 struct {
	*testing.Fake
}

func (c *FakeRouterV1alpha1) TemplateRouters(namespace string) v1alpha1.TemplateRouterInterface {
	return &FakeTemplateRouters{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeRouterV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
