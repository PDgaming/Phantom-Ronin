package game

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) drawUI() {
	switch g.state.menuState {
	case "intro":
		g.introDialogueManager.Draw()
		if g.introDialogueManager.Update(g.audioStreams.buttonSound) {
			g.state.menuState = "inGame"
			g.introDialogueManager.Reset()
			g.resetGame(g.state.Level)
		}
	case "startMenu":
		DrawStartMenu(g, screenWidth, screenHeight)
	case "settings":
		drawSettingsMenu(g)
	case "inGame":
		drawInGameMenu(g)
	case "paused":
		drawPausedMenu(g)
	case "levelTransition":
		rl.DrawText("Level Completed!", 80, 150, 80, rl.Red)
		g.buttons.transitionButton = gui.Button(rl.NewRectangle(float32(screenWidth)/2-50, 250, 100, 40), "Next")
		if g.buttons.transitionButton {
			rl.PlaySound(g.audioStreams.buttonSound)
			g.state.menuState = "inGame"
			g.state.Level++
			g.resetGame(g.state.Level)
		}
	case "gameOver":
		rl.DrawText("Game Completed!", 70, 190, 80, rl.Red)
		g.buttons.exitButton = gui.Button(rl.NewRectangle(float32(screenWidth)/2-50, 280, 100, 40), "Exit")
		if g.buttons.exitButton {
			rl.PlaySound(g.audioStreams.buttonSound)
			rl.CloseWindow()
		}
	}

	if g.displayMessageTimer > 0 && g.displayMessageText != "" {
		textWidth := rl.MeasureText(g.displayMessageText, 20)
		rl.DrawText(g.displayMessageText, int32(screenWidth)/2-textWidth/2, int32(screenHeight)-50, 20, rl.Green)
	}
}
