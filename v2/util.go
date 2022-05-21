package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
)

func GetVersionInfo() string {
	runtimeVersion := HTMLCVersion{
		0, 1, 1,
	}
	return fmt.Sprintf("HTMLC Compiler Version: %d.%d.%d\n", runtimeVersion.Major, runtimeVersion.Minor, runtimeVersion.Patch)
}

func GetConfig(ctx string) HTMLCConfigFile {
	var res HTMLCConfigFile
	contextFiles, err := ioutil.ReadDir(ctx)
	check(err)
	for _, file := range contextFiles {
		if !file.IsDir() {
			fname := file.Name()
			if fname == "htmlc.json" {
				content, err := ioutil.ReadFile(path.Join(ctx, fname))
				check(err)
				json.Unmarshal([]byte(content), &res)
			}
		}
	}
	return res
}

func JoinPaths(base string, child string) string {
	return path.Join(base, child)
}

func DefaultConfig() HTMLCConfig {
	return HTMLCConfig{
		"views", "chunks",
	}
}
