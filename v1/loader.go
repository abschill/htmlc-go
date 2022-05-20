package upv1

import (
	"fmt"
	"strings"
)

type Loader struct {
	config       LoaderOptions
	templateData []HTMLChunk
	partialData  []HTMLChunk
}

type PartialRaw struct {
	matcher string
	content string
}

/**
 * Returns Array of PartialRaw for the runtime to get ready to Load into a template
**/
func (loader *Loader) loadPartials(matcher string, index int) []PartialRaw {
	_name := FindPartialFileName(matcher)
	var partialsToLoad = make([]PartialRaw, len(loader.partialData))
	for _, p := range loader.partialData {
		if p.name == _name {
			partialsToLoad[index] = PartialRaw{
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
func (loader *Loader) retrievePartials(content string) string {
	temp := content
	matches := HasPartials(content)
	if matches {
		items := FindPartials(content)
		_items := FindPartialIndexes(content)
		for i := range _items {
			todos := loader.loadPartials(items[i], i)
			for _, todo := range todos {
				temp = strings.Replace(temp, todo.matcher, todo.content, -1)
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
func (loader *Loader) preRender(name string) string {
	fileName := NameHTML(name)
	for _, v := range loader.templateData {
		if v.name == fileName {
			rawContent := loader.retrievePartials(v.rawContent)
			return rawContent
		}
	}

	panic("No Template Match for" + fileName)
}

func (loader *Loader) Template(name string, data []Input) string {
	content := loader.preRender(name)
	hasKeys := HasKeys(content)
	hasLoops := HasLoop(content)
	if !hasKeys {
		return content
	}
	keys := FindKeys(content)
	_keys := FindKeyIndexes(content)

	if hasLoops {
		loopOpens := FindLoopOpenIndexes(content)
		loopCloses := FindLoopCloseIndexes(content)

		if len(loopOpens) == len(loopCloses) {
			for i, item := range loopOpens {
				_close := loopCloses[i]

				loopContent := content[item[0]:_close[1]]
				fmt.Println(loopContent)
			}
		} else {
			fmt.Println("error")
		}

	}

	for i := range _keys {
		keyName := FindKeyName(keys[i])
		for _, input := range data {
			if input.Key == keyName {
				content = strings.Replace(content, keys[i], input.Value, -1)
			}
		}
	}
	return content
}