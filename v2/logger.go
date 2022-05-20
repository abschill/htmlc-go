package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func printArgs() {
	argv := os.Args
	argc := len(argv)
	i := 1
	for i < argc {
		println(argv[i])
		i++
	}
}

func printScopes(root string, files []fs.FileInfo) {
	for _, file := range files {
		if !file.IsDir() {
			filePath := path.Join(root, file.Name())
			content, err := ioutil.ReadFile(filePath)
			check(err)
			println(filePath)
			println("Has Scope: ", hasScope(string(content)))
		}
	}
}

func printVersionInfo() {
	runtimeVersion := Version{
		0, 1, 1,
	}
	fmt.Printf("HTMLC Compiler Version: %d.%d.%d\n", runtimeVersion.Major, runtimeVersion.Minor, runtimeVersion.Patch)
	println(getAllowedExtensions())
}

func printContextInfo(args ...string) {
	i := 0
	for i < len(args) {
		println(args[i])
		i++
	}
}
