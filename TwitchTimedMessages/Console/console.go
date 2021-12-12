package console

import (
	"time"

	color "github.com/fatih/color"
)

func WriteLine(value string) {
	now := time.Now()
	color.Green("%d.%d.%d %d:%d:%d | %s", now.Day(), now.Month(), now.Year(), now.Hour(), now.Minute(), now.Second(), value)
}
