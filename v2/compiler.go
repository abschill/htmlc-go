package main

func CreateAST(chunk HTMLChunk) {
	//kReggie, err := regexp.Compile(KeyReggie)
	//check(err)
	//keys := kReggie.FindAllString(chunk.AsRaw, -1)
	// take parser result and compile output - todo
	LogRawChunkHeader(chunk)
	for _, key := range List() {
		matches := key.eMatcher.FindAllString(chunk.AsRaw, -1)
		isMatch := key.eMatcher.MatchString(chunk.AsRaw)
		if isMatch {
			println("Current Instruction:")
			println(key.InstructionType)
			println(key.rMatcher)
			println(matches[0])
			println("Next Instruction:")
			println(key.iNext)
			println("~~~~~~~~~~")
			// for _, match := range matches {
			// 	println(strings.Index(chunk.AsRaw, match))
			// }
		}
	}
	// println(List()[0].rMatcher)

}
