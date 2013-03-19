package main

import (
	"flag"
	"fmt"
	"github.com/mrnugget/watchgopher/watchgopher"
	"os"
	"time"
)

var interval = flag.Duration("interval", 60*time.Second, "when to check")
var path string

func main() {
	flag.Parse()

	path = flag.Arg(0)
	if path == "" {
		fmt.Println("You have to provide a path to watch")
		os.Exit(1)
	}
	if path == "." {
		path, _ = os.Getwd()
	}

	err := watchgopher.Watch(path, *interval)
	if err != nil {
		panic(err)
	}
}
