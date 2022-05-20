package main

type version struct {
	Major uint8
	Minor uint8
	Patch uint8
}

type Config struct {
	Root   string `json: "root"`
	Chunks string `json: "chunks"`
}
