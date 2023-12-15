package examples

import (
	"math/rand"

	lls "github.com/emirpasic/gods/stacks/linkedliststack"
	"github.com/firasjaber/ant-sim/config"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type FoodAnt struct {
	posX    int32
	posY    int32
	currDir Direction
	state   AntState
	path    *lls.Stack
}

func NewFoodAnt(posX int32, posY int32) *FoodAnt {
	// pick a random current direction
	directions := []Direction{UP, DOWN, LEFT, RIGHT, UP_LEFT, UP_RIGHT, DOWN_LEFT, DOWN_RIGHT}
	direction := directions[rl.GetRandomValue(0, 7)]

	return &FoodAnt{posX: posX, posY: posY, currDir: direction, state: SEEKER, path: lls.New()}
}

func (a *FoodAnt) Move(dir Direction) {
	// append the previous pos to the path
	a.path.Push(rl.Vector2{X: float32(a.posX), Y: float32(a.posY)})
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
func (a *FoodAnt) Wander() {
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
		// if true {
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

func (a *FoodAnt) FollowPathHome() {
	if a.path.Size() > 0 {
		lastPos, _ := a.path.Peek()
		a.posX = int32(lastPos.(rl.Vector2).X)
		a.posY = int32(lastPos.(rl.Vector2).Y)
		a.path.Pop()
	}
}

func (a *FoodAnt) Draw() {
	if a.state == RETURNER {
		rl.DrawRectangle(a.posX, a.posY, 10, 10, rl.Yellow)
		return
	}
	rl.DrawRectangle(a.posX, a.posY, 10, 10, rl.White)
}

func (a *FoodAnt) Update() {
	if a.state == SEEKER {
		a.Wander()
	} else if a.state == RETURNER {
		a.FollowPathHome()
	}
	a.Draw()
}

func (a *FoodAnt) GetRectangle() rl.Rectangle {
	return rl.Rectangle{X: float32(a.posX), Y: float32(a.posY), Width: 10, Height: 10}
}

func (a *FoodAnt) GetState() AntState {
	return a.state
}

func (a *FoodAnt) GetPosition() (int32, int32) {
	return a.posX, a.posY
}

func (a *FoodAnt) SetState(state AntState) {
	a.state = state
}

func (a *FoodAnt) ClearPath() {
	a.path.Clear()
}

func (a *FoodAnt) IsPreviousPosition(posX int32, posY int32) bool {
	prevPosX := posX
	prevPosY := posY
	switch a.currDir {
	case UP:
		prevPosX = posX
		prevPosY = posY - 1
	case DOWN:
		prevPosX = posX
		prevPosY = posY + 1
	case LEFT:
		prevPosX = posX - 1
		prevPosY = posY
	case RIGHT:
		prevPosX = posX + 1
		prevPosY = posY
	case UP_LEFT:
		prevPosX = posX - 1
		prevPosY = posY - 1
	case UP_RIGHT:
		prevPosX = posX + 1
		prevPosY = posY - 1
	case DOWN_LEFT:
		prevPosX = posX - 1
		prevPosY = posY + 1
	case DOWN_RIGHT:
		prevPosX = posX + 1
		prevPosY = posY + 1
	}
	if prevPosX == a.posX && prevPosY == a.posY {
		return true
	}
	return false
}
