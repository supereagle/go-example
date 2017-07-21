package kubernetes

import (
	v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	scheme "k8s.io/client-go/kubernetes/scheme"
)

// GetConfigMap takes name of the configMap, and returns the corresponding configMap object, and an error if there is any.
func (c *clientSet) GetConfigMap(namespace, name string) (result *v1.ConfigMap, err error) {
	result = &v1.ConfigMap{}
	err = c.coreClient.Get().
		Namespace(namespace).
		Resource("configmaps").
		Name(name).
		VersionedParams(&meta_v1.GetOptions{}, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// ListConfigMap takes label and field selectors, and returns the list of ConfigMaps that match those selectors.
func (c *clientSet) ListConfigMap(namespace string) (result *v1.ConfigMapList, err error) {
	result = &v1.ConfigMapList{}
	err = c.coreClient.Get().
		Namespace(namespace).
		Resource("configmaps").
		VersionedParams(&meta_v1.ListOptions{}, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}
