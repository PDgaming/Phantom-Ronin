package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = int32(800)
	screenHeight = int32(480)

	worldWidth  = float32(5)
	worldLength = float32(2)

	GRAVITY = -9.8

	cameraOffsetX = float32(6.0)
	cameraOffsetZ = float32(6.0)
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

	isSideView := true

	// backgroundTexture := loadBackground()
	background := Background{
		Position: rl.NewVector3(0.0, 0.0, -1.0),
		Height:   float32(screenHeight),
		Width:    worldWidth,
		Length:   0.1,
		Color:    rl.Blue,
	}

	ground := Ground{
		Position: rl.NewVector3(0.0, -1.5, 0.1),
		Height:   0.2,
		Width:    worldWidth,
		Length:   2.0,
		Color:    rl.Red,
	}

	player := Player{
		Position: rl.NewVector3(0.0, 0.0, 0.0),
		Width:    0.5,
		Height:   1.0,
		Length:   0.5,
		Color:    rl.Green,
	}

	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		if rl.IsKeyPressed(rl.KeyR) {
			isSideView = !isSideView
		}

		player.update(isSideView, &background, &ground)

		playerBox := GetBoundingBox(player.Position, player.Width, player.Height, player.Length)
		groundBox := GetBoundingBox(ground.Position, ground.Width, ground.Height, ground.Length)

		if rl.CheckCollisionBoxes(playerBox, groundBox) {
			if player.Velocity.Y <= 0 {
				player.IsGrounded = true
				player.jumpsUsed = 0
				player.Velocity.Y = 0.0

				player.Position.Y = ground.Position.Y + (ground.Height / 2) + (player.Height / 2)
			} else {
				player.Velocity.Y = 0.0
				player.Position.Y = ground.Position.Y - (ground.Height / 2) - (player.Height / 2)
			}
		} else {
			player.IsGrounded = false
		}

		if isSideView {
			clampY := rl.Clamp(camera.Position.Y, 0, background.Height-player.Height)

			// Camera position is along the Z axis (original view), no limits
			camera.Position = rl.NewVector3(player.Position.X, clampY, player.Position.Z+cameraOffsetZ)
			camera.Target = rl.NewVector3(player.Position.X, clampY, player.Position.Z)

			fmt.Printf("Camera Position: %v\n", camera.Position)
		} else {
			clampY := rl.Clamp(camera.Position.Y, 0, background.Height-player.Height)
			clampZ := rl.Clamp(player.Position.Z, 0, ground.Width-player.Width/2)
			// Camera position is along the X axis (rotated view), no limits
			camera.Position = rl.NewVector3(player.Position.X+cameraOffsetX, clampY, clampZ)
			camera.Target = rl.NewVector3(player.Position.X, clampY, clampZ)
			fmt.Printf("Camera Position: %v\n", camera.Position)
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.NewColor(255, 182, 193, 255))

		rl.BeginMode3D(camera)
		rl.DrawBoundingBox(playerBox, rl.Red)
		rl.DrawBoundingBox(groundBox, rl.Green)

		background.draw()
		ground.draw()

		player.draw()

		rl.EndMode3D()
		rl.DrawFPS(10, 10)
		rl.EndDrawing()
	}
}

// func loadBackground() rl.Texture2D {
// 	r, err := os.Open("./assets/background.png")
// 	if err != nil {
// 		rl.TraceLog(rl.LogError, err.Error())
// 	}
// 	defer r.Close()

// 	img, err := png.Decode(r)
// 	if err != nil {
// 		rl.TraceLog(rl.LogError, err.Error())
// 	}

// 	im := rl.NewImageFromImage(img)
// 	texture := rl.LoadTextureFromImage(im)

// 	rl.UnloadImage(im)

// 	return texture
// }
