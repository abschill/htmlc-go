package main

type KeyList = []string
type KeyMap2D = []KeyList

type HTMLCVersion struct {
	Major uint8
	Minor uint8
	Patch uint8
}

type HTMLCConfigFile = map[string]HTMLCConfig

type HTMLChunk struct {
	ChunkType     string
	FilePath      string
	FileExtension string
	IsStatic      bool
	IsValid       bool
	AsRaw         string
	AsRender      string
}

type TokenString = string

type HTMLCToken struct {
	Signature TokenString
}
