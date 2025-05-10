package entity

import (
	"github.com/firasjaber/ant-sim/config"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Home struct {
	posX      int32
	posY      int32
	foodCount int32
}

func NewHome(posX int32, posY int32) *Home {
	return &Home{posX: posX, posY: posY, foodCount: 0}
}

func (h *Home) Update() {
	if config.DrawHome {
		h.Draw()
	}
}

func (h *Home) Draw() {
	rl.DrawRectangle(h.posX, h.posY, 30, 30, rl.Brown)
}

func (h *Home) GetPosition() rl.Vector2 {
	return rl.NewVector2(float32(h.posX), float32(h.posY))
}

func (h *Home) GetFoodCount() int32 {
	return h.foodCount
}

func (h *Home) GetRectangle() rl.Rectangle {
	return rl.NewRectangle(float32(h.posX), float32(h.posY), 30, 30)
}

func (h *Home) AddFood() {
	h.foodCount++
}
