// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	jenkinsiov1 "github.com/jenkins-x/jx/v2/pkg/apis/jenkins.io/v1"
	versioned "github.com/jenkins-x/jx/v2/pkg/client/clientset/versioned"
	internalinterfaces "github.com/jenkins-x/jx/v2/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/jenkins-x/jx/v2/pkg/client/listers/jenkins.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// EnvironmentInformer provides access to a shared informer and lister for
// Environments.
type EnvironmentInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.EnvironmentLister
}

type environmentInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewEnvironmentInformer constructs a new informer for Environment type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewEnvironmentInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredEnvironmentInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredEnvironmentInformer constructs a new informer for Environment type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredEnvironmentInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.JenkinsV1().Environments(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.JenkinsV1().Environments(namespace).Watch(options)
			},
		},
		&jenkinsiov1.Environment{},
		resyncPeriod,
		indexers,
	)
}

func (f *environmentInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredEnvironmentInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *environmentInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&jenkinsiov1.Environment{}, f.defaultInformer)
}

func (f *environmentInformer) Lister() v1.EnvironmentLister {
	return v1.NewEnvironmentLister(f.Informer().GetIndexer())
}
