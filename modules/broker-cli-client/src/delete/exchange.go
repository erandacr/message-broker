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

package delete

import (
	"github.com/spf13/cobra"
	"strconv"
	"utils"
)

var DeleteExchangeCmd = &cobra.Command{
	Use:     "exchange",
	Short:   "Delete an exchange from MB",
	Long:    "Delete an exchange from MB",
	Example: utils.DeleteExchangeCmdExamples,
	Run:     deleteExchangeCommand,
}

var unused bool

func init() {
	DeleteExchangeCmd.Flags().BoolVarP(&unused, "unused", "u", false, "Delete only if the exchange is unused (default: false)")
}

func deleteExchangeCommand(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		utils.PrintUsageErrorAndExit("insufficient arguments", "delete exchange")
	}

	utils.Logln("delete exchange command executes\n")

	url := utils.BROKER_URL_CONTEXT + utils.EXCHANGES + utils.FILE_SEPERATOR + args[0] + "?ifUnused=" + strconv.FormatBool(unused)
	deleteExchange(url)
}

func deleteExchange(url string) {
	client := utils.CreateHttpsClient(true)
	request := utils.CreateHttpRequest("DELETE", url, nil)
	response, err := utils.DoHttpAction(request, client)
	if err != nil {
		utils.HandleErrorAndExit("Error connecting to the broker", err)
	}

	//ignore the logging error
	_ = utils.LogResponseMessage(*response)
}
