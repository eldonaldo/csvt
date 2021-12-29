package core

import (
	"fmt"
)

// CheckIfErrorAndPanic should be used to naively panics if an error is not nil.
func CheckIfErrorAndPanic(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	panic(err)
}

// Info should be used to describe the example commands that are about to run.
func Info(format string, args ...interface{}) {
	fmt.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

// Warning should be used to display a warning
func Warning(format string, args ...interface{}) {
	fmt.Printf("\x1b[36;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

// Error should be used to display an error
func Error(format string, args ...interface{}) {
	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}
