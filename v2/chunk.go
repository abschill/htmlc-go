package main

type HTMLChunk struct {
	ChunkType     string
	FilePath      string
	FileExtension string
	IsStatic      bool
	IsValid       bool
	AsRaw         string
	AsRender      string
}
