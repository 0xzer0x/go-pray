package notify

import (
	"bytes"
	"fmt"
	"os/exec"
	"time"
)

func Send(icon, body string, duration time.Duration) error {
	var err error
	var notifyCmd string
	if notifyCmd, err = exec.LookPath("notify-send"); err != nil {
		return fmt.Errorf("command '%s' does not exist", "notify-send")
	}

	var cmd *exec.Cmd
	var args []string = make([]string, 0, 5)
	if len(icon) > 0 {
		args = append(args, "--icon="+icon)
	}
	if duration > 0 {
		args = append(args, fmt.Sprintf("--expire-time=%d", duration.Milliseconds()))
	}
	args = append(args, "go-pray", body)

	cmd = exec.Command(notifyCmd, args...)
	return cmd.Run()
}

func SendInteractive(resChan chan Result, icon, body string, duration time.Duration) {
	defer close(resChan)
	var err error
	var notifyCmd string
	if notifyCmd, err = exec.LookPath("notify-send"); err != nil {
		resChan <- Result{
			Error: fmt.Errorf("command '%s' does not exist", "notify-send"),
		}
		return
	}

	var cmd *exec.Cmd
	var args []string = make([]string, 0, 6)
	if len(icon) > 0 {
		args = append(args, "--icon="+icon)
	}
	if duration > 0 {
		args = append(args, fmt.Sprintf("--expire-time=%d", duration.Milliseconds()))
	}
	args = append(args, "--action=clicked=trigger", "go-pray", body)

	cmd = exec.Command(notifyCmd, args...)
	output, err := cmd.Output()
	if err != nil {
		resChan <- Result{
			Error: fmt.Errorf("failed to execute command: %v", err),
		}
		return
	}

	resChan <- Result{
		Clicked: string(bytes.TrimSpace(output)) == "clicked",
		Error:   nil,
	}
}
