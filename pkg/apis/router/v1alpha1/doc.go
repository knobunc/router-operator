


// Api versions allow the api contract for a resource to be changed while keeping
// backward compatibility by support multiple concurrent versions
// of the same resource

// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=package,register
// +k8s:conversion-gen=github.com/knobunc/router-operator/pkg/apis/router
// +k8s:defaulter-gen=TypeMeta
// +groupName=router.operations.openshift.io
package v1alpha1 // import "github.com/knobunc/router-operator/pkg/apis/router/v1alpha1"
