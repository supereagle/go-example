package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello"))
	})

	http.HandleFunc("/proxy", func(writer http.ResponseWriter, request *http.Request) {
		targetURL, err := url.Parse("http://127.0.0.1:8080")
		if err != nil {
			log.Fatal(err.Error())
		}

		proxy := httputil.NewSingleHostReverseProxy(targetURL)
		// Change the request path.
		request.URL.Path = "/hello"
		proxy.ServeHTTP(writer, request)
	})

	http.ListenAndServe(":8080", nil)
}
