package main

import "os"

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
	PrintLoader(loader)
}
