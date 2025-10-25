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

type ShortFormatter struct {
	localizer *i18n.Localizer
}

func NewShortFormatter() (*ShortFormatter, error) {
	var err error
	var localizer *i18n.Localizer
	if localizer, err = i18n.GetInstance(); err != nil {
		return nil, fmt.Errorf("failed to initialize localizer: %w", err)
	}

	sf := &ShortFormatter{
		localizer,
	}
	return sf, nil
}

func (f *ShortFormatter) Calendar(calendar calc.PrayerTimes) (string, error) {
	localizedDate, err := f.localizer.Localize("date", nil)
	if err != nil {
		return "", fmt.Errorf("failed to localize: %v", err)
	}

	dateHeader := lipgloss.PlaceHorizontal(25, lipgloss.Center, fmt.Sprintf(
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
			f.localizer.LocalizeTime(calendar.TimeForPrayer(prayer), "03:04 PM"),
		)
		body += line + "\n"
	}

	return dateHeader + "\n" + body, nil
}

func (f *ShortFormatter) Prayer(calendar calc.PrayerTimes, prayer calc.Prayer) (string, error) {
	prayTime := calendar.TimeForPrayer(prayer)
	timeRemaining := time.Until(prayTime)

	localizedNextPrayer, err := f.localizer.Localize("prayer-next", &map[string]any{
		"CalendarName": common.CalendarName(calendar, prayer),
		"Remaining": f.localizer.LocalizeTimeString(fmt.Sprintf(
			"%02d:%02d",
			int(timeRemaining.Hours()),
			int(timeRemaining.Minutes())%60,
		)),
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
