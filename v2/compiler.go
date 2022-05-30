package main

import (
	"github.com/fatih/color"
)

func CreateAST(chunk HTMLChunk) {
	LogRawChunkHeader(chunk)
	println(chunk.AsRaw)
	for _, key := range List() {
		println("Checking Key")
		println(key.Name)
		matches := key.iMatchReggie.FindStringSubmatch(chunk.AsRaw)
		matchesIndices := key.iMatchReggie.FindStringSubmatchIndex(chunk.AsRaw)
		isMatch := key.iMatchReggie.MatchString(chunk.AsRaw)
		if len(matchesIndices) != 0 {
			print("Bounds:\n")
			print(matchesIndices[0])
			print(", ", matchesIndices[0]+key.SLen, "\n")
		}

		if isMatch {
			println("Current Instruction:")
			color.Blue("%s\n", key.InstructionType)
			print("Found ", len(matches), " ", key.iMatchString, " matches\n")
			println(matches[0])
			if len(key.IContext.rProps) != 0 {
				println("Validate Prior Ins:")
				color.Yellow("%s\n", key.IContext.rProps[0])
			}

			if len(key.IContext.rFollow) != 0 {
				println("Next Instruction:")
				color.Blue("%s\n", key.IContext.rFollow[0])
				println("~~~~~~~~~~")
			}
		}
	}
}
