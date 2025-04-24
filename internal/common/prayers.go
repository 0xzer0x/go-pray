package common

import (
	"time"

	"github.com/mnadev/adhango/pkg/calc"
	"github.com/spf13/viper"
)

var Prayers = map[string]calc.Prayer{
	"fajr":    calc.FAJR,
	"sunrise": calc.SUNRISE,
	"dhuhr":   calc.DHUHR,
	"asr":     calc.ASR,
	"maghrib": calc.MAGHRIB,
	"isha":    calc.ISHA,
}

var prayerNames = struct {
	En map[calc.Prayer]string
	Ar map[calc.Prayer]string
}{
	En: map[calc.Prayer]string{
		calc.FAJR:    "Fajr",
		calc.SUNRISE: "Sunrise",
		calc.DHUHR:   "Dhuhr",
		calc.ASR:     "Asr",
		calc.MAGHRIB: "Maghrib",
		calc.ISHA:    "Isha",
	},
	Ar: map[calc.Prayer]string{
		calc.FAJR:    "الفجر",
		calc.SUNRISE: "الشروق",
		calc.DHUHR:   "الظهر",
		calc.ASR:     "العصر",
		calc.MAGHRIB: "المغرب",
		calc.ISHA:    "العشاء",
	},
}

func IsJumuaa(prayerTimes calc.PrayerTimes) bool {
	return prayerTimes.Dhuhr.Weekday() == time.Friday
}

func PrayerName(prayer calc.Prayer) string {
	lang := viper.GetString("language")
	if lang == "ar" {
		return prayerNames.Ar[prayer]
	} else {
		return prayerNames.En[prayer]
	}
}

func CalendarName(calendar calc.PrayerTimes, prayer calc.Prayer) string {
	lang := viper.GetString("language")
	var name string
	if IsJumuaa(calendar) && prayer == calc.DHUHR {
		if lang == "ar" {
			name = "الجمعة"
		} else {
			name = "Jumuaa"
		}
	} else {
		name = PrayerName(prayer)
	}
	return name
}
