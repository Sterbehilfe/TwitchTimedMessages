package console

import (
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
)

func WriteLine(value string) {
	now := time.Now()
	builder := new(strings.Builder)
	if now.Hour() < 10 {
		builder.WriteString(fmt.Sprintf("0%d", now.Hour()))
	} else {
		builder.WriteString(fmt.Sprint(now.Hour()))
	}

	builder.WriteString(":")
	if now.Minute() < 10 {
		builder.WriteString(fmt.Sprintf("0%d", now.Minute()))
	} else {
		builder.WriteString(fmt.Sprint(now.Minute()))
	}

	builder.WriteString(":")
	if now.Second() < 10 {
		builder.WriteString(fmt.Sprintf("0%d", now.Second()))
	} else {
		builder.WriteString(fmt.Sprint(now.Second()))
	}

	builder.WriteString(" | ")
	builder.WriteString(value)

	color.Green("%s", builder.String())
}
