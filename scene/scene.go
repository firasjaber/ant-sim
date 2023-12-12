package scene

import (
	"slices"
	"strconv"

	"github.com/firasjaber/ant-sim/config"
	"github.com/firasjaber/ant-sim/entity"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Scene struct {
	ants       []*entity.Ant
	food       []*entity.Food
	home       *entity.Home
	pheromones []*entity.Pheromone
	grid       [][]interface{}
}

// implement a grid to represent the entities
// the grid will be used to check for collisions

func NewScene() *Scene {
	rl.InitWindow(config.WindowWidth, config.WindowHeight, config.WindowTitle)
	rl.SetTargetFPS(config.TargetFPS)
	rl.SetExitKey(0)

	// init the grid with null based on the window size
	grid := make([][]interface{}, config.WindowWidth)
	for i := 0; i < config.WindowWidth; i++ {
		// init a grid row
		grid[i] = make([]interface{}, config.WindowHeight)
		for j := 0; j < config.WindowHeight; j++ {
			grid[i][j] = nil
		}
	}
	// span entites
	// spawn home
	homeXPos := config.WindowWidth/2 - (config.WindowWidth / 4)
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
	// for i := 0; i < config.FoodCount; i++ {
	// 	// create a new food
	// 	f := entity.NewFood(int32(rl.GetRandomValue(10, config.WindowWidth-10)), int32(rl.GetRandomValue(10, config.WindowHeight-10)))
	// 	// add the food to the food list
	// 	food = append(food, f)
	// }

	initFoodSpawnXPos := config.WindowWidth/2 + config.WindowWidth/4
	initFoodSpawnYPos := config.WindowHeight / 4
	lastFoodSpawnXPos := initFoodSpawnXPos + config.FoodCount
	lastFoodSpawnYPos := initFoodSpawnYPos + config.FoodCount
	for i := initFoodSpawnXPos; i < lastFoodSpawnXPos; i++ {
		for j := initFoodSpawnYPos; j < lastFoodSpawnYPos; j++ {
			// create a new food
			f := entity.NewFood(int32(i), int32(j))
			// add the food to the food list
			food = append(food, f)
		}
	}

	initFoodSpawnTwoXPos := (config.WindowWidth / 2) + config.WindowWidth/4
	initFoodSpawnTwoYPos := (config.WindowHeight / 4) + config.WindowHeight/2
	lastFoodSpawnTwoXPos := initFoodSpawnTwoXPos + config.FoodCount
	lastFoodSpawnTwoYPos := initFoodSpawnTwoYPos + config.FoodCount
	for i := initFoodSpawnTwoXPos; i < lastFoodSpawnTwoXPos; i++ {
		for j := initFoodSpawnTwoYPos; j < lastFoodSpawnTwoYPos; j++ {
			// create a new food
			f := entity.NewFood(int32(i), int32(j))
			// add the food to the food list
			food = append(food, f)
		}
	}

	return &Scene{ants: ants, food: food, home: home, pheromones: []*entity.Pheromone{}, grid: grid}
}

func (s *Scene) Run() {
	for !rl.WindowShouldClose() {
		// begin the drawing and clear the screen
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		// draw hud
		s.updateHud()

		// update the entities
		s.updateFood()
		s.updateAnts()
		s.updatePheromones()
		s.home.Update()

		// end the drawing
		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func (s *Scene) AddPheromones(posX int32, posY int32) {
	// create a new pheromone
	p := entity.NewPheromone(posX, posY)
	// add the pheromone to the grid
	s.grid[posX][posY] = p
	// add the pheromone to the pheromones list
	s.pheromones = append(s.pheromones, p)
}

func (s *Scene) updatePheromones() {
	// remove the pheromones with concentration 0
	for i, p := range s.pheromones {
		if p == nil {
			continue
		}
		if p.GetConcentration() <= 0 {
			// s.pheromones = append(s.pheromones[:i], s.pheromones[i+1:]...)
			s.pheromones = slices.Delete(s.pheromones, i, i+1)
			s.grid[p.GetXPos()][p.GetYPos()] = nil
		}
	}
	for _, p := range s.pheromones {
		p.Update()
	}
}

func (s *Scene) updateAnts() {
	for _, ant := range s.ants {
		nearbyPheromones := []*entity.Pheromone{}
		if ant.GetState() == entity.SEEKER {
			checkAntCollision(ant, s.food, s.home)
			// check if there is a pheromone nearby
			nearbyPheromones = getPheromoneNearby(ant, s)
		}
		if ant.GetState() == entity.RETURNER {
			// drop pheromones
			antPosX, antPosY := ant.GetPosition()
			s.AddPheromones(antPosX, antPosY)
			checkAntCollision(ant, s.food, s.home)
		}
		ant.Update(nearbyPheromones)
	}
}

func (s *Scene) updateFood() {
	// remove the food that has been collected
	// TODO: fix this shit bruv
	// for i, f := range s.food {
	// 	if f.IsDestroyed() {
	// 		// s.food = append(s.food[:i], s.food[i+1:]...)
	// 		// s.food = slices.Delete(s.food, i, i+1)
	// 		s.grid[f.GetXPos()][f.GetYPos()] = nil
	// 	}
	// }
	for _, f := range s.food {
		f.Update()
	}
}

func (s *Scene) updateHud() {
	spawnedFood := "Spawned food: " + strconv.FormatInt(int64(config.FoodCount), 10)
	collectedFood := "Collected food: " + strconv.FormatInt(int64(s.home.GetFoodCount()), 10)
	fpsText := "FPS: " + strconv.FormatInt(int64(rl.GetFPS()), 10)

	rl.DrawText(spawnedFood, 10, 10, 15, rl.LightGray)
	rl.DrawText(collectedFood, 10, 30, 15, rl.LightGray)
	rl.DrawText(fpsText, 10, 50, 15, rl.LightGray)
}

func checkAntCollision(ant *entity.Ant, food []*entity.Food, home *entity.Home) {
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
	if ant.GetState() == entity.RETURNER && rl.CheckCollisionRecs(ant.GetRectangle(), home.GetRectangle()) {
		ant.SetState(entity.SEEKER)
		ant.ClearPath()
		home.AddFood()
	}
}

func getPheromoneNearby(ant *entity.Ant, s *Scene) []*entity.Pheromone {
	antPosX, antPosY := ant.GetPosition()

	nearbyPheromones := []*entity.Pheromone{}
	// check if there is a pheromone nearby in the range of 5 pixels
	for i := antPosX - 1; i <= antPosX+1; i++ {
		for j := antPosY - 1; j <= antPosY+1; j++ {
			if s.grid[i][j] != nil && !(i == antPosX && j == antPosY) {
				nearbyPheromones = append(nearbyPheromones, s.grid[i][j].(*entity.Pheromone))
			}
		}
	}

	return nearbyPheromones
}
