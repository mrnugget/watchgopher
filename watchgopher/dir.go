package watchgopher

import (
	"github.com/howeyc/fsnotify"
	"io/ioutil"
	"os"
	"path"
)

const fileChanBuf = 500

func NewDir(path string) (d *Dir, err error) {
	files := make(map[string]os.FileInfo)
	events := make(chan *fsnotify.FileEvent, fileChanBuf)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	return &Dir{path, files, events, watcher}, nil
}

type Dir struct {
	Path    string
	Files   map[string]os.FileInfo
	Events  chan *fsnotify.FileEvent
	watcher *fsnotify.Watcher
}

func (d *Dir) Listen() (err error) {
	err = d.scan()
	if err != nil {
		return
	}

	err = d.watcher.Watch(d.Path)
	if err != nil {
		return
	}

	go d.proxyEvents()
	return nil
}

func (d *Dir) Stop() (err error) {
	err = d.watcher.Close()
	return
}

func (d *Dir) proxyEvents() {
	for {
		select {
		case ev := <-d.watcher.Event:
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
		case <-d.watcher.Error:
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
