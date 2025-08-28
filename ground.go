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
	rl.DrawCube(g.Position, g.Width, g.Height, g.Length, rl.Red)
}
