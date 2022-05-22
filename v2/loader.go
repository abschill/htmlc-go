package main

import (
	"io/ioutil"
	"path"
	"strings"

	"github.com/fatih/color"
)

type HTMLCLoader struct {
	ProcessPath    string      `json:"processPath"`
	Config         HTMLCConfig `json:"config"`
	ResolvedChunks []HTMLChunk
	CallableChunks []HTMLChunk
}

// todo - set up chunk finder, determine which ones are valid syntax and add them to callable chunks
func CreateLoader(config HTMLCConfig, processPath string) HTMLCLoader {
	var rChunks []HTMLChunk
	var cChunks []HTMLChunk
	chunkPath := path.Join(processPath, config.BasePath, config.ChunkPath)
	files, err := ioutil.ReadDir(chunkPath)
	for _, file := range files {
		fValid := false
		// todo - set up path discovery sometime
		if !file.IsDir() {
			fname := file.Name()
			// todo - iterate over standalone extension types
			if strings.Contains(fname, ".html") {
				if !fValid {
					fValid = true
					fpath := path.Join(chunkPath, fname)
					fbytes, err := ioutil.ReadFile(fpath)
					check(err)
					// println(fname)
					content := string(fbytes)
					isValid := ValidSyntax(content)
					theChunk := HTMLChunk{
						FilePath:      fpath,
						FileExtension: strings.Split(fname, ".")[1],
						IsStatic:      !HasScope(content),
						IsValid:       isValid,
						AsRaw:         content,
					}
					rChunks = append(rChunks, theChunk)
					if isValid {
						cChunks = append(cChunks, theChunk)
					}
				}
			}
		}
	}

	check(err)
	return HTMLCLoader{
		ProcessPath: processPath,
		Config: HTMLCConfig{
			BasePath:  config.BasePath,
			ChunkPath: chunkPath,
		},
		ResolvedChunks: rChunks,
		CallableChunks: cChunks,
	}
}

func (loader HTMLCLoader) Print() {
	//todo
	color.Green("%s: \n", "Resolved Chunk Path")
	println(loader.Config.ChunkPath)
	color.Blue("%s\n", "Chunks Found:")
	println(loader.ResolvedChunks[0].AsRaw)

}
