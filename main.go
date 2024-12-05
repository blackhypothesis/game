package main

import (
	"embed"
	"image"
	_ "image/png"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const degree = math.Pi / 180

type Game struct {
	Player GameObject
}

type Vector struct {
	X float64
	Y float64
}

type GameObject struct {
	Position Vector
	Rotation float64
	Speed    Vector
	Sprite   *ebiten.Image
	HalfSize Vector
	Options  *ebiten.DrawImageOptions
}

type GameObjectEdit interface {
	RotateRelative(angle float64)
	Move()
}

//go:embed assets/*
var assets embed.FS
var PlayerSprite = loadImage("assets/PNG/playerShip1_blue.png")

func loadImage(name string) *ebiten.Image {
	f, err := assets.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}

func NewGameObject(assetPath string) *GameObject {
	spriteImage := loadImage(assetPath)

	return &GameObject{
		Position: Vector{0, 0},
		Rotation: 0,
		Speed:    Vector{0, 0},
		Sprite:   spriteImage,
		HalfSize: Vector{float64(spriteImage.Bounds().Dx()) / 2, float64(spriteImage.Bounds().Dy()) / 2},
		Options:  new(ebiten.DrawImageOptions),
	}
}

func (g *GameObject) RotateRelative(angle float64) {
	rad := angle * degree
	g.Rotation = g.Rotation + rad
	opts := new(ebiten.DrawImageOptions)
	opts.GeoM.Translate(-g.HalfSize.X, -g.HalfSize.y)
	opts.GeoM.Rotate(g.Rotation)
	opts.GeoM.Translate(g.Position.X, g.Position.X)
}

func (g *GameObject) Move() {

}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	// op.GeoM.Translate(50, 70)
	// op.GeoM.Rotate(-45.0 * math.Pi / 180.0)
	// op.GeoM.Scale(1, -1)

	op = RotateSprite(PlayerSprite, 68)

	screen.DrawImage(PlayerSprite, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	g := &Game{}

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
