package main

import (
    "log"
    "net/http"
    "encoding/json"
)

type HelloResponse struct {
    Message string `json:"message"`
}

func Hello(w http.ResponseWriter, r *http.Request) {
    log.Println(r.RequestURI)

    response := &HelloResponse{Message:"Hello World!"};
    w.Header().Add("Content-Type", "application/json")

    if err := json.NewEncoder(w).Encode(response); err != nil {
        panic(err)
    }
}
