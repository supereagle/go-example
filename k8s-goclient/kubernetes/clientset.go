package kubernetes

import (
	"fmt"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type clientSet struct {
	appsClient       rest.Interface
	coreClient       rest.Interface
	extensionsClient rest.Interface
}

func NewClientSet(server string) (*clientSet, error) {
	config := &rest.Config{
		Host: server,
	}

	fullClientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("new k8s client set with error %s\n", err.Error())
	}

	return &clientSet{
		appsClient:       fullClientSet.AppsV1beta1().RESTClient(),
		coreClient:       fullClientSet.CoreV1().RESTClient(),
		extensionsClient: fullClientSet.ExtensionsV1beta1().RESTClient(),
	}, nil
}
