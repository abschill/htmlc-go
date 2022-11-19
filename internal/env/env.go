package env

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
