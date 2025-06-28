package daemon

import (
	"log"
	"time"

	"github.com/mnadev/adhango/pkg/calc"
	"github.com/spf13/viper"

	"github.com/0xzer0x/go-pray/internal/adhan"
	"github.com/0xzer0x/go-pray/internal/common"
	"github.com/0xzer0x/go-pray/internal/notify"
	"github.com/0xzer0x/go-pray/internal/util"
)

func notifyPrayer(
	player adhan.Player,
	notifier notify.Notifier,
	calendar *calc.PrayerTimes,
	prayer calc.Prayer,
) {
	resultChan := make(chan notify.Result, 1)
	name := common.CalendarName(*calendar, prayer)
	prayerTime := calendar.TimeForPrayer(prayer)
	notification, err := notify.NewNotificationBuilder().
		SetIconName(viper.GetString("notification.icon")).
		SetTitleTemplate(viper.GetString("notification.title")).
		SetBodyTemplate(viper.GetString("notification.body")).
		SetDuration(player.Duration()).
		SetPrayer(calendar, prayer).
		Build()
	if err != nil {
		log.Fatalf("failed to build notification: %v", err)
	}

	log.Printf(
		"creating new timer: %s - time: %s\n",
		util.FindInMap(common.Prayers, prayer),
		prayerTime.Format(time.DateTime),
	)
	timer := time.NewTimer(time.Until(prayerTime))
	updateTicker := time.NewTicker(time.Minute)

	done := false
	for !done {
		select {
		case <-updateTicker.C:
			timer.Reset(time.Until(prayerTime))
		case <-timer.C:
			updateTicker.Stop()
			done = true
		}
	}

	if time.Now().After(prayerTime.Add(time.Minute)) {
		log.Printf("prayer time passed: %s, skipping notification\n", name)
		return
	}

	log.Printf("notification timer finished: %s\n", name)
	go notifier.Send(resultChan, notification)
	if err := player.Play(); err != nil {
		log.Fatalf("failed to play adhan: %v\n", err)
	}

	result := <-resultChan
	if result.Error != nil {
		log.Fatalf("failed to send notification: %v\n", result.Error)
	}
	if result.Clicked {
		if err := player.Stop(); err != nil {
			log.Fatalf("failed to stop adhan playback: %v\n", err)
		}
	}
}
