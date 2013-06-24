package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
)

const VERSION = "0.1"

var versionFlag = flag.Bool("version", false, "Print version of this build")

func usage() {
	fmt.Fprintf(os.Stderr, "Watchgopher %s - Listen to file changes and react to them\n", VERSION)
	fmt.Fprintf(os.Stderr, "\nUsage:\n")
	fmt.Fprintf(os.Stderr, "\t watchgopher [configuration file]\n")
	fmt.Fprintf(os.Stderr, "\nArguments:\n")
	fmt.Fprintf(os.Stderr, "\t --help or -h\t\tPrint help (this message) and exit\n")
	fmt.Fprintf(os.Stderr, "\t --version\t\tPrint version number and exit\n")
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if *versionFlag {
		fmt.Printf("Watchgopher %s\n", VERSION)
		os.Exit(0)
	}

	log.Printf("Starting watchgopher %s ...\n", VERSION)

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

	for payload := range queue {
		workOff(payload)
	}
}

func workOff(pl CmdPayload) {
	if pl.LogOutput {
		outp, err := pl.Cmd.StdoutPipe()
		if err != nil {
			logCmdErr(pl.Cmd, err)
		}

		errp, err := pl.Cmd.StderrPipe()
		if err != nil {
			logCmdErr(pl.Cmd, err)
		}

		_, filename := path.Split(pl.Cmd.Path)

		if outp != nil {
			go pipeToLog(filename, "STDOUT", outp)
		}

		if errp != nil {
			go pipeToLog(filename, "STDERR", errp)
		}
	}

	log.Printf("%s, ARGS: %s -- START\n", pl.Cmd.Path, pl.Cmd.Args[1:])

	if err := pl.Cmd.Start(); err != nil {
		logCmdErr(pl.Cmd, err)
		return
	}

	err := pl.Cmd.Wait()
	if err != nil {
		logCmdErr(pl.Cmd, err)
		return
	}
	log.Printf("%s, ARGS: %s -- SUCCESS\n", pl.Cmd.Path, pl.Cmd.Args[1:])
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
			break
		}
		log.Printf("[%s %s] %s", filename, prefix, line)
	}
}

func logCmdErr(cmd *exec.Cmd, err error) {
	log.Printf("%s, ARGS: %s -- ERROR: %s\n", cmd.Path, cmd.Args[1:], err)
}
