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
				cmd1 := exec.Command("ls", "-Rlh")
                                out1, err1 := cmd1.Output()

                                if err1 != nil {
					log.Fatal(err1.Error())
                                        fmt.Fprintf(w, "Cheese", err1.Error())
                                        return
                                }
                                fmt.Fprintf(w, string(out1))
                        }
                }
        })

        log.Fatal(http.ListenAndServe(":8080", nil))
}
