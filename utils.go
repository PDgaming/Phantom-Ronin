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
