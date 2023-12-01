package entity

import (
	"math/rand"

	"github.com/firasjaber/ant-sim/config"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Direction string

const (
	UP         Direction = "UP"
	DOWN       Direction = "DOWN"
	LEFT       Direction = "LEFT"
	RIGHT      Direction = "RIGHT"
	UP_LEFT    Direction = "UP_LEFT"
	UP_RIGHT   Direction = "UP_RIGHT"
	DOWN_LEFT  Direction = "DOWN_LEFT"
	DOWN_RIGHT Direction = "DOWN_RIGHT"
)

type AntState string

const (
	SEEKER   AntState = "SEEKER"
	RETURNER AntState = "RETURNER"
)

type Ant struct {
	posX    int32
	posY    int32
	currDir Direction
	state   AntState
	path    []rl.Vector2
}

func NewAnt(posX int32, posY int32) *Ant {
	// pick a random current direction
	directions := []Direction{UP, DOWN, LEFT, RIGHT, UP_LEFT, UP_RIGHT, DOWN_LEFT, DOWN_RIGHT}
	direction := directions[rl.GetRandomValue(0, 7)]
	return &Ant{posX: posX, posY: posY, currDir: direction, state: SEEKER, path: []rl.Vector2{}}
}

func (a *Ant) Move(dir Direction) {
	// append the previous pos to the path
	a.path = append(a.path, rl.Vector2{X: float32(a.posX), Y: float32(a.posY)})
	switch dir {
	case UP:
		a.posY -= 1
	case DOWN:
		a.posY += 1
	case LEFT:
		a.posX -= 1
	case RIGHT:
		a.posX += 1
	case UP_LEFT:
		a.posX -= 1
		a.posY -= 1
	case UP_RIGHT:
		a.posX += 1
		a.posY -= 1
	case DOWN_LEFT:
		a.posX -= 1
		a.posY += 1
	case DOWN_RIGHT:
		a.posX += 1
		a.posY += 1
	}
	a.currDir = dir
}

// wander randomly
func (a *Ant) Wander() {
	// if we are at the edge of the screen, turn around
	if a.posX <= 5 || a.posX >= (config.WindowWidth-5) || a.posY <= 5 || a.posY >= (config.WindowHeight-5) {
		oppisiteDir := getOppisiteDirection(a.currDir)
		a.Move(oppisiteDir)
		return
	}
	// pick if we should wander based on the wander rate
	// if we don't wander, move in the current direction
	wRand := int(config.WanderingRate * 100)
	cRand := rand.Intn(100)
	if cRand > wRand {
		a.Move(a.currDir)
		return
	}
	// pick a random direction that depends on the current ant direction so it can't rotate more than 90 degrees
	// which means elimate the possibliity of going in the opposite direction
	possibleDirections := getPossibleDirections(a.currDir)
	// print possible direction and curr dir
	direction := possibleDirections[rl.GetRandomValue(0, 1)]

	a.Move(direction)
}

func (a *Ant) FollowPathHome() {
	if len(a.path) > 0 {
		a.posX = int32(a.path[len(a.path)-1].X)
		a.posY = int32(a.path[len(a.path)-1].Y)
		a.path = a.path[:len(a.path)-1]
	}
}

func (a *Ant) Draw() {
	// draw the path
	for _, p := range a.path {
		rl.DrawCircle(int32(p.X), int32(p.Y), 1, rl.ColorAlpha(rl.Purple, 0.1))
	}
	// draw the ant
	if a.state == RETURNER {
		rl.DrawRectangle(a.posX, a.posY, 10, 10, rl.Yellow)
		return
	}
	rl.DrawRectangle(a.posX, a.posY, 10, 10, rl.White)
}

func (a *Ant) Update() {
	// if the ant is a seeker, wander
	// if the ant is a returner, follow the path to return home
	if a.state == SEEKER {
		a.Wander()
	} else if a.state == RETURNER {
		a.FollowPathHome()
	}
	a.Draw()
}

func (a *Ant) GetRectangle() rl.Rectangle {
	return rl.Rectangle{X: float32(a.posX), Y: float32(a.posY), Width: 10, Height: 10}
}

func (a *Ant) GetState() AntState {
	return a.state
}

func (a *Ant) SetState(state AntState) {
	a.state = state
}

func (a *Ant) ClearPath() {
	a.path = []rl.Vector2{}
}

func getOppisiteDirection(dir Direction) Direction {
	switch dir {
	case UP:
		return DOWN
	case DOWN:
		return UP
	case LEFT:
		return RIGHT
	case RIGHT:
		return LEFT
	case UP_LEFT:
		return DOWN_RIGHT
	case UP_RIGHT:
		return DOWN_LEFT
	case DOWN_LEFT:
		return UP_RIGHT
	case DOWN_RIGHT:
		return UP_LEFT
	}
	return dir
}

func getPossibleDirections(dir Direction) []Direction {
	switch dir {
	case UP:
		return []Direction{UP_LEFT, UP_RIGHT}
	case DOWN:
		return []Direction{DOWN_LEFT, DOWN_RIGHT}
	case LEFT:
		return []Direction{UP_LEFT, DOWN_LEFT}
	case RIGHT:
		return []Direction{UP_RIGHT, DOWN_RIGHT}
	case UP_LEFT:
		return []Direction{UP, LEFT}
	case UP_RIGHT:
		return []Direction{UP, RIGHT}
	case DOWN_LEFT:
		return []Direction{DOWN, LEFT}
	case DOWN_RIGHT:
		return []Direction{DOWN, RIGHT}
	}
	return []Direction{dir}
}
