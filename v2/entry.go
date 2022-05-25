package main

import (
	"os"
)

func main() {
	userArgs := GetProcessArgs()
	cwd, err := os.Getwd()
	check(err)
	if len(userArgs) > 0 {
		for _, arg := range userArgs {
			switch arg.Key {
			case "configPath":
				cwd = arg.Value
			}
		}
	}

	// this gets the full config file with .config as a property
	fsOptions := getFSOptions(cwd)
	// get .config property from full file options
	config := fsOptions.getOptionsFSToConfig()
	// create loader from the config
	loader := CreateLoader(config, cwd)
	loader.Print()
	// define keys from config that you want to bring into the runtime

	// println("user input from config:")
	/**
	ok so we need to take an array of key value pairs from the config file and then determine the type of the value in each
	once we have the value we will know how to render it against any template expression, so we will need to set up a struct of some sort
	that can identify the bindable methods for the given handle type.
	**/

}
