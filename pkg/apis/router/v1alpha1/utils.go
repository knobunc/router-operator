package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	defaultBaseImage = "openshift/origin-haproxy-router"
	// version format is "<upstream-version>-<our-version>"
	defaultVersion = "v3.9.0-alpha.4-0"
)

// SetDefaults sets the default vaules for the TemplateRouter spec and returns true if the spec was changed
func (t *TemplateRouterSpec) SetDefaults() bool {
	changed := false
	ts := &t.Spec
	if ts.Replicas == 0 {
		ts.Replicas = 1
		changed = true
	}
	if len(vs.BaseImage) == 0 {
		vs.BaseImage = defaultBaseImage
		changed = true
	}
	if len(vs.Version) == 0 {
		vs.Version = defaultVersion
		changed = true
	}
	return changed
}
