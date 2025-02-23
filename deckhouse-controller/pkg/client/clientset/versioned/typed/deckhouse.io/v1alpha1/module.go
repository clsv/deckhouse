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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/deckhouse/deckhouse/deckhouse-controller/pkg/apis/deckhouse.io/v1alpha1"
	scheme "github.com/deckhouse/deckhouse/deckhouse-controller/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ModulesGetter has a method to return a ModuleInterface.
// A group's client should implement this interface.
type ModulesGetter interface {
	Modules() ModuleInterface
}

// ModuleInterface has methods to work with Module resources.
type ModuleInterface interface {
	Create(ctx context.Context, module *v1alpha1.Module, opts v1.CreateOptions) (*v1alpha1.Module, error)
	Update(ctx context.Context, module *v1alpha1.Module, opts v1.UpdateOptions) (*v1alpha1.Module, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Module, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ModuleList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Module, err error)
	ModuleExpansion
}

// modules implements ModuleInterface
type modules struct {
	client rest.Interface
}

// newModules returns a Modules
func newModules(c *DeckhouseV1alpha1Client) *modules {
	return &modules{
		client: c.RESTClient(),
	}
}

// Get takes name of the module, and returns the corresponding module object, and an error if there is any.
func (c *modules) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Module, err error) {
	result = &v1alpha1.Module{}
	err = c.client.Get().
		Resource("modules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Modules that match those selectors.
func (c *modules) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ModuleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ModuleList{}
	err = c.client.Get().
		Resource("modules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested modules.
func (c *modules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("modules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a module and creates it.  Returns the server's representation of the module, and an error, if there is any.
func (c *modules) Create(ctx context.Context, module *v1alpha1.Module, opts v1.CreateOptions) (result *v1alpha1.Module, err error) {
	result = &v1alpha1.Module{}
	err = c.client.Post().
		Resource("modules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(module).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a module and updates it. Returns the server's representation of the module, and an error, if there is any.
func (c *modules) Update(ctx context.Context, module *v1alpha1.Module, opts v1.UpdateOptions) (result *v1alpha1.Module, err error) {
	result = &v1alpha1.Module{}
	err = c.client.Put().
		Resource("modules").
		Name(module.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(module).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the module and deletes it. Returns an error if one occurs.
func (c *modules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("modules").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *modules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("modules").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched module.
func (c *modules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Module, err error) {
	result = &v1alpha1.Module{}
	err = c.client.Patch(pt).
		Resource("modules").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
