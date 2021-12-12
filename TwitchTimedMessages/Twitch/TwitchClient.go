package twitch

import (
	console "Documents/Git/TwitchTimedMessages/TwitchTimedMessages/Console"
	settings "Documents/Git/TwitchTimedMessages/TwitchTimedMessages/Settings"
	"fmt"
	"time"

	linq "github.com/ahmetb/go-linq"
	color "github.com/fatih/color"
	irc "github.com/gempir/go-twitch-irc/v2"
)

type TwitchClient struct {
	_settings  settings.Settings
	_ircClient *irc.Client
}

func NewTwitchClient(settings settings.Settings) *TwitchClient {
	return &TwitchClient{
		_settings:  settings,
		_ircClient: irc.NewClient(settings.Username, settings.OAuthToken),
	}
}

func (client *TwitchClient) Initialize() {
	client.CheckMessagesForRateLimiting()
	client.SetEvents()
	client.JoinChannels()
	go client._ircClient.Connect()
	time.Sleep(time.Duration(5000) * time.Millisecond)
	client.CreateTimers()
}

func (client *TwitchClient) Send(message settings.Message) {
	client._ircClient.Say(message.Channel, message.Content)
	console.WriteLine("Sent message to <#" + message.Channel + ">: " + message.Content)
}

func (client *TwitchClient) GetChannels() []string {
	var result []string
	linq.From(client._settings.Messages).SelectT(func(m settings.Message) string {
		return m.Channel
	}).Distinct().ToSlice(&result)
	return result
}

func (client *TwitchClient) JoinChannels() {
	for _, channel := range client.GetChannels() {
		client._ircClient.Join(channel)
		console.WriteLine("Joined channel <#" + channel + ">")
	}
}

func (client *TwitchClient) SetEvents() {
	client._ircClient.OnConnect(func() {
		console.WriteLine("Client connected")
	})
}

func (client *TwitchClient) CreateTimers() {
	for _, m := range client._settings.Messages {
		ticker := time.NewTicker(time.Duration(m.Interval) * time.Millisecond)
		go func(m settings.Message) {
			client.Send(m)
			for {
				<-ticker.C
				client.Send(m)
			}
		}(m)
	}
}

func (client *TwitchClient) CheckMessagesForRateLimiting() {
	var intervalTooSmall []settings.Message
	linq.From(client._settings.Messages).WhereT(func(m settings.Message) bool {
		return m.Interval < 1300
	}).ToSlice(&intervalTooSmall)
	for _, m := range intervalTooSmall {
		color.Red("[WARNING] Interval for", m.Channel+":", `"`+m.Content+`"`, "too small.")
	}
	if len(intervalTooSmall) > 0 {
		fmt.Println("The interval shouldn't be smaller than 1300ms")
	}
	var messageTooLong []settings.Message
	linq.From(client._settings.Messages).WhereT(func(m settings.Message) bool {
		return len(m.Content) > 500
	}).ToSlice(&messageTooLong)
	for _, m := range messageTooLong {
		color.Red("[WARNING] Message for", m.Channel+":", `"`+m.Content+`"`, "too long.")
	}
	if len(messageTooLong) > 0 {
		fmt.Println("The length of the message can't exceed 500 chars")
	}
}
