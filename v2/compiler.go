package main

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
)

const htmlc_open_scope = "<!--@htmlc|"
const htmlc_close_scope = "|@htmlc-->"

func has_scope(content string) bool {
	return strings.Contains(content, htmlc_open_scope)
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	staticPath := path.Join(cwd, "static/basic")
	println(cwd)
	println(staticPath)

	files, err := ioutil.ReadDir(staticPath)
	for _, file := range files {
		if !file.IsDir() {
			filePath := path.Join(staticPath, file.Name())
			content, err := ioutil.ReadFile(filePath)
			if err != nil {
				panic(err)
			}
			println(filePath)
			println("Has Scope: ", has_scope(string(content)))
		}
	}
}
