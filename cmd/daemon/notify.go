package daemon

import (
	"log"
	"time"

	"github.com/mnadev/adhango/pkg/calc"

	"github.com/0xzer0x/go-pray/internal/common"
	"github.com/0xzer0x/go-pray/internal/notify"
	"github.com/0xzer0x/go-pray/internal/util"
)

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
