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
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"utils"
)

var ListBindingCmd = &cobra.Command{
	Use:     "binding",
	Short:   "List binding short",
	Long:    "List binding long",
	Example: "List binding examples",
	Run:     listBindingCommand,
}

func init() {
	ListBindingCmd.Flags().BoolVarP(&all, "pattern", "p", true, "List only specific binding pattern")
}

func listBindingCommand(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		utils.PrintUsageErrorAndExit("insufficient arguments", "list binding")
	}

	utils.Logln("list binding command executes\n")

	url := utils.BROKER_URL_CONTEXT + utils.EXCHANGES + utils.FILE_SEPERATOR + args[0] + utils.FILE_SEPERATOR + utils.BINDINGS
	logTableBindings(getBindings(url))
}

// retrieve bindings information from the MB REST Service
func getBindings(url string) (bindings []utils.BindingParent) {
	client := utils.CreateHttpsClient(true)
	request := utils.CreateHttpRequest("GET", url, nil)
	response, err := utils.DoHttpAction(request, client)
	if err != nil {
		utils.HandleErrorAndExit("Error connecting to the broker", err)
	}

	// do the parsing only if backend sends a valid response
	if strings.Contains(response.Status, "200") {
		decoder := json.NewDecoder(response.Body)
		err := decoder.Decode(&bindings)
		if err != nil {
			fmt.Println(err)
			defer response.Body.Close()
		}
	}
	return
}

// log the bindings in a table
func logTableBindings(bindings []utils.BindingParent) {
	table := tablewriter.NewWriter(os.Stderr)
	table.SetHeader([]string{"binding-pattern", "queue-name"})
	table.SetAutoFormatHeaders(false)

	for _, bindingPattern := range bindings {
		for _, binding := range bindingPattern.Bindings {
			table.Append([]string{bindingPattern.BindingPattern, binding.QueueName})
		}
	}

	table.Render()
}
