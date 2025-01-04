package calendar

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/0xzer0x/go-pray/internal/config"
	"github.com/0xzer0x/go-pray/internal/pfmt"
	"github.com/0xzer0x/go-pray/internal/ptime"
	"github.com/0xzer0x/go-pray/internal/util"
)

var specialDates = map[string]time.Time{
	"@yesterday": time.Now().AddDate(0, 0, -1),
	"@today":     time.Now(),
	"@tomorrow":  time.Now().AddDate(0, 0, 1),
}

var CalendarCmd = &cobra.Command{
	Use:   "calendar [date...]",
	Short: "Get prayer times calendar for a specific date",
	Long:  `Get prayer times calendar for specific date`,
	PreRun: func(cmd *cobra.Command, args []string) {
		validateDate(args)
		err := config.ValidateCalculationParams()
		if err != nil {
			util.ErrExit("%v", err)
		}
	},
	Run: calendarCmd,
}

func validateDate(args []string) {
	for _, date := range args {
		if slices.Contains(util.MapKeys(specialDates), date) {
			continue
		}
		if _, err := time.Parse("2006-01-02", date); err != nil {
			util.ErrExit(
				"invalid date '%s': must be in the format YYYY-MM-DD or one of %s",
				date,
				strings.Join(util.MapKeys(specialDates), ", "),
			)
		}
	}
}

func calendarCmd(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		args = append(args, "@today")
	}
	for _, date := range args {
		var calendarDate time.Time
		if specialTime, ok := specialDates[date]; ok {
			calendarDate = specialTime
		} else {
			calendarDate, _ = time.Parse("2006-01-02", date)
		}

		prayerTimes, err := ptime.DatePrayerTimes(calendarDate)
		if err != nil {
			util.ErrExit("failed to calculate prayer times: %v", err)
		}

		formatter := pfmt.NewPrayerTimesFormatterBuilder().
			SetCalendar(*prayerTimes).
			SetStrategy(config.FormatStrategy()).
			Build()

		fmt.Print(formatter.Calendar())
	}
}
