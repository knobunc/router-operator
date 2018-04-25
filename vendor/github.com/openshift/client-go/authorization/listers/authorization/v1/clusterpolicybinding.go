// This file was automatically generated by lister-gen

package v1

import (
	v1 "github.com/openshift/api/authorization/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterPolicyBindingLister helps list ClusterPolicyBindings.
type ClusterPolicyBindingLister interface {
	// List lists all ClusterPolicyBindings in the indexer.
	List(selector labels.Selector) (ret []*v1.ClusterPolicyBinding, err error)
	// Get retrieves the ClusterPolicyBinding from the index for a given name.
	Get(name string) (*v1.ClusterPolicyBinding, error)
	ClusterPolicyBindingListerExpansion
}

// clusterPolicyBindingLister implements the ClusterPolicyBindingLister interface.
type clusterPolicyBindingLister struct {
	indexer cache.Indexer
}

// NewClusterPolicyBindingLister returns a new ClusterPolicyBindingLister.
func NewClusterPolicyBindingLister(indexer cache.Indexer) ClusterPolicyBindingLister {
	return &clusterPolicyBindingLister{indexer: indexer}
}

// List lists all ClusterPolicyBindings in the indexer.
func (s *clusterPolicyBindingLister) List(selector labels.Selector) (ret []*v1.ClusterPolicyBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ClusterPolicyBinding))
	})
	return ret, err
}

// Get retrieves the ClusterPolicyBinding from the index for a given name.
func (s *clusterPolicyBindingLister) Get(name string) (*v1.ClusterPolicyBinding, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("clusterpolicybinding"), name)
	}
	return obj.(*v1.ClusterPolicyBinding), nil
}
