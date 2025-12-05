package main

import (
	"channels/logger"
	"channels/utils"
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Unbuffered Channel Example ===")
	unbufferedChannel()

	fmt.Println("\n=== Buffered Channel Example ===")
	bufferedChannel()
}

func unbufferedChannel() {
	c := make(chan string)
	log := logger.New()

	l := log.Child("main")
	go pusher("A", c, log.Child("push-A"))
	go pusher("B", c, log.Child("push-B"))

	for i := 0; i < 6; i++ {
		l.Log("Waiting for message...")
		msg := <-c
		l.Log(fmt.Sprintf("Received: %s", msg))

		duration, sleeper := utils.Sleep(100, 200)
		l.Log(fmt.Sprintf("Processing for %s...", duration))
		sleeper()
		l.Log("Done processing")
	}

	log.Flush("timeline_unbuffered.txt")
}

func bufferedChannel() {
	c := make(chan string, 3)
	log := logger.New()

	l := log.Child("main")
	go pusher("A", c, log.Child("push-A"))
	go pusher("B", c, log.Child("push-B"))

	for i := 0; i < 6; i++ {
		l.Log("Waiting for message...")
		msg := <-c
		l.Log(fmt.Sprintf("Received: %s", msg))

		duration, sleeper := utils.Sleep(100, 200)
		l.Log(fmt.Sprintf("Processing for %s...", duration))
		sleeper()
		l.Log("Done processing")
	}

	log.Flush("timeline_buffered.txt")
}

func pusher(name string, c chan string, l *logger.ChildLogger) {
	for i := 0; ; i++ {
		message := fmt.Sprintf("\"Pusher-%s #%d\"", name, i)

		duration, sleeper := utils.Sleep(10, 30)
		l.Log(fmt.Sprintf("Preparing %s for %s...", message, duration))
		sleeper()
		l.Log(fmt.Sprintf("Trying to send %s", message))

		start := time.Now()
		c <- message
		blocked := time.Since(start).Round(time.Millisecond)

		if blocked < time.Millisecond {
			l.Log(fmt.Sprintf("Sent %s (instant)", message))
		} else {
			l.Log(fmt.Sprintf("Sent %s (blocked for %s)", message, blocked))
		}
	}
}
