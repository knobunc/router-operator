


package v1alpha1

import (
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!
// Created by "kubebuilder create resource" for you to implement the TemplateRouter resource schema definition
// as a go struct.
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TemplateRouterSpec defines the desired state of TemplateRouter
type TemplateRouterSpec struct {
    // INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
}

// TemplateRouterStatus defines the observed state of TemplateRouter
type TemplateRouterStatus struct {
    // INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
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
