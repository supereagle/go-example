package auth

import (
	"fmt"
)

var AuthPlugins map[Service]Authorizer

func init() {
	AuthPlugins = make(map[Service]Authorizer)
}

func RegistryPlugin(service Service, authorizer Authorizer) error {
	if _, ok := AuthPlugins[service]; ok {
		return fmt.Errorf("Authorizer %s already exists.", service)
	}

	AuthPlugins[service] = authorizer
	return nil
}

type Authorizer interface {
	DoAuth() bool
}
