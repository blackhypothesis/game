package main

import (
	"image"
	"io/fs"
	"math"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

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

func loadImages(path string) []*ebiten.Image {
	matches, err := fs.Glob(assets, path)
	if err != nil {
		panic(err)
	}

	images := make([]*ebiten.Image, len(matches))
	for i, match := range matches {
		images[i] = loadImage(match)
	}

	return images
}

func loadFont(name string) font.Face {
	f, err := assets.ReadFile(name)
	if err != nil {
		panic(err)
	}

	tt, err := opentype.Parse(f)
	if err != nil {
		panic(err)
	}

	face, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    48,
		DPI:     72,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		panic(err)
	}

	return face
}

func randomFloat(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func checkCollision(obj1, obj2 GameObject) bool {
	if (math.Abs(obj1.Position.X-obj2.Position.X) < obj2.HalfSize.X) && (math.Abs(obj1.Position.Y-obj2.Position.Y) < obj2.HalfSize.Y) {
		return true
	}
	return false
}
