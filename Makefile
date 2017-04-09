# Makefile for creating blog posts

all: clean convert install

build: blog.go
	go build

clean:
	go clean
	rm -f fragments/*.html
	rm -f build/blog

install: convert build
	go install

convert: 
	mkdir -p ./fragments
	for file in $(shell find org-files/ -name "*.org"); do \
		f=$$(basename $$file .org).html; \
		pandoc -s -o fragments/$$f $$file; \
	done

deploy: clean convert
	mkdir -p ./build
	GOOS=linux GOARCH=amd64 go build -o build/blog
	rice append --exec build/blog
	ssh ubuntu@ec2-54-153-40-4.us-west-1.compute.amazonaws.com rm blog
	scp config/blog.service ubuntu@ec2-54-153-40-4.us-west-1.compute.amazonaws.com:~/
	scp build/blog ubuntu@ec2-54-153-40-4.us-west-1.compute.amazonaws.com:~/
	rsync -r fragments ubuntu@ec2-54-153-40-4.us-west-1.compute.amazonaws.com:~/
	ssh ubuntu@ec2-54-153-40-4.us-west-1.compute.amazonaws.com sudo service blog restart
