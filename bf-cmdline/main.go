// Copyright 2016, RadiantBlue Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"strconv"
	"github.com/spf13/cobra"
	"github.com/venicegeo/bf-algo"
)

var algoCmd = &cobra.Command{
	Use: "bf-cmdline",
	Long: " is a command-line interface to the GRiD database.",
}

var processCmd = &cobra.Command{
	Use:   "process",
	Short: "Initiate beachfront algorithm.  Example Format: bf-cmdline process \"{\\\"BoundBox\\\":[0,0,5,5],\\\"ImageLink\\\":\\\"dummy\\\"}\"",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(bfalgo.ProcessEdgeLine([]byte(args[0])))
	},
}

var statusCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of the GRiD CLI",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		idx, err := strconv.Atoi(args[0])
		if err != nil{
			fmt.Println("Error: Please specify an int procIndex.  If necessary, get help.")
		} else {
			fmt.Println(bfalgo.GetProcStatus(idx))
		}
	},
}

var resultCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of the GRiD CLI",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		idx, err := strconv.Atoi(args[0])
		if err != nil{
			fmt.Println("Error: Please specify an int resultIndex.  If necessary, get help.")
		} else {
			fmt.Println(bfalgo.GetResult(idx))
		}		
		
	},
}

var helpCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of the GRiD CLI",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		
		fmt.Println("grid v0.1 -- HEAD")
	},
}

func main() {
	algoCmd.AddCommand(processCmd)
	algoCmd.AddCommand(statusCmd)
	algoCmd.AddCommand(resultCmd)
	algoCmd.AddCommand(helpCmd)
	algoCmd.Execute()
}
