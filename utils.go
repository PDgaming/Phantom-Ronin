package main

import rl "github.com/gen2brain/raylib-go/raylib"

func GetBoundingBox(position rl.Vector3, width, height, length float32) rl.BoundingBox {
	halfWidth := width / 2.0
	halfHeight := height / 2.0
	halfLength := length / 2.0

	min := rl.NewVector3(
		position.X-halfWidth,
		position.Y-halfHeight,
		position.Z-halfLength,
	)

	max := rl.NewVector3(
		position.X+halfWidth,
		position.Y+halfHeight,
		position.Z+halfLength,
	)

	return rl.BoundingBox{Min: min, Max: max}
}

func DrawCubeTextureRec_Background(texture rl.Texture2D, source rl.Rectangle, position rl.Vector3, width float32, height float32, length float32, color rl.Color) {
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

func DrawCubeTextureRec_Platform(topTexture rl.Texture2D, sideTexture rl.Texture2D, source rl.Rectangle, position rl.Vector3, width float32, height float32, length float32, color rl.Color) {
	x := position.X
	y := position.Y
	z := position.Z
	texWidth := float32(source.Width)
	texHeight := float32(source.Height)

	rl.Begin(rl.Quads)
	rl.Color4ub(color.R, color.G, color.B, color.A)

	// Set texture for front, back, right, and left faces
	rl.SetTexture(sideTexture.ID)

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

	// Set texture for top and bottom faces
	rl.SetTexture(topTexture.ID)

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

	rl.SetTexture(sideTexture.ID)

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

	rl.End()

	rl.SetTexture(0)
}

func DrawCubeTextureRec_tiled(texture rl.Texture2D, source rl.Rectangle, position rl.Vector3, width float32, height float32, length float32, color rl.Color) {
	x := position.X
	y := position.Y
	z := position.Z
	texWidth := float32(texture.Width)
	texHeight := float32(texture.Height)

	rl.SetTexture(texture.ID)

	rl.Begin(rl.Quads)
	rl.Color4ub(color.R, color.G, color.B, color.A)

	// Calculate how many times the `source` tile fits on each face.
	// This is the core logic for "tiling".
	// For a Minecraft-style effect, `source.Width` and `source.Height`
	// would be the size of a single texture tile (e.g., 16x16).
	// tileX = number of tiles horizontally
	// tileY = number of tiles vertically
	// tileZ = number of tiles along the depth axis (length)
	tileX := float32(16)
	tileY := float32(4)
	tileZ := float32(4)

	// Front face
	rl.Normal3f(0.0, 0.0, 1.0)
	// The texture coordinates for the bottom-left corner are now scaled by tileX and tileY.
	// We start at the source X and Y, and extend to the source + (tile_count * source_size).
	// This tells OpenGL to repeat the texture that many times.
	rl.TexCoord2f(source.X/texWidth, (source.Y+source.Height*tileY)/texHeight)
	rl.Vertex3f(x-width/2, y-height/2, z+length/2)
	rl.TexCoord2f((source.X+source.Width*tileX)/texWidth, (source.Y+source.Height*tileY)/texHeight)
	rl.Vertex3f(x+width/2, y-height/2, z+length/2)
	rl.TexCoord2f((source.X+source.Width*tileX)/texWidth, source.Y/texHeight)
	rl.Vertex3f(x+width/2, y+height/2, z+length/2)
	rl.TexCoord2f(source.X/texWidth, source.Y/texHeight)
	rl.Vertex3f(x-width/2, y+height/2, z+length/2)

	// Back face
	rl.Normal3f(0.0, 0.0, -1.0)
	// ... all other faces follow the same pattern, using the appropriate tile variable ...
	rl.TexCoord2f((source.X+source.Width*tileX)/texWidth, (source.Y+source.Height*tileY)/texHeight)
	rl.Vertex3f(x-width/2, y-height/2, z-length/2)
	rl.TexCoord2f(source.X/texWidth, (source.Y+source.Height*tileY)/texHeight)
	rl.Vertex3f(x+width/2, y-height/2, z-length/2)
	rl.TexCoord2f(source.X/texWidth, source.Y/texHeight)
	rl.Vertex3f(x+width/2, y+height/2, z-length/2)
	rl.TexCoord2f((source.X+source.Width*tileX)/texWidth, source.Y/texHeight)
	rl.Vertex3f(x-width/2, y+height/2, z-length/2)

	// Top face (tiling along width and length)
	rl.Normal3f(0.0, 1.0, 0.0)
	rl.TexCoord2f(source.X/texWidth, source.Y/texHeight)
	rl.Vertex3f(x-width/2, y+height/2, z-length/2)
	rl.TexCoord2f(source.X/texWidth, (source.Y+source.Height*tileZ)/texHeight)
	rl.Vertex3f(x-width/2, y+height/2, z+length/2)
	rl.TexCoord2f((source.X+source.Width*tileX)/texWidth, (source.Y+source.Height*tileZ)/texHeight)
	rl.Vertex3f(x+width/2, y+height/2, z+length/2)
	rl.TexCoord2f((source.X+source.Width*tileX)/texWidth, source.Y/texHeight)
	rl.Vertex3f(x+width/2, y+height/2, z-length/2)

	// Bottom face (tiling along width and length)
	rl.Normal3f(0.0, -1.0, 0.0)
	rl.TexCoord2f((source.X+source.Width*tileX)/texWidth, source.Y/texHeight)
	rl.Vertex3f(x-width/2, y-height/2, z-length/2)
	rl.TexCoord2f(source.X/texWidth, source.Y/texHeight)
	rl.Vertex3f(x+width/2, y-height/2, z-length/2)
	rl.TexCoord2f(source.X/texWidth, (source.Y+source.Height*tileZ)/texHeight)
	rl.Vertex3f(x+width/2, y-height/2, z+length/2)
	rl.TexCoord2f((source.X+source.Width*tileX)/texWidth, (source.Y+source.Height*tileZ)/texHeight)
	rl.Vertex3f(x-width/2, y-height/2, z+length/2)

	// Right face (tiling along height and length)
	rl.Normal3f(1.0, 0.0, 0.0)
	rl.TexCoord2f((source.X+source.Width*tileZ)/texWidth, (source.Y+source.Height*tileY)/texHeight)
	rl.Vertex3f(x+width/2, y-height/2, z-length/2)
	rl.TexCoord2f((source.X+source.Width*tileZ)/texWidth, source.Y/texHeight)
	rl.Vertex3f(x+width/2, y+height/2, z-length/2)
	rl.TexCoord2f(source.X/texWidth, source.Y/texHeight)
	rl.Vertex3f(x+width/2, y+height/2, z+length/2)
	rl.TexCoord2f(source.X/texWidth, (source.Y+source.Height*tileY)/texHeight)
	rl.Vertex3f(x+width/2, y-height/2, z+length/2)

	// Left face (tiling along height and length)
	rl.Normal3f(-1.0, 0.0, 0.0)
	rl.TexCoord2f(source.X/texWidth, (source.Y+source.Height*tileY)/texHeight)
	rl.Vertex3f(x-width/2, y-height/2, z-length/2)
	rl.TexCoord2f((source.X+source.Width*tileZ)/texWidth, (source.Y+source.Height*tileY)/texHeight)
	rl.Vertex3f(x-width/2, y-height/2, z+length/2)
	rl.TexCoord2f((source.X+source.Width*tileZ)/texWidth, source.Y/texHeight)
	rl.Vertex3f(x-width/2, y+height/2, z+length/2)
	rl.TexCoord2f(source.X/texWidth, source.Y/texHeight)
	rl.Vertex3f(x-width/2, y+height/2, z-length/2)

	rl.End()

	rl.SetTexture(0)
}

func DrawCubeTextureRec_Player(topTexture rl.Texture2D, leftTexture rl.Texture2D, rightTexture rl.Texture2D, frontTexture rl.Texture2D, backTexture rl.Texture2D, bottomTexture rl.Texture2D, source rl.Rectangle, position rl.Vector3, width float32, height float32, length float32, color rl.Color) {
	x := position.X
	y := position.Y
	z := position.Z
	texWidth := float32(source.Width)
	texHeight := float32(source.Height)

	rl.Begin(rl.Quads)
	rl.Color4ub(color.R, color.G, color.B, color.A)

	// Set texture for front, back, right, and left faces
	rl.SetTexture(frontTexture.ID)

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

	rl.SetTexture(backTexture.ID)

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

	rl.SetTexture(rightTexture.ID)

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

	rl.SetTexture(leftTexture.ID)

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

	// Set texture for top and bottom faces
	rl.SetTexture(topTexture.ID)

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

	rl.SetTexture(bottomTexture.ID)

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

	rl.End()

	rl.SetTexture(0)
}
