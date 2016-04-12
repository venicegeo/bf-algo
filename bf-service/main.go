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
	//fmt.Fprintf(w, `{ "dataType": { "type": "text", "content": "%s", "mimeType": "text/plain" }, "metadata": {} }`, string(outs) )
	fmt.Fprintf(w, string(outs))
  }
}

func main() {

        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                switch r.URL.Query().Get("cheese"){
			case "ls": {
				cmdLs := exec.Command("ls", "-alh")
				outLs, errLs := cmdLs.Output()
				printError(errLs, w)
				printOutput(outLs, w)
			}
			case "prep": {
				cmd0 := exec.Command("./prep.sh")
                                out0, err0 := cmd0.CombinedOutput()
                                printError(err0, w)
                                printOutput(out0, w)
			}
			case "case1": {
				cmd0 := exec.Command("curl", "-L", "-o", "LC80090472014280LGN00_B3.TIF", "https://www.dropbox.com/s/9e8kk0af5lz5odk/LC80090472014280LGN00_B3.TIF?dl=1")
				cmd0.Run()
				cmd1 := exec.Command("curl", "-L", "-o", "LC80090472014280LGN00_B6.TIF", "https://www.dropbox.com/s/nghh5uo711dvhfz/LC80090472014280LGN00_B6.TIF?dl=1")
				cmd1.Run()

				cmd2 := exec.Command("ossim-cli", "shoreline", "--image", "LC80090472014280LGN00_B3.TIF,", "LC80090472014280LGN00_B6.TIF", "--projection", "geo-scaled", "--threshold", "0.5", "--tolerance", "0", "product.json")
				cmd2.Run()

				cmd3 := exec.Command("cat", "product.json")
				out3, err3 := cmd3.CombinedOutput()
				printError(err3, w)
				printOutput(out3, w)

				cmd4 := exec.Command("rm", "LC80090472014280LGN00_B3.TIF", "LC80090472014280LGN00_B6.TIF")
				cmd4.Run()
			}
			case "case2": {

                                cmd1 := exec.Command("curl", "-L", "-o", "LC80150442014002LGN00_B6.TIF", "https://www.dropbox.com/s/a7scvz5jdaytc39/LC80150442014002LGN00_B6.TIF?dl=1")
                                cmd1.Run()

				cmd2 := exec.Command("ossim-cli", "shoreline", "--image", "LC80150442014002LGN00_B3.TIF,", "LC80150442014002LGN00_B6.TIF", "--projection", "geo-scaled", "--threshold", "0.5", "--tolerance", "0", "product.json")
                                cmd2.Run()

                                //cmd3 := exec.Command("cat", "product.json")
                                //out3, err3 := cmd3.CombinedOutput()
                                //printError(err3, w)
                                //printOutput(out3, w)

				cmd4 := exec.Command("rm", "LC80150442014002LGN00_B3.TIF", "LC80150442014002LGN00_B6.TIF")
                                cmd4.Run()
			}
			case "case3": {
				cmd0 := exec.Command("curl", "-L", "-o", "LC80340432016061LGN00_B3.TIF", "https://www.dropbox.com/s/ejoydxx6zembuxe/LC80340432016061LGN00_B3.TIF?dl=1")
                                cmd0.Run()
                                cmd1 := exec.Command("curl", "-L", "-o", "LC80340432016061LGN00_B6.TIF", "https://www.dropbox.com/s/laewldtba1j73f4/LC80340432016061LGN00_B6.TIF?dl=1")
                                cmd1.Run()

				cmd2 := exec.Command("ossim-cli", "shoreline", "--image", "LC80340432016061LGN00_B3.TIF,", "LC80340432016061LGN00_B6.TIF", "--projection", "geo-scaled", "--threshold", "0.5", "--tolerance", "0", "product.json")
				cmd2.Run()

				cmd3 := exec.Command("cat", "product.json")
                                out3, err3 := cmd3.CombinedOutput()
                                printError(err3, w)
                                printOutput(out3, w)

				cmd4 := exec.Command("rm", "LC80340432016061LGN00_B3.TIF", "LC80340432016061LGN00_B6.TIF")
                                cmd4.Run()
			}
			case "case4": {
				cmd0 := exec.Command("curl", "-L", "-o", "LC81190532015078LGN00_B3.TIF", "https://www.dropbox.com/s/6vcjay8zrp2l0zo/LC81190532015078LGN00_B3.TIF?dl=1")
                                cmd0.Run()
                                cmd1 := exec.Command("curl", "-L", "-o", "LC81190532015078LGN00_B6.TIF", "https://www.dropbox.com/s/tx8ganamkiw3zm3/LC81190532015078LGN00_B6.TIF?dl=1")
                                cmd1.Run()

				cmd2 := exec.Command("ossim-cli", "shoreline", "--image", "LC81190532015078LGN00_B3.TIF,", "LC81190532015078LGN00_B6.TIF", "--projection", "geo-scaled", "--threshold", "0.5", "--tolerance", "0", "product.json")
				cmd2.Run()

                                cmd3 := exec.Command("cat", "product.json")
                                out3, err3 := cmd3.CombinedOutput()
                                printError(err3, w)
                                printOutput(out3, w)

				cmd4 := exec.Command("rm", "LC81190532015078LGN00_B3.TIF", "LC81190532015078LGN00_B6.TIF")
                                cmd4.Run()
			}
			case "case5": {
				cmd0 := exec.Command("curl", "-L", "-o", "LC81600422014314LGN00_B3.TIF", "https://www.dropbox.com/s/h2rlbovzfo5h1ki/LC81600422014314LGN00_B3.TIF?dl=1")
                                cmd0.Run()
                                cmd1 := exec.Command("curl", "-L", "-o", "LC81600422014314LGN00_B6.TIF", "https://www.dropbox.com/s/a90cmi2ommp1z9w/LC81600422014314LGN00_B6.TIF?dl=1")
                                cmd1.Run()

				cmd2 := exec.Command("ossim-cli", "shoreline", "--image", "LC81600422014314LGN00_B3.TIF,", "LC81600422014314LGN00_B6.TIF", "--projection", "geo-scaled", "--threshold", "0.5", "--tolerance", "0", "product.json")
                                cmd2.Run()

                                cmd3 := exec.Command("cat", "product.json")
                                out3, err3 := cmd3.CombinedOutput()
                                printError(err3, w)
                                printOutput(out3, w)

				cmd4 := exec.Command("rm", "LC81600422014314LGN00_B3.TIF", "LC81600422014314LGN00_B6.TIF")
                                cmd4.Run()
			}
			case "case6": {
				cmd0 := exec.Command("curl", "-L", "-o", "LC82010352014217LGN00_B3.TIF", "https://www.dropbox.com/s/bos9qa9w6fvo7sq/LC82010352014217LGN00_B3.TIF?dl=1")
                                cmd0.Run()
                                cmd1 := exec.Command("curl", "-L", "-o", "LC82010352014217LGN00_B6.TIF", "https://www.dropbox.com/s/31um105g38t1c5u/LC82010352014217LGN00_B6.TIF?dl=1")
                                cmd1.Run()

				cmd2 := exec.Command("ossim-cli", "shoreline", "--image", "LC82010352014217LGN00_B3.TIF,", "LC82010352014217LGN00_B6.TIF", "--projection", "geo-scaled", "--threshold", "0.5", "--tolerance", "0", "product.json")
                                cmd2.Run()

                                cmd3 := exec.Command("cat", "product.json")
                                out3, err3 := cmd3.CombinedOutput()
                                printError(err3, w)
                                printOutput(out3, w)

				cmd4 := exec.Command("rm", "LC82010352014217LGN00_B3.TIF", "LC82010352014217LGN00_B6.TIF")
                                cmd4.Run()
			}
			default: {
				//cmd0 := exec.Command("ossim-cli", "shoreline", "--image", "garden_b3.tif,", "garden_b6.tif", "--projection", "geo-scaled", "--threshold", "0.5", "--tolerance", "0", "product.json")
				//cmd0.Run()

                                //cmd1 := exec.Command("cat", "product.json")
                                //out1, err1 := cmd1.CombinedOutput()
                                //printError(err1, w)
                                //printOutput(out1, w)
				fmt.Fprintf(w, "Just so we have something...")
				fmt.Fprintf(w, string(r.URL.RawQuery))
			}
                }
        })

        log.Fatal(http.ListenAndServe(":8080", nil))
}
