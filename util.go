package up

import (
	"strings"
)


func NameHTML( n string ) string {
	return n + ".html"
}

func ParsePartialFileName( p string ) string {
	str := strings.Split( strings.Split( p, "=" )[1], "-->" )[0] + ".html" 
	return str
}

func ParseKeyName( k string ) string {
	str := strings.Split( strings.Split( k, "=" )[1], "-->" )[0] 
	return str
}