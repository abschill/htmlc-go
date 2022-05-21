package main

type HTMLCLoader struct {
	ProcessPath    string      `json:"processPath"`
	Config         HTMLCConfig `json:"config"`
	ResolvedChunks []HTMLChunk
	CallableChunks []HTMLChunk
}

// todo - set up chunk finder, determine which ones are valid syntax and add them to callable chunks
func CreateLoader(config HTMLCConfig, processPath string) HTMLCLoader {
	return HTMLCLoader{
		ProcessPath:    processPath,
		Config:         config,
		ResolvedChunks: nil,
		CallableChunks: nil,
	}
}
