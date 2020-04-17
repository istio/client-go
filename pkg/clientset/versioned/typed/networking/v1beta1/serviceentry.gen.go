// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	"time"

	v1beta1 "istio.io/client-go/pkg/apis/networking/v1beta1"
	scheme "istio.io/client-go/pkg/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ServiceEntriesGetter has a method to return a ServiceEntryInterface.
// A group's client should implement this interface.
type ServiceEntriesGetter interface {
	ServiceEntries(namespace string) ServiceEntryInterface
}

// ServiceEntryInterface has methods to work with ServiceEntry resources.
type ServiceEntryInterface interface {
	Create(ctx context.Context, serviceEntry *v1beta1.ServiceEntry, opts v1.CreateOptions) (*v1beta1.ServiceEntry, error)
	Update(ctx context.Context, serviceEntry *v1beta1.ServiceEntry, opts v1.UpdateOptions) (*v1beta1.ServiceEntry, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.ServiceEntry, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.ServiceEntryList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ServiceEntry, err error)
	ServiceEntryExpansion
}

// serviceEntries implements ServiceEntryInterface
type serviceEntries struct {
	client rest.Interface
	ns     string
}

// newServiceEntries returns a ServiceEntries
func newServiceEntries(c *NetworkingV1beta1Client, namespace string) *serviceEntries {
	return &serviceEntries{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the serviceEntry, and returns the corresponding serviceEntry object, and an error if there is any.
func (c *serviceEntries) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ServiceEntry, err error) {
	result = &v1beta1.ServiceEntry{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("serviceentries").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ServiceEntries that match those selectors.
func (c *serviceEntries) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ServiceEntryList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.ServiceEntryList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("serviceentries").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested serviceEntries.
func (c *serviceEntries) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("serviceentries").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a serviceEntry and creates it.  Returns the server's representation of the serviceEntry, and an error, if there is any.
func (c *serviceEntries) Create(ctx context.Context, serviceEntry *v1beta1.ServiceEntry, opts v1.CreateOptions) (result *v1beta1.ServiceEntry, err error) {
	result = &v1beta1.ServiceEntry{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("serviceentries").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(serviceEntry).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a serviceEntry and updates it. Returns the server's representation of the serviceEntry, and an error, if there is any.
func (c *serviceEntries) Update(ctx context.Context, serviceEntry *v1beta1.ServiceEntry, opts v1.UpdateOptions) (result *v1beta1.ServiceEntry, err error) {
	result = &v1beta1.ServiceEntry{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("serviceentries").
		Name(serviceEntry.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(serviceEntry).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the serviceEntry and deletes it. Returns an error if one occurs.
func (c *serviceEntries) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("serviceentries").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *serviceEntries) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("serviceentries").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched serviceEntry.
func (c *serviceEntries) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ServiceEntry, err error) {
	result = &v1beta1.ServiceEntry{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("serviceentries").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
