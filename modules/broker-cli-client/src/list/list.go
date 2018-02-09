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

package list

import (
	"github.com/spf13/cobra"
	"utils"
)

var ListCmd = &cobra.Command{
	Use:     "list",
	Short:   "List resources in MB",
	Long:    "List resources (exchanges, queues and bindings) in MB",
	Example: utils.ListCmdExamples,
	Run:     listCommand,
}

var all bool

func listCommand(cmd *cobra.Command, args []string) {
	utils.PrintUsageErrorAndExit("list command should be followed by MB resource type command.", "list")
}
