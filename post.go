package main

import (
	"html/template"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
)

type Post struct {
	Content string
	FirstP string
	Link string
	Title string
	Date string
}

func NewPost(fileLocation string) *Post {
	//	r := regexp.Compile("<p>[.]*</p>")
	data, _ := ioutil.ReadFile(fileLocation)
	link := fmt.Sprintf("<a href=\"%s\">%s</a>", fileLocation, "read more")
	html := string(data)
	return &Post{
		Content: html,
		Link: link,
		Title: FindTag("h1", html),
		Date: FindTag("h3", html),
		FirstP: FindTag("p", html),
	}
}

func (p *Post) ToIndexHTML() template.HTML {
	html := p.Title + p.Date + p.FirstP + p.Link
	return template.HTML(html)
}

// Shorten the post for the homepage by only showing
// the first paragraph
func (p *Post) shorten() {
	
}


func FindTag(tag string, html string) string {
	regex := fmt.Sprintf("<%s( class=.*)?( id=.*)?>.*</%s>", tag, tag)
	log.Printf("Regex: %v", regex)
	r, _ := regexp.Compile(regex)
	return r.FindString(html)
}
