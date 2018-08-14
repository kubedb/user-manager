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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/kubedb/user-manager/apis/authorization/v1alpha1"
	scheme "github.com/kubedb/user-manager/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MysqlRoleBindingsGetter has a method to return a MysqlRoleBindingInterface.
// A group's client should implement this interface.
type MysqlRoleBindingsGetter interface {
	MysqlRoleBindings(namespace string) MysqlRoleBindingInterface
}

// MysqlRoleBindingInterface has methods to work with MysqlRoleBinding resources.
type MysqlRoleBindingInterface interface {
	Create(*v1alpha1.MysqlRoleBinding) (*v1alpha1.MysqlRoleBinding, error)
	Update(*v1alpha1.MysqlRoleBinding) (*v1alpha1.MysqlRoleBinding, error)
	UpdateStatus(*v1alpha1.MysqlRoleBinding) (*v1alpha1.MysqlRoleBinding, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.MysqlRoleBinding, error)
	List(opts v1.ListOptions) (*v1alpha1.MysqlRoleBindingList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MysqlRoleBinding, err error)
	MysqlRoleBindingExpansion
}

// mysqlRoleBindings implements MysqlRoleBindingInterface
type mysqlRoleBindings struct {
	client rest.Interface
	ns     string
}

// newMysqlRoleBindings returns a MysqlRoleBindings
func newMysqlRoleBindings(c *AuthorizationV1alpha1Client, namespace string) *mysqlRoleBindings {
	return &mysqlRoleBindings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the mysqlRoleBinding, and returns the corresponding mysqlRoleBinding object, and an error if there is any.
func (c *mysqlRoleBindings) Get(name string, options v1.GetOptions) (result *v1alpha1.MysqlRoleBinding, err error) {
	result = &v1alpha1.MysqlRoleBinding{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mysqlrolebindings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MysqlRoleBindings that match those selectors.
func (c *mysqlRoleBindings) List(opts v1.ListOptions) (result *v1alpha1.MysqlRoleBindingList, err error) {
	result = &v1alpha1.MysqlRoleBindingList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mysqlrolebindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested mysqlRoleBindings.
func (c *mysqlRoleBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("mysqlrolebindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a mysqlRoleBinding and creates it.  Returns the server's representation of the mysqlRoleBinding, and an error, if there is any.
func (c *mysqlRoleBindings) Create(mysqlRoleBinding *v1alpha1.MysqlRoleBinding) (result *v1alpha1.MysqlRoleBinding, err error) {
	result = &v1alpha1.MysqlRoleBinding{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("mysqlrolebindings").
		Body(mysqlRoleBinding).
		Do().
		Into(result)
	return
}

// Update takes the representation of a mysqlRoleBinding and updates it. Returns the server's representation of the mysqlRoleBinding, and an error, if there is any.
func (c *mysqlRoleBindings) Update(mysqlRoleBinding *v1alpha1.MysqlRoleBinding) (result *v1alpha1.MysqlRoleBinding, err error) {
	result = &v1alpha1.MysqlRoleBinding{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mysqlrolebindings").
		Name(mysqlRoleBinding.Name).
		Body(mysqlRoleBinding).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *mysqlRoleBindings) UpdateStatus(mysqlRoleBinding *v1alpha1.MysqlRoleBinding) (result *v1alpha1.MysqlRoleBinding, err error) {
	result = &v1alpha1.MysqlRoleBinding{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mysqlrolebindings").
		Name(mysqlRoleBinding.Name).
		SubResource("status").
		Body(mysqlRoleBinding).
		Do().
		Into(result)
	return
}

// Delete takes name of the mysqlRoleBinding and deletes it. Returns an error if one occurs.
func (c *mysqlRoleBindings) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mysqlrolebindings").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *mysqlRoleBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mysqlrolebindings").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched mysqlRoleBinding.
func (c *mysqlRoleBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MysqlRoleBinding, err error) {
	result = &v1alpha1.MysqlRoleBinding{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("mysqlrolebindings").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
