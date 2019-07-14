package main

import (
	"sync"
	"time"
)

type TickCallback func(elapsed, target time.Duration)

type Stopwatch struct {
	tick TickCallback // callback fired once per second

	elapsed time.Duration // amount of time elapsed
	target  time.Duration // target total runtime

	ticker     *time.Ticker // periodic timer that updates state
	running    bool         // whether the timer is running or paused
	lastUpdate *time.Time
	lastTick *time.Time
	mutex      sync.Mutex
}

func NewStopwatch(tick TickCallback) *Stopwatch {
	stopwatch := &Stopwatch{
		tick:   tick,
		ticker: time.NewTicker(100 * time.Millisecond),
	}

	go func() {
		for {
			stopwatch.update()
		}
	}()

	return stopwatch
}

func (s *Stopwatch) update() {
	now := <-s.ticker.C

	s.mutex.Lock()
	defer s.mutex.Unlock()

	if !s.running {
		return
	}

	s.elapsed += now.Sub(*s.lastUpdate)
	s.lastUpdate = &now

	if now.Sub(*s.lastTick) >= 1 * time.Second {
		s.tick(s.elapsed, s.target)
		s.lastTick = &now
	}
}

func (s *Stopwatch) Start() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.unsafeStart()
}

func (s *Stopwatch) unsafeStart() {
	now := time.Now()
	s.lastUpdate = &now
	s.lastTick = &now

	s.tick(s.elapsed, s.target)

	s.running = true
}

func (s *Stopwatch) Pause() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.unsafePause()
}

func (s *Stopwatch) unsafePause() {
	s.running = false

	s.tick(s.elapsed, s.target)
}

func (s *Stopwatch) Reset() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.running = false

	s.elapsed = 0
	s.target = -1

	s.lastUpdate = nil
	s.lastTick = nil

	s.tick(s.elapsed, s.target)
}

func (s *Stopwatch) Toggle() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.running {
		s.unsafePause()
	} else {
		s.unsafeStart()
	}
}

func (s *Stopwatch) SetTarget(target time.Duration) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.target = target
}
