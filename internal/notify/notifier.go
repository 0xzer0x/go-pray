package notify

type Notifier interface {
	Initialize() error
	Send(Notification) error
	SendInteractive(chan<- Result, Notification)
	Close() error
}
