package main

import "strings"

type TokenString = string

const HTMLC_OPEN_SCOPE TokenString = "<!--@htmlc|"
const HTMLC_CLOSE_SCOPE TokenString = "|@htmlc-->"

type ParsableToken struct {
	str TokenString
}

func getAllowedExtensions() []string {
	return []string{
		".htm",
		".html",
		".htmlc",
		".hcl",
	}
}

func hasScope(content TokenString) bool {
	return strings.Contains(content, HTMLC_OPEN_SCOPE)
}
