package main

import (
	"testing"
)

var matchTests = []struct {
	rulePath, eventPath string
	match               bool
}{
	{"/a/b", "/a/b/z.txt", true},
	{"/a/b", "/a/c/z.txt", false},
	{"/a/b", "/a/b/d/z.txt", false},
}

func TestMatchingRule(t *testing.T) {
	for _, tt := range matchTests {
		rule := &Rule{tt.rulePath, "/bin/echo"}
		rules := []*Rule{rule}

		match := matchingRule(rules, tt.eventPath)
		if (tt.match && match == nil) || (!tt.match && match != nil) {
			t.Errorf("matchingRule(%v, %v) = %v, want %v", rules, tt.eventPath, match, tt.match)
		}
	}
}
