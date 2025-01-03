package common

import (
	"strings"
	"time"

	"github.com/mnadev/adhango/pkg/calc"
)

var Prayers = map[string]calc.Prayer{
	"fajr":    calc.FAJR,
	"sunrise": calc.SUNRISE,
	"dhuhr":   calc.DHUHR,
	"asr":     calc.ASR,
	"maghrib": calc.MAGHRIB,
	"isha":    calc.ISHA,
}

func IsJumuaa(prayerTimes calc.PrayerTimes) bool {
	return prayerTimes.Dhuhr.Weekday() == time.Friday
}

func PrayerName(prayer calc.Prayer) string {
	var name string
	for k, v := range Prayers {
		if v == prayer {
			name = k
		}
	}

	if len(name) > 0 {
		name = strings.ToUpper(string(name[0])) + name[1:]
	}
	return name
}

func CalendarName(calendar calc.PrayerTimes, prayer calc.Prayer) string {
	var name string
	if IsJumuaa(calendar) && prayer == calc.DHUHR {
		name = "Jumuaa"
	} else {
		name = PrayerName(prayer)
	}
	return name
}
