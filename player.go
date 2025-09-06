package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	JUMP_STRENGTH = 5.0
	SPEED         = 8.0
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

	IsGrounded bool
	jumpsUsed  int
	State      int

	TextureProvided bool
	topTexture      rl.Texture2D
	leftTexture     rl.Texture2D
	rightTexture    rl.Texture2D
	frontTexture    rl.Texture2D
	backTexture     rl.Texture2D
	bottomTexture   rl.Texture2D
}

func (p *Player) draw() {
	if !p.TextureProvided {
		rl.DrawCube(p.Position, p.Width, p.Height, p.Length, p.Color)
	} else {
		DrawCubeTextureRec_Player(p.topTexture, p.leftTexture, p.rightTexture, p.frontTexture, p.backTexture, p.bottomTexture, rl.Rectangle{X: 0, Y: 0, Width: float32(p.topTexture.Width), Height: float32(p.topTexture.Height)}, p.Position, p.Width, p.Height, p.Length, p.Color)
	}
}

func (p *Player) update(isSideView bool, b *Background, g *Ground) {
	p.Velocity.X = 0.0
	p.Velocity.Z = 0.0

	if isSideView {
		if rl.IsKeyDown(rl.KeyA) {
			p.Velocity.X = -SPEED
		}
		if rl.IsKeyDown(rl.KeyD) {
			p.Velocity.X = SPEED
		}
	} else {
		if rl.IsKeyDown(rl.KeyA) {
			p.Velocity.Z = SPEED
		}
		if rl.IsKeyDown(rl.KeyD) {
			p.Velocity.Z = -SPEED
		}
	}

	if rl.IsKeyPressed(rl.KeySpace) {
		if p.IsGrounded {
			p.Velocity.Y = JUMP_STRENGTH
			p.jumpsUsed = 1
			p.IsGrounded = false
		} else if p.jumpsUsed == 1 {
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
