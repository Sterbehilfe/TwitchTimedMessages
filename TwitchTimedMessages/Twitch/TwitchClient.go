package twitch

import (
	settings "Documents/Git/TwitchTimedMessages/TwitchTimedMessages/Settings"
	"fmt"
	"time"

	linq "github.com/ahmetb/go-linq"
	irc "github.com/gempir/go-twitch-irc"
)

type TwitchClient struct {
	_settings  settings.Settings
	_ircClient *irc.Client
	_timers    []time.Timer
}

func NewTwitchClient(settings settings.Settings) *TwitchClient {
	return &TwitchClient{
		_settings:  settings,
		_ircClient: irc.NewClient(settings.Username, settings.OAuthToken),
		_timers:    make([]time.Timer, len(settings.Messages)),
	}
}

func (client *TwitchClient) Initialize() {
	client.SetEvents()
	client.JoinChannels()
	err := client._ircClient.Connect()
	if err != nil {
		panic(err)
	}
}

func (client *TwitchClient) Send(message settings.Message) {
	client._ircClient.Say(message.Channel, message.Content)
	fmt.Println("Send message to " + message.Channel + ": " + message.Content)
}

func (client *TwitchClient) GetChannels() []string {
	var result []string
	linq.From(client._settings).SelectT(func(m settings.Message) string {
		return m.Channel
	}).Distinct().ToSlice(&result)
	return result
}

func (client *TwitchClient) JoinChannels() {
	for _, channel := range client.GetChannels() {
		client._ircClient.Join(channel)
	}
}

func (client *TwitchClient) SetEvents() {
	client._ircClient.OnConnect(func() {
		fmt.Println("Connected")
	})
}
