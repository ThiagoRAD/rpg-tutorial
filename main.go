package main

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Sprite struct {
	image  *ebiten.Image
	width  int
	height int
	X, Y   float64
}

type Enemy struct {
	*Sprite
	FollowsPlayer bool
}
type Potion struct {
	*Sprite
}
type Player struct {
	*Sprite
}
type Game struct {
	player  *Player
	enemies []*Enemy
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.player.Y -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.player.Y += 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.player.X -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.player.X += 2
	}
	for _, enemy := range g.enemies {
		if enemy.FollowsPlayer {
			if g.player.X < enemy.X {
				enemy.X -= 1
			} else if g.player.X > enemy.X {
				enemy.X += 1
			}
			if g.player.Y < enemy.Y {
				enemy.Y -= 1
			} else if g.player.Y > enemy.Y {
				enemy.Y += 1
			}
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{120, 180, 255, 255})

	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(g.player.X, g.player.Y)

	screen.DrawImage(g.player.image.SubImage(image.Rect(0, 0, g.player.image.Bounds().Dx()/g.player.width, g.player.image.Bounds().Dy()/g.player.height)).(*ebiten.Image), &opts)
	for _, enemy := range g.enemies {
		opts := ebiten.DrawImageOptions{}
		opts.GeoM.Translate(enemy.X, enemy.Y)
		screen.DrawImage(enemy.image.SubImage(image.Rect(0, 0, enemy.image.Bounds().Dx()/enemy.width, enemy.image.Bounds().Dy()/enemy.height)).(*ebiten.Image), &opts)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ebiten.WindowSize()
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	ninjaImage, _, err := ebitenutil.NewImageFromFile("./assets/images/Ninja/Walk.png")
	skeletonImage, _, err := ebitenutil.NewImageFromFile("./assets/images/Skeleton/Walk.png")

	if err != nil {
		log.Fatal(err)
	}

	player := &Sprite{
		image:  ninjaImage,
		width:  4,
		height: 4,
		X:      100,
		Y:      100,
	}

	enemy := &Enemy{
		Sprite: &Sprite{
			image:  skeletonImage,
			width:  4,
			height: 4,
			X:      200,
			Y:      200,
		},
		FollowsPlayer: true,
	}

	game := &Game{
		player:  &Player{Sprite: player},
		enemies: []*Enemy{enemy},
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
