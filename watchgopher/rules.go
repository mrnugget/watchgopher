package watchgopher

import (
	"os"
)

type Rule func(string, os.FileInfo)
