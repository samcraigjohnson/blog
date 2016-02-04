package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"regexp"
)

type Post struct {
	Content string
	FirstP  string
	Link    string
	Title   string
	Date    string
}

// Create a new post struct from an html file
// stored at fileLocation
func NewPost(fileLocation string) *Post {
	data, _ := ioutil.ReadFile(fileLocation)
	location := "/posts/" + fileLocation
	link := fmt.Sprintf("<a href=\"%s\">%s</a>", location, "read more")
	html := string(data)
	return &Post{
		Content: html,
		Link:    link,
		Title:   FindTag("h1", html),
		Date:    FindTag("h3", html),
		FirstP:  FindTag("p", html),
	}
}

// Create the shortened HTML for the homepage
func (p *Post) ToIndexHTML() template.HTML {
	html := p.Title + p.Date + p.FirstP + p.Link
	return template.HTML(html)
}

// Create full length HTML for post page
func (p *Post) ToHTML() template.HTML {
	return template.HTML(p.Content)
}

// Find the first HTML tag matching `tag` in an HTML string
func FindTag(tag string, html string) string {
	regex := fmt.Sprintf("<%s( class=.*)?( id=.*)?>.*</%s>", tag, tag)
	log.Printf("Regex: %v", regex)
	r, _ := regexp.Compile(regex)
	return r.FindString(html)
}
