package main

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

type HTMLChunkType string

const (
	Chunk HTMLChunkType = "chunk"
	Page  HTMLChunkType = "page"
)

type HTMLChunk struct {
	ChunkType     HTMLChunkType
	ChunkName     string
	FilePath      string
	FileExtension string
	IsStatic      bool
	IsValid       bool
	AsRaw         string
	AsRender      string
	Scopes        []HTMLCScope
}

type HTMLCScope struct {
	Raw     string
	Trimmed string
}

func (chunk HTMLChunk) Print() {
	LogSection()
	color.Green("Chunk Type: %s\nName: %s", chunk.ChunkType, chunk.ChunkName)
	if chunk.IsStatic {
		println("Static Content:")
		println(chunk.AsRaw)
		return
	}
	color.Yellow("Content Scopes:")
	chunk.GetScopes()
}

func (chunk HTMLChunk) GetScopes() {

	var buf string = chunk.AsRaw
	//var output string
	openSig := GetTokenName("HTML_OC_SCOPE").Signature
	closeSig := GetTokenName("HTML_CC_SCOPE").Signature
	oScopeCt := strings.Count(buf, openSig)
	cScopeCt := strings.Count(buf, closeSig)
	// scope closures dont match openings
	if oScopeCt != cScopeCt {
		panic("invalid scopes")
	}
	// split to create a delimiter then reinsert it during iteration to create the scope object for the chunk
	test := strings.Split(buf, openSig)
	for _, t := range test {
		if HasClosure(t) && !HasScope(t) {
			LogSubSection()
			reFmt := fmt.Sprintf("%s%s", openSig, t)
			reFmt = reFmt[0 : GetClosureIndex(reFmt)+CLOSURE_SIZE]
			println(reFmt)
		}
	}

	//	buf = strings.Replace(buf, openSig, "", oScopeCt)
	//	buf = strings.Replace(buf, closeSig, "", cScopeCt)
	//println(buf)
}
