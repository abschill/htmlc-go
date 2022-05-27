package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"

	"github.com/fatih/color"
)

type HTMLCConfigFile struct {
	Config      HTMLCConfig  `json:"config"`
	PreloadData []ProcessArg `json:"preload"`
}

type ProcessArg struct {
	Type  string `json:"type"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ProcessArgList = []ProcessArg

type HTMLCConfig struct {
	BasePath  string `json:"basePath"`  // root path that your htmlc files will be in (relative to config)
	ChunkPath string `json:"chunkPath"` // path to read in files for runtime / compiler
	WritePath string `json:"writePath"` // path to write any static files to
	LogPath   string `json:"logPath"`
}

func asProcessArg(key string, val string) ProcessArg {
	return ProcessArg{
		Key:   key,
		Value: val,
	}
}

// get arguments entered by the user into the process itself
func GetProcessArgs() ProcessArgList {
	validProcessArgs := GetValidProcessArgs()
	var validOptions ProcessArgList
	argv := os.Args
	argc := len(argv)
	//  if its greater than 2 it has named args to parse from the process itself
	if argc > 2 {
		for i, arg := range argv {
			for _, option := range validProcessArgs {
				if option[0] == arg || option[1] == arg {
					validOptions = append(validOptions, asProcessArg(arg, argv[i+1]))
				}
			}
		}
		return validOptions
	} else if argc == 2 {
		// if its equal to 2, handle the anonymous arg as the path lookup argument.
		validOptions = append(validOptions, asProcessArg("configPath", argv[argc-1]))
	}
	return validOptions
}

// print key/val pair from args
func PrintTuple(k string, v string) {
	color.Green("%s: %s", k, v)
}

// get options file as unmarshalled JSON
func getFSOptions() (HTMLCConfigFile, string) {
	userArgs := GetProcessArgs()
	ctx, err := os.Getwd()
	check(err)
	if len(userArgs) > 0 {
		for _, arg := range userArgs {
			switch arg.Key {
			case "configPath":
				ctx = arg.Value
			}
		}
	}
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
	return res, ctx
}

// get config object as struct from unmarshalled config file
func (config HTMLCConfigFile) getOptionsFSToConfig() HTMLCConfig {
	return config.Config
}

func (config HTMLCConfigFile) getOptionsFSToInput() []ProcessArg {
	return config.PreloadData
}

// get default config options
func DefaultConfig() HTMLCConfig {
	return HTMLCConfig{
		BasePath:  "views",
		ChunkPath: "chunks",
		WritePath: "htmlc-out",
		LogPath:   "htmlc-log",
	}
}
