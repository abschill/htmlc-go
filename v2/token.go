package main

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

var AssignedTokens = 0

// token that is resolved within a chunk
type HTMLCResolvedToken struct {
	Token     HTMLCToken
	StartLine int
	EndLine   int
	StartCol  int
	EndCol    int
	FromScope string
	FromChunk string
}

func (HTMLCToken) Exists() bool {
	return false
}

func (HTMLCToken) Get(input string) HTMLCResolvedToken {
	return HTMLCResolvedToken{}
}

func (HTMLCToken) Replace() bool {
	return false
}

func tokenize(name string, sig string, t IType) HTMLCToken {
	return HTMLCToken{
		Name:            name,
		Signature:       sig,
		InstructionType: t,
	}
}

var HTMLC_TOKENS = []HTMLCToken{
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
