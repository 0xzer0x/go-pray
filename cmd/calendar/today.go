package calendar

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/0xzer0x/go-pray/internal/config"
	"github.com/0xzer0x/go-pray/internal/pfmt"
	"github.com/0xzer0x/go-pray/internal/ptime"
	"github.com/0xzer0x/go-pray/internal/util"
)

var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "Get prayer times calendar for today",
	Long:  `Get prayer times calendar for today`,
	PreRun: func(cmd *cobra.Command, args []string) {
		err := config.ValidateCalculationParams()
		if err != nil {
			util.ErrExit("%v", err)
		}
	},
	Run: calendarTodayCmd,
}

func calendarTodayCmd(cmd *cobra.Command, args []string) {
	prayerTimes, err := ptime.CurrentPrayerTimes()
	if err != nil {
		util.ErrExit("failed to calculate prayer times: %v", err)
	}

	formatter := pfmt.NewPrayerTimesFormatterBuilder().
		SetCalendar(*prayerTimes).
		SetStrategy(config.FormatStrategy()).
		Build()
	fmt.Print(formatter.Calendar())
}
