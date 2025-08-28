package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = int32(800)
	screenHeight = int32(480)

	worldWidth  = float32(20.0)
	worldLength = float32(2.0)

	GRAVITY = -9.8

	cameraOffsetZ = float32(6.0)
	cameraOffsetY = float32(0.0)
	cameraOffsetX = float32(6.0)
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
		Position: rl.NewVector3(0.0, 0.0, 0.0),
		Height:   8.1,
		Width:    worldWidth,
		Length:   0.2,
		Color:    rl.Blue,
	}

	ground := Ground{
		Position: rl.NewVector3(0.0, -2.0, 1.0),
		Height:   0.5,
		Width:    worldWidth,
		Length:   2.0,
		Color:    rl.Red,
	}

	player := Player{
		Position: rl.NewVector3(0.0, 0.0, 1.0),
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

		player.update(isSideView)

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
			// Camera position is along the Z axis (original view)
			worldBoundaryMinX := ground.Position.X - ground.Width/2
			worldBoundaryMaxX := ground.Position.X + ground.Width/2
			worldBoundaryMinY := ground.Position.Y + cameraOffsetY
			worldBoundaryMaxY := background.Position.Y + background.Height/2

			clampedCameraX := rl.Clamp(player.Position.X, worldBoundaryMinX, worldBoundaryMaxX)
			clampedCameraY := rl.Clamp(player.Position.Y+cameraOffsetY, worldBoundaryMinY, worldBoundaryMaxY)

			camera.Position = rl.NewVector3(clampedCameraX, clampedCameraY-0.2, player.Position.Z+cameraOffsetZ)
			camera.Target = rl.NewVector3(clampedCameraX, clampedCameraY, player.Position.Z)

		} else {
			// Camera position is along the X axis (rotated view)
			worldBoundaryMinZ := ground.Position.Z - ground.Length/2
			worldBoundaryMaxZ := ground.Position.Z + ground.Length/2
			worldBoundaryMinY := ground.Position.Y + cameraOffsetY
			worldBoundaryMaxY := background.Position.Y + background.Height/2

			clampedCameraZ := rl.Clamp(player.Position.Z, worldBoundaryMinZ, worldBoundaryMaxZ)
			clampedCameraY := rl.Clamp(player.Position.Y+cameraOffsetY, worldBoundaryMinY, worldBoundaryMaxY)

			camera.Position = rl.NewVector3(player.Position.X+cameraOffsetX, clampedCameraY, clampedCameraZ)
			camera.Target = rl.NewVector3(player.Position.X, clampedCameraY, clampedCameraZ)

		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.NewColor(255, 182, 193, 255))
		rl.BeginMode3D(camera)

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
