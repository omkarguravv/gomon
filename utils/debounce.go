package utils

import (
	"sync"
	"time"
)

type Debouncer struct {
	delay time.Duration
	timer *time.Timer
	lock  sync.Mutex
}

func NewDebouncer(delay time.Duration) *Debouncer {
	return &Debouncer{delay: delay}
}

func (d *Debouncer) Call(fn func()) {
	d.lock.Lock()
	defer d.lock.Unlock()

	if d.timer != nil {
		d.timer.Stop()
	}
	d.timer = time.AfterFunc(d.delay, fn)
}
