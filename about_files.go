package go_koans

import (
	"io/ioutil"
)
import "strings"

func aboutFiles() {
	filename := "about_files.go"
	contents, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(contents), "\n")
	assert(lines[5] == `import "strings"`) // handling files is too trivial
}
