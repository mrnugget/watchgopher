package main

import (
	"github.com/howeyc/fsnotify"
	"os/exec"
	"path"
)

func Manage(events chan *fsnotify.FileEvent, rules []*Rule) (queue chan *exec.Cmd) {
	queue = make(chan *exec.Cmd)

	go func() {
		for ev := range events {
			matches := matchingRules(rules, ev.Name)
			if len(matches) > 0 {
				for _, rule := range matches {
					cmd := exec.Command(rule.Run, getEventType(ev), ev.Name)
					queue <- cmd
				}
			}
		}
	}()

	return
}

func matchingRules(rules []*Rule, filepath string) (matches []*Rule) {
	matches = make([]*Rule, 0)

	dir, file := path.Split(filepath)
	dir = stripTrailingSlash(dir)

	for _, r := range rules {
		if r.Path == dir && (r.Pattern == "*" || path.Ext(r.Pattern) == path.Ext(file)) {
			matches = append(matches, r)
		}
	}
	return matches
}

func getEventType(ev *fsnotify.FileEvent) string {
	switch {
	case ev.IsCreate():
		return "CREATE"
	case ev.IsModify():
		return "MODIFY"
	case ev.IsDelete():
		return "DELETE"
	case ev.IsRename():
		return "RENAME"
	}
	return ""
}
