package main

import (
	"io/ioutil"
	"path"
)

func defaultConfig() Config {
	return Config{
		"views", "chunks",
	}
}

func findConfig(
	ctx string,
) {
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
