package up
import (
	"strings"
)
/**
 * Returns Array of PartialRaw for the runtime to get ready to Load into a template
**/
func ( loader *Loader ) LoadPartials( matcher string, index int ) []PartialRaw {
	_name := ParsePartialFileName( matcher )
	var partialsToLoad = make( []PartialRaw, len( loader.partialData ) )
	for _, p := range loader.partialData {
		if p.name == _name {
			partialsToLoad[index] = PartialRaw {
				matcher: matcher,
				content: p.rawContent,
			}
		}
	}
	return partialsToLoad
}

/**
 * Returns Template as a string with partials inserted (pre-render)
**/
func ( loader *Loader ) retrievePartials( content string ) string {
	temp := content
	matches := HasPartials( content )
	if matches != false {
		items := FindPartials( content )
		_items := FindPartialIndexes( content )
		for i, _ := range _items {
			todos := loader.LoadPartials( items[i], i ) 
			for _, todo := range todos {
				temp = strings.Replace( temp, todo.matcher, todo.content, -1 )
			}
		}
	}	
	return temp
}

/**
 * Template Render Cycle ->
 * Pre Render & Render
 * TemplatePreRender handles pre-rendering partials into the template 
 * TemplateRender handles rendering key/values into the template with it's partials preloaded
**/
func ( loader *Loader ) TemplatePreRender( name string ) string {
	fileName := NameHTML( name )
	for _, v := range loader.templateData {
		if v.name == fileName {
			rawContent := loader.retrievePartials( v.rawContent )
			return rawContent
		}
	}

	panic( "No Template Match for" + fileName )
}

func ( loader *Loader ) TemplateRender( name string, data []Input ) string {
	content := loader.TemplatePreRender( name )
	hasKeys := HasKeys( content )
	if hasKeys == false {
		return content
	}
	keys := FindKeys( content )
	_keys := FindKeyIndexes( content )

	for i, _ := range _keys {
		keyName := ParseKeyName( keys[i] );
		for _, input := range data {
			if input.Key == keyName {
				content = strings.Replace( content, keys[i], input.Value, -1 )
			}
		}
	}
	return content
}