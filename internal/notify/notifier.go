package notify

type Notifier interface {
	Initialize() error
	Send(chan<- Result, Notification)
	Close() error
}
