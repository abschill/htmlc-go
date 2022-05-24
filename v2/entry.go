package main

import (
	"fmt"
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
	input := fsOptions.getOptionsFSToInput()
	// create loader from the config
	loader := CreateLoader(config, cwd)
	loader.Print()
	tt := []string{
		"foo",
	}
	println("user input from config:")
	for _, key := range tt {
		// todo - handle for non-string
		res := fmt.Sprintf("{\"%s\":\"%s\"}", key, input[key].(string))
		println(res)
	}

}
