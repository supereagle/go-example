package kubernetes

import (
	v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	scheme "k8s.io/client-go/kubernetes/scheme"
)

// CreateNamespace takes the representation of a namespace and creates it.  Returns the server's representation of the namespace, and an error, if there is any.
func (c *clientSet) CreateNamespace(namespace *v1.Namespace) (result *v1.Namespace, err error) {
	result = &v1.Namespace{}
	err = c.coreClient.Post().
		Resource("namespaces").
		Body(namespace).
		Do().
		Into(result)
	return
}

// DeleteNamespace takes name of the namespace and deletes it. Returns an error if one occurs.
func (c *clientSet) DeleteNamespace(name string) error {
	return c.coreClient.Delete().
		Resource("namespaces").
		Name(name).
		Body(&meta_v1.DeleteOptions{}).
		Do().
		Error()
}

// GetNamespace takes name of the namespace, and returns the corresponding namespace object, and an error if there is any.
func (c *clientSet) GetNamespace(name string) (result *v1.Namespace, err error) {
	result = &v1.Namespace{}
	err = c.coreClient.Get().
		Resource("namespaces").
		Name(name).
		VersionedParams(&meta_v1.GetOptions{}, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}
