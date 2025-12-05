package logger

import (
	"channels/utils"
	"fmt"
	"strings"
	"time"
)

var colors = []string{
	"\033[36m", // Cyan
	"\033[33m", // Yellow
	"\033[35m", // Magenta
	"\033[32m", // Green
	"\033[34m", // Blue
	"\033[31m", // Red
}

var colorIndex = 0

const colorReset = "\033[0m"

type ConsoleLogger struct {
	color string
}

func newConsoleLogger() *ConsoleLogger {
	color := colors[colorIndex%len(colors)]
	colorIndex++
	return &ConsoleLogger{color: color}
}

func (c *ConsoleLogger) log(name, message string, elapsed time.Duration) {
	label := utils.Center(strings.ToUpper(name), 6)
	display := elapsed.Round(time.Millisecond) // Round for display only
	fmt.Printf("%s%8s [%s] - %s%s\n", c.color, display, label, message, colorReset)
}
