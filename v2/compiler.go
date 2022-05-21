package main

import "os"

func main() {

	// this is our json file with the structed type

	//staticPath := JoinPaths(cwd, "static/basic")
	//

	//PrintContextInfo(cwd, staticPath)
	//PrintScopes(staticPath, files)

	cwd, err := os.Getwd()
	check(err)
	configJSON := GetConfig(cwd)
	// check(err)
	configStruct := configJSON["config"]
	loader := CreateLoader(configStruct, cwd)
	println(loader.ProcessPath)
}
