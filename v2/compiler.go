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
		//matches := key.iMatchReggie.FindStringSubmatch(chunk.AsRaw)
		//matchesIndices := key.iMatchReggie.FindStringSubmatchIndex(chunk.AsRaw)

		//isMatch := key.iMatchReggie.MatchString(chunk.AsRaw)
		isMatch, matcher := key.MatchFunc(chunk)

		if len(matcher.Starts) != 0 {
			print("Bounds:\n")
			print(matcher.Starts[0])
			print(", ", matcher.Ends[0], "\n")
		}

		if isMatch {
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
