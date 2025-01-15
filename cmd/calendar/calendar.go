package calendar

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/0xzer0x/go-pray/internal/config"
	"github.com/0xzer0x/go-pray/internal/ptime"
	"github.com/0xzer0x/go-pray/internal/util"
)

var specialDates = map[string]time.Time{
	"@yesterday": time.Now().AddDate(0, 0, -1),
	"@today":     time.Now(),
	"@tomorrow":  time.Now().AddDate(0, 0, 1),
}

var CalendarCmd = &cobra.Command{
	Use:       "calendar [date...]",
	Short:     "Get prayer times calendar for a specific date",
	Long:      `Get prayer times calendar for specific date`,
	ValidArgs: util.MapKeys(specialDates),
	PreRun:    validateCalendarArgs,
	Run:       execCalendar,
}

func execCalendar(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		args = append(args, "@today")
	}

	formatter := config.Formatter()
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

		output, err := formatter.Calendar(*prayerTimes)
		if err != nil {
			util.ErrExit("%v", err)
		}

		fmt.Print(output)
	}
}
