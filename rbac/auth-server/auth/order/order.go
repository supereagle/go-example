package order

import (
	"fmt"
	"regexp"
	"strings"

	log "github.com/Sirupsen/logrus"

	"github.com/supereagle/go-example/rbac/auth-server/auth"
	orderticket "github.com/supereagle/go-example/rbac/auth-server/ticket/order"
)

// Only support 2 types of users: master and common user.
const (
	orderUserPattern = "^order-(master|user)"
	masterPrefix     = "order-master"
)

func init() {
	auth.RegisterAuthorizer(auth.Order, NewAuthorizer())
}

type authorizer struct {
}

func NewAuthorizer() *authorizer {
	return new(authorizer)
}

func (a *authorizer) DoAuth(ar *auth.AuthRequest) error {
	// AuthN
	if !a.Authenticate(ar) {
		err := fmt.Errorf("AuthN failed for user %s", ar.Username)
		log.Error(err)
		return err
	}

	// AuthZ
	if !isMaster(ar.Username) {
		ar.Scope.Actions = auth.FilterActions(ar.Scope.Actions, "create", "delete", "cancel")

		if isCashOrder(ar.Scope.ResourceType) {
			ar.Scope.Actions = auth.FilterActions(ar.Scope.Actions, "refund")
		}
	}

	return nil
}

func (a *authorizer) Authenticate(ar *auth.AuthRequest) bool {
	if isCorrectUser(ar.Username, ar.Password) {
		return true
	}

	return false
}

func (a *authorizer) Authorize(ar *auth.AuthRequest) bool {
	return false
}

// isCorrectUser Verifies whether the user is correct. In this demo, just verify
// the username. In practice, need to verify the user in DB.
func isCorrectUser(username string, password string) bool {
	if username == "" || password == "" {
		log.Error("Username or password is empty")
		return false
	}

	if matched, _ := regexp.MatchString(orderUserPattern, username); matched {
		return true
	}

	return false
}

func isMaster(username string) bool {
	if strings.HasPrefix(username, masterPrefix) {
		return true
	}

	return false
}

func isCashOrder(orderType string) bool {
	if orderType == orderticket.Cash {
		return true
	}

	return false
}
