package main

import (
	"encoding/json"
	"io/ioutil"
)

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

	var f interface{}
	err = json.Unmarshal(c, &f)
	if err != nil {
		return nil, err
	}

	config := f.(map[string]interface{})
	for path, v := range config {
		attributes := v.(map[string]interface{})
		run := attributes["run"].(string)

		for len(path) > 0 && path[len(path)-1] == '/' {
			path = path[0 : len(path)-1]
		}

		rules = append(rules, &Rule{path, run})
	}

	return rules, nil
}
