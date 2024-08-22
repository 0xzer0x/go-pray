package praytimes

import (
	"fmt"
	"time"

	"github.com/mnadev/adhango/pkg/calc"
	"github.com/mnadev/adhango/pkg/util"
	"github.com/spf13/viper"

	"github.com/0xzer0x/go-pray/internal/config"
	cli "github.com/0xzer0x/go-pray/internal/util"
)

func calculationConfig() (*util.Coordinates, *calc.CalculationParameters, error) {
	coords, err := util.NewCoordinates(
		viper.GetFloat64("location.lat"),
		viper.GetFloat64("location.long"),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to construct coordinates: %v", err)
	}

	method, err := CalculationMethod(viper.GetString("method"))
	if err != nil {
		return nil, nil, err
	}

	params := calc.GetMethodParameters(method)
	return coords, params, nil
}

func PrayerName(prayer calc.Prayer) string {
	name := ""
	switch prayer {
	case calc.FAJR:
		name = "Fajr"
	case calc.DHUHR:
		name = "Dhuhr"
	case calc.ASR:
		name = "Asr"
	case calc.MAGHRIB:
		name = "Maghrib"
	case calc.ISHA:
		name = "Isha"
	}
	return name
}

func PrintCalendar(prayerTimes *calc.PrayerTimes) {
	config.ValidateSet("timezone")
	err := prayerTimes.SetTimeZone(viper.GetString("timezone"))
	if err != nil {
		cli.ErrExit("failed to use timezone: %v", err)
	}

	fmt.Printf("%-8s%7v\n", "Fajr:", prayerTimes.Fajr.Format(time.Kitchen))
	fmt.Printf("%-8s%7v\n", "Sunrise:", prayerTimes.Sunrise.Format(time.Kitchen))
	fmt.Printf("%-8s%7v\n", "Dhuhr:", prayerTimes.Dhuhr.Format(time.Kitchen))
	fmt.Printf("%-8s%7v\n", "Asr:", prayerTimes.Asr.Format(time.Kitchen))
	fmt.Printf("%-8s%7v\n", "Maghrib:", prayerTimes.Maghrib.Format(time.Kitchen))
	fmt.Printf("%-8s%7v\n", "Isha:", prayerTimes.Isha.Format(time.Kitchen))
}
