package main

import (
        "fmt"
        "log"
        "net/http"
	"strings"
        "os/exec"
)


func printCommand(cmd *exec.Cmd, w http.ResponseWriter) {
  fmt.Fprintf(w, "==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

func printError(err error, w http.ResponseWriter) {
  if err != nil {
    fmt.Fprintf(w, "==> Error: %s\n", err.Error())
  }
}

func printOutput(outs []byte, w http.ResponseWriter) {
  if len(outs) > 0 {
    fmt.Fprintf(w, "==> Output: %s\n", string(outs))
  }
}

func main() {

        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                switch r.URL.Path{
                        case "/": {


				cmd1 := exec.Command("ossim-cli", "shoreline", "--image", "garden_b3.tif,", "garden_b6.tif", "--projection", "geo-scaled", "--threshold", "0.5", "--tolerance", "0", "product.json")
				//printCommand(cmd1, w)
                                out1, err1 := cmd1.CombinedOutput()
				//printError(err1, w)
				//printOutput(out1, w)

				cmd2 := exec.Command("cat", "product.json")
				//printCommand(cmd2, w)
				out2, err2 := cmd2.CombinedOutput()
				printError(err2, w)
				printOutput(out2, w)

                        }
                }
        })

        log.Fatal(http.ListenAndServe(":8080", nil))
}
