package main

import "regexp"

func CreateAST(chunk HTMLChunk) {
	kReggie, err := regexp.Compile(KeyReggie)
	check(err)
	keys := kReggie.FindAllString(chunk.AsRaw, -1)
	// take parser result and compile output - todo

	for i, key := range keys {
		println(i)
		println(key)
	}
}
