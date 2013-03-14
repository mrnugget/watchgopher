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

	monitor := &Monitor{dir, ticker}
	monitor.start()
}

type Monitor struct {
	dir    *Dir
	ticker <-chan time.Time
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
	fmt.Println(m.dir.Files)
}
