package examples

import (
	"time"

	"github.com/firasjaber/ant-sim/config"
	"github.com/firasjaber/ant-sim/entity"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type FoodScene struct {
	ants []*FoodAnt
	food []*entity.Food
	home *entity.Home
}

// implement a grid to represent the entities
// the grid will be used to check for collisions

func NewFoodScene() *FoodScene {
	rl.InitWindow(config.WindowWidth, config.WindowHeight, config.WindowTitle)
	rl.SetTargetFPS(config.TargetFPS)
	rl.SetExitKey(0)

	// span entites
	// spawn home
	homeXPos := config.WindowWidth/2 - (config.WindowWidth / 4)
	homeYPos := config.WindowHeight / 2
	// create the home
	home := entity.NewHome(int32(homeXPos), int32(homeYPos))

	// spawn ants

	ants := []*FoodAnt{}
	ant := NewFoodAnt(int32(homeXPos), int32(homeYPos))
	// add the ant to the ants list
	ants = append(ants, ant)

	// spawn food
	// loop through the food number range and create new food on random positions
	food := []*entity.Food{}

	// random food positions
	for i := 0; i < 20; i++ {
		// create a new food
		f := entity.NewFood(int32(rl.GetRandomValue(0, config.WindowWidth)), int32(rl.GetRandomValue(0, config.WindowHeight)))
		// add the food to the food list
		food = append(food, f)
	}

	return &FoodScene{ants: ants, food: food, home: home}
}

func (s *FoodScene) Run() {
	time.Sleep(1 * time.Second)

	for !rl.WindowShouldClose() {
		// begin the drawing and clear the screen
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		// update the entities
		s.updateFood()
		s.updateAnts()
		s.home.Update()

		// end the drawing
		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func (s *FoodScene) updateAnts() {
	for _, ant := range s.ants {
		if ant.GetState() == SEEKER {
			// check collision with food
			checkAntCollision(ant, s.food, s.home)
		}
		if ant.GetState() == RETURNER {
			// check collision with home
			checkAntCollision(ant, s.food, s.home)
		}
		ant.Update()
	}
}

func (s *FoodScene) updateFood() {
	for _, f := range s.food {
		if f.IsDestroyed() {
			continue
		}
		f.Update()
	}
}

func checkAntCollision(ant *FoodAnt, food []*entity.Food, home *entity.Home) {
	if ant.GetState() == SEEKER {
		for _, f := range food {
			// TODO: remove the destroyed food from the food list so we don't waste collision checks
			if f.IsDestroyed() {
				continue
			}
			// ant collide with a food, change the ant state to RETURNER (collect it)
			if rl.CheckCollisionRecs(ant.GetRectangle(), f.GetRectangle()) {
				ant.SetState(RETURNER)
				f.Destroy()
			}
		}
	}

	// ant is a RETURNER and collide with the home, change the ant state to SEEKER (go find more food)
	if ant.GetState() == RETURNER && rl.CheckCollisionRecs(ant.GetRectangle(), home.GetRectangle()) {
		ant.SetState(SEEKER)
		ant.ClearPath()
		home.AddFood()
	}
}
