package watchgopher

import (
	"fmt"
	"time"
)

func Watch(path string, interval time.Duration) {
	fmt.Println("PATH:", path)
	fmt.Println("INTERVAL:", interval)
	fmt.Println("Watchgopher is watching...")

	dir := NewDir(path)
	ticker := time.Tick(interval)
	actions := []Action{Unzipper}

	monitor := &Monitor{dir, ticker, actions}
	monitor.start()
}

type Monitor struct {
	dir     *Dir
	ticker  <-chan time.Time
	actions []Action
}

func (m *Monitor) start() {
	err := m.dir.StartWatching()
	if err != nil {
		panic(err)
	}

	for {
		select {
		case <-m.dir.Events:
			m.run()
		case <-m.ticker:
			m.run()
		}
	}
}

func (m *Monitor) run() {
	for fpath, finfo := range m.dir.Files {
		for _, action := range m.actions {
			action(fpath, finfo)
		}
	}
}
