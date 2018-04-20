package v1alpha1

import (
	v1alpha1 "github.com/knobunc/router-operator/pkg/apis/router/v1alpha1"
	scheme "github.com/knobunc/router-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TemplateRoutersGetter has a method to return a TemplateRouterInterface.
// A group's client should implement this interface.
type TemplateRoutersGetter interface {
	TemplateRouters(namespace string) TemplateRouterInterface
}

// TemplateRouterInterface has methods to work with TemplateRouter resources.
type TemplateRouterInterface interface {
	Create(*v1alpha1.TemplateRouter) (*v1alpha1.TemplateRouter, error)
	Update(*v1alpha1.TemplateRouter) (*v1alpha1.TemplateRouter, error)
	UpdateStatus(*v1alpha1.TemplateRouter) (*v1alpha1.TemplateRouter, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.TemplateRouter, error)
	List(opts v1.ListOptions) (*v1alpha1.TemplateRouterList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.TemplateRouter, err error)
	TemplateRouterExpansion
}

// templateRouters implements TemplateRouterInterface
type templateRouters struct {
	client rest.Interface
	ns     string
}

// newTemplateRouters returns a TemplateRouters
func newTemplateRouters(c *RouterV1alpha1Client, namespace string) *templateRouters {
	return &templateRouters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the templateRouter, and returns the corresponding templateRouter object, and an error if there is any.
func (c *templateRouters) Get(name string, options v1.GetOptions) (result *v1alpha1.TemplateRouter, err error) {
	result = &v1alpha1.TemplateRouter{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("templaterouters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TemplateRouters that match those selectors.
func (c *templateRouters) List(opts v1.ListOptions) (result *v1alpha1.TemplateRouterList, err error) {
	result = &v1alpha1.TemplateRouterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("templaterouters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested templateRouters.
func (c *templateRouters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("templaterouters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a templateRouter and creates it.  Returns the server's representation of the templateRouter, and an error, if there is any.
func (c *templateRouters) Create(templateRouter *v1alpha1.TemplateRouter) (result *v1alpha1.TemplateRouter, err error) {
	result = &v1alpha1.TemplateRouter{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("templaterouters").
		Body(templateRouter).
		Do().
		Into(result)
	return
}

// Update takes the representation of a templateRouter and updates it. Returns the server's representation of the templateRouter, and an error, if there is any.
func (c *templateRouters) Update(templateRouter *v1alpha1.TemplateRouter) (result *v1alpha1.TemplateRouter, err error) {
	result = &v1alpha1.TemplateRouter{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("templaterouters").
		Name(templateRouter.Name).
		Body(templateRouter).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *templateRouters) UpdateStatus(templateRouter *v1alpha1.TemplateRouter) (result *v1alpha1.TemplateRouter, err error) {
	result = &v1alpha1.TemplateRouter{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("templaterouters").
		Name(templateRouter.Name).
		SubResource("status").
		Body(templateRouter).
		Do().
		Into(result)
	return
}

// Delete takes name of the templateRouter and deletes it. Returns an error if one occurs.
func (c *templateRouters) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("templaterouters").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *templateRouters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("templaterouters").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched templateRouter.
func (c *templateRouters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.TemplateRouter, err error) {
	result = &v1alpha1.TemplateRouter{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("templaterouters").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
