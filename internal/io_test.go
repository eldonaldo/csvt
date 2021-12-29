package internal

import (
	"encoding/csv"
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

func TestParseCSV_SingleRow(t *testing.T) {
	csvActual := ParseCSV("a,b,c\n1,2,3", true)
	csvExpected := make(CSVObject, 0)

	csvRow1 := make(CSVRow)
	csvRow1["a"] = "1"
	csvRow1["b"] = "2"
	csvRow1["c"] = "3"
	csvExpected = append(csvExpected, csvRow1)

	assert.Equal(t, true, reflect.DeepEqual(csvActual, csvExpected))
}

func TestParseCSV_MultipleRows(t *testing.T) {
	csvActual := ParseCSV("a,b,c\n1,2,3\n4,5,6", true)
	csvExpected := make(CSVObject, 0)

	csvRow1 := make(CSVRow)
	csvRow1["a"] = "1"
	csvRow1["b"] = "2"
	csvRow1["c"] = "3"
	csvExpected = append(csvExpected, csvRow1)

	csvRow2 := make(CSVRow)
	csvRow2["a"] = "4"
	csvRow2["b"] = "5"
	csvRow2["c"] = "6"
	csvExpected = append(csvExpected, csvRow2)

	assert.Equal(t, true, reflect.DeepEqual(csvActual, csvExpected))
}

func TestParseCSV_WithoutHeaders(t *testing.T) {
	csvActual := ParseCSV("1,2,3\n4,5,6", false)
	csvExpected := make(CSVObject, 0)

	csvRow1 := make(CSVRow)
	csvRow1["0"] = "1"
	csvRow1["1"] = "2"
	csvRow1["2"] = "3"
	csvExpected = append(csvExpected, csvRow1)

	csvRow2 := make(CSVRow)
	csvRow2["0"] = "4"
	csvRow2["1"] = "5"
	csvRow2["2"] = "6"
	csvExpected = append(csvExpected, csvRow2)

	assert.Equal(t, true, reflect.DeepEqual(csvActual, csvExpected))
}

func TestParseCSV_FromFile(t *testing.T) {
	csvFile, _ := os.Open("io_testdata/csv1.csv")
	r := csv.NewReader(csvFile)
	data, _ := r.ReadAll()

	csvActual := ParseCSV(CSVSliceToString(data), true)
	csvExpected := make(CSVObject, 0)

	csvRow1 := make(CSVRow)
	csvRow1["a"] = "1"
	csvRow1["b"] = "2"
	csvRow1["c"] = "3"
	csvExpected = append(csvExpected, csvRow1)

	csvRow2 := make(CSVRow)
	csvRow2["a"] = "4"
	csvRow2["b"] = "5"
	csvRow2["c"] = "6"
	csvExpected = append(csvExpected, csvRow2)

	assert.Equal(t, csvActual, csvExpected)
}

func TestCSVSliceToString(t *testing.T) {
	csvFile, _ := os.Open("io_testdata/csv1.csv")
	r := csv.NewReader(csvFile)
	data, _ := r.ReadAll()
	expected := "a,b,c\n1,2,3\n4,5,6"

	assert.Equal(t, expected, CSVSliceToString(data))
}
