package cmd

import (
	"github.com/eldonaldo/csvt/internal"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "csvt",
	Short: "CSV transformation tool",
	Long:  `A cli to transform your csv files. For example, you might rename columns`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&internal.FlagNoHeader, "no-header", "n", false, "If set indicates that the file has no header")
	rootCmd.PersistentFlags().StringVarP(&internal.FlagSeparator, "sep", "s", ",", "Separator, default is ','")
	rootCmd.PersistentFlags().BoolVarP(&internal.FlagUseTestData, "test-data", "t", false, "If true, uses hard coded test data as csv input instead of stdin")
	rootCmd.PersistentFlags().BoolVarP(&internal.FlagDebug, "debug", "d", false, "Print debug information")
}
