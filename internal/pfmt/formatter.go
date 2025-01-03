package pfmt

import "github.com/mnadev/adhango/pkg/calc"

type FormatStrategy interface {
	Calendar(calendar calc.PrayerTimes) string
	Prayer(calendar calc.PrayerTimes, prayer calc.Prayer) string
}

type PrayerTimesFormatter struct {
	calendar calc.PrayerTimes
	strategy FormatStrategy
}

func (fmtr *PrayerTimesFormatter) Calendar() string {
	return fmtr.strategy.Calendar(fmtr.calendar)
}

func (fmtr *PrayerTimesFormatter) Prayer(prayer calc.Prayer) string {
	return fmtr.strategy.Prayer(fmtr.calendar, prayer)
}

// NOTE: builder struct
type PrayerTimesFormatterBuilder struct {
	calendar calc.PrayerTimes
	strategy FormatStrategy
}

func (b *PrayerTimesFormatterBuilder) SetCalendar(
	cal calc.PrayerTimes,
) *PrayerTimesFormatterBuilder {
	b.calendar = cal
	return b
}

func (b *PrayerTimesFormatterBuilder) SetStrategy(
	strat FormatStrategy,
) *PrayerTimesFormatterBuilder {
	b.strategy = strat
	return b
}

func (b *PrayerTimesFormatterBuilder) Build() PrayerTimesFormatter {
	return PrayerTimesFormatter{
		calendar: b.calendar,
		strategy: b.strategy,
	}
}

func NewPrayerTimesFormatterBuilder() *PrayerTimesFormatterBuilder {
	return &PrayerTimesFormatterBuilder{}
}
