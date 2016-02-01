package main

import (
	"html/template"
	"fmt"
//	"regexp"
)

type Post struct {
	Content string
	FirstP string
	Link string
}

func NewPost(data string, fileLocation string) *Post {
	//	r := regexp.Compile("<p>[.]*</p>")
	link := fmt.Sprintf("<a href=\"%s\">%s</a>", fileLocation, "read more")
	return &Post{Content: data, Link: link}
}

func (p *Post) ToIndexHTML() template.HTML {
	html := p.Content + p.Link
	return template.HTML(html)
}

// Shorten the post for the hompage by only showing
// the first paragraph
func (p *Post) shorten() {
	
}
