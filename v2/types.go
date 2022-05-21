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
	BasePath  string `json: "basePath"`  // root path that your htmlc files will be in (relative to config)
	ChunkPath string `json: "chunkPath"` // path to read in files for runtime / compiler
	WritePath string `json: "writePath"` // path to write any static files to
	LogPath   string `json: "logPath"`
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
