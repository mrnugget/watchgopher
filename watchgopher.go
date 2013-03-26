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

	// @TODO: Parse configuration file to get which directories to watch,
	// which pattern to match for which directory, which scripts to run on
	// event
	_, err = ParseConfig(configPath)
	if err != nil {
		fmt.Println("Error: Could not parse config file")
	}

	// @TODO: Watch directories for events (see `dir_watcher.go`) and pass
	// events to a manager, which checks for appliance of configuration

	// @TODO: If filename matches a pattern (e.g. `*.jpg`), pass it to a worker,
	// that shells out and runs configured command with two arguments:
	// `~/bin/script EVENTTYPE FILENAME`, where EVENTTYPE can be CREATE, DELETE,
	// MODIFY, RENAME and FILENAME is the absolute path to the file which
	// triggered the event
}
