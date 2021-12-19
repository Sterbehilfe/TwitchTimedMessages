package main

import (
	"fmt"

	console "github.com/sterbehilfe/TwitchTimedMessages/TwitchTimedMessages/Console"
	settings "github.com/sterbehilfe/TwitchTimedMessages/TwitchTimedMessages/Settings"
	twitch "github.com/sterbehilfe/TwitchTimedMessages/TwitchTimedMessages/Twitch"
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
