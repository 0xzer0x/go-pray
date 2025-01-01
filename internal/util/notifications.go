package util

import "os/exec"

func SendNotification(body, icon string) error {
	var cmd *exec.Cmd
	if len(icon) > 0 {
		cmd = exec.Command("notify-send", "--icon="+icon, "go-pray", body)
	} else {
		cmd = exec.Command("notify-send", "go-pray", body)
	}
	return cmd.Run()
}
