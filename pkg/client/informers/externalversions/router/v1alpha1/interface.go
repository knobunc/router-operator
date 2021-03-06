// This file was automatically generated by informer-gen

package v1alpha1

import (
	internalinterfaces "github.com/knobunc/router-operator/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// TemplateRouters returns a TemplateRouterInformer.
	TemplateRouters() TemplateRouterInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// TemplateRouters returns a TemplateRouterInformer.
func (v *version) TemplateRouters() TemplateRouterInformer {
	return &templateRouterInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
