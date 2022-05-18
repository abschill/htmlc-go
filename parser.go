package main

import "strings"

const HTMLC_OPEN_SCOPE = "<!--@htmlc|"
const HTMLC_CLOSE_SCOPE = "|@htmlc-->"

type HTMLCToken interface{}

func getAllowedExtensions() []string {
	return []string{
		".htm",
		".html",
		".htmlc",
		".hcl",
	}
}

func hasScope(content string) bool {
	return strings.Contains(content, HTMLC_OPEN_SCOPE)
}
