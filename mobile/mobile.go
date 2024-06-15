package mobile

import (
	"github.com/hajimehoshi/ebiten/v2/mobile"

	"github.com/maripiyoko/2024-building/features/world"
)

func init() {
	game, err := world.NewGame()
	if err != nil {
		panic(err)
	}

	mobile.SetGame(game)
}

func Dummy() {}
