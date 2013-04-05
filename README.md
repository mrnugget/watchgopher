# Watchgopher [![Build Status](https://travis-ci.org/mrnugget/watchgopher.png)](https://travis-ci.org/mrnugget/watchgopher)

Watchgopher listens to file events in directories and dispatches these events
(including event type and file) to commands of your choice to handle them.

## Getting Started

1. Make sure you have [Go](http://golang.org/) running on your system and setup.
   Then use the `go` command to install Watchgopher:

    `$ go get -u github.com/mrnugget/watchgopher`

2. Create a configuration file for Watchgopher to use. Modify this example to
   your needs and save it:

   ```
   {
      "/Users/mrnugget/Downloads": [
         {"run": "/Users/mrnugget/bin/unzipper.rb", "pattern": "*.zip"}
      ]
   }
   ```

3. Make sure the specified directory to watch exists and the command is
   executable. Then point watchgopher to your newly created configuration file:
  
    `$ watchgopher watchgopher_config.json`

Watchgopher is now keeping track of all files in the `/Users/mrnugget/Downloads`
directory. As soon as something happens to a file, whose name matches the
specified pattern, Watchgopher will pass the type of the event and the absolute
path to the file to the specified command.

Creating a new `*.zip` file in `/Users/mrnugget/Downloads` will follow in
Watchgopher calling this command: 

  `/Users/mrnugget/bin/unzipper.rb CREATE /Users/mrnugget/Downloads/new.zip`

## Usage

Watchgopher will pass two arguments to every specified command, should a file
event happen whose filename matches the specified pattern. Those two arguments
will be:

1. Type of the event (`CREATE`, `MODIFY`, `DELETE` or `RENAME`)
2. Absolute path of the file triggering the event

To properly use Watchgopher, your specified commands should take care of those
arguments and act accordingly. What those scripts will do is entirely up to
them.

## Configuration

The basic pattern of a Watchgopher configuration file is this:

```
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
