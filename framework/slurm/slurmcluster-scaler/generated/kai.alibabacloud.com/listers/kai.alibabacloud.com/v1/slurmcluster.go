/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/AliyunContainerService/ai-models-on-ack/framework/slurm/slurmcluster-scaler/apis/kai.alibabacloud.com/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SlurmClusterLister helps list SlurmClusters.
// All objects returned here must be treated as read-only.
type SlurmClusterLister interface {
	// List lists all SlurmClusters in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.SlurmCluster, err error)
	// SlurmClusters returns an object that can list and get SlurmClusters.
	SlurmClusters(namespace string) SlurmClusterNamespaceLister
	SlurmClusterListerExpansion
}

// slurmClusterLister implements the SlurmClusterLister interface.
type slurmClusterLister struct {
	indexer cache.Indexer
}

// NewSlurmClusterLister returns a new SlurmClusterLister.
func NewSlurmClusterLister(indexer cache.Indexer) SlurmClusterLister {
	return &slurmClusterLister{indexer: indexer}
}

// List lists all SlurmClusters in the indexer.
func (s *slurmClusterLister) List(selector labels.Selector) (ret []*v1.SlurmCluster, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.SlurmCluster))
	})
	return ret, err
}

// SlurmClusters returns an object that can list and get SlurmClusters.
func (s *slurmClusterLister) SlurmClusters(namespace string) SlurmClusterNamespaceLister {
	return slurmClusterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SlurmClusterNamespaceLister helps list and get SlurmClusters.
// All objects returned here must be treated as read-only.
type SlurmClusterNamespaceLister interface {
	// List lists all SlurmClusters in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.SlurmCluster, err error)
	// Get retrieves the SlurmCluster from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.SlurmCluster, error)
	SlurmClusterNamespaceListerExpansion
}

// slurmClusterNamespaceLister implements the SlurmClusterNamespaceLister
// interface.
type slurmClusterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SlurmClusters in the indexer for a given namespace.
func (s slurmClusterNamespaceLister) List(selector labels.Selector) (ret []*v1.SlurmCluster, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.SlurmCluster))
	})
	return ret, err
}

// Get retrieves the SlurmCluster from the indexer for a given namespace and name.
func (s slurmClusterNamespaceLister) Get(name string) (*v1.SlurmCluster, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("slurmcluster"), name)
	}
	return obj.(*v1.SlurmCluster), nil
}
