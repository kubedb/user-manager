/*
Copyright 2018 The Attic Authors.

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
	v1alpha1 "github.com/kubedb/user-manager/apis/authorization/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MongoDBRoleBindingLister helps list MongoDBRoleBindings.
type MongoDBRoleBindingLister interface {
	// List lists all MongoDBRoleBindings in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.MongoDBRoleBinding, err error)
	// MongoDBRoleBindings returns an object that can list and get MongoDBRoleBindings.
	MongoDBRoleBindings(namespace string) MongoDBRoleBindingNamespaceLister
	MongoDBRoleBindingListerExpansion
}

// mongoDBRoleBindingLister implements the MongoDBRoleBindingLister interface.
type mongoDBRoleBindingLister struct {
	indexer cache.Indexer
}

// NewMongoDBRoleBindingLister returns a new MongoDBRoleBindingLister.
func NewMongoDBRoleBindingLister(indexer cache.Indexer) MongoDBRoleBindingLister {
	return &mongoDBRoleBindingLister{indexer: indexer}
}

// List lists all MongoDBRoleBindings in the indexer.
func (s *mongoDBRoleBindingLister) List(selector labels.Selector) (ret []*v1alpha1.MongoDBRoleBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MongoDBRoleBinding))
	})
	return ret, err
}

// MongoDBRoleBindings returns an object that can list and get MongoDBRoleBindings.
func (s *mongoDBRoleBindingLister) MongoDBRoleBindings(namespace string) MongoDBRoleBindingNamespaceLister {
	return mongoDBRoleBindingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MongoDBRoleBindingNamespaceLister helps list and get MongoDBRoleBindings.
type MongoDBRoleBindingNamespaceLister interface {
	// List lists all MongoDBRoleBindings in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.MongoDBRoleBinding, err error)
	// Get retrieves the MongoDBRoleBinding from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.MongoDBRoleBinding, error)
	MongoDBRoleBindingNamespaceListerExpansion
}

// mongoDBRoleBindingNamespaceLister implements the MongoDBRoleBindingNamespaceLister
// interface.
type mongoDBRoleBindingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MongoDBRoleBindings in the indexer for a given namespace.
func (s mongoDBRoleBindingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MongoDBRoleBinding, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MongoDBRoleBinding))
	})
	return ret, err
}

// Get retrieves the MongoDBRoleBinding from the indexer for a given namespace and name.
func (s mongoDBRoleBindingNamespaceLister) Get(name string) (*v1alpha1.MongoDBRoleBinding, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("mongodbrolebinding"), name)
	}
	return obj.(*v1alpha1.MongoDBRoleBinding), nil
}
