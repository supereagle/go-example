package main

import (
	log "github.com/golang/glog"
	gitlabv3 "github.com/xanzy/go-gitlab"
	gitlabv4 "gopkg.in/xanzy/go-gitlab.v0"
)

func main() {
	server := "https://gitlab.com"
	token := "1234567890"

	v3Client, err := newV3Client(server, token)
	if err != nil {
		panic(err)
	}

	listProjectsByV3Client(v3Client)

	v4Client, err := newV4Client(server, token)
	if err != nil {
		panic(err)
	}

	listProjectsByV4Client(v4Client)
}

func newV3Client(server, token string) (*gitlabv3.Client, error) {
	client := gitlabv3.NewClient(nil, token)
	if err := client.SetBaseURL(server + "/api/v3"); err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return client, nil
}

func listProjectsByV3Client(client *gitlabv3.Client) {
	// List first 30 projects accessible by the authenticated user.
	opt := &gitlabv3.ListProjectsOptions{
		ListOptions: gitlabv3.ListOptions{
			PerPage: 30,
		},
	}
	_, _, err := client.Projects.ListProjects(opt)
	if err != nil {
		panic(err)
	}
}

func newV4Client(server, token string) (*gitlabv4.Client, error) {
	client := gitlabv4.NewClient(nil, token)
	if err := client.SetBaseURL(server + "/api/v4"); err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return client, nil
}

func listProjectsByV4Client(client *gitlabv4.Client) {
	// List first 30 projects accessible by the authenticated user.
	// Must set membership as true, otherwise will include public projects.
	trueVar := true
	opt := &gitlabv4.ListProjectsOptions{
		ListOptions: gitlabv4.ListOptions{
			PerPage: 30,
		},
		Membership: &trueVar,
	}
	_, _, err := client.Projects.ListProjects(opt)
	if err != nil {
		panic(err)
	}
}
