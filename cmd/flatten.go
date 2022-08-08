/*
Copyright © 2021 Nico Previtali <nico.previtali@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"github.com/eldonaldo/csvt/internal"

	"github.com/spf13/cobra"
)

var (
	FlagFlattenSeparator string
)

// flattenCmd represents the flatten command
var flattenCmd = &cobra.Command{
	Use:   "flatten",
	Short: "Flattens the given csv using the specified separator",
	Long:  `Flattens the given csv using the specified separator.`,
	Run: func(cmd *cobra.Command, args []string) {
		csv := internal.ReadCSV()
		flat, err := csv.Flatten(FlagFlattenSeparator)
		internal.CheckIfErrorAndPanic(err)
		fmt.Println(flat)
	},
}

func init() {
	rootCmd.AddCommand(flattenCmd)
	colCmd.Flags().StringVar(&FlagFlattenSeparator, "flatten-sep", ",", "Flatten separator")
}
