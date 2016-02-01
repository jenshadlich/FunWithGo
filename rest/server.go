package main

import (
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "time"
)

func main() {
    requestRouter := mux.NewRouter().StrictSlash(true)
    requestRouter.HandleFunc("/hello", Hello).Methods("GET")

    server := &http.Server{
        Addr:           ":8000",
        Handler:        requestRouter,
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   10 * time.Second,
        MaxHeaderBytes: 1 << 13, // 8K
    }

    log.Println("Listening...")
    log.Fatal(server.ListenAndServe())
}
