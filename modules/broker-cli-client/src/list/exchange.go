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
	"encoding/json"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"strings"
	"utils"
)

var ListExchangeCmd = &cobra.Command{
	Use:     "exchange",
	Short:   "List exchange(s) in MB",
	Long:    "List exchange(s) in MB (all exchanges or selected one exchange)",
	Example: utils.ListExchangeCmdExamples,
	Run:     listExchangeCommand,
}

func init() {
	ListExchangeCmd.Flags().BoolVarP(&all, "all", "a", false, "Retrieve info on all exchanges")
}

func listExchangeCommand(cmd *cobra.Command, args []string) {
	url := utils.BROKER_URL_CONTEXT + utils.EXCHANGES

	utils.Logln("list exchange command executes\n")

	var exchanges []utils.Exchange

	if len(args) > 0 && !all {
		url = url + utils.FILE_SEPERATOR + args[0]
		exchanges = getExchanges(url, false)
	} else {
		exchanges = getExchanges(url, true)
	}

	logTableExchanges(exchanges)
}

// invoke the MB Rest Service and get exchange(s) information
func getExchanges(url string, isArray bool) (exchanges []utils.Exchange) {
	client := utils.CreateHttpsClient(true)
	request := utils.CreateHttpRequest("GET", url, nil)
	response, err := utils.DoHttpAction(request, client)
	if err != nil {
		utils.HandleErrorAndExit("Error connecting to the broker", err)
	}

	// do the parsing only if backend sends a valid response
	if strings.Contains(response.Status, "200") {
		decoder := json.NewDecoder(response.Body)

		if isArray {
			err := decoder.Decode(&exchanges)
			if err != nil {
				utils.HandleErrorAndExit("Error processing exchange information", err)
			}
		} else {
			var exchange utils.Exchange
			err := decoder.Decode(&exchange)
			if err != nil {
				utils.HandleErrorAndExit("Error processing exchange information", err)
			}
			exchanges = []utils.Exchange{exchange}
		}
		defer response.Body.Close()
	}
	return
}

// log exchanges in a table
func logTableExchanges(exchanges []utils.Exchange) {
	table := tablewriter.NewWriter(os.Stderr)
	table.SetHeader([]string{"name", "type", "durable"})
	table.SetAutoFormatHeaders(false)

	for _, exchange := range exchanges {
		table.Append([]string{exchange.Name, exchange.ExchangeType, strconv.FormatBool(exchange.Durable)})
	}

	table.Render()
}
