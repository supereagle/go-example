package kubernetes

import (
	v1beta1 "k8s.io/api/extensions/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	scheme "k8s.io/client-go/kubernetes/scheme"
)

// CreateIngress takes the representation of a ingress and creates it.  Returns the server's representation of the ingress, and an error, if there is any.
func (c *clientSet) CreateIngress(namespace string, ingress *v1beta1.Ingress) (result *v1beta1.Ingress, err error) {
	result = &v1beta1.Ingress{}
	err = c.extensionsClient.Post().
		Namespace(namespace).
		Resource("ingresses").
		Body(ingress).
		Do().
		Into(result)
	return
}

// DeleteIngress takes name of the ingress and deletes it. Returns an error if one occurs.
func (c *clientSet) DeleteIngress(namespace, name string) error {
	return c.extensionsClient.Delete().
		Namespace(namespace).
		Resource("ingresses").
		Name(name).
		Body(&v1.DeleteOptions{}).
		Do().
		Error()
}

// GetIngress takes name of the ingress, and returns the corresponding ingress object, and an error if there is any.
func (c *clientSet) GetIngress(namespace, name string) (result *v1beta1.Ingress, err error) {
	result = &v1beta1.Ingress{}
	err = c.extensionsClient.Get().
		Namespace(namespace).
		Resource("ingresses").
		Name(name).
		VersionedParams(&v1.GetOptions{}, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}
