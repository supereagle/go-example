package main

import (
	"flag"
	"net/http"

	"github.com/bndr/gojenkins"
	log "github.com/golang/glog"
)

var (
	jenkinsServer   = "http://127.0.0.1:8080"
	jenkinsUsername = "admin"
	jenkinsPassword = "passw0rd"
	jenkinsNodeURL  = "http://127.0.0.1:8080/computer/"
)

func main() {
	flag.Set("logtostderr", "true")
	flag.Parse()

	j := gojenkins.CreateJenkins(nil, jenkinsServer, jenkinsUsername, jenkinsPassword)
	_, err := j.Init()
	if err != nil {
		log.Fatal(err.Error())
	}

	ns, err := j.GetAllNodes()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Create gojenkins issue: error is not nil when the auth is wrong.
	if len(ns) > 0 {
		log.Infof("First node: %s", ns)
	}

	testCustomizedClient()

	testProxyClient()
}

func testCustomizedClient() {
	cc := NewCustomizedClient(jenkinsUsername, jenkinsPassword)
	resp, err := cc.do("GET", jenkinsNodeURL, nil)
	if err != nil {
		log.Error(err.Error())
		return
	}

	if resp.StatusCode == http.StatusUnauthorized {
		log.Error("Customized client fails to auth")
		return
	}

	log.Info("Customized client passes the auth")
}

func testProxyClient() {
	c := NewProxyClient(jenkinsUsername, jenkinsPassword)

	req, err := http.NewRequest("GET", jenkinsNodeURL, nil)
	if err != nil {
		log.Error(err.Error())
		return
	}

	resp, err := c.Do(req)
	if err != nil {
		log.Error(err.Error())
		return
	}

	if resp.StatusCode == http.StatusUnauthorized {
		log.Error("Proxy client fails to auth")
		return
	}

	log.Info("Proxy client passes the auth")
}
