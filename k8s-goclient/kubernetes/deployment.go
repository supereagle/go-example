package kubernetes

import (
	v1beta1 "k8s.io/api/apps/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	scheme "k8s.io/client-go/kubernetes/scheme"
)

// CreateDeployment takes the representation of a deployment and creates it.  Returns the server's representation of the deployment, and an error, if there is any.
func (c *clientSet) CreateDeployment(namespace string, deployment *v1beta1.Deployment) (result *v1beta1.Deployment, err error) {
	result = &v1beta1.Deployment{}
	err = c.appsClient.Post().
		Namespace(namespace).
		Resource("deployments").
		Body(deployment).
		Do().
		Into(result)
	return
}

// DeleteDeployment takes name of the deployment and deletes it. Returns an error if one occurs.
func (c *clientSet) DeleteDeployment(namespace, name string) error {
	propagation := v1.DeletePropagationForeground
	options := &v1.DeleteOptions{
		PropagationPolicy: &propagation,
	}
	return c.appsClient.Delete().
		Namespace(namespace).
		Resource("deployments").
		Name(name).
		Body(options).
		Do().
		Error()
}

// UpdateDeployment takes the representation of a deployment and updates it. Returns the server's representation of the deployment, and an error, if there is any.
func (c *clientSet) UpdateDeployment(namespace string, deployment *v1beta1.Deployment) (result *v1beta1.Deployment, err error) {
	result = &v1beta1.Deployment{}
	err = c.appsClient.Put().
		Namespace(namespace).
		Resource("deployments").
		Name(deployment.Name).
		Body(deployment).
		Do().
		Into(result)
	return
}

// GetDeployment takes name of the deployment, and returns the corresponding deployment object, and an error if there is any.
func (c *clientSet) GetDeployment(namespace, name string) (result *v1beta1.Deployment, err error) {
	result = &v1beta1.Deployment{}
	err = c.appsClient.Get().
		Namespace(namespace).
		Resource("deployments").
		Name(name).
		VersionedParams(&v1.GetOptions{}, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// WatchDeployment returns a watch.Interface that watches the requested deployments.
func (c *clientSet) WatchDeployment(namespace, name string) (watch.Interface, error) {
	opts := v1.ListOptions{
		Watch:         true,
		FieldSelector: "metadata.name=" + name,
	}
	return c.appsClient.Get().
		Namespace(namespace).
		Resource("deployments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}
