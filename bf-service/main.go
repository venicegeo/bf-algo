package main

import (
        "fmt"
        "log"
        "net/http"
        "os/exec"
)

func main() {

        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                switch r.URL.Path{
                        case "/": {
				cmd1 := exec.Command("chmod", "a+x", "command.sh")
                                out1, err1 := cmd1.CombinedOutput()

                                if err1 != nil {
					log.Fatal(err1.Error())
                                	fmt.Fprintf(w, "Cheese", err1.Error())
                                	return
                                }
                                fmt.Fprintf(w, string(out1))

				cmd2 := exec.Command("./command.sh")
				out2, err2 := cmd2.CombinedOutput()
				
				if err2 != nil {
					log.Fatal(err2.Error())
					fmt.Fprintf(w, "Cheese", err2.Error())
					return
				}
				fmt.Fprintf(w, string(out2))
                        }
                }
        })

        log.Fatal(http.ListenAndServe(":8080", nil))
}
