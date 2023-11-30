package entity

import rl "github.com/gen2brain/raylib-go/raylib"

type Direction string

const (
	UP Direction = "UP"
	DOWN = "DOWN"
	LEFT = "LEFT"
	RIGHT = "RIGHT"
)

type Ant struct {
	posX int32
	posY int32
}

func NewAnt(posX int32, posY int32) *Ant {
	return &Ant{posX: posX, posY: posY}
}

func (a *Ant) Move(dir Direction) {
	switch dir {
	case "UP":
		a.posY++
	case "DOWN":
		a.posY--
	case "LEFT":
		a.posX--
	case "RIGHT":
		a.posX++
	}
}

// wander randomly
func (a *Ant) Wander() {
	// pick a random direction
	directions := []Direction{UP, DOWN, LEFT, RIGHT}
	direction := directions[rl.GetRandomValue(0, 3)]
	a.Move(direction)
}

func (a *Ant) Draw() {
	rl.DrawRectangle(a.posX, a.posY, 10, 10, rl.White)
}

func (a *Ant) Update() {
	a.Wander()
	a.Draw()
}