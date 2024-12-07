package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Vector struct {
	X float64
	Y float64
}

type GameObject struct {
	Sprite   *ebiten.Image
	Position Vector
	Rotation float64
	Speed    Vector
	HalfSize Vector
	MsgQueue *MessageQueue
}

func NewGameObject(assetPath string, position Vector, rotation float64, speed Vector, msgQueue *MessageQueue) *GameObject {
	spriteImage := loadImage(assetPath)

	gameObject := new(GameObject)
	gameObject.Sprite = spriteImage
	gameObject.Position = position
	gameObject.Rotation = rotation
	gameObject.Speed = speed
	gameObject.HalfSize = Vector{float64(spriteImage.Bounds().Dx()) / 2, float64(spriteImage.Bounds().Dy()) / 2}
	gameObject.MsgQueue = msgQueue
	return gameObject
}

func (g *GameObject) Move() {
	g.Position.X += g.Speed.X
	g.Position.Y += g.Speed.Y

	if g.Position.X > SCREENWIDTH {
		g.Position.X -= SCREENWIDTH
	}

	if g.Position.X < 0 {
		g.Position.X += SCREENWIDTH
	}

	if g.Position.Y > SCREENHEIGHT {
		g.Position.Y -= SCREENHEIGHT
	}

	if g.Position.Y < 0 {
		g.Position.Y += SCREENHEIGHT
	}
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
