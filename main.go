package main

import (
	"image"
	"image/color"
	"log"
	"rpg-tutorial/entities"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var screenWidth, screenHeight int = 640, 480
var imageCache = make(map[string]*ebiten.Image)

type Game struct {
	player       *entities.Player
	enemies      []*entities.Enemy
	items        []*entities.Potion
	tilemap      *Tilemap
	tilemapImage *ebiten.Image
	camera       *Camera
}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return screenWidth, screenHeight
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
	g.camera.FollowTarget(g.player.X, g.player.Y, screenWidth, screenHeight)
	return nil
}

func (g *Game) RenderDrawable(screen *ebiten.Image, drawable *entities.Drawable) {
	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(drawable.X, drawable.Y)
	screen.DrawImage(drawable.Sprite.Image.SubImage(image.Rect(0, 0, drawable.Sprite.Image.Bounds().Dx()/drawable.Width, drawable.Sprite.Image.Bounds().Dy()/drawable.Height)).(*ebiten.Image), &opts)
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
			opts.GeoM.Translate(g.camera.X, g.camera.Y)
			screen.DrawImage(g.tilemapImage.SubImage(srcRect).(*ebiten.Image), &opts)

		}
	}
}
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{120, 180, 255, 255})

	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(g.player.X, g.player.Y)
	opts.GeoM.Translate(g.camera.X, g.camera.Y)

	g.DrawLayers(screen)
	g.RenderDrawable(screen, g.player.Drawable)

	for _, enemy := range g.enemies {
		opts := ebiten.DrawImageOptions{}
		opts.GeoM.Translate(enemy.X, enemy.Y)
		opts.GeoM.Translate(g.camera.X, g.camera.Y)
		g.RenderDrawable(screen, enemy.Drawable)
	}

	for _, item := range g.items {
		opts := ebiten.DrawImageOptions{}
		opts.GeoM.Translate(item.X, item.Y)
		opts.GeoM.Translate(g.camera.X, g.camera.Y)
		g.RenderDrawable(screen, item.Drawable)
	}
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
	item := &entities.Potion{
		Drawable: &entities.Drawable{
			Sprite: &entities.Sprite{
				Image:  potionImage,
				Width:  1,
				Height: 1,
			},
			Position: &entities.Position{
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
	enemy := &entities.Enemy{
		Drawable: &entities.Drawable{
			Sprite: &entities.Sprite{
				Image:  skeletonImage,
				Width:  4,
				Height: 4,
			},
			Position: &entities.Position{
				X: X,
				Y: Y,
			},
		},
		FollowsPlayer: true,
	}
	g.enemies = append(g.enemies, enemy)
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	tilemap, err := LoadTilemap("./assets/maps/spawn.json")
	if err != nil {
		log.Fatal(err)
	}
	ninjaImage := getImageFromFile("./assets/images/Ninja/Walk.png")
	tilemapImage := getImageFromFile("./assets/maps/TilesetFloor.png")

	player := &entities.Drawable{
		Sprite: &entities.Sprite{
			Image:  ninjaImage,
			Width:  4,
			Height: 4,
		},
		Position: &entities.Position{
			X: 100,
			Y: 100,
		},
	}
	game := &Game{
		player:       &entities.Player{Drawable: player, Health: 100},
		tilemap:      tilemap,
		tilemapImage: tilemapImage,
		camera:       NewCamera(),
	}
	game.AddSkeleton(200, 200)
	game.AddSkeleton(350, 400)
	game.AddSkeleton(-100, -400)
	game.AddPotion(300, 300)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
