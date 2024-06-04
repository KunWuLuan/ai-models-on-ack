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
	v1 "github.com/AliyunContainerService/ai-models-on-ack/framework/slurm/slurmcluster-scaler/apis/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SlurmJobLister helps list SlurmJobs.
// All objects returned here must be treated as read-only.
type SlurmJobLister interface {
	// List lists all SlurmJobs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.SlurmJob, err error)
	// SlurmJobs returns an object that can list and get SlurmJobs.
	SlurmJobs(namespace string) SlurmJobNamespaceLister
	SlurmJobListerExpansion
}

// slurmJobLister implements the SlurmJobLister interface.
type slurmJobLister struct {
	indexer cache.Indexer
}

// NewSlurmJobLister returns a new SlurmJobLister.
func NewSlurmJobLister(indexer cache.Indexer) SlurmJobLister {
	return &slurmJobLister{indexer: indexer}
}

// List lists all SlurmJobs in the indexer.
func (s *slurmJobLister) List(selector labels.Selector) (ret []*v1.SlurmJob, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.SlurmJob))
	})
	return ret, err
}

// SlurmJobs returns an object that can list and get SlurmJobs.
func (s *slurmJobLister) SlurmJobs(namespace string) SlurmJobNamespaceLister {
	return slurmJobNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SlurmJobNamespaceLister helps list and get SlurmJobs.
// All objects returned here must be treated as read-only.
type SlurmJobNamespaceLister interface {
	// List lists all SlurmJobs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.SlurmJob, err error)
	// Get retrieves the SlurmJob from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.SlurmJob, error)
	SlurmJobNamespaceListerExpansion
}

// slurmJobNamespaceLister implements the SlurmJobNamespaceLister
// interface.
type slurmJobNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SlurmJobs in the indexer for a given namespace.
func (s slurmJobNamespaceLister) List(selector labels.Selector) (ret []*v1.SlurmJob, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.SlurmJob))
	})
	return ret, err
}

// Get retrieves the SlurmJob from the indexer for a given namespace and name.
func (s slurmJobNamespaceLister) Get(name string) (*v1.SlurmJob, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("slurmjob"), name)
	}
	return obj.(*v1.SlurmJob), nil
}
