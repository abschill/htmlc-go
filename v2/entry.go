package main

func main() {
	// create loader from the config
	loader := CreateLoader()
	//loader.Print()
	loader.PreloadTemplateData()

}
