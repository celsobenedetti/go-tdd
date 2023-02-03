package main

import (
	"os"
	"time"

	"github.com/celso-patiri/go-tdd/mocking"
)

func main() {
    sleeper := mocking.ConfigurableSleeper{}
    sleeper.SetDuration(time.Second)
    sleeper.SetSleep(time.Sleep)

    mocking.Countdown(os.Stdout, &sleeper)
}

