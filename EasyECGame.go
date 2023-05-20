package main

import (
	E "EasyEC2"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"time"
	"math/rand"
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

type Player struct {
	E.Entity
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
	Player *E.Component[*Player]
}

var Comps *Components = &Components{}

func initComps() {
	Comps.Position = E.NewComp[*Position]()
	Comps.Image = E.NewComp[*Image]()
	Comps.Player = E.NewComp[*Player]()
}

func initSystems() {

	E.Systems = append(
		E.Systems,
		spawner2,
		TPS,
		//deleter2,
		move,
	)

	E.DrawSystems = append(
		E.DrawSystems,
		draw,
	)
}

var image1 *ebiten.Image

func init() {
	initComps()
	initSystems()

	var err error
	image1, _, err = ebitenutil.NewImageFromFile("gopher.png")
	if err != nil {
		log.Fatal(err)
	}

	Ent1 := E.NewEntity()
	addBasic(Ent1, 200, 200, image1)
	//Comps.Position.Add(&Position{Ent1, 200, 200})
	//Comps.Image.Add(&Image{Entity: Ent1, image: image1})

	Ent3 := E.NewEntity()
	Comps.Player.Add(&Player{Ent3})
	Comps.Position.Add(&Position{Ent3, 150, 150})
	Comps.Image.Add(&Image{Entity: Ent3, image: image1})

	//Comps.Position.Remove(Ent1.Getid())

	Ent2 := E.NewEntity()
	Comps.Position.Add(&Position{Ent2, 100, 100})
	Comps.Image.Add(&Image{Entity: Ent2, image: image1})

	//go spawner()
	//go deleter()

	
}

func addBasic(entity E.Entity, x, y int, image *ebiten.Image) {
	Comps.Position.Add(&Position{entity, x, y})
	Comps.Image.Add(&Image{Entity: entity, image: image})
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

func spawner() {
	for true {
		time.Sleep(2*time.Second)
		//E.Lock(Comps.Position)
		addBasic(E.NewEntity(), rand.Intn(500), rand.Intn(500), image1)
		//E.Unlock(Comps.Position)
	}
}

func deleter() {
	for true {
		time.Sleep(2*time.Second)
		E.Delete(rand.Intn(20))
	}
}

var deletetimer int
func deleter2() {
	deletetimer++
	if deletetimer > 50 {
		E.Delete(rand.Intn(20))
		deletetimer = 0
	}
}

var spawntimer int
func spawner2() {
	spawntimer++
	if spawntimer > 50 {
		addBasic(E.NewEntity(), rand.Intn(500), rand.Intn(500), image1)
		spawntimer = 0
	}
}

func TPS() {
	print(int(ebiten.ActualTPS()))
}

func move() {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		for _,s := range Comps.Position.GetArr() {
			if !Comps.Player.Contains(s.Getid()) {
				s.Y = s.Y + 5
			}
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		for _,s := range Comps.Position.GetArr() {
			if !Comps.Player.Contains(s.Getid()) {
				s.Y = s.Y - 5
			}
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		for _,s := range Comps.Position.GetArr() {
			if !Comps.Player.Contains(s.Getid()) {
				s.X = s.X - 5
			}
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		for _,s := range Comps.Position.GetArr() {
			if !Comps.Player.Contains(s.Getid()) {
				s.X = s.X + 5
			}
		}
	}
}
