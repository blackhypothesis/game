package main

import (
	"math"
	"math/rand/v2"
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

func NewBullet(player *GameObject) *GameObject {

	return NewGameObject("assets/PNG/Lasers/laserGreen05.png",
		Vector{player.Position.X + math.Sin(player.Angle)*player.HalfSize.X, player.Position.Y - math.Cos(player.Angle)*player.HalfSize.Y},
		Vector{1, 1},
		player.Angle,
		Vector{math.Sin(player.Angle) * 15, -math.Cos(player.Angle) * 15},
		0,
		NewMessageQueue())
}

func NewMeteor() *GameObject {
	meteorScale := randomFloat(0.5, 2)
	return NewGameObject("assets/PNG/Meteors/meteorBrown_big1.png",
		Vector{500, 500},
		Vector{meteorScale, meteorScale},
		rand.Float64()*2*math.Pi,
		Vector{randomFloat(-3, 3), randomFloat(-3, 3)},
		randomFloat(-0.05, 0.05),
		NewMessageQueue())
}

func NewMeteorDebris(position Vector, scale Vector) *GameObject {
	return NewGameObject("assets/PNG/Meteors/meteorBrown_big1.png",
		position,
		scale,
		rand.Float64()*2*math.Pi,
		Vector{randomFloat(-3, 3), randomFloat(-3, 3)},
		randomFloat(-0.05, 0.05),
		NewMessageQueue())
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
