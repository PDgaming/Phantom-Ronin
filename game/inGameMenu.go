package game

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func drawInGameMenu(g *Game) {
	g.buttons.pauseButton = gui.Button(rl.NewRectangle(float32(screenWidth)-60, 10, 50, 30), "Pause")
	if g.buttons.pauseButton {
		rl.PlaySound(g.audioStreams.buttonSound)
		g.state.menuState = "paused"
	}
}
