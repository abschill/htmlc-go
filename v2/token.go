package main

import "strings"

type IType string

const (
	ISTART IType = "start" // start open chunk scope
	_ISET  IType = "macro" // set metadata about chunk type
	IOPEN  IType = "open"  // start open expr scope inside chunk
	ICLOSE IType = "close" // close IOPEN scope
	IBRK   IType = "break" // break line & any existing scope
	ICALL  IType = "call"  // call directive as argument
	ISET   IType = "set"   // set directive input to argument follow up
	IWRAP  IType = "wrap"  // wrap expression with a condition or configuration hook
	IEND   IType = "end"   // end the html chunk scope
)

type HTMLCToken struct {
	Name            string
	Signature       string
	InstructionType IType
	POptions        []HTMLCToken
}

// token that is resolved within a scope
type HTMLCResolvedToken struct {
	Token     HTMLCToken
	StartLine int
	EndLine   int
	StartCol  int
	EndCol    int
	//FromScope string
	//FromChunk string
}

var RawList = []HTMLCToken{
	tokenize("HTML_OC_SCOPE", "<!--@htmlc|", ISTART),
	tokenize("HTML_CC_SCOPE", "|@htmlc-->", IEND),
	tokenize("HTMLC_CM_PREFIX", "~", _ISET),
	tokenize("HTMLC_TD_OSCOPE", "(", IOPEN),
	tokenize("HTMLC_TD_CSCOPE", ")", ICLOSE),
	tokenize("HTMLC_TD_RENDER", "#render", ICALL),
	tokenize("HTMLC_TD_CHUNK", "#chunk", ICALL),
	tokenize("HTMLC_TO_SET", "=", ISET),
	tokenize("HTMLC_TD_ENFORCE", "!", IWRAP),
	tokenize("HTMLC_TD_TRY", "?", IWRAP),
}

func List() []HTMLCToken {
	return RawList
}

func GetTokenName(tag string) HTMLCToken {
	var matcher HTMLCToken

	for _, token := range RawList {
		if token.Name == tag {
			matcher = token
		}
	}
	return matcher
}

func (t HTMLCToken) IsIn(input string) bool {
	return strings.Contains(input, t.Signature)
}

// todo
func (HTMLCToken) GetFrom(input string) HTMLCResolvedToken {
	return HTMLCResolvedToken{}
}

// todo
func (HTMLCToken) Replace() string {
	var output string = ""

	return output
}

// internal struct mapping
func tokenize(name string, sig string, t IType) HTMLCToken {
	return HTMLCToken{
		Name:            name,
		Signature:       sig,
		InstructionType: t,
	}
}
