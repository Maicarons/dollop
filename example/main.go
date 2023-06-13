package main

import (
	"github.com/Maicarons/dollop/est"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png"
	"log"
	"os"
)

var img *ebiten.Image
var Obj *est.Object

func init() {
	var err error
	getwd, err := os.Getwd()
	if err != nil {
		return
	}
	img, _, err = ebitenutil.NewImageFromFile(getwd + "/example/res/gopher.png")
	Obj = est.NewObject(img)
	if err != nil {
		log.Fatal(err)
	}
}

type Game est.Game

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(Obj.Image, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Render an image")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
