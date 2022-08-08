package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// filterCmd represents the filter command
var filterCmd = &cobra.Command{
	Use:   "filter",
	Short: "Filters the CSV according to columns names.",
	Long:  `Filters the CSV according to columns names. The filtered colums are retrained, all others are dropped.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("filter has no root command, call with --help")
	},
}

func init() {
	rootCmd.AddCommand(filterCmd)
}
