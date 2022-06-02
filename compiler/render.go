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

func PreRender(chunk HTMLChunk) {
	println(chunk.Raw)
	scopeList := chunk.GetScopes()
	if !chunk.IsStatic {
		for _, scope := range scopeList {
			for _, key := range TopLevelTokenList {
				isMatch, matcher := key.MatchFunc(scope.Raw)
				println("Checking Key")
				println(key.Name)

				if len(matcher.Matches) != 0 {
					print("Bounds:\n")
					for _, k := range matcher.Starts {
						fmt.Printf("%d, %d\n", k[0], k[1])
						expr := scope.Raw[k[0]:k[1]]
						switch key.Name {
						case "HTMLC_TD_RENDER":
							color.Green("%s\n", expr)
							color.Green("Data Key: %s\n", parseScopedExpressionKey(expr))
						default:
							color.Yellow("Todo: %s", key.Name)
						}
					}
				}

				if isMatch {
					println("~~~~~~~~~~")
				}
			}
		}
	} else {
		chunk.Render = chunk.Raw
	}

}
