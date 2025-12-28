package game

import (
	"Phantom_Ronin/game/dialogue"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) initialize() {
	rl.InitWindow(screenWidth, screenHeight, "Phantom-Ronin")
	rl.InitAudioDevice()
	g.initializeCamera()

	g.state = GameState{
		Level:      1,
		isSideView: true,
		isDebug:    true,
		menuState:  "startMenu",
		ShowIntro:  true,
	}
	rl.SetExitKey(rl.KeyNull)
}

func NewGame() *Game {
	g := &Game{}
	g.initialize()

	initializeAutoSave(g)

	// Initialize UseJiraiyaModel
	g.state.UseJiraiyaModel = false // Default to false

	// Load Jiraiya GLTF model
	g.gameObjects.playerModel = rl.LoadModel("assets/jiraiya/scene.gltf")
	g.gameObjects.playerModelLoaded = (g.gameObjects.playerModel.MeshCount > 0) // Check if model loaded successfully using MeshCount
	if !g.gameObjects.playerModelLoaded {
		fmt.Println("WARNING: Jiraiya GLTF model not loaded. Falling back to cube.")
	}

	// Load settings
	settings, err := LoadSettings()
	if err != nil {
		fmt.Printf("Error loading settings, using defaults: %v\n", err)
	} else {
		g.state.ShowIntro = settings.ShowIntro
		g.state.UseJiraiyaModel = settings.UseJiraiyaModel
	}

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

	g.initializeBackground()
	g.initializeGround()
	g.initializePlayer()
	g.initializeWalls()

	g.loadMusicAndSounds()

	g.resetGame(g.state.Level)

	return g
}

func (g *Game) resetGame(level int) {
	rl.StopSound(g.audioStreams.winSound)
	g.state.Level = level
	g.gameObjects.player.Position = rl.NewVector3(0.0, -1.0, 0.0)
	g.gameObjects.player.Velocity = rl.NewVector3(0.0, 0.0, 0.0)
	g.gameObjects.player.IsGrounded = true
	g.gameObjects.player.jumpsUsed = 0
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
		if g.state.menuState != "gameOver" {
			fmt.Println("Game Completed!")
		}
		g.state.Level = 0
		g.state.menuState = "gameOver"
		drawGameOverMenu()
	}
}

func (g *Game) Run() {
	rl.SetTargetFPS(200)
	for !rl.WindowShouldClose() {

		if rl.IsKeyPressed(rl.KeyEscape) {
			switch g.state.menuState {
			case "inGame":
				g.state.menuState = "paused"
			case "paused":
				g.state.menuState = "inGame"
			default:
				drawShouldCloseWindowMenu()
			}
		}

		g.update()
		g.draw()
	}

	g.unload()
}

func (g *Game) update() {
	if g.displayMessageTimer > 0 {
		g.displayMessageTimer -= rl.GetFrameTime()
	}

	isIntroState := g.state.menuState == "intro"
	isInGameState := g.state.menuState == "inGame"
	// isPausedState := g.state.menuState == "paused"
	isGameOverState := g.state.menuState == "gameOver"

	g.manageMusic()

	autoSave(isInGameState, g)

	if isInGameState || isGameOverState || isIntroState {
		if rl.IsKeyPressed(rl.KeyR) {
			g.state.isSideView = !g.state.isSideView
		}

		g.gameObjects.player.update(g.state.isSideView, &g.gameObjects.ground, g.audioStreams.jumpSound)
	}

	initializeBoundingBoxes(g)
	checkCollision(g)
	checkCollisionWithPlatforms(g)

	g.updateCamera()
}

func (g *Game) draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.NewColor(135, 206, 235, 255))

	rl.BeginMode3D(g.camera)

	initializeBoundingBoxes(g)

	g.gameObjects.background.draw()
	g.gameObjects.ground.draw()
	g.gameObjects.leftWall.draw()
	if g.state.isSideView {
		g.gameObjects.rightWall.draw()
	}
	drawPlatforms(g)
	g.gameObjects.player.draw()

	rl.EndMode3D()

	g.drawUI()

	if g.state.isDebug {
		drawBoundingBoxes(g)
		drawDebugInfo(g)
	} else {
		rl.DrawText(fmt.Sprintf("Level: %d", g.state.Level), 10, 30, 18, rl.Orange)
	}

	rl.DrawFPS(10, 10)
	rl.EndDrawing()
}

func (g *Game) unload() {
	rl.UnloadTexture(g.textures.backgroundTexture)
	rl.UnloadTexture(g.textures.groundTexture)
	rl.UnloadTexture(g.textures.wallTexture)
	rl.UnloadModel(g.gameObjects.playerModel)
	unloadAudio(g)
	rl.CloseWindow()
}
