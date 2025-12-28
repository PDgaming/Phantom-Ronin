package game

import (
	"fmt"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func drawPausedMenu(g *Game) {
	g.buttons.startButton = gui.Button(rl.NewRectangle(float32(screenWidth)/2-50, 200, 100, 40), "Resume")
	if g.buttons.startButton {
		rl.PlaySound(g.audioStreams.buttonSound)
		g.state.menuState = "inGame"
	}

	saveButton := gui.Button(rl.NewRectangle(float32(screenWidth)/2-50, 250, 100, 40), "Save")
	if saveButton {
		rl.PlaySound(g.audioStreams.buttonSound)
		err := g.saveGame()
		if err != nil {
			fmt.Println("Error saving game:", err)
		} else {
			fmt.Println("Game saved successfully")
			g.autoSaver.currentAutosaveInterval = g.autoSaver.autosaveDelayAfterManualSave // Set delay for next autosave
			g.autoSaver.autosaveTimer = g.autoSaver.currentAutosaveInterval                // Reset timer with new interval
		}
	}

	resetLevelButton := gui.Button(rl.NewRectangle(float32(screenWidth)/2-50, 300, 100, 40), "Reset Level")
	if resetLevelButton {
		rl.PlaySound(g.audioStreams.buttonSound)
		g.resetGame(g.state.Level)
		g.state.menuState = "inGame" // Optionally return to inGame state after reset
	}

	exitButton := gui.Button(rl.NewRectangle(float32(screenWidth)/2-50, 350, 100, 40), "Exit")
	if exitButton {
		rl.PlaySound(g.audioStreams.buttonSound)
		rl.CloseWindow()
	}
}
