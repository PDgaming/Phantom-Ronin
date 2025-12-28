package game

import (
	"fmt"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawStartMenu(g *Game, screenWidth int32, screenHeight int32) {
	rl.DrawText("Phanton Ronin", 80, 150, 80, rl.Red)
	g.buttons.startButton = gui.Button(rl.NewRectangle(float32(screenWidth)/2-50, 250, 100, 40), "Start")
	if g.buttons.startButton {
		rl.PlaySound(g.audioStreams.buttonSound)
		if g.state.ShowIntro {
			g.state.menuState = "intro"
		} else {
			g.state.menuState = "inGame"
			g.resetGame(g.state.Level)
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

	settingsButton := gui.Button(rl.NewRectangle(float32(screenWidth)/2-50, 350, 100, 40), "Settings")
	if settingsButton {
		rl.PlaySound(g.audioStreams.buttonSound)
		g.state.menuState = "settings"
	}

	g.buttons.exitButton = gui.Button(rl.NewRectangle(float32(screenWidth)/2-50, 400, 100, 40), "Exit")
	if g.buttons.exitButton {
		rl.PlaySound(g.audioStreams.buttonSound)
		rl.CloseWindow()
	}
}
