package est

import "github.com/hajimehoshi/ebiten/v2"

type Object struct {
	Image *ebiten.Image
	Pos   [2]int
}

func NewObject(image *ebiten.Image) *Object {
	return &Object{Image: image, Pos: [2]int{0, 0}}
}
