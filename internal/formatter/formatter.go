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

func New() (Formatter, error) {
	var formatter Formatter
	var err error

	value := viper.GetString("format")
	switch value {
	case "json":
		formatter, err = NewJSONFormatter()
	case "table":
		formatter, err = NewTableFormatter()
	default:
		formatter, err = NewShortFormatter()
	}

	return formatter, err
}
