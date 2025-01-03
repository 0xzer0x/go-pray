package pfmt

import (
	"fmt"
	"time"

	"github.com/mnadev/adhango/pkg/calc"

	"github.com/0xzer0x/go-pray/internal/util"
)

type TextFormatStrategy struct{}

func (f *TextFormatStrategy) Calendar(calendar calc.PrayerTimes) string {
	var output string
	for _, prayer := range []calc.Prayer{
		calc.FAJR,
		calc.SUNRISE,
		calc.DHUHR,
		calc.ASR,
		calc.MAGHRIB,
		calc.ISHA,
	} {
		output += fmt.Sprintf(
			"%-8s%7v\n",
			util.PrayerName(prayer, true)+":",
			calendar.TimeForPrayer(prayer).Format(time.Kitchen),
		)
	}
	return output
}

func (f *TextFormatStrategy) Prayer(calendar calc.PrayerTimes, prayer calc.Prayer) string {
	prayTime := calendar.TimeForPrayer(prayer)
	timeRemaining := time.Until(prayTime)

	return fmt.Sprintf("%s in %02d:%02d",
		util.PrayerName(prayer, true),
		int(timeRemaining.Hours()),
		int(timeRemaining.Minutes())%60,
	)
}
