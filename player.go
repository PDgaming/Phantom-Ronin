package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	Position rl.Vector3
	Width float32
	Height float32
	Rotation float32
	Collision rl.Rectangle
	Health float32

	Velocity rl.Vector3
	Acceleration rl.Vector3
	Mass float32

	IsGrounded bool
	State int
}

func (*Player) Update(dt float32) {
}