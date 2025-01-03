package pfmt

import (
	"encoding/json"
	"time"

	"github.com/mnadev/adhango/pkg/calc"

	"github.com/0xzer0x/go-pray/internal/util"
)

type prayerInfo struct {
	Name      string
	Time      string
	Remaining string
}

type calendarInfo struct {
	Fajr    prayerInfo `json:"fajr"`
	Sunrise prayerInfo `json:"sunrise"`
	Dhuhr   prayerInfo `json:"dhuhr"`
	Asr     prayerInfo `json:"asr"`
	Maghrib prayerInfo `json:"maghrib"`
	Isha    prayerInfo `json:"isha"`
}

type JsonFormatStrategy struct{}

func newPrayerInfo(calendar calc.PrayerTimes, prayer calc.Prayer) prayerInfo {
	prayerTime := calendar.TimeForPrayer(prayer)
	prayerInf := prayerInfo{
		Name:      util.PrayerName(prayer, false),
		Time:      prayerTime.Format(time.TimeOnly),
		Remaining: time.Until(prayerTime).Truncate(time.Second).String(),
	}
	return prayerInf
}

func (f *JsonFormatStrategy) Calendar(calendar calc.PrayerTimes) string {
	pt := calendarInfo{
		Fajr:    newPrayerInfo(calendar, calc.FAJR),
		Sunrise: newPrayerInfo(calendar, calc.SUNRISE),
		Dhuhr:   newPrayerInfo(calendar, calc.DHUHR),
		Asr:     newPrayerInfo(calendar, calc.ASR),
		Maghrib: newPrayerInfo(calendar, calc.MAGHRIB),
		Isha:    newPrayerInfo(calendar, calc.ISHA),
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
