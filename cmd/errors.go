package cmd

import (
	"fmt"
	"golog/internal/analyzer"
	"golog/internal/exporter"

	"github.com/spf13/cobra"
)

var errorsCmd = &cobra.Command{
	Use:   "errors",
	Short: "Show only error entries from a log file",
	Run: func(cmd *cobra.Command, args []string) {
		if filePath == "" {
			fmt.Println("Please provide a file using --file <path>")
			return
		}

		if outFile != "" && outAppend != "" {
			fmt.Println("Cannot use both --out and --out-append at the same time.")
			return
		}

		lines, total, err := analyzer.GetErrorLinesWithTotal(filePath)
		if err != nil {
			fmt.Println("Error scanning file:", err)
			return
		}

		
		exporter.PrintErrorsConsoleUTC(lines, total)

		
		if outFile != "" {
			if err := exporter.WriteErrorsReportUTC(outFile, false, lines, total); err != nil {
				fmt.Println("Error writing report:", err)
			} else {
				fmt.Println("Report written to", outFile)
			}
		} else if outAppend != "" {
			if err := exporter.WriteErrorsReportUTC(outAppend, true, lines, total); err != nil {
				fmt.Println("Error appending report:", err)
			} else {
				fmt.Println("Report appended to", outAppend)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(errorsCmd)
}
