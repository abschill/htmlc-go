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
const HTMLChunkExpandOpen = "{"
const HTMLChunkExpandClose = "}"
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
	iPrev           IType
	iNext           IType
	IsAnyNext       bool
}

type ITypeScope struct {
	rProps  []IType
	rFollow []IType
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
	tokenize("HTMLC_TD_EXPAND", HTMLChunkExpandOpen, IEXPAND, HTMLChunkExpandOpen),
	tokenize("HTMLC_TD_cEXPAND", HTMLChunkExpandClose, ICEXPAND, HTMLChunkExpandClose),
	tokenize("HTMLC_TD_IPUT", HTMLCValidCharset, IPUT, HTMLCValidCharset),
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

// internal struct mapping
func tokenize(name string, sig string, t IType, matcher string) HTMLCToken {
	reg, err := regexp.Compile(matcher)
	if err != nil {
		panic("error setting up tokenizer")
	}
	cType := t
	var iPrev IType
	var iNext IType
	// determine which type of instruction is to be expected after the current one to establish a valid syntax tree
	switch cType {
	case IWRAP:
		iPrev = INULL
		iNext = IOPEN // the thing you're wrapping
	case IOPEN:
		iPrev = INULL // pad data within closed scope
		iNext = ICALL // call scoped expression
	case ICLOSE:
		iPrev = IPUT // pad data within closed scope
		iNext = IBRK // should break line after scope end
	case ICALL:
		iPrev = IOPEN // should not be calling a directive outside of a scoped ()
		iNext = ISET  // should follow by setting the macro with an ISET
	case _ISET:
		iPrev = INULL // macros can come afte rother macros
		iNext = INULL // macros prepend the padded content
	case ISET:
		iPrev = ICALL // call directive to assign to the given data
		iNext = INULL // call data from preloaded / inlined data
	case IEXPAND:
		iPrev = IPUT
		iNext = INULL
	case ICEXPAND:
		iPrev = INULL
		iNext = ICLOSE
	default:
		iPrev = INULL
		iNext = INULL
	}
	return HTMLCToken{
		Name:            name,
		Signature:       sig,
		InstructionType: t,
		iMatchString:    matcher,
		iMatchReggie:    *reg,
		iNext:           iNext,
		iPrev:           iPrev,
	}
}
