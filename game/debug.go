package game

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func drawBoundingBoxes(g *Game) {
	rl.DrawBoundingBox(g.boxes.PlayerBox, rl.Red)
	rl.DrawBoundingBox(g.boxes.GroundBox, rl.Green)
	rl.DrawBoundingBox(g.boxes.LeftWallBox, rl.Blue)
	rl.DrawBoundingBox(g.boxes.RightWallBox, rl.Blue)
}

func drawDebugInfo(g *Game) {
	rl.DrawText(fmt.Sprintf("Player: %.2f, %.2f, %.2f", g.gameObjects.player.Position.X, g.gameObjects.player.Position.Y, g.gameObjects.player.Position.Z), 10, 40, 18, rl.Red)
	rl.DrawText(fmt.Sprintf("Camera: %.2f, %.2f, %.2f", g.camera.Position.X, g.camera.Position.Y, g.camera.Position.Z), 10, 60, 18, rl.Red)
	rl.DrawText(fmt.Sprintf("Level: %d", g.state.Level), 10, 80, 18, rl.Red)
}
