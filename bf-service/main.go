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
                                out, err := exec.Command("./ossim-info").Output()
                                if err != nil {
                                        log.Fatal(err)
                                }
                                fmt.Fprintf(w, "Cheese %s\n", out)
                        }
                }
        })

        log.Fatal(http.ListenAndServe(":8080", nil))
}
