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
	v1alpha1 "github.com/kubedb/user-manager/apis/users/v1alpha1"
	scheme "github.com/kubedb/user-manager/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MessagingServicesGetter has a method to return a MessagingServiceInterface.
// A group's client should implement this interface.
type MessagingServicesGetter interface {
	MessagingServices(namespace string) MessagingServiceInterface
}

// MessagingServiceInterface has methods to work with MessagingService resources.
type MessagingServiceInterface interface {
	Create(*v1alpha1.MessagingService) (*v1alpha1.MessagingService, error)
	Update(*v1alpha1.MessagingService) (*v1alpha1.MessagingService, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.MessagingService, error)
	List(opts v1.ListOptions) (*v1alpha1.MessagingServiceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MessagingService, err error)
	MessagingServiceExpansion
}

// messagingServices implements MessagingServiceInterface
type messagingServices struct {
	client rest.Interface
	ns     string
}

// newMessagingServices returns a MessagingServices
func newMessagingServices(c *UsersV1alpha1Client, namespace string) *messagingServices {
	return &messagingServices{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the messagingService, and returns the corresponding messagingService object, and an error if there is any.
func (c *messagingServices) Get(name string, options v1.GetOptions) (result *v1alpha1.MessagingService, err error) {
	result = &v1alpha1.MessagingService{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("messagingservices").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MessagingServices that match those selectors.
func (c *messagingServices) List(opts v1.ListOptions) (result *v1alpha1.MessagingServiceList, err error) {
	result = &v1alpha1.MessagingServiceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("messagingservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested messagingServices.
func (c *messagingServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("messagingservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a messagingService and creates it.  Returns the server's representation of the messagingService, and an error, if there is any.
func (c *messagingServices) Create(messagingService *v1alpha1.MessagingService) (result *v1alpha1.MessagingService, err error) {
	result = &v1alpha1.MessagingService{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("messagingservices").
		Body(messagingService).
		Do().
		Into(result)
	return
}

// Update takes the representation of a messagingService and updates it. Returns the server's representation of the messagingService, and an error, if there is any.
func (c *messagingServices) Update(messagingService *v1alpha1.MessagingService) (result *v1alpha1.MessagingService, err error) {
	result = &v1alpha1.MessagingService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("messagingservices").
		Name(messagingService.Name).
		Body(messagingService).
		Do().
		Into(result)
	return
}

// Delete takes name of the messagingService and deletes it. Returns an error if one occurs.
func (c *messagingServices) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("messagingservices").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *messagingServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("messagingservices").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched messagingService.
func (c *messagingServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MessagingService, err error) {
	result = &v1alpha1.MessagingService{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("messagingservices").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
