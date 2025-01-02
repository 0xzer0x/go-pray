package util

import (
	"fmt"
	"os/exec"
)

func SendNotification(icon, body string) error {
	var err error
	var notifyCmd string
	if notifyCmd, err = exec.LookPath("notify-send"); err != nil {
		return fmt.Errorf("command '%s' does not exist", "notify-send")
	}

	var cmd *exec.Cmd
	if len(icon) > 0 {
		cmd = exec.Command(notifyCmd, "--icon="+icon, "go-pray", body)
	} else {
		cmd = exec.Command(notifyCmd, "go-pray", body)
	}
	return cmd.Run()
}
