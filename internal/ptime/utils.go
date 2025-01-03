package ptime

import (
	"fmt"
	"time"

	"github.com/mnadev/adhango/pkg/calc"
	"github.com/mnadev/adhango/pkg/data"
	"github.com/mnadev/adhango/pkg/util"
	"github.com/spf13/viper"
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

func DatePrayerTimes(date time.Time) (*calc.PrayerTimes, error) {
	dateComponents := data.NewDateComponents(date)
	coords, params, err := calculationConfig()
	if err != nil {
		return nil, err
	}

	prayerTimes, err := calc.NewPrayerTimes(coords, dateComponents, params)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate prayer times: %v", err)
	}

	err = prayerTimes.SetTimeZone(viper.GetString("timezone"))
	if err != nil {
		return nil, fmt.Errorf("failed to use timezone: %v", err)
	}

	return prayerTimes, nil
}

func CurrentPrayerTimes() (*calc.PrayerTimes, error) {
	prayerTimes, err := DatePrayerTimes(time.Now())
	if err != nil {
		return nil, err
	}
	return prayerTimes, nil
}

func NextTime(prayer calc.Prayer) (calc.PrayerTimes, time.Time, error) {
	prayerTimes, err := CurrentPrayerTimes()
	if now := time.Now(); err == nil && !now.Before(prayerTimes.TimeForPrayer(prayer)) {
		prayerTimes, err = DatePrayerTimes(now.AddDate(0, 0, 1))
	}

	if err != nil {
		return calc.PrayerTimes{}, time.Time{}, err
	}

	return *prayerTimes, prayerTimes.TimeForPrayer(prayer), nil
}

func NextPrayer() (calc.PrayerTimes, calc.Prayer, error) {
	var err error
	var prayerTimes *calc.PrayerTimes
	nextPrayer := calc.NO_PRAYER
	date := time.Now()
	for prayerTimes, err = DatePrayerTimes(date); err == nil && nextPrayer == calc.NO_PRAYER; date = date.AddDate(0, 0, 1) {
		prayerTimes, err = DatePrayerTimes(date)
		nextPrayer = prayerTimes.NextPrayerNow()
	}

	if err != nil {
		return calc.PrayerTimes{}, calc.NO_PRAYER, err
	}
	return *prayerTimes, nextPrayer, nil
}
