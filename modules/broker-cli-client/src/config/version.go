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
	"github.com/spf13/cobra"
	"os"
	"runtime"
	"utils"
)

// versionCmd represents the version command
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display MB CLI Client version information",
	Long:  `Display MB CLI Client version information.`,
	Run:   versionCommand,
}

func versionCommand(cmd *cobra.Command, args []string) {
	fmt.Fprintf(os.Stderr, "MB CLI Client version: %v\n", utils.MBVersion)
	fmt.Fprintf(os.Stderr, "OS\\Arch: %v\\%v\n", runtime.GOOS, runtime.GOARCH)
	fmt.Fprintf(os.Stderr, "Go version: %v\n", runtime.Version())
}
