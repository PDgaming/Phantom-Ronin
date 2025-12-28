package game

import (
	"Phantom_Ronin/game/dialogue"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	state                GameState
	camera               rl.Camera3D
	introDialogueManager *dialogue.Manager
	currentLevel         Level

	gameObjects  gameObjects
	audioStreams audioStreams
	textures     textures
	buttons      buttons
	autoSaver    autoSaver

	displayMessageText  string
	displayMessageTimer float32

	boxes BoundingBox
}

type GameState struct {
	Level           int
	isSideView      bool
	isDebug         bool
	menuState       string
	ShowIntro       bool
	UseJiraiyaModel bool
}

type gameObjects struct {
	background        Background
	ground            Ground
	leftWall          Wall
	rightWall         Wall
	player            Player
	playerModel       rl.Model
	playerModelLoaded bool
}

type audioStreams struct {
	introMusic  rl.Music
	gameMusic   rl.Music
	jumpSound   rl.Sound
	buttonSound rl.Sound
	winSound    rl.Sound
}

type textures struct {
	backgroundTexture rl.Texture2D
	groundTexture     rl.Texture2D
	wallTexture       rl.Texture2D
}

type buttons struct {
	exitButton       bool
	startButton      bool
	transitionButton bool
	pauseButton      bool
}

type autoSaver struct {
	autosaveTimer                float32
	autosaveInterval             float32
	autosaveDelayAfterManualSave float32
	currentAutosaveInterval      float32
}

type BoundingBox struct {
	PlayerBox    rl.BoundingBox
	GroundBox    rl.BoundingBox
	LeftWallBox  rl.BoundingBox
	RightWallBox rl.BoundingBox
}
