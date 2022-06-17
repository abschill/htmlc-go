package htmlc

import (
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

func getKeyIn(key string, items []PreloadDataItem) PreloadDataItem {
	var ret PreloadDataItem
	for _, di := range items {
		if di.Key == key {
			return di
		}
	}
	return ret
}

func PreRender(chunk HTMLChunk, items []PreloadDataItem) string {
	var renderedChunk string
	if !chunk.IsStatic {
		scopeList := chunk.GetScopes()
		for _, scope := range scopeList {
			var buf string = scope.Raw
			for _, key := range TopLevelTokenList {
				isMatch, matcher := key.MatchFunc(scope.Raw)
				if len(matcher.Matches) != 0 {
					//print("Bounds:\n")
					for _, k := range matcher.Starts {
						//fmt.Printf("%d, %d\n", k[0], k[1])
						expr := scope.Raw[k[0]:k[1]]
						switch key.Name {
						case "HTML_OC_SCOPE":
							buf = strings.Replace(buf, HTMLCOpenScope+"|", "", -1)
						case "HTML_CC_SCOPE":
							buf = strings.Replace(buf, "|"+HTMLCCloseScope, "", -1)
						case "HTMLC_TD_RENDER":
							_key := parseScopedExpressionKey(expr)
							keyIn := getKeyIn(_key, items)
							if keyIn.Type != "value" {
								panic("render args not of type 'value'")
							}
							replaceVal := keyIn.Value
							replaceSeg := scope.Raw[k[0]-1 : k[1]+1]
							buf = strings.Replace(buf, replaceSeg, replaceVal, -1)

						default:
							color.Yellow("unidentified token: %s", key.Name)
						}
					}
				}
				if !isMatch {
					panic("error in prerender with matcher")
				}
			}
			renderedChunk += buf
		}
	} else {
		renderedChunk += chunk.Raw
	}

	return strings.ReplaceAll(strings.ReplaceAll(renderedChunk, "\n\t", ""), "\t", "")
}
