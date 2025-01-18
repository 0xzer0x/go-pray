package daemon

import (
	"log"
	"sync"
	"time"

	"github.com/mnadev/adhango/pkg/calc"
	"github.com/spf13/cobra"

	"github.com/0xzer0x/go-pray/internal/adhan"
	"github.com/0xzer0x/go-pray/internal/common"
	"github.com/0xzer0x/go-pray/internal/notify"
	"github.com/0xzer0x/go-pray/internal/ptime"
)

var DaemonCmd = &cobra.Command{
	Use:    "daemon",
	Short:  "Start the go-pray daemon to send desktop notifications at prayer times",
	PreRun: validateDaemonArgs,
	Run:    execDaemon,
}

func execDaemon(cmd *cobra.Command, ars []string) {
	log.Println("starting in daemon mode")
	player := adhan.NewPlayer()
	notifier := notify.NewNotifier()

	if err := player.Initialize(); err != nil {
		log.Fatalf("failed to initialize player: %v\n", err)
	}
	if err := notifier.Initialize(); err != nil {
		log.Fatalf("failed to initialize notifier: %v\n", err)
	}
	defer notifier.Close()

	var err error
	var prayerTimes calc.PrayerTimes
	for {
		// INFO: get prayers calendar which the upcoming prayer belongs to
		prayerTimes, _, err = ptime.NextPrayer()
		if err != nil {
			log.Fatalf("failed to calculate prayer times: %v\n", err)
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
					notifyPrayer(player, notifier, &prayerTimes, prayer)
				}()
			}
		}

		wg.Wait()
	}
}
