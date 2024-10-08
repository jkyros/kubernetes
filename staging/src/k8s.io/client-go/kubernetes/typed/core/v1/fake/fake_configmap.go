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
	context "context"
	json "encoding/json"
	fmt "fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	corev1 "k8s.io/client-go/applyconfigurations/core/v1"
	testing "k8s.io/client-go/testing"
)

// FakeConfigMaps implements ConfigMapInterface
type FakeConfigMaps struct {
	Fake *FakeCoreV1
	ns   string
}

var configmapsResource = v1.SchemeGroupVersion.WithResource("configmaps")

var configmapsKind = v1.SchemeGroupVersion.WithKind("ConfigMap")

// Get takes name of the configMap, and returns the corresponding configMap object, and an error if there is any.
func (c *FakeConfigMaps) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ConfigMap, err error) {
	emptyResult := &v1.ConfigMap{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(configmapsResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ConfigMap), err
}

// List takes label and field selectors, and returns the list of ConfigMaps that match those selectors.
func (c *FakeConfigMaps) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ConfigMapList, err error) {
	emptyResult := &v1.ConfigMapList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(configmapsResource, configmapsKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.ConfigMapList{ListMeta: obj.(*v1.ConfigMapList).ListMeta}
	for _, item := range obj.(*v1.ConfigMapList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested configMaps.
func (c *FakeConfigMaps) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(configmapsResource, c.ns, opts))

}

// Create takes the representation of a configMap and creates it.  Returns the server's representation of the configMap, and an error, if there is any.
func (c *FakeConfigMaps) Create(ctx context.Context, configMap *v1.ConfigMap, opts metav1.CreateOptions) (result *v1.ConfigMap, err error) {
	emptyResult := &v1.ConfigMap{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(configmapsResource, c.ns, configMap, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ConfigMap), err
}

// Update takes the representation of a configMap and updates it. Returns the server's representation of the configMap, and an error, if there is any.
func (c *FakeConfigMaps) Update(ctx context.Context, configMap *v1.ConfigMap, opts metav1.UpdateOptions) (result *v1.ConfigMap, err error) {
	emptyResult := &v1.ConfigMap{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(configmapsResource, c.ns, configMap, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ConfigMap), err
}

// Delete takes name of the configMap and deletes it. Returns an error if one occurs.
func (c *FakeConfigMaps) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(configmapsResource, c.ns, name, opts), &v1.ConfigMap{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeConfigMaps) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(configmapsResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.ConfigMapList{})
	return err
}

// Patch applies the patch and returns the patched configMap.
func (c *FakeConfigMaps) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ConfigMap, err error) {
	emptyResult := &v1.ConfigMap{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(configmapsResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ConfigMap), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied configMap.
func (c *FakeConfigMaps) Apply(ctx context.Context, configMap *corev1.ConfigMapApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ConfigMap, err error) {
	if configMap == nil {
		return nil, fmt.Errorf("configMap provided to Apply must not be nil")
	}
	data, err := json.Marshal(configMap)
	if err != nil {
		return nil, err
	}
	name := configMap.Name
	if name == nil {
		return nil, fmt.Errorf("configMap.Name must be provided to Apply")
	}
	emptyResult := &v1.ConfigMap{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(configmapsResource, c.ns, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ConfigMap), err
}
