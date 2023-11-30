package main

import (
	"github.com/firasjaber/ant-sim/config"
	"github.com/firasjaber/ant-sim/entity"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func updateAnts(ants []*entity.Ant) {
	for _, ant := range ants {
		ant.Update()
	}
}

func init() {
	rl.InitWindow(config.WindowWidth, config.WindowHeight, config.WindowTitle)
	rl.SetTargetFPS(config.TargetFPS)
	rl.SetExitKey(0) 
}

func main() {

	// spawn ants
	// loop through the ants number range and create new ants
	ants := []*entity.Ant{}
	for i := 0; i < config.AntsCount; i++ {
		// create a new ant
		ant := entity.NewAnt(100, 100)
		// add the ant to the ants list
		ants = append(ants, ant)
	} 

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