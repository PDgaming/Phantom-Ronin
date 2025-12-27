package game

import (
	"encoding/json"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type SaveData struct {
	Level      int
	Position   rl.Vector3
	Velocity   rl.Vector3
	IsGrounded bool
	JumpsUsed  int
}

func (g *Game) saveGame(filename string) error {
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

	return os.WriteFile(filename, data, 0644)
}

func (g *Game) loadGame(filename string) error {
	data, err := os.ReadFile(filename)
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
