package main

import (
	"net/http"
	"net/url"
	"strconv"
)

var logContent = `Started by user robin
Building in workspace /Users/robin/.jenkins/workspace/job2
[job2] $ /bin/sh -xe /var/folders/gc/7sjlxfsx4p1f0qlrtx_g3_j80000gn/T/jenkins7178629777156498390.sh
+ echo 'hello world'
hello world
Finished: SUCCESS
`

func main() {
	http.HandleFunc("/log", serveLog)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
}

func serveLog(w http.ResponseWriter, r *http.Request) {
	download, err := parseDownloadParam(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("Content-Type", "text/plain")
	if download {
		w.Header().Add("Content-Disposition", "attachment; filename=log.txt")
	}

	w.Write([]byte(logContent))
}

func parseDownloadParam(r *http.Request) (bool, error) {
	values, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		return false, err
	}

	downloadStr := values.Get("download")
	download := false
	if downloadStr != "" {
		var err error
		download, err = strconv.ParseBool(downloadStr)
		if err != nil {
			return false, err
		}
	}

	return download, nil
}
