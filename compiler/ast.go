package compiler

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

// get var name to load into expression
func parseScopedExpressionKey(expr string) string {
	splitter := strings.Split(expr, HTMLChunkEQ)
	if len(splitter) > 1 {
		return splitter[1]
	}
	return ""
}

func CreateAST(chunk HTMLChunk, scope HTMLCScope) {
	// println(chunk.AsRaw)
	println(scope.Raw)

	for _, key := range List() {
		isMatch, matcher := key.MatchFunc(scope.Raw)
		println("Checking Key")
		println(key.Name)

		if len(matcher.Matches) != 0 {
			print("Bounds:\n")
			for _, k := range matcher.Starts {
				fmt.Printf("%d, %d\n", k[0], k[1])
				expr := scope.Raw[k[0]:k[1]]
				if key.Name == "HTMLC_TD_RENDER" {
					color.Green("%s\n", expr)
					color.Green("Data Key: %s\n", parseScopedExpressionKey(expr))
				}
			}
		}

		if isMatch {
			// for _, x := range key.IContext.rProps {
			// 	println("Validate Prior:")
			// 	color.Yellow("%s\n", x)
			// }
			// for _, x := range key.IContext.rProps {
			// 	println("Needs Followup:")
			// 	color.Blue("%s\n", x)
			// }
			println("~~~~~~~~~~")
		}
	}
}
