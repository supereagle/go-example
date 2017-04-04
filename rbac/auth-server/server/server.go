package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/supereagle/go-example/rbac/auth-server/auth"
	_ "github.com/supereagle/go-example/rbac/auth-server/auth/order"
	_ "github.com/supereagle/go-example/rbac/auth-server/auth/product"
	"github.com/supereagle/go-example/rbac/auth-server/config"
	"github.com/supereagle/go-example/rbac/auth-server/ticket"
	_ "github.com/supereagle/go-example/rbac/auth-server/ticket/order"
	_ "github.com/supereagle/go-example/rbac/auth-server/ticket/product"
	"github.com/supereagle/go-example/rbac/auth-server/token"
)

func Run(cfg *config.Config) error {
	server := &Server{
		router: mux.NewRouter(),
		config: cfg,
	}

	server.RegisterRoutes()

	log.Infof("Start the server to listen on: %d", cfg.Port)
	return http.ListenAndServe(":"+strconv.Itoa(cfg.Port), server.router)
}

type Server struct {
	router *mux.Router
	config *config.Config
}

func (server *Server) RegisterRoutes() {
	server.router.Path("/token").Methods("GET").HandlerFunc(server.doAuth)
}

func (server *Server) doAuth(resp http.ResponseWriter, req *http.Request) {
	ar := parseRequest(req)
	if ar == nil {
		log.Error("Fail to parse auth info from request.")
		return
	}

	authorizer, ok := auth.AuthPlugins[ar.Service]
	if !ok {
		log.Errorf("Authorizer %s does not exist.", ar.Service)
		return
	}

	err := authorizer.DoAuth(ar)
	if err != nil {
		log.Errorf("Auth failed for user %s as %s.", ar.Username, err.Error())
		return
	}

	tokenStr, err := token.GenToken(ar)
	if err != nil {
		log.Errorf("Fail to generate token for user %s as %s.", ar.Username, err.Error())
	}

	tk := make(map[string]interface{})
	tk["token"] = tokenStr
	tk["expires_in"] = server.config.Expiration
	tk["issued_at"] = time.Now().UTC().Format(time.RFC3339)
	result, _ := json.Marshal(tk)

	resp.Header().Set("Content-Type", "application/json")
	resp.Write(result)

	fmt.Printf("AR is %v\n", ar)
	fmt.Printf("tokenStr is %s\n", tokenStr)
}

func parseRequest(req *http.Request) *auth.AuthRequest {
	username, password, ok := req.BasicAuth()
	if !ok {
		log.Error("Fail to get user info from request's Authorization header")
		return nil
	}

	service := auth.Service(req.URL.Query().Get("service"))
	scopeStr := req.URL.Query().Get("scope")

	scope, err := parseTicket(service, scopeStr)
	if err != nil {
		log.Errorf("Fail to parse ticket %s as %s", scopeStr, err.Error())
		return nil
	}

	ar := &auth.AuthRequest{
		Username: username,
		Password: password,
		Service:  service,
		Scope:    scope,
	}

	return ar
}

func parseTicket(service auth.Service, scopeStr string) (*auth.Scope, error) {
	ticketer, ok := ticket.TicketPlugins[service]
	if !ok {
		return nil, fmt.Errorf("Ticketer %s does not exist.", service)
	}

	return ticketer.ParseTicket(scopeStr)
}
