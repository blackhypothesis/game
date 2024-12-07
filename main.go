package main

import (
	"embed"
	_ "image/png"
	"math"
	"math/rand/v2"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	SCREENWIDTH  = 2000
	SCREENHEIGHT = 1500
)

type Game struct {
	Player      *GameObject
	Bullets     []GameObject
	Meteors     []GameObject
	BulletTimer *Timer
	MeteorTimer *Timer
}

//go:embed assets/*
var assets embed.FS

// var playerSprite = loadImage("assets/PNG/playerShip1_blue.png")
// var bulletSprite = loadImage("assets/PNG/Lasers/laserGreen05.png")
// var meteorSprites = loadImages("assets/PNG/Meteors/*")

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
		if g.BulletTimer.IsReady() {
			g.BulletTimer.Reset()
			g.Bullets = append(g.Bullets, *NewGameObject("assets/PNG/Lasers/laserGreen05.png",
				Vector{g.Player.Position.X + math.Sin(g.Player.Rotation)*g.Player.HalfSize.X, g.Player.Position.Y - math.Cos(g.Player.Rotation)*g.Player.HalfSize.Y},
				g.Player.Rotation,
				Vector{math.Sin(g.Player.Rotation) * 5, -math.Cos(g.Player.Rotation) * 5},
				NewMessageQueue()))
		}
	}

	g.Player.Update()
	for i := range g.Bullets {
		g.Bullets[i].Update()
	}

	for i := range g.Bullets {
		if time.Since(g.Bullets[i].CreatedAt) > 2*time.Second {
			g.Bullets = append(g.Bullets[:i], g.Bullets[i+1:]...)
			break
		}
	}

	for i := range g.Bullets {
		g.Bullets[i].Update()
	}

	for i := range g.Meteors {
		g.Meteors[i].Update()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.BulletTimer.Update()
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

	for i := range g.Bullets {
		g.Bullets[i].Draw(screen)
	}

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
			Vector{X: SCREENWIDTH / 2, Y: SCREENHEIGHT / 2},
			rand.Float64()*2*math.Pi,
			Vector{0, 0},
			mq),
		BulletTimer: NewTimer(80 * time.Millisecond),
		MeteorTimer: NewTimer(10000 * time.Millisecond),
	}

	ebiten.SetWindowSize(SCREENWIDTH, SCREENHEIGHT)
	ebiten.SetWindowTitle("My First Game")
	err := ebiten.RunGame(game)
	if err != nil {
		panic(err)
	}
}
