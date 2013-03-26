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

	// Rewrite file to trigger Modify event
	content, err := ioutil.ReadFile(sub1 + "/foobar.txt")
	checkErr(t, err)
	err = ioutil.WriteFile(sub1 + "/foobar.txt", content, 0644)
	checkErr(t, err)

	ev := <-watcher.Events
	if !ev.IsModify() && ev.Name != sub1 + "/foobar.txt" {
		t.Fatal("Wrong event")
	}

	// Create file to trigger create event
	err = ioutil.WriteFile(sub2 + "/hello.txt", []byte("Hello World!"), 0644)
	checkErr(t, err)

	ev = <-watcher.Events
	if !ev.IsCreate() && ev.Name != sub2 + "/hello.txt" {
		t.Fatal("Wrong event")
	}

	time.Sleep(1 * time.Millisecond)

	// Remove newly created file
	err = os.Remove(sub2 + "/hello.txt")
	checkErr(t, err)

	ev = <-watcher.Events
	if !ev.IsDelete() && ev.Name != sub2 + "/hello.txt" {
		t.Fatal("Wrong event")
	}

	watcher.Stop()
}

func checkErr(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
