package main

import (
	"fmt"
	"github.com/GeertJohan/go.rice"
	"html/template"
	"net/http"
)

const STATIC_BOX string = "static"

func LoadTemplate(name string, box *rice.Box) (*template.Template, error) {
	tName := fmt.Sprintf("%s.html", name)
	tString, err := box.String(tName)
	if err != nil {
		return nil, err
	}
	return template.New(name).Parse(tString)
}

type Data struct {
	Word string
}

func indexHandler(box *rice.Box) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, _ := LoadTemplate("index", box)
		t.Execute(w, &Data{Word: "helloworld"})
	}
}

func main() {
	// Web pages
	box := rice.MustFindBox("static")
	http.HandleFunc("/", indexHandler(box))

	// Static files
	staticFiles := http.StripPrefix("/static/", http.FileServer(box.HTTPBox()))
	http.Handle("/static/", staticFiles)

	http.ListenAndServe(":8080", nil)
}
