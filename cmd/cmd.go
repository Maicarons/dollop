package main

import (
	"fmt"
	"github.com/Maicarons/dollop/check"
	"github.com/Maicarons/dollop/est"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
)

func main() {
	TestCheck()
	TestEst()
}

func TestCheck() {
	ebiten.SetTPS(100)
	ebitengine, s2, err := check.Ebitengine()
	if err != nil {
		fmt.Printf("%v\n%v\n%s", ebitengine, s2, err.Error())
		return
	}
	fmt.Printf("%v %v\n", ebitengine, s2)
	b, s, err := check.Go()
	if err != nil {
		fmt.Println(b, s, err)
		return
	}
	fmt.Printf("%v %v\n", b, s)
}

func TestEst() {
	p1 := ebiten.NewImage(100, 100)
	p1.Fill(color.RGBA{R: 255, G: 0, B: 0, A: 255})
	o1 := est.NewObject(p1)
	s1 := est.NewScene([]*est.Object{o1, o1, o1, o1}, 1)
	s2 := est.NewScene([]*est.Object{o1, o1, o1, o1}, 2)
	settingjson := `
    {"info": {
        "GameName": "test",
        "Version": "1.0",
        "Platform": ["Windows"],
        "GoVersion": "1.20.5",
        "EbitengineVersion": "2.5.4"
    },
    "settings": {
     "ScreenSize":[
         300,300],
    "FullScreen": true ,
    "TPS": 30,
    "FPS": 30,
    "RunnableOnUnfocused": true,
    "ScreenClearedEveryFrame": true,
    "VsyncEnabled": true,
    "WindowClosingHandled": true,
    "WindowDecorated": true,
    "WindowFloating": true,
    "WindowIcon": "",
    "WindowPosition": [
        1,
        1
    ],
    "WindowResizingMode": 1,
    "WindowSize": [
        300,
        300
    ],
    "WindowSizeLimits": [
        300,
        300
    ],
    "WindowTitle": "test window"
    }}`
	setting, err := est.UnmarshalGameSetting([]byte(settingjson))
	if err != nil {
		log.Fatal(err.Error())
	}
	g1 := est.NewGame([]*est.Scene{s1, s2}, o1, &setting)
	err = ebiten.RunGame(g1)
	if err != nil {
		log.Fatal(err.Error())
	}
}
