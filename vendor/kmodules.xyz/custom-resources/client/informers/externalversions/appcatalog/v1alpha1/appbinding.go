/*
Copyright 2018 AppsCode Inc.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	appcatalogv1alpha1 "kmodules.xyz/custom-resources/apis/appcatalog/v1alpha1"
	versioned "kmodules.xyz/custom-resources/client/clientset/versioned"
	internalinterfaces "kmodules.xyz/custom-resources/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kmodules.xyz/custom-resources/client/listers/appcatalog/v1alpha1"
)

// AppBindingInformer provides access to a shared informer and lister for
// AppBindings.
type AppBindingInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.AppBindingLister
}

type appBindingInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewAppBindingInformer constructs a new informer for AppBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAppBindingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAppBindingInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredAppBindingInformer constructs a new informer for AppBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAppBindingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AppcatalogV1alpha1().AppBindings(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AppcatalogV1alpha1().AppBindings(namespace).Watch(options)
			},
		},
		&appcatalogv1alpha1.AppBinding{},
		resyncPeriod,
		indexers,
	)
}

func (f *appBindingInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAppBindingInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *appBindingInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&appcatalogv1alpha1.AppBinding{}, f.defaultInformer)
}

func (f *appBindingInformer) Lister() v1alpha1.AppBindingLister {
	return v1alpha1.NewAppBindingLister(f.Informer().GetIndexer())
}
