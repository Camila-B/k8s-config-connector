// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/documentai/v1alpha1"
	scheme "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DocumentAIProcessorsGetter has a method to return a DocumentAIProcessorInterface.
// A group's client should implement this interface.
type DocumentAIProcessorsGetter interface {
	DocumentAIProcessors(namespace string) DocumentAIProcessorInterface
}

// DocumentAIProcessorInterface has methods to work with DocumentAIProcessor resources.
type DocumentAIProcessorInterface interface {
	Create(ctx context.Context, documentAIProcessor *v1alpha1.DocumentAIProcessor, opts v1.CreateOptions) (*v1alpha1.DocumentAIProcessor, error)
	Update(ctx context.Context, documentAIProcessor *v1alpha1.DocumentAIProcessor, opts v1.UpdateOptions) (*v1alpha1.DocumentAIProcessor, error)
	UpdateStatus(ctx context.Context, documentAIProcessor *v1alpha1.DocumentAIProcessor, opts v1.UpdateOptions) (*v1alpha1.DocumentAIProcessor, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.DocumentAIProcessor, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.DocumentAIProcessorList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DocumentAIProcessor, err error)
	DocumentAIProcessorExpansion
}

// documentAIProcessors implements DocumentAIProcessorInterface
type documentAIProcessors struct {
	client rest.Interface
	ns     string
}

// newDocumentAIProcessors returns a DocumentAIProcessors
func newDocumentAIProcessors(c *DocumentaiV1alpha1Client, namespace string) *documentAIProcessors {
	return &documentAIProcessors{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the documentAIProcessor, and returns the corresponding documentAIProcessor object, and an error if there is any.
func (c *documentAIProcessors) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DocumentAIProcessor, err error) {
	result = &v1alpha1.DocumentAIProcessor{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("documentaiprocessors").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DocumentAIProcessors that match those selectors.
func (c *documentAIProcessors) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DocumentAIProcessorList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DocumentAIProcessorList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("documentaiprocessors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested documentAIProcessors.
func (c *documentAIProcessors) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("documentaiprocessors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a documentAIProcessor and creates it.  Returns the server's representation of the documentAIProcessor, and an error, if there is any.
func (c *documentAIProcessors) Create(ctx context.Context, documentAIProcessor *v1alpha1.DocumentAIProcessor, opts v1.CreateOptions) (result *v1alpha1.DocumentAIProcessor, err error) {
	result = &v1alpha1.DocumentAIProcessor{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("documentaiprocessors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(documentAIProcessor).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a documentAIProcessor and updates it. Returns the server's representation of the documentAIProcessor, and an error, if there is any.
func (c *documentAIProcessors) Update(ctx context.Context, documentAIProcessor *v1alpha1.DocumentAIProcessor, opts v1.UpdateOptions) (result *v1alpha1.DocumentAIProcessor, err error) {
	result = &v1alpha1.DocumentAIProcessor{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("documentaiprocessors").
		Name(documentAIProcessor.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(documentAIProcessor).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *documentAIProcessors) UpdateStatus(ctx context.Context, documentAIProcessor *v1alpha1.DocumentAIProcessor, opts v1.UpdateOptions) (result *v1alpha1.DocumentAIProcessor, err error) {
	result = &v1alpha1.DocumentAIProcessor{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("documentaiprocessors").
		Name(documentAIProcessor.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(documentAIProcessor).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the documentAIProcessor and deletes it. Returns an error if one occurs.
func (c *documentAIProcessors) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("documentaiprocessors").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *documentAIProcessors) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("documentaiprocessors").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched documentAIProcessor.
func (c *documentAIProcessors) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DocumentAIProcessor, err error) {
	result = &v1alpha1.DocumentAIProcessor{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("documentaiprocessors").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}