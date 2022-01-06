package main

import (
	"fmt"
	"github.com/abschill/underpin"
)

func main() {
	//Call up.DefaultLoaderOptions() to stick with the default options 
	options := up.CustomLoaderOptions( "root", "pages", "partials", false, false )

	//Define Loader with the generated LoaderOptions
	loader := up.DefineLoader( options )

	//Create a Slice for Template Data
	var data = make( []up.Input, 1 )
	//Fill in our page_title example as an Input
	data[0] = up.Input{
		Key: "page_title",
		Value: "Underpin",
	}
	//Retrieve the Template result from loader
	fmt.Println( loader.Template( "home", data ) );
	
}