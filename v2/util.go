package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"reflect"
)

func getType(x interface{}) string {
	return reflect.TypeOf(x).String()
}

func GetVersionInfo() string {
	runtimeVersion := HTMLCVersion{
		0, 1, 1,
	}
	return fmt.Sprintf("HTMLC Compiler Version: %d.%d.%d\n", runtimeVersion.Major, runtimeVersion.Minor, runtimeVersion.Patch)
}

func GetValidExtensions() []string {
	return []string{
		".htm",
		".html",
		".htmlc",
		".hcl",
	}
}

func GetFSOptions(ctx string) HTMLCConfigFile {
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

func getOptionsFSToConfig(config map[string]HTMLCConfig) HTMLCConfig {
	return config["config"]
}

func JoinPaths(base string, child string) string {
	return path.Join(base, child)
}

func DefaultConfig() HTMLCConfig {
	return HTMLCConfig{
		BasePath:  "views",
		ChunkPath: "chunks",
		WritePath: "htmlc-out",
		LogPath:   "htmlc-log",
	}
}
