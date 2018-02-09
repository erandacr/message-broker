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

package create

import (
	"github.com/spf13/cobra"
	"utils"
)

var CreateCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create resources in MB",
	Long:    "Create resources (exchanges, queues and bindings) in MB",
	Example: utils.CreateCmdExamples,
	Run:     createCommand,
}

func createCommand(cmd *cobra.Command, args []string) {
	utils.PrintUsageErrorAndExit("create command should be followed by MB resource type command.", "create")
}
