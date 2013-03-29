package main

import (
	"testing"
)

func TestMatchingRule(t *testing.T) {
	testingRule := &Rule{"/a/b", "/bin/date"}
	rules := []*Rule{testingRule}

	match := matchingRule(rules, "/a/b/z.txt")
	if match == nil {
		t.Fatal("Did not match the correct rule")
	}

	match = matchingRule(rules, "/a/c/z.txt")
	if match != nil {
		t.Fatal("Matched the wrong rule")
	}

	match = matchingRule(rules, "/a/b/d/z.txt")
	if match != nil {
		t.Fatal("Matched the wrong rule")
	}
}
