package daemon

import (
	"log"
	"sync"
	"time"

	"github.com/mnadev/adhango/pkg/calc"
	"github.com/spf13/cobra"

	"github.com/0xzer0x/go-pray/internal/adhan"
	"github.com/0xzer0x/go-pray/internal/common"
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
