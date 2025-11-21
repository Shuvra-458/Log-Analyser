package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Show statistics of a log file",
	Run: func(cmd *cobra.Command, args []string) {

		filePath, _ := cmd.Flags().GetString("file")
		if filePath == "" {
			fmt.Println("Please specify a log file using -f <path>")
			return
		}

		content, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		lines := splitLines(string(content))

		total := len(lines)
		errors := 0
		warnings := 0

		for _, line := range lines {
			if contains(line, "error") {
				errors++
			}
			if contains(line, "warn") {
				warnings++
			}
		}

		
		output := fmt.Sprintf(
			"Log Statistics\n---------------------------\nTotal lines:     %d\nError entries:   %d\nWarning entries: %d\n",
			total, errors, warnings,
		)

		fmt.Println("golog CLI starting…")
		fmt.Print(output)

		
		if outFile != "" {
			
			if isDir(outFile) {
				timestamp := time.Now().Format("2006-01-02_15-04-05")
				outFile = fmt.Sprintf("%s/golog_stats_%s.txt", outFile, timestamp)
			}

			err := os.WriteFile(outFile, []byte(output), 0644)
			if err != nil {
				fmt.Println("Failed to write output:", err)
				return
			}

			fmt.Println("Stats saved to:", outFile)
		}
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
	statsCmd.Flags().StringP("file", "f", "", "Log file to analyze")
	statsCmd.Flags().StringVarP(&outFile, "output-file", "o", "", "Save stats output to a file")
}


func isDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}
