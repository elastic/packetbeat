/*
Copyright The Kubernetes Authors.

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

package v1

import (
	"context"
	"time"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	scheme "k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
)

// ResourceQuotasGetter has a method to return a ResourceQuotaInterface.
// A group's client should implement this interface.
type ResourceQuotasGetter interface {
	ResourceQuotas(namespace string) ResourceQuotaInterface
}

// ResourceQuotaInterface has methods to work with ResourceQuota resources.
type ResourceQuotaInterface interface {
	Create(ctx context.Context, resourceQuota *v1.ResourceQuota, opts metav1.CreateOptions) (*v1.ResourceQuota, error)
	Update(ctx context.Context, resourceQuota *v1.ResourceQuota, opts metav1.UpdateOptions) (*v1.ResourceQuota, error)
	UpdateStatus(ctx context.Context, resourceQuota *v1.ResourceQuota, opts metav1.UpdateOptions) (*v1.ResourceQuota, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.ResourceQuota, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ResourceQuotaList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ResourceQuota, err error)
	ResourceQuotaExpansion
}

// resourceQuotas implements ResourceQuotaInterface
type resourceQuotas struct {
	client rest.Interface
	ns     string
}

// newResourceQuotas returns a ResourceQuotas
func newResourceQuotas(c *CoreV1Client, namespace string) *resourceQuotas {
	return &resourceQuotas{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the resourceQuota, and returns the corresponding resourceQuota object, and an error if there is any.
func (c *resourceQuotas) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ResourceQuota, err error) {
	result = &v1.ResourceQuota{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("resourcequotas").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ResourceQuotas that match those selectors.
func (c *resourceQuotas) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ResourceQuotaList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ResourceQuotaList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("resourcequotas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested resourceQuotas.
func (c *resourceQuotas) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("resourcequotas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a resourceQuota and creates it.  Returns the server's representation of the resourceQuota, and an error, if there is any.
func (c *resourceQuotas) Create(ctx context.Context, resourceQuota *v1.ResourceQuota, opts metav1.CreateOptions) (result *v1.ResourceQuota, err error) {
	result = &v1.ResourceQuota{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("resourcequotas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(resourceQuota).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a resourceQuota and updates it. Returns the server's representation of the resourceQuota, and an error, if there is any.
func (c *resourceQuotas) Update(ctx context.Context, resourceQuota *v1.ResourceQuota, opts metav1.UpdateOptions) (result *v1.ResourceQuota, err error) {
	result = &v1.ResourceQuota{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("resourcequotas").
		Name(resourceQuota.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(resourceQuota).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *resourceQuotas) UpdateStatus(ctx context.Context, resourceQuota *v1.ResourceQuota, opts metav1.UpdateOptions) (result *v1.ResourceQuota, err error) {
	result = &v1.ResourceQuota{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("resourcequotas").
		Name(resourceQuota.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(resourceQuota).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the resourceQuota and deletes it. Returns an error if one occurs.
func (c *resourceQuotas) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("resourcequotas").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *resourceQuotas) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("resourcequotas").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched resourceQuota.
func (c *resourceQuotas) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ResourceQuota, err error) {
	result = &v1.ResourceQuota{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("resourcequotas").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
