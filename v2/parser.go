package main

import (
	"strings"
)

var SCOPE_SIG = GetTokenName("HTML_OC_SCOPE").Signature
var CLOSURE_SIG = GetTokenName("HTML_CC_SCOPE").Signature
var SCOPE_SIZE = len(SCOPE_SIG)
var CLOSURE_SIZE = len(CLOSURE_SIG)

func CountScopes(c string) int {
	return strings.Count(c, SCOPE_SIG)

}
func CountClosures(c string) int {
	return strings.Count(c, CLOSURE_SIG)
}
func HasScope(content string) bool {
	return strings.Contains(content, GetTokenName("HTML_OC_SCOPE").Signature)
}

func HasClosure(content string) bool {
	return strings.Contains(content, GetTokenName("HTML_CC_SCOPE").Signature)
}

func GetClosureIndex(content string) int {
	return strings.Index(content, GetTokenName("HTML_CC_SCOPE").Signature)
}

// todo - find errors preemptively in syntax to validate
func ValidSyntax(content string) bool {
	openSig := GetTokenName("HTML_OC_SCOPE").Signature
	closeSig := GetTokenName("HTML_CC_SCOPE").Signature
	if strings.Contains(content, openSig) {
		// no closing signature to match the opened scope
		if !strings.Contains(content, closeSig) {
			return false
		}
		// scope improperly ordered
		if strings.Index(content, openSig) > strings.Index(content, closeSig) {
			return false
		}
	}
	// todo - validate more than just that
	return true
}

/**
TODO:
1. for inputs that established a valid scope and are non-static, determine the following:

	- opening/closing line/col of EACH scope in the given input
	- for each scope, find the opening and closing position of each token within the SCOPE UNIT (the scope's opening index is the 0th possible index- we can calculate the offsets based on the parent if we need but I think it will be easier to do it
	that way compared to the other way around
	- for each token in each scope, determine the method of replacement for the token, based on the decision tree that we need to rewrite in backus naur form from the js concept

**/
