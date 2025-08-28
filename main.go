package main

import (
	"image/png"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = int32(800)
	screenHeight = int32(480)
)

func main() {
	// Init
	rl.InitWindow(screenWidth, screenHeight, "Phantom-Ronin")
	defer rl.CloseWindow()

	// Camera
	camera := rl.Camera3D{}
	camera.Position = rl.NewVector3(0, 0, 10.0)
	camera.Target = rl.NewVector3(0.0, 0.0, 0.0)
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Fovy = 45.0
	camera.Projection = rl.CameraPerspective
	cameraSpeed := float32(0.1)

	backgroundTexture := loadBackground()
	background := Background{
		Position: rl.NewVector3(1.0, 0.0, 0.0),
		Height: 2.0,
		Width: 2.0,
		texture: backgroundTexture,
	}

	ground := Ground{
		Position: rl.NewVector3(0.0, 0.0, 0.0),
		Width:    2.0,
		Height:   2.0,
		Length:   2.0,
		Color: rl.Red,
	}

	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		if rl.IsKeyDown(rl.KeyA) {
			camera.Position.X -= cameraSpeed
			camera.Target.X -= cameraSpeed
		}
		if rl.IsKeyDown(rl.KeyD) {
			camera.Position.X += cameraSpeed
			camera.Target.X += cameraSpeed
		}
		
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)		

		rl.BeginMode3D(camera)

		background.draw()
		ground.draw()

		rl.EndMode3D()
		rl.DrawFPS(10, 10)
		rl.EndDrawing()
	}
}

func loadBackground() rl.Texture2D {
	r, err := os.Open("./assets/background.png")
	if err != nil {
		rl.TraceLog(rl.LogError, err.Error())
	}
	defer r.Close()

	img, err := png.Decode(r)
	if err != nil {
		rl.TraceLog(rl.LogError, err.Error())
	}

	im := rl.NewImageFromImage(img)
	texture := rl.LoadTextureFromImage(im)

	rl.UnloadImage(im)

	return texture
}