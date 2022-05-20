package main

import (
	"encoding/json"
	"io/ioutil"
	"path"
)

func findConfig(ctx string) ConfigFile {
	var res ConfigFile
	contextFiles, err := ioutil.ReadDir(ctx)
	check(err)
	for _, file := range contextFiles {
		if !file.IsDir() {
			fname := file.Name()
			if fname == "htmlc.json" {
				content, err := ioutil.ReadFile(path.Join(ctx, fname))
				check(err)

				json.Unmarshal([]byte(content), &res)

				check(err)
			}
		}
	}
	return res
}

func JoinPaths(base string, child string) string {
	return path.Join(base, child)
}

func DefaultConfig() Config {
	return Config{
		"views", "chunks",
	}
}
