package formatter

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/mnadev/adhango/pkg/calc"

	"github.com/0xzer0x/go-pray/internal/common"
	"github.com/0xzer0x/go-pray/internal/i18n"
	"github.com/0xzer0x/go-pray/internal/version"
)

type prayerInfo struct {
	Name      string `json:"name"`
	Time      string `json:"time"`
	Remaining string `json:"remaining"`
}

type calendarInfo struct {
	Date    string       `json:"date"`
	Prayers []prayerInfo `json:"prayers"`
}

type JSONFormatter struct {
	localizer *i18n.Localizer
}

func NewJSONFormatter() (*JSONFormatter, error) {
	var err error
	var localizer *i18n.Localizer
	if localizer, err = i18n.GetInstance(); err != nil {
		return nil, fmt.Errorf("failed to initialize localizer: %w", err)
	}

	jf := &JSONFormatter{
		localizer,
	}
	return jf, nil
}

func (f *JSONFormatter) newPrayerInfo(calendar calc.PrayerTimes, prayer calc.Prayer) prayerInfo {
	prayerTime := calendar.TimeForPrayer(prayer)
	prayerInf := prayerInfo{
		Name:      common.CalendarName(calendar, prayer),
		Time:      f.localizer.LocalizeTime(prayerTime, time.TimeOnly),
		Remaining: f.localizer.LocalizeDuration(time.Until(prayerTime).Truncate(time.Second)),
	}
	return prayerInf
}

func (f *JSONFormatter) Calendar(calendar calc.PrayerTimes) (string, error) {
	pt := calendarInfo{
		Date: f.localizer.LocalizeTime(calendar.Fajr, time.DateOnly),
		Prayers: []prayerInfo{
			f.newPrayerInfo(calendar, calc.FAJR),
			f.newPrayerInfo(calendar, calc.SUNRISE),
			f.newPrayerInfo(calendar, calc.DHUHR),
			f.newPrayerInfo(calendar, calc.ASR),
			f.newPrayerInfo(calendar, calc.MAGHRIB),
			f.newPrayerInfo(calendar, calc.ISHA),
		},
	}

	marshaled, err := json.MarshalIndent(pt, "", " ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal prayers calendar: %v", err)
	}

	return fmt.Sprintf("%s\n", marshaled), nil
}

func (f *JSONFormatter) Prayer(calendar calc.PrayerTimes, prayer calc.Prayer) (string, error) {
	prayerInf := f.newPrayerInfo(calendar, prayer)
	marshaled, err := json.MarshalIndent(prayerInf, "", " ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal prayer info: %v", err)
	}
	return fmt.Sprintf("%s\n", marshaled), nil
}

func (f *JSONFormatter) VersionInfo(versionInfo version.VersionInfo) (string, error) {
	marshaled, err := json.MarshalIndent(versionInfo, "", " ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal version info: %v", err)
	}
	return fmt.Sprintf("%s\n", marshaled), nil
}
