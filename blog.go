package main

import (
	"fmt"
	"github.com/GeertJohan/go.rice"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	STATIC_BOX string = "static"
)

func indexHandler(box *rice.Box) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		posts, _ := posts()
		log.Printf("Posts: %v", posts)
		t := LoadTemplate("index", box)
		t.ExecuteTemplate(w, "base", posts)
	}
}

// Read all the HTML post fragments and return them as Posts
func posts() (*[]template.HTML, error) {
	dir := "./fragments"
	files, _ := ioutil.ReadDir(dir)

	posts := make([]template.HTML, len(files))
	for i, file := range files {
		data, _ := ioutil.ReadFile(dir + "/" + file.Name())
		p := NewPost(string(data), file.Name())
		posts[i] = p.ToIndexHTML()
	}
	return &posts, nil
}


// Load an html template from a rice box
func LoadTemplate(name string, box *rice.Box) *template.Template {
	base := templateString("base", box)
	content := templateString(name, box)
	log.Printf("Content of content: %v", content)
	log.Printf("Content of base: %v", base)

	t := template.New(name)
	t = parseString(t, base)
	t = parseString(t, content)
	log.Printf("Template: %v", t)
	return t
}

// Get a template string from a rice box
func templateString(name string, box *rice.Box) string {
	tName := fmt.Sprintf("%s.html", name)
	tString, err := box.String(tName)
	if err != nil {
		log.Printf("Error getting template string: %v", err)
	}
	return tString
}

func parseString(t *template.Template, s string) (*template.Template) {
	new, err := t.Parse(s)
	if err != nil {
		log.Printf("Error parsing string: %v", err)
	}
	return new
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
