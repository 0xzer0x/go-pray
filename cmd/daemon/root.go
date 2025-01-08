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

var player *adhan.Player = adhan.NewPlayer()

var DaemonCmd = &cobra.Command{
	Use:    "daemon",
	Short:  "Start the go-pray daemon to send desktop notifications at prayer times",
	PreRun: validateDaemonArgs,
	Run:    execDaemon,
}

func validateDaemonArgs(cmd *cobra.Command, args []string) {
	var err error
	err = config.ValidateCalculationParams()
	err = config.ValidateKey("adhan")
	if err != nil {
		util.ErrExit("%v", err)
	}
}

func notifyPrayer(calendar *calc.PrayerTimes, prayer calc.Prayer) {
	notifyChan := make(chan notify.Result, 1)
	name := common.CalendarName(*calendar, prayer)
	prayerTime := calendar.TimeForPrayer(prayer)

	log.Printf("creating new timer: %s - time: %s\n", name, prayerTime.Format(time.DateTime))
	timer := time.NewTimer(time.Until(prayerTime))

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

func execDaemon(cmd *cobra.Command, ars []string) {
	log.Println("starting in daemon mode")

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
		for _, prayer := range common.Prayers {
			if prayer == calc.SUNRISE {
				continue
			}

			if time.Now().Before(prayerTimes.TimeForPrayer(prayer)) {
				wg.Add(1)
				go func() {
					defer wg.Done()
					notifyPrayer(&prayerTimes, prayer)
				}()
			}
		}

		wg.Wait()
	}
}
