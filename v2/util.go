package main

import (
	"io/ioutil"
	"path"
)

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

func JoinPaths(base string, child string) string {
	return path.Join(base, child)
}

func DefaultConfig() Config {
	return Config{
		"views", "chunks",
	}
}
