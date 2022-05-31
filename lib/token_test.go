package main

import (
	"testing"
)

func TestList(t *testing.T) {
	tokens := List()
	// test the head token item to be the opening matcher
	if tokens[0].Signature != "<!--@htmlc|" {
		t.Errorf("Error with signature: %s\n", tokens[0].Name)
	}
}

func TestGetTokenName(t *testing.T) {
	res := GetTokenName("HTML_OC_SCOPE")
	if res.InstructionType != ISTART {
		t.Errorf("Error sign ISTART token")
	}
	res = GetTokenName("HTML_CC_SCOPE")
	if res.InstructionType != IEND {
		t.Errorf("Error sign IEND token")
	}
}
