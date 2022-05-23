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
}

func (chunk HTMLChunk) Print() {
	color.Green("Type: %s\nName: %s", chunk.ChunkType, chunk.ChunkName)
	if chunk.IsStatic {
		println("static scope~~~~~~~~~~~~")
		println(chunk.AsRaw)
		return
	}
	color.Yellow("Scoped Bounds:")
	// scope := HTMLCScope{
	// 	Parent: chunk,
	// }
	//	scopeBounds := scope.ResolveNextScopeBounds()
	chunk.GetScopes()

}

func (chunk HTMLChunk) ResolveNextScopeBounds() []int {
	openSig := GetTokenName("HTML_OC_SCOPE").Signature
	closeSig := GetTokenName("HTML_CC_SCOPE").Signature
	headIndex := strings.Index(chunk.AsRaw, openSig)
	tailIndex := strings.Index(chunk.AsRaw, closeSig)
	return []int{headIndex, tailIndex}
}

func (chunk HTMLChunk) ResolveLastScopeBounds() []int {
	openSig := GetTokenName("HTML_OC_SCOPE").Signature
	closeSig := GetTokenName("HTML_CC_SCOPE").Signature
	headIndex := strings.LastIndex(chunk.AsRaw, openSig)
	tailIndex := strings.LastIndex(chunk.AsRaw, closeSig)
	return []int{headIndex, tailIndex}
}

func (chunk HTMLChunk) GetScopes() {
	var buf string
	//var output string
	openSig := GetTokenName("HTML_OC_SCOPE").Signature
	closeSig := GetTokenName("HTML_CC_SCOPE").Signature
	oScopeCt := strings.Count(chunk.AsRaw, openSig)
	cScopeCt := strings.Count(chunk.AsRaw, closeSig)
	nextBounds := chunk.ResolveNextScopeBounds()
	lastBounds := chunk.ResolveLastScopeBounds()
	if oScopeCt != cScopeCt {
		panic("invalid scopes")
	}

	fmt.Printf("Next Scope: %d, %d\n", nextBounds[0], nextBounds[1])
	fmt.Printf("Last Scope: %d, %d\n", lastBounds[0], lastBounds[1])
	println(len(openSig))
	println(len(closeSig))
	buf = strings.Replace(chunk.AsRaw, openSig, "", 1)
	buf = strings.Replace(buf, closeSig, "", 1)
	println(buf)
}
