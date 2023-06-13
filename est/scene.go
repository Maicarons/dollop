package est

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
