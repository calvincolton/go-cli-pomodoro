//go:build inmemory
// +build inmemory

package pomodoro_test

import (
	"testing"

	"github.com/calvincolton/go-cli-pomodoro/pomodoro"
	"github.com/calvincolton/go-cli-pomodoro/pomodoro/repository"
)

func getRepo(t *testing.T) (pomodoro.Repository, func()) {
	t.Helper()

	// would normally return a cleanup function, but this is not necessary for in memory data
	return repository.NewInMemoryRepo(), func() {}
}
