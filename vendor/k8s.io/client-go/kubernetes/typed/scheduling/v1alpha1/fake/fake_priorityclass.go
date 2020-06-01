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

package fake

import (
	"context"

	v1alpha1 "k8s.io/api/scheduling/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePriorityClasses implements PriorityClassInterface
type FakePriorityClasses struct {
	Fake *FakeSchedulingV1alpha1
}

var priorityclassesResource = schema.GroupVersionResource{Group: "scheduling.k8s.io", Version: "v1alpha1", Resource: "priorityclasses"}

var priorityclassesKind = schema.GroupVersionKind{Group: "scheduling.k8s.io", Version: "v1alpha1", Kind: "PriorityClass"}

// Get takes name of the priorityClass, and returns the corresponding priorityClass object, and an error if there is any.
func (c *FakePriorityClasses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PriorityClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(priorityclassesResource, name), &v1alpha1.PriorityClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PriorityClass), err
}

// List takes label and field selectors, and returns the list of PriorityClasses that match those selectors.
func (c *FakePriorityClasses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PriorityClassList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(priorityclassesResource, priorityclassesKind, opts), &v1alpha1.PriorityClassList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PriorityClassList{ListMeta: obj.(*v1alpha1.PriorityClassList).ListMeta}
	for _, item := range obj.(*v1alpha1.PriorityClassList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested priorityClasses.
func (c *FakePriorityClasses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(priorityclassesResource, opts))
}

// Create takes the representation of a priorityClass and creates it.  Returns the server's representation of the priorityClass, and an error, if there is any.
func (c *FakePriorityClasses) Create(ctx context.Context, priorityClass *v1alpha1.PriorityClass, opts v1.CreateOptions) (result *v1alpha1.PriorityClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(priorityclassesResource, priorityClass), &v1alpha1.PriorityClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PriorityClass), err
}

// Update takes the representation of a priorityClass and updates it. Returns the server's representation of the priorityClass, and an error, if there is any.
func (c *FakePriorityClasses) Update(ctx context.Context, priorityClass *v1alpha1.PriorityClass, opts v1.UpdateOptions) (result *v1alpha1.PriorityClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(priorityclassesResource, priorityClass), &v1alpha1.PriorityClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PriorityClass), err
}

// Delete takes name of the priorityClass and deletes it. Returns an error if one occurs.
func (c *FakePriorityClasses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(priorityclassesResource, name), &v1alpha1.PriorityClass{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePriorityClasses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(priorityclassesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PriorityClassList{})
	return err
}

// Patch applies the patch and returns the patched priorityClass.
func (c *FakePriorityClasses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PriorityClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(priorityclassesResource, name, pt, data, subresources...), &v1alpha1.PriorityClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PriorityClass), err
}
