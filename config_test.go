package main

import (
	"testing"
)

var testRules = []Rule{
	Rule{"/tmp/foo", "/usr/local/bar/foobar", "*.txt"},
	Rule{"/tmp/foo", "/usr/local/foo/foobar", "*.zip"},
	Rule{"/tmp/bar", "/usr/local/bar/barfoo", "*.jpg"},
}

func TestParseConfig(t *testing.T) {
	rules, err := ParseConfig(fixtures + "/example_config.json")
	if err != nil {
		t.Fatal(err)
	}

	for _, testRule := range testRules {
		found := false

		for _, rule := range rules {
			if testRule.Path == rule.Path && testRule.Run == rule.Run {
				found = true
			}
		}

		if !found {
			t.Errorf("Rule not found: %+v", testRule)
		}
	}
}
