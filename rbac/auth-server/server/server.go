package server

import (
	"net/http"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/supereagle/go-example/rbac/auth-server/auth"
	"github.com/supereagle/go-example/rbac/auth-server/config"
)

func Run(cfg *config.Config) error {
	server := &Server{
		router: mux.NewRouter(),
	}

	server.RegisterRoutes()

	log.Infof("Start the server to listen on: %d", cfg.Port)
	return http.ListenAndServe(":"+strconv.Itoa(cfg.Port), server.router)
}

type Server struct {
	router *mux.Router
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

}

func parseRequest(req *http.Request) *auth.AuthRequest {
	if username, password, ok := req.BasicAuth(); ok {
		ar := &auth.AuthRequest{
			Username: username,
			Password: password,
		}

		return ar
	}

	return nil
}
