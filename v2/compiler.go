package main

import (
	"io/ioutil"
	"os"
)

func main() {
	cwd, err := os.Getwd()
	check(err)

	// this is our json file with the structed type
	configJSON := findConfig(cwd)
	staticPath := JoinPaths(cwd, "static/basic")
	files, err := ioutil.ReadDir(staticPath)
	check(err)
	printContextInfo(cwd, staticPath)
	printScopes(staticPath, files)

	println(configJSON["config"].Root)
}
