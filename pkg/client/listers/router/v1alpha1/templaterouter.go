// This file was automatically generated by lister-gen

package v1alpha1

import (
	v1alpha1 "github.com/knobunc/router-operator/pkg/apis/router/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TemplateRouterLister helps list TemplateRouters.
type TemplateRouterLister interface {
	// List lists all TemplateRouters in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.TemplateRouter, err error)
	// TemplateRouters returns an object that can list and get TemplateRouters.
	TemplateRouters(namespace string) TemplateRouterNamespaceLister
	TemplateRouterListerExpansion
}

// templateRouterLister implements the TemplateRouterLister interface.
type templateRouterLister struct {
	indexer cache.Indexer
}

// NewTemplateRouterLister returns a new TemplateRouterLister.
func NewTemplateRouterLister(indexer cache.Indexer) TemplateRouterLister {
	return &templateRouterLister{indexer: indexer}
}

// List lists all TemplateRouters in the indexer.
func (s *templateRouterLister) List(selector labels.Selector) (ret []*v1alpha1.TemplateRouter, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TemplateRouter))
	})
	return ret, err
}

// TemplateRouters returns an object that can list and get TemplateRouters.
func (s *templateRouterLister) TemplateRouters(namespace string) TemplateRouterNamespaceLister {
	return templateRouterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TemplateRouterNamespaceLister helps list and get TemplateRouters.
type TemplateRouterNamespaceLister interface {
	// List lists all TemplateRouters in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.TemplateRouter, err error)
	// Get retrieves the TemplateRouter from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.TemplateRouter, error)
	TemplateRouterNamespaceListerExpansion
}

// templateRouterNamespaceLister implements the TemplateRouterNamespaceLister
// interface.
type templateRouterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TemplateRouters in the indexer for a given namespace.
func (s templateRouterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.TemplateRouter, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TemplateRouter))
	})
	return ret, err
}

// Get retrieves the TemplateRouter from the indexer for a given namespace and name.
func (s templateRouterNamespaceLister) Get(name string) (*v1alpha1.TemplateRouter, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("templaterouter"), name)
	}
	return obj.(*v1alpha1.TemplateRouter), nil
}
