package main

import (
        "fmt"
        "log"
        "net/http"
	"os"
        //"os/exec"
)

func main() {

        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                switch r.URL.Path{
                        case "/": {
				fmt.Fprintf(w, os.Getenv("HOME"))
                                //out, err := exec.Command("./ossim-info").CombinedOutput()
                                //if err != nil {
                                //        log.Fatal(err)
                                //}
                                //fmt.Fprintf(w, "Cheese %s\n", out)
                        }
                }
        })

        log.Fatal(http.ListenAndServe(":8080", nil))
}
