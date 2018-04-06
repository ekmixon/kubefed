/*
Copyright 2018 The Kubernetes Authors.

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

// This file was automatically generated by informer-gen

package internalversion

import (
	federation "github.com/marun/federation-v2/pkg/apis/federation"
	internalclientset "github.com/marun/federation-v2/pkg/client/clientset_generated/internalclientset"
	internalinterfaces "github.com/marun/federation-v2/pkg/client/informers_generated/internalversion/internalinterfaces"
	internalversion "github.com/marun/federation-v2/pkg/client/listers_generated/federation/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// FederatedServiceInformer provides access to a shared informer and lister for
// FederatedServices.
type FederatedServiceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.FederatedServiceLister
}

type federatedServiceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewFederatedServiceInformer constructs a new informer for FederatedService type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFederatedServiceInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredFederatedServiceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredFederatedServiceInformer constructs a new informer for FederatedService type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFederatedServiceInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Federation().FederatedServices(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Federation().FederatedServices(namespace).Watch(options)
			},
		},
		&federation.FederatedService{},
		resyncPeriod,
		indexers,
	)
}

func (f *federatedServiceInformer) defaultInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredFederatedServiceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *federatedServiceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&federation.FederatedService{}, f.defaultInformer)
}

func (f *federatedServiceInformer) Lister() internalversion.FederatedServiceLister {
	return internalversion.NewFederatedServiceLister(f.Informer().GetIndexer())
}
