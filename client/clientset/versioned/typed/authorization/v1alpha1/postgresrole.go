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

// PostgresRolesGetter has a method to return a PostgresRoleInterface.
// A group's client should implement this interface.
type PostgresRolesGetter interface {
	PostgresRoles(namespace string) PostgresRoleInterface
}

// PostgresRoleInterface has methods to work with PostgresRole resources.
type PostgresRoleInterface interface {
	Create(*v1alpha1.PostgresRole) (*v1alpha1.PostgresRole, error)
	Update(*v1alpha1.PostgresRole) (*v1alpha1.PostgresRole, error)
	UpdateStatus(*v1alpha1.PostgresRole) (*v1alpha1.PostgresRole, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.PostgresRole, error)
	List(opts v1.ListOptions) (*v1alpha1.PostgresRoleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PostgresRole, err error)
	PostgresRoleExpansion
}

// postgresRoles implements PostgresRoleInterface
type postgresRoles struct {
	client rest.Interface
	ns     string
}

// newPostgresRoles returns a PostgresRoles
func newPostgresRoles(c *AuthorizationV1alpha1Client, namespace string) *postgresRoles {
	return &postgresRoles{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the postgresRole, and returns the corresponding postgresRole object, and an error if there is any.
func (c *postgresRoles) Get(name string, options v1.GetOptions) (result *v1alpha1.PostgresRole, err error) {
	result = &v1alpha1.PostgresRole{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("postgresroles").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PostgresRoles that match those selectors.
func (c *postgresRoles) List(opts v1.ListOptions) (result *v1alpha1.PostgresRoleList, err error) {
	result = &v1alpha1.PostgresRoleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("postgresroles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested postgresRoles.
func (c *postgresRoles) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("postgresroles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a postgresRole and creates it.  Returns the server's representation of the postgresRole, and an error, if there is any.
func (c *postgresRoles) Create(postgresRole *v1alpha1.PostgresRole) (result *v1alpha1.PostgresRole, err error) {
	result = &v1alpha1.PostgresRole{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("postgresroles").
		Body(postgresRole).
		Do().
		Into(result)
	return
}

// Update takes the representation of a postgresRole and updates it. Returns the server's representation of the postgresRole, and an error, if there is any.
func (c *postgresRoles) Update(postgresRole *v1alpha1.PostgresRole) (result *v1alpha1.PostgresRole, err error) {
	result = &v1alpha1.PostgresRole{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("postgresroles").
		Name(postgresRole.Name).
		Body(postgresRole).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *postgresRoles) UpdateStatus(postgresRole *v1alpha1.PostgresRole) (result *v1alpha1.PostgresRole, err error) {
	result = &v1alpha1.PostgresRole{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("postgresroles").
		Name(postgresRole.Name).
		SubResource("status").
		Body(postgresRole).
		Do().
		Into(result)
	return
}

// Delete takes name of the postgresRole and deletes it. Returns an error if one occurs.
func (c *postgresRoles) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("postgresroles").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *postgresRoles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("postgresroles").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched postgresRole.
func (c *postgresRoles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PostgresRole, err error) {
	result = &v1alpha1.PostgresRole{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("postgresroles").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
