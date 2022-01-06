package main

import (
	"fmt"
	"underpin"
)

func main() {
	options := up.CustomLoaderOptions( "root", "pages", "partials", false, false )
	loader := up.DefineLoader( options )
	var data = make( []up.Input, 1 )
	data[0] = up.Input{
		Key: "page_title",
		Value: "Underpin",
	}
	fmt.Println( loader.Template( "home", data ) );
	
}