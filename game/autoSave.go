package game

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func initializeAutoSave(g *Game) {
	g.autoSaver.autosaveInterval = 10.0
	g.autoSaver.autosaveDelayAfterManualSave = 30.0
	g.autoSaver.currentAutosaveInterval = g.autoSaver.autosaveInterval
	g.autoSaver.autosaveTimer = g.autoSaver.currentAutosaveInterval
}

func autoSave(isInGameState bool, g *Game) {
	if isInGameState {
		// Log autosave timer in debug mode
		if g.state.isDebug {
			fmt.Printf("Autosave in: %.1f seconds\n", g.autoSaver.autosaveTimer)
		}

		// Update autosave timer
		g.autoSaver.autosaveTimer -= rl.GetFrameTime()
		if g.autoSaver.autosaveTimer <= 0 {
			err := g.saveGame()
			if err != nil {
				fmt.Printf("Error autosaving game: %v\n", err)
			} else {
				g.displayMessage("Game autosaved!", 2.0) // Display autosave message
			}
			g.autoSaver.autosaveTimer = g.autoSaver.currentAutosaveInterval
			g.autoSaver.currentAutosaveInterval = g.autoSaver.autosaveInterval // Reset to regular interval after any save
		}
	}
}
