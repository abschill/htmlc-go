package main

import (
	"strings"
)

type HTMLCScope struct {
	Parent HTMLChunk
	Outer  string
	Inner  string
}

func HasScope(content string) bool {
	return strings.Contains(content, GetTokenName("HTML_OC_SCOPE").Signature)
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

// when we get the scope of HTMLC within the html comment we want to get it as a full scope first, this will trim the closures from that initial scope that is retrieved to begin to compile the output
// func (s HTMLCScope) ToList() {
// 	var buf string
// 	//var output string
// 	openSig := GetTokenName("HTML_OC_SCOPE").Signature
// 	closeSig := GetTokenName("HTML_CC_SCOPE").Signature
// 	oScopeCt := strings.Count(s.Parent.AsRaw, openSig)
// 	cScopeCt := strings.Count(s.Parent.AsRaw, closeSig)
// 	nextBounds := s.ResolveNextScopeBounds()
// 	lastBounds := s.ResolveLastScopeBounds()
// 	// buf = s.Parent.AsRaw
// 	if oScopeCt != cScopeCt {
// 		panic("invalid scopes")
// 	}
// 	// make offset 1
// 	// for i < (oScopeCt + 1) {
// 	// 	println(i)
// 	// 	i++
// 	// }
// 	fmt.Printf("Next Scope: %d, %d\n", nextBounds[0], nextBounds[1])
// 	fmt.Printf("Last Scope: %d, %d\n", lastBounds[0], lastBounds[1])
// 	println(len(openSig))
// 	println(len(closeSig))
// 	buf = strings.Replace(s.Parent.AsRaw, openSig, "", 1)
// 	buf = strings.Replace(buf, closeSig, "", 1)
// 	println(buf)
// }
