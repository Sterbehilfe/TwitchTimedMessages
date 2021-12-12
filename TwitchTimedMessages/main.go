package main

import (
	console "Documents/Git/TwitchTimedMessages/TwitchTimedMessages/Console"
	settings "Documents/Git/TwitchTimedMessages/TwitchTimedMessages/Settings"
	twitch "Documents/Git/TwitchTimedMessages/TwitchTimedMessages/Twitch"
	"fmt"
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
