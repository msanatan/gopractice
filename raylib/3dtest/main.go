package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	entities "github.com/msanatan/gopractice/raylib/3dtest/entities"
)

// Global variables
var (
	camera         rl.Camera
	player         *entities.Player
	playerSpeed    float32 = 5.0
	playerVelocity rl.Vector3
)

func UpdateGame() {
	playerVelocity = rl.Vector3{X: 0.0, Y: 0.0, Z: 0.0}

	// Update player position based on arrow key input
	if rl.IsKeyDown(rl.KeyLeft) {
		playerVelocity.X -= 1.0
	}
	if rl.IsKeyDown(rl.KeyRight) {
		playerVelocity.X += 1.0
	}
	if rl.IsKeyDown(rl.KeyUp) {
		playerVelocity.Z -= 1.0
	}
	if rl.IsKeyDown(rl.KeyDown) {
		playerVelocity.Z += 1.0
	}

	playerVelocity = rl.Vector3Scale(rl.Vector3Normalize(playerVelocity), playerSpeed*rl.GetFrameTime())
	player.Move(playerVelocity)
}

func DrawGame() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	rl.BeginMode3D(camera)
	player.Draw()
	rl.EndMode3D()

	rl.DrawFPS(10, 10)
	rl.EndDrawing()
}

func main() {
	screenWidth := int32(1024)
	screenHeight := int32(768)
	rl.InitWindow(screenWidth, screenHeight, "Move Cube")

	// Set up camera
	camera = rl.Camera{
		Position: rl.Vector3{X: 0.0, Y: 10.0, Z: 10.0},
		Target:   rl.Vector3{X: 0.0, Y: 0.0, Z: 0.0},
		Up:       rl.Vector3{X: 0.0, Y: 1.0, Z: 0.0},
		Fovy:     45.0,
	}

	// Initialize player
	player = &entities.Player{
		Position: rl.Vector3{X: 0.0, Y: 0.0, Z: 0.0},
		Size:     2.0,
	}

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		UpdateGame()
		DrawGame()
	}

	rl.CloseWindow()
}
