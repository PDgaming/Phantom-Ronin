package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Ground struct {
	Position rl.Vector3
	Width float32
	Height float32
	Length float32
	Color rl.Color
}

func (b *Ground) draw() {
	rl.DrawCube(b.Position, b.Width, 2.0, 2.0, rl.Red)
}