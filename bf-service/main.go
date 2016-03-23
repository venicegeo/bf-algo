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
    "log"
    "net/http"
    "github.com/venicegeo/bf-algo"
)

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        switch r.URL.Path{
            case "/":
                fmt.Fprintf(w, "hello.")
            case "/dummyAlgo": {
                dummyAOI := []byte("{\"BoundBox\":[0,0,5,5],\"ImageLink\":\"dummy\"}")
                fmt.Fprintf(w, "%d", beachfront.ProcessEdgeLine(dummyAOI))
            }
            case "/checkStatus":
                fmt.Fprintf(w, beachfront.GetProcStatus(0))
            case "/getResult":
                fmt.Fprintf(w, beachfront.GetResult(0))
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


