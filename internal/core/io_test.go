package core

import (
	"github.com/stretchr/testify/assert"
	"os"
	"reflect"
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

func TestParseCSV(t *testing.T) {
	csvActual := ParseCSV("a,b,c\n1,2,3", true)
	csvExpected := make(CSVObject, 0)

	csvRow1 := make(CSVRow)
	csvRow1["a"] = "1"
	csvRow1["b"] = "2"
	csvRow1["c"] = "3"
	csvExpected = append(csvExpected, csvRow1)

	assert.Equal(t, true, reflect.DeepEqual(csvActual, csvExpected))
}
