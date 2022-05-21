package main

import "os"

func main() {

	userArgs := GetProcessArgs()
	if len(userArgs) > 0 {
		println(userArgs[0].Key)
	}

	cwd, err := os.Getwd()
	check(err)
	configJSON := GetConfig(cwd)
	// check(err)
	configStruct := configJSON["config"]
	loader := CreateLoader(configStruct, cwd)
	println(loader.ProcessPath)
}
