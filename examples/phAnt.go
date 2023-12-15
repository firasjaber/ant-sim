package examples

import (
	lls "github.com/emirpasic/gods/stacks/linkedliststack"
	"github.com/firasjaber/ant-sim/config"
	"github.com/firasjaber/ant-sim/entity"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type PhAnt struct {
	posX    int32
	posY    int32
	currDir Direction
	state   AntState
	path    *lls.Stack
}

func NewPhAnt(posX int32, posY int32) *PhAnt {
	// pick a random current direction
	if posX != 20 {
		return &PhAnt{posX: posX, posY: posY, currDir: DOWN, state: RETURNER, path: lls.New()}
	}
	return &PhAnt{posX: posX, posY: posY, currDir: RIGHT, state: SEEKER, path: lls.New()}
}

func (a *PhAnt) Move(dir Direction) {
	// append the previous pos to the path
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
func (a *PhAnt) Wander() {
	// if we are at the edge of the screen, turn around
	if a.posX <= 5 || a.posX >= (config.WindowWidth-5) || a.posY <= 5 || a.posY >= (config.WindowHeight-5) {
		oppisiteDir := getOppisiteDirection(a.currDir)
		a.Move(oppisiteDir)
		return
	}
	if true {
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

func (a *PhAnt) FollowPathHome() {
	a.Move(DOWN)
	if a.path.Size() > 0 {
		lastPos, _ := a.path.Peek()
		a.posX = int32(lastPos.(rl.Vector2).X)
		a.posY = int32(lastPos.(rl.Vector2).Y)
		a.path.Pop()
	}
}

func (a *PhAnt) Draw() {
	// draw the path
	// for _, p := range a.path {
	// 	rl.DrawCircle(int32(p.X), int32(p.Y), 1, rl.ColorAlpha(rl.Purple, 0.1))
	// }
	// draw the ant
	if a.state == RETURNER {
		rl.DrawRectangle(a.posX, a.posY, 10, 10, rl.Yellow)
		return
	}
	rl.DrawRectangle(a.posX, a.posY, 10, 10, rl.White)
}

func (a *PhAnt) Update(pheromones []*entity.Pheromone) {
	if a.state == SEEKER && len(pheromones) > 0 {
		// pick the pheromone with the highest concentration
		pWithHighestConcentration := pheromones[0]
		for _, p := range pheromones {
			if p.GetConcentration() < pWithHighestConcentration.GetConcentration() {
				pWithHighestConcentration = p
			}
		}
		pPosX := pWithHighestConcentration.GetXPos()
		pPosY := pWithHighestConcentration.GetYPos()
		if a.getDirOutOfTargetPosition(pPosX, pPosY) == getOppisiteDirection(a.currDir) {
			a.Wander()
		} else {
			// move the ant towards the pheromone
			a.currDir = a.getDirOutOfTargetPosition(pPosX, pPosY)
			a.posX = pPosX
			a.posY = pPosY
			a.path.Push(rl.Vector2{X: float32(a.posX), Y: float32(a.posY)})
		}

	} else if a.state == SEEKER && len(pheromones) == 0 {
		a.Wander()
	} else if a.state == RETURNER {
		a.FollowPathHome()
	}

	a.Draw()
}

func (a *PhAnt) GetRectangle() rl.Rectangle {
	return rl.Rectangle{X: float32(a.posX), Y: float32(a.posY), Width: 10, Height: 10}
}

func (a *PhAnt) GetState() AntState {
	return a.state
}

func (a *PhAnt) GetPosition() (int32, int32) {
	return a.posX, a.posY
}

func (a *PhAnt) SetState(state AntState) {
	a.state = state
}

func (a *PhAnt) ClearPath() {
	a.path.Clear()
}

func (a *PhAnt) getDirOutOfTargetPosition(targetX int32, targetY int32) Direction {
	if a.posX < targetX && a.posY < targetY {
		return DOWN_RIGHT
	} else if a.posX < targetX && a.posY > targetY {
		return UP_RIGHT
	} else if a.posX > targetX && a.posY < targetY {
		return DOWN_LEFT
	} else if a.posX > targetX && a.posY > targetY {
		return UP_LEFT
	} else if a.posX == targetX && a.posY < targetY {
		return DOWN
	} else if a.posX == targetX && a.posY > targetY {
		return UP
	} else if a.posX < targetX && a.posY == targetY {
		return RIGHT
	} else if a.posX > targetX && a.posY == targetY {
		return LEFT
	}
	return a.currDir
}

func (a *PhAnt) IsPreviousPosition(posX int32, posY int32) bool {
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
