package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"
	"path"
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

		if err = cmd.Start(); err != nil {
			log.Printf("%s, ARGS: %s -- ERROR: %s\n", cmd.Path, cmd.Args[1:], err)
			continue
		}

		// @TODO: This writes only after the pipes are closed. Would be better
		// to stream here.
		go pipeToLog(filename, "STDOUT", outp)
		go pipeToLog(filename, "STDERR", errp)

		err = cmd.Wait()
		if err != nil {
			log.Printf("%s, ARGS: %s -- ERROR: %s\n", cmd.Path, cmd.Args[1:], err)
			continue
		}
		log.Printf("%s, ARGS: %s -- SUCCESS\n", cmd.Path, cmd.Args[1:])
	}
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
