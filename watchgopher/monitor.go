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
	queue := make(chan string)

	monitor := &Monitor{dir, ticker, actions, queue}
	monitor.start()
}

type Monitor struct {
	dir     *Dir
	ticker  <-chan time.Time
	actions []Action
	queue   chan string
}

func (m *Monitor) start() {
	go m.workOff(m.queue)

	err := m.dir.StartWatching()
	if err != nil {
		panic(err)
	}

	for {
		select {
		case ev := <-m.dir.Events:
			m.queue <- ev.Name
		case <-m.ticker:
			m.dir.scan()
			for fpath, _ := range m.dir.Files {
				m.queue <- fpath
			}
		}
	}
}

func (m *Monitor) workOff(queue chan string) {
	for fpath := range queue {
		for _, action := range m.actions {
			action(fpath)
		}
	}
}
