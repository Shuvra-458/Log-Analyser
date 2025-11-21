package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	filePath  string
	outFile   string
	outAppend string
)

var rootCmd = &cobra.Command{
	Use:   "golog",
	Short: "A simple log analyzer CLI tool",
	Long:  "Golog analyzes log files and gives basic statistics or filtered outputs.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use a subcommand like: golog stats --file <path>")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&filePath, "file", "f", "", "Path to log file")
	rootCmd.PersistentFlags().StringVar(&outFile, "out", "", "Path to output file (overwrite)")
	rootCmd.PersistentFlags().StringVar(&outAppend, "out-append", "", "Path to output file (append)")
}
