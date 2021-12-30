package internal

import (
	"errors"
	"fmt"
	"strings"
)

// CSVField represents a column name to value map
type CSVField = map[string]string

// CSVRow is a column name to value map and a pointer to the csv object
type CSVRow struct {
	fields    CSVField
	csvObject *CSVObject
}

// CSVObject is a list of csv rows
type CSVObject struct {
	hasHeader bool
	Headers   []string
	Rows      []CSVRow
}

// Get returns the value of the given field
func (c *CSVRow) Get(field string) (string, error) {
	if !c.csvObject.hasHeader {
		return "", errors.New("CSV has no headers")
	}

	value, ok := c.fields[field]
	if !ok {
		return "", errors.New("CSV has no field " + field)
	}

	return value, nil
}

// String converts the csv object to string representation
func (c *CSVObject) String() string {
	var stringBuilder strings.Builder
	if c.hasHeader {
		stringBuilder.WriteString(strings.Join(c.Headers, FlagSeparator))
		stringBuilder.WriteString(CRLF)
	}

	for j, row := range c.Rows {
		r := make([]string, len(row.fields))
		for i, idx := range c.Headers {
			r[i] = row.fields[idx]
		}
		stringBuilder.WriteString(strings.Join(r, FlagSeparator))

		if j != len(c.Rows)-1 {
			stringBuilder.WriteString(CRLF)
		}
	}

	if FlagDebug {
		stringBuilder.WriteString(CRLF)
		stringBuilder.WriteString(CRLF)
		stringBuilder.WriteString("> DEBUG:")
		stringBuilder.WriteString(CRLF)
		stringBuilder.WriteString(fmt.Sprintf(">   Has Header: %t", c.hasHeader))
		stringBuilder.WriteString(CRLF)
		stringBuilder.WriteString(fmt.Sprintf(">   Headers: %s", strings.Join(c.Headers, ",")))
	}

	return stringBuilder.String()
}

// Print prints the csv to stdout
func (c *CSVObject) Print() {
	fmt.Println(c.String())
}

// FilterColumns retains only the given fields
func (c *CSVObject) FilterColumns(fields []string) *CSVObject {
	remove := Difference(c.Headers, fields)
	for _, row := range c.Rows {
		for _, f := range remove {
			delete(row.fields, f)
		}
	}

	newHeaders := make([]string, 0)
	for _, val := range c.Headers {
		retain := false
		for _, f := range fields {
			if val == f {
				retain = true
				break
			}
		}

		if retain {
			newHeaders = append(newHeaders, val)
		}
	}

	c.Headers = newHeaders
	return c
}
