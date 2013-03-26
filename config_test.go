package main

import (
	"testing"
)

var parseTests = []struct {
	index        int
	expectedPath string
	expectedRun  string
}{
	{
		0,
		"/tmp/foo",
		"/usr/local/bar/foobar",
	},
	{
		1,
		"/tmp/bar",
		"/usr/local/bar/barfoo",
	},
}

func TestParseConfig(t *testing.T) {
	rules, err := ParseConfig(fixtures + "/example_config.json")
	if err != nil {
		t.Fatal(err)
	}
	if len(rules) != 2 {
		t.Fatal("Did not include all rules")
	}

	for _, pt := range parseTests {
		actualPath := rules[pt.index].Path
		actualRun := rules[pt.index].Run

		if actualPath != pt.expectedPath && actualRun != pt.expectedRun {
			t.Errorf("ACTUALPATH %v - EXPECTEDPATH %v, ACTUALRUN %v - EXPECTEDRUN %v", actualPath, pt.expectedPath, actualRun, pt.expectedRun)
			continue
		}
	}
}
