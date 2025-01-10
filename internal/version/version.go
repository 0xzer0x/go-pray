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

type VersionInfo struct {
	Version     string    `json:"version"`
	BuildCommit string    `json:"build_commit"`
	BuildTime   time.Time `json:"build_time"`
	Os          string    `json:"os"`
	Arch        string    `json:"arch"`
	Runtime     string    `json:"runtime"`
}

func parseBuildTime() (time.Time, error) {
	if parsedBuildTime, err := time.Parse(time.RFC3339, buildTime); err != nil {
		return time.Time{}, err
	} else {
		return parsedBuildTime, nil
	}
}

func NewVersionInfo() (VersionInfo, error) {
	parsedTime, err := parseBuildTime()
	if err != nil {
		return VersionInfo{}, err
	}

	versionNum := version
	if string(versionNum[0]) == "v" {
		versionNum = versionNum[1:]
	}

	return VersionInfo{
		Version:     versionNum,
		BuildCommit: buildCommit,
		BuildTime:   parsedTime,
		Os:          runtime.GOOS,
		Arch:        runtime.GOARCH,
		Runtime:     runtime.Version(),
	}, nil
}
