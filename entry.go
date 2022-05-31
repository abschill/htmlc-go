package main

import (
	"github.com/abschill/htmlc-go"
)

func main() {
	// create loader from the config
	loader := htmlc.CreateLoader()
	//loader.Print()
	loader.PreloadTemplateData()

}
