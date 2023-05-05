package main

import (
	E "EasyEC2"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	//ebiten.SetVsyncEnabled(false)
	//ebiten.SetTPS(ebiten.SyncWithFPS)
	//ebiten.SetFPSMode(ebiten.)
	ebiten.SetWindowTitle("Render an image")
	if err := ebiten.RunGame(&E.Game{}); err != nil {
		log.Fatal(err)
	}
}

type Position struct {
	E.Entity
	X int
	Y int
}

type Image struct {
	E.Entity
	image *ebiten.Image
	op    ebiten.DrawImageOptions
}

type Components struct {
	Position *E.Component[*Position]
	Image    *E.Component[*Image]
}

var Comps *Components = &Components{}

func init() {

	Comps.Position = E.NewComp[*Position]()
	Comps.Image = E.NewComp[*Image]()

	E.DrawSystems = append(E.DrawSystems, draw)

	var err error
	image1, _, err := ebitenutil.NewImageFromFile("gopher.png")
	if err != nil {
		log.Fatal(err)
	}

	Ent1 := E.NewEntity()
	Comps.Position.Add(&Position{Ent1, 200, 200})
	Comps.Image.Add(&Image{Entity: Ent1, image: image1})

	Ent3 := E.NewEntity()
	Comps.Position.Add(&Position{Ent3, 300, 300})
	Comps.Image.Add(&Image{Entity: Ent3, image: image1})

	//Comps.Position.Remove(Ent3.Getid())

	Ent2 := E.NewEntity()
	Comps.Position.Add(&Position{Ent2, 100, 100})
	Comps.Image.Add(&Image{Entity: Ent2, image: image1})
}

func draw(screen *ebiten.Image) {
	for _, s := range Comps.Position.GetArr() {
		if a := Comps.Image.Get(s.Getid()); a != nil {
			a.op.GeoM.Reset()
			a.op.GeoM.Translate(float64(s.X), float64(s.Y))
			screen.DrawImage(a.image, &a.op)
		}
	}
}
