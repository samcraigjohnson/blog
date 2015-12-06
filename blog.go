package main

import (
	"fmt"
	"github.com/GeertJohan/go.rice"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

const STATIC_BOX string = "static"

func indexHandler(box *rice.Box) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		posts, _ := posts()
		log.Printf("Posts: %v", posts)
		t, _ := LoadTemplate("index", box)
		t.Execute(w, posts)
	}
}

// Read all the HTML post fragments and return them as Posts
func posts() (*[]template.HTML, error) {
	dir := "./fragments"
	files, _ := ioutil.ReadDir(dir)

	posts := make([]template.HTML, len(files))
	for i, file := range files {
		data, _ := ioutil.ReadFile(dir + "/" + file.Name())
		p := template.HTML(data)
		posts[i] = p
	}
	return &posts, nil
}

// Load an html template from a rice box
func LoadTemplate(name string, box *rice.Box) (*template.Template, error) {
	tName := fmt.Sprintf("%s.html", name)
	tString, err := box.String(tName)
	if err != nil {
		return nil, err
	}
	return template.New(name).Parse(tString)
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
