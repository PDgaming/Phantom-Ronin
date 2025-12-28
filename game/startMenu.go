package game

import (
	"fmt"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawStartMenu(g *Game, screenWidth int32, screenHeight int32) {
	rl.DrawText("Phanton Ronin", 80, 150, 80, rl.Red)

	// Button properties
	buttonWidth := float32(100)
	buttonHeight := float32(40)
	buttonX := float32(screenWidth)/2 - buttonWidth/2
	startY := float32(250) // Initial Y position for the first button
	buttonSpacing := float32(50)

	saveFileExists := CheckSaveFileExists()

	var firstButtonText string
	if saveFileExists {
		firstButtonText = "Continue"
	} else {
		firstButtonText = "Start"
	}

	// First button (Start/Continue)
	g.buttons.startButton = gui.Button(rl.NewRectangle(buttonX, startY, buttonWidth, buttonHeight), firstButtonText)
	if g.buttons.startButton {
		rl.PlaySound(g.audioStreams.buttonSound)
		if saveFileExists {
			err := g.loadGame()
			if err != nil {
				fmt.Println("Could not load save game:", err)
				// Fallback to starting a new game or showing an error
				g.state.menuState = "intro"
			}
		} else {
			if g.state.ShowIntro {
				g.state.menuState = "intro"
			} else {
				g.state.menuState = "inGame"
				g.resetGame(g.state.Level)
			}
		}
	}
	startY += buttonSpacing // Move Y position for the next button

	// Settings Button
	settingsButton := gui.Button(rl.NewRectangle(buttonX, startY, buttonWidth, buttonHeight), "Settings")
	if settingsButton {
		rl.PlaySound(g.audioStreams.buttonSound)
		g.state.menuState = "settings"
	}
	startY += buttonSpacing // Move Y position for the next button

	// Exit Button
	g.buttons.exitButton = gui.Button(rl.NewRectangle(buttonX, startY, buttonWidth, buttonHeight), "Exit")
	if g.buttons.exitButton {
		rl.PlaySound(g.audioStreams.buttonSound)
		rl.CloseWindow()
	}
}
