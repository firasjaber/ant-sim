package entity

import rl "github.com/gen2brain/raylib-go/raylib"

type Home struct {
	posX      int32
	posY      int32
	foodCount int32
}

func NewHome(posX int32, posY int32) *Home {
	return &Home{posX: posX, posY: posY, foodCount: 0}
}

func (h *Home) Update() {
	h.Draw()
}

func (h *Home) Draw() {
	rl.DrawRectangleLines(h.posX, h.posY, 15, 15, rl.Brown)
}

func (h *Home) GetPosition() rl.Vector2 {
	return rl.NewVector2(float32(h.posX), float32(h.posY))
}

func (h *Home) GetFoodCount() int32 {
	return h.foodCount
}

func (h *Home) GetRectangle() rl.Rectangle {
	return rl.NewRectangle(float32(h.posX), float32(h.posY), 15, 15)
}

func (h *Home) AddFood() {
	h.foodCount++
}
