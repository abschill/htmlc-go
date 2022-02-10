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
const LOOP_OPEN_MATCH = "<!--@loop=(.*?)"
const LOOP_CLOSE_MATCH = "\\)-->"
const LOOP_MATCH = "<!--@loop=(.*?)[\n*\t*]*?(<\\w+>.*)\n?\\)-->"

func HasLoop( content string ) bool {
	matches, _ := regexp.MatchString( LOOP_OPEN_MATCH, content )
	return matches;
}

func FindLoops( content string ) []string {
	r, _ := regexp.Compile( LOOP_MATCH )
	items := r.FindAllString( content, -1 )
	return items;
}

func FindLoopOpenIndexes( content string ) [][]int {
	r, _ := regexp.Compile( LOOP_OPEN_MATCH )
	items := r.FindAllStringIndex( content, -1 )
	return items;
}

func FindLoopCloseIndexes( content string ) [][]int {
	r, _ := regexp.Compile( LOOP_CLOSE_MATCH )
	items := r.FindAllStringIndex( content, -1 )
	return items;
}