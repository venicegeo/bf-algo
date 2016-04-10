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
                        case "/": {


				//cmd1 := exec.Command("ossim-cli", "shoreline", "--image", "garden_b3.tif,", "garden_b6.tif", "--projection", "geo-scaled", "--threshold", "0.5", "--tolerance", "0", "product.json")
				//printCommand(cmd1, w)
                                //out1, err1 := cmd1.CombinedOutput()
				//printError(err1, w)
				//printOutput(out1, w)
				//cmd1.Run()

				//cmd2 := exec.Command("cat", "product.json")
				//printCommand(cmd2, w)
				//out2, err2 := cmd2.CombinedOutput()
				//printError(err2, w)
				//printOutput(out2, w)

				cmd3 := exec.Command("ls", "-alh")
				printCommand(cmd3, w)
				out3, err3 := cmd3.CombinedOutput()
				printError(err3, w)
				printOutput(out3, w)

                        },
			case "prep10": {
				// LC80090472014280LGN00_B3.TIF
				cmd10 := exec.Command("curl", "-L", "-o", "LC80090472014280LGN00_B3.TIF", "https://www.dropbox.com/s/9e8kk0af5lz5odk/LC80090472014280LGN00_B3.TIF?dl=1")
				printCommand(cmd10, w)
				out10, err10 := cmd10.CombinedOutput()
				printError(err10, w)
				printOutput(cmd10, w)
				//cmd10.Run()
			},
			case "prep11": {
				// LC80090472014280LGN00_B6.TIF
                                cmd11 := exec.Command("curl", "-L", "-o", "LC80090472014280LGN00_B6.TIF", "https://www.dropbox.com/s/nghh5uo711dvhfz/LC80090472014280LGN00_B6.TIF?dl=1")
                                cmd11.Run()
			},
			case "prep12": {
				// LC80150442014002LGN00_B3.TIF
				cmd12 := exec.Command("curl", "-L", "-o", "LC80150442014002LGN00_B3.TIF", "https://www.dropbox.com/s/867cbaw0vmu7ssn/LC80150442014002LGN00_B3.TIF?dl=1")
                                cmd12.Run()
			},
			case "prep13": {
				// LC80150442014002LGN00_B6.TIF
                                cmd13 := exec.Command("curl", "-L", "-o", "LC80150442014002LGN00_B6.TIF", "https://www.dropbox.com/s/a7scvz5jdaytc39/LC80150442014002LGN00_B6.TIF?dl=1")
                                cmd13.Run()
			},
			case "prep14": {
				// LC80340432016061LGN00_B3.TIF
                                cmd14 := exec.Command("curl", "-L", "-o", "LC80340432016061LGN00_B3.TIF", "https://www.dropbox.com/s/ejoydxx6zembuxe/LC80340432016061LGN00_B3.TIF?dl=1")
                                cmd14.Run()
			},
			case "prep15": {
				// LC80340432016061LGN00_B6.TIF
                                cmd15 := exec.Command("curl", "-L", "-o", "LC80340432016061LGN00_B6.TIF", "https://www.dropbox.com/s/laewldtba1j73f4/LC80340432016061LGN00_B6.TIF?dl=1")
                                cmd15.Run()
			},
			case "prep16": {
				// LC81190532015078LGN00_B3.TIF
                                cmd16 := exec.Command("curl", "-L", "-o", "LC81190532015078LGN00_B3.TIF", "https://www.dropbox.com/s/6vcjay8zrp2l0zo/LC81190532015078LGN00_B3.TIF?dl=1")
                                cmd16.Run()
			},
			case "prep17": {
				// LC81190532015078LGN00_B6.TIF
                                cmd17 := exec.Command("curl", "-L", "-o", "LC81190532015078LGN00_B6.TIF", "https://www.dropbox.com/s/tx8ganamkiw3zm3/LC81190532015078LGN00_B6.TIF?dl=1")
                                cmd17.Run()
			},
			case "prep18": {
				// LC81600422014314LGN00_B3.TIF
                                cmd18 := exec.Command("curl", "-L", "-o", "LC81600422014314LGN00_B3.TIF", "https://www.dropbox.com/s/h2rlbovzfo5h1ki/LC81600422014314LGN00_B3.TIF?dl=1")
                                cmd18.Run()
			},
			case "prep19": {
				// LC81600422014314LGN00_B6.TIF
                                cmd19 := exec.Command("curl", "-L", "-o", "LC81600422014314LGN00_B6.TIF", "https://www.dropbox.com/s/a90cmi2ommp1z9w/LC81600422014314LGN00_B6.TIF?dl=1")
                                cmd19.Run()
			},
			case "prep20": {
				// LC82010352014217LGN00_B3.TIF?
                                cmd20 := exec.Command("curl", "-L", "-o", "LC82010352014217LGN00_B3.TIF", "https://www.dropbox.com/s/bos9qa9w6fvo7sq/LC82010352014217LGN00_B3.TIF?dl=1")
                                cmd20.Run()
			},
			case "prep21": {
				// LC82010352014217LGN00_B6.TIF
                                cmd21 := exec.Command("curl", "-L", "-o", "LC82010352014217LGN00_B6.TIF", "https://www.dropbox.com/s/31um105g38t1c5u/LC82010352014217LGN00_B6.TIF?dl=1")
                                cmd21.Run()
			}
                }
        })

        log.Fatal(http.ListenAndServe(":8080", nil))
}
