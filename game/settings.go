package game

import (
	"encoding/json"
	"fmt"
	"os"
)

const settingsFilePath = "settings.json"

type Settings struct {
	ShowIntro bool `json:"showIntro"`
}

// saveSettings saves the current game settings to a JSON file.
func saveSettings(settings Settings) error {
	file, err := json.MarshalIndent(settings, "", " ")
	if err != nil {
		return fmt.Errorf("failed to marshal settings: %w", err)
	}

	err = os.WriteFile(settingsFilePath, file, 0644)
	if err != nil {
		return fmt.Errorf("failed to write settings file: %w", err)
	}

	return nil
}

// loadSettings loads game settings from a JSON file.
func loadSettings() (Settings, error) {
	var settings Settings

	file, err := os.ReadFile(settingsFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			// If file doesn't exist, return default settings
			return Settings{ShowIntro: true}, nil
		}
		return settings, fmt.Errorf("failed to read settings file: %w", err)
	}

	err = json.Unmarshal(file, &settings)
	if err != nil {
		return settings, fmt.Errorf("failed to unmarshal settings: %w", err)
	}

	return settings, nil
}
