package main

import (
	"io/fs"
	"io/ioutil"
	"os"
	"path"

	"github.com/fatih/color"
)

type HTMLCLogger struct {
	Loader HTMLCLoader
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func (HTMLCLogger) PrintArgs() {
	argv := os.Args
	argc := len(argv)
	i := 1
	for i < argc {
		color.Blue("%s\n", argv[i])
		i++
	}
}

func (HTMLCLogger) PrintScopes(root string, files []fs.FileInfo) {
	for _, file := range files {
		if !file.IsDir() {
			filePath := path.Join(root, file.Name())
			content, err := ioutil.ReadFile(filePath)
			check(err)
			color.Green("%s\n", filePath)
			color.Yellow("Has Scope: ")
			println(HasScope(string(content)))
		}
	}
}

func (HTMLCLogger) PrintContextInfo(args ...string) {
	i := 0
	for i < len(args) {
		color.Blue("%s\n", args[i])
		i++
	}
}

func PrintArg(k string, v string) {
	color.Green("%s: %s", k, v)
}

func PrintLoader(loader HTMLCLoader) {
	//todo
	color.Green("%s: \n", "Resolved Chunk Path:")
	println(path.Join(loader.ProcessPath, loader.Config.BasePath, loader.Config.ChunkPath))
}
