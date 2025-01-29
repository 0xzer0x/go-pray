package formatter

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/mnadev/adhango/pkg/calc"

	"github.com/0xzer0x/go-pray/internal/common"
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

type JsonFormatter struct{}

func newPrayerInfo(calendar calc.PrayerTimes, prayer calc.Prayer) prayerInfo {
	prayerTime := calendar.TimeForPrayer(prayer)
	prayerInf := prayerInfo{
		Name:      common.CalendarName(calendar, prayer),
		Time:      prayerTime.Format(time.TimeOnly),
		Remaining: time.Until(prayerTime).Truncate(time.Second).String(),
	}
	return prayerInf
}

func (f *JsonFormatter) Calendar(calendar calc.PrayerTimes) (string, error) {
	pt := calendarInfo{
		Date: calendar.Fajr.Format(time.DateOnly),
		Prayers: []prayerInfo{
			newPrayerInfo(calendar, calc.FAJR),
			newPrayerInfo(calendar, calc.SUNRISE),
			newPrayerInfo(calendar, calc.DHUHR),
			newPrayerInfo(calendar, calc.ASR),
			newPrayerInfo(calendar, calc.MAGHRIB),
			newPrayerInfo(calendar, calc.ISHA),
		},
	}

	marshaled, err := json.MarshalIndent(pt, "", " ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal prayers calendar: %v", err)
	}

	return fmt.Sprintf("%s\n", marshaled), nil
}

func (f *JsonFormatter) Prayer(calendar calc.PrayerTimes, prayer calc.Prayer) (string, error) {
	prayerInf := newPrayerInfo(calendar, prayer)
	marshaled, err := json.MarshalIndent(prayerInf, "", " ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal prayer info: %v", err)
	}
	return fmt.Sprintf("%s\n", marshaled), nil
}

func (f *JsonFormatter) VersionInfo(versionInfo version.VersionInfo) (string, error) {
	marshaled, err := json.MarshalIndent(versionInfo, "", " ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal version info: %v", err)
	}
	return fmt.Sprintf("%s\n", marshaled), nil
}
