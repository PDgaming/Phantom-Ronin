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

	Velocity     rl.Vector3
	Acceleration rl.Vector3
	Mass         float32
	SPEED        float32

	IsGrounded bool
	jumpsUsed  int
	State      int

	Model    rl.Model
	UseModel bool

	modelRotationAxis  rl.Vector3
	modelRotationAngle float32
}

func (g *Game) initializePlayer() {
	g.gameObjects.player = Player{
		Position: rl.NewVector3(25.0, -1, 0.0),
		Width:    0.5,
		Height:   2.0,
		Length:   0.9,
		Color:    rl.Green,

		SPEED: 8.0,

		Model:    g.gameObjects.playerModel,
		UseModel: g.state.UseJiraiyaModel && g.gameObjects.playerModelLoaded,

		modelRotationAxis:  rl.NewVector3(0, 1, 0),
		modelRotationAngle: 90.0,
	}
}

func (p *Player) draw() {
	if p.UseModel {
		const JIRAIYA_MODEL_VISUAL_HEIGHT = 2.0

		modelPosition := rl.NewVector3(
			p.Position.X,
			p.Position.Y-p.Height/2+(JIRAIYA_MODEL_VISUAL_HEIGHT/2),
			p.Position.Z,
		)

		modelScale := rl.NewVector3(1, 1, 1)

		rl.DrawModelEx(p.Model, modelPosition, p.modelRotationAxis, p.modelRotationAngle, modelScale, rl.White)

	} else {
		rl.DrawCube(p.Position, p.Width, p.Height, p.Length, p.Color)
	}
}

func (p *Player) update(isSideView bool, g *Ground, jumpSound rl.Sound) {
	p.Velocity.X = 0.0
	p.Velocity.Z = 0.0

	if isSideView {
		if rl.IsKeyDown(rl.KeyA) {
			p.modelRotationAngle = -90.0
			p.Velocity.X = -p.SPEED
		}
		if rl.IsKeyDown(rl.KeyD) {
			p.modelRotationAngle = 90.0
			p.Velocity.X = p.SPEED
		}
	} else {
		if rl.IsKeyDown(rl.KeyA) {
			p.modelRotationAngle = 0.0
			p.Velocity.Z = p.SPEED
		}
		if rl.IsKeyDown(rl.KeyD) {
			p.modelRotationAngle = 180.0
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
