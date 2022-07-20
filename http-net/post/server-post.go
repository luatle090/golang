package main

// https://www.digitalocean.com/community/tutorials/how-to-make-http-requests-in-go

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

const serverPort = 3333

func main() {
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("server: %s /\n", r.Method)

			// get query string
			fmt.Printf("server: query id: %s\n", r.URL.Query().Get("id"))

			// get content type
			fmt.Printf("server: content-type: %s\n", r.Header.Get("content-type"))
			fmt.Printf("server: headers:\n")
			for headerName, headerValue := range r.Header {
				fmt.Printf("\t%s = %s\n", headerName, strings.Join(headerValue, ", "))
			}

			// get request body received from the client client
			reqBody, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Printf("server: could not read request body: %s\n", err)
			}
			fmt.Printf("server: request body: %s\n", reqBody)

			fmt.Fprintf(w, `{"message": "hello!"}`)
		})
		server := http.Server{
			Addr:    fmt.Sprintf(":%d", serverPort),
			Handler: mux,
		}
		if err := server.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				fmt.Printf("error running http server: %s\n", err)
			}
		}
	}()

	time.Sleep(100 * time.Millisecond)

	jsonBody := []byte(`{"client_message": "hello, server!"}`)
	bodyReader := bytes.NewReader(jsonBody)

	requestURL := fmt.Sprintf("http://localhost:%d?id=1234", serverPort)

	// use http.MethodPost
	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)

	// // Once the http.Request is created and configured
	// use the Do method of http.DefaultClient to send the request to the server
	// The http.DefaultClient value is Go’s default HTTP client, the same you’ve been using with http.Get
	// The Do method of the HTTP client returns the same values you received from the http.Get function so that you can handle the response in the same way.
	// Go’s default http.DefaultClient doesn’t specify a timeout, so if you make a request using that client,
	// it will wait until it receives a response
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("client: response body: %s\n", resBody)
}
