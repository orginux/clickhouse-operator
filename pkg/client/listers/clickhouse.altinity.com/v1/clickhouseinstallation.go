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
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClickHouseInstallationLister helps list ClickHouseInstallations.
type ClickHouseInstallationLister interface {
	// List lists all ClickHouseInstallations in the indexer.
	List(selector labels.Selector) (ret []*v1.ClickHouseInstallation, err error)
	// ClickHouseInstallations returns an object that can list and get ClickHouseInstallations.
	ClickHouseInstallations(namespace string) ClickHouseInstallationNamespaceLister
	ClickHouseInstallationListerExpansion
}

// clickHouseInstallationLister implements the ClickHouseInstallationLister interface.
type clickHouseInstallationLister struct {
	indexer cache.Indexer
}

// NewClickHouseInstallationLister returns a new ClickHouseInstallationLister.
func NewClickHouseInstallationLister(indexer cache.Indexer) ClickHouseInstallationLister {
	return &clickHouseInstallationLister{indexer: indexer}
}

// List lists all ClickHouseInstallations in the indexer.
func (s *clickHouseInstallationLister) List(selector labels.Selector) (ret []*v1.ClickHouseInstallation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ClickHouseInstallation))
	})
	return ret, err
}

// ClickHouseInstallations returns an object that can list and get ClickHouseInstallations.
func (s *clickHouseInstallationLister) ClickHouseInstallations(namespace string) ClickHouseInstallationNamespaceLister {
	return clickHouseInstallationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ClickHouseInstallationNamespaceLister helps list and get ClickHouseInstallations.
type ClickHouseInstallationNamespaceLister interface {
	// List lists all ClickHouseInstallations in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.ClickHouseInstallation, err error)
	// Get retrieves the ClickHouseInstallation from the indexer for a given namespace and name.
	Get(name string) (*v1.ClickHouseInstallation, error)
	ClickHouseInstallationNamespaceListerExpansion
}

// clickHouseInstallationNamespaceLister implements the ClickHouseInstallationNamespaceLister
// interface.
type clickHouseInstallationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ClickHouseInstallations in the indexer for a given namespace.
func (s clickHouseInstallationNamespaceLister) List(selector labels.Selector) (ret []*v1.ClickHouseInstallation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ClickHouseInstallation))
	})
	return ret, err
}

// Get retrieves the ClickHouseInstallation from the indexer for a given namespace and name.
func (s clickHouseInstallationNamespaceLister) Get(name string) (*v1.ClickHouseInstallation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("clickhouseinstallation"), name)
	}
	return obj.(*v1.ClickHouseInstallation), nil
}
