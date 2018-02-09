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

package main

import (
	"config"
	"create"
	"delete"
	"github.com/renstrom/dedent"
	"github.com/spf13/cobra"
	"list"
	"os"
	"time"
	"utils"
)

var Verbose bool = false

var (
	rootLongDesc = dedent.Dedent(`
	WSO2 MB Admin Service Client.
	`)
)

var RootCmd = &cobra.Command{
	Use:   "mb",
	Short: "WSO2 MB Admin Service Client.",
	Long:  "WSO2 MB Admin Service Client.",
}

func init() {
	cobra.OnInitialize(initConfig)
	cobra.EnableCommandSorting = false
	RootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Enable verbose mode")
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// config commands
	RootCmd.AddCommand(config.VersionCmd)
	RootCmd.AddCommand(config.InitCmd)

	//list commands
	RootCmd.AddCommand(list.ListCmd)
	list.ListCmd.AddCommand(list.ListExchangeCmd)
	list.ListCmd.AddCommand(list.ListQueueCmd)
	list.ListCmd.AddCommand(list.ListBindingCmd)

	//create commands
	RootCmd.AddCommand(create.CreateCmd)
	create.CreateCmd.AddCommand(create.CreateExchangeCmd)

	//delete commands
	RootCmd.AddCommand(delete.DeleteCmd)
	delete.DeleteCmd.AddCommand(delete.DeleteExchangeCmd)

	// this will create a bash completion file
	//RootCmd.GenBashCompletionFile("mb_completion.sh")

	if err := RootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

func initConfig() {
	if Verbose {
		utils.EnableVerboseMode()
		t := time.Now()
		utils.Logf("Executing mb admin on : %v\n\n", t.Format(time.RFC1123))
	}
}
