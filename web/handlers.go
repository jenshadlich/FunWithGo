package main

import (
    "html/template"
    "log"
    "net/http"
)

type IndexPage struct {
}

func Index(w http.ResponseWriter, r *http.Request) {
    log.Println(r.RequestURI)
    t, _ := template.ParseFiles("templates/index.html")

    t.Execute(w, &IndexPage{})
}
