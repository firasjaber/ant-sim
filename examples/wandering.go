package examples

import (
	"time"

	"github.com/firasjaber/ant-sim/config"
	"github.com/firasjaber/ant-sim/entity"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type WanderingScene struct {
	home *entity.Home
	ants []*DumbAnt
}

func NewWanderingScene() *WanderingScene {
	rl.InitWindow(config.WindowWidth, config.WindowHeight, config.WindowTitle)
	rl.SetTargetFPS(config.TargetFPS)
	rl.SetExitKey(0)

	// spawn home
	homeXPos := config.WindowWidth / 2
	homeYPos := config.WindowHeight / 2
	// create the home
	home := entity.NewHome(int32(homeXPos), int32(homeYPos))

	// spawn ants
	ants := []*DumbAnt{}

	ant := NewDumbAnt(int32(homeXPos), int32(homeYPos))
	ants = append(ants, ant)

	return &WanderingScene{home: home, ants: ants}
}

func (s *WanderingScene) Run() {
	time.Sleep(2 * time.Second)

	for !rl.WindowShouldClose() {
		// begin the drawing and clear the screen
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		// update entities
		s.home.Update()
		s.updateAnts()

		// end the drawing
		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func (s *WanderingScene) updateAnts() {
	for _, ant := range s.ants {
		ant.Update()
	}
}
