package up

import "testing"

func TestHasPartials( t *testing.T ) {
	content := "<!--@render-partial=test-->"
	match := HasPartials( content )

	if match != true {
		t.Errorf( "Has Partials Failed" )
	}
}

func TestHasKeys( t *testing.T ) {
	content := "<!--@render=test-->"
	match := HasKeys( content )

	if match != true {
		t.Errorf( "Has Keys Failed" )
	}
}