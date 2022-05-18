package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
)

type version struct {
	Major uint8
	Minor uint8
	Patch uint8
}

type Config struct {
	root   string `json: "root"`
	chunks string `json: "chunks"`
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func defaultConfig() Config {
	return Config{
		"views", "chunks",
	}
}

func printVersionInfo() {
	runtimeVersion := version{
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

func findConfig(ctx string) {
	contextFiles, err := ioutil.ReadDir(ctx)
	check(err)
	for _, file := range contextFiles {
		if !file.IsDir() {
			fname := file.Name()
			if fname == "htmlc.json" {
				content, err := ioutil.ReadFile(path.Join(ctx, fname))
				check(err)
				println(content)
			}
		}
	}
}

func printScopes(staticPath string, files []fs.FileInfo) {
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
