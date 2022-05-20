package main

import (
	"io/ioutil"
	"os"
)

func main() {
	cwd, err := os.Getwd()
	printArgs()
	printVersionInfo()
	findConfig(cwd)
	check(err)
	staticPath := JoinPaths(cwd, "static/basic")
	files, err := ioutil.ReadDir(staticPath)
	printContextInfo(cwd, staticPath)
	printScopes(staticPath, files)
	check(err)
}
