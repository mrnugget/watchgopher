package main

import (
	"fmt"
	"github.com/howeyc/fsnotify"
	"path"
)

func Manage(events chan *fsnotify.FileEvent, rules []*Rule) {
	for ev := range events {
		fmt.Println("RECEIVED EVENT")

		dir, file := path.Split(ev.Name)
		dir = stripTrailingSlash(dir)

		rule := ruleWithPath(rules, dir)
		if rule != nil {
			fmt.Printf("RUN: %s, FILE ARG: %s\n", rule.Run, file)
		}
	}
}

func ruleWithPath(rules []*Rule, path string) (rule *Rule) {
	for _, rule := range rules {
		if rule.Path == path {
			return rule
		}
	}
	return nil
}
