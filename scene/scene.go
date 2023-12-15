package scene

import (
	"strconv"
	"sync"
	"time"

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

func NewScene(mapId int) *Scene {
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
	entities := GetEntitiesByMapId(mapId)

	return &Scene{ants: entities.Ants, food: entities.Food, home: entities.Home, pheromones: []*entity.Pheromone{}, grid: grid}
}

func (s *Scene) Run() {
	time.Sleep(1 * time.Second)
	for !rl.WindowShouldClose() {
		// begin the drawing and clear the screen
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		// draw hud
		if config.DrawHUD {
			s.updateHud()
		}

		// update the entities
		s.updateFood()
		s.updateAnts()
		s.updatePheromonesOptimized()
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

func (s *Scene) updatePheromonesOptimized() {
	//log the frame number
	numPheromones := len(s.pheromones)
	if numPheromones == 0 {
		return
	}

	numChunks := 1
	var wg sync.WaitGroup

	chunkSize := (numPheromones + numChunks - 1) / numChunks

	for i := 0; i < numPheromones; i += chunkSize {
		end := i + chunkSize
		if end > numPheromones {
			end = numPheromones
		}

		wg.Add(1)
		go s.processPheromoneChunk(i, end, &wg)
	}

	wg.Wait()
	for _, p := range s.pheromones {
		if p.GetConcentration() > 0 {
			p.Update()
		}
	}
}

func (s *Scene) processPheromoneChunk(start, end int, wg *sync.WaitGroup) {
	for i := start; i < end; i++ {
		p := s.pheromones[i]
		if p != nil {
			if p.GetConcentration() <= 0 {
				gridP := s.grid[p.GetXPos()][p.GetYPos()]
				if gridP != nil && gridP.(*entity.Pheromone).GetConcentration() <= p.GetConcentration() {
					s.grid[p.GetXPos()][p.GetYPos()] = nil
				}
			}
		}
	}
	wg.Done()
}

func (s *Scene) updateAnts() {
	for _, ant := range s.ants {
		nearbyPheromones := []*entity.Pheromone{}
		if ant.GetState() == entity.SEEKER {
			// check collision with food
			checkAntCollision(ant, s.food, s.home)
			// check if there is a pheromone nearby
			nearbyPheromones = getPheromoneNearby(ant, s)
		}
		if ant.GetState() == entity.RETURNER {
			// drop pheromones
			antPosX, antPosY := ant.GetPosition()
			s.AddPheromones(antPosX, antPosY)
			// check collision with home
			checkAntCollision(ant, s.food, s.home)
		}
		ant.Update(nearbyPheromones)
	}
}

func (s *Scene) updateFood() {
	for _, f := range s.food {
		if f.IsDestroyed() {
			continue
		}
		f.Update()
	}
}

func (s *Scene) updateHud() {
	collectedFood := "Collected food: " + strconv.FormatInt(int64(s.home.GetFoodCount()), 10)
	fpsText := "FPS: " + strconv.FormatInt(int64(rl.GetFPS()), 10)

	rl.DrawText(collectedFood, 10, 10, 15, rl.LightGray)
	rl.DrawText(fpsText, 10, 30, 15, rl.LightGray)
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
