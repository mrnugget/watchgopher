package main

import (
	"testing"
)

type expectedRule map[string]string

var expectedRules = map[string]expectedRule{
	"/tmp/foo": map[string]string{
		"run": "/usr/local/bar/foobar",
	},
	"/tmp/bar": map[string]string{
		"run": "/usr/local/bar/barfoo",
	},
}

func TestParseConfig(t *testing.T) {
	rules, err := ParseConfig(fixtures + "/example_config.json")
	if err != nil {
		t.Fatal(err)
	}

	for _, rule := range rules {
		expected, ok := expectedRules[rule.Path]
		if !ok {
			t.Error("Rule not found")
			continue
		}
		if rule.Run != expected["run"] {
			t.Error("Wrong Run. Expected: %v Actual: %v", expected["run"],
				rule.Run)
		}
	}
}
