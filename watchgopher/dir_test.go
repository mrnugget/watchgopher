package watchgopher

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"
)

var (
	_, filename, _, _ = runtime.Caller(0)
	fixtures          = filepath.Dir(filename) + "/fixtures"
)

func TestFiles(t *testing.T) {
	dir, err := NewDir(fixtures)
	checkErr(t, err)

	err = dir.Listen()
	checkErr(t, err)
	defer dir.Stop()

	_, ok := dir.Files[fixtures+"/foobar.txt"]
	if !ok {
		t.Errorf("dir.Files does not include right files")
	}

	_, ok = dir.Files[fixtures+"/hello_world.txt"]
	if !ok {
		t.Errorf("dir.Files does not include right files")
	}
}

func TestEvents(t *testing.T) {
	testfilepath := fixtures + "/testfile.txt"

	dir, err := NewDir(fixtures)
	checkErr(t, err)

	err = dir.Listen()
	checkErr(t, err)
	defer dir.Stop()

	err = ioutil.WriteFile(testfilepath, []byte("Hello World!"), 0644)
	checkErr(t, err)

	time.Sleep(1 * time.Millisecond)

	_, ok := dir.Files[testfilepath]
	if !ok {
		t.Errorf("Did not add the created file to dir.Files")
	}

	err = os.Remove(testfilepath)
	checkErr(t, err)

	createev := <-dir.Events
	if createev.Name != testfilepath && !createev.IsCreate() {
		t.Fatal("Did not receive the right event")
	}

	deleteev := <-dir.Events
	if deleteev.Name != testfilepath && !deleteev.IsDelete() {
		t.Fatal("Did not receive the right event")
	}

	file, ok := dir.Files[testfilepath]
	if file != nil && ok {
		t.Errorf("Did not remove the deleted file from dir.Files")
	}
}

func checkErr(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
