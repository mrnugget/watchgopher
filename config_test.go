package main

import (
	"testing"
)

var testRules = []Rule{
	Rule{"/tmp/foo", "/usr/local/bar/foobar", "*.txt", false, false},
	Rule{"/tmp/foo", "/usr/local/foo/foobar", "*.zip", true, true},
	Rule{"/tmp/bar", "/usr/local/bar/barfoo", "*.jpg", false, false},
}

func TestParseConfig(t *testing.T) {
	rules, err := ParseConfig(fixtures + "/example_config.json")
	if err != nil {
		t.Fatal(err)
	}

	for _, testRule := range testRules {
		found := false

		for _, rule := range rules {
			found = found || (testRule.Path == rule.Path &&
			testRule.Run == rule.Run &&
			testRule.Pattern == rule.Pattern &&
			testRule.ChangePwd == rule.ChangePwd)
		}

		if !found {
			t.Errorf("Rule not found: %+v", testRule)
		}
	}
}
