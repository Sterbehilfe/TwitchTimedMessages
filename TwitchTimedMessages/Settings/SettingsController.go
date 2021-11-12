package settings

import (
	"encoding/json"
	"fmt"
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
		fmt.Println(err)
	}
	err = json.Unmarshal(content, &settings)
	if err != nil {
		fmt.Println(err)
	}
	return settings
}
