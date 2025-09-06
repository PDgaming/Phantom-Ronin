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

	TextureProvided bool
	TopTexture      rl.Texture2D
	SideTexture     rl.Texture2D

	final bool
}

func (p *Platform) draw() {
	if !p.TextureProvided {
		rl.DrawCube(p.Position, p.Width, p.Height, p.Length, p.Color)
	} else {
		DrawCubeTextureRec_Platform(p.TopTexture, p.SideTexture, rl.NewRectangle(0, 0, float32(p.TopTexture.Width), float32(p.SideTexture.Height)), p.Position, p.Width, p.Height, p.Length, rl.White)
	}
}
