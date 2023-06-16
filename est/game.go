package est

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

type Game struct {
	Scene       []Scene
	SceneIndex  uint
	NowScene    Scene
	RootObject  Object
	GameSetting GameSetting
}

func (g *Game) ChooseScene(Index uint) {
	a := g.Scene
	for _, scene := range a {
		if scene.Index == Index {
			g.NowScene = scene
			g.SceneIndex = Index
			return
		}
	}
}

func NewGame(scene []Scene, rootObject Object, gameSetting GameSetting) *Game {
	return &Game{Scene: scene, RootObject: rootObject, GameSetting: gameSetting}
}

func (g *Game) Update() error {
	err := g.NowScene.Update()
	if err != nil {
		return err
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.NowScene.Draw(screen)
}
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return int(g.GameSetting.Settings.ScreenSize[0]), int(g.GameSetting.Settings.ScreenSize[1])
}

func (g *Game) SceneSort() {
	t := g.Scene
	bubbleSort(t)
	for i, i2 := range t {
		log.Println(i, i2.Index)
	}
}

func bubbleSort(arr []Scene) {
	n := len(arr)
	for i := 0; i < n-1; i++ { // 外层循环控制排序轮数
		for j := 0; j < n-i-1; j++ { // 内层循环控制每轮比较次数
			if arr[j].Index > arr[j+1].Index { // 如果前一个元素大于后一个元素，则交换这两个元素
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func (g *Game) GameStart() error {

	err := ebiten.RunGame(g)
	if err != nil {
		return err
	}
	return nil
}
