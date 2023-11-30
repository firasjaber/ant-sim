package entity

import rl "github.com/gen2brain/raylib-go/raylib"

type Food struct {
	posX int32
	posY int32
}

func NewFood(posX int32, posY int32) *Food {
	return &Food{posX: posX, posY: posY}
}

func (f *Food) Update() {
	f.Draw()
}

func (f *Food) Draw() {
	rl.DrawRectangle(f.posX, f.posY, 5, 5, rl.Green)
}

func (f *Food) GetPosition() rl.Vector2 {
	return rl.Vector2{X: float32(f.posX), Y: float32(f.posY)}
}