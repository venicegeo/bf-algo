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

package bf-algo

import (
    "fmt"
    "encoding/json"
)

type jsonHolder struct {
    algPayload []byte
    status int
}

var dumbDB []jsonHolder
var nextIndex = 0

// takes input in Json
// gives output as GeoJson
func ProcessEdgeLine(jsonAOI []byte) int {

    type aoiStruct struct {
        BoundBox [4]float64 // {minX, minY, maxX, maxY}
        ImageLink string // URL of image to be examined
    }

    var dataAOI aoiStruct  

    err:= json.Unmarshal(jsonAOI, &dataAOI)
    if err != nil {
        fmt.Println("error:", err)
    }

    dataLoad := []byte ( fmt.Sprintf(
        "{ \"type\": \"Feature\", \"geometry\": { \"type\": \"LineString\", \"coordinates\": [ [%f, %f], [%f, %f] ] }, \"properties\": { \"algorithm\": \"dummy\" } }", 
        dataAOI.BoundBox[0],
        dataAOI.BoundBox[1],
        dataAOI.BoundBox[2],
        dataAOI.BoundBox[3] ) )

    dumbDB = append (dumbDB, jsonHolder{dataLoad, 3}) 
    nextIndex++
    return nextIndex
}

func GetProcStatus(procId int) string {
    if procId >= nextIndex {
        return "Error: Process not initiated."
    }
    if dumbDB[procId].status == 0 {
        return "done"
    } else {
        dumbDB[procId].status--
        return "processing"
    }
}

func GetResult(resId int) string {
    if resId >= nextIndex {
        return "Error: Process not initiated."
    }
    if dumbDB[resId].status == 0 {
        return string(dumbDB[0].algPayload)
    } else {
        return "Error: Process incomplete."
    }
}

