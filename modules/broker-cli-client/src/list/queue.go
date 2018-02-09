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
	"strconv"
	"strings"
	"utils"
)

var ListQueueCmd = &cobra.Command{
	Use:     "queue",
	Short:   "List queue(s) in MB",
	Long:    "List queue(s) in MB (all queues or selected one queue)",
	Example: "List queue examples",
	Run:     listQueueCommand,
}

func init() {
	ListQueueCmd.Flags().BoolVarP(&all, "all", "a", false, "Whether we need all the queues")
}

func listQueueCommand(cmd *cobra.Command, args []string) {
	url := utils.BROKER_URL_CONTEXT + utils.QUEUES

	utils.Logln("list queue command executes\n")

	var queues []utils.Queue

	if len(args) > 0 {
		url = url + utils.FILE_SEPERATOR + args[0]
		queues = getQueues(url, false)
	} else {
		queues = getQueues(url, true)
	}

	logTableQueues(queues)
}

// retrieve queue(s) information from the MB REST Service
func getQueues(url string, isAll bool) (queues []utils.Queue) {
	client := utils.CreateHttpsClient(true)
	request := utils.CreateHttpRequest("GET", url, nil)
	response, err := utils.DoHttpAction(request, client)
	if err != nil {
		utils.HandleErrorAndExit("Error connecting to the broker", err)
	}

	// do the parsing only if backend sends a valid response
	if strings.Contains(response.Status, "200") {
		decoder := json.NewDecoder(response.Body)

		if isAll {
			err2 := decoder.Decode(&queues)
			if err2 != nil {
				fmt.Println(err2)
			}
		} else {
			var queue utils.Queue
			err2 := decoder.Decode(&queue)
			if err2 != nil {
				fmt.Println(err2)
			}
			queues = []utils.Queue{queue}
		}
		defer response.Body.Close()
	}
	return
}

// log the queues in a table
func logTableQueues(queues []utils.Queue) {
	table := tablewriter.NewWriter(os.Stderr)
	table.SetHeader([]string{"name", "durable", "auto-delete", "consumer-count", "capacity", "size"})
	table.SetAutoFormatHeaders(false)

	for _, queue := range queues {
		table.Append([]string{queue.Name, strconv.FormatBool(queue.Durable),
			strconv.FormatBool(queue.AutoDelete), strconv.Itoa(queue.ConsumerCount),
			strconv.Itoa(queue.Capacity), strconv.Itoa(queue.Size)})
	}

	table.Render()
}
