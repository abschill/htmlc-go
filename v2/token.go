package main

import (
	"regexp"
	"strings"
)

type IType string

const HTMLCOpenScope = "<!--@htmlc|"
const HTMLCCloseScope = "|@htmlc-->"
const HTMLChunkMacroPrefix = "~"
const HTMLChunkScopeOpen = "\\("
const HTMLChunkScopeClose = "\\)"
const HTMLChunkRender = "@render"
const HTMLChunkLoop = "@loop"
const HTMLChunkRenderChunk = "@chunk"
const HTMLChunkEQ = "="
const HTMLChunkEnf = "!"
const HTMLChunkTry = "\\?"
const HTMLCValidCharset = "[a-z | 0-9 | _ | -]"

/**
 * Chunk Calls to other chunks within the given lodaer
**/
const ChunkReggie = HTMLChunkRenderChunk + HTMLChunkEQ + "((" + HTMLCValidCharset + "+))"

/**
 * Keys that map to preloads / inlines
**/
const KeyReggie = HTMLChunkRender + HTMLChunkEQ + "((" + HTMLCValidCharset + "+))"

/**
 * Iterators that map to preloads / inlines
**/
const LoopOpenReggie = HTMLChunkLoop + HTMLChunkEQ + "((" + HTMLCValidCharset + "+))"

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
	INULL  IType = "null"  // no-op, pad data (mask slots), so this just accumulates everything into a scope
)

type HTMLCToken struct {
	Name            string
	Signature       string
	InstructionType IType
	rMatcher        string
	eMatcher        regexp.Regexp
	iNext           IType
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
	tokenize("HTML_OC_SCOPE", HTMLCOpenScope, ISTART, HTMLCOpenScope),
	tokenize("HTML_CC_SCOPE", HTMLCCloseScope, IEND, HTMLCCloseScope),
	tokenize("HTMLC_CM_PREFIX", HTMLChunkMacroPrefix, _ISET, HTMLChunkMacroPrefix),
	tokenize("HTMLC_TD_OSCOPE", HTMLChunkScopeOpen, IOPEN, HTMLChunkScopeOpen),
	tokenize("HTMLC_TD_CSCOPE", HTMLChunkScopeClose, ICLOSE, HTMLChunkScopeClose),
	tokenize("HTMLC_TD_RENDER", HTMLChunkRender, ICALL, HTMLChunkRender),
	tokenize("HTMLC_TD_CHUNK", HTMLChunkRenderChunk, ICALL, HTMLChunkRenderChunk),
	tokenize("HTMLC_TO_SET", HTMLChunkEQ, ISET, HTMLChunkEQ),
	tokenize("HTMLC_TD_ENFORCE", HTMLChunkEnf, IWRAP, HTMLChunkEnf),
	tokenize("HTMLC_TD_TRY", HTMLChunkTry, IWRAP, HTMLChunkTry),
}

func List() []HTMLCToken {
	return RawList
}

// get a token by its name
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
func tokenize(name string, sig string, t IType, matcher string) HTMLCToken {
	reg, err := regexp.Compile(matcher)
	if err != nil {
		panic("error setting up tokenizer")
	}
	cType := t
	var iNext IType
	// determine which type of instruction is to be expected after the current one to establish a valid syntax tree
	switch cType {
	case IWRAP:
		iNext = IOPEN
	case IOPEN:
		iNext = ICALL
	case ICLOSE:
		iNext = IBRK
	case ICALL:
		iNext = ISET
	case _ISET:
		iNext = INULL
	default:
		iNext = INULL
	}
	return HTMLCToken{
		Name:            name,
		Signature:       sig,
		InstructionType: t,
		rMatcher:        matcher,
		eMatcher:        *reg,
		iNext:           iNext,
	}
}
