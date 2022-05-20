package main

type Version struct {
	Major uint8
	Minor uint8
	Patch uint8
}

type Config struct {
	Root   string `json: "root"`   // root path that your htmlc files will be in (relative to config)
	Chunks string `json: "chunks"` // path to read in files for runtime / compiler
}

type ConfigFile = map[string]Config
