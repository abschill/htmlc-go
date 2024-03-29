package core

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
	Raw           string
	Render        string
	Scopes        []HTMLCScope
	MinifyLoad    bool
}

type HTMLCScope struct {
	Raw        string
	TrimCache0 string
	CtxStart   int
	CtxEnd     int
}

// utility for sectioning off standard out in debugging
func (c HTMLChunk) Log() {
	color.Magenta("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	color.Green("Chunk Type: %s\nName: %s", c.ChunkType, c.ChunkName)
	color.Magenta("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
}

func (chunk HTMLChunk) GetScopeBounds() [][]int {
	var bounds [][]int
	for _, scope := range chunk.GetScopes() {
		bounds = append(bounds, []int{scope.CtxStart, scope.CtxEnd})
	}
	return bounds
}

func (chunk HTMLChunk) Print() {
	chunk.Log()
	if chunk.IsStatic {
		println("Static Content:")
		println(chunk.Raw)
		return
	}
	color.Yellow("Content Scopes:")
	for i, s := range chunk.GetScopes() {
		println(i)
		fmt.Printf("Scope Bounds: %d, %d\n", s.CtxStart, s.CtxEnd)
		println(s.Raw)
	}
}

func (chunk HTMLChunk) GetScopes() []HTMLCScope {
	var buf string = chunk.Raw
	oScopeCt := CountScopes(buf)
	cScopeCt := CountClosures(buf)
	// scope closures dont match openings
	if oScopeCt != cScopeCt {
		// we checked for this earlier but in case the input changes for some reason we check again here
		panic(fmt.Errorf("invalid scopes: closures dont match up | (scopes:closures) = %d:%d\nin file %s", oScopeCt, cScopeCt, chunk.FilePath))
	}
	// split to create a delimiter then reinsert it during iteration to create the scope object for the chunk
	temp := strings.Split(buf, SCOPE_SIG)
	for _, t := range temp {
		if HasClosure(t) && !HasScope(t) {
			// first, format the signature back into the iterator
			reFmt := fmt.Sprintf("%s%s", SCOPE_SIG, t)
			reFmt = reFmt[0 : GetClosureIndex(reFmt)+CLOSURE_SIZE]
			startPos := strings.Index(chunk.Raw, reFmt)
			chunk.Scopes = append(chunk.Scopes, HTMLCScope{
				Raw:        reFmt,
				TrimCache0: strings.Replace(strings.Replace(buf, SCOPE_SIG, "", oScopeCt), CLOSURE_SIG, "", cScopeCt),
				CtxStart:   startPos,
				CtxEnd:     startPos + len(reFmt),
			})
		}
	}
	return chunk.Scopes
}
