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

// ScalePlanLister helps list ScalePlans.
// All objects returned here must be treated as read-only.
type ScalePlanLister interface {
	// List lists all ScalePlans in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ScalePlan, err error)
	// ScalePlans returns an object that can list and get ScalePlans.
	ScalePlans(namespace string) ScalePlanNamespaceLister
	ScalePlanListerExpansion
}

// scalePlanLister implements the ScalePlanLister interface.
type scalePlanLister struct {
	indexer cache.Indexer
}

// NewScalePlanLister returns a new ScalePlanLister.
func NewScalePlanLister(indexer cache.Indexer) ScalePlanLister {
	return &scalePlanLister{indexer: indexer}
}

// List lists all ScalePlans in the indexer.
func (s *scalePlanLister) List(selector labels.Selector) (ret []*v1alpha1.ScalePlan, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ScalePlan))
	})
	return ret, err
}

// ScalePlans returns an object that can list and get ScalePlans.
func (s *scalePlanLister) ScalePlans(namespace string) ScalePlanNamespaceLister {
	return scalePlanNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ScalePlanNamespaceLister helps list and get ScalePlans.
// All objects returned here must be treated as read-only.
type ScalePlanNamespaceLister interface {
	// List lists all ScalePlans in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ScalePlan, err error)
	// Get retrieves the ScalePlan from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ScalePlan, error)
	ScalePlanNamespaceListerExpansion
}

// scalePlanNamespaceLister implements the ScalePlanNamespaceLister
// interface.
type scalePlanNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ScalePlans in the indexer for a given namespace.
func (s scalePlanNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ScalePlan, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ScalePlan))
	})
	return ret, err
}

// Get retrieves the ScalePlan from the indexer for a given namespace and name.
func (s scalePlanNamespaceLister) Get(name string) (*v1alpha1.ScalePlan, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("scaleplan"), name)
	}
	return obj.(*v1alpha1.ScalePlan), nil
}
