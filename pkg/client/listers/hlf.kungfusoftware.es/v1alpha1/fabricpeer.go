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

// FabricPeerLister helps list FabricPeers.
// All objects returned here must be treated as read-only.
type FabricPeerLister interface {
	// List lists all FabricPeers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FabricPeer, err error)
	// FabricPeers returns an object that can list and get FabricPeers.
	FabricPeers(namespace string) FabricPeerNamespaceLister
	FabricPeerListerExpansion
}

// fabricPeerLister implements the FabricPeerLister interface.
type fabricPeerLister struct {
	indexer cache.Indexer
}

// NewFabricPeerLister returns a new FabricPeerLister.
func NewFabricPeerLister(indexer cache.Indexer) FabricPeerLister {
	return &fabricPeerLister{indexer: indexer}
}

// List lists all FabricPeers in the indexer.
func (s *fabricPeerLister) List(selector labels.Selector) (ret []*v1alpha1.FabricPeer, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FabricPeer))
	})
	return ret, err
}

// FabricPeers returns an object that can list and get FabricPeers.
func (s *fabricPeerLister) FabricPeers(namespace string) FabricPeerNamespaceLister {
	return fabricPeerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FabricPeerNamespaceLister helps list and get FabricPeers.
// All objects returned here must be treated as read-only.
type FabricPeerNamespaceLister interface {
	// List lists all FabricPeers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FabricPeer, err error)
	// Get retrieves the FabricPeer from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.FabricPeer, error)
	FabricPeerNamespaceListerExpansion
}

// fabricPeerNamespaceLister implements the FabricPeerNamespaceLister
// interface.
type fabricPeerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FabricPeers in the indexer for a given namespace.
func (s fabricPeerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FabricPeer, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FabricPeer))
	})
	return ret, err
}

// Get retrieves the FabricPeer from the indexer for a given namespace and name.
func (s fabricPeerNamespaceLister) Get(name string) (*v1alpha1.FabricPeer, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("fabricpeer"), name)
	}
	return obj.(*v1alpha1.FabricPeer), nil
}
