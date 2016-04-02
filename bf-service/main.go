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
				cmd1 := exec.Command("ls", "-alh", "/etc/")
                                out1, err1 := cmd1.Output()

                                if err1 != nil {
                                        fmt.Fprintf(w, "Cheese", err1.Error())
                                        return
                                }
                                fmt.Fprintf(w, string(out1))



				cmd2 := exec.Command("cat", "/etc/*-release")
                                out2, err2 := cmd2.Output()

                                if err2 != nil {
                                        fmt.Fprintf(w, "Cheese", err2.Error())
                                        return
                                }
                                fmt.Fprintf(w, string(out2))
                        }
                }
        })

        log.Fatal(http.ListenAndServe(":8080", nil))
}
