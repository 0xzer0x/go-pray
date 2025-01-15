package format

import (
	"github.com/mnadev/adhango/pkg/calc"

	"github.com/0xzer0x/go-pray/internal/version"
)

type Formatter interface {
	Calendar(calendar calc.PrayerTimes) (string, error)
	Prayer(calendar calc.PrayerTimes, prayer calc.Prayer) (string, error)
	VersionInfo(versionInfo version.VersionInfo) (string, error)
}
