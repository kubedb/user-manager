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

// PostgresRoleBindingsGetter has a method to return a PostgresRoleBindingInterface.
// A group's client should implement this interface.
type PostgresRoleBindingsGetter interface {
	PostgresRoleBindings(namespace string) PostgresRoleBindingInterface
}

// PostgresRoleBindingInterface has methods to work with PostgresRoleBinding resources.
type PostgresRoleBindingInterface interface {
	Create(*v1alpha1.PostgresRoleBinding) (*v1alpha1.PostgresRoleBinding, error)
	Update(*v1alpha1.PostgresRoleBinding) (*v1alpha1.PostgresRoleBinding, error)
	UpdateStatus(*v1alpha1.PostgresRoleBinding) (*v1alpha1.PostgresRoleBinding, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.PostgresRoleBinding, error)
	List(opts v1.ListOptions) (*v1alpha1.PostgresRoleBindingList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PostgresRoleBinding, err error)
	PostgresRoleBindingExpansion
}

// postgresRoleBindings implements PostgresRoleBindingInterface
type postgresRoleBindings struct {
	client rest.Interface
	ns     string
}

// newPostgresRoleBindings returns a PostgresRoleBindings
func newPostgresRoleBindings(c *AuthorizationV1alpha1Client, namespace string) *postgresRoleBindings {
	return &postgresRoleBindings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the postgresRoleBinding, and returns the corresponding postgresRoleBinding object, and an error if there is any.
func (c *postgresRoleBindings) Get(name string, options v1.GetOptions) (result *v1alpha1.PostgresRoleBinding, err error) {
	result = &v1alpha1.PostgresRoleBinding{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("postgresrolebindings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PostgresRoleBindings that match those selectors.
func (c *postgresRoleBindings) List(opts v1.ListOptions) (result *v1alpha1.PostgresRoleBindingList, err error) {
	result = &v1alpha1.PostgresRoleBindingList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("postgresrolebindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested postgresRoleBindings.
func (c *postgresRoleBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("postgresrolebindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a postgresRoleBinding and creates it.  Returns the server's representation of the postgresRoleBinding, and an error, if there is any.
func (c *postgresRoleBindings) Create(postgresRoleBinding *v1alpha1.PostgresRoleBinding) (result *v1alpha1.PostgresRoleBinding, err error) {
	result = &v1alpha1.PostgresRoleBinding{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("postgresrolebindings").
		Body(postgresRoleBinding).
		Do().
		Into(result)
	return
}

// Update takes the representation of a postgresRoleBinding and updates it. Returns the server's representation of the postgresRoleBinding, and an error, if there is any.
func (c *postgresRoleBindings) Update(postgresRoleBinding *v1alpha1.PostgresRoleBinding) (result *v1alpha1.PostgresRoleBinding, err error) {
	result = &v1alpha1.PostgresRoleBinding{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("postgresrolebindings").
		Name(postgresRoleBinding.Name).
		Body(postgresRoleBinding).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *postgresRoleBindings) UpdateStatus(postgresRoleBinding *v1alpha1.PostgresRoleBinding) (result *v1alpha1.PostgresRoleBinding, err error) {
	result = &v1alpha1.PostgresRoleBinding{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("postgresrolebindings").
		Name(postgresRoleBinding.Name).
		SubResource("status").
		Body(postgresRoleBinding).
		Do().
		Into(result)
	return
}

// Delete takes name of the postgresRoleBinding and deletes it. Returns an error if one occurs.
func (c *postgresRoleBindings) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("postgresrolebindings").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *postgresRoleBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("postgresrolebindings").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched postgresRoleBinding.
func (c *postgresRoleBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PostgresRoleBinding, err error) {
	result = &v1alpha1.PostgresRoleBinding{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("postgresrolebindings").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
