package main

import (
	"strconv"

	"github.com/firasjaber/ant-sim/config"
	"github.com/firasjaber/ant-sim/entity"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func updateAnts(ants []*entity.Ant, food []*entity.Food, home *entity.Home) {
	for _, ant := range ants {
		checkFoodCollision(ant, food, home)
		ant.Update()
	}
}

func updateFood(food []*entity.Food) {
	for _, f := range food {
		f.Update()
	}
}

func updateHud(home *entity.Home, fps int32) {
	spawnedFood := "Spawned food: " + strconv.FormatInt(int64(config.FoodCount), 10)
	collectedFood := "Collected food: " + strconv.FormatInt(int64(home.GetFoodCount()), 10)
	fpsText := "FPS: " + strconv.FormatInt(int64(fps), 10)

	rl.DrawText(spawnedFood, 10, 10, 15, rl.LightGray)
	rl.DrawText(collectedFood, 10, 30, 15, rl.LightGray)
	rl.DrawText(fpsText, 10, 50, 15, rl.LightGray)
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
		f := entity.NewFood(int32(rl.GetRandomValue(10, config.WindowWidth-10)), int32(rl.GetRandomValue(10, config.WindowHeight-10)))
		// add the food to the food list
		food = append(food, f)
	}

	for !rl.WindowShouldClose() {
		// begin the drawing and clear the screen
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		// draw hud
		updateHud(home, rl.GetFPS())

		// update the entities
		updateFood(food)
		updateAnts(ants, food, home)
		home.Update()

		// end the drawing
		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func checkFoodCollision(ant *entity.Ant, food []*entity.Food, home *entity.Home) {
	if ant.GetState() == entity.SEEKER {
		for _, f := range food {
			// TODO: remove the destroyed food from the food list so we don't waste collision checks
			if f.IsDestroyed() {
				continue
			}
			// ant collide with a food, change the ant state to RETURNER (collect it)
			if rl.CheckCollisionRecs(ant.GetRectangle(), f.GetRectangle()) {
				ant.SetState(entity.RETURNER)
				f.Destroy()
			}
		}
	}

	// ant is a RETURNER and collide with the home, change the ant state to SEEKER (go find more food)
	if ant.GetState() == entity.RETURNER && rl.CheckCollisionCircleRec(home.GetPosition(), 15, ant.GetRectangle()) {
		ant.SetState(entity.SEEKER)
		ant.ClearPath()
		home.AddFood()
	}
}
