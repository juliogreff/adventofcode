package debug

import "fmt"

func Print(format string, a ...any) {
	format += "\n"
	fmt.Printf(format, a...)
}
