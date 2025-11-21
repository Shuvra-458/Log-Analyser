package exporter

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Helper: UTC timestamp header (Theme 3)
func utcHeader() string {
	return fmt.Sprintf("[Report Generated: %s]\n\n", time.Now().UTC().Format("2006-01-02 15:04:05 UTC"))
}

// ----------------------- STATS output -----------------------
func PrintStatsConsoleUTC(total, errs, warns int) {
	fmt.Print(utcHeader())
	fmt.Println("Log Statistics")
	fmt.Println("---------------------------")
	printStatsBody(total, errs, warns)
}

func WriteStatsReportUTC(path string, appendMode bool, total, errs, warns int) error {
	w, err := openWriter(path, appendMode)
	if err != nil {
		return err
	}
	defer w.Close()

	_, _ = io.WriteString(w, utcHeader())
	_, _ = io.WriteString(w, "Log Statistics\n")
	_, _ = io.WriteString(w, "---------------------------\n")
	printStatsBodyToWriter(w, total, errs, warns)
	_, _ = io.WriteString(w, "\n")
	return nil
}

func printStatsBody(total, errs, warns int) {
	printStatsBodyToWriter(os.Stdout, total, errs, warns)
}

func printStatsBodyToWriter(w io.Writer, total, errs, warns int) {
	// percentages
	var pErr, pWarn float64
	if total > 0 {
		pErr = (float64(errs) / float64(total)) * 100.0
		pWarn = (float64(warns) / float64(total)) * 100.0
	}
	_, _ = fmt.Fprintf(w, "Total Lines:       %d\n", total)
	_, _ = fmt.Fprintf(w, "Error Entries:     %d   (%.1f%%)\n", errs, pErr)
	_, _ = fmt.Fprintf(w, "Warning Entries:   %d   (%.1f%%)\n", warns, pWarn)
}

// ----------------------- ERRORS output -----------------------
func PrintErrorsConsoleUTC(lines []string, total int) {
	fmt.Print(utcHeader())
	fmt.Println("Error Entries")
	fmt.Println("---------------------------")
	for _, l := range lines {
		fmt.Println(l)
	}
	_, _ = fmt.Fprintf(os.Stdout, "\nTotal Errors: %d\n", len(lines))
	_, _ = fmt.Fprintf(os.Stdout, "Scanned Lines: %d\n", total)
}

func WriteErrorsReportUTC(path string, appendMode bool, lines []string, total int) error {
	w, err := openWriter(path, appendMode)
	if err != nil {
		return err
	}
	defer w.Close()

	_, _ = io.WriteString(w, utcHeader())
	_, _ = io.WriteString(w, "Error Entries\n")
	_, _ = io.WriteString(w, "---------------------------\n")
	for _, l := range lines {
		_, _ = io.WriteString(w, l+"\n")
	}
	_, _ = fmt.Fprintf(w, "\nTotal Errors: %d\n", len(lines))
	_, _ = fmt.Fprintf(w, "Scanned Lines: %d\n", total)
	return nil
}

// ----------------------- file helper -----------------------
func openWriter(path string, appendMode bool) (*os.File, error) {
	if appendMode {
		return os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	}
	// overwrite
	return os.Create(path)
}
