//go:build !containers && !disable_notification
// +build !containers,!disable_notification

package app

import notify "github.com/calvincolton/go-cli-notify"

func send_notification(msg string) {
	n := notify.New("Pomodoro", msg, notify.SeverityNormal)
	n.Send()
}
