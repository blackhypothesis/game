package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Vector struct {
	X float64
	Y float64
}

type GameObject struct {
	Sprite        *ebiten.Image
	Position      Vector
	Scale         Vector
	Angle         float64
	Speed         Vector
	RotationSpeed float64
	HalfSize      Vector
	MsgQueue      *MessageQueue
	CreatedAt     time.Time
}

func NewGameObject(sprite string, position Vector, scale Vector, angle float64, speed Vector, rotationSpeed float64, msgQueue *MessageQueue) *GameObject {
	spriteImage := loadImage(sprite)

	gameObject := new(GameObject)
	gameObject.Sprite = spriteImage
	gameObject.Position = position
	gameObject.Scale = scale
	gameObject.Angle = angle
	gameObject.Speed = speed
	gameObject.RotationSpeed = rotationSpeed
	gameObject.HalfSize = Vector{float64(spriteImage.Bounds().Dx()) / 2, float64(spriteImage.Bounds().Dy()) / 2}
	gameObject.MsgQueue = msgQueue
	gameObject.CreatedAt = time.Now()
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

	g.Angle += g.RotationSpeed
}

func (g *GameObject) Update() {
	g.Move()
}

func (g *GameObject) Draw(screen *ebiten.Image) {
	opts := new(ebiten.DrawImageOptions)
	opts.GeoM.Translate(-g.HalfSize.X, -g.HalfSize.Y)
	opts.GeoM.Scale(g.Scale.X, g.Scale.Y)
	opts.GeoM.Rotate(g.Angle)
	opts.GeoM.Translate(g.Position.X, g.Position.Y)
	screen.DrawImage(g.Sprite, opts)
}
