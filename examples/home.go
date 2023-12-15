package examples

import (
	"time"

	"github.com/firasjaber/ant-sim/config"
	"github.com/firasjaber/ant-sim/entity"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type HomeScene struct {
	home *entity.Home
}

func NewHomeScene() *HomeScene {
	rl.InitWindow(config.WindowWidth, config.WindowHeight, config.WindowTitle)
	rl.SetTargetFPS(config.TargetFPS)
	rl.SetExitKey(0)

	// spawn home
	homeXPos := config.WindowWidth / 2
	homeYPos := config.WindowHeight / 2
	// create the home
	home := entity.NewHome(int32(homeXPos), int32(homeYPos))

	return &HomeScene{home: home}
}

func (s *HomeScene) Run() {
	time.Sleep(2 * time.Second)

	for !rl.WindowShouldClose() {
		// begin the drawing and clear the screen
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		s.home.Update()

		// end the drawing
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
