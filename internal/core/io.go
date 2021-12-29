package core

import (
	"encoding/csv"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// CSVRow represents a column name to value map
type CSVRow map[string]string

// CSVObject is a list of csv rows
type CSVObject []CSVRow

// ReadFromStdIn reads piped data from std in
func ReadFromStdIn() string {
	data, err := ioutil.ReadAll(os.Stdin)
	CheckIfErrorAndPanic(err)
	return string(data)
}

// ParseCSV parses the given string as CSV object
func ParseCSV(data string, hasHeader bool) CSVObject {
	stringReader := strings.NewReader(data)
	csvReader := csv.NewReader(stringReader)
	headerMap := make(map[int]string)

	// Parse header, if any
	if hasHeader {
		row, err := csvReader.Read()
		CheckIfErrorAndPanic(err)

		for i, name := range row {
			headerMap[i] = name
		}
	}

	// Parse data
	csvRows := make(CSVObject, 0)
	for {
		row, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
		}

		csvRow := make(CSVRow)
		for i, value := range row {
			key := strconv.Itoa(i)
			if hasHeader {
				key = headerMap[i]
			}

			csvRow[key] = value
		}

		csvRows = append(csvRows, csvRow)
	}

	return csvRows
}

// CSVSliceToString converts a csv [][]string to a string
func CSVSliceToString(slice [][]string) string {
	var stringBuilder strings.Builder
	for j, row := range slice {
		for i, value := range row {
			if i != 0 {
				stringBuilder.WriteString(",")
			}
			stringBuilder.WriteString(value)
		}

		if j != len(slice)-1 {
			stringBuilder.WriteString("\n")
		}
	}

	return stringBuilder.String()
}
