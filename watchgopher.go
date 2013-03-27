package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	configPath := flag.Arg(0)
	_, err := os.Stat(configPath)
	if err != nil && os.IsNotExist(err) {
		fmt.Println("Please pass the path to the config file to watchgopher")
		os.Exit(1)
	}

	rules, err := ParseConfig(configPath)
	if err != nil {
		fmt.Println("Error: Could not parse config file")
		os.Exit(1)
	}

	dirs := make([]string, 0)
	for _, rule := range rules {
		dirs = append(dirs, rule.Path)
	}

	watcher, err := WatchDirs(dirs)
	if err != nil {
		fmt.Println("Error: Could not start watching directories")
		fmt.Println(err)
		os.Exit(1)
	}

	// @TODO: Manage the events, check if a rule applies (file is in path),
	// then run script with arguments
	for ev := range watcher.Events {
		fmt.Println("EVENT - File:", ev.Name)
	}

	// @TODO: If filename matches a pattern (e.g. `*.jpg`), pass it to a worker,
	// that shells out and runs configured command with two arguments:
	// `~/bin/script EVENTTYPE FILENAME`, where EVENTTYPE can be CREATE, DELETE,
	// MODIFY, RENAME and FILENAME is the absolute path to the file which
	// triggered the event
}
