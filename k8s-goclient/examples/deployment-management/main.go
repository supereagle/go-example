package main

import (
	"fmt"

	"github.com/supereagle/go-example/k8s-goclient/kubernetes"
)

func main() {
	// Set the apiserver address of your cluster
	server := "server.k8s.com:8080"

	clientSet, err := kubernetes.NewClientSet(server)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	result, err := clientSet.CreateDeployment("test", nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())
}
