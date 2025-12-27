package game

import (
	"Phantom_Ronin/game/dialogue"
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

type GameState struct {
	Level      int
	isSideView bool
	isDebug    bool
	menuState  string
}

type Game struct {
	state                GameState
	camera               rl.Camera3D
	introDialogueManager *dialogue.Manager

	background   Background
	ground       Ground
	leftWall     Wall
	rightWall    Wall
	player       Player
	currentLevel Level

	introMusic  rl.Music
	gameMusic   rl.Music
	jumpSound   rl.Sound
	buttonSound rl.Sound
	winSound    rl.Sound

	backgroundTexture rl.Texture2D
	groundTexture     rl.Texture2D
	wallTexture       rl.Texture2D

	playerTopTexture    rl.Texture2D
	playerLeftTexture   rl.Texture2D
	playerRightTexture  rl.Texture2D
	playerFrontTexture  rl.Texture2D
	playerBackTexture   rl.Texture2D
	playerBottomTexture rl.Texture2D

	exitButton       bool
	startButton      bool
	transitionButton bool
	pauseButton      bool
}

func NewGame() *Game {
	g := &Game{}
	g.initialize()

	g.introDialogueManager = dialogue.NewManager([]string{
		"Kaito: My village didn't fall to swords or steel. It fell to\na whisper of black magic...",
		"Kaito: The Mage sits in his high tower, weaving a spell\nthat anchors our world to the void.",
		"Kaito: But the path to him is guarded by the Magical\nMountains...",
		"Kaito: To most, this world is flat. A simple path of left\nand right.",
		"Kaito: But I carry the Gaze of the Void. I can see the\nworld as it truly is.",
		"Kaito: When a wall blocks my path, I do not turn back. I\nshift my reality.",
		"Press [R] to Shift your Perspective. See the world from\na new angle.",
		"Kaito: The Mage thinks he is safe. He forgot that a\nNinja strikes from the angle you least expect.",
	}, screenWidth, screenHeight)

	g.backgroundTexture = rl.LoadTexture("./assets/background.png")
	g.background = Background{
		Position: rl.NewVector3(0, 0, -1.0),
		Height:   worldHeight,
		Width:    worldWidth,
		Length:   0.1,
		Color:    rl.Blue,

		Texture:         g.backgroundTexture,
		TextureProvided: true,
	}

	g.groundTexture = rl.LoadTexture("./assets/grass.jpg")
	g.ground = Ground{
		Position: rl.NewVector3(0.0, -3.5, 0.1),
		Height:   0.2,
		Width:    worldWidth,
		Length:   2.0,
		Color:    rl.Red,

		TextureProvided: true,
		Texture:         g.groundTexture,
	}

	g.playerTopTexture = rl.LoadTexture("./assets/topTexture.png")
	g.playerLeftTexture = rl.LoadTexture("./assets/backTexture.png")
	g.playerRightTexture = rl.LoadTexture("./assets/frontTexture.png")
	g.playerFrontTexture = rl.LoadTexture("./assets/leftTexture.png")
	g.playerBackTexture = rl.LoadTexture("./assets/rightTexture.png")
	g.playerBottomTexture = rl.LoadTexture("./assets/bottomTexture.png")

	g.player = Player{
		Position: rl.NewVector3(25.0, -1, 0.0),
		Width:    0.5,
		Height:   1.0,
		Length:   0.5,
		Color:    rl.Green,

		SPEED: 8.0,

		TextureProvided: false,
		topTexture:      g.playerTopTexture,
		leftTexture:     g.playerLeftTexture,
		rightTexture:    g.playerRightTexture,
		frontTexture:    g.playerFrontTexture,
		backTexture:     g.playerBackTexture,
		bottomTexture:   g.playerBottomTexture,
	}

	g.wallTexture = rl.LoadTexture("./assets/wall.jpg")
	g.leftWall = Wall{
		Position:        rl.NewVector3(g.ground.Position.X, g.ground.Position.Y+2.5+g.ground.Height/2, 0.1),
		Width:           1,
		Height:          5,
		Length:          g.ground.Length,
		Color:           rl.DarkBrown,
		TextureProvided: true,
		Texture:         g.wallTexture,
	}
	g.rightWall = Wall{
		Position:        rl.NewVector3(g.ground.Width, g.ground.Position.Y+2.5+g.ground.Height/2, 0.1),
		Width:           1,
		Height:          5,
		Length:          g.ground.Length,
		Color:           rl.DarkBrown,
		TextureProvided: true,
		Texture:         g.wallTexture,
	}

	g.introMusic = rl.LoadMusicStream("./assets/that-game-arcade.mp3")
	g.gameMusic = rl.LoadMusicStream("./assets/356-8-bit-chiptune-game-music.mp3")
	g.jumpSound = rl.LoadSound("./assets/pixel-jump.mp3")
	g.buttonSound = rl.LoadSound("./assets/button-press.mp3")
	g.winSound = rl.LoadSound("./assets/piglevelwin2.mp3")

	g.resetGame()

	return g
}

func (g *Game) initialize() {
	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.InitWindow(screenWidth, screenHeight, "Phantom-Ronin")
	rl.InitAudioDevice()
	g.initializeCamera()

	g.state = GameState{
		Level:      1,
		isSideView: true,
		isDebug:    false,
		menuState:  "startMenu",
	}
}

func (g *Game) initializeCamera() {
	g.camera = rl.Camera3D{
		Position:   rl.NewVector3(8.5, 0, 2),
		Target:     rl.NewVector3(8.5, 0.0, 0.0),
		Up:         rl.NewVector3(0.0, 1.0, 0.0),
		Fovy:       10.0,
		Projection: rl.CameraOrthographic,
	}
}

func (g *Game) manageMusic() {
	var activeMusic rl.Music
	var inactiveMusic rl.Music

	if g.state.menuState == "startMenu" || g.state.menuState == "intro" {
		activeMusic = g.introMusic
		inactiveMusic = g.gameMusic
	} else if g.state.menuState == "inGame" {
		activeMusic = g.gameMusic
		inactiveMusic = g.introMusic
	}

	if rl.IsMusicStreamPlaying(inactiveMusic) {
		rl.StopMusicStream(inactiveMusic)
	}

	if !rl.IsMusicStreamPlaying(activeMusic) {
		rl.PlayMusicStream(activeMusic)
	}

	rl.UpdateMusicStream(activeMusic)

	if rl.GetMusicTimePlayed(activeMusic) >= rl.GetMusicTimeLength(activeMusic) {
		rl.StopMusicStream(activeMusic)
		rl.PlayMusicStream(activeMusic)
	}
}

func (g *Game) resetGame() {
	g.player.Position = rl.NewVector3(0.0, -1.0, 0.0)
	g.player.Velocity = rl.NewVector3(0.0, 0.0, 0.0)
	g.player.IsGrounded = true
	g.player.jumpsUsed = 0
	g.currentLevel.resetLevel()

	switch g.state.Level {
	case 1:
		g.currentLevel.loadLevel("./level-maps/level1.csv")
	case 2:
		g.currentLevel.loadLevel("./level-maps/level2.csv")
	case 3:
		g.currentLevel.loadLevel("./level-maps/level3.csv")
	case 4:
		g.currentLevel.loadLevel("./level-maps/level4.csv")
	case 5:
		g.currentLevel.loadLevel("./level-maps/level5.csv")
	case 6:
		g.currentLevel.loadLevel("./level-maps/level6.csv")
	case 7:
		g.currentLevel.loadLevel("./level-maps/level7.csv")
	case 8:
		g.currentLevel.loadLevel("./level-maps/level8.csv")
	case 9:
		g.currentLevel.loadLevel("./level-maps/level9.csv")
	case 10:
		g.currentLevel.loadLevel("./level-maps/level10.csv")
	default:
		fmt.Println("Game Completed!")
		g.state.Level = 0
		g.state.menuState = "gameOver"
	}
}

func (g *Game) Run() {
	rl.SetTargetFPS(200)
	for !rl.WindowShouldClose() {
		g.update()
		g.draw()
	}

	g.unload()
}

func (g *Game) update() {
	isInGameState := g.state.menuState == "inGame"
	isGameOverState := g.state.menuState == "gameOver"
	isIntroState := g.state.menuState == "intro"
	isPausedState := g.state.menuState == "paused"

	g.manageMusic()

	if isPausedState {
		// PAUSED STATE
	} else if isInGameState || isGameOverState {
		if rl.IsKeyPressed(rl.KeyR) {
			g.state.isSideView = !g.state.isSideView
		}

		g.player.update(g.state.isSideView, &g.ground, g.jumpSound)
	}

	if isIntroState {
		if rl.IsKeyPressed(rl.KeyR) {
			g.state.isSideView = !g.state.isSideView
		}
	}

	playerBox := GetBoundingBox(g.player.Position, g.player.Width, g.player.Height, g.player.Length)
	groundBox := GetBoundingBox(g.ground.Position, g.ground.Width, g.ground.Height, g.ground.Length)
	leftWallBox := GetBoundingBox(g.leftWall.Position, g.leftWall.Width, g.leftWall.Height, g.leftWall.Length)
	rightWallBox := GetBoundingBox(g.rightWall.Position, g.rightWall.Width, g.rightWall.Height, g.rightWall.Length)

	if rl.CheckCollisionBoxes(playerBox, groundBox) {
		if g.player.Velocity.Y <= 0 {
			g.player.IsGrounded = true
			g.player.jumpsUsed = 0
			g.player.Velocity.Y = 0.0

			g.player.Position.Y = g.ground.Position.Y + (g.ground.Height / 2) + (g.player.Height / 2)
		} else {
			g.player.Velocity.Y = 0.0
			g.player.Position.Y = g.ground.Position.Y - (g.ground.Height / 2) - (g.player.Height / 2)
		}
	} else {
		g.player.IsGrounded = false
	}

	if rl.CheckCollisionBoxes(playerBox, leftWallBox) {
		g.player.Position.X = g.leftWall.Position.X + g.leftWall.Width/2 + g.player.Width/2
	}
	if rl.CheckCollisionBoxes(playerBox, rightWallBox) {
		g.player.Position.X = g.rightWall.Position.X - g.rightWall.Width/2 - g.player.Width/2
	}

	if g.state.isSideView {
		g.camera.Projection = rl.CameraOrthographic
		g.camera.Fovy = 10

		clampX := rl.Clamp(g.player.Position.X, 8.5, g.background.Width-8.6)
		clampY := rl.Clamp(g.player.Position.Y, 0.1, g.background.Height-g.player.Height)

		g.camera.Position = rl.NewVector3(clampX, clampY, 2)
		g.camera.Target = rl.NewVector3(clampX, clampY, g.player.Position.Z)

		g.player.SPEED = 8.0
	} else {
		g.camera.Projection = rl.CameraPerspective
		g.camera.Fovy = 45.0

		g.camera.Position = rl.NewVector3(g.player.Position.X+5, g.player.Position.Y+2, 4)
		g.camera.Target = rl.NewVector3(g.player.Position.X, g.player.Position.Y, 0)

		g.player.SPEED = 4.0
	}

	for _, platform := range g.currentLevel.Platforms {
		platformBox := GetBoundingBox(platform.Position, platform.Width, platform.Height, platform.Length)

		if rl.CheckCollisionBoxes(playerBox, platformBox) {
			playerBottom := g.player.Position.Y - g.player.Height/2
			platformTop := platform.Position.Y + platform.Height/2
			platformBottom := platform.Position.Y - platform.Height/2

			if playerBottom >= platformTop-0.05 && g.player.Velocity.Y <= 0 {
				g.player.Position.Y = platformTop + g.player.Height/2
				g.player.IsGrounded = true
				g.player.jumpsUsed = 0
				g.player.Velocity.Y = 0.0

				if platform.final {
					if g.state.menuState != "levelTransition" {
						rl.PlaySound(g.winSound)
						if g.state.isDebug {
							fmt.Printf("Transitioning to Level %d\n", g.state.Level)
						}
						g.state.menuState = "levelTransition"
					}
				}
			} else if (g.player.Position.Y+g.player.Height/2) <= platformBottom+0.05 && g.player.Velocity.Y > 0 {
				g.player.Position.Y = platformBottom - g.player.Height/2
				g.player.Velocity.Y = 0.0
			} else {
				playerTop := g.player.Position.Y + g.player.Height/2

				if playerTop > platformBottom && platformBottom < platformTop {
					if g.state.isSideView {
						playerLeft := g.player.Position.X - g.player.Width/2
						playerRight := g.player.Position.X + g.player.Width/2
						platformLeft := platform.Position.X - platform.Width/2
						platformRight := platform.Position.X + platform.Width/2

						overlapLeft := playerRight - platformLeft
						overlapRight := platformRight - playerLeft

						if overlapLeft < overlapRight {
							g.player.Position.X -= overlapLeft
						} else {
							g.player.Position.X += overlapRight
						}
					} else {
						playerFront := g.player.Position.Z + g.player.Length/2
						playerBack := g.player.Position.Z - g.player.Length/2
						platformFront := platform.Position.Z + platform.Length/2
						platformBack := platform.Position.Z - platform.Length/2

						overlapBack := playerFront - platformBack
						overlapFront := platformFront - playerBack

						if overlapBack < overlapFront {
							g.player.Position.Z -= overlapBack
						} else {
							g.player.Position.Z += overlapFront
						}
					}
				}
			}
		}
	}
}

func (g *Game) draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.NewColor(135, 206, 235, 255))

	rl.BeginMode3D(g.camera)

	playerBox := GetBoundingBox(g.player.Position, g.player.Width, g.player.Height, g.player.Length)
	groundBox := GetBoundingBox(g.ground.Position, g.ground.Width, g.ground.Height, g.ground.Length)
	leftWallBox := GetBoundingBox(g.leftWall.Position, g.leftWall.Width, g.leftWall.Height, g.leftWall.Length)
	rightWallBox := GetBoundingBox(g.rightWall.Position, g.rightWall.Width, g.rightWall.Height, g.rightWall.Length)

	if g.state.isDebug {
		rl.DrawBoundingBox(playerBox, rl.Red)
		rl.DrawBoundingBox(groundBox, rl.Green)
		rl.DrawBoundingBox(leftWallBox, rl.Blue)
		rl.DrawBoundingBox(rightWallBox, rl.Blue)
	}

	g.background.draw()
	g.ground.draw()
	g.leftWall.draw()
	if g.state.isSideView {
		g.rightWall.draw()
	}

	for _, platform := range g.currentLevel.Platforms {
		platform.draw()

		if g.state.isDebug {
			platformBox := GetBoundingBox(platform.Position, platform.Width, platform.Height, platform.Length)
			rl.DrawBoundingBox(platformBox, rl.Red)
		}
	}

	g.player.draw()

	rl.EndMode3D()

	g.drawUI()

	if g.state.isDebug {
		rl.DrawText(fmt.Sprintf("Player: %.2f, %.2f, %.2f", g.player.Position.X, g.player.Position.Y, g.player.Position.Z), 10, 40, 18, rl.Red)
		rl.DrawText(fmt.Sprintf("Camera: %.2f, %.2f, %.2f", g.camera.Position.X, g.camera.Position.Y, g.camera.Position.Z), 10, 60, 18, rl.Red)
		rl.DrawText(fmt.Sprintf("Level: %d", g.state.Level), 10, 80, 18, rl.Red)
	} else {
		rl.DrawText(fmt.Sprintf("Level: %d", g.state.Level), 10, 30, 18, rl.Orange)
	}

	rl.DrawFPS(10, 10)
	rl.EndDrawing()
}

func (g *Game) drawUI() {
	switch g.state.menuState {
	case "intro":
		g.introDialogueManager.Draw()
		if g.introDialogueManager.Update(g.buttonSound) {
			g.state.menuState = "inGame"
			g.introDialogueManager.Reset()
			g.resetGame()
		}
	case "startMenu":
		rl.DrawText("Phanton Ronin", 80, 150, 80, rl.Red)
		g.startButton = gui.Button(rl.NewRectangle(float32(screenWidth)/2-50, 250, 100, 40), "Start")
		if g.startButton {
			rl.PlaySound(g.buttonSound)
			g.state.menuState = "intro"
		}
		g.exitButton = gui.Button(rl.NewRectangle(float32(screenWidth)/2-50, 300, 100, 40), "Exit")
		if g.exitButton {
			rl.PlaySound(g.buttonSound)
			rl.CloseWindow()
		}
	case "inGame":
		g.pauseButton = gui.Button(rl.NewRectangle(float32(screenWidth)-60, 10, 50, 30), "Pause")
		if g.pauseButton {
			rl.PlaySound(g.buttonSound)
			g.state.menuState = "paused"
		}
	case "paused":
		g.startButton = gui.Button(rl.NewRectangle(float32(screenWidth)/2-50, 200, 100, 40), "Resume")
		if g.startButton {
			rl.PlaySound(g.buttonSound)
			g.state.menuState = "inGame"
		}

		// Placeholder for Save Button
		gui.Button(rl.NewRectangle(float32(screenWidth)/2-50, 250, 100, 40), "Save")

		g.exitButton = gui.Button(rl.NewRectangle(float32(screenWidth)/2-50, 300, 100, 40), "Exit")
		if g.exitButton {
			rl.PlaySound(g.buttonSound)
			rl.CloseWindow()
		}
	case "levelTransition":
		rl.DrawText("Level Completed!", 80, 150, 80, rl.Red)
		g.transitionButton = gui.Button(rl.NewRectangle(float32(screenWidth)/2-50, 250, 100, 40), "Next")
		if g.transitionButton {
			rl.PlaySound(g.buttonSound)
			g.state.menuState = "inGame"
			g.state.Level++
			g.resetGame()
		}
	case "gameOver":
		rl.DrawText("Game Completed!", 70, 190, 80, rl.Red)
		g.exitButton = gui.Button(rl.NewRectangle(float32(screenWidth)/2-50, 280, 100, 40), "Exit")
		if g.exitButton {
			rl.PlaySound(g.buttonSound)
			rl.CloseWindow()
		}
	}
}

func (g *Game) unload() {
	rl.UnloadTexture(g.backgroundTexture)
	rl.UnloadTexture(g.groundTexture)
	rl.UnloadTexture(g.wallTexture)
	rl.UnloadTexture(g.playerTopTexture)
	rl.UnloadTexture(g.playerLeftTexture)
	rl.UnloadTexture(g.playerRightTexture)
	rl.UnloadTexture(g.playerFrontTexture)
	rl.UnloadTexture(g.playerBackTexture)
	rl.UnloadTexture(g.playerBottomTexture)

	rl.UnloadMusicStream(g.introMusic)
	rl.UnloadMusicStream(g.gameMusic)
	rl.UnloadSound(g.jumpSound)
	rl.UnloadSound(g.buttonSound)
	rl.UnloadSound(g.winSound)
	rl.CloseAudioDevice()
	rl.CloseWindow()
}
