package main

import (
	"io/ioutil"
	"os"
	"path"
)

func main() {
	cwd, err := os.Getwd()
	argv := os.Args
	argc := len(argv)
	printVersionInfo()
	findConfig(cwd)
	// which arg to offset from
	i := 1
	for i < argc {
		println(argv[i])
		i++
	}

	check(err)
	staticPath := path.Join(cwd, "static/basic")
	files, err := ioutil.ReadDir(staticPath)
	printContextInfo(cwd, staticPath)
	printScopes(staticPath, files)
	check(err)

}
