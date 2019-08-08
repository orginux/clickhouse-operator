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

package v1

import (
	v1 "github.com/altinity/clickhouse-operator/pkg/apis/clickhouse.altinity.com/v1"
	scheme "github.com/altinity/clickhouse-operator/pkg/client/clientset/versioned/scheme"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClickHouseInstallationsGetter has a method to return a ClickHouseInstallationInterface.
// A group's client should implement this interface.
type ClickHouseInstallationsGetter interface {
	ClickHouseInstallations(namespace string) ClickHouseInstallationInterface
}

// ClickHouseInstallationInterface has methods to work with ClickHouseInstallation resources.
type ClickHouseInstallationInterface interface {
	Create(*v1.ClickHouseInstallation) (*v1.ClickHouseInstallation, error)
	Update(*v1.ClickHouseInstallation) (*v1.ClickHouseInstallation, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.ClickHouseInstallation, error)
	List(opts meta_v1.ListOptions) (*v1.ClickHouseInstallationList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ClickHouseInstallation, err error)
	ClickHouseInstallationExpansion
}

// clickHouseInstallations implements ClickHouseInstallationInterface
type clickHouseInstallations struct {
	client rest.Interface
	ns     string
}

// newClickHouseInstallations returns a ClickHouseInstallations
func newClickHouseInstallations(c *ClickhouseV1Client, namespace string) *clickHouseInstallations {
	return &clickHouseInstallations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the clickHouseInstallation, and returns the corresponding clickHouseInstallation object, and an error if there is any.
func (c *clickHouseInstallations) Get(name string, options meta_v1.GetOptions) (result *v1.ClickHouseInstallation, err error) {
	result = &v1.ClickHouseInstallation{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clickhouseinstallations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClickHouseInstallations that match those selectors.
func (c *clickHouseInstallations) List(opts meta_v1.ListOptions) (result *v1.ClickHouseInstallationList, err error) {
	result = &v1.ClickHouseInstallationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clickhouseinstallations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clickHouseInstallations.
func (c *clickHouseInstallations) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("clickhouseinstallations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a clickHouseInstallation and creates it.  Returns the server's representation of the clickHouseInstallation, and an error, if there is any.
func (c *clickHouseInstallations) Create(clickHouseInstallation *v1.ClickHouseInstallation) (result *v1.ClickHouseInstallation, err error) {
	result = &v1.ClickHouseInstallation{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("clickhouseinstallations").
		Body(clickHouseInstallation).
		Do().
		Into(result)
	return
}

// Update takes the representation of a clickHouseInstallation and updates it. Returns the server's representation of the clickHouseInstallation, and an error, if there is any.
func (c *clickHouseInstallations) Update(clickHouseInstallation *v1.ClickHouseInstallation) (result *v1.ClickHouseInstallation, err error) {
	result = &v1.ClickHouseInstallation{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clickhouseinstallations").
		Name(clickHouseInstallation.Name).
		Body(clickHouseInstallation).
		Do().
		Into(result)
	return
}

// Delete takes name of the clickHouseInstallation and deletes it. Returns an error if one occurs.
func (c *clickHouseInstallations) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clickhouseinstallations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clickHouseInstallations) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clickhouseinstallations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched clickHouseInstallation.
func (c *clickHouseInstallations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ClickHouseInstallation, err error) {
	result = &v1.ClickHouseInstallation{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("clickhouseinstallations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
