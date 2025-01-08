package version

import (
	"runtime"
	"time"
)

// WARN: must be set during build using -ldflags
var (
	version     = ""
	buildCommit = ""
	buildTime   = ""
)

func Version() string {
	if string(version[0]) == "v" {
		return version[1:]
	}

	return version
}

func BuildCommit() string {
	return buildCommit
}

func BuildTime() (time.Time, error) {
	if parsedBuildTime, err := time.Parse(time.DateTime, buildTime); err != nil {
		return time.Time{}, err
	} else {
		return parsedBuildTime, nil
	}
}

func OsArch() string {
	return runtime.GOOS + "/" + runtime.GOARCH
}

func Runtime() string {
	return runtime.Version()
}
