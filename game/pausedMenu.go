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

	loadButton := gui.Button(rl.NewRectangle(float32(screenWidth)/2-50, 300, 100, 40), "Load Game")
	if loadButton {
		rl.PlaySound(g.audioStreams.buttonSound)
		err := g.loadGame()
		if err != nil {
			fmt.Println("Could not load save game:", err)
			g.state.menuState = "intro"
		}
	}

	resetLevelButton := gui.Button(rl.NewRectangle(float32(screenWidth)/2-50, 350, 100, 40), "Reset Level")
	if resetLevelButton {
		rl.PlaySound(g.audioStreams.buttonSound)
		g.resetGame(g.state.Level)
		g.state.menuState = "inGame" // Optionally return to inGame state after reset
	}

	mainMenuButton := gui.Button(rl.NewRectangle(float32(screenWidth)/2-50, 400, 100, 40), "Main Menu")
	if mainMenuButton {
		rl.PlaySound(g.audioStreams.buttonSound)
		g.state.menuState = "startMenu"
	}
}
