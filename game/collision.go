package game

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func GetBoundingBox(position rl.Vector3, width, height, length float32) rl.BoundingBox {
	halfWidth := width / 2.0
	halfHeight := height / 2.0
	halfLength := length / 2.0

	min := rl.NewVector3(
		position.X-halfWidth,
		position.Y-halfHeight,
		position.Z-halfLength,
	)

	max := rl.NewVector3(
		position.X+halfWidth,
		position.Y+halfHeight,
		position.Z+halfLength,
	)

	return rl.BoundingBox{Min: min, Max: max}
}

func initializeBoundingBoxes(g *Game) {
	g.boxes.PlayerBox = GetBoundingBox(g.gameObjects.player.Position, g.gameObjects.player.Width, g.gameObjects.player.Height, g.gameObjects.player.Length)
	g.boxes.GroundBox = GetBoundingBox(g.gameObjects.ground.Position, g.gameObjects.ground.Width, g.gameObjects.ground.Height, g.gameObjects.ground.Length)
	g.boxes.LeftWallBox = GetBoundingBox(g.gameObjects.leftWall.Position, g.gameObjects.leftWall.Width, g.gameObjects.leftWall.Height, g.gameObjects.leftWall.Length)
	g.boxes.RightWallBox = GetBoundingBox(g.gameObjects.rightWall.Position, g.gameObjects.rightWall.Width, g.gameObjects.rightWall.Height, g.gameObjects.rightWall.Length)
}

func checkCollision(g *Game) {
	if rl.CheckCollisionBoxes(g.boxes.PlayerBox, g.boxes.GroundBox) {
		if g.gameObjects.player.Velocity.Y <= 0 {
			g.gameObjects.player.IsGrounded = true
			g.gameObjects.player.jumpsUsed = 0
			g.gameObjects.player.Velocity.Y = 0.0

			g.gameObjects.player.Position.Y = g.gameObjects.ground.Position.Y + (g.gameObjects.ground.Height / 2) + (g.gameObjects.player.Height / 2)
		} else {
			g.gameObjects.player.Velocity.Y = 0.0
			g.gameObjects.player.Position.Y = g.gameObjects.ground.Position.Y - (g.gameObjects.ground.Height / 2) - (g.gameObjects.player.Height / 2)
		}
	} else {
		g.gameObjects.player.IsGrounded = false
	}

	if rl.CheckCollisionBoxes(g.boxes.PlayerBox, g.boxes.LeftWallBox) {
		g.gameObjects.player.Position.X = g.gameObjects.leftWall.Position.X + g.gameObjects.leftWall.Width/2 + g.gameObjects.player.Width/2
	}
	if rl.CheckCollisionBoxes(g.boxes.PlayerBox, g.boxes.RightWallBox) {
		g.gameObjects.player.Position.X = g.gameObjects.rightWall.Position.X - g.gameObjects.rightWall.Width/2 - g.gameObjects.player.Width/2
	}
}

func checkCollisionWithPlatforms(g *Game) {
	for _, platform := range g.currentLevel.Platforms {
		platformBox := GetBoundingBox(platform.Position, platform.Width, platform.Height, platform.Length)

		if rl.CheckCollisionBoxes(g.boxes.PlayerBox, platformBox) {
			playerBottom := g.gameObjects.player.Position.Y - g.gameObjects.player.Height/2
			platformTop := platform.Position.Y + platform.Height/2
			platformBottom := platform.Position.Y - platform.Height/2

			if playerBottom >= platformTop-0.05 && g.gameObjects.player.Velocity.Y <= 0 {
				g.gameObjects.player.Position.Y = platformTop + g.gameObjects.player.Height/2
				g.gameObjects.player.IsGrounded = true
				g.gameObjects.player.jumpsUsed = 0
				g.gameObjects.player.Velocity.Y = 0.0

				if platform.final {
					if g.state.menuState != "levelTransition" {
						rl.PlaySound(g.audioStreams.winSound)
						if g.state.isDebug {
							fmt.Printf("Transitioning to Level %d\n", g.state.Level)
						}
						g.state.menuState = "levelTransition"
					}
				}
			} else if (g.gameObjects.player.Position.Y+g.gameObjects.player.Height/2) <= platformBottom+0.05 && g.gameObjects.player.Velocity.Y > 0 {
				g.gameObjects.player.Position.Y = platformBottom - g.gameObjects.player.Height/2
				g.gameObjects.player.Velocity.Y = 0.0
			} else {
				playerTop := g.gameObjects.player.Position.Y + g.gameObjects.player.Height/2

				if playerTop > platformBottom && platformBottom < platformTop {
					if g.state.isSideView {
						playerLeft := g.gameObjects.player.Position.X - g.gameObjects.player.Width/2
						playerRight := g.gameObjects.player.Position.X + g.gameObjects.player.Width/2
						platformLeft := platform.Position.X - platform.Width/2
						platformRight := platform.Position.X + platform.Width/2

						overlapLeft := playerRight - platformLeft
						overlapRight := platformRight - playerLeft

						if overlapLeft < overlapRight {
							g.gameObjects.player.Position.X -= overlapLeft
						} else {
							g.gameObjects.player.Position.X += overlapRight
						}
					} else {
						playerFront := g.gameObjects.player.Position.Z + g.gameObjects.player.Length/2
						playerBack := g.gameObjects.player.Position.Z - g.gameObjects.player.Length/2
						platformFront := platform.Position.Z + platform.Length/2
						platformBack := platform.Position.Z - platform.Length/2

						overlapBack := playerFront - platformBack
						overlapFront := platformFront - playerBack

						if overlapBack < overlapFront {
							g.gameObjects.player.Position.Z -= overlapBack
						} else {
							g.gameObjects.player.Position.Z += overlapFront
						}
					}
				}
			}
		}
	}
}
