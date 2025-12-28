package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	JUMP_STRENGTH = 5.0
)

type Player struct {
	Position rl.Vector3
	Width    float32
	Height   float32
	Length   float32
	Rotation float32
	Color    rl.Color

	Model    rl.Model // New field for the loaded GLTF model
	UseModel bool     // New field to toggle between model and cube

	Velocity     rl.Vector3
	Acceleration rl.Vector3
	Mass         float32
	SPEED        float32

	IsGrounded bool
	jumpsUsed  int
	State      int
}

func (g *Game) initializePlayer() {
	g.gameObjects.player = Player{
		Position: rl.NewVector3(25.0, -1, 0.0),
		Width:    0.5,
		Height:   2.0,
		Length:   1,
		Color:    rl.Green,

		SPEED: 8.0,

		Model:    g.gameObjects.playerModel,                                  // Pass the loaded model
		UseModel: g.state.UseJiraiyaModel && g.gameObjects.playerModelLoaded, // Pass the flag, also considering if model loaded
	}
}

func (p *Player) draw() {
	if p.UseModel {
		// Define scaling and offset constants for the Jiraiya model
		// Define scaling and offset constants for the Jiraiya model
		const JIRAIYA_MODEL_VISUAL_HEIGHT = 2.0 // Actual visual height of the Jiraiya model in its default scale

		// The player's physics height (p.Height) is 2.0 (from Game struct initialization)
		// The visual model's height is JIRAIYA_MODEL_VISUAL_HEIGHT = 2.0
		// This makes the visual model 2.0 units tall, matching p.Height.

		// Model's feet should be at p.Position.Y - p.Height/2 (bottom of the physics box)
		// If the model's origin is at its center, and its scaled height is 2.0, its feet are at modelPosition.Y - 1.0.
		// So, modelPosition.Y - 1.0 = p.Position.Y - p.Height/2
		// modelPosition.Y = p.Position.Y - p.Height/2 + 1.0
		// Since p.Height = 2.0, p.Height/2 = 1.0.
		// modelPosition.Y = p.Position.Y - 1.0 + 1.0 = p.Position.Y

		// Corrected model position calculation:
		// Target Y for model's base: p.Position.Y - p.Height/2
		// Model's inherent Y offset from its center to its base (after scaling): JIRAIYA_MODEL_VISUAL_HEIGHT/2 = 1.0
		// So, the model's center should be at: (p.Position.Y - p.Height/2) + (JIRAIYA_MODEL_VISUAL_HEIGHT/2)
		modelPosition := rl.NewVector3(
			p.Position.X,
			p.Position.Y-p.Height/2+(JIRAIYA_MODEL_VISUAL_HEIGHT/2),
			p.Position.Z,
		)

		modelScale := rl.NewVector3(1, 1, 1)

		// Rotate the model so it faces forward (assuming GLTF +X is forward, need to rotate to +Z)
		modelRotationAxis := rl.NewVector3(0, 1, 0) // Y-axis rotation
		modelRotationAngle := float32(90.0)         // Rotate 90 degrees around Y to align X with Z

		rl.DrawModelEx(p.Model, modelPosition, modelRotationAxis, modelRotationAngle, modelScale, rl.White)

	} else {
		rl.DrawCube(p.Position, p.Width, p.Height, p.Length, p.Color)
	}
}

func (p *Player) update(isSideView bool, g *Ground, jumpSound rl.Sound) {
	p.Velocity.X = 0.0
	p.Velocity.Z = 0.0

	if isSideView {
		if rl.IsKeyDown(rl.KeyA) {
			p.Velocity.X = -p.SPEED
		}
		if rl.IsKeyDown(rl.KeyD) {
			p.Velocity.X = p.SPEED
		}
	} else {
		if rl.IsKeyDown(rl.KeyA) {
			p.Velocity.Z = p.SPEED
		}
		if rl.IsKeyDown(rl.KeyD) {
			p.Velocity.Z = -p.SPEED
		}
	}

	if rl.IsKeyPressed(rl.KeySpace) {
		if p.IsGrounded {
			rl.PlaySound(jumpSound)
			p.Velocity.Y = JUMP_STRENGTH
			p.jumpsUsed = 1
			p.IsGrounded = false
		} else if p.jumpsUsed == 1 {
			rl.PlaySound(jumpSound)
			p.Velocity.Y = JUMP_STRENGTH
			p.jumpsUsed = 2
		}
	}

	if !p.IsGrounded {
		p.Velocity.Y += GRAVITY * rl.GetFrameTime()
	}

	p.Velocity.X *= 0.5

	p.Position.X += p.Velocity.X * rl.GetFrameTime()
	p.Position.Y += p.Velocity.Y * rl.GetFrameTime()
	p.Position.Z += p.Velocity.Z * rl.GetFrameTime()

	minZ := g.Position.Z - g.Length/2 + p.Length/2
	maxZ := g.Position.Z + g.Length/2 - p.Length/2

	clampedZ := rl.Clamp(p.Position.Z, minZ, maxZ)

	p.Position.Z = clampedZ
}
