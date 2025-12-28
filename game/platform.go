package game

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

	Model    rl.Model
	UseModel bool

	modelRotationAxis  rl.Vector3
	modelRotationAngle float32
	modelScale         rl.Vector3
}

func (p *Platform) draw() {
	if p.UseModel {
		modelPosition := p.Position
		modelPosition.Y -= p.Height / 2 // Adjust Y position to align model with the bottom of the bounding box

		// Use a more uniform scaling factor
		modelScale := rl.NewVector3(p.Width*0.01, p.Height*0.03, p.Length*0.01)

		modelRotationAxis := rl.NewVector3(0, 1, 0)
		modelRotationAngle := float32(0)

		rl.DrawModelEx(p.Model, modelPosition, modelRotationAxis, modelRotationAngle, modelScale, rl.White)

	} else if !p.TextureProvided {
		rl.DrawCube(p.Position, p.Width, p.Height, p.Length, p.Color)
	} else {
		DrawCubeTextureRec_Platform(p.TopTexture, p.SideTexture, rl.NewRectangle(0, 0, float32(p.TopTexture.Width), float32(p.SideTexture.Height)), p.Position, p.Width, p.Height, p.Length, rl.White)
	}
}

func drawPlatforms(g *Game) {
	for _, platform := range g.currentLevel.Platforms {
		platform.draw()

		if g.state.isDebug {
			platformBox := GetBoundingBox(platform.Position, platform.Width, platform.Height, platform.Length)
			rl.DrawBoundingBox(platformBox, rl.Red)
		}
	}
}
