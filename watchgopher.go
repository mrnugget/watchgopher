package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	flag.Parse()
	configPath := flag.Arg(0)
	_, err := os.Stat(configPath)
	if err != nil && os.IsNotExist(err) {
		log.Fatal("ERROR: Could not find configuration file")
	}

	rules, err := ParseConfig(configPath)
	if err != nil {
		log.Fatalf("ERROR: Could not parse config file: %s", err)
	}
	log.Printf("Successfully loaded configuration file. Number of rules: %d", len(rules))

	dirs := make([]string, 0)
	for _, rule := range rules {
		dirs = append(dirs, rule.Path)
	}

	watcher, err := WatchDirs(dirs)
	if err != nil {
		log.Println("ERROR: Could not start watching directories")
		log.Fatal(err)
	}

	defer func() {
		err = watcher.Stop()
		if err != nil {
			log.Fatal("ERROR: Did not shut down cleanly")
		}
	}()

	queue := Manage(watcher.Events, rules)

	log.Println("Watchgopher is now ready process file events")

	for cmd := range queue {
		err = cmd.Run()
		if err != nil {
			log.Printf("%s, ARGS: %s, -- ERROR: %s\n", cmd.Path, cmd.Args[1:], err)
			continue
		}
		log.Printf("%s, ARGS: %s, -- SUCCESS\n", cmd.Path, cmd.Args[1:])
	}
}
