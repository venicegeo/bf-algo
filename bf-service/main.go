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
				cmd := exec.Command("ossim-info")
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
