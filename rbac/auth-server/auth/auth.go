package auth

import (
	"fmt"
)

var AuthPlugins map[Service]Authorizer

func init() {
	AuthPlugins = make(map[Service]Authorizer)
}

func RegisterAuthorizer(service Service, authorizer Authorizer) error {
	if _, ok := AuthPlugins[service]; ok {
		return fmt.Errorf("Authorizer %s already exists.", service)
	}

	AuthPlugins[service] = authorizer
	return nil
}

type Authorizer interface {
	DoAuth(ar *AuthRequest) error
}

func FilterActions(actions []string, exclusions ...string) (filteredActions []string) {
	for _, action := range actions {
		excluded := false
		for _, exclusion := range exclusions {
			if action == exclusion {
				excluded = true
				break
			}
		}
		if !excluded {
			filteredActions = append(filteredActions, action)
		}
	}

	return filteredActions
}
