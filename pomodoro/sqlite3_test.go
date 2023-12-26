//go:build !inmemory
// +build !inmemory

package pomodoro_test

import (
	"os"
	"testing"

	"github.com/calvincolton/go-cli-pomodoro/pomodoro"
	"github.com/calvincolton/go-cli-pomodoro/pomodoro/repository"
)

func getRepo(t *testing.T) (pomodoro.Repository, func()) {
	t.Helper()

	tf, err := os.CreateTemp("", "pomo")
	if err != nil {
		t.Fatal(err)
	}
	tf.Close()

	dbRepo, err := repository.NewSQLite3Repo(tf.Name())
	if err != nil {
		t.Fatal(err)
	}

	return dbRepo, func() {
		os.Remove(tf.Name())
	}
}
