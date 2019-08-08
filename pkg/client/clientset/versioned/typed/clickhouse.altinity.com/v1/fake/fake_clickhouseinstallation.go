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

package fake

import (
	clickhouse_altinity_com_v1 "github.com/altinity/clickhouse-operator/pkg/apis/clickhouse.altinity.com/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClickHouseInstallations implements ClickHouseInstallationInterface
type FakeClickHouseInstallations struct {
	Fake *FakeClickhouseV1
	ns   string
}

var clickhouseinstallationsResource = schema.GroupVersionResource{Group: "clickhouse.altinity.com", Version: "v1", Resource: "clickhouseinstallations"}

var clickhouseinstallationsKind = schema.GroupVersionKind{Group: "clickhouse.altinity.com", Version: "v1", Kind: "ClickHouseInstallation"}

// Get takes name of the clickHouseInstallation, and returns the corresponding clickHouseInstallation object, and an error if there is any.
func (c *FakeClickHouseInstallations) Get(name string, options v1.GetOptions) (result *clickhouse_altinity_com_v1.ClickHouseInstallation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(clickhouseinstallationsResource, c.ns, name), &clickhouse_altinity_com_v1.ClickHouseInstallation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*clickhouse_altinity_com_v1.ClickHouseInstallation), err
}

// List takes label and field selectors, and returns the list of ClickHouseInstallations that match those selectors.
func (c *FakeClickHouseInstallations) List(opts v1.ListOptions) (result *clickhouse_altinity_com_v1.ClickHouseInstallationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(clickhouseinstallationsResource, clickhouseinstallationsKind, c.ns, opts), &clickhouse_altinity_com_v1.ClickHouseInstallationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &clickhouse_altinity_com_v1.ClickHouseInstallationList{ListMeta: obj.(*clickhouse_altinity_com_v1.ClickHouseInstallationList).ListMeta}
	for _, item := range obj.(*clickhouse_altinity_com_v1.ClickHouseInstallationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clickHouseInstallations.
func (c *FakeClickHouseInstallations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(clickhouseinstallationsResource, c.ns, opts))

}

// Create takes the representation of a clickHouseInstallation and creates it.  Returns the server's representation of the clickHouseInstallation, and an error, if there is any.
func (c *FakeClickHouseInstallations) Create(clickHouseInstallation *clickhouse_altinity_com_v1.ClickHouseInstallation) (result *clickhouse_altinity_com_v1.ClickHouseInstallation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(clickhouseinstallationsResource, c.ns, clickHouseInstallation), &clickhouse_altinity_com_v1.ClickHouseInstallation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*clickhouse_altinity_com_v1.ClickHouseInstallation), err
}

// Update takes the representation of a clickHouseInstallation and updates it. Returns the server's representation of the clickHouseInstallation, and an error, if there is any.
func (c *FakeClickHouseInstallations) Update(clickHouseInstallation *clickhouse_altinity_com_v1.ClickHouseInstallation) (result *clickhouse_altinity_com_v1.ClickHouseInstallation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(clickhouseinstallationsResource, c.ns, clickHouseInstallation), &clickhouse_altinity_com_v1.ClickHouseInstallation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*clickhouse_altinity_com_v1.ClickHouseInstallation), err
}

// Delete takes name of the clickHouseInstallation and deletes it. Returns an error if one occurs.
func (c *FakeClickHouseInstallations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(clickhouseinstallationsResource, c.ns, name), &clickhouse_altinity_com_v1.ClickHouseInstallation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClickHouseInstallations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(clickhouseinstallationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &clickhouse_altinity_com_v1.ClickHouseInstallationList{})
	return err
}

// Patch applies the patch and returns the patched clickHouseInstallation.
func (c *FakeClickHouseInstallations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *clickhouse_altinity_com_v1.ClickHouseInstallation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clickhouseinstallationsResource, c.ns, name, data, subresources...), &clickhouse_altinity_com_v1.ClickHouseInstallation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*clickhouse_altinity_com_v1.ClickHouseInstallation), err
}
