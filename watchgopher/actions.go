package watchgopher

import (
	"fmt"
	"os"
	"regexp"
)

type Action func(string, os.FileInfo)

func Unzipper(path string, fi os.FileInfo) {
	ok, err := regexp.MatchString(`^.*\.zip$`, path)

	if err == nil && ok {
		fmt.Println("IT IS A ZIP!")
	}
}
