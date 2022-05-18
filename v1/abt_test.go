package upv1

import "testing"

func TestHasPartials(t *testing.T) {
	content0 := "<!--@partial=test-->"
	match0 := HasPartials(content0)

	if match0 != true {
		t.Errorf("Has Partials Failed - Content 0")
	}

	content1 := "hello world"
	match1 := HasPartials(content1)

	if match1 == true {
		t.Errorf("Has Partials Failed - Content 1")
	}
}

func TestHasKeys(t *testing.T) {
	match0 := HasKeys("<!--@render=test-->")
	match1 := HasKeys("hello")

	if match0 != true {
		t.Errorf("Has Keys Failed - Match 0")
	}

	if match1 == true {
		t.Errorf("Has Keys Failed - Match 1")
	}
}

func TestHasLoop(t *testing.T) {
	match0 := HasLoop(`<!--@loop=test(
<div>{_}</foo>
)-->`)
	if match0 != true {
		t.Errorf("Has Loop Failed - Match 0")
	}

	match1 := HasLoop(`<!--@loop=test(
		<div>{_}</foo>
)-->`)

	if match1 != true {
		t.Errorf("Has Loop Failed - Match 1")
	}

	match2 := HasLoop(`<!--@loop=test(
		<div>{_}</foo>)-->`)

	if match2 != true {
		t.Errorf("Has Loop Failed - Match 2")
	}

	match3 := HasLoop(`<!--@loop=test(<div>{_}</foo>)-->`)

	if match3 != true {
		t.Errorf("Has Loop Failed - Match 3")
	}
}
