/*
Copyright 2022 The Knative Authors

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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	rabbitmqcomv1beta1 "knative.dev/eventing-rabbitmq/third_party/pkg/apis/rabbitmq.com/v1beta1"
	versioned "knative.dev/eventing-rabbitmq/third_party/pkg/client/clientset/versioned"
	internalinterfaces "knative.dev/eventing-rabbitmq/third_party/pkg/client/informers/externalversions/internalinterfaces"
	v1beta1 "knative.dev/eventing-rabbitmq/third_party/pkg/client/listers/rabbitmq.com/v1beta1"
)

// PermissionInformer provides access to a shared informer and lister for
// Permissions.
type PermissionInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.PermissionLister
}

type permissionInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPermissionInformer constructs a new informer for Permission type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPermissionInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPermissionInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPermissionInformer constructs a new informer for Permission type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPermissionInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RabbitmqV1beta1().Permissions(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RabbitmqV1beta1().Permissions(namespace).Watch(context.TODO(), options)
			},
		},
		&rabbitmqcomv1beta1.Permission{},
		resyncPeriod,
		indexers,
	)
}

func (f *permissionInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPermissionInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *permissionInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&rabbitmqcomv1beta1.Permission{}, f.defaultInformer)
}

func (f *permissionInformer) Lister() v1beta1.PermissionLister {
	return v1beta1.NewPermissionLister(f.Informer().GetIndexer())
}
