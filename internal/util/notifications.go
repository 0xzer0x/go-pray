package util

import "os/exec"

func SendNotification(body string) error {
	return exec.Command("notify-send", "go-pray", body).Run()
}
