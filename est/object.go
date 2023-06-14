package est

import "github.com/hajimehoshi/ebiten/v2"

type Object struct {
	Image *ebiten.Image
	Pos   [2]float64
	Do    func() error
}

func NewObject(image *ebiten.Image) *Object {
	return &Object{Image: image, Pos: [2]float64{0, 0}, Do: func() error {
		println("New Object created")
		return nil
	}}
}

func (O *Object) Draw(img *ebiten.Image) {
	Opt := ebiten.DrawImageOptions{}
	Opt.GeoM.Translate(O.Pos[0], O.Pos[1])
	img.DrawImage(O.Image, &Opt)
}
