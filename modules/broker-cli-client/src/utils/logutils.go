// Copyright (c) 2018, WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var loglnFunc = func(v ...interface{}) {}

var logfFunc = func(format string, a ...interface{}) {}

var logRequest = func(r *http.Request) {}

var logResponse = func(r *http.Response) {}

func Logln(a ...interface{}) {
	loglnFunc(a...)
}

func Logf(format string, a ...interface{}) {
	logfFunc(format, a...)
}

func LogRequest(r *http.Request) {
	logRequest(r)
}

func LogResponse(r *http.Response) {
	logResponse(r)
}

func EnableVerboseMode() {
	loglnFunc = func(a ...interface{}) { fmt.Fprintln(os.Stderr, a...) }
	logfFunc = func(format string, a ...interface{}) { fmt.Fprintf(os.Stderr, format, a...) }
	logRequest = func(request *http.Request) { fmt.Fprintln(os.Stderr, formatRequest(request)) }
	logResponse = func(response *http.Response) { fmt.Fprintln(os.Stderr, formatResponse(response)) }
}

// formatters

// formatRequest generates ascii representation of a request
func formatRequest(r *http.Request) string {
	// Create return string
	var request []string
	// Add the request string
	url := fmt.Sprintf("> %v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	// Add the host
	request = append(request, fmt.Sprintf("> Host: %v", r.Host))
	// Loop through headers
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("> %v: %v", name, h))
		}
	}

	// If this is a POST, add post data
	if r.Method == "POST" {
		r.ParseForm()
		request = append(request, "\n")
		request = append(request, r.Form.Encode())
	}
	// Return the request as a string
	return strings.Join(request, "\n") + "\n"
}

// formatResponse generates ascii representation of a response
func formatResponse(r *http.Response) string {
	// Create return string
	var response []string
	// Add the response string
	url := fmt.Sprintf("< %v %d %v", r.Proto, r.StatusCode, r.Status)
	response = append(response, url)
	// Loop through headers
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			response = append(response, fmt.Sprintf("< %v: %v", name, h))
		}
	}

	// Read the content and restore
	if r.Body != nil {
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			HandleErrorAndExit("Error reading body from http response in verbose mode", err)
		}
		// Restore the io.ReadCloser to its original state
		r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		// Use the content
		response = append(response, "\n")
		response = append(response, string(string(bodyBytes)))
	}

	// Return the response as a string
	return strings.Join(response, "\n") + "\n"
}
