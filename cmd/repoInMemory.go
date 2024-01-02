//go:build inmemory || containers
// +build inmemory containers

package cmd

import (
	"github.com/calvincolton/go-cli-pomodoro/pomodoro"
	"github.com/calvincolton/go-cli-pomodoro/pomodoro/repository"
)

func getRepo() (pomodoro.Repository, error) {
	return repository.NewInMemoryRepo(), nil
}
