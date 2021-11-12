package twitch

import (
	settings "Documents/Git/TwitchTimedMessages/TwitchTimedMessages/Settings"

	linq "github.com/ahmetb/go-linq"
	irc "github.com/gempir/go-twitch-irc"
)

type TwitchClient struct {
	_settings settings.Settings
	_client   *irc.Client
}

func NewTwitchClient(settings settings.Settings) *TwitchClient {
	return &TwitchClient{
		_settings: settings,
		_client:   irc.NewClient(settings.Username, settings.OAuthToken),
	}
}

func (client *TwitchClient) Initialize() {
	for _, channel := range client.GetChannels() {
		client._client.Join(channel)
	}
}

func (client *TwitchClient) GetChannels() []string {
	var result []string
	linq.From(client._settings).SelectT(func(m settings.Message) string {
		return m.Channel
	}).Distinct().ToSlice(&result)
	return result
}
