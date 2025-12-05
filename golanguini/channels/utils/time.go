package utils

import (
	"math/rand"
	"time"
)

func Sleep(minMs, maxMs int) (time.Duration, func()) {
	duration := time.Duration(minMs+rand.Intn(maxMs-minMs)) * time.Millisecond
	return duration, func() { time.Sleep(duration) }
}
