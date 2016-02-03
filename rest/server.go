package main

import (
    "log"
    "net/http"
    "time"
)

func main() {
    server := &http.Server {
        Addr:           ":8000",
        Handler:        NewRouter(),
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   10 * time.Second,
        MaxHeaderBytes: 1 << 13, // 8K
    }

    log.Println("Listening...")
    log.Fatal(server.ListenAndServe())
}
