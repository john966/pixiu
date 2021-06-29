/*
Copyright 2021 The Pixiu Authors.

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

	v1alpha1 "github.com/caoyingjunz/pixiu/pkg/apis/apps/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAdvancedImages implements AdvancedImageInterface
type FakeAdvancedImages struct {
	Fake *FakeAppsV1alpha1
	ns   string
}

var advancedimagesResource = schema.GroupVersionResource{Group: "apps.pixiu.io", Version: "v1alpha1", Resource: "advancedimages"}

var advancedimagesKind = schema.GroupVersionKind{Group: "apps.pixiu.io", Version: "v1alpha1", Kind: "AdvancedImage"}

// Get takes name of the advancedImage, and returns the corresponding advancedImage object, and an error if there is any.
func (c *FakeAdvancedImages) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AdvancedImage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(advancedimagesResource, c.ns, name), &v1alpha1.AdvancedImage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AdvancedImage), err
}

// List takes label and field selectors, and returns the list of AdvancedImages that match those selectors.
func (c *FakeAdvancedImages) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AdvancedImageList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(advancedimagesResource, advancedimagesKind, c.ns, opts), &v1alpha1.AdvancedImageList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AdvancedImageList{ListMeta: obj.(*v1alpha1.AdvancedImageList).ListMeta}
	for _, item := range obj.(*v1alpha1.AdvancedImageList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested advancedImages.
func (c *FakeAdvancedImages) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(advancedimagesResource, c.ns, opts))

}

// Create takes the representation of a advancedImage and creates it.  Returns the server's representation of the advancedImage, and an error, if there is any.
func (c *FakeAdvancedImages) Create(ctx context.Context, advancedImage *v1alpha1.AdvancedImage, opts v1.CreateOptions) (result *v1alpha1.AdvancedImage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(advancedimagesResource, c.ns, advancedImage), &v1alpha1.AdvancedImage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AdvancedImage), err
}

// Update takes the representation of a advancedImage and updates it. Returns the server's representation of the advancedImage, and an error, if there is any.
func (c *FakeAdvancedImages) Update(ctx context.Context, advancedImage *v1alpha1.AdvancedImage, opts v1.UpdateOptions) (result *v1alpha1.AdvancedImage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(advancedimagesResource, c.ns, advancedImage), &v1alpha1.AdvancedImage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AdvancedImage), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAdvancedImages) UpdateStatus(ctx context.Context, advancedImage *v1alpha1.AdvancedImage, opts v1.UpdateOptions) (*v1alpha1.AdvancedImage, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(advancedimagesResource, "status", c.ns, advancedImage), &v1alpha1.AdvancedImage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AdvancedImage), err
}

// Delete takes name of the advancedImage and deletes it. Returns an error if one occurs.
func (c *FakeAdvancedImages) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(advancedimagesResource, c.ns, name), &v1alpha1.AdvancedImage{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAdvancedImages) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(advancedimagesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AdvancedImageList{})
	return err
}

// Patch applies the patch and returns the patched advancedImage.
func (c *FakeAdvancedImages) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AdvancedImage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(advancedimagesResource, c.ns, name, pt, data, subresources...), &v1alpha1.AdvancedImage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AdvancedImage), err
}
