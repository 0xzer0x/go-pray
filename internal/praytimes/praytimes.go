package praytimes

import (
	"fmt"
	"time"

	"github.com/mnadev/adhango/pkg/calc"
	"github.com/mnadev/adhango/pkg/data"
)

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

	return prayerTimes, nil
}

func CurrentPrayerTimes() (*calc.PrayerTimes, error) {
	prayerTimes, err := DatePrayerTimes(time.Now())
	if err != nil {
		return nil, err
	}
	return prayerTimes, nil
}
