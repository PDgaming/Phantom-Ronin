package game

import (
	"encoding/json"
	"os"
	"path/filepath"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const saveDir = "save"

type SaveData struct {
	Level      int
	Position   rl.Vector3
	Velocity   rl.Vector3
	IsGrounded bool
	JumpsUsed  int
}

type SettingsData struct {
	ShowIntro bool
}

func ensureSaveDirectory() error {
	if _, err := os.Stat(saveDir); os.IsNotExist(err) {
		return os.MkdirAll(saveDir, 0755)
	}
	return nil
}

func (g *Game) saveGame() error {
	if err := ensureSaveDirectory(); err != nil {
		return err
	}

	saveData := SaveData{
		Level:      g.state.Level,
		Position:   g.player.Position,
		Velocity:   g.player.Velocity,
		IsGrounded: g.player.IsGrounded,
		JumpsUsed:  g.player.jumpsUsed,
	}

	data, err := json.MarshalIndent(saveData, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filepath.Join(saveDir, "savegame.json"), data, 0644)
}

func (g *Game) loadGame() error {
	data, err := os.ReadFile(filepath.Join(saveDir, "savegame.json"))
	if err != nil {
		return err
	}

	var saveData SaveData
	if err := json.Unmarshal(data, &saveData); err != nil {
		return err
	}

	g.resetGame(saveData.Level) // reset to the level from the save

	// After resetting, apply the saved player state
	g.player.Position = saveData.Position
	g.player.Velocity = saveData.Velocity
	g.player.IsGrounded = saveData.IsGrounded
	g.player.jumpsUsed = saveData.JumpsUsed

	g.state.menuState = "inGame"

	return nil
}

func SaveSettings(settings SettingsData) error {
	if err := ensureSaveDirectory(); err != nil {
		return err
	}

	data, err := json.MarshalIndent(settings, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filepath.Join(saveDir, "settings.json"), data, 0644)
}

func LoadSettings() (SettingsData, error) {
	data, err := os.ReadFile(filepath.Join(saveDir, "settings.json"))
	if err != nil {
		return SettingsData{ShowIntro: true}, err // Default to true if file not found
	}

	var settings SettingsData
	if err := json.Unmarshal(data, &settings); err != nil {
		return SettingsData{ShowIntro: true}, err // Default to true if unmarshal fails
	}
	return settings, nil
}
