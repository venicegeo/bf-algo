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
	"bufio"
	"fmt"
	"log"
	"strconv"
	"net/http"
	"github.com/venicegeo/bf-algo"
	"os"
	"os/exec"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path{
			case "/": {
				//fmt.Fprintf(w, "hello.")
				cmdName := "ls"
				cmdArgs := []string{""}

				cmd := exec.Command(cmdName, cmdArgs...)
				cmdReader, err := cmd.StdoutPipe()
				if err != nil {
					fmt.Fprintf(w, "Error creating StdoutPipe for Cmd")
					os.Exit(1)
				}

				scanner := bufio.NewScanner(cmdReader)
				go func() {
					for scanner.Scan() {
						fmt.Fprintf(w, "ossim-info out \n")
						//fmt.Fprintf(w, "ossim-info out | %s\n", scanner.Text())
					}
				}()

				err = cmd.Start()
				if err != nil {
					fmt.Fprintf(w, "Error starting Cmd")
					os.Exit(1)
				}

				err = cmd.Wait()
				if err != nil {
					fmt.Fprintf(w, "Error waiting for Cmd")
					os.Exit(1)
				}
			}
			case "/dummyAlgo": {
				aoiString := r.URL.Query().Get("aoi")
				if aoiString == "" {
					fmt.Fprintf(w, "Error: Please specify an aoi in your query parameters\n")
					fmt.Fprintf(w, "Example: aoi={\"BoundBox\":[0,0,5,5],\"ImageLink\":\"dummy\"}\n")
				} else {
				//dummyAOI := []byte("{\"BoundBox\":[0,0,5,5],\"ImageLink\":\"dummy\"}")
					idx, err := bfalgo.ProcessEdgeLine([]byte(aoiString))
					if err != nil {
						fmt.Fprintf(w, err.Error())
					} else {
						fmt.Fprintf(w, "%d", idx)
					}
				}
			}
			case "/checkStatus":
				idxString := r.URL.Query().Get("procIndex")
				idx, err := strconv.Atoi(idxString)
				if err != nil{
					fmt.Fprintf(w, "Error: Please specify an int procIndex in your query parameters")
				} else {
					fmt.Fprintf(w, bfalgo.GetProcStatus(idx))
				}
			case "/getResult":
				idxString := r.URL.Query().Get("resultIndex")
				idx, err := strconv.Atoi(idxString)
				if err != nil{
					fmt.Fprintf(w, "Error: Please specify an int resultIndex in your query parameters")
				} else {
					fmt.Fprintf(w, bfalgo.GetResult(idx))
				}
			case "/help":
				help(w)
			default:
				other(w)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func other(w http.ResponseWriter) {
	fmt.Fprintf(w, "Command undefined.  Try help?\n")
}

func help(w http.ResponseWriter) {
	fmt.Fprintf(w, "We're sorry, help is not yet implemented\n")
}


