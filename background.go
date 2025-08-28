package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Background struct {
	Position rl.Vector3
	Height float32
	Width float32
	Length float32
	Color rl.Color
}

func (b *Background) draw() {
	rl.DrawCube(b.Position, b.Width, b.Height, b.Length, b.Color)
}

// func (b *Background) Unload() {
// 	rl.UnloadTexture(b.texture)
// }