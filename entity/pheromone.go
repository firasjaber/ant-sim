package entity

import (
	"github.com/firasjaber/ant-sim/config"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Pheromone struct {
	xPos          int32
	yPos          int32
	droppedAt     float64
	concentration float32
}

func NewPheromone(xPos int32, yPos int32) *Pheromone {
	return &Pheromone{xPos: xPos, yPos: yPos, droppedAt: rl.GetTime(), concentration: 1}
}

func (p *Pheromone) Update() {
	// lower the concentration of the pheromone based on the time it was dropped and the current time
	p.concentration = float32(1 - (rl.GetTime()-p.droppedAt)*config.PheromoneDecayRate)
	p.Draw()
}

func (p *Pheromone) Draw() {
	rl.DrawCircle(p.xPos, p.yPos, 1, rl.ColorAlpha(rl.Pink, p.concentration/2))
	// draw a bigger circle around the pheromone with lower indensity
	// rl.DrawCircle(p.xPos, p.yPos, 4, rl.ColorAlpha(rl.Pink, p.concentration/10))
}

func (p *Pheromone) GetConcentration() float32 {
	return p.concentration
}

func (p *Pheromone) GetXPos() int32 {
	return p.xPos
}

func (p *Pheromone) GetYPos() int32 {
	return p.yPos
}
