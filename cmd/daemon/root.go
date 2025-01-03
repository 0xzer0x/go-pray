package daemon

import (
	"fmt"
	"strings"
	"time"

	"github.com/mnadev/adhango/pkg/calc"
	"github.com/spf13/cobra"

	"github.com/0xzer0x/go-pray/internal/common"
	"github.com/0xzer0x/go-pray/internal/config"
	"github.com/0xzer0x/go-pray/internal/ptime"
	"github.com/0xzer0x/go-pray/internal/util"
)

var DaemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Start the go-pray daemon to send desktop notifications at prayer times",
	PreRun: func(cmd *cobra.Command, args []string) {
		err := config.ValidateCalculationParams()
		if err != nil {
			util.ErrExit("%v", err)
		}
	},
	Run: runNotify,
}

func runNotify(cmd *cobra.Command, ars []string) {
	err := util.SendNotification("bell", "Started in daemon mode")
	if err != nil {
		util.ErrExit("failed to send notification: %v", err)
	}

	prayerTimes, err := ptime.CurrentPrayerTimes()
	if err != nil {
		util.ErrExit("%v", err)
	}

	for {
		if nxt := prayerTimes.NextPrayerNow(); nxt == calc.NO_PRAYER {
			prayerTimes, err = ptime.DatePrayerTimes(time.Now().UTC().AddDate(0, 0, 1))
			if err != nil {
				util.ErrExit("%v", err)
			}
		}

		nextPrayer := prayerTimes.NextPrayerNow()
		nextName := common.CalendarName(*prayerTimes, nextPrayer)

		fmt.Printf("next prayer: %s\nstarting timer...\n", strings.ToLower(nextName))
		timeRemaining := prayerTimes.TimeForPrayer(nextPrayer).Sub(time.Now().UTC())
		notifyTimer := time.NewTimer(timeRemaining)

		<-notifyTimer.C
		err = util.SendNotification("clock-applet-symbolic", "Time for "+nextName+" prayer 🕌")
		if err != nil {
			util.ErrExit("%v", err)
		}

	}
}
