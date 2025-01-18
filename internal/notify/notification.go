package notify

import "time"

type Notification struct {
	iconName, title, body string
	duration              time.Duration
}

type NotificationBuilder struct {
	iconName, title, body string
	duration              time.Duration
}

func NewNotificationBuilder() *NotificationBuilder {
	return &NotificationBuilder{}
}

// TODO: check for icon existence
func (b *NotificationBuilder) SetIconName(name string) *NotificationBuilder {
	b.iconName = name
	return b
}

func (b *NotificationBuilder) SetTitle(title string) *NotificationBuilder {
	b.title = title
	return b
}

func (b *NotificationBuilder) SetBody(body string) *NotificationBuilder {
	b.body = body
	return b
}

func (b *NotificationBuilder) SetDuration(duration time.Duration) *NotificationBuilder {
	b.duration = duration
	return b
}

func (b *NotificationBuilder) Build() Notification {
	return Notification{
		iconName: b.iconName,
		title:    b.title,
		body:     b.body,
		duration: b.duration,
	}
}
