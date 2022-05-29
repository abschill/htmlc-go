package main

type KeyList = []string
type KeyMap2D = []KeyList

type HTMLCVersion struct {
	Major uint8
	Minor uint8
	Patch uint8
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

const HTMLC_VERSION_MAJOR uint8 = 0
const HTMLC_VERSION_MINOR uint8 = 1
const HTMLC_VERSION_PATCH uint8 = 0

var HTMLC_VALID_EXTENSIONS = []string{
	".htm",
	".html",
	".htmlc",
	".hcl",
}

var HTMLC_VALID_pARGS = KeyMap2D{
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
