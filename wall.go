package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Wall struct {
	Position rl.Vector3
	Width    float32
	Height   float32
	Length   float32
	Color    rl.Color

	TextureProvided bool
	Texture         rl.Texture2D
}

func (p *Wall) draw() {
	if !p.TextureProvided {
		rl.DrawCube(p.Position, p.Width, p.Height, p.Length, p.Color)
	} else {
		DrawCubeTextureRec_tiled(p.Texture, rl.NewRectangle(0, 0, float32(p.Texture.Width), float32(p.Texture.Height)), p.Position, p.Width, p.Height, p.Length, rl.White)
	}
}
