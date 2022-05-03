package up

import "regexp"

func HasLoop(content string) bool {
	matches, _ := regexp.MatchString(ASTLoopOpenReggie, content)
	return matches
}

func FindLoops(content string) []string {
	r, _ := regexp.Compile(ASTLoopReggie)
	items := r.FindAllString(content, -1)
	return items
}

func FindLoopOpenIndexes(content string) [][]int {
	r, _ := regexp.Compile(ASTLoopOpenReggie)
	items := r.FindAllStringIndex(content, -1)
	return items
}

func FindLoopCloseIndexes(content string) [][]int {
	r, _ := regexp.Compile(ASTLoopCloseReggie)
	items := r.FindAllStringIndex(content, -1)
	return items
}

func HasPartials(content string) bool {
	matches, _ := regexp.MatchString(ASTPartialReggie, content)
	return matches
}

func FindPartials(content string) []string {
	r, _ := regexp.Compile(ASTPartialReggie)
	return r.FindAllString(content, -1)
}

func FindPartialIndexes(content string) [][]int {
	r, _ := regexp.Compile(ASTPartialReggie)
	return r.FindAllStringIndex(content, -1)
}

func HasKeys(content string) bool {
	matches, _ := regexp.MatchString(ASTKeyReggie, content)
	return matches
}

func FindKeys(content string) []string {
	r, _ := regexp.Compile(ASTKeyReggie)
	items := r.FindAllString(content, -1)
	return items
}

func FindKeyIndexes(content string) [][]int {
	r, _ := regexp.Compile(ASTKeyReggie)
	return r.FindAllStringIndex(content, -1)
}
