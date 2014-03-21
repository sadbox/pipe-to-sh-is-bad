package main

import (
    "net/http"
    "strings"
    "log"
    "fmt"
)

func main() {
    http.HandleFunc("/install.sh", func(w http.ResponseWriter, r *http.Request) { 
        log.Printf("Served %s to %s, user-agent is %s", r.URL, r.RemoteAddr, r.Header.Get("user-agent"))
        if strings.Contains(r.Header.Get("user-agent"), "curl") {
            fmt.Fprint(w, "{INSERT BAD THINGS HERE}\n")
        } else {
            fmt.Fprint(w, "Regular script for people to look at\n")
        }
    })
    http.ListenAndServe(":8080", nil)
}
