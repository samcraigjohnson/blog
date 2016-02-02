package main

import (
	"html/template"
	"fmt"
	"io/ioutil"
//	"regexp"
)

type Post struct {
	Content string
	FirstP string
	Link string
}

func NewPost(fileLocation string) *Post {
	//	r := regexp.Compile("<p>[.]*</p>")
	data, _ := ioutil.ReadFile(fileLocation)
	link := fmt.Sprintf("<a href=\"%s\">%s</a>", fileLocation, "read more")
	return &Post{Content: string(data), Link: link}
}

func (p *Post) ToIndexHTML() template.HTML {
	html := p.Content + p.Link
	return template.HTML(html)
}

// Shorten the post for the homepage by only showing
// the first paragraph
func (p *Post) shorten() {
	
}
