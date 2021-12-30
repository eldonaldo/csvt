package internal

import (
	"encoding/csv"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// ReadFromStdIn reads piped data from std in
func ReadFromStdIn() string {
	data, err := ioutil.ReadAll(os.Stdin)
	CheckIfErrorAndPanic(err)
	return string(data)
}

// ParseCSV parses the given string as CSV object
func ParseCSV(data string, hasHeader bool) *CSVObject {
	stringReader := strings.NewReader(data)
	csvReader := csv.NewReader(stringReader)
	header := make([]string, 0)
	csvObject := &CSVObject{}

	// Parse header, if any
	if hasHeader {
		row, err := csvReader.Read()
		CheckIfErrorAndPanic(err)

		for _, name := range row {
			header = append(header, name)
		}
	}

	// Parse data
	csvRows := make([]CSVRow, 0)
	orderRead := false

	for {
		row, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
		}

		csvRow := make(CSVField)
		for i, value := range row {
			strIdx := strconv.Itoa(i)

			if !orderRead && !hasHeader {
				header = append(header, strIdx)
			}

			key := strIdx
			if hasHeader {
				key = header[i]
			}

			csvRow[key] = value
		}

		csvRows = append(csvRows, CSVRow{fields: csvRow, csvObject: csvObject})
		orderRead = true
	}

	csvObject.Rows = csvRows
	csvObject.Headers = header
	csvObject.hasHeader = hasHeader
	return csvObject
}

// CSVSliceToString converts a csv [][]string to a string
func CSVSliceToString(slice [][]string) string {
	var stringBuilder strings.Builder
	for j, row := range slice {
		for i, value := range row {
			if i != 0 {
				stringBuilder.WriteString(FlagSeparator)
			}
			stringBuilder.WriteString(value)
		}

		if j != len(slice)-1 {
			stringBuilder.WriteString("\n")
		}
	}

	return stringBuilder.String()
}

// ReadCSV reads the csv from stdin or returns test data if the test flag is set
func ReadCSV() *CSVObject {
	hasHeader := !FlagNoHeader
	if FlagUseTestData {
		return ParseCSV("col1,col2,col3\nval1,val2,val3\nval4,val5,val6", hasHeader)
	} else {
		return ParseCSV(ReadFromStdIn(), hasHeader)
	}
}
