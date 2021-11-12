package twitch

import (
	settings "Documents/Git/TwitchTimedMessages/TwitchTimedMessages/Settings"
	"time"

	linq "github.com/ahmetb/go-linq"
	irc "github.com/gempir/go-twitch-irc"
)

type TwitchClient struct {
	_settings settings.Settings
	_client   *irc.Client
	_timers   []time.Timer
}

func NewTwitchClient(settings settings.Settings) *TwitchClient {
	return &TwitchClient{
		_settings: settings,
		_client:   irc.NewClient(settings.Username, settings.OAuthToken),
		_timers:   make([]time.Timer, len(settings.Messages)),
	}
}

func (client *TwitchClient) Initialize() {
	client.JoinChannels()
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
		client._client.Join(channel)
	}
}
