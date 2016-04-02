package main

import (
        "fmt"
        "log"
        "net/http"
        "os/exec"
	"os"
)

func main() {

        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                switch r.URL.Path{
                        case "/": {
				cmd := exec.Command("which", "ossim-info")
                                out, err := cmd.Output()

                                if err != nil {
                                        fmt.Fprintf(w, "Cheese", err.Error())
                                        return
                                }
                                fmt.Fprintf(w, string(out))



				cmd := exec.Command("cat", "/etc/*-release")
                                out, err := cmd.Output()

                                if err != nil {
                                        fmt.Fprintf(w, "Cheese", err.Error())
                                        return
                                }
                                fmt.Fprintf(w, string(out))
                        }
                }
        })

        log.Fatal(http.ListenAndServe(":8080", nil))
}
