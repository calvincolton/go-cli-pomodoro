//go:build !inmemory && !containers
// +build !inmemory,!containers

package cmd

import (
	"github.com/calvincolton/go-cli-pomodoro/pomodoro"
	"github.com/calvincolton/go-cli-pomodoro/pomodoro/repository"
	"github.com/spf13/viper"
)

func getRepo() (pomodoro.Repository, error) {
	repo, err := repository.NewSQLite3Repo(viper.GetString("db"))
	if err != nil {
		return nil, err
	}

	return repo, nil
}
