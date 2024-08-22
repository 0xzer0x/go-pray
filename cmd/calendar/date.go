package calendar

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/0xzer0x/go-pray/internal/config"
	"github.com/0xzer0x/go-pray/internal/praytimes"
	"github.com/0xzer0x/go-pray/internal/util"
)

var dateCmd = &cobra.Command{
	Use:   "date [dates...]",
	Short: "Get prayer times calendar for a specific date",
	Long:  `Get prayer times calendar for specific date`,
	Args:  cobra.MinimumNArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		validateDate(args)
		err := config.ValidateCalculationParams()
		if err != nil {
			util.ErrExit("%v", err)
		}
	},
	Run: calendarDateCmd,
}

func validateDate(args []string) {
	for _, date := range args {
		if _, err := time.Parse("2006-01-02", date); err != nil {
			util.ErrExit("invalid date '%s': must be in the format YYYY-MM-DD", date)
		}
	}
}

func calendarDateCmd(cmd *cobra.Command, args []string) {
	for _, date := range args {
		calendarDate, err := time.Parse("2006-01-02", date)
		if err != nil {
			util.ErrExit("%v", err)
		}

		prayerTimes, err := praytimes.DatePrayerTimes(calendarDate)
		if err != nil {
			util.ErrExit("failed to calculate prayer times: %v", err)
		}

		fmt.Printf("-- Date: %s --\n", calendarDate.Format("Jan 2, 2006"))
		praytimes.PrintCalendar(prayerTimes)
	}
}
