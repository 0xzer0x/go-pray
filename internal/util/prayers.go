package util

import (
	"time"

	"github.com/mnadev/adhango/pkg/calc"
)

func PrayerName(prayer calc.Prayer, considerJumuaa bool) string {
	var name string
	switch prayer {
	case calc.FAJR:
		name = "Fajr"
	case calc.SUNRISE:
		name = "Sunrise"
	case calc.DHUHR:
		if considerJumuaa && time.Now().UTC().Weekday() == time.Friday {
			name = "Jumuaa"
		} else {
			name = "Dhuhr"
		}
	case calc.ASR:
		name = "Asr"
	case calc.MAGHRIB:
		name = "Maghrib"
	case calc.ISHA:
		name = "Isha"
	}
	return name
}
