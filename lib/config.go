package main

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"os"
	"path"

	"github.com/fatih/color"
)

type HTMLCTLOpts struct {
	Config      HTMLCConfig       `json:"config"`
	PreloadData []PreloadDataItem `json:"preload"`
}

type PreloadDataItem struct {
	Type  string `json:"type"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ProcessArg struct {
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

// check the file name to see if it matches the possible config names and if so, read it into struct ptr
func checkForConfig(ctx string, f fs.FileInfo, res *HTMLCTLOpts) {
	fileName := f.Name()
	for _, fno := range GetConfigFNameOptions() {
		if fileName == fno {
			topLevelOptions, err := ioutil.ReadFile(path.Join(ctx, fileName))
			check(err)
			json.Unmarshal([]byte(topLevelOptions), &res)
			return
		}
	}
}

// get options file as unmarshalled JSON
func getTopLevelOptions() (HTMLCTLOpts, string) {
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
	var res HTMLCTLOpts
	contextFiles, err := ioutil.ReadDir(ctx)
	check(err)
	for _, file := range contextFiles {
		if !file.IsDir() {
			checkForConfig(ctx, file, &res)
		} else {
			childCtx, err := ioutil.ReadDir(path.Join(ctx, file.Name()))
			check(err)
			for _, ffile := range childCtx {
				// dont recur it because they may not have actually meant to leave config out
				checkForConfig(ctx, ffile, &res)
			}
		}
	}
	return res, ctx
}

// get config data key as struct from unmarshalled config
func (config HTMLCTLOpts) GetConfig() HTMLCConfig {
	return config.Config
}

// get preload data key from the unmarshalled config
func (config HTMLCTLOpts) GetPreloads() []PreloadDataItem {
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

// get all possible names for config file
func GetConfigFNameOptions() []string {
	return []string{
		"htmlc.json",
		"hcl-config.json",
		"htmlc-config.json",
		".htmlc",
		"htmlc.conf",
	}
}
