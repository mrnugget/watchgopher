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
	{"/a/b", "*.txt", "/a/b/Z.TXT", false},
	{"/a/b", "*.TXT", "/a/b/Z.TXT", true},
	{"/a/b", "*.txt", "/a/b/d/z.txt", false},
	{"/a/b", "*.jpg", "/a/b/z.txt", false},
	{"/a/b", "*.jpg", "/a/b/z.jpg", true},
	{"/a/b", "*.jpg", "/a/b/z.jpeg", false},
	{"/a/b", "*", "/a/b/z.jpg", true},
	{"/a/b", "*", "/a/b/z.txt", true},
	{"/a/b", "*", "/a/c/z.jpg", false},
	{"/a/b", "IMG_12*.jpg", "/a/b/IMG_123.jpg", true},
	{"/a/b", "IMG_13*.jpg", "/a/b/IMG_123.jpg", false},
	{"/a/b", "*_chapter.md", "/a/b/1st_chapter.md", true},
	{"/a/b", "*_chapter.md", "/a/b/1.chapter.md", false},
	{"/a/b", "1st_chapter.md", "/a/b/1st_chapter.md", true},
}

func TestMatchingRules(t *testing.T) {
	for _, tt := range matchTests {
		rule := &Rule{tt.rulePath, "/bin/echo", tt.rulePattern}
		rules := []*Rule{rule}

		matches := matchingRules(rules, tt.eventPath)
		if (tt.match && len(matches) == 0) || (!tt.match && len(matches) > 0) {
			t.Errorf("(len(matchingRules(%+v, %v)) > 0) == %v, want %v", rules, tt.eventPath, len(matches) > 0, tt.match)
		}
	}
}
