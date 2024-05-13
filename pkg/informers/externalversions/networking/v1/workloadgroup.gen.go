// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	networkingv1 "istio.io/client-go/pkg/apis/networking/v1"
	versioned "istio.io/client-go/pkg/clientset/versioned"
	internalinterfaces "istio.io/client-go/pkg/informers/externalversions/internalinterfaces"
	v1 "istio.io/client-go/pkg/listers/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// WorkloadGroupInformer provides access to a shared informer and lister for
// WorkloadGroups.
type WorkloadGroupInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.WorkloadGroupLister
}

type workloadGroupInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewWorkloadGroupInformer constructs a new informer for WorkloadGroup type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewWorkloadGroupInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredWorkloadGroupInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredWorkloadGroupInformer constructs a new informer for WorkloadGroup type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredWorkloadGroupInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkingV1().WorkloadGroups(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkingV1().WorkloadGroups(namespace).Watch(context.TODO(), options)
			},
		},
		&networkingv1.WorkloadGroup{},
		resyncPeriod,
		indexers,
	)
}

func (f *workloadGroupInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredWorkloadGroupInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *workloadGroupInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&networkingv1.WorkloadGroup{}, f.defaultInformer)
}

func (f *workloadGroupInformer) Lister() v1.WorkloadGroupLister {
	return v1.NewWorkloadGroupLister(f.Informer().GetIndexer())
}