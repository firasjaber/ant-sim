package main

import (
	"github.com/firasjaber/ant-sim/entity"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func updateAnts(ants []*entity.Ant) {
	for _, ant := range ants {
		ant.Update()
	}
}

func init() {
	rl.InitWindow(200, 200, "raylib [core] example - basic window")
	rl.SetTargetFPS(60)
	rl.SetExitKey(0) 
}

func main() {
	// spawn an ant
	ant := entity.NewAnt(100, 100)

	// build ants list
	ants := []*entity.Ant{ant}

	for !rl.WindowShouldClose() {
		// begin the drawing and clear the screen
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		
		// update the entities
		updateAnts(ants)

		// end the drawing
		rl.EndDrawing()
	}

	rl.CloseWindow()
}