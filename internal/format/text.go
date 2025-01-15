package format

import (
	"fmt"
	"time"

	"github.com/mnadev/adhango/pkg/calc"

	"github.com/0xzer0x/go-pray/internal/common"
	"github.com/0xzer0x/go-pray/internal/version"
)

type TextFormatter struct{}

func (f *TextFormatter) Calendar(calendar calc.PrayerTimes) (string, error) {
	var output string
	output = fmt.Sprintf("-- Date: %s --\n", calendar.Fajr.Format("Jan 02, 2006"))
	for _, prayer := range []calc.Prayer{
		calc.FAJR,
		calc.SUNRISE,
		calc.DHUHR,
		calc.ASR,
		calc.MAGHRIB,
		calc.ISHA,
	} {
		output += fmt.Sprintf(
			"%-17s%7v\n",
			common.CalendarName(calendar, prayer)+":",
			calendar.TimeForPrayer(prayer).Format(time.Kitchen),
		)
	}
	return output, nil
}

func (f *TextFormatter) Prayer(calendar calc.PrayerTimes, prayer calc.Prayer) (string, error) {
	prayTime := calendar.TimeForPrayer(prayer)
	timeRemaining := time.Until(prayTime)

	return fmt.Sprintf("%s in %02d:%02d\n",
		common.CalendarName(calendar, prayer),
		int(timeRemaining.Hours()),
		int(timeRemaining.Minutes())%60,
	), nil
}

func (f *TextFormatter) VersionInfo(versionInfo version.VersionInfo) (string, error) {
	return fmt.Sprintf(
		"%-15s%s\n%-15s%s\n%-15s%s\n%-15s%s\n%-15s%s\n",
		"Version:",
		versionInfo.Version,
		"Go Version:",
		versionInfo.Runtime,
		"Git Commit:",
		versionInfo.BuildCommit,
		"Built:",
		versionInfo.BuildTime.Format("Mon Jan 02 15:04:05 2006"),
		"OS/Arch:",
		versionInfo.Os+"/"+versionInfo.Arch,
	), nil
}
