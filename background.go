package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

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
		DrawCubeTextureRec(b.Texture, rl.NewRectangle(0, 0, float32(b.Texture.Width), float32(b.Texture.Height)), b.Position, b.Width, 2.0, 2.0, rl.White)
	}
}

func DrawCubeTextureRec(texture rl.Texture2D, source rl.Rectangle, position rl.Vector3, width float32, height float32, length float32, color rl.Color) {
	x := position.X
	y := position.Y
	z := position.Z
	texWidth := float32(texture.Width)
	texHeight := float32(texture.Height)

	rl.SetTexture(texture.ID)

	rl.Begin(rl.Quads)
	rl.Color4ub(color.R, color.G, color.B, color.A)

	// Front face
	rl.Normal3f(0.0, 0.0, 1.0)
	rl.TexCoord2f(source.X/texWidth, (source.Y+source.Height)/texHeight)
	rl.Vertex3f(x-width/2, y-height/2, z+length/2)
	rl.TexCoord2f((source.X+source.Width)/texWidth, (source.Y+source.Height)/texHeight)
	rl.Vertex3f(x+width/2, y-height/2, z+length/2)
	rl.TexCoord2f((source.X+source.Width)/texWidth, source.Y/texHeight)
	rl.Vertex3f(x+width/2, y+height/2, z+length/2)
	rl.TexCoord2f(source.X/texWidth, source.Y/texHeight)
	rl.Vertex3f(x-width/2, y+height/2, z+length/2)

	// Back face
	rl.Normal3f(0.0, 0.0, -1.0)
	rl.TexCoord2f((source.X+source.Width)/texWidth, (source.Y+source.Height)/texHeight)
	rl.Vertex3f(x-width/2, y-height/2, z-length/2)
	rl.TexCoord2f(source.X/texWidth, source.Y/texHeight)
	rl.Vertex3f(x+width/2, y-height/2, z-length/2)
	rl.TexCoord2f(source.X/texWidth, (source.Y+source.Height)/texHeight)
	rl.Vertex3f(x+width/2, y+height/2, z-length/2)
	rl.TexCoord2f((source.X+source.Width)/texWidth, (source.Y+source.Height)/texHeight)
	rl.Vertex3f(x-width/2, y+height/2, z-length/2)

	// Top face
	rl.Normal3f(0.0, 1.0, 0.0)
	rl.TexCoord2f(source.X/texWidth, source.Y/texHeight)
	rl.Vertex3f(x-width/2, y+height/2, z-length/2)
	rl.TexCoord2f(source.X/texWidth, (source.Y+source.Height)/texHeight)
	rl.Vertex3f(x-width/2, y+height/2, z+length/2)
	rl.TexCoord2f((source.X+source.Width)/texWidth, (source.Y+source.Height)/texHeight)
	rl.Vertex3f(x+width/2, y+height/2, z+length/2)
	rl.TexCoord2f((source.X+source.Width)/texWidth, source.Y/texHeight)
	rl.Vertex3f(x+width/2, y+height/2, z-length/2)

	// Bottom face
	rl.Normal3f(0.0, -1.0, 0.0)
	rl.TexCoord2f((source.X+source.Width)/texWidth, source.Y/texHeight)
	rl.Vertex3f(x-width/2, y-height/2, z-length/2)
	rl.TexCoord2f(source.X/texWidth, source.Y/texHeight)
	rl.Vertex3f(x+width/2, y-height/2, z-length/2)
	rl.TexCoord2f(source.X/texWidth, (source.Y+source.Height)/texHeight)
	rl.Vertex3f(x+width/2, y-height/2, z+length/2)
	rl.TexCoord2f((source.X+source.Width)/texWidth, (source.Y+source.Height)/texHeight)
	rl.Vertex3f(x-width/2, y-height/2, z+length/2)

	// Right face
	rl.Normal3f(1.0, 0.0, 0.0)
	rl.TexCoord2f((source.X+source.Width)/texWidth, (source.Y+source.Height)/texHeight)
	rl.Vertex3f(x+width/2, y-height/2, z-length/2)
	rl.TexCoord2f((source.X+source.Width)/texWidth, source.Y/texHeight)
	rl.Vertex3f(x+width/2, y+height/2, z-length/2)
	rl.TexCoord2f(source.X/texWidth, source.Y/texHeight)
	rl.Vertex3f(x+width/2, y+height/2, z+length/2)
	rl.TexCoord2f(source.X/texWidth, (source.Y+source.Height)/texHeight)
	rl.Vertex3f(x+width/2, y-height/2, z+length/2)

	// Left face
	rl.Normal3f(-1.0, 0.0, 0.0)
	rl.TexCoord2f(source.X/texWidth, (source.Y+source.Height)/texHeight)
	rl.Vertex3f(x-width/2, y-height/2, z-length/2)
	rl.TexCoord2f((source.X+source.Width)/texWidth, (source.Y+source.Height)/texHeight)
	rl.Vertex3f(x-width/2, y-height/2, z+length/2)
	rl.TexCoord2f((source.X+source.Width)/texWidth, source.Y/texHeight)
	rl.Vertex3f(x-width/2, y+height/2, z+length/2)
	rl.TexCoord2f(source.X/texWidth, source.Y/texHeight)
	rl.Vertex3f(x-width/2, y+height/2, z-length/2)

	rl.End()

	rl.SetTexture(0)
}
