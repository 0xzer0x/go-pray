package daemon

import (
	"fmt"
	"strings"
	"time"

	"github.com/mnadev/adhango/pkg/calc"
	"github.com/spf13/cobra"

	"github.com/0xzer0x/go-pray/internal/adhan"
	"github.com/0xzer0x/go-pray/internal/common"
	"github.com/0xzer0x/go-pray/internal/config"
	"github.com/0xzer0x/go-pray/internal/notify"
	"github.com/0xzer0x/go-pray/internal/ptime"
	"github.com/0xzer0x/go-pray/internal/util"
)

var DaemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Start the go-pray daemon to send desktop notifications at prayer times",
	PreRun: func(cmd *cobra.Command, args []string) {
		var err error
		err = config.ValidateCalculationParams()
		err = config.ValidateKey("adhan")
		if err != nil {
			util.ErrExit("%v", err)
		}
	},
	Run: daemonCmd,
}

func daemonCmd(cmd *cobra.Command, ars []string) {
	fmt.Println("starting in daemon mode")

	player := adhan.NewPlayer()
	if err := player.Initialize(); err != nil {
		util.ErrExit("%v", err)
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

		fmt.Printf("next prayer: %s\n", strings.ToLower(nextName))
		timeRemaining := time.Until(prayerTimes.TimeForPrayer(nextPrayer))
		fmt.Printf("time remaining: %s\n", timeRemaining.String())
		notifyTimer := time.NewTimer(timeRemaining)
		notifyChan := make(chan notify.Result)

		<-notifyTimer.C
		fmt.Println("sending prayer notification")
		go notify.SendInteractive(
			notifyChan,
			"clock-applet-symbolic",
			"Time for "+nextName+" prayer 🕌",
			player.Duration(),
		)
		if err := player.Play(); err != nil {
			util.ErrExit("%v", err)
		}

		result := <-notifyChan
		if result.Error != nil {
			util.ErrExit("failed to send notification: %v", result.Error)
		}
		if result.Clicked {
			player.Stop()
		}
	}
}
