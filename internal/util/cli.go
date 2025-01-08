package util

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func ErrExit(format string, a ...any) {
	fmt.Fprintln(os.Stderr, color.RedString("error: "+format, a...))
	os.Exit(1)
}
