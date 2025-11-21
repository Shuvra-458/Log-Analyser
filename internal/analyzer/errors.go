package analyzer

import (
	"bufio"
	"os"
	"strings"
)

func GetErrorLinesWithTotal(path string) (lines []string, total int, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	total = 0
	lines = []string{}

	for scanner.Scan() {
		line := scanner.Text()
		total++
		if strings.Contains(strings.ToLower(line), "error") {
			lines = append(lines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return lines, total, err
	}

	return lines, total, nil
}
