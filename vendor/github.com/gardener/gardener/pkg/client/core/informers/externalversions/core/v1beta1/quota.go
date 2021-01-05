/*
Copyright (c) 2021 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

package v1beta1

import (
	"context"
	time "time"

	corev1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	versioned "github.com/gardener/gardener/pkg/client/core/clientset/versioned"
	internalinterfaces "github.com/gardener/gardener/pkg/client/core/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/gardener/gardener/pkg/client/core/listers/core/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// QuotaInformer provides access to a shared informer and lister for
// Quotas.
type QuotaInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.QuotaLister
}

type quotaInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewQuotaInformer constructs a new informer for Quota type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewQuotaInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredQuotaInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredQuotaInformer constructs a new informer for Quota type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredQuotaInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1beta1().Quotas(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1beta1().Quotas(namespace).Watch(context.TODO(), options)
			},
		},
		&corev1beta1.Quota{},
		resyncPeriod,
		indexers,
	)
}

func (f *quotaInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredQuotaInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *quotaInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&corev1beta1.Quota{}, f.defaultInformer)
}

func (f *quotaInformer) Lister() v1beta1.QuotaLister {
	return v1beta1.NewQuotaLister(f.Informer().GetIndexer())
}
