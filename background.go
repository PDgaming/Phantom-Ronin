package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Background struct {
	Position rl.Vector3
	Height   float32
	Width    float32
	Length   float32
	Color    rl.Color
}

func (b *Background) draw() {
	b.Position.X = 0 + (b.Width / 2) - 0.25
	rl.DrawCube(b.Position, b.Width, b.Height, b.Length, b.Color)
}
