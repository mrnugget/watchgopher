package main

import (
	"testing"
)

var matchTests = []struct {
	rulePath, rulePattern string
	eventPath             string
	match                 bool
}{
	{"/a/b", "*.txt", "/a/b/z.txt", true},
	{"/a/b", "*.txt", "/a/c/z.txt", false},
	{"/a/b", "*.txt", "/a/b/d/z.txt", false},
	{"/a/b", "*.jpg", "/a/b/f.txt", false},
	{"/a/b", "*.jpg", "/a/b/z.jpg", true},
	{"/a/b", "*.jpg", "/a/b/z.jpeg", false},
	{"/a/b", "*", "/a/b/z.jpg", true},
	{"/a/b", "*", "/a/c/z.jpg", false},
}

func TestMatchingRule(t *testing.T) {
	for _, tt := range matchTests {
		rule := &Rule{tt.rulePath, "/bin/echo", tt.rulePattern}
		rules := []*Rule{rule}

		match := matchingRule(rules, tt.eventPath)
		if (tt.match && match == nil) || (!tt.match && match != nil) {
			t.Errorf("matchingRule(%+v, %v) = %v, want %v", rules, tt.eventPath, match, tt.match)
		}
	}
}
