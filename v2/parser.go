package main

import "strings"

func HasScope(content TokenString) bool {
	return strings.Contains(content, HTMLC_OPEN_SCOPE)
}
