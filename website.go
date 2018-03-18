package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
)

var addr = flag.String("addr", ":1718", "http service address") // Q=17, R=18

func main() {
	http.Handle("/", http.HandlerFunc(index))
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func index(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html ; charset=utf-8")
	wd, _ := os.Getwd()
	file := wd + "/templates/index.html"
	t, err := template.New("root").ParseFiles(file)
	if err != nil {
		log.Fatal("Template error :", err)
	}
	t.Execute(w, nil)
}
