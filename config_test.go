package main

import (
	"testing"
)

func TestParseConfig(t *testing.T) {
	_, err := ParseConfig(fixtures + "/example_config.json")
	if err != nil {
		t.Fatal(err)
	}
}
