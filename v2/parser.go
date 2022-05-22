package main

import "strings"

func HasScope(content string) bool {
	return strings.Contains(content, "@htmlc")
}

/**
TODO:
1. for inputs that established a valid scope and are non-static, determine the following:

	- opening/closing line/col of EACH scope in the given input
	- for each scope, find the opening and closing position of each token within the SCOPE UNIT (the scope's opening index is the 0th possible index- we can calculate the offsets based on the parent if we need but I think it will be easier to do it
	that way compared to the other way around
	- for each token in each scope, determine the method of replacement for the token, based on the decision tree that we need to rewrite in backus naur form from the js concept

**/
