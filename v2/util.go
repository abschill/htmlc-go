package main

import (
	"fmt"
	"reflect"
)

type HTMLCInternal = interface {
	GetType() string
	IsType() bool
	GetVersion()
	GetValidExtensions() []string
	GetValidProcessArgs() KeyMap2D
}

const pkgName = "main"

func GetType(x interface{}) string {
	return reflect.TypeOf(x).String()
}

// prefixes the packge into the typename ie isType("Foo") checks "main.Foo" bc thats how go lang does it
func IsType(x interface{}, t string) bool {
	return GetType(x) == fmt.Sprintf("%s.%s", pkgName, t)
}

func GetVersion() string {
	runtimeVersion := HTMLCVersion{
		0, 1, 1,
	}
	return fmt.Sprintf("HTMLC Compiler Version: %d.%d.%d\n", runtimeVersion.Major, runtimeVersion.Minor, runtimeVersion.Patch)
}

func GetValidExtensions() []string {
	return []string{
		".htm",
		".html",
		".htmlc",
		".hcl",
	}
}

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
