package main

// todo - set up chunk finder, determine which ones are valid syntax and add them to callable chunks
func CreateLoader(config HTMLCConfig, processPath string) HTMLCLoader {
	return HTMLCLoader{
		ProcessPath:    processPath,
		Config:         config,
		ResolvedChunks: nil,
		CallableChunks: nil,
	}
}
