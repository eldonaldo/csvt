package internal

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestReadFromStdIn(t *testing.T) {
	input := []byte("Alice")
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	_, err = w.Write(input)
	if err != nil {
		t.Error(err)
	}
	w.Close()

	stdin := os.Stdin
	// Restore stdin right after the test.
	defer func() { os.Stdin = stdin }()
	os.Stdin = r

	d := ReadFromStdIn()
	assert.Equal(t, "Alice", d)
}
