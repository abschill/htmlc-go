package main

import (
	"io/ioutil"
	"path"
	"strings"

	"github.com/fatih/color"
)

// structure for the loader to call on the server side
type HTMLCLoader struct {
	ProcessPath    string      `json:"processPath"`
	Config         HTMLCConfig `json:"config"`
	ResolvedChunks []HTMLChunk
	CallableChunks []HTMLChunk
}

// create loader(s) for process
func CreateLoader(config HTMLCConfig, processPath string) HTMLCLoader {
	var rChunks []HTMLChunk
	var cChunks []HTMLChunk
	chunkPath := path.Join(processPath, config.BasePath, config.ChunkPath)
	files, err := ioutil.ReadDir(chunkPath)
	check(err)
	for _, file := range files {
		// todo - set up path discovery sometime
		if !file.IsDir() {
			fname := file.Name()
			// todo - iterate over standalone extension types
			if strings.Contains(fname, ".html") {
				fpath := path.Join(chunkPath, fname)
				fbytes, err := ioutil.ReadFile(fpath)
				check(err)
				content := string(fbytes)
				isValid := ValidSyntax(content)
				splitName := strings.Split(fname, ".")
				theChunk := HTMLChunk{
					ChunkType:     "chunk",
					ChunkName:     splitName[0],
					FilePath:      fpath,
					FileExtension: splitName[1],
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

// print data of loader to standard out
func (loader HTMLCLoader) Print() {
	//todo
	color.HiBlue("%s: %s", "chunk path results from:", loader.Config.ChunkPath)

	for _, chunk := range loader.ResolvedChunks {
		chunk.Print()
	}
}
