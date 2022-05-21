package main

type KeyList = []string
type KeyMap2D = []KeyList

type ProcessArg struct {
	Key   string
	Value string
}

type ProcessArgList = []ProcessArg

type HTMLCVersion struct {
	Major uint8
	Minor uint8
	Patch uint8
}

type HTMLCConfig struct {
	PathRoot  string `json: pathRoot`    // root path that your htmlc files will be in (relative to config)
	ChunkRoot string `json: "chunkRoot"` // path to read in files for runtime / compiler
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

type HTMLCLoader struct {
	ProcessPath    string      `json: "process_path"`
	Config         HTMLCConfig `json: "config"`
	ResolvedChunks []HTMLChunk
	CallableChunks []HTMLChunk
}

type TokenString = string

type HTMLCToken struct {
	Signature TokenString
}

type HTMLCDebugger struct {
	Loader HTMLCLoader
}
