package common

import (
	"time"

	"github.com/mnadev/adhango/pkg/calc"

	"github.com/0xzer0x/go-pray/internal/i18n"
	"github.com/0xzer0x/go-pray/internal/util"
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
	localizer, err := i18n.GetInstance()
	if err != nil {
		return ""
	}

	messageId := util.FindInMap(Prayers, prayer)

	var localizedName string
	if localizedName, err = localizer.Localize(messageId, nil); err != nil {
		return ""
	}

	return localizedName
}

func CalendarName(calendar calc.PrayerTimes, prayer calc.Prayer) string {
	localizer, err := i18n.GetInstance()
	if err != nil {
		return ""
	}

	var messageID string
	if IsJumuaa(calendar) && prayer == calc.DHUHR {
		messageID = "jumuaa"
	} else {
		messageID = util.FindInMap(Prayers, prayer)
	}

	var localizedName string
	if localizedName, err = localizer.Localize(messageID, nil); err != nil {
		return ""
	}

	return localizedName
}
