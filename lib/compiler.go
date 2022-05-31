package main

import (
	"fmt"

	"github.com/fatih/color"
)

func CreateAST(chunk HTMLChunk) {
	LogRawChunkHeader(chunk)
	println(chunk.AsRaw)
	for _, key := range List() {
		println("Checking Key")
		println(key.Name)
		isMatch, matcher := key.MatchFunc(chunk)
		if len(matcher.Matches) != 0 {
			print("Bounds:\n")
			for _, k := range matcher.Starts {
				fmt.Printf("%d, %d\n", k[0], k[1])
			}
		}

		if isMatch {
			for _, x := range key.IContext.rProps {
				println("Validate Prior:")
				color.Yellow("%s\n", x)
			}
			for _, x := range key.IContext.rProps {
				println("Needs Followup:")
				color.Blue("%s\n", x)
			}
			println("~~~~~~~~~~")
		}
	}
}
