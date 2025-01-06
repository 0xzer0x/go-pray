package daemon

import (
	"log"
	"sync"
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

func notifyPrayer(
	player *adhan.Player,
	calendar *calc.PrayerTimes,
	prayer calc.Prayer,
	fromNow time.Duration,
) {
	timer := time.NewTimer(fromNow)
	name := common.CalendarName(*calendar, prayer)
	notifyChan := make(chan notify.Result)

	<-timer.C
	log.Printf("notification timer finished: %s\n", name)
	go notify.SendInteractive(
		notifyChan,
		"clock-applet-symbolic",
		"Time for "+name+" prayer 🕌",
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

func daemonCmd(cmd *cobra.Command, ars []string) {
	log.Println("starting in daemon mode")

	player := adhan.NewPlayer()
	if err := player.Initialize(); err != nil {
		util.ErrExit("%v", err)
	}

	var err error
	var prayerTimes calc.PrayerTimes
	for {
		// INFO: get prayers calendar which the upcoming prayer belongs to
		prayerTimes, _, err = ptime.NextPrayer()
		if err != nil {
			util.ErrExit("%v", err)
		}

		// INFO: create a WaitGroup for the prayers in the calendar with a future date
		var wg sync.WaitGroup
		for name, prayer := range common.Prayers {
			if prayer == calc.SUNRISE {
				continue
			}

			remainingForPrayer := time.Until(prayerTimes.TimeForPrayer(prayer))
			if remainingForPrayer > 0 {
				log.Printf(
					"creating prayer timer: %s - remaining: %s\n",
					name,
					remainingForPrayer.Truncate(time.Second).String(),
				)

				wg.Add(1)
				go func() {
					defer wg.Done()
					notifyPrayer(player, &prayerTimes, prayer, remainingForPrayer)
				}()
			}
		}

		wg.Wait()
	}
}
