package main

import (
	"embed"
	"image"
	_ "image/png"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	SCREENWIDTH  = 3200
	SCREENHEIGHT = 2000
)

type Game struct {
	Player *GameObject
}

type Vector struct {
	X float64
	Y float64
}

type GameObject struct {
	Sprite   *ebiten.Image
	Position Vector
	Rotation float64
	Speed    float64
	HalfSize Vector
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

func NewGameObject(assetPath string, position Vector, rotation float64, speed float64) *GameObject {
	spriteImage := loadImage(assetPath)

	gameObject := new(GameObject)
	gameObject.Sprite = spriteImage
	gameObject.Position = position
	gameObject.Rotation = rotation
	gameObject.Speed = speed
	gameObject.HalfSize = Vector{float64(spriteImage.Bounds().Dx()) / 2, float64(spriteImage.Bounds().Dy()) / 2}
	return gameObject
}

func (g *GameObject) Move() {
	g.Position.X += math.Sin(g.Rotation) * g.Speed
	g.Position.Y += -math.Cos(g.Rotation) * g.Speed
}

func (g *GameObject) Update() {
	g.Move()

}

func (g *GameObject) Draw(screen *ebiten.Image) {
	opts := new(ebiten.DrawImageOptions)
	opts.GeoM.Translate(-g.HalfSize.X, -g.HalfSize.Y)
	opts.GeoM.Rotate(g.Rotation)
	opts.GeoM.Translate(g.Position.X, g.Position.Y)
	screen.DrawImage(g.Sprite, opts)
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.Player.Speed -= 0.1
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.Player.Speed += 0.1
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.Player.Rotation += math.Pi / 90
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.Player.Rotation -= math.Pi / 90
	}

	g.Player.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREENWIDTH, SCREENHEIGHT
}

func main() {
	game := &Game{
		Player: NewGameObject("assets/PNG/playerShip1_blue.png", Vector{X: 100, Y: 100}, 180*math.Pi/180, 0.5),
	}
	game.Player.Position = Vector{
		X: 100,
		Y: 100,
	}

	ebiten.SetWindowSize(3200, 2000)
	ebiten.SetWindowTitle("My First Game")
	err := ebiten.RunGame(game)
	if err != nil {
		panic(err)
	}
}
