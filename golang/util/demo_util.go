package demo_util

import (
	"math/rand"
	"time"
)

func MockExecuteByRandTime(subject string, randScope int) {
	randTime := rand.Intn(randScope)
	time.Sleep(time.Duration(randTime) * time.Millisecond)
	println(subject, " execution time lasted ", randTime, " milliseconds")
}
