package notify

import (
	"fmt"
	"strings"
	"text/template"
	"time"

	"github.com/mnadev/adhango/pkg/calc"

	"github.com/0xzer0x/go-pray/internal/common"
)

type Notification struct {
	iconName, title, body string
	duration              time.Duration
}

type NotificationBuilder struct {
	iconName, title, body string
	duration              time.Duration
	calendar              *calc.PrayerTimes
	prayer                calc.Prayer
}

func NewNotificationBuilder() *NotificationBuilder {
	return &NotificationBuilder{}
}

// TODO: check for icon existence
func (b *NotificationBuilder) SetIconName(name string) *NotificationBuilder {
	b.iconName = name
	return b
}

// Sets the title for the notification, will be treated as template if prayer is set
func (b *NotificationBuilder) SetTitle(title string) *NotificationBuilder {
	b.title = title
	return b
}

// Sets the body for the notification, will be treated as template if prayer is set
func (b *NotificationBuilder) SetBody(body string) *NotificationBuilder {
	b.body = body
	return b
}

func (b *NotificationBuilder) SetDuration(duration time.Duration) *NotificationBuilder {
	b.duration = duration
	return b
}

func (b *NotificationBuilder) SetPrayer(
	calendar *calc.PrayerTimes,
	prayer calc.Prayer,
) *NotificationBuilder {
	b.calendar = calendar
	b.prayer = prayer
	return b
}

func (b *NotificationBuilder) renderTemplate(
	name, templateString string,
	data any,
) (string, error) {
	var builder strings.Builder
	var err error

	tmpl, err := template.New(name).Parse(templateString)
	if err != nil {
		return "", err
	}

	if err = tmpl.Execute(&builder, data); err != nil {
		return "", err
	}
	return builder.String(), nil
}

func (b *NotificationBuilder) Build() (Notification, error) {
	if b.prayer == calc.NO_PRAYER {
		return Notification{
			iconName: b.iconName,
			title:    b.title,
			body:     b.body,
			duration: b.duration,
		}, nil
	}

	// NOTE: render title and body templates
	var err error
	prayerData := struct {
		PrayerName, CalendarName string
	}{
		PrayerName:   common.PrayerName(b.prayer),
		CalendarName: common.CalendarName(*b.calendar, b.prayer),
	}

	notification := Notification{iconName: b.iconName, duration: b.duration}
	notification.title, err = b.renderTemplate("notification-title", b.title, prayerData)
	if err != nil {
		return Notification{}, fmt.Errorf("failed to execute title template: %v", err)
	}
	notification.body, err = b.renderTemplate("notification-body", b.body, prayerData)
	if err != nil {
		return Notification{}, fmt.Errorf("failed to execute body template: %v", err)
	}

	return notification, nil
}
