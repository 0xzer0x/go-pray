package notify

import (
	"fmt"
	"strings"
	"time"

	"github.com/mnadev/adhango/pkg/calc"
	"github.com/spf13/cobra"

	"github.com/0xzer0x/go-pray/internal/config"
	"github.com/0xzer0x/go-pray/internal/praytimes"
	"github.com/0xzer0x/go-pray/internal/util"
)

var NotifyCmd = &cobra.Command{
	Use:   "notify",
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
	err := util.SendNotification("Started in daemon mode", "")
	if err != nil {
		util.ErrExit("failed to send notification: %v", err)
	}

	prayerTimes, err := praytimes.CurrentPrayerTimes()
	if err != nil {
		util.ErrExit("%v", err)
	}

	for {
		nextPrayer := prayerTimes.NextPrayerNow()
		nextName := praytimes.PrayerName(nextPrayer)
		fmt.Printf("next prayer: %v\nstarting timer...\n", strings.ToLower(nextName))

		timeRemaining := prayerTimes.TimeForPrayer(nextPrayer).Sub(time.Now().UTC())
		notifyTimer := time.NewTimer(timeRemaining)

		<-notifyTimer.C
		err = util.SendNotification("Time for "+nextName+" prayer 🕌", "clock-applet-symbolic")
		if err != nil {
			util.ErrExit("%v", err)
		}

		if nextPrayer == calc.NO_PRAYER {
			prayerTimes, err = praytimes.CurrentPrayerTimes()
			if err != nil {
				util.ErrExit("%v", err)
			}
		}
	}
}
