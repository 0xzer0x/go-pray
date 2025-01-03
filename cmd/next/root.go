package next

import (
	"fmt"
	"time"

	"github.com/mnadev/adhango/pkg/calc"
	"github.com/spf13/cobra"

	"github.com/0xzer0x/go-pray/internal/config"
	"github.com/0xzer0x/go-pray/internal/pfmt"
	"github.com/0xzer0x/go-pray/internal/ptime"
	"github.com/0xzer0x/go-pray/internal/util"
)

var NextCommand = &cobra.Command{
	Use:   "next",
	Short: "Get the next prayer",
	PreRun: func(cmd *cobra.Command, args []string) {
		err := config.ValidateCalculationParams()
		if err != nil {
			util.ErrExit("%v", err)
		}
	},
	Run: nextCmd,
}

func nextCmd(cmd *cobra.Command, args []string) {
	prayerTimes, err := ptime.CurrentPrayerTimes()
	if err != nil {
		util.ErrExit("failed to calculate prayer times: %v", err)
	}

	var nextPrayer calc.Prayer
	for nextPrayer = prayerTimes.NextPrayerNow(); nextPrayer == calc.NO_PRAYER; nextPrayer = prayerTimes.NextPrayerNow() {
		nextDay := time.Date(
			prayerTimes.DateComponent.Year,
			time.Month(prayerTimes.DateComponent.Month),
			prayerTimes.DateComponent.Day,
			0,
			0,
			0,
			0,
			time.UTC,
		).AddDate(0, 0, 1)
		prayerTimes, err = ptime.DatePrayerTimes(nextDay)
		if err != nil {
			util.ErrExit("failed to calculate prayer times: %v", err)
		}
	}

	formatter := pfmt.NewPrayerTimesFormatterBuilder().
		SetCalendar(*prayerTimes).
		SetStrategy(config.FormatStrategy()).
		Build()

	fmt.Println(formatter.Prayer(nextPrayer))
}
