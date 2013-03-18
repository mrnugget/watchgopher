package watchgopher

import (
	"github.com/howeyc/fsnotify"
	"io/ioutil"
	"os"
	"path"
)

func NewDir(path string) (d *Dir) {
	files := make(map[string]os.FileInfo)
	events := make(chan *fsnotify.FileEvent)

	return &Dir{path, files, events}
}

type Dir struct {
	Path   string
	Files  map[string]os.FileInfo
	Events chan *fsnotify.FileEvent
}

func (d *Dir) Listen() (err error) {
	err = d.scan()
	if err != nil {
		return
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return
	}

	err = watcher.Watch(d.Path)
	if err != nil {
		return
	}

	go d.handleEvents(watcher.Event, watcher.Error)
	return nil
}

func (d *Dir) handleEvents(events chan *fsnotify.FileEvent, errs chan error) {
	for {
		select {
		case ev := <-events:
			if ev.IsRename() || ev.IsDelete() {
				delete(d.Files, ev.Name)
			}
			if ev.IsCreate() || ev.IsModify() {
				fi, err := os.Stat(ev.Name)
				if err != nil {
					break
				}
				d.Files[ev.Name] = fi
			}
			d.Events <- ev
		case <-errs:
			break
		}
	}
}

func (d *Dir) scan() (err error) {
	files, err := ioutil.ReadDir(d.Path)
	if err != nil {
		return
	}

	for _, f := range files {
		d.Files[path.Join(d.Path, f.Name())] = f
	}

	return nil
}
