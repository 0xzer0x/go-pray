package next

import (
	"fmt"

	"github.com/mnadev/adhango/pkg/calc"
	"github.com/spf13/cobra"

	"github.com/0xzer0x/go-pray/internal/common"
	"github.com/0xzer0x/go-pray/internal/config"
	"github.com/0xzer0x/go-pray/internal/ptime"
	"github.com/0xzer0x/go-pray/internal/util"
)

var nextPrayerArg = ""

var NextCommand = &cobra.Command{
	Use:    "next [prayer]",
	Short:  "Get the next prayer time or the next occurrence of a specific prayer",
	Args:   cobra.MaximumNArgs(1),
	PreRun: validateNextArgs,
	Run:    execNext,
}

func execNext(cmd *cobra.Command, args []string) {
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

	formatter := config.Formatter()
	output, err := formatter.Prayer(prayerTimes, nextPrayer)
	if err != nil {
		util.ErrExit("%v", err)
	}

	fmt.Print(output)
}
