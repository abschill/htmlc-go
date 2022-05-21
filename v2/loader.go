package main

import (
	"os"
)

// todo - set up chunk finder, determine which ones are valid syntax and add them to callable chunks
func CreateLoader(config HTMLCConfig, processPath string) HTMLCLoader {
	return HTMLCLoader{
		ProcessPath:    processPath,
		Config:         config,
		ResolvedChunks: nil,
		CallableChunks: nil,
	}
}

func main() {
	userArgs := GetProcessArgs()
	cwd, err := os.Getwd()
	check(err)
	if len(userArgs) > 0 {
		for _, arg := range userArgs {
			if arg.Key == "configPath" {
				cwd = arg.Value
			}
		}
	}

	// this gets the full config file with .config as a property
	fsOptions := GetFSOptions(cwd)
	// get .config property from full file options
	config := getOptionsFSToConfig(fsOptions)
	// create loader from the config
	loader := CreateLoader(config, cwd)
	println(getType(loader))
}
