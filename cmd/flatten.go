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
