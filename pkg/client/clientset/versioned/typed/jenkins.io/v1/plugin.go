// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/jenkins-x/jx/v2/pkg/apis/jenkins.io/v1"
	scheme "github.com/jenkins-x/jx/v2/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PluginsGetter has a method to return a PluginInterface.
// A group's client should implement this interface.
type PluginsGetter interface {
	Plugins(namespace string) PluginInterface
}

// PluginInterface has methods to work with Plugin resources.
type PluginInterface interface {
	Create(*v1.Plugin) (*v1.Plugin, error)
	Update(*v1.Plugin) (*v1.Plugin, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.Plugin, error)
	List(opts metav1.ListOptions) (*v1.PluginList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Plugin, err error)
	PluginExpansion
}

// plugins implements PluginInterface
type plugins struct {
	client rest.Interface
	ns     string
}

// newPlugins returns a Plugins
func newPlugins(c *JenkinsV1Client, namespace string) *plugins {
	return &plugins{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the plugin, and returns the corresponding plugin object, and an error if there is any.
func (c *plugins) Get(name string, options metav1.GetOptions) (result *v1.Plugin, err error) {
	result = &v1.Plugin{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("plugins").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Plugins that match those selectors.
func (c *plugins) List(opts metav1.ListOptions) (result *v1.PluginList, err error) {
	result = &v1.PluginList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("plugins").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested plugins.
func (c *plugins) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("plugins").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a plugin and creates it.  Returns the server's representation of the plugin, and an error, if there is any.
func (c *plugins) Create(plugin *v1.Plugin) (result *v1.Plugin, err error) {
	result = &v1.Plugin{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("plugins").
		Body(plugin).
		Do().
		Into(result)
	return
}

// Update takes the representation of a plugin and updates it. Returns the server's representation of the plugin, and an error, if there is any.
func (c *plugins) Update(plugin *v1.Plugin) (result *v1.Plugin, err error) {
	result = &v1.Plugin{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("plugins").
		Name(plugin.Name).
		Body(plugin).
		Do().
		Into(result)
	return
}

// Delete takes name of the plugin and deletes it. Returns an error if one occurs.
func (c *plugins) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("plugins").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *plugins) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("plugins").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched plugin.
func (c *plugins) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Plugin, err error) {
	result = &v1.Plugin{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("plugins").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
