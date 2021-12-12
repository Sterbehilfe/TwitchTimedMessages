package settings

import (
	console "Documents/Git/TwitchTimedMessages/TwitchTimedMessages/Console"
	"encoding/json"
	"os"
)

type Settings struct {
	Username   string
	OAuthToken string
	Messages   []Message
}

func LoadSettings() Settings {
	var settings Settings
	content, err := os.ReadFile("./Settings.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(content, &settings)
	if err != nil {
		panic(err)
	}
	console.WriteLine("Loaded settings")
	return settings
}
