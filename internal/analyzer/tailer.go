package analyzer

import (
	"bufio"
	"os"
	"strings"
	"time"

	"golog/internal/color"
)

func TailFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Move to the end of file
	offset, _ := file.Seek(0, os.SEEK_END)
	_ = offset

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			time.Sleep(500 * time.Millisecond)
			continue
		}

		trimmed := strings.TrimSpace(line)

		// Color-highlighted streaming output
		lower := strings.ToLower(trimmed)
		switch {
		case strings.Contains(lower, "error"):
			println(color.RedText(trimmed))
		case strings.Contains(lower, "warn"):
			println(color.YellowText(trimmed))
		default:
			println(trimmed)
		}
	}
}
