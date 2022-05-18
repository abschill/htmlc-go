package upv1

/**
* Main Loader Options
**/
type LoaderOptions struct {
	base      string
	templates string
	partials  string
	debug     bool
	watch     bool
}

const (
	Single    = 0
	MultiPart = 1
)

// @render-* input
type Input struct {
	Key   string
	Value string
}

type ArrInput struct {
	Key   string
	Value []Input
}

// Template/Partial Input Object
type InputFile struct {
	name      string
	insertMap []Input
}

type HTMLChunk struct {
	name       string
	path       string
	_type      string
	rawContent string
}

type LoopMatch struct {
	Open, Close int
}

//reserved
type DebugOptions struct {
	LogFile  string
	LogLevel uint8
}
