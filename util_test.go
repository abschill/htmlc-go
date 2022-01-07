package up

import "testing"


func TestNameHTML( t *testing.T ) {
	name := NameHTML( "test" )
	want := "test.html"

	if name != want {
		t.Errorf( "Name HTML Failed got %q, wanted %q", name, want )
	}
}

func TestFindPartialFileName( t *testing.T ) {
	name := FindPartialFileName( "<!--@render-partial=test-->" )
	want := "test.html"

	if name != want {
		t.Errorf( "Find Partial Name Failed, got %q wanted %q", name, want )
	}
}

func TestFindKeyName( t *testing.T ) {
	name := FindKeyName( "<!--@render=test-->" )
	want := "test"

	if name != want {
		t.Errorf( "Find Partial Name Failed, got %q wanted %q", name, want )
	}
}