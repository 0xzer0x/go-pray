package pfmt

import (
	"encoding/json"
	"time"

	"github.com/mnadev/adhango/pkg/calc"

	"github.com/0xzer0x/go-pray/internal/common"
	"github.com/0xzer0x/go-pray/internal/util"
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

type JsonFormatStrategy struct{}

func newPrayerInfo(calendar calc.PrayerTimes, prayer calc.Prayer) prayerInfo {
	prayerTime := calendar.TimeForPrayer(prayer)
	prayerInf := prayerInfo{
		Name:      common.PrayerName(prayer),
		Time:      prayerTime.Format(time.TimeOnly),
		Remaining: time.Until(prayerTime).Truncate(time.Second).String(),
	}
	return prayerInf
}

func (f *JsonFormatStrategy) Calendar(calendar calc.PrayerTimes) string {
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

	marshaled, err := json.Marshal(pt)
	if err != nil {
		util.ErrExit("failed to marshal prayers calendar: %v", err)
	}

	return string(marshaled)
}

func (f *JsonFormatStrategy) Prayer(calendar calc.PrayerTimes, prayer calc.Prayer) string {
	prayerInf := newPrayerInfo(calendar, prayer)
	marshaled, err := json.Marshal(prayerInf)
	if err != nil {
		util.ErrExit("failed to marshal prayer info: %v", err)
	}
	return string(marshaled)
}
