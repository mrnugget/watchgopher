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
			rule := ruleForEvent(rules, ev)
			if rule != nil {
				cmd := exec.Command(rule.Run, getEventType(ev), ev.Name)
				queue <- cmd
			}
		}
	}()

	return
}

func ruleForEvent(rules []*Rule, ev *fsnotify.FileEvent) (rule *Rule) {
	path, _ := path.Split(ev.Name)
	path = stripTrailingSlash(path)

	for _, rule := range rules {
		if rule.Path == path {
			return rule
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
