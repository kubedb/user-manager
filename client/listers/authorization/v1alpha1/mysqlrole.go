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

// MysqlRoleLister helps list MysqlRoles.
type MysqlRoleLister interface {
	// List lists all MysqlRoles in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.MysqlRole, err error)
	// MysqlRoles returns an object that can list and get MysqlRoles.
	MysqlRoles(namespace string) MysqlRoleNamespaceLister
	MysqlRoleListerExpansion
}

// mysqlRoleLister implements the MysqlRoleLister interface.
type mysqlRoleLister struct {
	indexer cache.Indexer
}

// NewMysqlRoleLister returns a new MysqlRoleLister.
func NewMysqlRoleLister(indexer cache.Indexer) MysqlRoleLister {
	return &mysqlRoleLister{indexer: indexer}
}

// List lists all MysqlRoles in the indexer.
func (s *mysqlRoleLister) List(selector labels.Selector) (ret []*v1alpha1.MysqlRole, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MysqlRole))
	})
	return ret, err
}

// MysqlRoles returns an object that can list and get MysqlRoles.
func (s *mysqlRoleLister) MysqlRoles(namespace string) MysqlRoleNamespaceLister {
	return mysqlRoleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MysqlRoleNamespaceLister helps list and get MysqlRoles.
type MysqlRoleNamespaceLister interface {
	// List lists all MysqlRoles in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.MysqlRole, err error)
	// Get retrieves the MysqlRole from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.MysqlRole, error)
	MysqlRoleNamespaceListerExpansion
}

// mysqlRoleNamespaceLister implements the MysqlRoleNamespaceLister
// interface.
type mysqlRoleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MysqlRoles in the indexer for a given namespace.
func (s mysqlRoleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MysqlRole, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MysqlRole))
	})
	return ret, err
}

// Get retrieves the MysqlRole from the indexer for a given namespace and name.
func (s mysqlRoleNamespaceLister) Get(name string) (*v1alpha1.MysqlRole, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("mysqlrole"), name)
	}
	return obj.(*v1alpha1.MysqlRole), nil
}
