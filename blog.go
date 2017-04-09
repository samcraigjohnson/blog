package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/GeertJohan/go.rice"
)

const (
	STATIC_BOX string = "static"
	BLOG_DIR   string = "./fragments"
)

func indexHandler(box *rice.Box) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request for index /")
		posts, _ := posts()
		t := LoadTemplate("index", box)
		t.ExecuteTemplate(w, "base", posts)
	}
}

func postHandler(box *rice.Box) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Reqest for posts /posts")
		fileLocation := strings.Split(r.URL.Path, "/")[2]
		post := NewPost(postLocation(fileLocation))
		t := LoadTemplate("post", box)
		t.ExecuteTemplate(w, "base", post.ToHTML())
	}
}

// Read all the HTML post fragments and return them as Posts
func posts() (*[]template.HTML, error) {
	files, _ := ioutil.ReadDir(BLOG_DIR)

	// Create posts and sort by date
	posts := make([]Post, len(files))
	for i, file := range files {
		posts[i] = NewPost(postLocation(file.Name()))
	}
	sort.Sort(ByDate(posts))

	// Create HTML from posts
	postTemplates := make([]template.HTML, len(files))
	for i, p := range posts {
		postTemplates[i] = p.ToIndexHTML()
	}

	return &postTemplates, nil
}

func postLocation(name string) string {
	return fmt.Sprintf("%s/%s", BLOG_DIR, name)
}

// Load an html template from a rice box
func LoadTemplate(name string, box *rice.Box) *template.Template {
	base := templateString("base", box)
	content := templateString(name, box)

	t := template.New(name)
	t = parseString(t, base)
	t = parseString(t, content)
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

func parseString(t *template.Template, s string) *template.Template {
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

	port := os.Getenv("PORT")
	log.Printf("Listening on port %v...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
