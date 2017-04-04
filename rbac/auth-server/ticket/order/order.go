package order

import (
	"fmt"

	log "github.com/Sirupsen/logrus"

	"github.com/supereagle/go-example/rbac/auth-server/auth"
	"github.com/supereagle/go-example/rbac/auth-server/ticket"
	setsUtil "github.com/supereagle/go-example/rbac/auth-server/utils/sets"
)

const (
	Cash   string = "cash"
	Alipay        = "alipay"
	Weixin        = "weixin"
)

var (
	SupportedResourceTypes = []string{Cash, Alipay, Weixin}
	ResourceTypes          = setsUtil.NewStringSets(SupportedResourceTypes)

	SupportedActions = []string{"create", "delete", "update", "view", "refund", "exchange", "send", "cancel"}
	Actions          = setsUtil.NewStringSets(SupportedActions)
)

func init() {
	ticket.RegisterTicketer(auth.Order, NewTicketer())
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

	if !ResourceTypes.Contains(scope.ResourceType) {
		err = fmt.Errorf("Resource type %s is not supported.", scope.ResourceType)
		log.Error(err)
		return nil, err
	}

	for _, action := range scope.Actions {
		if !Actions.Contains(action) {
			err = fmt.Errorf("Action %s is not supported.", action)
			log.Error(err)
			return nil, err
		}
	}

	return scope, nil
}
