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
const HTMLCAnyChars = "[a-z | 0-9 | _ | -]"

/**
 * Chunk Calls to other chunks within the given lodaer
**/
const ChunkReggie = HTMLChunkRenderChunk + HTMLChunkEQ + "((" + HTMLCAnyChars + "+))"

/**
 * Keys that map to preloads / inlines
**/
const KeyReggie = HTMLChunkRender + HTMLChunkEQ + "((" + HTMLCAnyChars + "+))"

/**
 * Iterators that map to preloads / inlines
**/
const LoopOpenReggie = HTMLChunkLoop + HTMLChunkEQ + "((" + HTMLCAnyChars + "+))"

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
	SLen            int
	InstructionType IType
	iMatchString    string
	iMatchReggie    regexp.Regexp
	IContext        ITypeScope
	IsAnyP          bool
	IsAnyN          bool
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
	tokenize("HTMLC_TD_IPUT", HTMLCAnyChars, IPUT, HTMLCAnyChars),
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
	var ctx ITypeScope
	// determine which type of instruction is to be expected after the current one to establish a valid syntax tree
	switch cType {
	case IWRAP:
		ctx.rProps = append(ctx.rProps, INULL)
		ctx.rFollow = append(ctx.rFollow, IOPEN) // the thing you're wrapping
	case IOPEN:
		ctx.rProps = append(ctx.rProps, INULL)   // pad data within closed scope
		ctx.rFollow = append(ctx.rFollow, ICALL) // call scoped expression
	case ICLOSE:
		ctx.rProps = append(ctx.rProps, IPUT)   // pad data within closed scope
		ctx.rFollow = append(ctx.rFollow, IBRK) // should break line after scope end
	case ICALL:
		ctx.rProps = append(ctx.rProps, IOPEN)  // should not be calling a directive outside of a scoped ()
		ctx.rFollow = append(ctx.rFollow, ISET) // should follow by setting the macro with an ISET
	case _ISET:
		ctx.rProps = append(ctx.rProps, INULL)   // macros can come afte rother macros
		ctx.rFollow = append(ctx.rFollow, INULL) // macros prepend the padded content
	case ISET:
		ctx.rProps = append(ctx.rProps, ICALL)   // call directive to assign to the given data
		ctx.rFollow = append(ctx.rFollow, INULL) // call data from preloaded / inlined data
	case IEXPAND:
		ctx.rProps = append(ctx.rProps, IPUT)
		ctx.rFollow = append(ctx.rFollow, INULL)
	case ICEXPAND:
		ctx.rProps = append(ctx.rProps, INULL)
		ctx.rFollow = append(ctx.rFollow, ICLOSE)
	default:

	}
	return HTMLCToken{
		Name:            name,
		Signature:       sig,
		SLen:            len(sig),
		InstructionType: t,
		iMatchString:    matcher,
		iMatchReggie:    *reg,
		IContext:        ctx,
	}
}
