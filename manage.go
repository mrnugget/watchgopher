package main

import (
	"bufio"
	"github.com/howeyc/fsnotify"
	"io"
	"log"
	"os/exec"
	"path"
	"path/filepath"
)

func Manage(events chan *fsnotify.FileEvent, rules []*Rule) (queue chan *exec.Cmd) {
	queue = make(chan *exec.Cmd)

	go func() {
		for ev := range events {
			matches := matchingRules(rules, ev.Name)
			if len(matches) > 0 {
				for _, rule := range matches {
					cmd := exec.Command(rule.Run, getEventType(ev), ev.Name)

					outp, err := cmd.StdoutPipe()
					if err != nil {
						log.Printf("%s, ARGS: %s -- ERROR: %s\n", cmd.Path, cmd.Args[1:], err)
						continue
					}

					errp, err := cmd.StderrPipe()
					if err != nil {
						log.Printf("%s, ARGS: %s -- ERROR: %s\n", cmd.Path, cmd.Args[1:], err)
						continue
					}

					_, filename := path.Split(cmd.Path)

					go pipeToLog(filename, "STDOUT", outp)
					go pipeToLog(filename, "STDERR", errp)

					queue <- cmd
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
			if r.Pattern == "*"  {
				matches = append(matches, r)
				continue
			}

			match, err := filepath.Match(r.Pattern, file)
			if r.Path == dir && match && err == nil {
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

func pipeToLog(filename, prefix string, pipe io.ReadCloser) {
	reader := bufio.NewReader(pipe)
	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("[%s %s] Reading Error: %s", filename, prefix, err)
		}
		log.Printf("[%s %s] %s", filename, prefix, line)
	}
}
