package watchgopher

import (
	"fmt"
	"regexp"
)

type Action func(path string)

func Unzipper(path string) {
	ok, err := regexp.MatchString(`^.*\.zip$`, path)

	if err == nil && ok {
		fmt.Println("IT IS A ZIP!")
	}
}
