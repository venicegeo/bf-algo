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
    fmt.Fprintf(w, string(outs))
  }
}

func main() {

        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                switch r.URL.Path{
			case "/ls": {
				cmdLs := exec.Command("ls", "-alh")
				outLs := cmdLs.Output()
				printOutput(outLs, w)
			}
			case "/case1": {
				cmd0 := exec.Command("ossim-cli", "shoreline", "--image", "LC80090472014280LGN00_B3.TIF,", "LC80090472014280LGN00_B6.TIF", "--projection", "geo-scaled", "--threshold", "0.5", "--tolerance", "0", "product.json")
				cmd0.Run()

				cmd1 := exec.Command("cat", "product.json")
				out1, err1 := cmd1.CombinedOutput()
				printError(err1, w)
				printOutput(out1, w)
			}
			case "/case2": {
				cmd0 := exec.Command("ossim-cli", "shoreline", "--image", "LC80150442014002LGN00_B3.TIF,", "LC80150442014002LGN00_B6.TIF", "--projection", "geo-scaled", "--threshold", "0.5", "--tolerance", "0", "product.json")
                                cmd0.Run()

                                cmd1 := exec.Command("cat", "product.json")
                                out1, err1 := cmd1.CombinedOutput()
                                printError(err1, w)
                                printOutput(out1, w)
			}
			case "/case3": {
				cmd0 := exec.Command("ossim-cli", "shoreline", "--image", "LC80340432016061LGN00_B3.TIF,", "LC80340432016061LGN00_B6.TIF", "--projection", "geo-scaled", "--threshold", "0.5", "--tolerance", "0", "product.json")
				cmd0.Run()

				cmd1 := exec.Command("cat", "product.json")
                                out1, err1 := cmd1.CombinedOutput()
                                printError(err1, w)
                                printOutput(out1, w)
			}
			case "/case4": {
				cmd0 := exec.Command("ossim-cli", "shoreline", "--image", "LC81190532015078LGN00_B3.TIF,", "LC81190532015078LGN00_B6.TIF", "--projection", "geo-scaled", "--threshold", "0.5", "--tolerance", "0", "product.json")
				cmd0.Run()

                                cmd1 := exec.Command("cat", "product.json")
                                out1, err1 := cmd1.CombinedOutput()
                                printError(err1, w)
                                printOutput(out1, w)
			}
			case "/case5": {
				cmd0 := exec.Command("ossim-cli", "shoreline", "--image", "LC81600422014314LGN00_B3.TIF,", "LC81600422014314LGN00_B6.TIF", "--projection", "geo-scaled", "--threshold", "0.5", "--tolerance", "0", "product.json")
                                cmd0.Run()

                                cmd1 := exec.Command("cat", "product.json")
                                out1, err1 := cmd1.CombinedOutput()
                                printError(err1, w)
                                printOutput(out1, w)
			}
			case "/case6": {
				cmd0 := exec.Command("ossim-cli", "shoreline", "--image", "LC82010352014217LGN00_B3.TIF,", "LC82010352014217LGN00_B6.TIF", "--projection", "geo-scaled", "--threshold", "0.5", "--tolerance", "0", "product.json")
                                cmd0.Run()

                                cmd1 := exec.Command("cat", "product.json")
                                out1, err1 := cmd1.CombinedOutput()
                                printError(err1, w)
                                printOutput(out1, w)
			}
			default: {
				cmd0 := exec.Command("ossim-cli", "shoreline", "--image", "LC82010352014217LGN00_B3.TIF,", "LC82010352014217LGN00_B6.TIF", "--projection", "geo-scaled", "--threshold", "0.5", "--tolerance", "0", "product.json")
                                cmd0.Run()

                                cmd1 := exec.Command("cat", "product.json")
                                out1, err1 := cmd1.CombinedOutput()
                                printError(err1, w)
                                printOutput(out1, w)
			}
                }
        })

        log.Fatal(http.ListenAndServe(":8080", nil))
}
