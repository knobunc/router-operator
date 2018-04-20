package fake

import (
	v1alpha1 "github.com/knobunc/router-operator/pkg/apis/router/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTemplateRouters implements TemplateRouterInterface
type FakeTemplateRouters struct {
	Fake *FakeRouterV1alpha1
	ns   string
}

var templateroutersResource = schema.GroupVersionResource{Group: "router.operations.openshift.io", Version: "v1alpha1", Resource: "templaterouters"}

var templateroutersKind = schema.GroupVersionKind{Group: "router.operations.openshift.io", Version: "v1alpha1", Kind: "TemplateRouter"}

// Get takes name of the templateRouter, and returns the corresponding templateRouter object, and an error if there is any.
func (c *FakeTemplateRouters) Get(name string, options v1.GetOptions) (result *v1alpha1.TemplateRouter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(templateroutersResource, c.ns, name), &v1alpha1.TemplateRouter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TemplateRouter), err
}

// List takes label and field selectors, and returns the list of TemplateRouters that match those selectors.
func (c *FakeTemplateRouters) List(opts v1.ListOptions) (result *v1alpha1.TemplateRouterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(templateroutersResource, templateroutersKind, c.ns, opts), &v1alpha1.TemplateRouterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TemplateRouterList{}
	for _, item := range obj.(*v1alpha1.TemplateRouterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested templateRouters.
func (c *FakeTemplateRouters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(templateroutersResource, c.ns, opts))

}

// Create takes the representation of a templateRouter and creates it.  Returns the server's representation of the templateRouter, and an error, if there is any.
func (c *FakeTemplateRouters) Create(templateRouter *v1alpha1.TemplateRouter) (result *v1alpha1.TemplateRouter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(templateroutersResource, c.ns, templateRouter), &v1alpha1.TemplateRouter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TemplateRouter), err
}

// Update takes the representation of a templateRouter and updates it. Returns the server's representation of the templateRouter, and an error, if there is any.
func (c *FakeTemplateRouters) Update(templateRouter *v1alpha1.TemplateRouter) (result *v1alpha1.TemplateRouter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(templateroutersResource, c.ns, templateRouter), &v1alpha1.TemplateRouter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TemplateRouter), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTemplateRouters) UpdateStatus(templateRouter *v1alpha1.TemplateRouter) (*v1alpha1.TemplateRouter, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(templateroutersResource, "status", c.ns, templateRouter), &v1alpha1.TemplateRouter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TemplateRouter), err
}

// Delete takes name of the templateRouter and deletes it. Returns an error if one occurs.
func (c *FakeTemplateRouters) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(templateroutersResource, c.ns, name), &v1alpha1.TemplateRouter{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTemplateRouters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(templateroutersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.TemplateRouterList{})
	return err
}

// Patch applies the patch and returns the patched templateRouter.
func (c *FakeTemplateRouters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.TemplateRouter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(templateroutersResource, c.ns, name, data, subresources...), &v1alpha1.TemplateRouter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TemplateRouter), err
}
