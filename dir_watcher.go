package main

import (
	"github.com/howeyc/fsnotify"
)

const fileChanBuf = 500

func WatchDirs(dirs []string) (d *DirWatcher, err error) {
	bufevents := make(chan *fsnotify.FileEvent, fileChanBuf)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	d = &DirWatcher{bufevents, watcher}

	go d.proxyEvents()

	for _, dir := range dirs {
		err = watcher.Watch(dir)
		if err != nil {
			return nil, err
		}
	}
	return
}

type DirWatcher struct {
	Events  chan *fsnotify.FileEvent
	watcher *fsnotify.Watcher
}

func (d *DirWatcher) Stop() (err error) {
	err = d.watcher.Close()
	return
}

func (d *DirWatcher) proxyEvents() {
	for {
		select {
		case ev := <-d.watcher.Event:
			d.Events <- ev
		case <-d.watcher.Error:
			break
		}
	}
}
