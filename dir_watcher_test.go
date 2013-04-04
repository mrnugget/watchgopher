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

	// Change content of file to trigger Modify event
	f, err := os.OpenFile(sub1+"/foobar.txt", os.O_WRONLY, 0666)
	checkErr(t, err)
	f.Sync()
	time.Sleep(time.Millisecond)
	f.WriteString("Hello")
	f.Sync()
	f.Close()

	time.Sleep(100 * time.Millisecond)

	ev := <-watcher.Events
	if !ev.IsModify() && ev.Name != sub1+"/foobar.txt" {
		t.Fatalf("Wrong event: %s", ev)
	}

	time.Sleep(100 * time.Millisecond)

	// Create file to trigger create event
	err = ioutil.WriteFile(sub2+"/hello.txt", []byte("Hello World!"), 0644)
	checkErr(t, err)

	ev = <-watcher.Events
	if !ev.IsCreate() && ev.Name != sub2+"/hello.txt" {
		t.Fatalf("Wrong event: %s", ev)
	}

	time.Sleep(100 * time.Millisecond)

	// Remove newly created file
	err = os.Remove(sub2 + "/hello.txt")
	checkErr(t, err)

	time.Sleep(1 * time.Millisecond)

	ev = <-watcher.Events
	if !ev.IsDelete() && ev.Name != sub2+"/hello.txt" {
		t.Fatalf("Wrong event: %s", ev)
	}

	watcher.Stop()
}

func checkErr(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
