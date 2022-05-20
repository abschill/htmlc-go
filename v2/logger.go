package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"path"
)

func check(
	err error,
) {
	if err != nil {
		panic(err)
	}
}

func printScopes(
	staticPath string,
	files []fs.FileInfo,
) {
	for _, file := range files {
		if !file.IsDir() {
			filePath := path.Join(staticPath, file.Name())
			content, err := ioutil.ReadFile(filePath)
			check(err)
			println(filePath)
			println("Has Scope: ", hasScope(string(content)))
		}
	}
}

func printVersionInfo() {
	runtimeVersion := version{
		0, 1, 1,
	}
	fmt.Printf("HTMLC Compiler Version: %d.%d.%d\n", runtimeVersion.Major, runtimeVersion.Minor, runtimeVersion.Patch)
	println(getAllowedExtensions())
}

func printContextInfo(
	args ...string,
) {
	i := 0
	for i < len(args) {
		println(args[i])
		i++
	}
}
