package ticket

import (
	"fmt"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/supereagle/go-example/rbac/auth-server/auth"
)

var TicketPlugins map[auth.Service]Ticketer

func init() {
	TicketPlugins = make(map[auth.Service]Ticketer)
}

func RegisterTicketer(service auth.Service, ticketer Ticketer) error {
	if _, ok := TicketPlugins[service]; ok {
		return fmt.Errorf("Ticketer %s already exists.", service)
	}

	TicketPlugins[service] = ticketer
	return nil
}

type Ticketer interface {
	ParseTicket(scopeStr string) (*auth.Scope, error)
}

func ParseScope(scopeStr string) (*auth.Scope, error) {
	if scopeStr == "" {
		err := fmt.Errorf("The scope is empty.")
		log.Error(err)
		return nil, err
	}

	parts := strings.Split(scopeStr, ":")
	if len(parts) != 3 {
		err := fmt.Errorf("The scope format has error.")
		log.Error(err)
		return nil, err
	}

	actions := strings.Split(parts[2], ",")
	scope := &auth.Scope{
		ResourceType: parts[0],
		ResourceName: parts[1],
		Actions:      actions,
	}

	return scope, nil
}
