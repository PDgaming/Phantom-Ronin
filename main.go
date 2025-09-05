package main

import (
	"fmt"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = int32(800)
	screenHeight = int32(480)

	worldWidth  = float32(30)
	worldLength = float32(2)

	GRAVITY = -9.8
)

var (
	currentLevel     Level
	exitButton       bool
	startButton      bool
	transitionButton bool
)

type GameState struct {
	Level      int
	isSideView bool
	isDebug    bool
	menuState  string
}

func main() {
	// Init
	rl.InitWindow(screenWidth, screenHeight, "Phantom-Ronin")
	defer rl.CloseWindow()

	state := GameState{
		Level:      1,
		isSideView: true,
		isDebug:    false,
		menuState:  "startMenu",
	}

	// Camera
	camera := rl.Camera3D{}
	camera.Position = rl.NewVector3(0, 0, 10.0)
	camera.Target = rl.NewVector3(0.0, 0.0, 0.0)
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Fovy = 45.0
	camera.Projection = rl.CameraPerspective

	backgroundTexture := rl.LoadTexture("./assets/background.png")

	background := Background{
		Position:        rl.NewVector3(0.0, 0.0, -1.0),
		Height:          float32(screenHeight),
		Width:           worldWidth,
		Length:          0.1,
		Color:           rl.Blue,
		Texture:         backgroundTexture,
		TextureProvided: false,
	}

	ground := Ground{
		Position: rl.NewVector3(0.0, -2, 0.1),
		Height:   0.2,
		Width:    worldWidth,
		Length:   2.0,
		Color:    rl.Red,
	}

	player := Player{
		Position: rl.NewVector3(0.0, -1.0, 0.0),
		Width:    0.5,
		Height:   1.0,
		Length:   0.5,
		Color:    rl.Green,
	}

	leftWall := Wall{
		Position: rl.NewVector3(-0.5, -0.6, 0.0),
		Width:    0.5,
		Height:   3.0,
		Length:   2.2,
		Color:    rl.DarkBrown,
	}

	rightWall := Wall{
		Position: rl.NewVector3(worldWidth, -0.6, 0.0),
		Width:    0.5,
		Height:   3.0,
		Length:   2.2,
		Color:    rl.DarkBrown,
	}

	resetGame(&state, &player, &currentLevel)

	rl.SetTargetFPS(120)
	for !rl.WindowShouldClose() {
		if state.menuState == "inGame" || state.menuState == "gameOver" {
			if rl.IsKeyPressed(rl.KeyR) {
				state.isSideView = !state.isSideView
			}

			player.update(state.isSideView, &background, &ground)
		}

		playerBox := GetBoundingBox(player.Position, player.Width, player.Height, player.Length)
		groundBox := GetBoundingBox(ground.Position, ground.Width, ground.Height, ground.Length)
		leftWallBox := GetBoundingBox(leftWall.Position, leftWall.Width, leftWall.Height, leftWall.Length)
		rightWallBox := GetBoundingBox(rightWall.Position, rightWall.Width, rightWall.Height, rightWall.Length)

		if rl.CheckCollisionBoxes(playerBox, groundBox) {
			if player.Velocity.Y <= 0 {
				player.IsGrounded = true
				player.jumpsUsed = 0
				player.Velocity.Y = 0.0

				player.Position.Y = ground.Position.Y + (ground.Height / 2) + (player.Height / 2)
			} else {
				player.Velocity.Y = 0.0
				player.Position.Y = ground.Position.Y - (ground.Height / 2) - (player.Height / 2)
			}
		} else {
			player.IsGrounded = false
		}

		if rl.CheckCollisionBoxes(playerBox, leftWallBox) {
			player.Position.X = leftWall.Position.X + leftWall.Width/2 + player.Width/2
		}
		if rl.CheckCollisionBoxes(playerBox, rightWallBox) {
			player.Position.X = rightWall.Position.X - rightWall.Width/2 - player.Width/2
		}

		if state.isSideView {
			clampX := rl.Clamp(player.Position.X, 3.15, background.Width-3.7)
			clampY := rl.Clamp(player.Position.Y, 0.1, background.Height-player.Height)

			camera.Position = rl.NewVector3(clampX, clampY, 6)
			camera.Target = rl.NewVector3(clampX, clampY, player.Position.Z)
		} else {
			camera.Position = rl.NewVector3(player.Position.X+5, player.Position.Y+2, 4)
			camera.Target = rl.NewVector3(player.Position.X, player.Position.Y, 0)
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.NewColor(255, 182, 193, 255))

		rl.BeginMode3D(camera)

		if state.isDebug {
			rl.DrawBoundingBox(playerBox, rl.Red)
			rl.DrawBoundingBox(groundBox, rl.Green)
			rl.DrawBoundingBox(leftWallBox, rl.Blue)
			rl.DrawBoundingBox(rightWallBox, rl.Blue)
		}

		background.draw()
		ground.draw()
		leftWall.draw()
		rightWall.draw()

		for _, platform := range currentLevel.Platforms {
			platform.draw()

			platformBox := GetBoundingBox(platform.Position, platform.Width, platform.Height, platform.Length)

			if state.isDebug {
				rl.DrawBoundingBox(platformBox, rl.Red)
			}

			if rl.CheckCollisionBoxes(playerBox, platformBox) {
				// Only allow landing on top of the platform if falling down
				playerBottom := player.Position.Y - player.Height/2
				platformTop := platform.Position.Y + platform.Height/2
				platformBottom := platform.Position.Y - platform.Height/2

				// Check if player is above the platform and moving down
				if playerBottom >= platformTop-0.05 && player.Velocity.Y <= 0 {
					// Landing on top of the platform
					player.Position.Y = platformTop + player.Height/2
					player.IsGrounded = true
					player.jumpsUsed = 0
					player.Velocity.Y = 0.0

					if platform.final {
						fmt.Printf("Transitioning to Level %d\n", state.Level)
						state.menuState = "levelTransition"
					}
				} else if (player.Position.Y+player.Height/2) <= platformBottom+0.05 && player.Velocity.Y > 0 {
					// Hitting the platform from below while moving up
					player.Position.Y = platformBottom - player.Height/2
					player.Velocity.Y = 0.0
				} else {
					// Prevent horizontal movement through the platform sides
					playerTop := player.Position.Y + player.Height/2

					// Properly resolve horizontal collision by calculating overlap and moving player out by the minimal axis
					if playerTop > platformBottom && playerBottom < platformTop {
						if state.isSideView {
							// Calculate overlap on X axis
							playerLeft := player.Position.X - player.Width/2
							playerRight := player.Position.X + player.Width/2
							platformLeft := platform.Position.X - platform.Width/2
							platformRight := platform.Position.X + platform.Width/2

							overlapLeft := playerRight - platformLeft
							overlapRight := platformRight - playerLeft

							// Move player out by the minimal overlap
							if overlapLeft < overlapRight {
								player.Position.X -= overlapLeft
							} else {
								player.Position.X += overlapRight
							}
						} else {
							// Calculate overlap on Z axis
							playerFront := player.Position.Z + player.Length/2
							playerBack := player.Position.Z - player.Length/2
							platformFront := platform.Position.Z + platform.Length/2
							platformBack := platform.Position.Z - platform.Length/2

							overlapBack := playerFront - platformBack
							overlapFront := platformFront - playerBack

							// Move player out by the minimal overlap
							if overlapBack < overlapFront {
								player.Position.Z -= overlapBack
							} else {
								player.Position.Z += overlapFront
							}
						}
					}
				}
			}
		}

		player.draw()

		rl.EndMode3D()

		switch state.menuState {
		case "startMenu":
			rl.DrawText("Phanton Ronin", 80, 150, 80, rl.Red)
			startButton = gui.Button(rl.NewRectangle(float32(screenWidth)/2-50, 250, 100, 40), "Start")
			if startButton {
				state.menuState = "inGame"
				resetGame(&state, &player, &currentLevel)
			}
			exitButton = gui.Button(rl.NewRectangle(float32(screenWidth)/2-50, 300, 100, 40), "Exit")
			if exitButton {
				rl.CloseWindow()
			}
		case "levelTransition":
			rl.DrawText("Level Completed!", 80, 150, 80, rl.Red)
			transitionButton = gui.Button(rl.NewRectangle(float32(screenWidth)/2-50, 250, 100, 40), "Next")
			if transitionButton {
				state.menuState = "inGame"
				state.Level++
				resetGame(&state, &player, &currentLevel)
			}
		case "gameOver":
			rl.DrawText("Game Completed!", 70, 190, 80, rl.Red)
			exitButton = gui.Button(rl.NewRectangle(float32(screenWidth)/2-50, 280, 100, 40), "Exit")
			if exitButton {
				rl.CloseWindow()
			}
		}

		if state.isDebug {
			rl.DrawText(fmt.Sprintf("Player: %.2f, %.2f, %.2f", player.Position.X, player.Position.Y, player.Position.Z), 10, 40, 18, rl.Red)
			rl.DrawText(fmt.Sprintf("Camera: %.2f, %.2f, %.2f", camera.Position.X, camera.Position.Y, camera.Position.Z), 10, 60, 18, rl.Red)
			rl.DrawText(fmt.Sprintf("Level: %d", state.Level), 10, 80, 18, rl.Red)
		}

		rl.DrawText(fmt.Sprintf("Level: %d", state.Level), 10, 30, 18, rl.Orange)

		rl.DrawFPS(10, 10)
		rl.EndDrawing()
	}
	defer rl.UnloadTexture(backgroundTexture)
}

func resetGame(state *GameState, player *Player, currentLevel *Level) {
	// Reset player's position to the start of the new level
	// This is a placeholder; you might want to read the starting position from the level file
	player.Position = rl.NewVector3(0.0, -1.0, 0.0)
	player.Velocity = rl.NewVector3(0.0, 0.0, 0.0)
	player.IsGrounded = true
	player.jumpsUsed = 0
	currentLevel.resetLevel()

	// Load the new level
	switch state.Level {
	case 1:
		currentLevel.loadLevel("./level-maps/level1.csv")
	case 2:
		currentLevel.loadLevel("./level-maps/level2.csv")
	// case 3:
	//     currentLevel.loadLevel("./level-maps/level3.csv")
	default:
		fmt.Println("Game Completed!")
		state.Level = 0
		state.menuState = "gameOver"
	}
}
