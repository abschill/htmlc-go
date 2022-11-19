package main

import "github.com/abschill/htmlc-go/pkg/core"

func main() {
	// create loader from the config
	loader := core.CreateLoader()
	loader.Preload()
}
