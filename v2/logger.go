package main

import (
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

func (HTMLCDebugger) PrintArgs() {
	argv := os.Args
	argc := len(argv)
	i := 1
	for i < argc {
		println(argv[i])
		i++
	}
}

func (HTMLCDebugger) PrintScopes(root string, files []fs.FileInfo) {
	for _, file := range files {
		if !file.IsDir() {
			filePath := path.Join(root, file.Name())
			content, err := ioutil.ReadFile(filePath)
			check(err)
			println(filePath)
			println("Has Scope: ", HasScope(string(content)))
		}
	}
}

func (HTMLCDebugger) PrintContextInfo(args ...string) {
	i := 0
	for i < len(args) {
		println(args[i])
		i++
	}
}
