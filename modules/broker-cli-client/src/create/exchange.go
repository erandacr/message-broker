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
	"bytes"
	"encoding/json"
	"github.com/spf13/cobra"
	"utils"
)

var CreateExchangeCmd = &cobra.Command{
	Use:     "exchange",
	Short:   "Create an exchange in MB",
	Long:    "Create an exchange in MB with specific exchange properties",
	Example: utils.CreateExchangeCmdExamples,
	Run:     createExchangeCommand,
}

var exchangeType string
var durable bool

func init() {
	CreateExchangeCmd.Flags().StringVarP(&exchangeType, "type", "t", "direct", "Type of the exchange (default: direct)")
	CreateExchangeCmd.Flags().BoolVarP(&durable, "durable", "d", false, "Exchange durable or not (default: false)")
}

func createExchangeCommand(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		utils.PrintUsageErrorAndExit("insufficient arguments", "create exchange")
	}

	utils.Logln("create exchange command executes\n")

	exchange := utils.Exchange{Name: args[0], ExchangeType: exchangeType, Durable: durable}

	url := utils.BROKER_URL_CONTEXT + utils.EXCHANGES
	postExchange(url, exchange)
}

func postExchange(url string, exchange utils.Exchange) {
	// building payload
	byteBuf := new(bytes.Buffer)
	json.NewEncoder(byteBuf).Encode(exchange)

	client := utils.CreateHttpsClient(true)
	request := utils.CreateHttpRequest("POST", url, byteBuf)
	response, err := utils.DoHttpAction(request, client)
	if err != nil {
		utils.HandleErrorAndExit("Error connecting to the broker", err)
	}

	//ignore the logging error
	_ = utils.LogResponseMessage(*response)
}
