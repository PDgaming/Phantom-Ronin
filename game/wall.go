package game

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

func (g *Game) initializeWalls() {
	g.textures.wallTexture = rl.LoadTexture("./assets/images/wall.jpg")
	g.gameObjects.leftWall = Wall{
		Position:        rl.NewVector3(g.gameObjects.ground.Position.X, g.gameObjects.ground.Position.Y+2.5+g.gameObjects.ground.Height/2, 1.1),
		Width:           1,
		Height:          5,
		Length:          g.gameObjects.ground.Length,
		Color:           rl.DarkBrown,
		TextureProvided: true,
		Texture:         g.textures.wallTexture,
	}
	g.gameObjects.rightWall = Wall{
		Position:        rl.NewVector3(g.gameObjects.ground.Width, g.gameObjects.ground.Position.Y+2.5+g.gameObjects.ground.Height/2, 0.1),
		Width:           1,
		Height:          5,
		Length:          g.gameObjects.ground.Length,
		Color:           rl.DarkBrown,
		TextureProvided: true,
		Texture:         g.textures.wallTexture,
	}
}

func (p *Wall) draw() {
	if !p.TextureProvided {
		rl.DrawCube(p.Position, p.Width, p.Height, p.Length, p.Color)
	} else {
		DrawCubeTextureRec_tiled(p.Texture, rl.NewRectangle(0, 0, float32(p.Texture.Width), float32(p.Texture.Height)), p.Position, p.Width, p.Height, p.Length, rl.White)
	}
}
