package utils

import "fmt"

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Bold   = "\033[1m"
)

func Print(message string, color string) {
	fmt.Printf("%s%s%s\n", color, message, Reset)
}
