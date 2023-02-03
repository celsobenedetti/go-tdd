package mocking

import (
	"fmt"
	"io"
	"time"
)

const (
	finalWord           = "Go!"
	countdownIterations = 3
)

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownIterations; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	fmt.Fprintln(out, finalWord)
}

type Sleeper interface {
	Sleep()
}

type sleeperFunc = func(time.Duration)

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    sleeperFunc
}

func (s *ConfigurableSleeper) SetDuration(d time.Duration)  {
    s.duration = d
}

func (s *ConfigurableSleeper) SetSleep(sleep sleeperFunc)  {
    s.sleep = sleep
}


func (s *ConfigurableSleeper) Sleep()  {
    s.sleep(s.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept += duration
}
