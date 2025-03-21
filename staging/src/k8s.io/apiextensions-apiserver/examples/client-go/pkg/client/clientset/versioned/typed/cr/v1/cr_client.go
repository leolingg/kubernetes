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

package v1

import (
	http "net/http"

	crv1 "k8s.io/apiextensions-apiserver/examples/client-go/pkg/apis/cr/v1"
	scheme "k8s.io/apiextensions-apiserver/examples/client-go/pkg/client/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type CrV1Interface interface {
	RESTClient() rest.Interface
	ExamplesGetter
}

// CrV1Client is used to interact with features provided by the cr.example.apiextensions.k8s.io group.
type CrV1Client struct {
	restClient rest.Interface
}

func (c *CrV1Client) Examples(namespace string) ExampleInterface {
	return newExamples(c, namespace)
}

// NewForConfig creates a new CrV1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*CrV1Client, error) {
	config := *c
	setConfigDefaults(&config)
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new CrV1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*CrV1Client, error) {
	config := *c
	setConfigDefaults(&config)
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &CrV1Client{client}, nil
}

// NewForConfigOrDie creates a new CrV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *CrV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new CrV1Client for the given RESTClient.
func New(c rest.Interface) *CrV1Client {
	return &CrV1Client{c}
}

func setConfigDefaults(config *rest.Config) {
	gv := crv1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = rest.CodecFactoryForGeneratedClient(scheme.Scheme, scheme.Codecs).WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *CrV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
