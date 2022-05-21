package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func GetVersionInfo() string {
	runtimeVersion := HTMLCVersion{
		0, 1, 1,
	}
	return fmt.Sprintf("HTMLC Compiler Version: %d.%d.%d\n", runtimeVersion.Major, runtimeVersion.Minor, runtimeVersion.Patch)
}

func GetAllowedExtensions() []string {
	return []string{
		".htm",
		".html",
		".htmlc",
		".hcl",
	}
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

func GetProcessArgs() ProcessArgList {

	validProcessArgs := KeyMap2D{
		{"-c", "--config-file"},
		{"-l", "--loader-file"},
		{"-d", "--debug-file"},
	}
	var validOptions ProcessArgList
	// iterate over validProcessArgs if the user has entered args inline, otherwise ignore

	argv := os.Args
	argc := len(argv)

	if argc > 2 {
		for i, arg := range argv {
			for _, option := range validProcessArgs {
				isMatch := option[0] == arg || option[1] == arg
				if isMatch {
					validOptions = append(validOptions, ProcessArg{
						Key:   arg,
						Value: argv[i+1],
					})
				}
			}
		}
	}
	return validOptions
}

func JoinPaths(base string, child string) string {
	return path.Join(base, child)
}

func DefaultConfig() HTMLCConfig {
	return HTMLCConfig{
		PathRoot:  "views",
		ChunkRoot: "chunks",
	}
}
