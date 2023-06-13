package main

import (
	"fmt"
	"github.com/Maicarons/dollop/check"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
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
