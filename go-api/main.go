package main

import (
	"context"
	"io/ioutil"

	"github.com/google/go-github/github"
)

func main() {
	md, err := ioutil.ReadFile("sample.md")
	if err != nil {
		panic(err)
	}

	client := github.NewClient(nil)
	opt := &github.MarkdownOptions{Mode: "gfm", Context: "google/go-github"}

	html, _, err := client.Markdown(context.Background(), string(md), opt)
	if err != nil {
		panic(err)
	}

	ioutil.WriteFile("sample.html", []byte(html), 0644)
}
