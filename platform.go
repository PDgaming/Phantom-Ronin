package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Platform struct {
	Position rl.Vector3
	Width    float32
	Height   float32
	Length   float32
	Color    rl.Color
}

func (p *Platform) draw() {
	rl.DrawCube(p.Position, p.Width, p.Height, p.Length, p.Color)
}
