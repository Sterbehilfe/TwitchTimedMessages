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
	if now.Day() < 10 {
		builder.WriteString(fmt.Sprintf("0%d", now.Day()))
	} else {
		builder.WriteString(fmt.Sprint(now.Day()))
	}
	if now.Month() < 10 {
		builder.WriteString(fmt.Sprintf("0%d", now.Month()))
	} else {
		builder.WriteString(fmt.Sprint(now.Month()))
	}
	builder.WriteString(fmt.Sprint(now.Year()))
	if now.Hour() < 10 {
		builder.WriteString(fmt.Sprintf("0%d", now.Hour()))
	} else {
		builder.WriteString(fmt.Sprint(now.Hour()))
	}
	if now.Minute() < 10 {
		builder.WriteString(fmt.Sprintf("0%d", now.Minute()))
	} else {
		builder.WriteString(fmt.Sprint(now.Minute()))
	}
	if now.Second() < 10 {
		builder.WriteString(fmt.Sprintf("0%d", now.Second()))
	} else {
		builder.WriteString(fmt.Sprint(now.Second()))
	}
	builder.WriteString(" | ")
	builder.WriteString(value)

	color.Green("%s", builder.String())
}
