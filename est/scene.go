package est

import "github.com/hajimehoshi/ebiten/v2"

func init() {

}

type Scene struct {
	Object []*Object
	Index  uint
}

func NewScene(object []*Object, index uint) *Scene {
	return &Scene{Object: object, Index: index}
}

/*,Index: uint(rand.New(rand.NewSource(time.Now().UnixNano())).Uint32()*/

func (g *Scene) Update() error {
	for _, object := range g.Object {
		err := object.Do()
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *Scene) Draw(screen *ebiten.Image) {
	for _, object := range g.Object {
		object.Draw(screen)

	}
}
