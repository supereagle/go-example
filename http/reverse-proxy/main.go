package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
)

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			writer.Write([]byte("Hello"))
		default:
			user := &User{}
			err := GetJsonPayload(request, user)
			if err != nil {
				ResponseWithError(writer, err)
				return
			}
			log.Printf("Request body after proxy: %+v", user)
			defer request.Body.Close()
		}
	})

	http.HandleFunc("/proxy", func(writer http.ResponseWriter, request *http.Request) {
		targetURL, err := url.Parse("http://127.0.0.1:8080")
		if err != nil {
			log.Fatal(err.Error())
			ResponseWithError(writer, err)
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(targetURL)
		// Change the request path.
		request.URL.Path = "/hello"
		switch request.Method {
		case http.MethodPost:
			// Get and print the request body.
			user := &User{}
			if err := GetJsonPayloadAndKeepState(request, user); err != nil {
				ResponseWithError(writer, err)
				return
			}
			log.Printf("Request body before proxy: %+v", user)
		case http.MethodPatch:
			// Get and update the request body.
			user := &User{}
			err := GetJsonPayload(request, user)
			if err != nil {
				ResponseWithError(writer, err)
				return
			}
			log.Printf("Request body before proxy: %+v", user)

			// Handle the request body before proxy.
			if user.Age > 30 {
				user.Young = false
			} else {
				user.Young = true
			}
			err = SetJsonPayload(request, user)
			if err != nil {
				ResponseWithError(writer, err)
				return
			}
		}

		// Proxy the request.
		proxy.ServeHTTP(writer, request)
	})

	http.ListenAndServe(":8080", nil)
}

// User represents the data in request body.
type User struct {
	Username string `json:"username"`
	Age      int    `json:"age"`
	Young    bool   `json:"young"`
}

// ResponseWithError responses the HTTP request with error. Please ensure that error is not nil when call this method.
func ResponseWithError(writer http.ResponseWriter, err error) {
	writer.WriteHeader(http.StatusInternalServerError)
	writer.Write([]byte(err.Error()))
}

// GetJsonPayload reads json payload from request and unmarshal it into entity. The request body will be empty.
func GetJsonPayload(request *http.Request, entity interface{}) error {
	content, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(content, entity)
	if err != nil {
		return err
	}

	return nil
}

// GetJsonPayloadAndKeepState reads json payload from request and unmarshal it into entity, and keep the request body.
func GetJsonPayloadAndKeepState(request *http.Request, entity interface{}) error {
	content, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return err
	}

	// Restore the io.ReadCloser to its original state.
	request.Body = ioutil.NopCloser(bytes.NewBuffer(content))

	err = json.Unmarshal(content, entity)
	if err != nil {
		return err
	}

	return nil
}

// SetJsonPayload marshals entity into json payload, and writes it into request body.
func SetJsonPayload(request *http.Request, entity interface{}) error {
	content, err := json.Marshal(entity)
	if err != nil {
		return err
	}

	request.Body = ioutil.NopCloser(bytes.NewBuffer(content))
	contentLength := len(content)
	request.Header.Set(http.CanonicalHeaderKey("Content-Length"), strconv.Itoa(contentLength))
	request.ContentLength = int64(contentLength)

	return nil
}
