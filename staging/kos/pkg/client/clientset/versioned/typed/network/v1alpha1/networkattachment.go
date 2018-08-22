/*
Copyright 2019 The Kubernetes Authors.

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
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "k8s.io/examples/staging/kos/pkg/apis/network/v1alpha1"
	scheme "k8s.io/examples/staging/kos/pkg/client/clientset/versioned/scheme"
)

// NetworkAttachmentsGetter has a method to return a NetworkAttachmentInterface.
// A group's client should implement this interface.
type NetworkAttachmentsGetter interface {
	NetworkAttachments(namespace string) NetworkAttachmentInterface
}

// NetworkAttachmentInterface has methods to work with NetworkAttachment resources.
type NetworkAttachmentInterface interface {
	Create(*v1alpha1.NetworkAttachment) (*v1alpha1.NetworkAttachment, error)
	Update(*v1alpha1.NetworkAttachment) (*v1alpha1.NetworkAttachment, error)
	UpdateStatus(*v1alpha1.NetworkAttachment) (*v1alpha1.NetworkAttachment, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.NetworkAttachment, error)
	List(opts v1.ListOptions) (*v1alpha1.NetworkAttachmentList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NetworkAttachment, err error)
	NetworkAttachmentExpansion
}

// networkAttachments implements NetworkAttachmentInterface
type networkAttachments struct {
	client rest.Interface
	ns     string
}

// newNetworkAttachments returns a NetworkAttachments
func newNetworkAttachments(c *NetworkV1alpha1Client, namespace string) *networkAttachments {
	return &networkAttachments{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the networkAttachment, and returns the corresponding networkAttachment object, and an error if there is any.
func (c *networkAttachments) Get(name string, options v1.GetOptions) (result *v1alpha1.NetworkAttachment, err error) {
	result = &v1alpha1.NetworkAttachment{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("networkattachments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NetworkAttachments that match those selectors.
func (c *networkAttachments) List(opts v1.ListOptions) (result *v1alpha1.NetworkAttachmentList, err error) {
	result = &v1alpha1.NetworkAttachmentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("networkattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested networkAttachments.
func (c *networkAttachments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("networkattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a networkAttachment and creates it.  Returns the server's representation of the networkAttachment, and an error, if there is any.
func (c *networkAttachments) Create(networkAttachment *v1alpha1.NetworkAttachment) (result *v1alpha1.NetworkAttachment, err error) {
	result = &v1alpha1.NetworkAttachment{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("networkattachments").
		Body(networkAttachment).
		Do().
		Into(result)
	return
}

// Update takes the representation of a networkAttachment and updates it. Returns the server's representation of the networkAttachment, and an error, if there is any.
func (c *networkAttachments) Update(networkAttachment *v1alpha1.NetworkAttachment) (result *v1alpha1.NetworkAttachment, err error) {
	result = &v1alpha1.NetworkAttachment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("networkattachments").
		Name(networkAttachment.Name).
		Body(networkAttachment).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *networkAttachments) UpdateStatus(networkAttachment *v1alpha1.NetworkAttachment) (result *v1alpha1.NetworkAttachment, err error) {
	result = &v1alpha1.NetworkAttachment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("networkattachments").
		Name(networkAttachment.Name).
		SubResource("status").
		Body(networkAttachment).
		Do().
		Into(result)
	return
}

// Delete takes name of the networkAttachment and deletes it. Returns an error if one occurs.
func (c *networkAttachments) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("networkattachments").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *networkAttachments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("networkattachments").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched networkAttachment.
func (c *networkAttachments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NetworkAttachment, err error) {
	result = &v1alpha1.NetworkAttachment{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("networkattachments").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
