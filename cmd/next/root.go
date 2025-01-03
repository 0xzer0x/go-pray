package next

import (
	"fmt"
	"strings"

	"github.com/mnadev/adhango/pkg/calc"
	"github.com/spf13/cobra"

	"github.com/0xzer0x/go-pray/internal/common"
	"github.com/0xzer0x/go-pray/internal/config"
	"github.com/0xzer0x/go-pray/internal/pfmt"
	"github.com/0xzer0x/go-pray/internal/ptime"
	"github.com/0xzer0x/go-pray/internal/util"
)

var nextPrayerArg = ""

var NextCommand = &cobra.Command{
	Use:    "next [prayer]",
	Short:  "Get the next prayer time or the next occurrence of a specific prayer",
	Args:   cobra.MaximumNArgs(1),
	PreRun: nextCmdValidations,
	Run:    nextCmd,
}

func nextCmdValidations(cmd *cobra.Command, args []string) {
	err := config.ValidateCalculationParams()
	if err != nil {
		util.ErrExit("%v", err)
	}

	if len(args) > 0 {
		nextPrayerArg = strings.ToLower(args[0])
		validArg := false
		for name := range common.Prayers {
			if nextPrayerArg == name {
				validArg = true
				break
			}
		}
		if !validArg {
			util.ErrExit(
				"invalid prayer '%s', allowed values are: %s",
				args[0],
				strings.Join(util.MapKeys(common.Prayers), ", "),
			)
		}
	}
}

func nextCmd(cmd *cobra.Command, args []string) {
	var prayerTimes calc.PrayerTimes
	var nextPrayer calc.Prayer
	var err error

	if len(nextPrayerArg) > 0 {
		nextPrayer = common.Prayers[nextPrayerArg]
		prayerTimes, _, err = ptime.NextTime(nextPrayer)
	} else {
		prayerTimes, nextPrayer, err = ptime.NextPrayer()
	}

	if err != nil {
		util.ErrExit("failed to calculate next prayer time: %v", err)
	}

	formatter := pfmt.NewPrayerTimesFormatterBuilder().
		SetCalendar(prayerTimes).
		SetStrategy(config.FormatStrategy()).
		Build()

	fmt.Println(formatter.Prayer(nextPrayer))
}
