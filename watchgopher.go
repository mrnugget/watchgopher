package main

func main() {
	// @TODO: Read configuration file either from `~/.watchgopher` or from
	// first command line argument

	// @TODO: Parse configuration file to get which directories to watch,
	// which pattern to match for which directory, which scripts to run on
	// event

	// @TODO: Watch directories for events (see `dir_watcher.go`) and pass
	// events to a manager, which checks for appliance of configuration

	// @TODO: If filename matches a pattern (e.g. `*.jpg`), pass it to a worker,
	// that shells out and runs configured command with two arguments:
	// `~/bin/script EVENTTYPE FILENAME`, where EVENTTYPE can be CREATE, DELETE,
	// MODIFY, RENAME and FILENAME is the absolute path to the file which
	// triggered the event
}
