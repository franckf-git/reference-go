package xkcd

import (
	"context"
	"sync"
	"time"
)

// DefaultWatchPeriod is the default value used for the global watcher.
const DefaultWatchPeriod = 10 * time.Minute

var (
	globalWatcher *Watcher
	once          sync.Once
)

// W returns the global watcher, initializing it if necessary with the default
// period.
// This function panics if initialization fails.
func W() *Watcher {
	once.Do(func() {
		var err error
		globalWatcher, err = NewWatcher(context.Background(), DefaultWatchPeriod)
		if err != nil {
			panic(err)
		}
	})
	return globalWatcher
}

// SetWatchPeriod sets the poll period
func SetWatchPeriod(p time.Duration) {
	W().SetPeriod(p)
}

// OnNewComic adds a function to be called when a new xkcd comic is detected.
func OnNewComic(cbk Callback) {
	W().AddCallback(cbk)
}

// LatestNum returns the latest xkcd number.
func LatestNum() int {
	return W().LatestNum()
}
