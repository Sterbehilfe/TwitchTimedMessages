package settings

import (
	"encoding/json"
	"os"

	console "github.com/sterbehilfe/TwitchTimedMessages/TwitchTimedMessages/Console"
)

type Settings struct {
	Username   string
	OAuthToken string
	Messages   []Message
}

func LoadSettings() Settings {
	content, err := os.ReadFile("./Settings.json")
	if err != nil {
		panic(err)
	}
	var settings Settings
	err = json.Unmarshal(content, &settings)
	if err != nil {
		panic(err)
	}
	console.WriteLine("Loaded settings")
	return settings
}
