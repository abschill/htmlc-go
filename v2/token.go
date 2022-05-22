package main

type IType string

const (
	ENDL IType = "end"
	CALL IType = "call"
	SET  IType = "set"
)

type HTMLCToken struct {
	Name            string
	Signature       string
	InstructionType IType
	POptions        []HTMLCToken
}

// token that is resolved within a chunk
type HTMLCResolvedToken struct {
	StartLine int
	EndLine   int
	StartCol  int
	EndCol    int
}

func (HTMLCToken) Exists() bool {
	return false
}

func (HTMLCToken) Get() HTMLCResolvedToken {
	return HTMLCResolvedToken{}
}

func (HTMLCToken) Replace() bool {
	return false
}

func tokenize(name string, sig string) HTMLCToken {
	return HTMLCToken{
		Name:      name,
		Signature: sig,
	}
}

var HTMLC_TOKENS = []HTMLCToken{
	tokenize("HTML_OC_SCOPE", "<!--@htmlc|"),
	tokenize("HTML_CC_SCOPE", "|@htmlc-->"),
	tokenize("HTMLC_TD_OSCOPE", "("),
	tokenize("HTMLC_TD_CSCOPE", ")"),
	tokenize("HTMLC_TD_RENDER", "#render"),
	tokenize("HTMLC_TD_CHUNK", "#chunk"),
	tokenize("HTMLC_TO_SET", "="),
	tokenize("HTMLC_TD_ENFORCE", "!"),
	tokenize("HTMLC_TD_TRY", "?"),
	tokenize("HTMLC_CM_PREFIX", "~"),
}
