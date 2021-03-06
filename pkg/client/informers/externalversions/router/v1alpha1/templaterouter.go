// This file was automatically generated by informer-gen

package v1alpha1

import (
	router_v1alpha1 "github.com/knobunc/router-operator/pkg/apis/router/v1alpha1"
	versioned "github.com/knobunc/router-operator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/knobunc/router-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/knobunc/router-operator/pkg/client/listers/router/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// TemplateRouterInformer provides access to a shared informer and lister for
// TemplateRouters.
type TemplateRouterInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.TemplateRouterLister
}

type templateRouterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewTemplateRouterInformer constructs a new informer for TemplateRouter type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTemplateRouterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTemplateRouterInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredTemplateRouterInformer constructs a new informer for TemplateRouter type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTemplateRouterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RouterV1alpha1().TemplateRouters(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RouterV1alpha1().TemplateRouters(namespace).Watch(options)
			},
		},
		&router_v1alpha1.TemplateRouter{},
		resyncPeriod,
		indexers,
	)
}

func (f *templateRouterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTemplateRouterInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *templateRouterInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&router_v1alpha1.TemplateRouter{}, f.defaultInformer)
}

func (f *templateRouterInformer) Lister() v1alpha1.TemplateRouterLister {
	return v1alpha1.NewTemplateRouterLister(f.Informer().GetIndexer())
}
