package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Ground struct {
	Position rl.Vector3
	Width    float32
	Height   float32
	Length   float32
	Color    rl.Color
}

func (g *Ground) draw() {
	g.Position.X = 0 + (g.Width / 2) - 0.25

	rl.DrawCube(g.Position, g.Width, g.Height, g.Length, g.Color)
}
