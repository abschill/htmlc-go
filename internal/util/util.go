package util

import (
	"fmt"
	"reflect"

	"github.com/abschill/htmlc-go/internal/env"
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
	runtimeVersion := env.HTMLCVersion{
		Major: env.HTMLC_VERSION_MAJOR, Minor: env.HTMLC_VERSION_MINOR, Patch: env.HTMLC_VERSION_PATCH,
	}
	return fmt.Sprintf("HTMLC Compiler Version: %d.%d.%d\n", runtimeVersion.Major, runtimeVersion.Minor, runtimeVersion.Patch)
}

func GetVersion(seg string) uint8 {
	switch seg {
	case "patch":
		return env.HTMLC_VERSION_PATCH
	case "minor":
		return env.HTMLC_VERSION_MINOR
	case "major":
		return env.HTMLC_VERSION_MAJOR
	default:
		return env.HTMLC_VERSION_MAJOR
	}
}

func GetValidChunkExtensions() []string {
	return env.HTMLC_VALID_EXTENSIONS
}

func GetValidProcessArgs() env.KeyMap2D {
	return env.HTMLC_VALID_pARGS
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}
