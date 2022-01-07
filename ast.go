package up

import "regexp"

/**
 * Partials
**/
const PARTIAL_MATCH = "<!--@render-partial=([a-z | 0-9 | _ | -]+)-->"

func HasPartials( content string ) bool {
	matches, _ := regexp.MatchString( PARTIAL_MATCH , content )
	return matches;
}

func FindPartials( content string ) []string {
	r, _ := regexp.Compile( PARTIAL_MATCH )
	items := r.FindAllString( content, -1 )
	return items;
}

func FindPartialIndexes( content string ) [][]int {
	r, _ := regexp.Compile( PARTIAL_MATCH )
	items := r.FindAllStringIndex( content, -1 )
	return items;
}

/**
 * Keys
**/
const KEY_MATCH = "<!--@render=([a-z | 0-9 | _ | -]+)-->"

func HasKeys( content string ) bool {
	matches, _ := regexp.MatchString( KEY_MATCH, content )
	return matches;
}

func FindKeys( content string ) []string {
	r, _ := regexp.Compile( KEY_MATCH )
	items := r.FindAllString( content, -1 )
	return items;
}

func FindKeyIndexes( content string ) [][]int {
	r, _ := regexp.Compile( KEY_MATCH )
	items := r.FindAllStringIndex( content, -1 )
	return items;
}


/**
 * Iterators
**/
const FOR_NAME_MATCH = "<!--@for((.*?)){[\n?|\t|\r| .]+\n?}"

const FOR_CONTENT_MATCH = "<!--@for((.*?)){\n?(\n*[\n*|\t*]*.+?)\n?}"

func HasLoop( content string ) bool {
	matches, _ := regexp.MatchString( FOR_CONTENT_MATCH, content )
	return matches;
}

func FindLoops( content string ) []string {
	r, _ := regexp.Compile( FOR_CONTENT_MATCH )
	items := r.FindAllString( content, -1 )
	return items;
}

func FindLoopIndexes( content string ) [][]int {
	r, _ := regexp.Compile( FOR_CONTENT_MATCH )
	items := r.FindAllStringIndex( content, -1 )
	return items;
}