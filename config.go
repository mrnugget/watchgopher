package main

import (
	"encoding/json"
	"io/ioutil"
)

var f interface{}

type Rule struct {
	Path string
	Run string
}

func ParseConfig(path string) (rules []*Rule, err error) {
	rules = make([]*Rule, 0)

	c, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(c, &f)
	if err != nil {
		return nil, err
	}

	config := f.(map[string]interface{})
	for path, v := range config {
		attributes := v.(map[string]interface{})
		run := attributes["run"].(string)

		rules = append(rules, &Rule{path, run})
	}

	return rules, nil
}
