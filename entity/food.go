package entity

import rl "github.com/gen2brain/raylib-go/raylib"

type Food struct {
	posX      int32
	posY      int32
	destroyed bool
}

func NewFood(posX int32, posY int32) *Food {
	return &Food{posX: posX, posY: posY, destroyed: false}
}

func (f *Food) Update() {
	f.Draw()
}

func (f *Food) Draw() {
	if f.destroyed {
		return
	}
	rl.DrawRectangle(f.posX, f.posY, 5, 5, rl.Green)
}

func (f *Food) GetPosition() rl.Vector2 {
	return rl.Vector2{X: float32(f.posX), Y: float32(f.posY)}
}

func (f *Food) GetXPos() int32 {
	return f.posX
}

func (f *Food) GetYPos() int32 {
	return f.posY
}

func (f *Food) Destroy() {
	f.destroyed = true
}

func (f *Food) IsDestroyed() bool {
	return f.destroyed
}

func (a *Food) GetRectangle() rl.Rectangle {
	return rl.Rectangle{X: float32(a.posX), Y: float32(a.posY), Width: 5, Height: 5}
}
