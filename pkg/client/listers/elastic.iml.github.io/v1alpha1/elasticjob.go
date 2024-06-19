/*
Copyright 2022.

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

package v1alpha1

import (
	v1alpha1 "github.com/kidlj/dlrover-operator/api/elastic.iml.github.io/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ElasticJobLister helps list ElasticJobs.
// All objects returned here must be treated as read-only.
type ElasticJobLister interface {
	// List lists all ElasticJobs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ElasticJob, err error)
	// ElasticJobs returns an object that can list and get ElasticJobs.
	ElasticJobs(namespace string) ElasticJobNamespaceLister
	ElasticJobListerExpansion
}

// elasticJobLister implements the ElasticJobLister interface.
type elasticJobLister struct {
	indexer cache.Indexer
}

// NewElasticJobLister returns a new ElasticJobLister.
func NewElasticJobLister(indexer cache.Indexer) ElasticJobLister {
	return &elasticJobLister{indexer: indexer}
}

// List lists all ElasticJobs in the indexer.
func (s *elasticJobLister) List(selector labels.Selector) (ret []*v1alpha1.ElasticJob, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ElasticJob))
	})
	return ret, err
}

// ElasticJobs returns an object that can list and get ElasticJobs.
func (s *elasticJobLister) ElasticJobs(namespace string) ElasticJobNamespaceLister {
	return elasticJobNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ElasticJobNamespaceLister helps list and get ElasticJobs.
// All objects returned here must be treated as read-only.
type ElasticJobNamespaceLister interface {
	// List lists all ElasticJobs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ElasticJob, err error)
	// Get retrieves the ElasticJob from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ElasticJob, error)
	ElasticJobNamespaceListerExpansion
}

// elasticJobNamespaceLister implements the ElasticJobNamespaceLister
// interface.
type elasticJobNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ElasticJobs in the indexer for a given namespace.
func (s elasticJobNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ElasticJob, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ElasticJob))
	})
	return ret, err
}

// Get retrieves the ElasticJob from the indexer for a given namespace and name.
func (s elasticJobNamespaceLister) Get(name string) (*v1alpha1.ElasticJob, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("elasticjob"), name)
	}
	return obj.(*v1alpha1.ElasticJob), nil
}
