package main

import (
	"fmt"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = int32(800)
	screenHeight = int32(480)

	worldHeight = float32(15)
	worldWidth  = float32(30)
	worldLength = float32(2)

	GRAVITY = -9.8
)

var (
	currentLevel     Level
	exitButton       bool
	startButton      bool
	transitionButton bool
	state            GameState
	camera           rl.Camera3D
)

var introDialogues = []string{
	"Kaito: My village didn't fall to swords or steel. It fell to\na whisper of black magic...",
	"Kaito: The Mage sits in his high tower, weaving a spell\nthat anchors our world to the void.",
	"Kaito: But the path to him is guarded by the Magical\nMountains...",
	"Kaito: To most, this world is flat. A simple path of left\nand right.",
	"Kaito: But I carry the Gaze of the Void. I can see the\nworld as it truly is.",
	"Kaito: When a wall blocks my path, I do not turn back. I\nshift my reality.",
	"Press [R] to Shift your Perspective. See the world from\na new angle.",
	"Kaito: The Mage thinks he is safe. He forgot that a\nNinja strikes from the angle you least expect.",
}

var currentDialogueIdx = 0

type GameState struct {
	Level      int
	isSideView bool
	isDebug    bool
	menuState  string
}

func initialize() {
	rl.InitWindow(screenWidth, screenHeight, "Phantom-Ronin")
	rl.InitAudioDevice()

	state = GameState{
		Level:      1,
		isSideView: true,
		isDebug:    false,
		menuState:  "startMenu",
	}
}

func initializeCamera() {
	camera = rl.Camera3D{
		Position:   rl.NewVector3(8.5, 0, 2),
		Target:     rl.NewVector3(8.5, 0.0, 0.0),
		Up:         rl.NewVector3(0.0, 1.0, 0.0),
		Fovy:       10.0,
		Projection: rl.CameraOrthographic,
	}
}

func manageMusic(state string, introMusic, gameMusic rl.Music) {
	var activeMusic rl.Music
	var inactiveMusic rl.Music

	// 1. Determine which song should be playing
	if state == "startMenu" || state == "intro" {
		activeMusic = introMusic
		inactiveMusic = gameMusic
	} else if state == "inGame" {
		activeMusic = gameMusic
		inactiveMusic = introMusic
	}

	// 2. Stop the song that shouldn't be playing
	if rl.IsMusicStreamPlaying(inactiveMusic) {
		rl.StopMusicStream(inactiveMusic)
	}

	// 3. Start the active song if it's not playing
	if !rl.IsMusicStreamPlaying(activeMusic) {
		rl.PlayMusicStream(activeMusic)
	}

	// 4. Update the buffer
	rl.UpdateMusicStream(activeMusic)

	// 5. MANUAL LOOPING LOGIC
	// Check if the time played has reached the total length
	if rl.GetMusicTimePlayed(activeMusic) >= rl.GetMusicTimeLength(activeMusic) {
		rl.StopMusicStream(activeMusic)
		rl.PlayMusicStream(activeMusic)
	}
}

func main() {
	initialize()
	defer rl.CloseWindow()
	initializeCamera()

	backgroundTexture := rl.LoadTexture("./assets/background.png")

	background := Background{
		Position: rl.NewVector3(0, 0, -1.0),
		Height:   worldHeight,
		Width:    worldWidth,
		Length:   0.1,
		Color:    rl.Blue,

		Texture:         backgroundTexture,
		TextureProvided: true,
	}

	groundTexture := rl.LoadTexture("./assets/grass.jpg")

	ground := Ground{
		Position: rl.NewVector3(0.0, -3.5, 0.1),
		Height:   0.2,
		Width:    worldWidth,
		Length:   2.0,
		Color:    rl.Red,

		TextureProvided: true,
		Texture:         groundTexture,
	}

	playerTopTexture := rl.LoadTexture("./assets/topTexture.png")
	playerLeftTexture := rl.LoadTexture("./assets/backTexture.png")
	playerRightTexture := rl.LoadTexture("./assets/frontTexture.png")
	playerFrontTexture := rl.LoadTexture("./assets/leftTexture.png")
	playerBackTexture := rl.LoadTexture("./assets/rightTexture.png")
	playerBottomTexture := rl.LoadTexture("./assets/bottomTexture.png")

	player := Player{
		Position: rl.NewVector3(25.0, -1, 0.0),
		Width:    0.5,
		Height:   1.0,
		Length:   0.5,
		Color:    rl.Green,

		SPEED: 8.0,

		TextureProvided: false,
		topTexture:      playerTopTexture,
		leftTexture:     playerLeftTexture,
		rightTexture:    playerRightTexture,
		frontTexture:    playerFrontTexture,
		backTexture:     playerBackTexture,
		bottomTexture:   playerBottomTexture,
	}

	wallTexture := rl.LoadTexture("./assets/wall.jpg")

	leftWall := Wall{
		Position:        rl.NewVector3(ground.Position.X, ground.Position.Y+2.5+ground.Height/2, 0.1),
		Width:           1,
		Height:          5,
		Length:          ground.Length,
		Color:           rl.DarkBrown,
		TextureProvided: true,
		Texture:         wallTexture,
	}

	rightWall := Wall{
		Position:        rl.NewVector3(ground.Width, ground.Position.Y+2.5+ground.Height/2, 0.1),
		Width:           1,
		Height:          5,
		Length:          ground.Length,
		Color:           rl.DarkBrown,
		TextureProvided: true,
		Texture:         wallTexture,
	}

	introMusic := rl.LoadMusicStream("./assets/that-game-arcade.mp3")
	gameMusic := rl.LoadMusicStream("./assets/356-8-bit-chiptune-game-music.mp3")
	jumpSound := rl.LoadSound("./assets/pixel-jump.mp3")
	buttonSound := rl.LoadSound("./assets/button-press.mp3")
	winSound := rl.LoadSound("./assets/piglevelwin2.mp3")

	resetGame(&state, &player, &currentLevel)

	rl.SetTargetFPS(200)
	for !rl.WindowShouldClose() {
		isInGameState := state.menuState == "inGame"
		isGameOverState := state.menuState == "gameOver"
		isIntroState := state.menuState == "intro"

		manageMusic(state.menuState, introMusic, gameMusic)

		if isInGameState || isGameOverState {
			if rl.IsKeyPressed(rl.KeyR) {
				state.isSideView = !state.isSideView
			}

			player.update(state.isSideView, &background, &ground, jumpSound)
		}

		if isIntroState {
			if rl.IsKeyPressed(rl.KeyR) {
				state.isSideView = !state.isSideView
			}
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
			camera.Projection = rl.CameraOrthographic
			camera.Fovy = 10

			clampX := rl.Clamp(player.Position.X, 8.5, background.Width-8.6)
			clampY := rl.Clamp(player.Position.Y, 0.1, background.Height-player.Height)

			camera.Position = rl.NewVector3(clampX, clampY, 2)
			camera.Target = rl.NewVector3(clampX, clampY, player.Position.Z)

			player.SPEED = 8.0
		} else {
			camera.Projection = rl.CameraPerspective
			camera.Fovy = 45.0

			camera.Position = rl.NewVector3(player.Position.X+5, player.Position.Y+2, 4)
			camera.Target = rl.NewVector3(player.Position.X, player.Position.Y, 0)

			player.SPEED = 4.0
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.NewColor(135, 206, 235, 255)) // Sky blue

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
		if state.isSideView {
			rightWall.draw()
		}

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
						if state.menuState != "levelTransition" {
							rl.PlaySound(winSound)
							if state.isDebug {
								fmt.Printf("Transitioning to Level %d\n", state.Level)
							}
							state.menuState = "levelTransition"
						}
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
		case "intro":
			boxX := (int32(screenWidth) / 2) - (600 / 2)
			boxY := int32(screenHeight) - 180
			rl.DrawRectangle(boxX, boxY, 600, 150, rl.Black)
			rl.DrawRectangle(boxX+5, boxY+5, 590, 140, rl.Gray)

			rl.DrawText(introDialogues[currentDialogueIdx], boxX+10, boxY+20, 20, rl.Black)
			nextBtnRect := rl.NewRectangle(float32(boxX+480), float32(boxY+100), 100, 40)

			buttonLabel := "Next"
			if currentDialogueIdx == len(introDialogues)-1 {
				buttonLabel = "Begin"
			}

			if gui.Button(nextBtnRect, buttonLabel) || rl.IsKeyPressed(rl.KeySpace) {
				rl.PlaySound(buttonSound)
				if currentDialogueIdx < len(introDialogues)-1 {
					currentDialogueIdx++
				} else {
					state.menuState = "inGame"
					currentDialogueIdx = 0 // Reset for next time
					resetGame(&state, &player, &currentLevel)
				}
			}
		case "startMenu":
			rl.DrawText("Phanton Ronin", 80, 150, 80, rl.Red)
			startButton = gui.Button(rl.NewRectangle(float32(screenWidth)/2-50, 250, 100, 40), "Start")
			if startButton {
				rl.PlaySound(buttonSound)
				state.menuState = "intro"
			}
			exitButton = gui.Button(rl.NewRectangle(float32(screenWidth)/2-50, 300, 100, 40), "Exit")
			if exitButton {
				rl.PlaySound(buttonSound)
				rl.CloseWindow()
			}
		case "levelTransition":
			rl.DrawText("Level Completed!", 80, 150, 80, rl.Red)
			transitionButton = gui.Button(rl.NewRectangle(float32(screenWidth)/2-50, 250, 100, 40), "Next")
			if transitionButton {
				rl.PlaySound(buttonSound)
				state.menuState = "inGame"
				state.Level++
				resetGame(&state, &player, &currentLevel)
			}
		case "gameOver":
			rl.DrawText("Game Completed!", 70, 190, 80, rl.Red)
			exitButton = gui.Button(rl.NewRectangle(float32(screenWidth)/2-50, 280, 100, 40), "Exit")
			if exitButton {
				rl.PlaySound(buttonSound)
				rl.CloseWindow()
			}
		}

		if state.isDebug {
			rl.DrawText(fmt.Sprintf("Player: %.2f, %.2f, %.2f", player.Position.X, player.Position.Y, player.Position.Z), 10, 40, 18, rl.Red)
			rl.DrawText(fmt.Sprintf("Camera: %.2f, %.2f, %.2f", camera.Position.X, camera.Position.Y, camera.Position.Z), 10, 60, 18, rl.Red)
			rl.DrawText(fmt.Sprintf("Level: %d", state.Level), 10, 80, 18, rl.Red)
		} else {
			rl.DrawText(fmt.Sprintf("Level: %d", state.Level), 10, 30, 18, rl.Orange)
		}

		rl.DrawFPS(10, 10)
		rl.EndDrawing()
	}

	defer rl.UnloadTexture(backgroundTexture)

	rl.UnloadMusicStream(introMusic)
	rl.UnloadMusicStream(gameMusic)
	defer rl.UnloadSound(jumpSound)
	defer rl.UnloadSound(buttonSound)
	defer rl.UnloadSound(winSound)
	rl.CloseAudioDevice()
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
	case 3:
		currentLevel.loadLevel("./level-maps/level3.csv")
	case 4:
		currentLevel.loadLevel("./level-maps/level4.csv")
	case 5:
		currentLevel.loadLevel("./level-maps/level5.csv")
	case 6:
		currentLevel.loadLevel("./level-maps/level6.csv")
	case 7:
		currentLevel.loadLevel("./level-maps/level7.csv")
	case 8:
		currentLevel.loadLevel("./level-maps/level8.csv")
	case 9:
		currentLevel.loadLevel("./level-maps/level9.csv")
	case 10:
		currentLevel.loadLevel("./level-maps/level10.csv")
	default:
		fmt.Println("Game Completed!")
		state.Level = 0
		state.menuState = "gameOver"
	}
}
