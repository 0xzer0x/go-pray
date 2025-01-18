package daemon

import (
	"log"
	"time"

	"github.com/mnadev/adhango/pkg/calc"

	"github.com/0xzer0x/go-pray/internal/adhan"
	"github.com/0xzer0x/go-pray/internal/common"
	"github.com/0xzer0x/go-pray/internal/notify"
)

func notifyPrayer(
	player adhan.Player,
	notifier notify.Notifier,
	calendar *calc.PrayerTimes,
	prayer calc.Prayer,
) {
	notifyChan := make(chan notify.Result, 1)
	name := common.CalendarName(*calendar, prayer)
	prayerTime := calendar.TimeForPrayer(prayer)
	notification := notify.NewNotificationBuilder().
		SetIconName("clock-applet-symbolic").
		SetTitle("Prayer").
		SetBody("Time for " + name + " prayer 🕌").
		SetDuration(player.Duration()).
		Build()

	log.Printf("creating new timer: %s - time: %s\n", name, prayerTime.Format(time.DateTime))
	timer := time.NewTimer(time.Until(prayerTime))

	<-timer.C
	log.Printf("notification timer finished: %s\n", name)
	go notifier.SendInteractive(notifyChan, notification)
	if err := player.Play(); err != nil {
		log.Fatalf("failed to play adhan: %v\n", err)
	}

	result := <-notifyChan
	if result.Error != nil {
		log.Fatalf("failed to send notification: %v\n", result.Error)
	}
	if result.Clicked {
		if err := player.Stop(); err != nil {
			log.Fatalf("failed to stop adhan playback: %v\n", err)
		}
	}
}
