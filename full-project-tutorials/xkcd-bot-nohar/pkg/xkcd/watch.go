package xkcd

import (
	"context"
	"log"
	"sync"
	"time"
)

// Callback is the type of functions that get called when a new comic is published
type Callback func(*Comic)

// A Watcher periodically polls xkcd.com and triggers callbacks when a new
// comic is published.
type Watcher struct {
	// latestNum is the number of the latest xkcd number
	latestNum int
	// callbacks are called when a new xkcd comic is detected
	callbacks []Callback
	// period is this watcher's watch period
	period time.Duration

	// Internal state
	mtx    sync.Mutex
	ctx    context.Context
	cancel context.CancelFunc
}

// NewWatcher creates a new watcher and initializes it.
// An error is returned if we fail at fetching the latest xkcd comic during
// initialization.
func NewWatcher(ctx context.Context, period time.Duration) (*Watcher, error) {
	latest, err := Get(Latest)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithCancel(ctx)
	w := &Watcher{
		latestNum: latest.Num,
		period:    time.Minute,
		ctx:       ctx,
		cancel:    cancel,
	}
	go w.poll()
	return w, nil
}

// Close stops the watcher
func (w *Watcher) Close() {
	w.cancel()
}

// LatestNum returns the latest xkcd number this watcher has seen
func (w *Watcher) LatestNum() int {
	w.mtx.Lock()
	defer w.mtx.Unlock()
	return w.latestNum
}

// Period returns the watcher's polling period
func (w *Watcher) Period() time.Duration {
	w.mtx.Lock()
	defer w.mtx.Unlock()
	return w.period
}

// SetPeriod modifies the watcher's polling period
func (w *Watcher) SetPeriod(p time.Duration) {
	w.mtx.Lock()
	defer w.mtx.Unlock()
	w.period = p
}

// AddCallback adds a callback function to be executed when new comics are out.
func (w *Watcher) AddCallback(cbk Callback) {
	w.mtx.Lock()
	defer w.mtx.Unlock()

	if w.callbacks == nil {
		w.callbacks = make([]Callback, 0, 1)
	}
	w.callbacks = append(w.callbacks, cbk)
}

func (w *Watcher) poll() {
	for {
		select {
		case <-w.ctx.Done():
			log.Printf("watcher closed")
			return
		case <-time.After(w.Period()):
			w.checkUpdates()
		}
	}
}

func (w *Watcher) checkUpdates() {
	c, err := Get(Latest)
	if err != nil {
		log.Println("while polling xkcd for updates:", err)
		return
	}

	w.mtx.Lock()
	defer w.mtx.Unlock()

	var lastSeen int
	w.latestNum, lastSeen = c.Num, w.latestNum
	if w.latestNum != lastSeen {
		log.Printf("new xkcd: #%d (%s)", c.Num, c.SafeTitle)
		for _, cbk := range w.callbacks {
			// Copy comic contents to avoid side effects
			tmp := *c
			cbk(&tmp)
		}
	}
}
