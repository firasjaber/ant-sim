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

func updateFood(food []*entity.Food) {
	for _, f := range food {
		f.Update()
	}
}

func init() {
	rl.InitWindow(config.WindowWidth, config.WindowHeight, config.WindowTitle)
	rl.SetTargetFPS(config.TargetFPS)
	rl.SetExitKey(0) 
}

func main() {
	homeXPos := config.WindowWidth / 2
	homeYPos := config.WindowHeight / 2
	// create the home
	home := entity.NewHome(int32(homeXPos), int32(homeYPos))

	// spawn ants
	// loop through the ants number range and create new ants
	ants := []*entity.Ant{}
	for i := 0; i < config.AntsCount; i++ {
		// create a new ant
		ant := entity.NewAnt(int32(homeXPos), int32(homeYPos))
		// add the ant to the ants list
		ants = append(ants, ant)
	}

	// spawn food
	// loop through the food number range and create new food on random positions
	food := []*entity.Food{}
	for i := 0; i < config.FoodCount; i++ {
	 		// create a new food
		f := entity.NewFood(int32(rl.GetRandomValue(10, 190)), int32(rl.GetRandomValue(10, 190)))
		// add the food to the food list
		food = append(food, f)
	}

	for !rl.WindowShouldClose() {
		// begin the drawing and clear the screen
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		
		// update the entities
		updateFood(food)
		updateAnts(ants, food, home)
		home.Update()

		// end the drawing
		rl.EndDrawing()
	}

	rl.CloseWindow()
}