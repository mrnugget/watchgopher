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
			rule := matchingRule(rules, ev.Name)
			if rule != nil {
				cmd := exec.Command(rule.Run, getEventType(ev), ev.Name)
				queue <- cmd
			}
		}
	}()

	return
}

func matchingRule(rules []*Rule, filepath string) (rule *Rule) {
	dir, file := path.Split(filepath)
	dir = stripTrailingSlash(dir)

	for _, r := range rules {
		if r.Path == dir && (r.Pattern == "*" || path.Ext(r.Pattern) == path.Ext(file)) {
			return r
		}
	}
	return nil
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
