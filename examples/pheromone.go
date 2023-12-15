package examples

import (
	"slices"
	"time"

	"github.com/firasjaber/ant-sim/config"
	"github.com/firasjaber/ant-sim/entity"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type PhScene struct {
	ants       []*PhAnt
	pheromones []*entity.Pheromone
	grid       [][]interface{}
}

// implement a grid to represent the entities
// the grid will be used to check for collisions

func NewPhScene() *PhScene {
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
	// spawn ants
	// loop through the ants number range and create new ants
	ants := []*PhAnt{}
	// create a new ant
	ant1 := NewPhAnt(20, 150)
	ant2 := NewPhAnt(150, 20)
	// add the ant to the ants list
	ants = append(ants, ant1)
	ants = append(ants, ant2)

	return &PhScene{ants: ants, pheromones: []*entity.Pheromone{}, grid: grid}
}

func (s *PhScene) Run() {
	time.Sleep(1 * time.Second)

	for !rl.WindowShouldClose() {
		// begin the drawing and clear the screen
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		// draw hud

		// update the entities
		s.updateAnts()
		s.updatePheromones()

		// end the drawing
		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func (s *PhScene) AddPheromones(posX int32, posY int32) {
	// create a new pheromone
	p := entity.NewPheromone(posX, posY)
	// add the pheromone to the grid
	s.grid[posX][posY] = p
	// add the pheromone to the pheromones list
	s.pheromones = append(s.pheromones, p)
}

func (s *PhScene) updatePheromones() {
	// remove the pheromones with concentration 0
	for i, p := range s.pheromones {
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

func (s *PhScene) updateAnts() {
	for _, ant := range s.ants {
		nearbyPheromones := []*entity.Pheromone{}
		if ant.GetState() == SEEKER {
			// check if there is a pheromone nearby
			nearbyPheromones = getPheromoneNearby(ant, s)
		}
		if ant.GetState() == RETURNER {
			// drop pheromones
			antPosX, antPosY := ant.GetPosition()
			s.AddPheromones(antPosX, antPosY)
		}
		ant.Update(nearbyPheromones)
	}
}

func getPheromoneNearby(ant *PhAnt, s *PhScene) []*entity.Pheromone {
	antPosX, antPosY := ant.GetPosition()

	// eligblePheromonesCoords := ant.GetPossiblePheromonesCoordsToFollow()
	nearbyPheromones := []*entity.Pheromone{}
	// check if there is a pheromone nearby in the range of 5 pixels
	for i := antPosX - 1; i <= antPosX+1; i++ {
		for j := antPosY - 1; j <= antPosY+1; j++ {
			if s.grid[i][j] != nil && !(i == antPosX && j == antPosY) {
				// currCoord := utils.Coord{X: i, Y: j}
				// if isOneOfCoords(currCoord, eligblePheromonesCoords) {
				nearbyPheromones = append(nearbyPheromones, s.grid[i][j].(*entity.Pheromone))
				// }
			}
		}
	}
	return nearbyPheromones
}
