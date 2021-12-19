package main

import (
	"fmt"

	console "TwitchTimedMessages/Console"
	settings "TwitchTimedMessages/Settings"
	twitch "TwitchTimedMessages/Twitch"
)

func main() {
	console.WriteLine("TwitchTimedMessages started")
	settings := settings.LoadSettings()
	twitchClient := twitch.NewTwitchClient(settings)
	twitchClient.Initialize()

	for {
		var input string
		fmt.Scan(&input)
	}
}
