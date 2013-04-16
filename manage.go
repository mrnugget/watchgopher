package main

import (
	"github.com/howeyc/fsnotify"
	"os/exec"
	"path"
	"path/filepath"
)

type CmdPayload struct {
	Cmd       *exec.Cmd
	LogOutput bool
}

func Manage(events chan *fsnotify.FileEvent, rules []*Rule) (queue chan CmdPayload) {
	queue = make(chan CmdPayload)

	go func() {
		for ev := range events {
			matches := matchingRules(rules, ev.Name)
			if len(matches) > 0 {
				for _, rule := range matches {
					cmd := exec.Command(rule.Run, getEventType(ev), ev.Name)
					payload := CmdPayload{cmd, rule.LogOutput}

					queue <- payload
				}
			}
		}
	}()

	return
}

func matchingRules(rules []*Rule, filename string) (matches []*Rule) {
	matches = make([]*Rule, 0)

	dir, file := path.Split(filename)
	dir = stripTrailingSlash(dir)

	for _, r := range rules {
		if r.Path == dir {
			if r.Pattern == "*" {
				matches = append(matches, r)
				continue
			}

			match, err := filepath.Match(r.Pattern, file)
			if match && err == nil {
				matches = append(matches, r)
			}
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
