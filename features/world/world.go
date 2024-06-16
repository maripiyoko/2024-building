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

	baseLineHeight = 34
	startLinePosX  = 30

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
	bgImage         *ebiten.Image

	bgWidth, bgHeight int
)

type Game struct {
	count int

	player   *Player
	viewPort *ViewPort

	reachEnd bool
}

type Player struct {
	idleCount int
	walkCount int

	isWalking        bool
	walkingBackwards bool

	positionX int
}

type ViewPort struct {
	positionX int
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
	bgImage = ebiten.NewImageFromImage(img)

	s := bgImage.Bounds().Size()
	bgWidth = s.X
	bgHeight = s.Y
	fmt.Printf("bg width %d, height %d", bgWidth, bgHeight)
}

func (g *Game) Update() error {
	g.count++

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		g.player.isWalking = true
		g.player.walkingBackwards = false
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyArrowRight) {
		g.player.isWalking = false
		g.player.walkingBackwards = false
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		g.player.isWalking = true
		g.player.walkingBackwards = true
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyArrowLeft) {
		g.player.isWalking = false
		g.player.walkingBackwards = false
	}

	if g.player.isWalking {
		g.player.walkCount++
		if g.player.walkingBackwards {
			if g.player.positionX > -40 {
				g.player.positionX--
				g.reachEnd = false
			} else { // 画面左端に到達
				g.reachEnd = true
			}
		} else {
			if g.player.positionX < 438 {
				g.player.positionX++
				g.reachEnd = false
			} else { // 画面右端に到達
				g.reachEnd = true
			}
		}
	} else {
		g.player.idleCount++
		g.player.walkCount = 0
	}

	// 画面がスクロール可能かチェック
	if g.player.positionX > 0 && g.player.positionX < 140 {
		g.viewPort.positionX = g.player.positionX
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	drawBackground(g, screen)

	op := &ebiten.DrawImageOptions{}
	if g.player.walkingBackwards { // 左右どちら向きか判定
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(frameWidth, 0)
	}
	// キャラの位置を初期値へ移動
	op.GeoM.Translate(-float64(frameWidth)/2, -float64(frameHeight)/2)
	op.GeoM.Translate(startLinePosX, screenHeight-baseLineHeight)
	// 画面がスクロール不可の場合、キャラのみ移動
	diffX := g.player.positionX - g.viewPort.positionX
	op.GeoM.Translate(float64(diffX), 0)

	var i = 0
	var targetImage *ebiten.Image
	if g.player.isWalking {
		i = (g.player.walkCount / 5) % frameCountWalking
		targetImage = archerWalkImage
	} else {
		i = (g.player.idleCount / 5) % frameCountIdle
		targetImage = archerIdleImage
	}
	sx, sy := frameOX+i*frameWidth, frameOY
	rect := image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)
	screen.DrawImage(targetImage.SubImage(rect).(*ebiten.Image), op)

	debugPrint(g, screen)
}

func drawBackground(g *Game, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	sx := g.viewPort.positionX
	op.GeoM.Scale(0.5, 0.5)
	op.GeoM.Translate(-float64(sx), -float64(154))

	screen.DrawImage(bgImage, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func NewGame() (*Game, error) {
	game := &Game{
		player:   &Player{},
		viewPort: &ViewPort{},
	}
	return game, nil
}

func debugPrint(g *Game, screen *ebiten.Image) {
	/*var playserStatus = ""
	if g.player.isWalking {
		playserStatus = "walking"
	} else {
		playserStatus = "idle"
	}*/

	var end = ""
	if g.reachEnd {
		if g.player.positionX > 0 {
			end = "forwardEnd"
		} else {
			end = "backwardEnd"
		}
	}

	var msg = fmt.Sprintf("Player x,y=%d,%d viewPortX=%d end=%s", g.player.positionX, 0, g.viewPort.positionX, end)
	ebitenutil.DebugPrint(screen, msg)
}
