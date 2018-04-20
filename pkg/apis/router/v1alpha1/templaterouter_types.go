package v1alpha1

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!
// Created by "kubebuilder create resource" for you to implement the TemplateRouter resource schema definition
// as a go struct.
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TemplateRouterSpec defines the desired state of TemplateRouter
type TemplateRouterSpec struct {
	// Number of replicas to deploy for a TemplateRouter deployment.
	// Default: 1.
	Replicas int32 `json:"replicas,omitempty"`

	// Base image to use for a Router deployment.
	BaseImage string `json:"baseImage"`

	// Version of Router to be deployed.
	Version string `json:"version"`

	// PodPolicy defines the policy for pods owned by the router operator.
	// This field cannot be updated once the CR is created. (TODO: Why?)
	Pod *PodPolicy `json:"pod,omitempty"`
}

// PodPolicy defines the policy for pods owned by vault operator.
type PodPolicy struct {
	// Resources is the resource requirements for the containers.
	Resources v1.ResourceRequirements `json:"resources,omitempty"`
}

// TemplateRouterStatus defines the observed state of TemplateRouter
type TemplateRouterStatus struct {
	// PodNames of the routers.
	Routers []string `json:"routers"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TemplateRouter
// +k8s:openapi-gen=true
// +resource:path=templaterouters
type TemplateRouter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TemplateRouterSpec   `json:"spec,omitempty"`
	Status TemplateRouterStatus `json:"status,omitempty"`
}
