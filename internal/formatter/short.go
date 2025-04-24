package formatter

import (
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/mnadev/adhango/pkg/calc"

	"github.com/0xzer0x/go-pray/internal/common"
	"github.com/0xzer0x/go-pray/internal/i18n"
	"github.com/0xzer0x/go-pray/internal/version"
)

type ShortFormatter struct{}

func (f *ShortFormatter) Calendar(calendar calc.PrayerTimes) (string, error) {
	localizer, err := i18n.GetInstance()
	if err != nil {
		return "", fmt.Errorf("failed to instantiate localizer: %v", err)
	}

	localizedDate, err := localizer.Localize("date", nil)
	if err != nil {
		return "", fmt.Errorf("failed to localize: %v", err)
	}

	var dateHeader string
	dateHeader = lipgloss.PlaceHorizontal(25, lipgloss.Center, fmt.Sprintf(
		" %s: %s ",
		localizedDate,
		calendar.Fajr.Format(time.DateOnly),
	), lipgloss.WithWhitespaceChars("-"))

	var body string
	for _, prayer := range []calc.Prayer{
		calc.FAJR,
		calc.SUNRISE,
		calc.DHUHR,
		calc.ASR,
		calc.MAGHRIB,
		calc.ISHA,
	} {
		line := lipgloss.PlaceHorizontal(
			17,
			lipgloss.Left,
			common.CalendarName(calendar, prayer)+":",
		) + lipgloss.PlaceHorizontal(
			8,
			lipgloss.Right,
			calendar.TimeForPrayer(prayer).Format(time.Kitchen),
		)
		body += line + "\n"
	}

	return dateHeader + "\n" + body, nil
}

func (f *ShortFormatter) Prayer(calendar calc.PrayerTimes, prayer calc.Prayer) (string, error) {
	localizer, err := i18n.GetInstance()
	if err != nil {
		return "", fmt.Errorf("failed to instantiate localizer: %v", err)
	}

	prayTime := calendar.TimeForPrayer(prayer)
	timeRemaining := time.Until(prayTime)

	localizedNextPrayer, err := localizer.Localize("prayer-next", &map[string]any{
		"CalendarName": common.CalendarName(calendar, prayer),
		"Remaining": fmt.Sprintf(
			"%02d:%02d",
			int(timeRemaining.Hours()),
			int(timeRemaining.Minutes())%60,
		),
	})
	if err != nil {
		return "", fmt.Errorf("failed to localize message: %v", err)
	}

	return localizedNextPrayer, nil
}

func (f *ShortFormatter) VersionInfo(versionInfo version.VersionInfo) (string, error) {
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
