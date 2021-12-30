package internal

import (
	"strings"
)

// CSVRow represents a column name to value map
type CSVRow map[string]string

// CSVObject is a list of csv rows
type CSVObject struct {
	hasHeader bool
	Header    []string
	Rows      []CSVRow
}

//func (c *CSVRow) Get(field string) (string, error) {
//	if !c.HasHeader() {
//		return "", errors.New("CSV has no headers")
//	}
//}

// String prints the csv object as
func (c *CSVObject) String() string {
	var stringBuilder strings.Builder
	if c.hasHeader {
		stringBuilder.WriteString(strings.Join(c.Header, FlagSeparator))
		stringBuilder.WriteString(CRLF)
	}

	for j, row := range c.Rows {
		r := make([]string, len(row))
		for i, idx := range c.Header {
			r[i] = row[idx]
		}
		stringBuilder.WriteString(strings.Join(r, FlagSeparator))

		if j != len(c.Rows)-1 {
			stringBuilder.WriteString(CRLF)
		}
	}

	return stringBuilder.String()
}
