package main

import (
	"io/ioutil"
	"os"
)

func main() {
	// content := []byte("hello world\n")
	// ioutil.WriteFile("/tmp/go-file", content, os.ModePerm)
	content := []byte(
		`hello world
 hello world
 hello world
 hello world
 `)
	ioutil.WriteFile("/tmp/go-file", content, os.ModePerm)
}
