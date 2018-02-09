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

import "github.com/renstrom/dedent"

var InitCmdExamples = dedent.Dedent(`# Enter MB REST service connection details and user credentials as follows.
  mb init -H localhost -P 9000 -u admin -p admin123

# In cause you don't enter the password inline, you can enter it through the secured prompt.
  mb init -H localhost -P 9000 -u admin
  Password for 'admin':`)

var ListCmdExamples = dedent.Dedent(`# List all Exchanges in MB.
  mb list exchange
  or
  mb list exchange --all

# List information of the exchange named 'myExchange'.
  mb list exchange myExchange`)

var CreateCmdExamples = dedent.Dedent(`# Create a durable direct Exchange in MB.
  mb create exchange myExchange -t direct -d`)

var DeleteCmdExamples = dedent.Dedent(`# Delete the exchange named 'myExchange' from MB, only if its not used (no bindings).
  mb delete exchange myExchange -u`)

var ListExchangeCmdExamples = dedent.Dedent(`# List all Exchanges in MB.
  mb list exchange
  or
  mb list exchange --all

# List information of the exchange named 'myExchange'.
  mb list exchange myExchange`)

var CreateExchangeCmdExamples = dedent.Dedent(`# Create a durable direct Exchange in MB.
  mb create exchange myExchange -t direct -d`)

var DeleteExchangeCmdExamples = dedent.Dedent(`# Delete the exchange named 'myExchange' from MB, only if its not used (no bindings).
  mb delete exchange myExchange -u`)

var ListQueueCmdExamples = dedent.Dedent(`# List all Queues in MB.
  mb list queue
  or
  mb list queue --all

# List information of the queue named 'myQueue'.
  mb list queue myQueue`)

var CreateQueueCmdExamples = dedent.Dedent(`# Create a durable auto-delete Queue in MB.
  mb create queue myQueue -a -d`)

var DeleteQueueCmdExamples = dedent.Dedent(`# Delete the queue named 'myQueue' from MB, only if its not used (no bindings) and empty.
  mb delete queue myQueue -u -e`)
