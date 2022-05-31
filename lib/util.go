package main

import (
	"fmt"
	"reflect"

	"github.com/fatih/color"
)

// split off a section of standard out with a yellow line - for logging
func LogSection() {
	color.Yellow("%s", "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
}

// split off a subsection of standard out for logging
func LogSubSection() {
	println("_________________________________________________")
}

func GetType(x interface{}) string {
	return reflect.TypeOf(x).String()
}

func GetVersionString() string {
	runtimeVersion := HTMLCVersion{
		HTMLC_VERSION_MAJOR, HTMLC_VERSION_MINOR, HTMLC_VERSION_PATCH,
	}
	return fmt.Sprintf("HTMLC Compiler Version: %d.%d.%d\n", runtimeVersion.Major, runtimeVersion.Minor, runtimeVersion.Patch)
}

func GetVersion(seg string) uint8 {
	switch seg {
	case "patch":
		return HTMLC_VERSION_PATCH
	case "minor":
		return HTMLC_VERSION_MINOR
	case "major":
		return HTMLC_VERSION_MAJOR
	default:
		return HTMLC_VERSION_MAJOR
	}
}

func GetValidChunkExtensions() []string {
	return HTMLC_VALID_EXTENSIONS
}

func GetValidProcessArgs() KeyMap2D {
	return HTMLC_VALID_pARGS
}
