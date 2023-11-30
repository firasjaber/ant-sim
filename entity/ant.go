package entity

import (
	"log"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const wanderRate = 0.1

type Direction string
const (
	UP Direction = "UP"
	DOWN Direction = "DOWN"
	LEFT Direction = "LEFT"
	RIGHT Direction = "RIGHT"
	UP_LEFT Direction = "UP_LEFT"
	UP_RIGHT Direction = "UP_RIGHT"
	DOWN_LEFT Direction = "DOWN_LEFT"
	DOWN_RIGHT Direction = "DOWN_RIGHT"
)

type Ant struct {
	posX int32
	posY int32
	currDir Direction
}

func NewAnt(posX int32, posY int32) *Ant {
	// pick a random current direction
	directions := []Direction{UP, DOWN, LEFT, RIGHT, UP_LEFT, UP_RIGHT, DOWN_LEFT, DOWN_RIGHT}
	direction := directions[rl.GetRandomValue(0, 7)]
	return &Ant{posX: posX, posY: posY, currDir: direction}
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
	a.currDir = dir
}

// wander randomly
func (a *Ant) Wander() {
	// pick if we should wander based on the wander rate
	// if we don't wander, move in the current direction
	wRand := int(wanderRate * 10)
	cRand := rand.Intn(10)
	if cRand > wRand {
		a.Move(a.currDir)
		log.Println("moving in current direction")
		return
	}
	log.Println("wandering")
	// pick a random direction that depends on the current ant direction so it can't rotate more than 90 degrees
	// which means elimate the possibliity of going in the opposite direction
	possibleDirections := []Direction{}
	switch a.currDir {
	case UP:
		possibleDirections = []Direction{UP_LEFT, UP_RIGHT}
	case DOWN:
		possibleDirections = []Direction{DOWN_LEFT, DOWN_RIGHT}
	case LEFT:
		possibleDirections = []Direction{UP_LEFT, DOWN_LEFT}
	case RIGHT:
		possibleDirections = []Direction{UP_RIGHT, DOWN_RIGHT}
	case UP_LEFT:
		possibleDirections = []Direction{UP, LEFT}
	case UP_RIGHT:
		possibleDirections = []Direction{UP, RIGHT}
	case DOWN_LEFT:
		possibleDirections = []Direction{DOWN, LEFT}
	case DOWN_RIGHT:
		possibleDirections = []Direction{DOWN, RIGHT}
	}
	// print possible direction and curr dir
	direction := possibleDirections[rl.GetRandomValue(0, 1)]

	a.Move(direction)
}

func (a *Ant) Draw() {
	rl.DrawRectangle(a.posX, a.posY, 10, 10, rl.White)
}

func (a *Ant) Update() {
	a.Wander()
	a.Draw()
}