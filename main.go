package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Ninja struct {
	AttackImage   *ebiten.Image
	DeadImage     *ebiten.Image
	IdleImage     *ebiten.Image
	ItemImage     *ebiten.Image
	JumpImage     *ebiten.Image
	Special1Image *ebiten.Image
	Special2Image *ebiten.Image
	WalkImage     *ebiten.Image
}
type Game struct {
	Ninja *Ninja
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{120, 180, 255, 255})
	screen.DrawImage(g.Ninja.IdleImage, &ebiten.DrawImageOptions{})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ebiten.WindowSize()
}

func NewNinja() *Ninja {
	basePath := "./assets/images/Ninja/"
	AttackImage, _, err := ebitenutil.NewImageFromFile(basePath + "Attack.png")
	if err != nil {
		log.Fatal(err)
	}
	DeadImage, _, err := ebitenutil.NewImageFromFile(basePath + "Dead.png")
	if err != nil {
		log.Fatal(err)
	}
	IdleImage, _, err := ebitenutil.NewImageFromFile(basePath + "Idle.png")
	if err != nil {
		log.Fatal(err)
	}
	ItemImage, _, err := ebitenutil.NewImageFromFile(basePath + "Item.png")
	if err != nil {
		log.Fatal(err)
	}
	JumpImage, _, err := ebitenutil.NewImageFromFile(basePath + "Jump.png")
	if err != nil {
		log.Fatal(err)
	}
	Special1Image, _, err := ebitenutil.NewImageFromFile(basePath + "Special1.png")
	if err != nil {
		log.Fatal(err)
	}
	Special2Image, _, err := ebitenutil.NewImageFromFile(basePath + "Special2.png")
	if err != nil {
		log.Fatal(err)
	}
	WalkImage, _, err := ebitenutil.NewImageFromFile(basePath + "Walk.png")
	if err != nil {
		log.Fatal(err)
	}
	return &Ninja{
		AttackImage:   AttackImage,
		DeadImage:     DeadImage,
		IdleImage:     IdleImage,
		ItemImage:     ItemImage,
		JumpImage:     JumpImage,
		Special1Image: Special1Image,
		Special2Image: Special2Image,
		WalkImage:     WalkImage,
	}
}
func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	ninja := NewNinja()

	if err := ebiten.RunGame(&Game{Ninja: ninja}); err != nil {
		log.Fatal(err)
	}
}
