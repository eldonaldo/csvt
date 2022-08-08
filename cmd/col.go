package cmd

import (
	"github.com/eldonaldo/csvt/internal"
	"github.com/spf13/cobra"
	"strings"
)

var (
	FlagColNames string
)

// colCmd represents the col command
var colCmd = &cobra.Command{
	Use:   "col",
	Short: "Filter operations on the columns",
	Long:  `Column filtering operations such as dropping columns etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		fields := strings.Split(FlagColNames, ",")
		csv := internal.ReadCSV()
		csv.FilterColumns(fields)
		csv.Print()
	},
}

func init() {
	filterCmd.AddCommand(colCmd)
	colCmd.Flags().StringVar(&FlagColNames, "keep", "", "Colum names to keep after filtering")
	_ = colCmd.MarkFlagRequired("keep")
}
