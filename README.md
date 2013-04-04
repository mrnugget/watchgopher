# Watchgopher [![Build Status](https://travis-ci.org/mrnugget/watchgopher.png)]

Watchgopher listens to file events in directories and dispatches these events
(including event type and file) to commands of your choice to handle them.

## Installation

Make sure you have [Go](http://golang.org/) running on your system and setup.
Then use the `go` command to install Watchgopher:

    go get -u github.com/mrnugget/watchgopher

The `watchgopher` command should now be available to you.

## Usage

Before starting Watchgopher you need a configuration file to use. Modify this
example to your needs:

```json
{
  "/Users/mrnugget/tmp": [
    {"run": "/Users/mrnugget/bin/ruby_run.rb", "pattern": "*"},
    {"run": "/Users/mrnugget/bin/shell_run.sh", "pattern": "*.zip"}
  ],
  "/Users/mrnugget/Downloads": [
    {"run": "/Users/mrnugget/bin/gif_archiver", "pattern": "*.gif"}
  ]
}
```

Now point the `watchgopher` executable to this configuration file
and see it load up:

```bash
$ watchgopher tmp/watchgopher.json
2013/04/03 08:54:29 Successfully loaded configuration file. Number of rules: 3
2013/04/03 08:54:29 Watchgopher is now ready process file events
```

Watchgopher is now watching over those two directories and dispatch events if
anything happens in them!

Whenever a file event is triggered in any of the watched directories,
Watchgopher will dispatch this event to the defined commands (`"run"`) with **two
arguments**:

1. Type of the event (`CREATE`, `MODIFY`, `DELETE` or `RENAME`)
2. Absolute path of the file triggering the event

Watchgopher will output the following after successfully dispatching an event:

    2013/04/03 08:58:11 /Users/mrnugget/bin/gif_archiver, ARGS: [CREATE /Users/mrnugget/Downloads/otter.gif] -- SUCCESS

Well done!

## Configuration

The basic pattern of a Watchgopher configuration file is this:

```json
{
  "[PATH OF DIRECTORY TO WATCH]": [
    {"run": "[PATH OF COMMAND HANDLING THE EVENT]", "pattern": "[FILE NAME PATTERN]"},
    {"run": "[PATH OF COMMAND HANDLING THE EVENT]", "pattern": "[FILE NAME PATTERN]"}
  ],
  "[PATH OF DIRECTORY TO WATCH]": [
    {"run": "[PATH OF COMMAND HANDLING THE EVENT]", "pattern": "[FILE NAME PATTERN]"},
    {"run": "[PATH OF COMMAND HANDLING THE EVENT]", "pattern": "[FILE NAME PATTERN]"},
    ...
  ],
  ...
}
```

Basic explanation:

1. `[PATH OF DIRECTORY TO WATCH]`: Absolute path to the directory to watch
2. `[PATH OF COMMAND HANDLING THE EVENT]`: Absolute path to command
3. `[FILE NAME PATTERN]`: Only dispatch events to the defined command if the
   file name matches this pattern. [See this](http://golang.org/pkg/path/filepath/#Match) for an explanation of possible
   patterns.

**Every watched directory can have several rules (`{"run":"[...]", "pattern":"[...]"}`)
defined.**

Whenever an event is triggered in a directory, watchgopher checks which rules
apply to this event (by checking against the `"pattern"`). If a rule applies,
because the defined pattern matches the file events absolute filename,
Watchgopher will dispatch the event to the defined command (`"run"`).

## Thanks

Thanks to [howeyc](https://github.com/howeyc) for building the
[fsnotify](https://github.com/howeyc/fsnotify) package.

## License

MIT, see [LICENSE](LICENSE)
