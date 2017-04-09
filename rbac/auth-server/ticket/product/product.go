package product

import (
	"fmt"

	log "github.com/Sirupsen/logrus"

	"github.com/supereagle/go-example/rbac/auth-server/auth"
	"github.com/supereagle/go-example/rbac/auth-server/ticket"
	setsUtil "github.com/supereagle/go-example/rbac/auth-server/utils/sets"
)

var (
	SupportedResourceTypes = []string{"fruits", "vegetables"}
	ResourceTypes          = setsUtil.NewStringSet(SupportedResourceTypes)

	SupportedActions = []string{"create", "delete", "update", "view"}
	Actions          = setsUtil.NewStringSet(SupportedActions)
)

func init() {
	ticket.RegisterTicketer(auth.Product, NewTicketer())
}

type ticketer struct {
}

func NewTicketer() *ticketer {
	return new(ticketer)
}

func (t *ticketer) ParseTicket(scopeStr string) (*auth.Scope, error) {
	scope, err := ticket.ParseScope(scopeStr)
	if err != nil {
		return nil, err
	}

	if !ResourceTypes.Has(scope.ResourceType) {
		err = fmt.Errorf("Resource type %s is not supported.", scope.ResourceType)
		log.Error(err)
		return nil, err
	}

	for _, action := range scope.Actions {
		if !Actions.Has(action) {
			err = fmt.Errorf("Action %s is not supported.", action)
			log.Error(err)
			return nil, err
		}
	}

	return scope, nil
}
