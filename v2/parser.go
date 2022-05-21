package main

import "strings"

func GetAllowedExtensions() []string {
	return []string{
		".htm",
		".html",
		".htmlc",
		".hcl",
	}
}

func HasScope(content TokenString) bool {
	return strings.Contains(content, HTMLC_OPEN_SCOPE)
}
