package main

import "github.com/fatih/color"

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
	color.Yellow("Type: %s\nName: %s", chunk.ChunkType, chunk.ChunkName)
	color.Green("Raw Content: ")
	color.Magenta("Scoped Bounds:")
	scopeBounds := ResolveNextScopeBounds(chunk.AsRaw)

	if scopeBounds[1] > -1 {
		println("scope 1~~~~~~~~~~~~")
		println(chunk.AsRaw[scopeBounds[0]:])
		println("~~~~~~~~~~~~")
		var subTest string
		// close scope signature
		eSig := GetTokenName("HTML_CC_SCOPE").Signature
		offset := (scopeBounds[1] + len(eSig))
		subTest = chunk.AsRaw[offset:]
		println("scope 2~~~~~~~~~~~~")
		println(subTest)
		println("~~~~~~~~~~~~")
	} else {
		println("static scope~~~~~~~~~~~~")
		println(chunk.AsRaw)
		println("~~~~~~~~~~~~")
	}

}
