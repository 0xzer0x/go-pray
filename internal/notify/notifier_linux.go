package notify

import (
	"errors"
	"fmt"
	"log"
	"slices"

	"github.com/esiqveland/notify"
	"github.com/godbus/dbus/v5"
)

type DbusNotifier struct {
	conn       *dbus.Conn
	hasActions bool
}

func NewNotifier() Notifier {
	return &DbusNotifier{}
}

func (n *DbusNotifier) Initialize() error {
	var err error
	n.conn, err = dbus.SessionBusPrivate()
	if err != nil {
		return fmt.Errorf("failed to initialize dbus connection: %v", err)
	}
	if err = n.conn.Auth(nil); err != nil {
		return fmt.Errorf("failed to authenticate to dbus: %v", err)
	}
	if err = n.conn.Hello(); err != nil {
		return fmt.Errorf("failed to send dbus hello message: %v", err)
	}

	// NOTE: check server has `actions` capability
	caps, err := notify.GetCapabilities(n.conn)
	if err != nil {
		return fmt.Errorf("failed to fetch capabilities: %v", err)
	}
	n.hasActions = slices.Contains(caps, "actions")

	return nil
}

func (n *DbusNotifier) Send(resChan chan<- Result, notification Notification) {
	sendResult := func(clicked bool, err error) {
		defer close(resChan)
		resChan <- Result{
			Clicked: clicked,
			Error:   err,
		}
	}

	if n.conn == nil {
		sendResult(false, errors.New("dbus connection not initialized"))
		return
	}
	if !n.hasActions {
		sendResult(false, errors.New("notification daemon does not support actions capability"))
		return
	}

	dbusNotification := notify.Notification{
		AppName:       "go-pray",
		AppIcon:       notification.iconName,
		Summary:       notification.title,
		Body:          notification.body,
		ExpireTimeout: notification.duration,
	}

	done := make(chan struct{})
	dbusNotifier, err := notify.New(
		n.conn,
		notify.WithOnClosed(func(closer *notify.NotificationClosedSignal) {
			defer close(done)
			log.Printf("notification close reason: %v", closer.Reason.String())
			sendResult(closer.Reason == notify.ReasonDismissedByUser, nil)
		}),
	)
	if err != nil {
		sendResult(false, err)
		return
	}
	defer dbusNotifier.Close()

	_, err = dbusNotifier.SendNotification(dbusNotification)
	if err != nil {
		sendResult(false, err)
		return
	}

	<-done
}

func (n *DbusNotifier) Close() error {
	if n.conn != nil {
		if err := n.conn.Close(); err != nil {
			return fmt.Errorf("failed to close dbus connection: %v", err)
		}
	}

	return nil
}
