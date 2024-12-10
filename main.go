package main

import (
	"embed"
	_ "image/png"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	SCREENWIDTH  = 1500
	SCREENHEIGHT = 800
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
		g.Player.Speed.X += -math.Sin(g.Player.Angle) / 10
		g.Player.Speed.Y += math.Cos(g.Player.Angle) / 10
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.Player.Speed.X += math.Sin(g.Player.Angle) / 10
		g.Player.Speed.Y += -math.Cos(g.Player.Angle) / 10
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.Player.Angle -= math.Pi / 90
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.Player.Angle += math.Pi / 90
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.BulletTimer.Update()
		if g.BulletTimer.IsReady() {
			g.BulletTimer.Reset()
			g.Bullets = append(g.Bullets, *NewBullet(g.Player))
		}
	}

	g.MeteorTimer.Update()
	if g.MeteorTimer.IsReady() {
		g.MeteorTimer.Reset()
		g.Meteors = append(g.Meteors, *NewMeteor())
	}

	for i := range g.Bullets {
		if time.Since(g.Bullets[i].CreatedAt) > 2*time.Second {
			g.Bullets = append(g.Bullets[:i], g.Bullets[i+1:]...)
			break
		}
	}

	for i := range g.Meteors {
		g.Meteors[i].Update()
	}

	g.Player.Update()

	for i := range g.Bullets {
		g.Bullets[i].Update()
	}

out:
	for b := range g.Bullets {
		for m := range g.Meteors {
			if checkCollision(g.Bullets[b], g.Meteors[m]) {
				if g.Meteors[m].Scale.X > 1 {
					for i := 2; i < 6; i++ {
						g.Meteors = append(g.Meteors,
							*NewMeteorDebris(g.Meteors[m].Position,
								Vector{g.Meteors[m].Scale.X / float64(i), g.Meteors[m].Scale.Y / float64(i)}))
					}
				}
				g.Meteors = append(g.Meteors[:m], g.Meteors[m+1:]...)
				g.Bullets = append(g.Bullets[:b], g.Bullets[b+1:]...)
				break out
			}
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for i := range g.Bullets {
		g.Bullets[i].Draw(screen)
	}

	for i := range g.Meteors {
		g.Meteors[i].Draw(screen)
	}

	g.Player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREENWIDTH, SCREENHEIGHT
}

func main() {

	mq := NewMessageQueue()

	game := &Game{
		Player: NewGameObject("assets/PNG/playerShip1_blue.png",
			Vector{SCREENWIDTH / 2, SCREENHEIGHT / 2},
			Vector{1, 1},
			0,
			Vector{0, 0},
			0,
			mq),
		BulletTimer: NewTimer(250 * time.Millisecond),
		MeteorTimer: NewTimer(4000 * time.Millisecond),
	}

	ebiten.SetWindowSize(SCREENWIDTH, SCREENHEIGHT)
	ebiten.SetWindowTitle("My First Game")
	err := ebiten.RunGame(game)
	if err != nil {
		panic(err)
	}
}
