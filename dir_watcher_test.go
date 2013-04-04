package main

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestEvents(t *testing.T) {
	sub1 := fixtures + "/sub1"
	sub2 := fixtures + "/sub2"

	dirs := []string{sub1, sub2}

	watcher, err := WatchDirs(dirs)
	checkErr(t, err)
	defer func() {
		err = watcher.Stop()
		if err != nil {
			t.Fatalf("watcher.Stop() error: %s", err)
		}
	}()

	// Rewrite file to trigger Modify event
	content, err := ioutil.ReadFile(sub1 + "/foobar.txt")
	checkErr(t, err)
	err = ioutil.WriteFile(sub1+"/foobar.txt", content, 0644)
	checkErr(t, err)

	// Create file to trigger create event
	err = ioutil.WriteFile(sub2+"/hello.txt", []byte("Hello World!"), 0644)
	checkErr(t, err)

	// Wait for filesystem to sync changes
	time.Sleep(20 * time.Millisecond)

	// Remove newly created file
	err = os.Remove(sub2 + "/hello.txt")
	checkErr(t, err)


	// Wait again to make sure all triggered events are queued up
	time.Sleep(20 * time.Millisecond)
	eventslen := len(watcher.Events)
	if eventslen != 3 {
		t.Fatalf("len(watcher.Events) = %s, wanted %s", eventslen, 3)
	}

	ev := <-watcher.Events
	if !ev.IsModify() && ev.Name != sub1+"/foobar.txt" {
		t.Fatalf("Wrong event: %s", ev)
	}

	ev = <-watcher.Events
	if !ev.IsCreate() && ev.Name != sub2+"/hello.txt" {
		t.Fatalf("Wrong event: %s", ev)
	}

	ev = <-watcher.Events
	if !ev.IsDelete() && ev.Name != sub2+"/hello.txt" {
		t.Fatalf("Wrong event: %s", ev)
	}
}

func checkErr(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
