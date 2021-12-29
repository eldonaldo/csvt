package internal

import (
	"io/ioutil"
	"os"
)

// ReadFromStdIn reads piped data from std in
func ReadFromStdIn() string {
	data, err := ioutil.ReadAll(os.Stdin)
	CheckIfErrorAndPanic(err)
	return string(data)
}
