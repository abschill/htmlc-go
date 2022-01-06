package up

import (
	"io/ioutil"
)

func Check( e error ) {
	if e != nil {
		panic( e )
	}
}

func DefaultLoaderOptions() LoaderOptions {
	return LoaderOptions { 
		base: "views",
		templates: "pages",
		partials: "partials",
		debug: false,
		watch: false,
	};
}

func CustomLoaderOptions( base string, templates string, partials string, debug bool, watch bool ) LoaderOptions {
	return LoaderOptions { 
		base: base,
		templates: templates,
		partials: partials,
		debug: debug,
		watch: watch,
	};
}

func DefineLoader( config LoaderOptions ) Loader {

	baseDir := config.base + "/"
	templateDir := baseDir + config.templates + "/"
	partialDir := baseDir + config.partials + "/"
	
	templates, err := ioutil.ReadDir( templateDir )
	Check( err )

	partials, err := ioutil.ReadDir( partialDir )
	Check( err )

	var templateRawData = make( []HTMLChunk, len( templates ) )
	var partialRawData = make( []HTMLChunk, len( partials ) )
	for i, f := range templates {
		name := f.Name()
		path := templateDir + name
		data, err := ioutil.ReadFile( path )
		Check( err )
		d := string( data );
		templateRawData[i] = HTMLChunk {
			name: name,
			path: path,
			_type: "template",
			rawContent: d,
		}
	}

	for i, f := range partials {
		name := f.Name()
		path := partialDir + name
		data, err := ioutil.ReadFile( path )
		Check( err )
		d := string( data );
		partialRawData[i] = HTMLChunk {
			name: name,
			path: path,
			_type: "partial",
			rawContent: d,
		}
	}
	return Loader {
		config: config,
		partialData: partialRawData,
		templateData: templateRawData,
	}
}