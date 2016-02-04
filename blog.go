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
	BLOG_DIR string = "./fragments"
)

func indexHandler(box *rice.Box) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		posts, _ := posts()
		log.Printf("Posts: %v", posts)
		t := LoadTemplate("index", box)
		t.ExecuteTemplate(w, "base", posts)
	}
}

func postHandler(box *rice.Box) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Path: %v", r.URL.Path)
		post := NewPost(postLocation(r.URL.Path))
		t := LoadTemplate("post", box)
		t.ExecuteTemplate(w, "base", post.ToHTML())
	}
}

// Read all the HTML post fragments and return them as Posts
func posts() (*[]template.HTML, error) {
	files, _ := ioutil.ReadDir(BLOG_DIR)

	posts := make([]template.HTML, len(files))
	for i, file := range files {
		p := NewPost(postLocation(file.Name()))
		posts[i] = p.ToIndexHTML()
	}
	return &posts, nil
}

func postLocation(name string) string {
	return fmt.Sprintf("%s/%s", BLOG_DIR, name)
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
	http.HandleFunc("/posts/", postHandler(box))

	// Static files
	staticFiles := http.StripPrefix("/static/", http.FileServer(box.HTTPBox()))
	http.Handle("/static/", staticFiles)

	http.ListenAndServe(":8080", nil)
}
