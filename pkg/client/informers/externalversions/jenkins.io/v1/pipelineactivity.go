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

// PipelineActivityInformer provides access to a shared informer and lister for
// PipelineActivities.
type PipelineActivityInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.PipelineActivityLister
}

type pipelineActivityInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPipelineActivityInformer constructs a new informer for PipelineActivity type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPipelineActivityInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPipelineActivityInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPipelineActivityInformer constructs a new informer for PipelineActivity type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPipelineActivityInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.JenkinsV1().PipelineActivities(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.JenkinsV1().PipelineActivities(namespace).Watch(options)
			},
		},
		&jenkinsiov1.PipelineActivity{},
		resyncPeriod,
		indexers,
	)
}

func (f *pipelineActivityInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPipelineActivityInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *pipelineActivityInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&jenkinsiov1.PipelineActivity{}, f.defaultInformer)
}

func (f *pipelineActivityInformer) Lister() v1.PipelineActivityLister {
	return v1.NewPipelineActivityLister(f.Informer().GetIndexer())
}
