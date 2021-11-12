package twitch

import (
	settings "Documents/Git/TwitchTimedMessages/TwitchTimedMessages/Settings"
)

type TwitchClient struct {
	_settings settings.Settings
}

func NewTwitchClient(settings settings.Settings) *TwitchClient {
	return &TwitchClient{
		_settings: settings,
	}
}

func (client TwitchClient) Initialize() {

}
