package main

import (
	"io/ioutil"
)

type HTMLChunk struct {
	ChunkType     string
	FilePath      string
	FileExtension string
	IsStatic      bool
	IsValid       bool
	AsRaw         string
	AsRender      string
}

func RegisterChunk(chunkPath string) HTMLChunk {
	file, err := ioutil.ReadFile(chunkPath)
	cString := string(file)
	check(err)
	println(cString)
	return HTMLChunk{
		FilePath: chunkPath,
		AsRaw:    cString,
	}
}
