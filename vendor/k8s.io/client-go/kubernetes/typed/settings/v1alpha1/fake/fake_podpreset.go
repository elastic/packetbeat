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

	v1alpha1 "k8s.io/api/settings/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePodPresets implements PodPresetInterface
type FakePodPresets struct {
	Fake *FakeSettingsV1alpha1
	ns   string
}

var podpresetsResource = schema.GroupVersionResource{Group: "settings.k8s.io", Version: "v1alpha1", Resource: "podpresets"}

var podpresetsKind = schema.GroupVersionKind{Group: "settings.k8s.io", Version: "v1alpha1", Kind: "PodPreset"}

// Get takes name of the podPreset, and returns the corresponding podPreset object, and an error if there is any.
func (c *FakePodPresets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PodPreset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(podpresetsResource, c.ns, name), &v1alpha1.PodPreset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodPreset), err
}

// List takes label and field selectors, and returns the list of PodPresets that match those selectors.
func (c *FakePodPresets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PodPresetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(podpresetsResource, podpresetsKind, c.ns, opts), &v1alpha1.PodPresetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PodPresetList{ListMeta: obj.(*v1alpha1.PodPresetList).ListMeta}
	for _, item := range obj.(*v1alpha1.PodPresetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested podPresets.
func (c *FakePodPresets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(podpresetsResource, c.ns, opts))

}

// Create takes the representation of a podPreset and creates it.  Returns the server's representation of the podPreset, and an error, if there is any.
func (c *FakePodPresets) Create(ctx context.Context, podPreset *v1alpha1.PodPreset, opts v1.CreateOptions) (result *v1alpha1.PodPreset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(podpresetsResource, c.ns, podPreset), &v1alpha1.PodPreset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodPreset), err
}

// Update takes the representation of a podPreset and updates it. Returns the server's representation of the podPreset, and an error, if there is any.
func (c *FakePodPresets) Update(ctx context.Context, podPreset *v1alpha1.PodPreset, opts v1.UpdateOptions) (result *v1alpha1.PodPreset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(podpresetsResource, c.ns, podPreset), &v1alpha1.PodPreset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodPreset), err
}

// Delete takes name of the podPreset and deletes it. Returns an error if one occurs.
func (c *FakePodPresets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(podpresetsResource, c.ns, name), &v1alpha1.PodPreset{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePodPresets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(podpresetsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PodPresetList{})
	return err
}

// Patch applies the patch and returns the patched podPreset.
func (c *FakePodPresets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PodPreset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(podpresetsResource, c.ns, name, pt, data, subresources...), &v1alpha1.PodPreset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodPreset), err
}
