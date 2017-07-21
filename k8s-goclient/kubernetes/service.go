package kubernetes

import (
	v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	scheme "k8s.io/client-go/kubernetes/scheme"
)

// CreateService takes the representation of a service and creates it.  Returns the server's representation of the service, and an error, if there is any.
func (c *clientSet) CreateService(namespace string, service *v1.Service) (result *v1.Service, err error) {
	result = &v1.Service{}
	err = c.coreClient.Post().
		Namespace(namespace).
		Resource("services").
		Body(service).
		Do().
		Into(result)
	return
}

// DeleteService takes name of the service and deletes it. Returns an error if one occurs.
func (c *clientSet) DeleteService(namespace, name string) error {
	return c.coreClient.Delete().
		Namespace(namespace).
		Resource("services").
		Name(name).
		Body(&meta_v1.DeleteOptions{}).
		Do().
		Error()
}

// GetService takes name of the service, and returns the corresponding service object, and an error if there is any.
func (c *clientSet) GetService(namespace, name string) (result *v1.Service, err error) {
	result = &v1.Service{}
	err = c.coreClient.Get().
		Namespace(namespace).
		Resource("services").
		Name(name).
		VersionedParams(&meta_v1.GetOptions{}, scheme.ParameterCodec).
		Do().
		Into(result)

	return
}

// ListService takes label and field selectors, and returns the list of Services that match those selectors.
func (c *clientSet) ListService(namespace string) (result *v1.ServiceList, err error) {
	result = &v1.ServiceList{}
	err = c.coreClient.Get().
		Namespace(namespace).
		Resource("services").
		VersionedParams(&meta_v1.ListOptions{}, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}
