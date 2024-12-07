package main

import (
	"embed"
	"fmt"
	_ "image/png"
	"math"
	"math/rand/v2"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	SCREENWIDTH  = 3200
	SCREENHEIGHT = 2000
)

type Game struct {
	Player      *GameObject
	Bullets     []GameObject
	Meteors     []GameObject
	AttackTimer *Timer
	MeteorTimer *Timer
}

//go:embed assets/*
var assets embed.FS
var PlayerSprite = loadImage("assets/PNG/playerShip1_blue.png")
var MeteorSprites = loadImages("assets/PNG/Meteors/*")

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.Player.Speed.X += -math.Sin(g.Player.Rotation) / 10
		g.Player.Speed.Y += math.Cos(g.Player.Rotation) / 10
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.Player.Speed.X += math.Sin(g.Player.Rotation) / 10
		g.Player.Speed.Y += -math.Cos(g.Player.Rotation) / 10
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.Player.Rotation += math.Pi / 90
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.Player.Rotation -= math.Pi / 90
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		fmt.Println("Shoot, ...")
	}

	g.Player.Update()
	for i := range g.Bullets {
		g.Bullets[i].Update()
	}

	for i := range g.Meteors {
		g.Meteors[i].Update()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.MeteorTimer.Update()
	if g.MeteorTimer.IsReady() {
		g.MeteorTimer.Reset()
		g.Meteors = append(g.Meteors, *NewGameObject("assets/PNG/Meteors/meteorBrown_big1.png",
			Vector{X: 500, Y: 500},
			rand.Float64()*2*math.Pi,
			Vector{randomFloat(-3, 3), randomFloat(-3, 3)},
			NewMessageQueue()))
	}
	g.Player.Draw(screen)

	for i := range g.Meteors {
		g.Meteors[i].Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREENWIDTH, SCREENHEIGHT
}

func main() {

	mq := NewMessageQueue()

	game := &Game{
		Player: NewGameObject("assets/PNG/playerShip1_blue.png",
			Vector{X: 100, Y: 100},
			rand.Float64()*2*math.Pi,
			Vector{0, 0},
			mq),
		MeteorTimer: NewTimer(100 * time.Millisecond),
	}

	ebiten.SetWindowSize(3200, 2000)
	ebiten.SetWindowTitle("My First Game")
	err := ebiten.RunGame(game)
	if err != nil {
		panic(err)
	}
}
