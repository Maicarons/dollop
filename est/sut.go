package est

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
	Scene []*Scene
}

type Scene struct {
	Object []*Object
}

type Object struct {
	Image *ebiten.Image
}
