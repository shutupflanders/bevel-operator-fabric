/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

import (
	v1alpha1 "github.com/kfsoftware/hlf-operator/api/hlf.kungfusoftware.es/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FabricOrderingServiceLister helps list FabricOrderingServices.
// All objects returned here must be treated as read-only.
type FabricOrderingServiceLister interface {
	// List lists all FabricOrderingServices in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FabricOrderingService, err error)
	// FabricOrderingServices returns an object that can list and get FabricOrderingServices.
	FabricOrderingServices(namespace string) FabricOrderingServiceNamespaceLister
	FabricOrderingServiceListerExpansion
}

// fabricOrderingServiceLister implements the FabricOrderingServiceLister interface.
type fabricOrderingServiceLister struct {
	indexer cache.Indexer
}

// NewFabricOrderingServiceLister returns a new FabricOrderingServiceLister.
func NewFabricOrderingServiceLister(indexer cache.Indexer) FabricOrderingServiceLister {
	return &fabricOrderingServiceLister{indexer: indexer}
}

// List lists all FabricOrderingServices in the indexer.
func (s *fabricOrderingServiceLister) List(selector labels.Selector) (ret []*v1alpha1.FabricOrderingService, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FabricOrderingService))
	})
	return ret, err
}

// FabricOrderingServices returns an object that can list and get FabricOrderingServices.
func (s *fabricOrderingServiceLister) FabricOrderingServices(namespace string) FabricOrderingServiceNamespaceLister {
	return fabricOrderingServiceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FabricOrderingServiceNamespaceLister helps list and get FabricOrderingServices.
// All objects returned here must be treated as read-only.
type FabricOrderingServiceNamespaceLister interface {
	// List lists all FabricOrderingServices in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FabricOrderingService, err error)
	// Get retrieves the FabricOrderingService from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.FabricOrderingService, error)
	FabricOrderingServiceNamespaceListerExpansion
}

// fabricOrderingServiceNamespaceLister implements the FabricOrderingServiceNamespaceLister
// interface.
type fabricOrderingServiceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FabricOrderingServices in the indexer for a given namespace.
func (s fabricOrderingServiceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FabricOrderingService, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FabricOrderingService))
	})
	return ret, err
}

// Get retrieves the FabricOrderingService from the indexer for a given namespace and name.
func (s fabricOrderingServiceNamespaceLister) Get(name string) (*v1alpha1.FabricOrderingService, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("fabricorderingservice"), name)
	}
	return obj.(*v1alpha1.FabricOrderingService), nil
}
