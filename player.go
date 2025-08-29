package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	JUMP_STRENGTH_1 = 5.0
	JUMP_STRENGTH_2 = 6.0
	SPEED           = 10.0
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
}

func (p *Player) draw() {
	rl.DrawCube(p.Position, p.Width, p.Height, p.Length, p.Color)
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
			p.Velocity.Y = JUMP_STRENGTH_1
			p.jumpsUsed = 1
			p.IsGrounded = false
		} else if p.jumpsUsed == 1 {
			p.Velocity.Y = JUMP_STRENGTH_2
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

	worldMaxX := b.Width - p.Width/2

	// Clamp player Z so it stays above the ground only
	minZ := g.Position.Z - g.Length/2 + p.Length/2
	maxZ := g.Position.Z + g.Length/2 - p.Length/2

	clampedX := rl.Clamp(p.Position.X, 0, worldMaxX-p.Width/2)
	clampedZ := rl.Clamp(p.Position.Z, minZ, maxZ)

	p.Position.X = clampedX
	p.Position.Z = clampedZ

	// fmt.Printf("Player Position: %v\n", p.Position)
}
