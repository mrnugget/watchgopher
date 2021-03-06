package main

import (
	"encoding/json"
	"io/ioutil"
)

type Rule struct {
	Path      string
	Run       string
	Pattern   string
	LogOutput bool
	ChangePwd bool
}

func ParseConfig(configpath string) (rules []*Rule, err error) {
	rules = make([]*Rule, 0)

	c, err := ioutil.ReadFile(configpath)
	if err != nil {
		return nil, err
	}

	var f interface{}
	err = json.Unmarshal(c, &f)
	if err != nil {
		return nil, err
	}

	paths := f.(map[string]interface{})
	for path, pathRules := range paths {
		for _, ruleAttributes := range pathRules.([]interface{}) {
			attributes := ruleAttributes.(map[string]interface{})
			rules = append(rules, attributesToRule(path, attributes))
		}
	}

	return rules, nil
}

func attributesToRule(p string, attr map[string]interface{}) *Rule {
	run := attr["run"].(string)
	pattern := attr["pattern"].(string)

	log := false
	if attr["log_output"] != nil {
		log = attr["log_output"].(bool)
	}

	changePwd := false
	if attr["change_pwd"] != nil {
		changePwd = attr["change_pwd"].(bool)
	}

	return &Rule{stripTrailingSlash(p), run, pattern, log, changePwd}
}

func stripTrailingSlash(path string) string {
	for len(path) > 0 && path[len(path)-1] == '/' {
		path = path[0 : len(path)-1]
	}
	return path
}
