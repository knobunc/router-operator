package v1alpha1

const (
	defaultBaseImage = "openshift/origin-haproxy-router"
	// TODO: version format is "<upstream-version>-<our-version>"
	defaultVersion = "v3.9.0-alpha.4"
)

// SetDefaults sets the default vaules for the TemplateRouter spec and returns true if the spec was changed
func (t *TemplateRouter) SetDefaults() bool {
	changed := false
	ts := &t.Spec
	if len(ts.BaseImage) == 0 {
		ts.BaseImage = defaultBaseImage
		changed = true
	}
	if len(ts.Version) == 0 {
		ts.Version = defaultVersion
		changed = true
	}
	return changed
}
