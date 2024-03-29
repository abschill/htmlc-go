package core

import (
	"regexp"
	"strings"
)

type IType string

const HTMLCOpenScope = "<!--@htmlc"
const HTMLCCloseScope = "@htmlc-->"
const HTMLChunkMacroPrefix = "~"
const HTMLChunkScopeOpen = "(("
const HTMLChunkScopeClose = "))"
const HTMLChunkRender = "@render"
const HTMLChunkLoop = "@loop"
const HTMLChunkRenderChunk = "@chunk"
const HTMLChunkDirective = "[" + HTMLChunkRender + "|" + HTMLChunkRenderChunk + "|" + HTMLChunkLoop + "]"
const HTMLChunkEQ = "="
const HTMLChunkEnf = "!"
const HTMLChunkTry = "\\?"
const HTMLChunkExpandOpen = "{"
const HTMLChunkExpandClose = "}"
const HTMLCAnyChars = "[a-z | 0-9 | _ | -]"

/**
 * Keys that map to preloads / inlines
**/
const HTMLChunkRenderMatch = HTMLChunkScopeOpen + HTMLChunkRender + HTMLChunkEQ + HTMLCAnyChars + "+" + HTMLChunkScopeClose

const HTMLChunkChunkMatch = HTMLChunkScopeOpen + HTMLChunkRenderChunk + HTMLChunkEQ + HTMLCAnyChars + "+" + HTMLChunkScopeClose

const (
	ISTART   IType = "start"   // start open chunk scope
	_ISET    IType = "macro"   // set metadata about chunk type
	IOPEN    IType = "open"    // start open expr scope inside chunk
	ICLOSE   IType = "close"   // close IOPEN scope
	IBRK     IType = "break"   // break line & any existing scope
	ICALL    IType = "call"    // call directive as argument
	ISET     IType = "set"     // set directive input to argument follow up
	IWRAP    IType = "wrap"    // wrap expression with a condition or configuration hook
	IEND     IType = "end"     // end the html chunk scope
	IEXPAND  IType = "expand"  // expand a scope to new line
	ICEXPAND IType = "cexpand" // close expansion
	INULL    IType = "dpad"    // no-op, pad data (mask slots), so this just accumulates everything into a scope
	IPUT     IType = "put"     // user input
)

type HTMLCToken struct {
	Name            string
	Signature       string
	InstructionType IType
	iMatchString    string
	iMatchReggie    regexp.Regexp
	ValidChildren   HTMLCSubToken
}

type HTMLCSubToken struct {
	Name            string
	Signature       string
	InstructionType IType
	iMatchString    string
	iMatchReggie    regexp.Regexp
}

type ITypeScope struct {
	rProps  []IType
	rFollow []IType
}

var TopLevelTokenList = []HTMLCToken{
	tokenize("HTML_OC_SCOPE", HTMLCOpenScope, ISTART, HTMLCOpenScope),
	tokenize("HTMLC_TD_RENDER", HTMLChunkRender, ICALL, HTMLChunkRenderMatch),
	tokenize("HTMLC_TD_CHUNK", HTMLChunkRenderChunk, ICALL, HTMLChunkChunkMatch),
	tokenize("HTML_CC_SCOPE", HTMLCCloseScope, IEND, HTMLCCloseScope),
}

// get a token by its name
func GetTokenName(tag string) HTMLCToken {
	var matcher HTMLCToken

	for _, token := range TopLevelTokenList {
		if token.Name == tag {
			matcher = token
		}
	}
	return matcher
}

func (t HTMLCToken) IsIn(input string) bool {
	return strings.Contains(input, t.Signature)
}

type TokenMatchData struct {
	Starts  [][]int
	Matches []string
}

func (t HTMLCToken) MatchFunc(scopeString string) (bool, TokenMatchData) {
	isMatch := t.iMatchReggie.MatchString(scopeString)
	if isMatch {
		matches := t.iMatchReggie.FindAllString(scopeString, -1)
		matchesIndices := t.iMatchReggie.FindAllStringIndex(scopeString, -1)
		return true, TokenMatchData{
			Starts:  matchesIndices,
			Matches: matches,
		}
	} else {
		return true, TokenMatchData{}
	}
}

// internal struct mapping
func tokenize(name string, sig string, t IType, matcher string) HTMLCToken {
	reg, err := regexp.Compile(matcher)
	if err != nil {
		panic(err)
	}

	return HTMLCToken{
		Name:            name,
		Signature:       sig,
		InstructionType: t,
		iMatchString:    matcher,
		iMatchReggie:    *reg,
	}
}
