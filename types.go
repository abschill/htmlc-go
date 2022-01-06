package up

/**
* Main Loader Options
**/
type LoaderOptions struct {
	base string
	templates string
	partials string
	debug bool
	watch bool
}

// @render-* input
type Input struct {
	Key string
	Value string
}

// Template/Partial Input Object
type InputFile struct {
	name string
	insertMap []Input
}

type HTMLChunk struct {
	name string
	path string
	_type string
	rawContent string
}

type Loader struct {
	config LoaderOptions
	templateData []HTMLChunk
	partialData []HTMLChunk
}

type PartialRaw struct {
	matcher string
	content string
}