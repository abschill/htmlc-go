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
	println(chunk.AsRaw)
	color.Magenta("Scoped Bounds:")
	println(ResolveNextScopeBounds(chunk.AsRaw))
}
