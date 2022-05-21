package main

import (
	"os"
)

func main() {

	userArgs := GetProcessArgs()
	if len(userArgs) > 0 {
		for _, arg := range userArgs {
			PrintArg(arg.Key, arg.Value)
		}
	}

	cwd, err := os.Getwd()
	check(err)
	configJSON := GetConfig(cwd)
	configStruct := configJSON["config"]
	loader := CreateLoader(configStruct, cwd)
	println(loader.ProcessPath)
}
