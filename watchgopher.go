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

	defer func() {
		err = watcher.Stop()
		if err != nil {
			fmt.Println("Error: Did not shut down cleanly")
			os.Exit(1)
		}
	}()

	queue := Manage(watcher.Events, rules)

	for cmd := range queue {
		cmd.Stdout = os.Stdout
		fmt.Printf("NOW RUNNING: %s, ARGS: %s\n", cmd.Path, cmd.Args[1:])
		err = cmd.Run()
		if err != nil {
			fmt.Println("ERROR:", err)
			continue
		}
	}
}
