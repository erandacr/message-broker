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

package config

import (
	"fmt"
	"github.com/renstrom/dedent"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"
	"syscall"
	"utils"
)

var (
	initLongDesc = dedent.Dedent(`Initialize MB Admin Client with Connection Details and User Credentials.
	`)
)

// initCmd represents the init command
var InitCmd = &cobra.Command{
	Use:     "init",
	Short:   "Initialize MB admin client",
	Long:    initLongDesc,
	Example: utils.InitCmdExamples,
	Run:     initCommand,
}

var hostname string
var port int
var username string
var password string

func init() {
	InitCmd.Flags().StringVarP(&hostname, "hostname", "H", "localhost", "Specify MB REST API hostname")
	InitCmd.Flags().IntVarP(&port, "port", "P", 9000, "Specify MB REST API port")
	InitCmd.Flags().StringVarP(&username, "username", "u", "admin", "Specify your username")
	InitCmd.Flags().StringVarP(&password, "password", "p", "", "Specify your password")
	InitCmd.Flags().SortFlags = false
	InitCmd.PersistentFlags().SortFlags = false
}

func initCommand(cmd *cobra.Command, args []string) {
	if len(password) == 0 {
		password = getPassword(username)
	}

	utils.Logf("%v init command is executed with hostname: %v, port: %d, username: %v\n", utils.ROOT_CMD, hostname, port, username)
	utils.GenerateConfigurationFile(hostname, port, username, password)
}

// prompt for password
func getPassword(username string) string {
	var password []byte
	fmt.Fprintf(os.Stderr, "Password for '%v': ", strings.TrimSpace(username))
	password, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		utils.HandleErrorAndExit("Unable to read your user password", err)
	}

	fmt.Fprintln(os.Stderr)
	return string(password)
}
