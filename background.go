package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Background struct {
	Position rl.Vector3
	Height float32
	Width float32
	texture rl.Texture2D
}

func (b *Background) draw() {
	rl.DrawCube(b.Position, b.Width, b.Height, b.Height, rl.Blue)
}

func (b *Background) Unload() {
	rl.UnloadTexture(b.texture)
}