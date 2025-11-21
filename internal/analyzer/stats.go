package analyzer

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ShowStats(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	total := 0
	errors := 0
	warnings := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		total++

		if strings.Contains(strings.ToLower(line), "error") {
			errors++
		}
		if strings.Contains(strings.ToLower(line), "warn") {
			warnings++
		}
	}

	fmt.Println("Log Statistics")
	fmt.Println("---------------------------")
	fmt.Printf("Total lines:     %d\n", total)
	fmt.Printf("Error entries:   %d\n", errors)
	fmt.Printf("Warning entries: %d\n", warnings)
}
