package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) initializeBackground() {
	g.textures.backgroundTexture = rl.LoadTexture("./assets/images/background.png")
	g.gameObjects.background = Background{
		Position: rl.NewVector3(0, 0, -1.0),
		Height:   worldHeight,
		Width:    worldWidth,
		Length:   0.1,
		Color:    rl.Blue,

		Texture:         g.textures.backgroundTexture,
		TextureProvided: true,
	}
}

type Background struct {
	Position rl.Vector3
	Height   float32
	Width    float32
	Length   float32
	Color    rl.Color

	TextureProvided bool
	Texture         rl.Texture2D
}

func (b *Background) draw() {
	b.Position.X = 0 + (b.Width / 2) - 0.25

	if !b.TextureProvided {
		rl.DrawCube(b.Position, b.Width, b.Height, b.Length, b.Color)
	} else {
		DrawCubeTextureRec_Background(b.Texture, rl.NewRectangle(0, 0, float32(b.Texture.Width), float32(b.Texture.Height)), b.Position, b.Width, b.Height, b.Length, rl.White)
	}
}
