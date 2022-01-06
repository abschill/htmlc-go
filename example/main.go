package main

import (
	"fmt"
	"underpin"
)

func main() {
	// options := up.DefaultLoaderOptions()
	options := up.CustomLoaderOptions( "views2", "pages", "partials", false, false )
	loader := up.DefineLoader( options )
	var data = make( []up.Input, 1 )
	data[0] = up.Input{
		Key: "page_title",
		Value: "Underpin",
	}
	fmt.Println( loader.TemplateRender( "home", data ) );
	
}