package game

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

func (g *Game) initializeGround() {
	g.textures.groundTexture = rl.LoadTexture("./assets/images/grass.jpg")
	g.gameObjects.ground = Ground{
		Position: rl.NewVector3(0.0, -3.5, 0.1),
		Height:   0.2,
		Width:    worldWidth,
		Length:   2.0,
		Color:    rl.Red,

		TextureProvided: true,
		Texture:         g.textures.groundTexture,
	}
}

func (g *Ground) draw() {
	g.Position.X = 0 + (g.Width / 2) - 0.25

	if !g.TextureProvided {
		rl.DrawCube(g.Position, g.Width, g.Height, g.Length, g.Color)
	} else {
		DrawCubeTextureRec_tiled(g.Texture, rl.NewRectangle(0, 0, float32(g.Texture.Width), float32(g.Texture.Height)), g.Position, g.Width, g.Height, g.Length, rl.White)
	}
}
