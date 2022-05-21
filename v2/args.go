package main

import "os"

func GetValidProcessArgs() KeyMap2D {
	return KeyMap2D{
		{
			"-c",
			"--config-file",
		},
		{
			"-l",
			"--loader-file",
		},
		{
			"-d",
			"--debug-file",
		},
	}
}

func GetProcessArgs() ProcessArgList {
	validProcessArgs := GetValidProcessArgs()
	var validOptions ProcessArgList
	argv := os.Args
	argc := len(argv)
	//  if its greater than 2 it has named args to parse from the process itself
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
		return validOptions
	} else if argc == 2 {
		// if its equal to 2, handle the anonymous arg as the path lookup argument.
		validOptions = append(validOptions, ProcessArg{
			Key:   "configPath",
			Value: argv[argc-1],
		})
	}
	return validOptions
}
