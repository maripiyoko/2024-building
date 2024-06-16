package world

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"github.com/maripiyoko/2024-building/resources/images"
)

const (
	screenWidth  = 320
	screenHeight = 240

	frameOX           = 0
	frameOY           = 0
	frameWidth        = 100
	frameHeight       = 100
	frameCountWalking = 8
	frameCountIdle    = 6
)

var (
	archerWalkImage *ebiten.Image
	archerIdleImage *ebiten.Image
	backgroundImage *ebiten.Image
)

type Game struct {
	count          int
	walkCount      int
	isWalking      bool
	walkingForward bool

	playerPositionX int
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(images.ArcherWalk_png))
	if err != nil {
		log.Fatal(err)
	}
	archerWalkImage = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(images.ArcherIdle_png))
	if err != nil {
		log.Fatal(err)
	}
	archerIdleImage = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(images.Background_png))
	if err != nil {
		log.Fatal(err)
	}
	backgroundImage = ebiten.NewImageFromImage(img)
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		g.isWalking = true
		g.walkingForward = true
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyArrowRight) {
		g.isWalking = false
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		g.isWalking = true
		g.walkingForward = false
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyArrowLeft) {
		g.isWalking = false
		g.walkingForward = true
	}

	if g.isWalking {
		g.walkCount++
		if g.walkingForward {
			if g.playerPositionX < 290 {
				g.playerPositionX++
			}
		} else {
			if g.playerPositionX > 0 {
				g.playerPositionX--
			}
		}
	} else {
		g.count++
		g.walkCount = 0
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	drawBackground(g, screen)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(screenWidth)/2, float64(30))
	op.GeoM.Translate(screenWidth/2, screenHeight/2)

	var i = 0
	var targetImage *ebiten.Image
	if g.isWalking {
		i = (g.walkCount / 5) % frameCountWalking
		targetImage = archerWalkImage
	} else {
		i = (g.count / 5) % frameCountIdle
		targetImage = archerIdleImage
	}
	sx, sy := frameOX+i*frameWidth, frameOY
	rect := image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)
	screen.DrawImage(targetImage.SubImage(rect).(*ebiten.Image), op)

	debugPrint(g, screen)
}

func drawBackground(g *Game, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	sx := g.playerPositionX
	op.GeoM.Translate(-float64(sx), -float64(620)/2)
	op.GeoM.Scale(0.5, 0.5)
	screen.DrawImage(backgroundImage, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func NewGame() (*Game, error) {
	game := &Game{}
	return game, nil
}

func debugPrint(g *Game, screen *ebiten.Image) {
	var playserStatus = ""
	if g.isWalking {
		playserStatus = "walking"
	} else {
		playserStatus = "idle"
	}
	var msg = fmt.Sprintf("Player x,y=%d,%d (%s) ", g.playerPositionX, 0, playserStatus)
	ebitenutil.DebugPrint(screen, msg)
}
