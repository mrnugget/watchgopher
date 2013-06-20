# Watchgopher [![Build Status](https://travis-ci.org/mrnugget/watchgopher.png)](https://travis-ci.org/mrnugget/watchgopher)

Watchgopher is a small, lightweight tool that listens to file events in
directories and dispatches these events (including event type and file) to
commands of your choice to handle them.

## Getting Started

1. Make sure you have Watchgopher installed (see the "Installation" section
   further down for instructions).

2. Create a configuration file for Watchgopher to use. Modify this example
   according to your needs and save it:

   ```
   {
      "/Users/mrnugget/Downloads": [
         {
           "run": "/Users/mrnugget/bin/unzipper.rb",
           "pattern": "*.zip",
           "log_output": true
         }
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

Check out the [examples](https://github.com/mrnugget/watchgopher/tree/master/examples)
directory for a config file and an example script to unzip new `*.zip` files in
a folder.

## Installation
### Binary Releases

Download one of the pre-built binary releases for Linux or OS X (darwin):

 * [watchgopher-0.1-darwin-amd64.tar.gz](https://s3-us-west-2.amazonaws.com/watchgopher/watchgopher-0.1-darwin-amd64.tar.gz)
 * [watchgopher-0.1-linux-amd64.tar.gz](https://s3-us-west-2.amazonaws.com/watchgopher/watchgopher-0.1-linux-amd64.tar.gz)

### Compiling From Source

Make sure you have [Go](http://golang.org/) running on your system and setup.
Then use the `go` command to install Watchgopher and its dependencies:

  `$ go get -u github.com/mrnugget/watchgopher`

## Usage

Watchgopher will pass two arguments to every specified command, should a file
event happen whose filename matches the specified pattern. Those two arguments
will be:

1. Type of the event (`CREATE`, `MODIFY`, `DELETE` or `RENAME`)
2. Absolute path of the file triggering the event

To properly use Watchgopher, your specified commands should take care of those
arguments and act accordingly. What those scripts will do is entirely up to
them. For a more thorough explanation read [this blog post](http://mrnugget.github.io/blog/2013/04/07/watchgopher/).

## Configuration

The basic, required pattern of a Watchgopher configuration file is this:

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

`[PATH OF DIRECTORY TO WATCH]` is the absolute path of a directory which
Watchgopher should keep track of.

**Every defined directory can have several rules**.

Every **rule** requires two attributes:

1. **"run"**: Absolute path to command which will handle the event. If the
   command is in your `$PATH` you won't need to provide the absolute path, just
   the name of the command.
2. **"pattern"**: Defines which pattern a filename of an event has to match in
   order to get dispatched to he **"run"** command. [See
   this](http://golang.org/pkg/path/filepath/#Match) for an explanation of possible
   patterns.

Every **rule** can use optional attributes:

1. **"log_output"**: Tells Watchgopher whether to log the output of the
   specified **"run"** command or not. If not specified, the default is
   **false**. If it's **true**, then Watchgopher will log the commands STDERR and
   STDOUT to its logoutput, prefixed with the commands filename.
2. **"change_pwd"**: If this is **true** (default, when unspecified, is **false**)
   Watchgopher changes the working directory of the executed command to the path
   of the directory to watch.

Whenever an event is triggered in a directory, watchgopher checks which rules
apply to this event (by checking against the `"pattern"`). If a rule applies,
because the defined pattern matches the file events filename, Watchgopher will
dispatch the event to the defined command (`"run"`).

## Thanks

Thanks to [howeyc](https://github.com/howeyc) for building the
[fsnotify](https://github.com/howeyc/fsnotify) package.

## License

MIT, see [LICENSE](LICENSE)
