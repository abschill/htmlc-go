package main

import (
	"fmt"
	"reflect"

	"github.com/fatih/color"
)

type KeyList = []string
type KeyMap2D = []KeyList

type HTMLCVersion struct {
	Major uint8
	Minor uint8
	Patch uint8
}

// name of pkg for type string resolution
const pkgName = "main"

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// split off a section of standard out with a yellow line - for logging
func LogSection() {
	color.Yellow("%s", "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
}

// split off a subsection of standard out for logging
func LogSubSection() {
	println("_________________________________________________")
}

// utility for sectioning off standard out in debugging
func LogRawChunkHeader(c HTMLChunk) {
	color.Magenta("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	color.Green("Chunk Type: %s\nName: %s", c.ChunkType, c.ChunkName)
	color.Magenta("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
}

func GetType(x interface{}) string {
	return reflect.TypeOf(x).String()
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
