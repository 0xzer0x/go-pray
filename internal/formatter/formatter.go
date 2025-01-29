package formatter

import (
	"github.com/mnadev/adhango/pkg/calc"
	"github.com/spf13/viper"

	"github.com/0xzer0x/go-pray/internal/version"
)

type Formatter interface {
	Calendar(calendar calc.PrayerTimes) (string, error)
	Prayer(calendar calc.PrayerTimes, prayer calc.Prayer) (string, error)
	VersionInfo(versionInfo version.VersionInfo) (string, error)
}

func New() Formatter {
	value := viper.GetString("format")
	switch value {
	case "json":
		return &JsonFormatter{}
	case "table":
		return &TableFormatter{}
	default:
		return &ShortFormatter{}
	}
}
