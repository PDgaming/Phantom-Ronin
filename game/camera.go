package game

import rl "github.com/gen2brain/raylib-go/raylib"

func (g *Game) initializeCamera() {
	g.camera = rl.Camera3D{
		Position:   rl.NewVector3(8.5, 0, 2),
		Target:     rl.NewVector3(8.5, 0.0, 0.0),
		Up:         rl.NewVector3(0.0, 1.0, 0.0),
		Fovy:       10.0,
		Projection: rl.CameraOrthographic,
	}
}

func (g *Game) updateCamera() {
	if g.state.isSideView {
		g.camera.Projection = rl.CameraOrthographic
		g.camera.Fovy = 10

		clampX := rl.Clamp(g.gameObjects.player.Position.X, 8.5, g.gameObjects.background.Width-8.6)
		clampY := rl.Clamp(g.gameObjects.player.Position.Y, 0.1, g.gameObjects.background.Height-g.gameObjects.player.Height)

		g.camera.Position = rl.NewVector3(clampX, clampY, 2)
		g.camera.Target = rl.NewVector3(clampX, clampY, g.gameObjects.player.Position.Z)

		g.gameObjects.player.SPEED = 8.0
	} else {
		g.camera.Projection = rl.CameraPerspective
		g.camera.Fovy = 45.0

		g.camera.Position = rl.NewVector3(g.gameObjects.player.Position.X+5, g.gameObjects.player.Position.Y+2, 4)
		g.camera.Target = rl.NewVector3(g.gameObjects.player.Position.X, g.gameObjects.player.Position.Y, 0)

		g.gameObjects.player.SPEED = 4.0
	}
}
