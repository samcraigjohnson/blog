# Makefile for creating blog posts

all: clean convert install

build: blog.go
	go build

clean:
	go clean
	rm fragments/*.html

install: convert build
	go install

convert: 
	# Convert org-mode files into html fragments
	for file in $(shell find org-files/ -name "*.org"); do \
		f=$$(basename $$file .org).html; \
		pandoc -o fragments/$$f $$file; \
	done
