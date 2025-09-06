package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Ground struct {
	Position rl.Vector3
	Width    float32
	Height   float32
	Length   float32
	Color    rl.Color

	TextureProvided bool
	Texture         rl.Texture2D
}

func (g *Ground) draw() {
	g.Position.X = 0 + (g.Width / 2) - 0.25

	if !g.TextureProvided {
		rl.DrawCube(g.Position, g.Width, g.Height, g.Length, g.Color)
	} else {
		DrawCubeTextureRec_tiled(g.Texture, rl.NewRectangle(0, 0, float32(g.Texture.Width), float32(g.Texture.Height)), g.Position, g.Width, g.Height, g.Length, rl.White)
	}
}
