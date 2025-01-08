package util

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/fatih/color"
)

func ErrExit(format string, a ...any) {
	fmt.Fprintln(os.Stderr, color.RedString("error: "+format, a...))
	os.Exit(1)
}

func Version() (version, sum string) {
	info, ok := debug.ReadBuildInfo()
	if !ok || info.Main.Version == "" {
		version = "unknown"
	} else {
		version = info.Main.Version
		sum = info.Main.Sum
	}
	return version, sum
}
