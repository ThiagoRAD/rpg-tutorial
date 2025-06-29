package main

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var imageCache = make(map[string]*ebiten.Image)

type Sprite struct {
	image  *ebiten.Image
	width  int
	height int
}

type Drawable struct {
	*Sprite
	*Position
}
type Position struct {
	X, Y float64
}
type Enemy struct {
	*Drawable
	FollowsPlayer bool
}
type Player struct {
	*Drawable
	Health float64
}
type Potion struct {
	*Drawable
	HealAmount float64
}
type Game struct {
	player       *Player
	enemies      []*Enemy
	items        []*Potion
	tilemap      *Tilemap
	tilemapImage *ebiten.Image
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

func (g *Game) RenderDrawable(screen *ebiten.Image, drawable *Drawable) {
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(drawable.X, drawable.Y)
	screen.DrawImage(drawable.image.SubImage(image.Rect(0, 0, drawable.image.Bounds().Dx()/drawable.width, drawable.image.Bounds().Dy()/drawable.height)).(*ebiten.Image), &opts)
}

func (g *Game) DrawLayers(screen *ebiten.Image) {
	for _, layer := range g.tilemap.Tiles {
		for index, id := range layer.Data {
			x := index % layer.Width
			y := index / layer.Width
			x *= 16 // the width of each tile in the map
			y *= 16
			srcX := (id - 1) % 22 // the tileimage has 22 squares
			srcY := (id - 1) / 22
			srcX *= 16
			srcY *= 16

			srcRect := image.Rect(srcX, srcY, srcX+16, srcY+16)
			opts := ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(x), float64(y))
			screen.DrawImage(g.tilemapImage.SubImage(srcRect).(*ebiten.Image), &opts)

		}
	}
}
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{120, 180, 255, 255})

	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(g.player.X, g.player.Y)

	g.DrawLayers(screen)
	g.RenderDrawable(screen, g.player.Drawable)

	for _, enemy := range g.enemies {
		opts := ebiten.DrawImageOptions{}
		opts.GeoM.Translate(enemy.X, enemy.Y)
		g.RenderDrawable(screen, enemy.Drawable)
	}

	for _, item := range g.items {
		opts := ebiten.DrawImageOptions{}
		opts.GeoM.Translate(item.X, item.Y)
		g.RenderDrawable(screen, item.Drawable)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ebiten.WindowSize()
}

func getImageFromFile(path string) *ebiten.Image {
	imageCached, ok := imageCache[path]
	if ok {
		return imageCached
	}
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	imageCache[path] = img
	return img
}

func (g *Game) AddPotion(X, Y float64) {
	potionImage := getImageFromFile("./assets/images/Potion/LifePot.png")
	item := &Potion{
		Drawable: &Drawable{
			Sprite: &Sprite{
				image:  potionImage,
				width:  1,
				height: 1,
			},
			Position: &Position{
				X: X,
				Y: Y,
			},
		},
		HealAmount: 20,
	}
	g.items = append(g.items, item)
}

func (g *Game) AddSkeleton(X, Y float64) {
	skeletonImage := getImageFromFile("./assets/images/Skeleton/Walk.png")
	enemy := &Enemy{
		Drawable: &Drawable{
			Sprite: &Sprite{
				image:  skeletonImage,
				width:  4,
				height: 4,
			},
			Position: &Position{
				X: X,
				Y: Y,
			},
		},
		FollowsPlayer: true,
	}
	g.enemies = append(g.enemies, enemy)
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	tilemap, err := LoadTilemap("./assets/maps/map.json")
	if err != nil {
		log.Fatal(err)
	}
	ninjaImage := getImageFromFile("./assets/images/Ninja/Walk.png")
	tilemapImage := getImageFromFile("./assets/maps/TilesetFloor.png")

	player := &Drawable{
		Sprite: &Sprite{
			image:  ninjaImage,
			width:  4,
			height: 4,
		},
		Position: &Position{
			X: 100,
			Y: 100,
		},
	}
	game := &Game{
		player:       &Player{Drawable: player, Health: 100},
		tilemap:      tilemap,
		tilemapImage: tilemapImage,
	}
	game.AddSkeleton(200, 200)
	game.AddSkeleton(350, 400)
	game.AddSkeleton(-100, -400)
	game.AddPotion(300, 300)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
