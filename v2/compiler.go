package main

import (
	"github.com/fatih/color"
)

func CreateAST(chunk HTMLChunk) {
	//kReggie, err := regexp.Compile(KeyReggie)
	//check(err)
	//keys := kReggie.FindAllString(chunk.AsRaw, -1)
	// take parser result and compile output - todo
	LogRawChunkHeader(chunk)
	//color.Magenta("Raw Content:")
	println(chunk.AsRaw)
	for _, key := range List() {
		println("Checking Key")
		println(key.Name)
		matches := key.iMatchReggie.FindStringSubmatch(chunk.AsRaw)
		isMatch := key.iMatchReggie.MatchString(chunk.AsRaw)

		if isMatch {
			println("Current Instruction:")
			//color.Blue("%s\n", key.InstructionType)
			print("Found ", len(matches), " ", key.iMatchString, " matches\n")
			println(matches[0])
			if key.iPrev != INULL {
				println("Validate Prior Ins:")
				color.Yellow("%s\n", key.iPrev)
			}

			println("Next Instruction:")
			color.Blue("%s\n", key.iNext)
			println("~~~~~~~~~~")
			// for _, match := range matches {
			// 	println(strings.Index(chunk.AsRaw, match))
			// }
		}
	}
	// println(List()[0].rMatcher)

}
