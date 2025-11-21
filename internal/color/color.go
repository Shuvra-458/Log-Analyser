package color

import "fmt"

const (
	Red    = "\033[31m"
	Yellow = "\033[33m"
	Green  = "\033[32m"
	Blue   = "\033[34m"
	Reset  = "\033[0m"
)

func RedText(s string) string {
	return fmt.Sprintf("%s%s%s", Red, s, Reset)
}

func YellowText(s string) string {
	return fmt.Sprintf("%s%s%s", Yellow, s, Reset)
}

func GreenText(s string) string {
	return fmt.Sprintf("%s%s%s", Green, s, Reset)
}

func BlueText(s string) string {
	return fmt.Sprintf("%s%s%s", Blue, s, Reset)
}
