package cmd

import (
	"fmt"
	"golog/internal/analyzer"

	"github.com/spf13/cobra"
)

var tailCmd = &cobra.Command{
	Use:   "tail",
	Short: "Real-time monitoring of log files (like tail -f)",
	Run: func(cmd *cobra.Command, args []string) {
		filePath, _ := cmd.Flags().GetString("file")
		if filePath == "" {
			fmt.Println("Please specify a log file using -f or --file")
			return
		}

		fmt.Println("Starting real-time log tailing…")
		err := analyzer.TailFile(filePath)
		if err != nil {
			fmt.Println("Error tailing file:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(tailCmd)
	tailCmd.Flags().StringP("file", "f", "", "Log file to monitor")
}
