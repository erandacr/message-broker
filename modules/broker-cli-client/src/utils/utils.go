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
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

var MBVersion string

// create generic httpRequest
func CreateHttpRequest(method string, context string, body io.Reader) (req *http.Request) {
	url, username, password := getServiceUrlAndAuth()
	req, err := http.NewRequest(method, url+context, body)
	if err != nil {
		HandleErrorAndExit("Error reading config yaml", err)
	}
	req.SetBasicAuth(username, password)
	req.Header.Set(HeaderContentType, "application/json")
	req.Header.Set(HeaderAccept, "application/json")
	return
}

// create httpClient
func CreateHttpsClient(insecure bool) *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: insecure},
	}
	client := &http.Client{Transport: tr}
	return client
}

func DoHttpAction(request *http.Request, client *http.Client) (response *http.Response, err error) {
	LogRequest(request)
	response, err = client.Do(request)
	if response != nil {
		LogResponse(response)
	}
	return
}

// return full url and basic auth header value
func getServiceUrlAndAuth() (string, string, string) {
	config := readConfigurationFile()
	url := "https://" + config.Hostname + ":" + strconv.Itoa(config.Port)

	return url, config.Username, config.Password
}

// Read config file
func readConfigurationFile() (config Configuration) {
	config = Configuration{}
	dat, err := ioutil.ReadFile(CONFIG_FILE_NAME)
	if err != nil {
		HandleErrorAndExit("Error reading config yaml", err)
	}

	err = yaml.Unmarshal(dat, &config)
	if err != nil {
		HandleErrorAndExit("Error processing configuration data", err)
	}
	return
}

// Store config file
func GenerateConfigurationFile(hostname string, port int, username string, password string) {
	config := Configuration{hostname, port, username, password}
	// Convert struct to YAML.
	y, err := yaml.Marshal(config)
	if err != nil {
		HandleErrorAndExit("Error processing configuration data", err)
	}
	err = ioutil.WriteFile(CONFIG_FILE_NAME, y, 0644)
	if err != nil {
		HandleErrorAndExit("Error creating config yaml", err)
	}
}

// print error and exist
func PrintUsageErrorAndExit(msg, commandName string) {
	fmt.Fprintf(os.Stderr, ROOT_CMD+": %v\n", msg)
	fmt.Fprintf(os.Stderr, "Try '"+ROOT_CMD+" %v --help' for more information.\n", commandName)
	os.Exit(1)
}

// handle a given err by printing the error message to the standard error (os.Stderr)
// as well as make a call to os.Exit() with the exit status 1.
func HandleErrorAndExit(msg string, err error) {
	if err == nil {
		fmt.Fprintf(os.Stderr, ROOT_CMD+": %v\n", msg)
	} else {
		fmt.Fprintf(os.Stderr, ROOT_CMD+": %v reason: %v\n", msg, err.Error())
	}
	os.Exit(1)
}

// print json message on http responses
func LogResponseMessage(response http.Response) error {

	if response.Body == nil {
		return errors.New("empty response message")
	}

	var message Message
	decoder := json.NewDecoder(response.Body)
	err := decoder.Decode(&message)
	defer response.Body.Close()

	if err == nil {
		fmt.Fprintf(os.Stderr, ROOT_CMD+": %v", message.Message)
	}

	return err
}
